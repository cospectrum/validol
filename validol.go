package validol

import "fmt"

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
