
package typeconv

import (
	"strconv"
)

// IntToUint8 as the name implies takes a int and converts it into a uint8
func IntToUint8(input int) (uint8, error) {

	return uint8(input), nil

}


// Int8ToUint8 as the name implies takes a int8 and converts it into a uint8
func Int8ToUint8(input int8) (uint8, error) {

	return uint8(input), nil

}


// Int16ToUint8 as the name implies takes a int16 and converts it into a uint8
func Int16ToUint8(input int16) (uint8, error) {

	return uint8(input), nil

}


// Int32ToUint8 as the name implies takes a int32 and converts it into a uint8
func Int32ToUint8(input int32) (uint8, error) {

	return uint8(input), nil

}


// Int64ToUint8 as the name implies takes a int64 and converts it into a uint8
func Int64ToUint8(input int64) (uint8, error) {

	return uint8(input), nil

}


// UintToUint8 as the name implies takes a uint and converts it into a uint8
func UintToUint8(input uint) (uint8, error) {

	return uint8(input), nil

}


// Uint16ToUint8 as the name implies takes a uint16 and converts it into a uint8
func Uint16ToUint8(input uint16) (uint8, error) {

	return uint8(input), nil

}


// Uint32ToUint8 as the name implies takes a uint32 and converts it into a uint8
func Uint32ToUint8(input uint32) (uint8, error) {

	return uint8(input), nil

}


// Uint64ToUint8 as the name implies takes a uint64 and converts it into a uint8
func Uint64ToUint8(input uint64) (uint8, error) {

	return uint8(input), nil

}


// Float32ToUint8 as the name implies takes a float32 and converts it into a uint8
func Float32ToUint8(input float32) (uint8, error) {

	return uint8(input), nil

}


// Float64ToUint8 as the name implies takes a float64 and converts it into a uint8
func Float64ToUint8(input float64) (uint8, error) {

	return uint8(input), nil

}


// StringToUint8 as the name implies takes a string and converts it into a uint8
func StringToUint8(input string) (uint8, error) {

	result, err := strconv.ParseUint(input, 10, 8)
	if err != nil {
		return uint8(0), err
	}
	return uint8(result), nil

}


