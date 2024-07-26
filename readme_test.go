package validol_test

import (
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
		vd.Len[string](vd.Gt(5)),
		vd.Len[string](vd.Lte(100)),
		vd.Email,
	)(string(e))
}

type Info struct {
	Email email
	Sex   Sex
}

func run() {
	var info Info
	if err := vd.Validate(info); err != nil {
		panic(err)
	}
}

func TestReadme(t *testing.T) {
	defer func() {
		assert.NotNil(t, recover())
	}()
	run()
}
