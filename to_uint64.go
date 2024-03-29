
package typeconv

import (
	"strconv"
)

// IntToUint64 as the name implies takes a int and converts it into a uint64
func IntToUint64(input int) (uint64, error) {

	return uint64(input), nil

}


// Int8ToUint64 as the name implies takes a int8 and converts it into a uint64
func Int8ToUint64(input int8) (uint64, error) {

	return uint64(input), nil

}


// Int16ToUint64 as the name implies takes a int16 and converts it into a uint64
func Int16ToUint64(input int16) (uint64, error) {

	return uint64(input), nil

}


// Int32ToUint64 as the name implies takes a int32 and converts it into a uint64
func Int32ToUint64(input int32) (uint64, error) {

	return uint64(input), nil

}


// Int64ToUint64 as the name implies takes a int64 and converts it into a uint64
func Int64ToUint64(input int64) (uint64, error) {

	return uint64(input), nil

}


// UintToUint64 as the name implies takes a uint and converts it into a uint64
func UintToUint64(input uint) (uint64, error) {

	return uint64(input), nil

}


// Uint8ToUint64 as the name implies takes a uint8 and converts it into a uint64
func Uint8ToUint64(input uint8) (uint64, error) {

	return uint64(input), nil

}


// Uint16ToUint64 as the name implies takes a uint16 and converts it into a uint64
func Uint16ToUint64(input uint16) (uint64, error) {

	return uint64(input), nil

}


// Uint32ToUint64 as the name implies takes a uint32 and converts it into a uint64
func Uint32ToUint64(input uint32) (uint64, error) {

	return uint64(input), nil

}


// Float32ToUint64 as the name implies takes a float32 and converts it into a uint64
func Float32ToUint64(input float32) (uint64, error) {

	return uint64(input), nil

}


// Float64ToUint64 as the name implies takes a float64 and converts it into a uint64
func Float64ToUint64(input float64) (uint64, error) {

	return uint64(input), nil

}


// StringToUint64 as the name implies takes a string and converts it into a uint64
func StringToUint64(input string) (uint64, error) {

	return strconv.ParseUint(input, 10, 64)

}


