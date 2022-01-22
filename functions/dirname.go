package functions

import "path/filepath"

// Returns a parent directory's path
func Dirname(p string, args ...int) string {
	if p == "/" || p == "." || p == "" {
		return p[:]
	}

	if p[len(p)-1:] == "/" {
		p = p[:len(p)-1]
	}
	levels := 1
	if len(args) > 0 {
		levels = args[0]
	}
	for levels > 0 {
		levels--
		p, _ = filepath.Split(p)

		if p == "/" || p == "." || p == "" {
			return p[:]
		}
		if p[len(p)-1:] == "/" {
			p = p[:len(p)-1]
		}
	}
	return p
}
