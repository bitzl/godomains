package godomains

import "testing"

type testpair struct {
	line        string
	isValid     bool
	description string
}

var lines = []testpair{
	{"something", true, "is valid because it starts with a letter."},
	{"something/123", true, "is valid because it starts with a letter."},
	{"<something", false, "is not valid because it starts with '<'"},
	{"a something", false, "is not valid because it starts with a space"},
}

func TestIfSlashesEndWords(t *testing.T) {
	line := "test/s"
	word := extractWordFrom(line)
	if word != "test" {
		t.Errorf("Extracting word from line '%s' failed: Expected '%s' but got '%s'.", line, "test", word)
	}
}

func TestIsValid(t *testing.T) {
	for _, pair := range lines {
		if isValidLine(pair.line) != pair.isValid {
			t.Errorf("Line '%s' is %s", pair.line, pair.description)
		}
	}
}
