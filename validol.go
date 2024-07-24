package validol

import (
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
		return fmt.Errorf("validol.OneOf(%s)(%+v) failed", fmtVarargs(vals), t)
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
