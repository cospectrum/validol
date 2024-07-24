package validol

import "reflect"

func lenOf[T any](t T) int {
	return reflect.ValueOf(t).Len()
}
