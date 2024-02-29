package typeconv

import (
	"strconv"
)

// IntToFloat64 as the name implies takes a int and converts it into a float64
func IntToFloat64(input int) (float64, error) {

	return float64(input), nil

}


// Int8ToFloat64 as the name implies takes a int8 and converts it into a float64
func Int8ToFloat64(input int8) (float64, error) {

	return float64(input), nil

}


// Int16ToFloat64 as the name implies takes a int16 and converts it into a float64
func Int16ToFloat64(input int16) (float64, error) {

	return float64(input), nil

}


// Int32ToFloat64 as the name implies takes a int32 and converts it into a float64
func Int32ToFloat64(input int32) (float64, error) {

	return float64(input), nil

}


// Int64ToFloat64 as the name implies takes a int64 and converts it into a float64
func Int64ToFloat64(input int64) (float64, error) {

	return float64(input), nil

}


// UintToFloat64 as the name implies takes a uint and converts it into a float64
func UintToFloat64(input uint) (float64, error) {

	return float64(input), nil

}


// Uint8ToFloat64 as the name implies takes a uint8 and converts it into a float64
func Uint8ToFloat64(input uint8) (float64, error) {

	return float64(input), nil

}


// Uint16ToFloat64 as the name implies takes a uint16 and converts it into a float64
func Uint16ToFloat64(input uint16) (float64, error) {

	return float64(input), nil

}


// Uint32ToFloat64 as the name implies takes a uint32 and converts it into a float64
func Uint32ToFloat64(input uint32) (float64, error) {

	return float64(input), nil

}


// Uint64ToFloat64 as the name implies takes a uint64 and converts it into a float64
func Uint64ToFloat64(input uint64) (float64, error) {

	return float64(input), nil

}


// Float32ToFloat64 as the name implies takes a float32 and converts it into a float64
func Float32ToFloat64(input float32) (float64, error) {

	return float64(input), nil

}


// StringToFloat64 as the name implies takes a string and converts it into a float64
func StringToFloat64(input string) (float64, error) {

	return strconv.ParseFloat(input, 64)

}


