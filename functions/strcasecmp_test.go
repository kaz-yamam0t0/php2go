package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrcasecmp(t *testing.T) {
	assert.Equal(t, 0, Strcasecmp("hello", "hello"))
	assert.Equal(t, 0, Strcasecmp("HeLLo", "hello"))
	assert.Equal(t, 0, Strcasecmp("hello", "hEllO"))
	assert.Equal(t, 17, Strcasecmp("Yello\n", "hello"))
	assert.Equal(t, 17, Strcasecmp("Yello", "hello\n"))

	// different results from the case of C language
	assert.Equal(t, 1, Strcasecmp("hello\n", "Hello"))
	assert.Equal(t, -1, Strcasecmp("Hello", "hello\n"))
}
