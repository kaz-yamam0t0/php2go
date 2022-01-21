package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMd5(t *testing.T) {
	assert.Equal(t, "d41d8cd98f00b204e9800998ecf8427e", Md5(""))
	assert.Equal(t, "ab56b4d92b40713acc5af89985d4b786", Md5("abcde"))
	assert.Equal(t, "bd817371d2b9bf5cc9740069695af96d", Md5("今日はいい天気"))
	assert.Equal(t, "be0990750a27212be2dc9e364a9edb2c", Md5("😃😃☀️"))
}
