package validol_test

import (
	"testing"
	"validol"

	"github.com/stretchr/testify/assert"
)

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

func chain[T any](slices ...[]T) []T {
	out := make([]T, 0)
	for _, slice := range slices {
		out = append(out, slice...)
	}
	return out
}
