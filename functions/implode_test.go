package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImplode(t *testing.T) {
	assert.Equal(t, "lastname,email,phone", Implode(",", []string{"lastname", "email", "phone"}))
}
