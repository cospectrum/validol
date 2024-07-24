package validol_test

import (
	"errors"
	"testing"

	vd "github.com/cospectrum/validol"
	"github.com/stretchr/testify/assert"
)

type Sex string

func (s Sex) Validate() error {
	return vd.OneOf[Sex]("male", "female", "other")(s)
}

type Email string

func (e Email) Validate() error {
	return vd.All(
		vd.Len[string](vd.All(vd.Gt(5), vd.Lte(100))),
		vd.Email,
	)(string(e))
}

type Info struct {
	Email email
	Sex   Sex
	age   uint
}

func (info Info) Validate() error {
	return errors.Join(
		vd.Walk(info),
		vd.Gte(uint(18))(info.age),
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
