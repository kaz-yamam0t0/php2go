/**
 * Golang equivalent to php `explode`
 *
 * Split a string by a string
 * @see https://www.php.net/manual/en/function.explode.php
 *
 * @param string separator
 * @param string s
 * @param int limit
 * @return interface{}
 */
package functions

import "strings"

func Explode(separator string, s string, args ...int) []string {
	limit := 9223372036854775807
	if len(args) > 0 {
		limit = args[0]
	}
	remove_num := 0
	if limit < 0 {
		remove_num = -limit
		limit = 9223372036854775807
	}
	if limit == 0 {
		limit = 1
	}

	ret := strings.Split(s, separator)

	if len(ret) > limit {
		ret_ := make([]string, limit)

		for i := 0; i < limit-1; i++ {
			ret_[i] = ret[i]
		}
		ret_[limit-1] = strings.Join(ret[limit-1:], separator)

		ret = ret_
	}
	if remove_num > 0 {
		if len(ret) > remove_num {
			return ret[0 : len(ret)-remove_num]
		} else {
			return []string{}
		}
	}
	return ret
}
