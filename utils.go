package validol

import "reflect"

func lenOf[T any](t T) int {
	return reflect.ValueOf(t).Len()
}

func isNil(val reflect.Value) bool {
	switch val.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice, reflect.UnsafePointer:
		return val.IsNil()
	default:
		return false
	}
}
