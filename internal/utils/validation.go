package utils

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func ValidationErrors(e validator.FieldError, obj any) (int, map[string]string) {
	errorsMap := map[string]string{}
	var key string

	elem := reflect.TypeOf(obj).Elem()
	field, _ := elem.FieldByName(e.Field())
	jsonTag := field.Tag.Get("json")
	if jsonTag != "" {
		key = jsonTag
	} else {
		key = e.Field()
	}
	errorsMap[key] = e.Tag()
	if e.Tag() == "required" {
		errorsMap[key] = fmt.Sprintf("%s is required", key)
	} else if e.Tag() == "email" {
		errorsMap[key] = "invalid email format"
	} else if e.Tag() == "password" {
		errorsMap[key] = "password must be at least 8 characters long, contain at least one uppercase letter, one number, and one special character"
	} else if e.Tag() == "min" {
		errorsMap[key] = fmt.Sprintf("%s must be at least %s characters long", key, e.Param())
	} else if e.Tag() == "max" {
		errorsMap[key] = fmt.Sprintf("%s must be at most %s characters long", key, e.Param())
	} else if e.Tag() == "alphanum" {
		errorsMap[key] = fmt.Sprintf("%s must contain only alphanumeric characters", key)
	} else if e.Tag() == "eqfield" {
		errorsMap[key] = fmt.Sprintf("%s must be equal to %s", key, e.Param())
	}

	if len(errorsMap) == 0 {
		return 400, map[string]string{"error": "validation failed"}
	}
	return 400, errorsMap
}
