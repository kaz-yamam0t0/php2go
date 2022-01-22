package functions

import "path/filepath"

// Returns trailing name component of path
func Basename(p string, args ...string) string {
	if p == "/" || p == "" {
		return ""
	}
	if p[len(p)-1:] == "/" {
		p = p[:len(p)-1]
	}

	_, name := filepath.Split(p)
	if len(args) >= 1 {
		ext := args[0]
		pos := len(name) - len(ext)
		if pos >= 0 && name[pos:] == ext {
			name = name[:pos]
		}
	}

	return name
}
