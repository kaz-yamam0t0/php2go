package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNl2br(t *testing.T) {
	assert.Equal(t, "foo<br />\nbar", Nl2br("foo\nbar"))
	assert.Equal(t, "foo<br />\rbar", Nl2br("foo\rbar"))
	assert.Equal(t, "foo<br />\r\nbar", Nl2br("foo\r\nbar"))
	assert.Equal(t, "foo<br />\n\rbar", Nl2br("foo\n\rbar"))
	assert.Equal(t, "foo<br>\nbar", Nl2br("foo\nbar", false))
	assert.Equal(t, "foo<br>\rbar", Nl2br("foo\rbar", false))
	assert.Equal(t, "foo<br>\r\nbar", Nl2br("foo\r\nbar", false))
	assert.Equal(t, "foo<br>\n\rbar", Nl2br("foo\n\rbar", false))
	assert.Equal(t, "foo<br />\n<br />\n<br />\nbar", Nl2br("foo\n\n\nbar"))
	assert.Equal(t, "foo<br />\r<br />\r<br />\rbar", Nl2br("foo\r\r\rbar"))
	assert.Equal(t, "foo<br />\r\n<br />\rbar", Nl2br("foo\r\n\rbar"))
	assert.Equal(t, "foo<br />\n\r<br />\rbar", Nl2br("foo\n\r\rbar"))
	assert.Equal(t, "foo<br />\n\r<br />\nbar", Nl2br("foo\n\r\nbar"))
}
