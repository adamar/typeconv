
package typeconv

import (
    "testing"
)

func TestIntToUint64(t *testing.T) {

    input := int(1)
    expected := uint64(1)

    result, err := IntToUint64(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt8ToUint64(t *testing.T) {

    input := int8(1)
    expected := uint64(1)

    result, err := Int8ToUint64(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt16ToUint64(t *testing.T) {

    input := int16(1)
    expected := uint64(1)

    result, err := Int16ToUint64(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt32ToUint64(t *testing.T) {

    input := int32(1)
    expected := uint64(1)

    result, err := Int32ToUint64(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt64ToUint64(t *testing.T) {

    input := int64(1)
    expected := uint64(1)

    result, err := Int64ToUint64(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUintToUint64(t *testing.T) {

    input := uint(1)
    expected := uint64(1)

    result, err := UintToUint64(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint8ToUint64(t *testing.T) {

    input := uint8(1)
    expected := uint64(1)

    result, err := Uint8ToUint64(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint16ToUint64(t *testing.T) {

    input := uint16(1)
    expected := uint64(1)

    result, err := Uint16ToUint64(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint32ToUint64(t *testing.T) {

    input := uint32(1)
    expected := uint64(1)

    result, err := Uint32ToUint64(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestFloat32ToUint64(t *testing.T) {

    input := float32(1)
    expected := uint64(1)

    result, err := Float32ToUint64(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestFloat64ToUint64(t *testing.T) {

    input := float64(1)
    expected := uint64(1)

    result, err := Float64ToUint64(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestStringToUint64(t *testing.T) {

    input := string(1)
    expected := uint64(1)

    result, err := StringToUint64(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


