package validol_test

import (
	"testing"

	"github.com/cospectrum/validol"
	"github.com/stretchr/testify/assert"
)

type email string

func (e email) Validate() error {
	return validol.Email(string(e))
}

type wrapper[T any] struct {
	Value T
}

func (w wrapper[T]) Validate() error {
	return validol.Walk(w)
}

func wrap[T any](val T) wrapper[T] {
	return wrapper[T]{
		Value: val,
	}
}

func TestEmail(t *testing.T) {
	e := email("valid@gmail.com")
	assert.NoError(t, e.Validate())
	assert.NoError(t, wrap(e).Validate())

	e = email("invalid|gmail.com")
	assert.Error(t, e.Validate())
	assert.Error(t, wrap(e).Validate())
}
