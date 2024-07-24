package validol_test

import (
	"testing"
	"validol"

	"github.com/stretchr/testify/assert"
)

func TestOneOf(t *testing.T) {
	assert.NoError(t, validol.OneOf(2, 1, 3)(1))
	assert.Error(t, validol.OneOf(2, 3)(1))

	a := []int{2, 1, 3}
	b := []int{5, 4}
	oneOfA := validol.OneOf(a...)
	oneOfB := validol.OneOf(b...)
	oneOfAorB := validol.OneOf(chain(b, a)...)
	for _, el := range a {
		assert.NoError(t, oneOfA(el))
		assert.Error(t, oneOfB(el))

		assert.NoError(t, oneOfAorB(el))
	}
	for _, el := range b {
		assert.Error(t, oneOfA(el))
		assert.NoError(t, oneOfB(el))

		assert.NoError(t, oneOfAorB(el))
	}
}

func TestAll(t *testing.T) {
	a := []int{2, 1, 3}

	oneOfA := validol.All(validol.OneOf(a...))
	oneOfAandA := validol.All(oneOfA, oneOfA)
	for _, el := range a {
		assert.NoError(t, oneOfA(el))
		assert.NoError(t, oneOfAandA(el))
	}
}

func chain[T any](slices ...[]T) []T {
	out := make([]T, 0)
	for _, slice := range slices {
		out = append(out, slice...)
	}
	return out
}
