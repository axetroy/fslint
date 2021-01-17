package fslint

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/yosuke-furukawa/json5/encoding/json5"
)

type Mode string
type Level string

const (
	ModeBigCamelCase    Mode = "CamelCase"
	ModeLittleCamelCase Mode = "camelCase"
	ModeBigKebab        Mode = "Kebab-Kebab"
	ModeLittleKebab     Mode = "kebab-kebab"
	ModeBigSnakeCase    Mode = "Snake_Case"
	ModeLittleSnakeCase Mode = "snake_case"

	LevelWarn  Level = "warn"
	LevelError Level = "error"
)

type Config struct {
	Exclude []string   `json:"exclude"`
	Include []Selector `json:"include"`
}

type Selector struct {
	File    string `json:"file"`
	Folder  string `json:"folder"`
	Pattern Mode   `json:"pattern"`
	Level   Level  `json:"level"`
}

func readConfig(configFilepath string) (Config, error) {
	var config = Config{}

	b, err := ioutil.ReadFile(configFilepath)

	if err != nil {
		return config, errors.WithStack(err)
	}

	if err = json5.Unmarshal(b, &config); err != nil {
		return config, errors.WithStack(err)
	}

	return config, nil
}
