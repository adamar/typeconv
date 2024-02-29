
package typeconv

import (
    "testing"
)

func TestIntToString(t *testing.T) {

    input := int(1)
    expected := string(1)

    result, err := IntToString(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt8ToString(t *testing.T) {

    input := int8(1)
    expected := string(1)

    result, err := Int8ToString(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt16ToString(t *testing.T) {

    input := int16(1)
    expected := string(1)

    result, err := Int16ToString(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt32ToString(t *testing.T) {

    input := int32(1)
    expected := string(1)

    result, err := Int32ToString(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestInt64ToString(t *testing.T) {

    input := int64(1)
    expected := string(1)

    result, err := Int64ToString(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUintToString(t *testing.T) {

    input := uint(1)
    expected := string(1)

    result, err := UintToString(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint8ToString(t *testing.T) {

    input := uint8(1)
    expected := string(1)

    result, err := Uint8ToString(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint16ToString(t *testing.T) {

    input := uint16(1)
    expected := string(1)

    result, err := Uint16ToString(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint32ToString(t *testing.T) {

    input := uint32(1)
    expected := string(1)

    result, err := Uint32ToString(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestUint64ToString(t *testing.T) {

    input := uint64(1)
    expected := string(1)

    result, err := Uint64ToString(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestFloat32ToString(t *testing.T) {

    input := float32(1)
    expected := string(1)

    result, err := Float32ToString(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}


func TestFloat64ToString(t *testing.T) {

    input := float64(1)
    expected := string(1)

    result, err := Float64ToString(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was Incorrect, got: %v, wanted: %v.", result, expected)
    }

}
