package main

import "strings"

func pathHasPrefix(path string, prefix string) bool {
	if len(path) == 0 || len(prefix) == 0 || len(prefix) > len(path) {
		return false
	}
	if len(prefix) == len(path) {

		if path == prefix {
			return true
		} else {
			return false
		}
	}

	switch {
	case strings.HasPrefix(prefix, "/"):

		if strings.HasPrefix(path, prefix+"/") {
			return true
		}
		return false
	case strings.HasPrefix(prefix, ".."):
		if strings.HasPrefix(path, prefix+"/") {
			return true
		}
		return false
	case strings.HasPrefix(prefix, "."):
		if strings.HasPrefix(path, "..") || strings.HasPrefix(path, "/") {
			return false
		}
		if !strings.HasPrefix(path, "./") {
			path = "./" + path
		}
		if strings.HasPrefix(path, prefix+"/") {
			return true
		}
		return false

	}
	return false
}
