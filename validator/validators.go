package validator

import (
	validator2 "github.com/go-playground/validator/v10"
	"strings"
)

// CoolValidation - Custom validation for struct tags ("is-cool")
func CoolValidation(field validator2.FieldLevel)bool{
	value := field.Field().String()
	value = strings.ToLower(value)
	return strings.Contains(value, "cool")
}
