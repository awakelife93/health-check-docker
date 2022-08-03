package utils

import "reflect"

func IsFunc(value interface{}) bool {
	return reflect.TypeOf(value).Kind() == reflect.Func
}

func IsString(value interface{}) bool {
	return reflect.TypeOf(value).Kind() == reflect.String
}

func IsStringAndNotEmpty(value interface{}) bool {
	return IsString(value) && value != ""
}
