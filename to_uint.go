
package typeconv

import (
	"strconv"
)

// IntToUint as the name implies takes a int and converts it into a uint
func IntToUint(input int) (uint, error) {

	// NOT IMPLEMENTED YET
	return uint(input), nil

}


// Int8ToUint as the name implies takes a int8 and converts it into a uint
func Int8ToUint(input int8) (uint, error) {

	// NOT IMPLEMENTED YET
	return uint(input), nil

}


// Int16ToUint as the name implies takes a int16 and converts it into a uint
func Int16ToUint(input int16) (uint, error) {

	// NOT IMPLEMENTED YET
	return uint(input), nil

}


// Int32ToUint as the name implies takes a int32 and converts it into a uint
func Int32ToUint(input int32) (uint, error) {

	// NOT IMPLEMENTED YET
	return uint(input), nil

}


// Int64ToUint as the name implies takes a int64 and converts it into a uint
func Int64ToUint(input int64) (uint, error) {

	// NOT IMPLEMENTED YET
	return uint(input), nil

}


// Uint8ToUint as the name implies takes a uint8 and converts it into a uint
func Uint8ToUint(input uint8) (uint, error) {

	// NOT IMPLEMENTED YET
	return uint(input), nil

}


// Uint16ToUint as the name implies takes a uint16 and converts it into a uint
func Uint16ToUint(input uint16) (uint, error) {

	// NOT IMPLEMENTED YET
	return uint(input), nil

}


// Uint32ToUint as the name implies takes a uint32 and converts it into a uint
func Uint32ToUint(input uint32) (uint, error) {

	// NOT IMPLEMENTED YET
	return uint(input), nil

}


// Uint64ToUint as the name implies takes a uint64 and converts it into a uint
func Uint64ToUint(input uint64) (uint, error) {

	// NOT IMPLEMENTED YET
	return uint(input), nil

}


// Float32ToUint as the name implies takes a float32 and converts it into a uint
func Float32ToUint(input float32) (uint, error) {

	// NOT IMPLEMENTED YET
	return uint(input), nil

}


// Float64ToUint as the name implies takes a float64 and converts it into a uint
func Float64ToUint(input float64) (uint, error) {

	// NOT IMPLEMENTED YET
	return uint(input), nil

}


// StringToUint as the name implies takes a string and converts it into a uint
func StringToUint(input string) (uint, error) {

	result, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(result), nil

}


