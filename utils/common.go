package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
)

func IsProdEnv() bool {
	return os.Getenv("ENV") == "prod"
}

func ConvertBindingErr(err error) string {
	if errs, ok := err.(validator.ValidationErrors); ok {
		if len(errs) > 0 {
			verr := errs[0]
			fieldName := verr.Field()
			tagName := verr.ActualTag()
			if tagName == "required" {
				return fmt.Sprintf("`%s` cannot be empty", strings.ToLower(fieldName))
			} else if tagName == "max" {
				return fmt.Sprintf("`%s` length cannot greater than %s", strings.ToLower(fieldName), verr.Param())
			}
		}
	}
	return ""
}
