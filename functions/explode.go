package functions

import "strings"

// Split a string by a string
//
// Actually the sypnosis of Explode is like this:
//
//  Explode(separator string, s string[, int limit=-1]) []string
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
