package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrRepeat(t *testing.T) {
	assert.Equal(t, "-=-=-=-=-=-=-=-=-=-=", StrRepeat("-=", 10))
}
