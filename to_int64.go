
package typeconv

import (
	"strconv"
)

// IntToInt64 as the name implies takes a int and converts it into a int64
func IntToInt64(input int) (int64, error) {

	// NOT IMPLEMENTED YET
	return int64(input), nil

}


// Int8ToInt64 as the name implies takes a int8 and converts it into a int64
func Int8ToInt64(input int8) (int64, error) {

	// NOT IMPLEMENTED YET
	return int64(input), nil

}


// Int16ToInt64 as the name implies takes a int16 and converts it into a int64
func Int16ToInt64(input int16) (int64, error) {

	// NOT IMPLEMENTED YET
	return int64(input), nil

}


// Int32ToInt64 as the name implies takes a int32 and converts it into a int64
func Int32ToInt64(input int32) (int64, error) {

	// NOT IMPLEMENTED YET
	return int64(input), nil

}


// UintToInt64 as the name implies takes a uint and converts it into a int64
func UintToInt64(input uint) (int64, error) {

	// NOT IMPLEMENTED YET
	return int64(input), nil

}


// Uint8ToInt64 as the name implies takes a uint8 and converts it into a int64
func Uint8ToInt64(input uint8) (int64, error) {

	// NOT IMPLEMENTED YET
	return int64(input), nil

}


// Uint16ToInt64 as the name implies takes a uint16 and converts it into a int64
func Uint16ToInt64(input uint16) (int64, error) {

	// NOT IMPLEMENTED YET
	return int64(input), nil

}


// Uint32ToInt64 as the name implies takes a uint32 and converts it into a int64
func Uint32ToInt64(input uint32) (int64, error) {

	// NOT IMPLEMENTED YET
	return int64(input), nil

}


// Uint64ToInt64 as the name implies takes a uint64 and converts it into a int64
func Uint64ToInt64(input uint64) (int64, error) {

	// NOT IMPLEMENTED YET
	return int64(input), nil

}


// Float32ToInt64 as the name implies takes a float32 and converts it into a int64
func Float32ToInt64(input float32) (int64, error) {

	// NOT IMPLEMENTED YET
	return int64(input), nil

}


// Float64ToInt64 as the name implies takes a float64 and converts it into a int64
func Float64ToInt64(input float64) (int64, error) {

	// NOT IMPLEMENTED YET
	return int64(input), nil

}


// StringToInt64 as the name implies takes a string and converts it into a int64
func StringToInt64(input string) (int64, error) {

	return strconv.ParseInt(input, 10, 64)

}


