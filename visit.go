package validol

import "reflect"

func visitChildren(in any) error {
	return innerVisit(in, false)
}

func visit(in any) error {
	return innerVisit(in, true)
}

func innerVisit(in any, validateItself bool) error {
	val, ok := in.(reflect.Value)
	if !ok {
		val = reflect.ValueOf(in)
	}
	if validateItself && val.CanInterface() {
		if v, ok := val.Interface().(Validatable); ok {
			return v.Validate()
		}
	}

	switch val.Kind() {
	case reflect.Pointer:
		return visit(val.Elem())
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			item := val.Index(i)
			if err := visit(item); err != nil {
				return err
			}
			return nil
		}
	case reflect.Map:
		it := val.MapRange()
		for it.Next() {
			if err := visit(it.Key()); err != nil {
				return err
			}
			if err := visit(it.Value()); err != nil {
				return err
			}
		}
		return nil
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			if err := visit(field); err != nil {
				return err
			}
		}
		return nil
	default:
		return nil
	}
	return nil
}
