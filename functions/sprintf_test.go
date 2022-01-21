package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSprintf(t *testing.T) {

	var _err error
	var _ret string

	// Testcases for some common usages
	_ret, _err = Sprintf("%d monkeys in the %s", 5, "tree")
	assert.NoError(t, _err)
	assert.Equal(t, "5 monkeys in the tree", _ret)

	_ret, _err = Sprintf("%2$d monkeys in the %1$s", "tree", 4)
	assert.NoError(t, _err)
	assert.Equal(t, "4 monkeys in the tree", _ret)

	// Testcases for each formats
	_ret, _err = Sprintf("%b", 43951789)
	assert.NoError(t, _err)
	assert.Equal(t, "10100111101010011010101101", _ret)

	_ret, _err = Sprintf("%c", 65)
	assert.NoError(t, _err)
	assert.Equal(t, "A", _ret)

	_ret, _err = Sprintf("%d", 43951789)
	assert.NoError(t, _err)
	assert.Equal(t, "43951789", _ret)

	_ret, _err = Sprintf("%e", 43951789)
	assert.NoError(t, _err)
	assert.Equal(t, "4.395179e+7", _ret)

	_ret, _err = Sprintf("%u", 43951789)
	assert.NoError(t, _err)
	assert.Equal(t, "43951789", _ret)

	_ret, _err = Sprintf("%u", -43951789)
	assert.NoError(t, _err)
	assert.Equal(t, "18446744073665599827", _ret)

	_ret, _err = Sprintf("%u", -1)
	assert.NoError(t, _err)
	assert.Equal(t, "18446744073709551615", _ret)

	_ret, _err = Sprintf("%f", 43951789)
	assert.NoError(t, _err)
	assert.Equal(t, "43951789.000000", _ret)

	_ret, _err = Sprintf("%o", 43951789)
	assert.NoError(t, _err)
	assert.Equal(t, "247523255", _ret)

	_ret, _err = Sprintf("%s", 43951789)
	assert.NoError(t, _err)
	assert.Equal(t, "43951789", _ret)

	_ret, _err = Sprintf("%x", 43951789)
	assert.NoError(t, _err)
	assert.Equal(t, "29ea6ad", _ret)

	_ret, _err = Sprintf("%X", 43951789)
	assert.NoError(t, _err)
	assert.Equal(t, "29EA6AD", _ret)

	// "l" format is always ignored in php sprintf
	_ret, _err = Sprintf("%lb", 43951789)
	assert.NoError(t, _err)
	assert.Equal(t, "10100111101010011010101101", _ret)

	_ret, _err = Sprintf("%ld", 43951789)
	assert.NoError(t, _err)
	assert.Equal(t, "43951789", _ret)

	_ret, _err = Sprintf("%le", 43951789)
	assert.NoError(t, _err)
	assert.Equal(t, "4.395179e+7", _ret)

	_ret, _err = Sprintf("%lu", 43951789)
	assert.NoError(t, _err)
	assert.Equal(t, "43951789", _ret)

	_ret, _err = Sprintf("%lu", -43951789)
	assert.NoError(t, _err)
	assert.Equal(t, "18446744073665599827", _ret)

	_ret, _err = Sprintf("%lu", -1)
	assert.NoError(t, _err)
	assert.Equal(t, "18446744073709551615", _ret)

	_ret, _err = Sprintf("%lf", 43951789)
	assert.NoError(t, _err)
	assert.Equal(t, "43951789.000000", _ret)

	_ret, _err = Sprintf("%lo", 43951789)
	assert.NoError(t, _err)
	assert.Equal(t, "247523255", _ret)

	_ret, _err = Sprintf("%ls", 43951789) // l is ignored even with s
	assert.NoError(t, _err)
	assert.Equal(t, "43951789", _ret)

	_ret, _err = Sprintf("%lx", 43951789)
	assert.NoError(t, _err)
	assert.Equal(t, "29ea6ad", _ret)

	_ret, _err = Sprintf("%lX", 43951789)
	assert.NoError(t, _err)
	assert.Equal(t, "29EA6AD", _ret)

	// Testcases for "+"
	_ret, _err = Sprintf("%+b", 43951789)
	assert.NoError(t, _err)
	assert.Equal(t, "10100111101010011010101101", _ret)

	_ret, _err = Sprintf("%+c", 65) // no "+" with chars or strings
	assert.NoError(t, _err)
	assert.Equal(t, "A", _ret)

	_ret, _err = Sprintf("%+d", 43951789)
	assert.NoError(t, _err)
	assert.Equal(t, "+43951789", _ret)

	_ret, _err = Sprintf("%+e", 43951789)
	assert.NoError(t, _err)
	assert.Equal(t, "+4.395179e+7", _ret)

	_ret, _err = Sprintf("%+u", 43951789)
	assert.NoError(t, _err)
	assert.Equal(t, "43951789", _ret)

	_ret, _err = Sprintf("%+u", -43951789)
	assert.NoError(t, _err)
	assert.Equal(t, "18446744073665599827", _ret)

	_ret, _err = Sprintf("%+u", -1)
	assert.NoError(t, _err)
	assert.Equal(t, "18446744073709551615", _ret)

	_ret, _err = Sprintf("%+f", 43951789)
	assert.NoError(t, _err)
	assert.Equal(t, "+43951789.000000", _ret)

	_ret, _err = Sprintf("%+o", 43951789)
	assert.NoError(t, _err)
	assert.Equal(t, "247523255", _ret)

	_ret, _err = Sprintf("%+s", 43951789)
	assert.NoError(t, _err)
	assert.Equal(t, "43951789", _ret)

	_ret, _err = Sprintf("%+x", 43951789)
	assert.NoError(t, _err)
	assert.Equal(t, "29ea6ad", _ret)

	_ret, _err = Sprintf("%+X", 43951789)
	assert.NoError(t, _err)
	assert.Equal(t, "29EA6AD", _ret)

	// Testcases for numbers
	_ret, _err = Sprintf("%02d", 1)
	assert.NoError(t, _err)
	assert.Equal(t, "01", _ret)

	_ret, _err = Sprintf("%02.2f", 1)
	assert.NoError(t, _err)
	assert.Equal(t, "1.00", _ret)

	_ret, _err = Sprintf("%02.2f", 1.1234)
	assert.NoError(t, _err)
	assert.Equal(t, "1.12", _ret)

	_ret, _err = Sprintf("%06.2f", 1.1234)
	assert.NoError(t, _err)
	assert.Equal(t, "001.12", _ret)

	// Testcases for numbers with general formats
	_ret, _err = Sprintf("%g", 123.4567)
	assert.NoError(t, _err)
	assert.Equal(t, "123.457", _ret)

	_ret, _err = Sprintf("%.5g", 123.456)
	assert.NoError(t, _err)
	assert.Equal(t, "123.46", _ret)

	_ret, _err = Sprintf("%.2g", 1234500000000)
	assert.NoError(t, _err)
	assert.Equal(t, "1.2e+12", _ret)

	_ret, _err = Sprintf("%.2g", "123.45e+10")
	assert.NoError(t, _err)
	assert.Equal(t, "1.2e+12", _ret)

	_ret, _err = Sprintf("%.2g", 1.2345e-8)
	assert.NoError(t, _err)
	assert.Equal(t, "1.2e-8", _ret)

	_ret, _err = Sprintf("%.2g", "123.45e-10")
	assert.NoError(t, _err)
	assert.Equal(t, "1.2e-8", _ret)

	_ret, _err = Sprintf("%.2G", 1234500000000)
	assert.NoError(t, _err)
	assert.Equal(t, "1.2E+12", _ret)

	_ret, _err = Sprintf("%.2G", "123.45e+10")
	assert.NoError(t, _err)
	assert.Equal(t, "1.2E+12", _ret)

	_ret, _err = Sprintf("%.2G", 1.2345e-8)
	assert.NoError(t, _err)
	assert.Equal(t, "1.2E-8", _ret)

	_ret, _err = Sprintf("%.2G", "123.45e-10")
	assert.NoError(t, _err)
	assert.Equal(t, "1.2E-8", _ret)

	// Testcases for numbers with "+"
	_ret, _err = Sprintf("%+02d", 1)
	assert.NoError(t, _err)
	assert.Equal(t, "+1", _ret)

	_ret, _err = Sprintf("%+02.2f", 1)
	assert.NoError(t, _err)
	assert.Equal(t, "+1.00", _ret)

	_ret, _err = Sprintf("%+02.2f", 1.1234)
	assert.NoError(t, _err)
	assert.Equal(t, "+1.12", _ret)

	_ret, _err = Sprintf("%+06.2f", 1.1234)
	assert.NoError(t, _err)
	assert.Equal(t, "+01.12", _ret)

	// Testcases for numbers with "+" and "-"
	_ret, _err = Sprintf("%-+02d", 1)
	assert.NoError(t, _err)
	assert.Equal(t, "+1", _ret)

	_ret, _err = Sprintf("%-+03d", 1) // "0" is evaluated only when the number does not change
	assert.NoError(t, _err)
	assert.Equal(t, "+1 ", _ret)

	_ret, _err = Sprintf("%-+02.2f", 1)
	assert.NoError(t, _err)
	assert.Equal(t, "+1.00", _ret)

	_ret, _err = Sprintf("%-+02.2f", 1.1234)
	assert.NoError(t, _err)
	assert.Equal(t, "+1.12", _ret)

	_ret, _err = Sprintf("%-+06.2f", 1.1234) // "0" is evaluated only when the number does not change
	assert.NoError(t, _err)
	assert.Equal(t, "+1.120", _ret)

	// Testcases for string positions
	_ret, _err = Sprintf("[%s]", "monkey")
	assert.NoError(t, _err)
	assert.Equal(t, "[monkey]", _ret)

	_ret, _err = Sprintf("[%10s]", "monkey")
	assert.NoError(t, _err)
	assert.Equal(t, "[    monkey]", _ret)

	_ret, _err = Sprintf("[%-10s]", "monkey")
	assert.NoError(t, _err)
	assert.Equal(t, "[monkey    ]", _ret)

	_ret, _err = Sprintf("[%010s]", "monkey")
	assert.NoError(t, _err)
	assert.Equal(t, "[0000monkey]", _ret)

	_ret, _err = Sprintf("[%'#10s]", "monkey")
	assert.NoError(t, _err)
	assert.Equal(t, "[####monkey]", _ret)

	_ret, _err = Sprintf("[%-010s]", "monkey")
	assert.NoError(t, _err)
	assert.Equal(t, "[monkey0000]", _ret)

	_ret, _err = Sprintf("[%-'#10s]", "monkey")
	assert.NoError(t, _err)
	assert.Equal(t, "[monkey####]", _ret)

}
