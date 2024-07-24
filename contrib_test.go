package validol_test

import (
	"testing"

	vd "github.com/cospectrum/validol"
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

func wrap[T any](val T) wrapper[T] {
	return wrapper[T]{
		Value: val,
	}
}

func TestEmail(t *testing.T) {
	e := email("valid@gmail.com")
	assert.NoError(t, e.Validate())
	assert.NoError(t, vd.Walk(wrap(e)))
	assert.NoError(t, vd.Walk(wrap(wrap(e))))

	e = email("invalid|gmail.com")
	assert.Error(t, e.Validate())
	assert.Error(t, vd.Walk(wrap(e)))
	assert.Error(t, vd.Walk(wrap(wrap(e))))
}
