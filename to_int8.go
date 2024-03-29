
package typeconv

import (
	"strconv"
)

// IntToInt8 as the name implies takes a int and converts it into a int8
func IntToInt8(input int) (int8, error) {

	return int8(input), nil

}


// Int16ToInt8 as the name implies takes a int16 and converts it into a int8
func Int16ToInt8(input int16) (int8, error) {

	return int8(input), nil

}


// Int32ToInt8 as the name implies takes a int32 and converts it into a int8
func Int32ToInt8(input int32) (int8, error) {

	return int8(input), nil

}


// Int64ToInt8 as the name implies takes a int64 and converts it into a int8
func Int64ToInt8(input int64) (int8, error) {

	return int8(input), nil

}


// UintToInt8 as the name implies takes a uint and converts it into a int8
func UintToInt8(input uint) (int8, error) {

	return int8(input), nil

}


// Uint8ToInt8 as the name implies takes a uint8 and converts it into a int8
func Uint8ToInt8(input uint8) (int8, error) {

	return int8(input), nil

}


// Uint16ToInt8 as the name implies takes a uint16 and converts it into a int8
func Uint16ToInt8(input uint16) (int8, error) {

	return int8(input), nil

}


// Uint32ToInt8 as the name implies takes a uint32 and converts it into a int8
func Uint32ToInt8(input uint32) (int8, error) {

	return int8(input), nil

}


// Uint64ToInt8 as the name implies takes a uint64 and converts it into a int8
func Uint64ToInt8(input uint64) (int8, error) {

	return int8(input), nil

}


// Float32ToInt8 as the name implies takes a float32 and converts it into a int8
func Float32ToInt8(input float32) (int8, error) {

	return int8(input), nil

}


// Float64ToInt8 as the name implies takes a float64 and converts it into a int8
func Float64ToInt8(input float64) (int8, error) {

	return int8(input), nil

}


// StringToInt8 as the name implies takes a string and converts it into a int8
func StringToInt8(input string) (int8, error) {

	result, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return int8(0), err
	}
	return int8(result), nil



}


