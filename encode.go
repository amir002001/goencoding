package goencoding

import (
	"fmt"
	"reflect"
)

type dataType struct {
	symbol byte
	size   int
}

var (
	int8Symbol    = dataType{0x10, 1}
	int16Symbol   = dataType{0x11, 2}
	int32Symbol   = dataType{0x12, 4}
	int64Symbol   = dataType{0x13, 8}
	float32Symbol = dataType{0x14, 4}
	float64Symbol = dataType{0x15, 8}
)

func Encode(args ...interface{}) {
	encoded := []byte{}
	for _, arg := range args {
		switch argType := reflect.ValueOf(arg).Type().Kind(); argType {
		case reflect.Int8:
			EncodeInt8(arg.(int8), &encoded)
		case reflect.Int16:
		case reflect.Int32:
		case reflect.Int64:
		case reflect.Float32:
		case reflect.Float64:
		}
	}
}

func Decode(encoded []byte) ([]interface{}, error) {
	idx := 0
	values := []interface{}{}
	for len(values) != 0 {
		valueType := encoded[idx]
		idx++
		switch valueType {
		default:
			return nil, fmt.Errorf("unexpected type when decoding byte array")
		}

	}
	return values, nil
}

func EncodeInt8(arg int8, encoded *[]byte) {
}

func EncodeInt16(arg int16, encoded *[]byte) {
}

func EncodeInt32(arg int32, encoded *[]byte) {
}

func EncodeInt64(arg int64, encoded *[]byte) {
}

func EncodeFloat32(arg float32, encoded *[]byte) {
}

func EncodeFloat64(arg float64, encoded *[]byte) {
}
