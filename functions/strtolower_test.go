package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrtolower(t *testing.T) {
	assert.Equal(t, "mary had a little lamb and she loved it so", Strtolower("Mary Had A Little Lamb and She LOVED It So"))
}
