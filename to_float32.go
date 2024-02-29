
package typeconv

// IntToFloat32 as the name implies takes a int and converts it into a float32
func IntToFloat32(input int) (float32, error) {

	// NOT IMPLEMENTED YET
	return float32(input), nil

}


// Int8ToFloat32 as the name implies takes a int8 and converts it into a float32
func Int8ToFloat32(input int8) (float32, error) {

	// NOT IMPLEMENTED YET
	return float32(input), nil

}


// Int16ToFloat32 as the name implies takes a int16 and converts it into a float32
func Int16ToFloat32(input int16) (float32, error) {

	// NOT IMPLEMENTED YET
	return float32(input), nil

}


// Int32ToFloat32 as the name implies takes a int32 and converts it into a float32
func Int32ToFloat32(input int32) (float32, error) {

	// NOT IMPLEMENTED YET
	return float32(input), nil

}


// Int64ToFloat32 as the name implies takes a int64 and converts it into a float32
func Int64ToFloat32(input int64) (float32, error) {

	// NOT IMPLEMENTED YET
	return float32(input), nil

}


// UintToFloat32 as the name implies takes a uint and converts it into a float32
func UintToFloat32(input uint) (float32, error) {

	// NOT IMPLEMENTED YET
	return float32(input), nil

}


// Uint8ToFloat32 as the name implies takes a uint8 and converts it into a float32
func Uint8ToFloat32(input uint8) (float32, error) {

	// NOT IMPLEMENTED YET
	return float32(input), nil

}


// Uint16ToFloat32 as the name implies takes a uint16 and converts it into a float32
func Uint16ToFloat32(input uint16) (float32, error) {

	// NOT IMPLEMENTED YET
	return float32(input), nil

}


// Uint32ToFloat32 as the name implies takes a uint32 and converts it into a float32
func Uint32ToFloat32(input uint32) (float32, error) {

	// NOT IMPLEMENTED YET
	return float32(input), nil

}


// Uint64ToFloat32 as the name implies takes a uint64 and converts it into a float32
func Uint64ToFloat32(input uint64) (float32, error) {

	// NOT IMPLEMENTED YET
	return float32(input), nil

}


// Float64ToFloat32 as the name implies takes a float64 and converts it into a float32
func Float64ToFloat32(input float64) (float32, error) {

	// NOT IMPLEMENTED YET
	return float32(input), nil

}


// Complex64ToFloat32 as the name implies takes a complex64 and converts it into a float32
func Complex64ToFloat32(input complex64) (float32, error) {

	return float32(real(c)), nil

}


// Complex128ToFloat32 as the name implies takes a complex128 and converts it into a float32
func Complex128ToFloat32(input complex128) (float32, error) {

	return float32(real(c)), nil

}


// StringToFloat32 as the name implies takes a string and converts it into a float32
func StringToFloat32(input string) (float32, error) {


	result, err := strconv.ParseFloat(input, 32)
	if err != nil {
		return 0, err
	}
	return float32(result), nil

}


