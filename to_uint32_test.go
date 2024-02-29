
package typeconv

import (
    "testing"
)

func TestIntToUint32(t *testing.T) {

    input := int(1)
    expected := uint32(1)

    result, err := IntToUint32(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt8ToUint32(t *testing.T) {

    input := int8(1)
    expected := uint32(1)

    result, err := Int8ToUint32(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt16ToUint32(t *testing.T) {

    input := int16(1)
    expected := uint32(1)

    result, err := Int16ToUint32(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt32ToUint32(t *testing.T) {

    input := int32(1)
    expected := uint32(1)

    result, err := Int32ToUint32(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt64ToUint32(t *testing.T) {

    input := int64(1)
    expected := uint32(1)

    result, err := Int64ToUint32(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUintToUint32(t *testing.T) {

    input := uint(1)
    expected := uint32(1)

    result, err := UintToUint32(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint8ToUint32(t *testing.T) {

    input := uint8(1)
    expected := uint32(1)

    result, err := Uint8ToUint32(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint16ToUint32(t *testing.T) {

    input := uint16(1)
    expected := uint32(1)

    result, err := Uint16ToUint32(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint64ToUint32(t *testing.T) {

    input := uint64(1)
    expected := uint32(1)

    result, err := Uint64ToUint32(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestFloat32ToUint32(t *testing.T) {

    input := float32(1)
    expected := uint32(1)

    result, err := Float32ToUint32(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestFloat64ToUint32(t *testing.T) {

    input := float64(1)
    expected := uint32(1)

    result, err := Float64ToUint32(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestComplex64ToUint32(t *testing.T) {

    input := complex64(1)
    expected := uint32(1)

    result, err := Complex64ToUint32(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestComplex128ToUint32(t *testing.T) {

    input := complex128(1)
    expected := uint32(1)

    result, err := Complex128ToUint32(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestStringToUint32(t *testing.T) {

    input := string(1)
    expected := uint32(1)

    result, err := StringToUint32(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


