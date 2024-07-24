package validol_test

import (
	"errors"
	"testing"

	"github.com/cospectrum/validol"
	"github.com/stretchr/testify/assert"
)

type Sex string

func (s Sex) Validate() error {
	return validol.OneOf[Sex]("male", "female", "other")(s)
}

type Email string

func (e Email) Validate() error {
	bound := validol.All(validol.Gt(5), validol.Lte(100))
	return validol.All(
		validol.Len[string](bound),
		validol.Email,
	)(string(e))
}

type Info struct {
	Email email
	Sex   Sex
	age   uint
}

func (info Info) Validate() error {
	return errors.Join(
		validol.Walk(info),
		validol.Gte(uint(18))(info.age),
	)
}

func run() {
	var info Info
	if err := info.Validate(); err != nil {
		panic(err)
	}
}

func TestReadme(t *testing.T) {
	defer func() {
		assert.NotNil(t, recover())
	}()
	run()
}
