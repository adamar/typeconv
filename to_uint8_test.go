
package typeconv

import (
    "testing"
)

func TestIntToUint8(t *testing.T) {

    input := int(1)
    expected := uint8(1)

    result, err := IntToUint8(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt8ToUint8(t *testing.T) {

    input := int8(1)
    expected := uint8(1)

    result, err := Int8ToUint8(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt16ToUint8(t *testing.T) {

    input := int16(1)
    expected := uint8(1)

    result, err := Int16ToUint8(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt32ToUint8(t *testing.T) {

    input := int32(1)
    expected := uint8(1)

    result, err := Int32ToUint8(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt64ToUint8(t *testing.T) {

    input := int64(1)
    expected := uint8(1)

    result, err := Int64ToUint8(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUintToUint8(t *testing.T) {

    input := uint(1)
    expected := uint8(1)

    result, err := UintToUint8(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint16ToUint8(t *testing.T) {

    input := uint16(1)
    expected := uint8(1)

    result, err := Uint16ToUint8(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint32ToUint8(t *testing.T) {

    input := uint32(1)
    expected := uint8(1)

    result, err := Uint32ToUint8(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint64ToUint8(t *testing.T) {

    input := uint64(1)
    expected := uint8(1)

    result, err := Uint64ToUint8(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestFloat32ToUint8(t *testing.T) {

    input := float32(1)
    expected := uint8(1)

    result, err := Float32ToUint8(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestFloat64ToUint8(t *testing.T) {

    input := float64(1)
    expected := uint8(1)

    result, err := Float64ToUint8(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestStringToUint8(t *testing.T) {

    input := "1"
    expected := uint8(1)

    result, err := StringToUint8(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


