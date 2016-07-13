package godomains

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

var validLine = regexp.MustCompile("^[a-zA-Z].*")

// HunspellWordSource is a WordSource that returns words from a Hunspell dictionary
type HunspellWordSource struct {
	scanner *bufio.Scanner
}

// NewHunspellWordSource creates a new HunspellWordSource
func NewHunspellWordSource(file io.Reader) *HunspellWordSource {
	scanner := bufio.NewScanner(file)
	return &HunspellWordSource{scanner}
}

// Next loads the next word
func (source HunspellWordSource) Next() bool {
	hasToken := source.scanner.Scan()
	if hasToken && !isValidLine(source.scanner.Text()) {
		hasToken = source.Next()
	}
	return hasToken
}

// Word returns the currently loaded word
func (source HunspellWordSource) Word() string {
	line := source.scanner.Text()
	return extractWordFrom(line)
}

// Err returns an error if one happened
func (source HunspellWordSource) Err() error {
	return source.scanner.Err()
}

// extractWordFrom strips all characters from line which are not part of the word
func extractWordFrom(line string) string {
	parts := strings.Split(line, "/")
	return parts[0]
}

func isValidLine(line string) bool {
	return validLine.MatchString(line)
}
