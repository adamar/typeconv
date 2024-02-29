
package typeconv

import (
    "testing"
)

func TestIntToInt16(t *testing.T) {

    input := int(1)
    expected := int16(1)

    result, err := IntToInt16(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt8ToInt16(t *testing.T) {

    input := int8(1)
    expected := int16(1)

    result, err := Int8ToInt16(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt32ToInt16(t *testing.T) {

    input := int32(1)
    expected := int16(1)

    result, err := Int32ToInt16(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt64ToInt16(t *testing.T) {

    input := int64(1)
    expected := int16(1)

    result, err := Int64ToInt16(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUintToInt16(t *testing.T) {

    input := uint(1)
    expected := int16(1)

    result, err := UintToInt16(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint8ToInt16(t *testing.T) {

    input := uint8(1)
    expected := int16(1)

    result, err := Uint8ToInt16(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint16ToInt16(t *testing.T) {

    input := uint16(1)
    expected := int16(1)

    result, err := Uint16ToInt16(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint32ToInt16(t *testing.T) {

    input := uint32(1)
    expected := int16(1)

    result, err := Uint32ToInt16(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint64ToInt16(t *testing.T) {

    input := uint64(1)
    expected := int16(1)

    result, err := Uint64ToInt16(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestFloat32ToInt16(t *testing.T) {

    input := float32(1)
    expected := int16(1)

    result, err := Float32ToInt16(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestFloat64ToInt16(t *testing.T) {

    input := float64(1)
    expected := int16(1)

    result, err := Float64ToInt16(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestStringToInt16(t *testing.T) {

    input := "1"
    expected := int16(1)

    result, err := StringToInt16(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


