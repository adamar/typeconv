
package typeconv

import (
	"fmt"
	"strconv"
	"strings"
)

// IntToComplex128 as the name implies takes a int and converts it into a complex128
func IntToComplex128(input int) (complex128, error) {

	return complex(float64(input), 0), nil

}


// Int8ToComplex128 as the name implies takes a int8 and converts it into a complex128
func Int8ToComplex128(input int8) (complex128, error) {

	return complex(float64(input), 0), nil

}


// Int16ToComplex128 as the name implies takes a int16 and converts it into a complex128
func Int16ToComplex128(input int16) (complex128, error) {

	return complex(float64(input), 0), nil

}


// Int32ToComplex128 as the name implies takes a int32 and converts it into a complex128
func Int32ToComplex128(input int32) (complex128, error) {

	return complex(float64(input), 0), nil

}


// Int64ToComplex128 as the name implies takes a int64 and converts it into a complex128
func Int64ToComplex128(input int64) (complex128, error) {

	return complex(float64(input), 0), nil

}


// UintToComplex128 as the name implies takes a uint and converts it into a complex128
func UintToComplex128(input uint) (complex128, error) {

	return complex(float64(input), 0), nil

}


// Uint8ToComplex128 as the name implies takes a uint8 and converts it into a complex128
func Uint8ToComplex128(input uint8) (complex128, error) {

	return complex(float64(input), 0), nil

}


// Uint16ToComplex128 as the name implies takes a uint16 and converts it into a complex128
func Uint16ToComplex128(input uint16) (complex128, error) {

	return complex(float64(input), 0), nil

}


// Uint32ToComplex128 as the name implies takes a uint32 and converts it into a complex128
func Uint32ToComplex128(input uint32) (complex128, error) {

	return complex(float64(input), 0), nil

}


// Uint64ToComplex128 as the name implies takes a uint64 and converts it into a complex128
func Uint64ToComplex128(input uint64) (complex128, error) {

	return complex(float64(input), 0), nil

}


// Float32ToComplex128 as the name implies takes a float32 and converts it into a complex128
func Float32ToComplex128(input float32) (complex128, error) {

	return complex(float64(input), 0), nil

}


// Float64ToComplex128 as the name implies takes a float64 and converts it into a complex128
func Float64ToComplex128(input float64) (complex128, error) {

	return complex(float64(input), 0), nil

}


// Complex64ToComplex128 as the name implies takes a complex64 and converts it into a complex128
func Complex64ToComplex128(input complex64) (complex128, error) {

	return complex128(input), nil

}


// StringToComplex128 as the name implies takes a string and converts it into a complex128
func StringToComplex128(input string) (complex128, error) {

	// Remove whitespace and lowercase the string
	s := strings.TrimSpace(strings.ToLower(input))

	// Find the index of the '+' or '-' sign
	signIndex := strings.LastIndexAny(input, "+-")
	if signIndex == -1 || signIndex == 0 {
		return 0, fmt.Errorf("invalid format: %s", s)
	}

	// Extract the real and imaginary parts
	realPart := s[:signIndex]
	imaginaryPart := s[signIndex : len(s)-1] // Exclude the 'i' at the end

	// Convert the real and imaginary parts to float64
	real, err := strconv.ParseFloat(realPart, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid real part: %s", realPart)
	}

	imaginary, err := strconv.ParseFloat(imaginaryPart, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid imaginary part: %s", imaginaryPart)
	}

	return complex(real, imaginary), nil

}


