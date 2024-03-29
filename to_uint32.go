
package typeconv

import (
	"strconv"
)

// IntToUint32 as the name implies takes a int and converts it into a uint32
func IntToUint32(input int) (uint32, error) {

	return uint32(input), nil

}


// Int8ToUint32 as the name implies takes a int8 and converts it into a uint32
func Int8ToUint32(input int8) (uint32, error) {

	return uint32(input), nil

}


// Int16ToUint32 as the name implies takes a int16 and converts it into a uint32
func Int16ToUint32(input int16) (uint32, error) {

	return uint32(input), nil

}


// Int32ToUint32 as the name implies takes a int32 and converts it into a uint32
func Int32ToUint32(input int32) (uint32, error) {

	return uint32(input), nil

}


// Int64ToUint32 as the name implies takes a int64 and converts it into a uint32
func Int64ToUint32(input int64) (uint32, error) {

	return uint32(input), nil

}


// UintToUint32 as the name implies takes a uint and converts it into a uint32
func UintToUint32(input uint) (uint32, error) {

	return uint32(input), nil

}


// Uint8ToUint32 as the name implies takes a uint8 and converts it into a uint32
func Uint8ToUint32(input uint8) (uint32, error) {

	return uint32(input), nil

}


// Uint16ToUint32 as the name implies takes a uint16 and converts it into a uint32
func Uint16ToUint32(input uint16) (uint32, error) {

	return uint32(input), nil

}


// Uint64ToUint32 as the name implies takes a uint64 and converts it into a uint32
func Uint64ToUint32(input uint64) (uint32, error) {

	return uint32(input), nil

}


// Float32ToUint32 as the name implies takes a float32 and converts it into a uint32
func Float32ToUint32(input float32) (uint32, error) {

	return uint32(input), nil

}


// Float64ToUint32 as the name implies takes a float64 and converts it into a uint32
func Float64ToUint32(input float64) (uint32, error) {

	return uint32(input), nil

}


// StringToUint32 as the name implies takes a string and converts it into a uint32
func StringToUint32(input string) (uint32, error) {

	result, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		return uint32(0), err
	}
	return uint32(result), nil

}


