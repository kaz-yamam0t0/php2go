package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChr(t *testing.T) {
	assert.Equal(t, "a", Chr(97))
	assert.Equal(t, "c", Chr(99))

	// Overflow behavior
	//assert.Equal(t, "a", Chr(-159)) // Overflow behavior
	//assert.Equal(t, "A", Chr(833)) // Overflow behavior
}
