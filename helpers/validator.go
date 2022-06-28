package helpers

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	Message string
}

func EmptyChecker(models interface{}) []*Validator {
	var errors []*Validator

	err := validator.New().Struct(models)
	if err != nil {
		for _, i := range err.(validator.ValidationErrors) {
			var e Validator
			result := strings.Split(i.StructNamespace(), ".")
			e.Message = result[1] + " tidak boleh kosong!"
			errors = append(errors, &e)
		}
	}

	return errors
}
