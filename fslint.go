package fslint

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/axetroy/fslint/internal/parser"
	zglob "github.com/mattn/go-zglob"
	"github.com/pkg/errors"
	glob "github.com/ryanuber/go-glob"
)

func isGlob(pattern string) bool {
	return strings.Contains(pattern, "*")
}

func handleMatchFile(results *Results, selector Selector, exclude *[]string) error {
	var (
		isFolder bool = false
	)

	selectTarget := selector.File

	if selector.Folder != "" {
		selectTarget = selector.Folder
		isFolder = true
	}

	if selectTarget == "" {
		return errors.New("missing 'file' od 'folder' file in include block")
	}

	matchers, err := zglob.Glob(selectTarget)

	if err != nil {
		return errors.WithStack(err)
	}

loop:
	for _, file := range matchers {

		// exclude
		if exclude != nil {
			paths := strings.Split(file, string(filepath.Separator))
			for _, pattern := range *exclude {
				if isGlob(pattern) {
					if glob.Glob(pattern, file) {
						continue loop
					}
				} else {
					for _, p := range paths {
						if p == pattern {
							continue loop
						}
					}
				}
			}
		}

		// ignore for rules
		if selector.Ignore != nil {
			paths := strings.Split(file, string(filepath.Separator))
			for _, pattern := range selector.Ignore {
				if isGlob(pattern) {
					if glob.Glob(pattern, file) {
						continue loop
					}
				} else {
					for _, p := range paths {
						if p == pattern {
							continue loop
						}
					}
				}
			}
		}

		var testTarget string

		if isFolder {
			stat, err := os.Stat(file)

			if err != nil {
				return errors.WithStack(err)
			}

			if !stat.IsDir() {
				continue
			}

			splits := strings.Split(file, "/")

			testTarget = splits[len(splits)-1]
		} else {
			filenameWitoutExtension := filepath.Base(file)

			filenameWitoutExtension = strings.TrimRight(filenameWitoutExtension, filepath.Ext(filenameWitoutExtension))

			testTarget = filenameWitoutExtension
		}

		switch selector.Pattern {
		case ModePascalCase:
			if !parser.IsPascalCase(testTarget) {
				results.Append(LintResult{
					FilePath: file,
					Expect:   ModePascalCase,
					Level:    selector.Level,
				})
			}
		case ModeCamelCase:
			if !parser.IsCamelCase(testTarget) {
				results.Append(LintResult{
					FilePath: file,
					Expect:   ModeCamelCase,
					Level:    selector.Level,
				})
			}
		case ModeKebabCase:
			if !parser.IsKebabCase(testTarget) {
				results.Append(LintResult{
					FilePath: file,
					Expect:   ModeKebabCase,
					Level:    selector.Level,
				})
			}
		case ModeSnakeCase:
			if !parser.IsSnakeCase(testTarget) {
				results.Append(LintResult{
					FilePath: file,
					Expect:   ModeSnakeCase,
					Level:    selector.Level,
				})
			}
		default:
			if isRegExpStr(string(selector.Pattern)) {
				val := strings.TrimLeft(string(selector.Pattern), "/")
				val = strings.TrimRight(val, "/")
				reg, err := regexp.Compile(val)

				if err != nil {
					return errors.WithMessage(err, "invalid regexpression")
				}

				if !reg.MatchString(testTarget) {
					results.Append(LintResult{
						FilePath: file,
						Expect:   ModeRegExp,
						Level:    selector.Level,
					})
				}
			}
		}

	}

	return nil
}

func Lint(configFilepath string) (*Results, error) {
	var (
		results = NewResults()
	)

	config, err := readConfig(configFilepath)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	for _, selector := range config.Include {
		if err := handleMatchFile(results, selector, config.Exclude); err != nil {
			return nil, errors.WithStack(err)
		}
	}

	return results, nil
}
