package cyclic

import (
	"reflect"
	"unsafe"
)

type link struct {
	tail *link
}

func cyclic(x reflect.Value, seen map[ptr]bool) bool {
	if !x.IsValid() {
		return x.IsValid()
	}

	if x.CanAddr() {
		xptr := ptr{unsafe.Pointer(x.UnsafeAddr()), x.Type()}

		if seen[xptr] {
			return true // already seen
		}
		seen[xptr] = true
	}

	switch x.Kind() {
	case reflect.Ptr, reflect.Interface:
		return cyclic(x.Elem(), seen)

	case reflect.Array, reflect.Slice:
		for i := 0; i < x.Len(); i++ {
			if cyclic(x.Index(i), seen) {
				return true
			}
		}
	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if cyclic(x.Field(i), seen) {
				return true
			}
		}
	case reflect.Map:
		for _, k := range x.MapKeys() {
			if cyclic(x.MapIndex(k), seen) {
				return true
			}
		}
	default:
		return false
	}
	return false
}

func Cyclic(x interface{}) bool {
	seen := make(map[ptr]bool)
	return cyclic(reflect.ValueOf(x), seen)
}

type ptr struct {
	x unsafe.Pointer
	t reflect.Type
}
