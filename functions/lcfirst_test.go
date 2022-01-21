package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLcfirst(t *testing.T) {
	assert.Equal(t, "hello World!", Lcfirst("Hello World!"))
	assert.Equal(t, "hELLO WORLD!", Lcfirst("HELLO WORLD!"))
	assert.Equal(t, "hello world!", Lcfirst("hello world!"))
	assert.Equal(t, "  Hello world!", Lcfirst("  Hello world!"))
}
