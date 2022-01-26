package functions

import (
	"errors"
	//"bytes"
	"reflect"
)

// phpEqual
func phpEqual(a interface{}, b interface{}, strict_flg bool) bool {
	if strict_flg == false {

	}
	return reflect.DeepEqual(a, b)
}

// Return all the keys or a subset of the keys of an array.
//
// Actually ArrayKeys is called like this:
//
//  ArrayKeys(array interface{} , search_value interface{} [, string_flg bool]) []interface{}
func ArrayKeys(args ...interface{}) ([]interface{}, error) {
	//var values []interface{}
	var keys []interface{}
	var search_value interface{}
	strict_flg := false
	//var err error
	//_ = err
	//_ = strict_flg

	// arguments
	if len(args) <= 0 {
		return nil, errors.New("value is not an array")
	}
	if len(args) >= 2 {
		search_value = args[1]
	}
	if len(args) >= 3 {
		switch a := args[2].(type) {
		case bool:
			strict_flg = a
		}
	}
	_append := func(_keys []interface{}, _k interface{}, _v interface{}) {
		if search_value == nil {
			keys = append(_keys, _k)
			// values = append(values, interface{}(_v))
		} else {
			if phpEqual(_v, search_value, strict_flg) {
				keys = append(_keys, _k)
			}
		}
	}

	switch a := args[0].(type) {
	// []*
	case []int:
		for _i, _v := range a {
			_append(keys, _i, _v)
		}
	case []string:
		for _i, _v := range a {
			_append(keys, _i, _v)
		}
	case []interface{}:
		for _i, _v := range a {
			_append(keys, _i, _v)
		}
	// map[string]*
	case map[string]string:
		for _k, _v := range a {
			_append(keys, _k, _v)
		}
	case map[string]int:
		for _k, _v := range a {
			_append(keys, _k, _v)
		}
	case map[string]interface{}:
		for _k, _v := range a {
			_append(keys, _k, _v)
		}
	default:
		return nil, errors.New("value is not an array")
	}

	return keys, nil
}
