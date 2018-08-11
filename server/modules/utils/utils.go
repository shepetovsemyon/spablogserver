package utils

import (
	"strconv"
)

//Abs of int
func Abs(value int) int {
	if value < 0 {
		value = -value
	}
	return value
}

//Str2UInt converting string to positive int
func Str2UInt(value string, defValue int) int {
	_value, err := strconv.Atoi(value)
	if err != nil {
		_value = defValue
	}

	return Abs(_value)
}
