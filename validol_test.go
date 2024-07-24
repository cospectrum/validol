package validol_test

import (
	"testing"

	"github.com/cospectrum/validol"
	"github.com/stretchr/testify/assert"
)

func chain[T any](slices ...[]T) []T {
	out := make([]T, 0)
	for _, slice := range slices {
		out = append(out, slice...)
	}
	return out
}

func TestOneOf(t *testing.T) {
	assert.NoError(t, validol.OneOf(2, 1, 3)(1))
	assert.Error(t, validol.OneOf(2, 3)(1))

	A := []int{2, 1, 3}
	B := []int{5, 4}
	oneOfA := validol.OneOf(A...)
	oneOfB := validol.OneOf(B...)
	oneOfAorB := validol.OneOf(chain(B, A)...)
	for _, a := range A {
		assert.NoError(t, oneOfA(a))
		assert.Error(t, oneOfB(a))

		assert.NoError(t, oneOfAorB(a))
	}
	for _, b := range B {
		assert.Error(t, oneOfA(b))
		assert.NoError(t, oneOfB(b))

		assert.NoError(t, oneOfAorB(b))
	}
}

func TestAll(t *testing.T) {
	A := []int{2, 1, 3}
	B := []int{5, 4, 6}

	oneOfA := validol.All(validol.OneOf(A...))
	oneOfAandA := validol.All(oneOfA, oneOfA)

	oneOfB := validol.All(validol.OneOf(B...))
	oneOfAandB := validol.All(oneOfA, oneOfB)

	oneOfBandB := validol.All(oneOfB, oneOfB)

	for _, a := range A {
		assert.NoError(t, oneOfA(a))
		assert.NoError(t, oneOfAandA(a))

		assert.Error(t, oneOfB(a))
		assert.Error(t, oneOfBandB(a))

		assert.Error(t, oneOfAandB(a))
	}
	for _, b := range B {
		assert.Error(t, oneOfA(b))
		assert.Error(t, oneOfAandA(b))

		assert.NoError(t, oneOfB(b))
		assert.NoError(t, oneOfBandB(b))

		assert.Error(t, oneOfAandB(b))
	}
}

func TestAny(t *testing.T) {
	A := []int{2, 1, 3}
	B := []int{4, 5, 6}

	oneOfA := validol.Any(validol.OneOf(A...))
	oneOfAorA := validol.Any(oneOfA, oneOfA)

	oneOfB := validol.Any(validol.OneOf(B...))
	oneOfBorB := validol.Any(validol.OneOf(B...), oneOfB)

	oneOfAorB := validol.Any(oneOfA, oneOfB)

	for _, a := range A {
		assert.NoError(t, oneOfA(a))
		assert.NoError(t, oneOfAorA(a))

		assert.Error(t, oneOfB(a))
		assert.Error(t, oneOfBorB(a))

		assert.NoError(t, oneOfAorB(a))
	}

	for _, b := range B {
		assert.Error(t, oneOfA(b))
		assert.Error(t, oneOfAorA(b))

		assert.NoError(t, oneOfB(b))
		assert.NoError(t, oneOfBorB(b))

		assert.NoError(t, oneOfAorB(b))
	}
}

type M struct {
	Pub     int
	private int
}

var _ validol.Validatable = M{}
var _ validol.Validatable = &M{}

func (m M) Validate() error {
	return validol.Visit(m)
}

func TestVisit(t *testing.T) {
	m := &M{Pub: 0, private: 1}
	assert.NoError(t, validol.Visit(*m))
	assert.NoError(t, validol.Visit(m))
	assert.NoError(t, validol.Visit(&m))
}

func TestGt(t *testing.T) {
	gt3 := validol.Gt(3)
	assert.NoError(t, gt3(4))
	assert.Error(t, gt3(3))
	assert.Error(t, gt3(2))
}

func TestGte(t *testing.T) {
	gte3 := validol.Gte(3)
	assert.NoError(t, gte3(4))
	assert.NoError(t, gte3(3))
	assert.Error(t, gte3(2))
}

func TestLt(t *testing.T) {
	lt3 := validol.Lt(3)
	assert.Error(t, lt3(4))
	assert.Error(t, lt3(3))
	assert.NoError(t, lt3(2))
}

func TestLte(t *testing.T) {
	lte3 := validol.Lte(3)
	assert.Error(t, lte3(4))
	assert.NoError(t, lte3(3))
	assert.NoError(t, lte3(2))
}

func TestEq(t *testing.T) {
	eq3 := validol.Eq(3)
	assert.Error(t, eq3(4))
	assert.NoError(t, eq3(3))
	assert.Error(t, eq3(2))
}

func TestNe(t *testing.T) {
	ne3 := validol.Ne(3)
	assert.NoError(t, ne3(4))
	assert.Error(t, ne3(3))
	assert.NoError(t, ne3(2))
}

func TestNot(t *testing.T) {
	ne3 := validol.Ne(3)
	eq3 := validol.Not(ne3)
	assert.NoError(t, eq3(3))
	assert.Error(t, eq3(2))
}

func TestLen(t *testing.T) {
	lenLte3 := validol.Len[string](validol.Lte(3))

	assert.Error(t, lenLte3("12345"))
	assert.Error(t, lenLte3("1234"))

	assert.NoError(t, lenLte3("123"))
	assert.NoError(t, lenLte3("12"))
	assert.NoError(t, lenLte3("1"))
	assert.NoError(t, lenLte3(""))
}
