package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrtoupper(t *testing.T) {
	assert.Equal(t, "MARY HAD A LITTLE LAMB AND SHE LOVED IT SO", Strtoupper("Mary Had A Little Lamb and She LOVED It So"))
}
