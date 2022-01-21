package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrPad(t *testing.T) {
	assert.Equal(t, "Alien     ", StrPad("Alien", 10))
	assert.Equal(t, "-=-=-Alien", StrPad("Alien", 10, "-=", STR_PAD_LEFT))
	assert.Equal(t, "__Alien___", StrPad("Alien", 10, "_", STR_PAD_BOTH))
	assert.Equal(t, "Alien_", StrPad("Alien", 6, "___"))
	assert.Equal(t, "Alien", StrPad("Alien", 3, "*"))
	assert.Equal(t, "12341Alien123412", StrPad("Alien", 16, "1234", STR_PAD_BOTH))
}
