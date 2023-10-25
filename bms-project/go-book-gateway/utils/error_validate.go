package utils

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func ParseError(errs ...error) []string {
	var out []string
	for _, err := range errs {
		switch typedError := err.(type) {
		case validator.ValidationErrors:
			// if the type is validator.ValidationErrors then it's actually an array of validator.FieldError so we'll
			// loop through each of those and convert them one by one
			for _, e := range typedError {
				out = append(out, setFieldLevel(e))
			}
		case *json.UnmarshalTypeError:
			// similarly, if the error is an unmarshalling error we'll parse it into another, more readable string format
			out = append(out, err.Error())
		default:
			out = append(out, err.Error())
		}
	}
	return out
}
func setFieldLevel(e validator.FieldError) string {

	tag := strings.Split(e.Tag(), "|")[0]

	switch tag {
	case "required":
		return fmt.Sprintf("%s is an required field", e.StructField())

	default:
		english := en.New()
		translator := ut.New(english, english)
		if translatorInstance, found := translator.GetTranslator("en"); found {
			return e.Translate(translatorInstance)
		} else {
			return fmt.Errorf("%v", e).Error()
		}

	}

}
