
package typeconv

import (
	"strconv"
)

// IntToString as the name implies takes a int and converts it into a string
func IntToString(input int) (string, error) {

	return strconv.Itoa(input), nil

}


// Int8ToString as the name implies takes a int8 and converts it into a string
func Int8ToString(input int8) (string, error) {

	return strconv.FormatInt(int64(input), 10), nil

}


// Int16ToString as the name implies takes a int16 and converts it into a string
func Int16ToString(input int16) (string, error) {

	return strconv.FormatInt(int64(input), 10), nil

}


// Int32ToString as the name implies takes a int32 and converts it into a string
func Int32ToString(input int32) (string, error) {

	return strconv.FormatInt(int64(input), 10), nil

}


// Int64ToString as the name implies takes a int64 and converts it into a string
func Int64ToString(input int64) (string, error) {

	return strconv.FormatInt(int64(input), 10), nil

}


// UintToString as the name implies takes a uint and converts it into a string
func UintToString(input uint) (string, error) {

	return strconv.FormatInt(int64(input), 10), nil

}


// Uint8ToString as the name implies takes a uint8 and converts it into a string
func Uint8ToString(input uint8) (string, error) {

	return strconv.FormatInt(int64(input), 10), nil

}


// Uint16ToString as the name implies takes a uint16 and converts it into a string
func Uint16ToString(input uint16) (string, error) {

	return strconv.FormatInt(int64(input), 10), nil

}


// Uint32ToString as the name implies takes a uint32 and converts it into a string
func Uint32ToString(input uint32) (string, error) {

	return strconv.FormatInt(int64(input), 10), nil

}


// Uint64ToString as the name implies takes a uint64 and converts it into a string
func Uint64ToString(input uint64) (string, error) {

	return strconv.FormatInt(int64(input), 10), nil

}


// Float32ToString as the name implies takes a float32 and converts it into a string
func Float32ToString(input float32) (string, error) {

	return strconv.FormatFloat(float64(input), 'f', -1, 32), nil

}


// Float64ToString as the name implies takes a float64 and converts it into a string
func Float64ToString(input float64) (string, error) {

	return strconv.FormatFloat(input, 'f', -1, 32), nil

}
