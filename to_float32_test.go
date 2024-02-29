
package typeconv

import (
    "testing"
)

func TestIntToFloat32(t *testing.T) {

    input := int(1)
    expected := float32(1)

    result, err := IntToFloat32(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt8ToFloat32(t *testing.T) {

    input := int8(1)
    expected := float32(1)

    result, err := Int8ToFloat32(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt16ToFloat32(t *testing.T) {

    input := int16(1)
    expected := float32(1)

    result, err := Int16ToFloat32(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt32ToFloat32(t *testing.T) {

    input := int32(1)
    expected := float32(1)

    result, err := Int32ToFloat32(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt64ToFloat32(t *testing.T) {

    input := int64(1)
    expected := float32(1)

    result, err := Int64ToFloat32(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUintToFloat32(t *testing.T) {

    input := uint(1)
    expected := float32(1)

    result, err := UintToFloat32(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint8ToFloat32(t *testing.T) {

    input := uint8(1)
    expected := float32(1)

    result, err := Uint8ToFloat32(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint16ToFloat32(t *testing.T) {

    input := uint16(1)
    expected := float32(1)

    result, err := Uint16ToFloat32(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint32ToFloat32(t *testing.T) {

    input := uint32(1)
    expected := float32(1)

    result, err := Uint32ToFloat32(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint64ToFloat32(t *testing.T) {

    input := uint64(1)
    expected := float32(1)

    result, err := Uint64ToFloat32(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestFloat64ToFloat32(t *testing.T) {

    input := float64(1)
    expected := float32(1)

    result, err := Float64ToFloat32(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestStringToFloat32(t *testing.T) {

    input := string(1)
    expected := float32(1)

    result, err := StringToFloat32(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


