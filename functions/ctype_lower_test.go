package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCtypeLower(t *testing.T) {
	assert.Equal(t, true, CtypeLower("abcde"))
	assert.Equal(t, false, CtypeLower("abcDe"))
	assert.Equal(t, false, CtypeLower("abcde#\x24"))
	assert.Equal(t, false, CtypeLower("abc_de"))
	assert.Equal(t, false, CtypeLower("abc-de"))

	// Irregular arguments
	assert.Equal(t, false, CtypeLower(""))
}
