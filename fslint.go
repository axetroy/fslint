package fslint

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/axetroy/fslint/parser"
	"github.com/pkg/errors"
)

type Mode string

var (
	ModeBigCamelCase    Mode = "CamelCase"
	ModeLittleCamelCase Mode = "camelCase"
	ModeBigKebab        Mode = "Kebab-Kebab"
	ModeLittleKebab     Mode = "kebab-kebab"
	ModeBigSnakeCase    Mode = "Snake_Case"
	ModeLittleSnakeCase Mode = "snake_case"

	TestMap map[Mode]*regexp.Regexp = map[Mode]*regexp.Regexp{
		ModeBigCamelCase:    regexp.MustCompile(``),
		ModeLittleCamelCase: regexp.MustCompile(``),
		ModeBigKebab:        regexp.MustCompile(``),
		ModeLittleKebab:     regexp.MustCompile(``),
		ModeBigSnakeCase:    regexp.MustCompile(``),
		ModeLittleSnakeCase: regexp.MustCompile(``),
	}
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

func Lint(configFilepath string) ([]LintResult, error) {
	var (
		results = make([]LintResult, 0)
	)

	config, err := readConfig(configFilepath)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	for pattern, mode := range config.Include {
		matchers, err := filepath.Glob(pattern)

		if err != nil {
			return nil, errors.WithStack(err)
		}

		for _, file := range matchers {
			filename := filepath.Base(file)

			extName := filepath.Ext(filename)

			if extName != "" {
				filename = strings.TrimRight(filename, extName)
			}

			switch mode {
			case ModeBigCamelCase:
				if !parser.IsCamelCase(filename, true) {
					results = append(results, LintResult{
						FileName: filename,
						FilePath: file,
						Expect:   ModeBigCamelCase,
					})
				}
			case ModeLittleCamelCase:
				if !parser.IsCamelCase(filename, false) {
					results = append(results, LintResult{
						FileName: filename,
						FilePath: file,
						Expect:   ModeLittleCamelCase,
					})
				}
			case ModeBigKebab:
				if !parser.IsKebab(filename, true) {
					results = append(results, LintResult{
						FileName: filename,
						FilePath: file,
						Expect:   ModeBigKebab,
					})
				}
			case ModeLittleKebab:
				if !parser.IsKebab(filename, false) {
					results = append(results, LintResult{
						FileName: filename,
						FilePath: file,
						Expect:   ModeLittleKebab,
					})
				}
			case ModeBigSnakeCase:
			case ModeLittleSnakeCase:
			}
		}
	}

	return results, nil
}
