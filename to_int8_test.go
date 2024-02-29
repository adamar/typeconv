
package typeconv

import (
    "testing"
)

func TestIntToInt8(t *testing.T) {

    input := int(1)
    expected := int8(1)

    result, err := IntToInt8(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt16ToInt8(t *testing.T) {

    input := int16(1)
    expected := int8(1)

    result, err := Int16ToInt8(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt32ToInt8(t *testing.T) {

    input := int32(1)
    expected := int8(1)

    result, err := Int32ToInt8(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt64ToInt8(t *testing.T) {

    input := int64(1)
    expected := int8(1)

    result, err := Int64ToInt8(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUintToInt8(t *testing.T) {

    input := uint(1)
    expected := int8(1)

    result, err := UintToInt8(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint8ToInt8(t *testing.T) {

    input := uint8(1)
    expected := int8(1)

    result, err := Uint8ToInt8(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint16ToInt8(t *testing.T) {

    input := uint16(1)
    expected := int8(1)

    result, err := Uint16ToInt8(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint32ToInt8(t *testing.T) {

    input := uint32(1)
    expected := int8(1)

    result, err := Uint32ToInt8(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint64ToInt8(t *testing.T) {

    input := uint64(1)
    expected := int8(1)

    result, err := Uint64ToInt8(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestFloat32ToInt8(t *testing.T) {

    input := float32(1)
    expected := int8(1)

    result, err := Float32ToInt8(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestFloat64ToInt8(t *testing.T) {

    input := float64(1)
    expected := int8(1)

    result, err := Float64ToInt8(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestStringToInt8(t *testing.T) {

    input := "1"
    expected := int8(1)

    result, err := StringToInt8(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


