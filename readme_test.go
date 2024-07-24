package validol_test

import (
	"errors"
	"testing"

	"github.com/cospectrum/validol"
	"github.com/stretchr/testify/assert"
)

type Sex string

func (s Sex) Validate() error {
	isSex := validol.OneOf("male", "female", "other")
	return isSex(string(s))
}

type Email string

func (e Email) Validate() error {
	return validol.Email(string(e))
}

type Info struct {
	Email email
	Sex   Sex
	age   uint
}

func (i Info) Validate() error {
	if i.age < 18 {
		return errors.New("18+")
	}
	return validol.Visit(i)
}

func run() {
	var i Info
	if err := i.Validate(); err != nil {
		panic(err)
	}
}

func TestReadme(t *testing.T) {
	defer func() {
		assert.NotNil(t, recover())
	}()
	run()
}
