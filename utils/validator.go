package utils

import (
	"github.com/go-playground/validator/v10"
)

type EnumValid interface {
	Valid() bool
}

func Register(v *validator.Validate) {
	v.RegisterValidation("enum", ValidateEnum)
}

func ValidateEnum(fl validator.FieldLevel) bool {
	if enum, ok := fl.Field().Interface().(EnumValid); ok {
		return enum.Valid()
	}
	return false
}

func CreateValidator() *validator.Validate {
	v := validator.New()
	Register(v)
	return v
}
