package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExplode(t *testing.T) {
	assert.Equal(t, []string{"piece1", "piece2", "piece3", "piece4", "piece5", "piece6"}, Explode(" ", "piece1 piece2 piece3 piece4 piece5 piece6"))
	assert.Equal(t, []string{"piece1", "piece2", "piece3 piece4 piece5 piece6"}, Explode(" ", "piece1 piece2 piece3 piece4 piece5 piece6", 3))
	assert.Equal(t, []string{"piece1 piece2 piece3 piece4 piece5 piece6"}, Explode(" ", "piece1 piece2 piece3 piece4 piece5 piece6", 1))
	assert.Equal(t, []string{"piece1 piece2 piece3 piece4 piece5 piece6"}, Explode(" ", "piece1 piece2 piece3 piece4 piece5 piece6", 0))
	assert.Equal(t, []string{"piece1", "piece2", "piece3", "piece4", "piece5"}, Explode(" ", "piece1 piece2 piece3 piece4 piece5 piece6", -1))
	assert.Equal(t, []string{"piece1", "piece2", "piece3"}, Explode(" ", "piece1 piece2 piece3 piece4 piece5 piece6", -3))
}
