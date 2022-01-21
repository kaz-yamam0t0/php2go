package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrncasecmp(t *testing.T) {
	assert.Equal(t, 0, Strncasecmp("hello", "hello", 3))
	assert.Equal(t, 0, Strncasecmp("HeLLo", "hello", 3))
	assert.Equal(t, 0, Strncasecmp("hello", "hEllO", 3))
	assert.Equal(t, 0, Strncasecmp("Yello\n", "hello", 0))
	assert.Equal(t, 17, Strncasecmp("Yello\n", "hello", 6))
	assert.Equal(t, 0, Strncasecmp("Yello", "hello\n", 0))
	assert.Equal(t, 17, Strncasecmp("Yello", "hello\n", 6))

	// different results from the case of C language
	assert.Equal(t, 1, Strncasecmp("hello\n", "Hello", 7))
	assert.Equal(t, -1, Strncasecmp("Hello", "hello\n", 7))
}
