package validol_test

import (
	"testing"
	"validol"

	"github.com/stretchr/testify/assert"
)

func TestOneOf(t *testing.T) {
	assert.NoError(t, validol.OneOf(2, 1, 3)(1))
	assert.Error(t, validol.OneOf(2, 3)(1))
}
