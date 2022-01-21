package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddslashes(t *testing.T) {
	assert.Equal(t, "Is your name O\\'Reilly?", Addslashes("Is your name O'Reilly?"))
	assert.Equal(t, "\\\"Web Ninja\\\"", Addslashes("\"Web Ninja\""))
	assert.Equal(t, "\\\"Web\\\\ Ninja\\\"", Addslashes("\"Web\\ Ninja\""))
}
