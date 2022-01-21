package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrncmp(t *testing.T) {
	assert.Equal(t, 0, Strncmp("helLo", "hello", 3))
	assert.Equal(t, -32, Strncmp("Hello", "helLo", 3))
	assert.Equal(t, 32, Strncmp("hEllo", "Hello", 3))
	assert.Equal(t, 0, Strncmp("Hello\n", "hello", 0))
	assert.Equal(t, -32, Strncmp("Hello\n", "hello", 6))
	assert.Equal(t, 0, Strncmp("Hello", "hello\n", 0))
	assert.Equal(t, -32, Strncmp("Hello", "hello\n", 6))

	// different results from the case of C language
	assert.Equal(t, 1, Strncmp("hello\n", "hello", 7))
	assert.Equal(t, -1, Strncmp("hello", "hello\n", 7))
}
