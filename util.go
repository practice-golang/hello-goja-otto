package main

import (
	"path/filepath"
	"strings"
)

func GetFileExtensions(filename string) (nameWithoutExt, fullExt string) {
	ext := filepath.Ext(filename)
	if ext == "" {
		return filename, ""
	}

	nameWithoutExt = strings.TrimSuffix(filename, ext)

	secondExt := filepath.Ext(nameWithoutExt)

	if secondExt != "" {
		fullExt = secondExt + ext
		nameWithoutExt = strings.TrimSuffix(nameWithoutExt, secondExt)
	} else {
		fullExt = ext
	}

	if fullExt != "" {
		fullExt = fullExt[1:]
	}

	return nameWithoutExt, fullExt
}
