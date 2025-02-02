package main

import (
	"path/filepath"
	"regexp"
	"strings"
)

func isImage(file string) bool {
	ext := strings.ToLower(filepath.Ext(file))
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".tiff", ".webp":
		return true
	default:
		return false
	}
}

func sanitizeFilename(str string) string {
	re := regexp.MustCompile(`[?*<>|":/\\]+`)
	return re.ReplaceAllString(str, "_")
}
