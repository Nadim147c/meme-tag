package main

import (
	_ "embed"
	"regexp"
	"slices"
	"strings"
	"sync"

	"github.com/otiai10/gosseract/v2"
)

//go:embed words.txt
var wordsString string
var wordOnce sync.Once

var words []string

func normalizeText(word string) string {
	re := regexp.MustCompile("[^a-zA-Z0-9]")

	word = re.ReplaceAllString(word, "")

	if word == "|" {
		return "I"
	}

	if word == "a" || word == "A" {
		return "a"
	}

	if len(word) == 1 {
		return ""
	}

	return word
}

// filter takes a b gosseract.BoundingBox. Normalizes the word with search and replace.
// Counts a score from b.Confidence. Adds 10 to scores if resulting value is a valid word.
// Returns a normalized word and bool if the score is more than 90
func filter(b gosseract.BoundingBox) (string, float64) {
	socre := b.Confidence

	word := normalizeText(b.Word)

	wordOnce.Do(func() { words = strings.Split(wordsString, "\n") })
	_, exists := slices.BinarySearchFunc(words, strings.ToLower(word), func(a, b string) int {
		return strings.Compare(a, b)
	})

	if exists {
		socre += 15
	}

	if word == "" {
		return word, 0
	}

	return word, socre
}
