package fslint

import (
	"io/ioutil"
	"log"
	"reflect"
	"strings"

	"github.com/fatih/color"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
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
	Exclude *[]string  `json:"exclude"`
	Include []Selector `json:"include" validate:"required,dive,required"`
}

type Selector struct {
	File    string   `json:"file" validate:"required_if_field_empty=Folder"`
	Folder  string   `json:"folder" validate:"required_if_field_empty=File"`
	Pattern Mode     `json:"pattern" validate:"required,oneof=CamelCase camelCase Kebab-Kebab kebab-kebab Snake_Case snake_case dot.dot Dot.Dot"`
	Level   Level    `json:"level" validate:"oneof=warn error,required"`
	Ignore  []string `json:"ignore"`
}

func requiredIfFieldNotEmpty(field validator.FieldLevel) bool {
	val := field.Field().String()

	if field.Param() != "" {
		targetVal := field.Parent().FieldByName(field.Param()).String()

		if val != "" && targetVal != "" {
			return false
		}

		if val == "" && targetVal == "" {
			return false
		}
	}

	return true
}

var (
	validate = validator.New()
	trans    ut.Translator
)

func init() {
	z := en.New()
	uni := ut.New(z, z)
	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ = uni.GetTranslator("zh")
	if err := enTranslations.RegisterDefaultTranslations(validate, trans); err != nil {
		log.Fatalln(err)
	}

	// register function to get tag name from json tags.
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	if err := validate.RegisterValidation("required_if_field_empty", requiredIfFieldNotEmpty); err != nil {
		panic(err)
	}

	if err := validate.RegisterTranslation("required_if_field_empty", trans, func(ut ut.Translator) error {
		return ut.Add("required_if_field_empty", "{0} is required when '{1}' is empty!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required_if_field_empty", fe.StructNamespace(), strings.ToLower(fe.Param()))
		return t
	}); err != nil {
		panic(err)
	}

	if err := validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is required!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.StructNamespace())
		return t
	}); err != nil {
		panic(err)
	}

	if err := validate.RegisterTranslation("oneof", trans, func(ut ut.Translator) error {
		return ut.Add("oneof", "{0} must be one of [{1}]", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("oneof", fe.StructNamespace(), fe.Param())
		return t
	}); err != nil {
		panic(err)
	}
}

func NewConfig(content []byte) (*Config, error) {
	var (
		config = Config{}
		err    error
	)

	if err = json5.Unmarshal(content, &config); err != nil {
		return nil, errors.WithStack(err)
	}

	err = validate.Struct(config)

	if err != nil {
		// translate all error at once
		errs := err.(validator.ValidationErrors)

		errorsMap := errs.Translate(trans)

		msg := []string{}

		for _, e := range errorsMap {
			msg = append(msg, color.RedString("[config]: "+e))
		}

		return nil, errors.New(strings.Join(msg, "\n"))
	}

	return &config, nil
}

func readConfig(configFilepath string) (*Config, error) {
	b, err := ioutil.ReadFile(configFilepath)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return NewConfig(b)
}
