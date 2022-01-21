package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUcfirst(t *testing.T) {
	assert.Equal(t, "Hello world!", Ucfirst("hello world!"))
	assert.Equal(t, "HELLO WORLD!", Ucfirst("HELLO WORLD!"))
	assert.Equal(t, "  hello world!", Ucfirst("  hello world!"))
}
