package validol

import "reflect"

func walkDescendants(in any) error {
	return validationWalk(in, false)
}

func walk(in any) error {
	return validationWalk(in, true)
}

func validationWalk(in any, validateItself bool) error {
	val := toReflectValue(in)
	if isNil(val) {
		return nil
	}
	if validateItself && val.CanInterface() {
		if v, ok := val.Interface().(Validatable); ok {
			return v.Validate()
		}
	}

	switch val.Kind() {
	case reflect.Pointer:
		return validationWalk(val.Elem(), validateItself)
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			item := val.Index(i)
			if err := walk(item); err != nil {
				return err
			}
		}
		return nil
	case reflect.Map:
		it := val.MapRange()
		for it.Next() {
			if err := walk(it.Key()); err != nil {
				return err
			}
			if err := walk(it.Value()); err != nil {
				return err
			}
		}
		return nil
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			if err := walk(field); err != nil {
				return err
			}
		}
		return nil
	default:
		return nil
	}
}
