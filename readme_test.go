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
		vd.Len[string](vd.Gt(5)),
		vd.Len[string](vd.Lte(100)),
		vd.Email,
	)(string(e))
}

type User struct {
	Email Email
	Sex   Sex
	age   int
}

func (u User) Validate() error {
	return errors.Join(
		vd.Gte(18)(u.age),
		vd.Walk(u), // continue the `Walk` using the children's `Validate` methods
	)
}

func run() {
	users := []User{
		{Email: "first_user@mail.com", age: 22},
		{Email: "second_user@mail.com", Sex: "male"},
	}
	if err := vd.Validate(users); err != nil {
		panic(err)
	}
}

func TestReadme(t *testing.T) {
	defer func() {
		assert.NotNil(t, recover())
	}()
	run()
}
