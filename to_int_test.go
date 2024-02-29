
package typeconv

import (
    "testing"
)

func TestInt8ToInt(t *testing.T) {

    input := int8(1)
    expected := int(1)

    result, err := Int8ToInt(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt16ToInt(t *testing.T) {

    input := int16(1)
    expected := int(1)

    result, err := Int16ToInt(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt32ToInt(t *testing.T) {

    input := int32(1)
    expected := int(1)

    result, err := Int32ToInt(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt64ToInt(t *testing.T) {

    input := int64(1)
    expected := int(1)

    result, err := Int64ToInt(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUintToInt(t *testing.T) {

    input := uint(1)
    expected := int(1)

    result, err := UintToInt(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint8ToInt(t *testing.T) {

    input := uint8(1)
    expected := int(1)

    result, err := Uint8ToInt(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint16ToInt(t *testing.T) {

    input := uint16(1)
    expected := int(1)

    result, err := Uint16ToInt(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint32ToInt(t *testing.T) {

    input := uint32(1)
    expected := int(1)

    result, err := Uint32ToInt(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint64ToInt(t *testing.T) {

    input := uint64(1)
    expected := int(1)

    result, err := Uint64ToInt(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestFloat32ToInt(t *testing.T) {

    input := float32(1)
    expected := int(1)

    result, err := Float32ToInt(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestFloat64ToInt(t *testing.T) {

    input := float64(1)
    expected := int(1)

    result, err := Float64ToInt(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestStringToInt(t *testing.T) {

    input := "1"
    expected := int(1)

    result, err := StringToInt(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


