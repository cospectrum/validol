package validol

import (
	"cmp"
	"fmt"
)

type Validator[Of any] func(Of) error

type Validatable interface {
	Validate() error
}

func OneOf[T comparable](vals ...T) Validator[T] {
	return func(t T) error {
		for _, val := range vals {
			if val == t {
				return nil
			}
		}
		return failed(fmt.Sprintf("validol.OneOf(%s)(%+v)", fmtVarargs(vals), t))
	}
}

func All[T any](validators ...Validator[T]) Validator[T] {
	return func(t T) error {
		for _, f := range validators {
			if err := f(t); err != nil {
				return err
			}
		}
		return nil
	}
}

func Any[T any](validators ...Validator[T]) Validator[T] {
	return func(t T) error {
		var lastErr error
		for _, f := range validators {
			if err := f(t); err != nil {
				lastErr = err
				continue
			}
			return nil
		}
		return lastErr
	}
}

func Visit[T any](t T) error {
	return visitChildren(t)
}

func Not[T any](fn Validator[T]) Validator[T] {
	return func(t T) error {
		err := fn(t)
		if err != nil {
			return nil //nolint:nilerr
		}
		return failed(fmt.Sprintf("validol.Not(...)(%+v)", t))
	}
}

func Gt[T cmp.Ordered](val T) Validator[T] {
	return func(t T) error {
		ok := t > val
		if !ok {
			return failed(fmt.Sprintf("validol.Gt(%+v)(%+v)", val, t))
		}
		return nil
	}
}

func Gte[T cmp.Ordered](val T) Validator[T] {
	return func(t T) error {
		ok := t >= val
		if !ok {
			return failed(fmt.Sprintf("validol.Gte(%+v)(%+v)", val, t))
		}
		return nil
	}
}

func Lt[T cmp.Ordered](val T) Validator[T] {
	return func(t T) error {
		ok := t < val
		if !ok {
			return failed(fmt.Sprintf("validol.Lt(%+v)(%+v)", val, t))
		}
		return nil
	}
}

func Lte[T cmp.Ordered](val T) Validator[T] {
	return func(t T) error {
		ok := t <= val
		if !ok {
			return failed(fmt.Sprintf("validol.Lte(%+v)(%+v)", val, t))
		}
		return nil
	}
}

func Eq[T comparable](val T) Validator[T] {
	return func(t T) error {
		ok := t == val
		if !ok {
			return failed(fmt.Sprintf("validol.Eq(%+v)(%+v)", val, t))
		}
		return nil
	}
}

func Ne[T comparable](val T) Validator[T] {
	return func(t T) error {
		ok := t != val
		if !ok {
			return failed(fmt.Sprintf("validol.Ne(%+v)(%+v)", val, t))
		}
		return nil
	}
}

func Len[T any](validateLen Validator[int]) Validator[T] {
	return func(t T) error {
		if err := validateLen(lenOf(t)); err != nil {
			return err
		}
		return nil
	}
}
