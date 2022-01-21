/**
 * php `sprintf` with Golang
 *
 * Return a formatted string
 * @see https://www.php.net/manual/en/function.sprintf.php
 *
 * [Note]
 * This is a golang alternative of php `sprintf`, but I recommend `fmt`, golang standard library for formatting variables.
 *
 * @param string format
 * @param ...interface{} args
 * @return string
 */
package functions

import (
	// "regexp"
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Sprintf(format string, args ...interface{}) (string, error) {
	// fmt.doPrintf
	// https://cs.opensource.google/go/go/+/refs/tags/go1.17.5:src/fmt/print.go;drc=refs%2Ftags%2Fgo1.17.5;l=974
	end := len(format)
	var ret bytes.Buffer

	current_index := 0

	for i := 0; i < end; i++ {
		var lasti int
		lasti = i
		for i < end && format[i] != '%' {
			i++
		}
		if lasti < i {
			ret.Write([]byte(format[lasti:i]))
		}
		if i >= end {
			break
		}

		// %---
		start_i := i
		i++ // %

		argidx := -1
		plus_flg := false
		minus_flg := false
		var char byte = 0
		width := -1
		precision := -1

		// 0$
		lasti = i
		for ; i < end; i++ {
			c := format[i]
			if c == '$' {
				if lasti >= i {
					return "", fmt.Errorf("invalid format \"%s\" a", format[lasti:i])
				}
				argidx, _ = strconv.Atoi(format[lasti:i])
				i++
				break
			}
			if !('0' <= c && c <= '9') {
				break
			}
		}
		if argidx < 0 { // not found $
			i = lasti
		}

		if i >= end {
			return "", fmt.Errorf("invalid format \"%s\" b", format[start_i:end])
		}
		// +-
		for ; i < end; i++ {
			c := format[i]
			if c == '+' {
				plus_flg = true
			} else if c == '-' {
				minus_flg = true
			} else {
				break
			}
		}
		if i >= end {
			return "", fmt.Errorf("invalid format \"%s\" c", format[start_i:end])
		}
		// 0 | \x20 | '.
		if format[i] == '0' || format[i] == ' ' {
			char = format[i]
			i++
		} else if format[i] == '\'' {
			i++
			if i >= end {
				return "", fmt.Errorf("invalid format \"%s\" d", format[start_i:end])
			}
			char = format[i]
			i++
		}
		if i >= end {
			return "", fmt.Errorf("invalid format \"%s\" e", format[start_i:end])
		}
		// \d+
		lasti = i
		for ; i < end; i++ {
			c := format[i]
			if !('0' <= c && c <= '9') {
				break
			}
		}
		if lasti < i {
			width, _ = strconv.Atoi(format[lasti:i])
		}
		if i >= end {
			return "", fmt.Errorf("invalid format \"%s\" f", format[start_i:end])
		}
		if width > 0 && char == 0 {
			char = ' '
		}

		// .\d+
		if format[i] == '.' {
			i++
			lasti = i
			for ; i < end; i++ {
				c := format[i]
				if !('0' <= c && c <= '9') {
					break
				}
			}
			if lasti < i {
				precision, _ = strconv.Atoi(format[lasti:i])
			}
			if i >= end {
				return "", fmt.Errorf("invalid format \"%s\" g", format[start_i:end])
			}
		}

		// formatter
		if format[i] == 'l' {
			i++
		} // ignore l
		formatter := format[i]

		// argument
		var idx int
		if argidx > 0 {
			idx = argidx - 1
		} else {
			idx = current_index
			current_index++
		}
		if len(args) <= idx {
			return "", errors.New("Too few arguments")
		}
		a := args[idx]

		// formatting
		var err_ error
		var ret_ string
		switch a.(type) {
		case int, int8, int16, int32, int64:
			v := int64(a.(int))
			ret_, err_ = formatInteger(v, plus_flg, minus_flg, char, width, precision, formatter)
			if err_ != nil {
				return "", err_
			}
		case uint, uint8, uint16, uint32, uint64:
			v := uint64(a.(uint))
			ret_, err_ = formatUnsignedInteger(v, plus_flg, minus_flg, char, width, precision, formatter)
			if err_ != nil {
				return "", err_
			}
		case float32, float64:
			v := a.(float64)
			ret_, err_ = formatFloat(v, plus_flg, minus_flg, char, width, precision, formatter)
			if err_ != nil {
				return "", err_
			}

		case string:
			v := a.(string)
			ret_, err_ = formatString(v, plus_flg, minus_flg, char, width, precision, formatter)
			if err_ != nil {
				return "", err_
			}
		case []byte:
			v := string(a.([]byte))
			ret_, err_ = formatString(v, plus_flg, minus_flg, char, width, precision, formatter)
			if err_ != nil {
				return "", err_
			}

		// doesn't care for complex64, complex128, bool
		default:
			return "", errors.New("Invalid arg")
		}

		ret_, err_ = paddingChars(ret_, plus_flg, minus_flg, char, width, precision, formatter)
		if err_ != nil {
			return "", err_
		}

		ret.Write([]byte(ret_))
	}

	return ret.String(), nil
}

func formatString(v string, lus_flg bool, minus_flg bool, char byte, width int, precision int, f byte) (string, error) {
	var dst string
	switch f {
	case 's':
		dst = v
	case 'c':
		dst = string(v[0])
	case 'd':
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return "", err
		}
		dst = strconv.FormatInt(i, 10)
	case 'u':
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return "", err
		}
		dst = strconv.FormatUint(uint64(i), 10) // Itoa
	case 'b':
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return "", err
		}
		dst = strconv.FormatInt(i, 2)
	case 'o':
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return "", err
		}
		dst = strconv.FormatInt(i, 8)
	case 'x', 'X':
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return "", err
		}
		dst = strconv.FormatInt(i, 16)
		if f == 'X' {
			dst = strings.ToUpper(dst)
		}
	case 'f', 'F', 'e', 'E', 'g', 'G':
		if precision < 0 {
			precision = 6
		}
		if f == 'F' {
			f = 'f'
		}
		i, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return "", err
		}
		dst = normalizeExponent(strconv.FormatFloat(i, f, precision, 64))
	default:
		return "", fmt.Errorf("Invalid Format '%c' for String", f)
	}

	return dst, nil
}
func formatInteger(v int64, plus_flg bool, minus_flg bool, char byte, width int, precision int, f byte) (string, error) {
	var dst string
	switch f {
	case 'd', 's':
		dst = strconv.FormatInt(v, 10) // Itoa
	case 'u':
		dst = strconv.FormatUint(uint64(v), 10) // Itoa
	case 'b':
		dst = strconv.FormatInt(v, 2)
	case 'o':
		dst = strconv.FormatInt(v, 8)
	case 'x', 'X':
		dst = strconv.FormatInt(v, 16)
		if f == 'X' {
			dst = strings.ToUpper(dst)
		}
	case 'c':
		if v < 0 || 255 < v {
			return "", fmt.Errorf("%d cannot be converted to char.", v)
		}
		dst = string(byte(v))
	case 'f', 'F', 'e', 'E', 'g', 'G':
		if precision < 0 {
			precision = 6
		}
		if f == 'F' {
			f = 'f'
		}
		dst = normalizeExponent(strconv.FormatFloat(float64(v), f, precision, 64))
	default:
		return "", fmt.Errorf("Invalid Format '%c' for Number", f)
	}
	return dst, nil
}
func formatFloat(v float64, plus_flg bool, minus_flg bool, char byte, width int, precision int, f byte) (string, error) {
	var dst string
	switch f {
	case 'd', 's':
		dst = strconv.FormatInt(int64(v), 10) // Itoa
	case 'u':
		dst = strconv.FormatUint(uint64(v), 10) // Itoa
	case 'b':
		dst = strconv.FormatInt(int64(v), 2)
	case 'o':
		dst = strconv.FormatInt(int64(v), 8)
	case 'x', 'X':
		dst = strconv.FormatInt(int64(v), 16)
		if f == 'X' {
			dst = strings.ToUpper(dst)
		}
	case 'c':
		if v < 0 || 255 < v {
			return "", fmt.Errorf("%f cannot be converted to char.", v)
		}
		dst = string(byte(v))
	case 'f', 'F', 'e', 'E', 'g', 'G':
		if precision < 0 {
			precision = 6
		}
		if f == 'F' {
			f = 'f'
		}
		dst = normalizeExponent(strconv.FormatFloat(v, f, precision, 64))
	default:
		return "", fmt.Errorf("Invalid Format '%c' for Number", f)
	}

	return dst, nil
}

