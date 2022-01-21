package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrcmp(t *testing.T) {
	assert.Equal(t, 0, Strcmp("hello", "hello"))
	assert.Equal(t, -32, Strcmp("Hello", "hello"))
	assert.Equal(t, 32, Strcmp("hello", "Hello"))
	assert.Equal(t, -32, Strcmp("Hello\n", "hello"))
	assert.Equal(t, -32, Strcmp("Hello", "hello\n"))

	// different results from the case of C language
	assert.Equal(t, 1, Strcmp("hello\n", "hello"))
	assert.Equal(t, -1, Strcmp("hello", "hello\n"))
}
