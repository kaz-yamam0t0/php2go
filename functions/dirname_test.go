package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDirname(t *testing.T) {
	assert.Equal(t, "/etc", Dirname("/etc/passwd"))
	assert.Equal(t, "/", Dirname("/etc/"))
	assert.Equal(t, ".", Dirname("."))
	assert.Equal(t, "/usr", Dirname("/usr/local/lib", 2))
}
