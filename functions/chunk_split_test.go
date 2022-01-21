package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChunkSplit(t *testing.T) {
	assert.Equal(t, "abcd\r\nefgh\r\nijkl\r\nmn\r\n", ChunkSplit("abcdefghijklmn", 4))
	assert.Equal(t, "abcd\r\nefgh\r\nijkl\r\nmnop\r\n", ChunkSplit("abcdefghijklmnop", 4))
	assert.Equal(t, "abcd\n\nefg\nhijk\nlmn\n", ChunkSplit("abcd\nefghijklmn", 4, "\n"))
}
