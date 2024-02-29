
package typeconv

import (
	"strconv"
)

// IntToInt16 as the name implies takes a int and converts it into a int16
func IntToInt16(input int) (int16, error) {

	return int16(input), nil

}


// Int8ToInt16 as the name implies takes a int8 and converts it into a int16
func Int8ToInt16(input int8) (int16, error) {

	return int16(input), nil

}


// Int32ToInt16 as the name implies takes a int32 and converts it into a int16
func Int32ToInt16(input int32) (int16, error) {

	return int16(input), nil

}


// Int64ToInt16 as the name implies takes a int64 and converts it into a int16
func Int64ToInt16(input int64) (int16, error) {

	return int16(input), nil

}


// UintToInt16 as the name implies takes a uint and converts it into a int16
func UintToInt16(input uint) (int16, error) {

	return int16(input), nil

}


// Uint8ToInt16 as the name implies takes a uint8 and converts it into a int16
func Uint8ToInt16(input uint8) (int16, error) {

	return int16(input), nil

}


// Uint16ToInt16 as the name implies takes a uint16 and converts it into a int16
func Uint16ToInt16(input uint16) (int16, error) {

	return int16(input), nil

}


// Uint32ToInt16 as the name implies takes a uint32 and converts it into a int16
func Uint32ToInt16(input uint32) (int16, error) {

	return int16(input), nil

}


// Uint64ToInt16 as the name implies takes a uint64 and converts it into a int16
func Uint64ToInt16(input uint64) (int16, error) {

	return int16(input), nil

}


// Float32ToInt16 as the name implies takes a float32 and converts it into a int16
func Float32ToInt16(input float32) (int16, error) {

	return int16(input), nil

}


// Float64ToInt16 as the name implies takes a float64 and converts it into a int16
func Float64ToInt16(input float64) (int16, error) {

	return int16(input), nil

}


// StringToInt16 as the name implies takes a string and converts it into a int16
func StringToInt16(input string) (int16, error) {

	result, err := strconv.ParseInt(input, 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(result), nil

}


