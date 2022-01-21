package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUcwords(t *testing.T) {
	assert.Equal(t, "Hello World!", Ucwords("hello world!"))
	assert.Equal(t, "HELLO WORLD!", Ucwords("HELLO WORLD!"))
	assert.Equal(t, "HELLO ..woRLD!", Ucwords("HELLO ..woRLD!"))
	assert.Equal(t, "  Hello World!", Ucwords("  hello world!"))

	// Separators
	assert.Equal(t, "Hello|World!", Ucwords("hello|world!", "|"))
	assert.Equal(t, "|Hello_World!", Ucwords("|hello_world!", "_|"))
}