func formatUnsignedInteger(v uint64, plus_flg bool, minus_flg bool, char byte, width int, precision int, f byte) (string, error) {
	return formatInteger(int64(v), plus_flg, minus_flg, char, width, precision, f)
}

func normalizeExponent(s string) string {
	// strip "0" after "e+" or "e-"
	len_ := len(s)
	b := make([]byte, len_)
	idx := 0
	for i := 0; i < len_; i++ {
		if i+2 < len_ && (s[i] == 'e' || s[i] == 'E') && (s[i+1] == '+' || s[i+1] == '-') && s[i+2] == '0' {
			b[idx] = s[i]
			b[idx+1] = s[i+1]
			i += 2 // skip "0"
			idx += 2
		} else {
			b[idx] = s[i]
			idx++
		}
	}
	return string(b[0:idx])
}

func paddingChars(v string, plus_flg bool, minus_flg bool, char byte, width int, precision int, f byte) (string, error) {
	len_ := len(v)
	if len_ <= 0 {
		return "", nil
	}
	// allocate max size
	size := len(v) + 1
	if size < width {
		size = width
	}
	res := make([]byte, size)

	is_number := f == 'd' || f == 'e' || f == 'E' || f == 'f' || f == 'F' || f == 'g' || f == 'G'

	if is_number && minus_flg && char == '0' {
		if strings.Index(v, ".") < 0 {
			char = ' '
		}
	}

	pos := 0
	if plus_flg {
		if is_number && v[0] != '-' {
			res[pos] = '+'
			pos++
			width--
		}
	}
	if minus_flg == false {
		if width > len_ {
			for i := 0; i < width-len_; i++ {
				res[pos] = char
				pos++
			}
		}
	}
	for i := 0; i < len_; i++ {
		res[pos] = v[i]
		pos++
	}
	if minus_flg == true {
		if width > len_ {
			for i := 0; i < width-len_; i++ {
				res[pos] = char
				pos++
			}
		}
	}

	return string(res[0:pos]), nil
}
