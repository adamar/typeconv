
package typeconv

import (
    "testing"
)

func TestIntToInt64(t *testing.T) {

    input := int(1)
    expected := int64(1)

    result, err := IntToInt64(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt8ToInt64(t *testing.T) {

    input := int8(1)
    expected := int64(1)

    result, err := Int8ToInt64(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt16ToInt64(t *testing.T) {

    input := int16(1)
    expected := int64(1)

    result, err := Int16ToInt64(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt32ToInt64(t *testing.T) {

    input := int32(1)
    expected := int64(1)

    result, err := Int32ToInt64(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUintToInt64(t *testing.T) {

    input := uint(1)
    expected := int64(1)

    result, err := UintToInt64(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint8ToInt64(t *testing.T) {

    input := uint8(1)
    expected := int64(1)

    result, err := Uint8ToInt64(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint16ToInt64(t *testing.T) {

    input := uint16(1)
    expected := int64(1)

    result, err := Uint16ToInt64(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint32ToInt64(t *testing.T) {

    input := uint32(1)
    expected := int64(1)

    result, err := Uint32ToInt64(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint64ToInt64(t *testing.T) {

    input := uint64(1)
    expected := int64(1)

    result, err := Uint64ToInt64(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestFloat32ToInt64(t *testing.T) {

    input := float32(1)
    expected := int64(1)

    result, err := Float32ToInt64(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestFloat64ToInt64(t *testing.T) {

    input := float64(1)
    expected := int64(1)

    result, err := Float64ToInt64(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestComplex64ToInt64(t *testing.T) {

    input := complex64(1)
    expected := int64(1)

    result, err := Complex64ToInt64(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestComplex128ToInt64(t *testing.T) {

    input := complex128(1)
    expected := int64(1)

    result, err := Complex128ToInt64(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestStringToInt64(t *testing.T) {

    input := string(1)
    expected := int64(1)

    result, err := StringToInt64(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


