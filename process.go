package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/otiai10/gosseract/v2"
)

const (
	MaxPathLength = 4096       // Linux max path length
	ProcessedTag  = "[tagged]" // Add this add end of the file
)

var logMutex sync.Mutex

func makeUniquePath(dir, base, ext string) string {
	newPath := filepath.Join(dir, base+ext)
	counter := 1
	for {
		if _, err := os.Stat(newPath); os.IsNotExist(err) {
			return newPath
		}
		newPath = filepath.Join(dir, fmt.Sprintf("%s %d%s", base, counter, ext))
		counter++
	}
}

func processImage(filePath string, minScore float64, force bool) {
	if strings.Contains(filePath, ProcessedTag) && !force {
		fmt.Fprintf(os.Stderr, "Skipping already processed file: %s\n", filePath)
		return
	}

	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage(filePath)

	boxes, err := client.GetBoundingBoxesVerbose()

	var str []string
	for _, box := range boxes {
		if w, s := filter(box); s > minScore {
			str = append(str, w)
		}
	}

	text := strings.Join(str, " ")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", filePath, err)
		return
	}

	text = sanitizeFilename(strings.TrimSpace(text))
	if text == "" {
		text = uuid.New().String()
	}

	ext := filepath.Ext(filePath)
	taggedExt := " " + ProcessedTag + ext

	dir := filepath.Dir(filePath)

	newPath := makeUniquePath(dir, text, taggedExt)

	if len(newPath) > MaxPathLength {
		fmt.Fprintf(os.Stderr, "New path exceeds limit: %s\n", newPath)
		newPath = newPath[MaxPathLength-len(taggedExt):] + taggedExt
	}

	err = os.Rename(filePath, newPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to rename '%s': %v\n", filePath, err)
	} else {
		fmt.Fprintf(os.Stderr, "Renamed: '%s' -> '%s'\n", filePath, newPath)
	}
}

func walkDirectory(dir string, minScore float64) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && isImage(path) {
			processImage(path, minScore, false)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Error walking directory: %v", err)
	}
}
