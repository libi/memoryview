package memoryview

import (
	"errors"
	"reflect"
	"unsafe"
)

func MemView(value interface{}) ([]byte, error) {
	val := reflect.ValueOf(value)
	if val.Kind() != reflect.Ptr {
		return []byte{}, errors.New("value must be pointer")
	}
	pStart := val.Pointer()
	size := val.Elem().Type().Size()
	view := make([]byte, 0)
	for i := 0; i < int(size); i++ {
		viewByte := *(*byte)((unsafe.Pointer(uintptr(int(pStart) + i))))
		view = append(view, viewByte)
	}
	return view, nil
}
