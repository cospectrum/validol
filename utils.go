package validol

import "reflect"

func toReflectValue(t any) reflect.Value {
	val, ok := t.(reflect.Value)
	if !ok {
		val = reflect.ValueOf(t)
	}
	return val
}

func lenOf[T any](t T) int {
	return toReflectValue(t).Len()
}

func isNil[T any](t T) bool {
	val := toReflectValue(t)
	switch val.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice, reflect.UnsafePointer:
		return val.IsNil()
	default:
		return false
	}
}
