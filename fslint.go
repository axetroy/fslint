package fslint

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/axetroy/fslint/parser"
	zglob "github.com/mattn/go-zglob"
	"github.com/pkg/errors"
	glob "github.com/ryanuber/go-glob"
)

type Mode string

var (
	ModeBigCamelCase    Mode = "CamelCase"
	ModeLittleCamelCase Mode = "camelCase"
	ModeBigKebab        Mode = "Kebab-Kebab"
	ModeLittleKebab     Mode = "kebab-kebab"
	ModeBigSnakeCase    Mode = "Snake_Case"
	ModeLittleSnakeCase Mode = "snake_case"
)

type Config struct {
	Exclude []string        `json:"exclude"`
	Include map[string]Mode `json:"include"`
}

type LintResult struct {
	FileName string `json:"filename"`
	FilePath string `json:"filepath"`
	Expect   Mode   `json:"expect"`
	Actually Mode   `json:"actually"`
}

func readConfig(configFilepath string) (Config, error) {
	var config = Config{}

	b, err := ioutil.ReadFile(configFilepath)

	if err != nil {
		return config, errors.WithStack(err)
	}

	if err = json.Unmarshal(b, &config); err != nil {
		return config, errors.WithStack(err)
	}

	return config, nil
}

func hasResultExist(results []LintResult, filepath string) bool {
	for _, result := range results {
		if result.FilePath == filepath {
			return true
		}
	}

	return false
}

func Lint(configFilepath string) ([]LintResult, error) {
	var (
		results = make([]LintResult, 0)
	)

	config, err := readConfig(configFilepath)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	for pattern, mode := range config.Include {
		matchers, err := zglob.Glob(pattern)

		if err != nil {
			return nil, errors.WithStack(err)
		}

	loop:
		for _, file := range matchers {

			// exclude
			paths := strings.Split(file, string(filepath.Separator))
			for _, pattern := range config.Exclude {
				if strings.Contains(pattern, "*") {
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

			filenameWithExtension := filepath.Base(file)
			filenameWitoutExtension := filenameWithExtension

			extName := filepath.Ext(filenameWitoutExtension)

			if extName != "" {
				filenameWitoutExtension = strings.TrimRight(filenameWitoutExtension, extName)
			}

			switch mode {
			case ModeBigCamelCase:
				if !parser.IsCamelCase(filenameWitoutExtension, true) {
					if !hasResultExist(results, file) {
						results = append(results, LintResult{
							FileName: filenameWithExtension,
							FilePath: file,
							Expect:   ModeBigCamelCase,
						})
					}
				}
			case ModeLittleCamelCase:
				if !parser.IsCamelCase(filenameWitoutExtension, false) {
					if !hasResultExist(results, file) {
						results = append(results, LintResult{
							FileName: filenameWithExtension,
							FilePath: file,
							Expect:   ModeLittleCamelCase,
						})
					}
				}
			case ModeBigKebab:
				if !parser.IsKebab(filenameWitoutExtension, true) {
					if !hasResultExist(results, file) {
						results = append(results, LintResult{
							FileName: filenameWithExtension,
							FilePath: file,
							Expect:   ModeBigKebab,
						})
					}
				}
			case ModeLittleKebab:
				if !parser.IsKebab(filenameWitoutExtension, false) {
					if !hasResultExist(results, file) {
						results = append(results, LintResult{
							FileName: filenameWithExtension,
							FilePath: file,
							Expect:   ModeLittleKebab,
						})
					}
				}
			case ModeBigSnakeCase:
				if !parser.IsSnakeCase(filenameWitoutExtension, true) {
					if !hasResultExist(results, file) {
						results = append(results, LintResult{
							FileName: filenameWithExtension,
							FilePath: file,
							Expect:   ModeBigSnakeCase,
						})
					}
				}
			case ModeLittleSnakeCase:
				if !parser.IsSnakeCase(filenameWitoutExtension, false) {
					if !hasResultExist(results, file) {
						results = append(results, LintResult{
							FileName: filenameWithExtension,
							FilePath: file,
							Expect:   ModeLittleSnakeCase,
						})
					}
				}
			}
		}
	}

	return results, nil
}
