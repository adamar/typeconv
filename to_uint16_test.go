
package typeconv

import (
    "testing"
)

func TestIntToUint16(t *testing.T) {

    input := int(1)
    expected := uint16(1)

    result, err := IntToUint16(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt8ToUint16(t *testing.T) {

    input := int8(1)
    expected := uint16(1)

    result, err := Int8ToUint16(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt16ToUint16(t *testing.T) {

    input := int16(1)
    expected := uint16(1)

    result, err := Int16ToUint16(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt32ToUint16(t *testing.T) {

    input := int32(1)
    expected := uint16(1)

    result, err := Int32ToUint16(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt64ToUint16(t *testing.T) {

    input := int64(1)
    expected := uint16(1)

    result, err := Int64ToUint16(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUintToUint16(t *testing.T) {

    input := uint(1)
    expected := uint16(1)

    result, err := UintToUint16(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint8ToUint16(t *testing.T) {

    input := uint8(1)
    expected := uint16(1)

    result, err := Uint8ToUint16(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint32ToUint16(t *testing.T) {

    input := uint32(1)
    expected := uint16(1)

    result, err := Uint32ToUint16(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint64ToUint16(t *testing.T) {

    input := uint64(1)
    expected := uint16(1)

    result, err := Uint64ToUint16(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestFloat32ToUint16(t *testing.T) {

    input := float32(1)
    expected := uint16(1)

    result, err := Float32ToUint16(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestFloat64ToUint16(t *testing.T) {

    input := float64(1)
    expected := uint16(1)

    result, err := Float64ToUint16(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestStringToUint16(t *testing.T) {

    input := string(1)
    expected := uint16(1)

    result, err := StringToUint16(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


