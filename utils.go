package validol

import "reflect"

func toReflectValue[T any](t T) reflect.Value {
	if val, ok := any(t).(reflect.Value); ok {
		return val
	}
	return reflect.ValueOf(&t).Elem()
}

func lenOf[T any](t T) int {
	return reflect.ValueOf(t).Len()
}

func isNil[T any](t T) bool {
	val := toReflectValue(t)
	switch val.Kind() {
	case reflect.Invalid:
		return true
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map,
		reflect.Pointer, reflect.Slice, reflect.UnsafePointer:
		return val.IsNil()
	default:
		return false
	}
}

func isEmpty[T any](t T) bool {
	val := toReflectValue(t)
	switch val.Kind() {
	case reflect.Invalid:
		return true
	default:
		return val.IsZero()
	}
}
