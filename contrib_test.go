package validol_test

import (
	"testing"

	vd "github.com/cospectrum/validol"
	"github.com/stretchr/testify/assert"
)

type wrapper[T any] struct {
	Value T
}

func wrap[T any](val T) wrapper[T] {
	return wrapper[T]{
		Value: val,
	}
}

type email string

func (e email) Validate() error {
	return vd.Email(string(e))
}

func TestEmail(t *testing.T) {
	t.Parallel()

	valid := []string{
		"valid@gmail.com",
		"test@mail.com",
		"Dörte@Sörensen.example.com",
		"θσερ@εχαμπλε.ψομ",
		"юзер@екзампл.ком",
		"उपयोगकर्ता@उदाहरण.कॉम",
		"用户@例子.广告",
		`"test test"@email.com`,
	}
	for _, s := range valid {
		e := email(s)
		assert.NoError(t, vd.Email(s))
		assert.NoError(t, e.Validate())
		assert.NoError(t, vd.Walk(wrap(e)))
		assert.NoError(t, vd.Walk(wrap(wrap(e))))
	}
	invalid := []string{
		"",
		"invalid|gmail.com",
		"mail@domain_with_underscores.org",
		"test@email",
		"test@email.",
		"@email.com",
		`"@email.com`,
	}
	for _, s := range invalid {
		e := email(s)
		assert.Error(t, vd.Email(s))
		assert.Error(t, e.Validate())
		assert.Error(t, vd.Walk(wrap(e)))
		assert.Error(t, vd.Walk(wrap(wrap(e))))
	}
}

type uuid4 string

func (u uuid4) Validate() error {
	return vd.UUID4(string(u))
}

type void struct{}

func mapF[T any, U any](elems []T, f func(T) U) []U {
	out := make([]U, 0, len(elems))
	for _, el := range elems {
		out = append(out, f(el))
	}
	return out
}

func TestUUID4(t *testing.T) {
	t.Parallel()

	valid := []string{
		"57b73598-8764-4ad0-a76a-679bb6640eb1",
		"625e63f3-58f5-40b7-83a1-a72ad31acffb",
	}
	for _, s := range valid {
		u := uuid4(s)
		assert.NoError(t, vd.UUID4(s))
		assert.NoError(t, u.Validate())
		assert.NoError(t, vd.Walk(wrap(u)))
		assert.NoError(t, vd.Walk(wrap(wrap(u))))
		assert.NoError(t, vd.Walk(struct{ Val uuid4 }{u}))
		assert.NoError(t, vd.Walk([]uuid4{u}))
		assert.NoError(t, vd.Walk(map[uuid4]void{u: {}}))
		assert.NoError(t, vd.Walk(map[void]uuid4{{}: u}))
	}
	assert.NoError(t, vd.Walk(mapF(valid, func(s string) uuid4 {
		return uuid4(s)
	})))
	assert.NoError(t, vd.Walk(mapF(valid, func(s string) *uuid4 {
		u := uuid4(s)
		return &u
	})))
	invalid := []string{
		"",
		"xxxa987fbc9-4bed-3078-cf07-9141ba07c9f3",
		"a987fbc9-4bed-5078-af07-9141ba07c9f3",
		"934859",
	}
	for _, s := range invalid {
		u := uuid4(s)
		assert.Error(t, vd.UUID4(s))
		assert.Error(t, u.Validate())
		assert.Error(t, vd.Walk(wrap(u)))
		assert.Error(t, vd.Walk(wrap(wrap(u))))
		assert.Error(t, vd.Walk(struct{ Val uuid4 }{u}))
		assert.Error(t, vd.Walk([]uuid4{u}))
		assert.Error(t, vd.Walk(map[uuid4]void{u: {}}))
		assert.Error(t, vd.Walk(map[void]uuid4{{}: u}))
	}
	assert.Error(t, vd.Walk(mapF(invalid, func(s string) uuid4 {
		return uuid4(s)
	})))
	assert.Error(t, vd.Walk(mapF(invalid, func(s string) *uuid4 {
		u := uuid4(s)
		return &u
	})))
}
