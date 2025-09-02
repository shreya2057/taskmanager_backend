package utils

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func Validate(validate *validator.Validate, obj any) (int, map[string]string) {

	if err := validate.Struct(obj); err != nil {
		errorsMap := map[string]string{}
		var key string
		for _, e := range err.(validator.ValidationErrors) {

			elem := reflect.TypeOf(obj).Elem()
			if elem.Kind() == reflect.Ptr {
				elem = elem.Elem()
			}
			field, _ := elem.FieldByName(e.Field())
			jsonTag := field.Tag.Get("json")
			if jsonTag != "" {
				key = jsonTag
			} else {
				key = e.Field()
			}
			errorsMap[key] = ValidationError(e, key)
		}

		if len(errorsMap) == 0 {
			return 400, map[string]string{"error": "validation failed"}
		}
		return 400, errorsMap
	}
	return 200, nil
}

func ValidationError(e validator.FieldError, key string) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", key)
	case "email":
		return "invalid email format"
	case "password":
		return "password must be at least 8 characters long, contain at least one uppercase letter, one number, and one special character"
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long", key, e.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters long", key, e.Param())
	case "alphanum":
		return fmt.Sprintf("%s must contain only alphanumeric characters", key)
	case "eqfield":
		return fmt.Sprintf("%s must be equal to %s", key, e.Param())
	default:
		return fmt.Sprintf("%s is not valid", key)

	}
}
