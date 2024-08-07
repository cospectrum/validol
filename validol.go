package validol

import (
	"cmp"
	"fmt"
	"strings"
)

type Validator[Of any] func(Of) error

type Validatable interface {
	Validate() error
}

func Validate[T any](t T) error {
	if val, ok := any(t).(Validatable); ok {
		return val.Validate()
	}
	return Walk(t)
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

func And[T any](first, second Validator[T]) Validator[T] {
	return All(first, second)
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

func Or[T any](first, second Validator[T]) Validator[T] {
	return Any(first, second)
}

func True(b bool) error {
	if b {
		return nil
	}
	return failed(fmt.Sprintf("validol.True(%v)", b))
}

func False(b bool) error {
	if !b {
		return nil
	}
	return failed(fmt.Sprintf("validol.False(%v)", b))
}

func Walk[T any](t T) error {
	return walkDescendants(t)
}

var _ Validator[any] = Nil

func Nil[T any](t T) error {
	if isNil(t) {
		return nil
	}
	return failed(fmt.Sprintf("validol.Nil(%+v)", t))
}

var _ Validator[any] = NotNil

func NotNil[T any](t T) error {
	notNil := !isNil(t)
	if notNil {
		return nil
	}
	return failed(fmt.Sprintf("validol.NotNil(%+v)", t))
}

var _ Validator[any] = Empty

func Empty[T any](t T) error {
	if isEmpty(t) {
		return nil
	}
	return failed(fmt.Sprintf("validol.Empty(%+v)", t))
}

var _ Validator[any] = Required

func Required[T any](t T) error {
	notEmpty := !isEmpty(t)
	if notEmpty {
		return nil
	}
	return failed(fmt.Sprintf("validol.Required(%+v)", t))
}

func Not[T any](fn Validator[T]) Validator[T] {
	return func(t T) error {
		if err := fn(t); err != nil {
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

func StartsWith(prefix string) Validator[string] {
	return func(s string) error {
		if strings.HasPrefix(s, prefix) {
			return nil
		}
		return failed(fmt.Sprintf("validol.StartsWith(%q)(%q)", prefix, s))
	}
}

func EndsWith(suffix string) Validator[string] {
	return func(s string) error {
		if strings.HasSuffix(s, suffix) {
			return nil
		}
		return failed(fmt.Sprintf("validol.EndsWith(%q)(%q)", suffix, s))
	}
}

func Contains(substr string) Validator[string] {
	return func(s string) error {
		if strings.Contains(s, substr) {
			return nil
		}
		return failed(fmt.Sprintf("validol.Contains(%q)(%q)", substr, s))
	}
}

func ContainsRune(r rune) Validator[string] {
	return func(s string) error {
		if strings.ContainsRune(s, r) {
			return nil
		}
		return failed(fmt.Sprintf("validol.ContainsRune(%q)(%q)", r, s))
	}
}
