package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubstr(t *testing.T) {
	assert.Equal(t, "bcdef", Substr("abcdef", 1))
	assert.Equal(t, "bc", Substr("abcdef", 1, 2))
	assert.Equal(t, "f", Substr("abcdef", -1))
	assert.Equal(t, "d", Substr("abcdef", -3, 1))
}
