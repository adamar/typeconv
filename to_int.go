package typeconv

import (
	"strconv"
)

// Int8ToInt as the name implies takes a int8 and converts it into a int
func Int8ToInt(input int8) (int, error) {

	return int(input), nil

}


// Int16ToInt as the name implies takes a int16 and converts it into a int
func Int16ToInt(input int16) (int, error) {

	return int(input), nil

}


// Int32ToInt as the name implies takes a int32 and converts it into a int
func Int32ToInt(input int32) (int, error) {

	return int(input), nil

}


// Int64ToInt as the name implies takes a int64 and converts it into a int
func Int64ToInt(input int64) (int, error) {

	return int(input), nil

}


// UintToInt as the name implies takes a uint and converts it into a int
func UintToInt(input uint) (int, error) {

	return int(input), nil

}


// Uint8ToInt as the name implies takes a uint8 and converts it into a int
func Uint8ToInt(input uint8) (int, error) {

	return int(input), nil

}


// Uint16ToInt as the name implies takes a uint16 and converts it into a int
func Uint16ToInt(input uint16) (int, error) {

	return int(input), nil

}


// Uint32ToInt as the name implies takes a uint32 and converts it into a int
func Uint32ToInt(input uint32) (int, error) {

	return int(input), nil

}


// Uint64ToInt as the name implies takes a uint64 and converts it into a int
func Uint64ToInt(input uint64) (int, error) {

	return int(input), nil

}


// Float32ToInt as the name implies takes a float32 and converts it into a int
func Float32ToInt(input float32) (int, error) {

	return int(input), nil

}


// Float64ToInt as the name implies takes a float64 and converts it into a int
func Float64ToInt(input float64) (int, error) {

	return int(input), nil

}


// StringToInt as the name implies takes a string and converts it into a int
func StringToInt(input string) (int, error) {

	return strconv.Atoi(input)

}


