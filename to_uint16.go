
package typeconv

import (
	"strconv"
)

// IntToUint16 as the name implies takes a int and converts it into a uint16
func IntToUint16(input int) (uint16, error) {

	return uint16(input), nil

}


// Int8ToUint16 as the name implies takes a int8 and converts it into a uint16
func Int8ToUint16(input int8) (uint16, error) {

	return uint16(input), nil

}


// Int16ToUint16 as the name implies takes a int16 and converts it into a uint16
func Int16ToUint16(input int16) (uint16, error) {

	return uint16(input), nil

}


// Int32ToUint16 as the name implies takes a int32 and converts it into a uint16
func Int32ToUint16(input int32) (uint16, error) {

	return uint16(input), nil

}


// Int64ToUint16 as the name implies takes a int64 and converts it into a uint16
func Int64ToUint16(input int64) (uint16, error) {

	return uint16(input), nil

}


// UintToUint16 as the name implies takes a uint and converts it into a uint16
func UintToUint16(input uint) (uint16, error) {

	return uint16(input), nil

}


// Uint8ToUint16 as the name implies takes a uint8 and converts it into a uint16
func Uint8ToUint16(input uint8) (uint16, error) {

	return uint16(input), nil

}


// Uint32ToUint16 as the name implies takes a uint32 and converts it into a uint16
func Uint32ToUint16(input uint32) (uint16, error) {

	return uint16(input), nil

}


// Uint64ToUint16 as the name implies takes a uint64 and converts it into a uint16
func Uint64ToUint16(input uint64) (uint16, error) {

	return uint16(input), nil

}


// Float32ToUint16 as the name implies takes a float32 and converts it into a uint16
func Float32ToUint16(input float32) (uint16, error) {

	return uint16(input), nil

}


// Float64ToUint16 as the name implies takes a float64 and converts it into a uint16
func Float64ToUint16(input float64) (uint16, error) {

	return uint16(input), nil

}


// StringToUint16 as the name implies takes a string and converts it into a uint16
func StringToUint16(input string) (uint16, error) {

	result, err := strconv.ParseUint(input, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(result), nil

}


