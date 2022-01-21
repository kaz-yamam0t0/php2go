package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberFormat(t *testing.T) {
	assert.Equal(t, "1,234", NumberFormat(1234))
	assert.Equal(t, "1,235", NumberFormat(1234.56))
	assert.Equal(t, "1 234,56", NumberFormat(1234.56, 2, ",", " ")) // French notation
	assert.Equal(t, "1234.57", NumberFormat(1234.5678, 2, ".", "")) // English notation without thousands separator

	// negative numbers
	assert.Equal(t, "-1,234", NumberFormat(-1234))
	assert.Equal(t, "-1,235", NumberFormat(-1234.56))
	assert.Equal(t, "-1 234,56", NumberFormat(-1234.56, 2, ",", " ")) // French notation
	assert.Equal(t, "-1234.57", NumberFormat(-1234.5678, 2, ".", "")) // English notation without thousands separator
}
