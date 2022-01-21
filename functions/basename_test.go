package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasename(t *testing.T) {
	assert.Equal(t, "sudoers", Basename("/etc/sudoers.d", ".d"))
	assert.Equal(t, "sudoers.d", Basename("/etc/sudoers.d"))
	assert.Equal(t, "passwd", Basename("/etc/passwd"))
	assert.Equal(t, "etc", Basename("/etc/"))
	assert.Equal(t, ".", Basename("."))
	assert.Equal(t, "", Basename("/"))
}
