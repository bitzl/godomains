package godomains

// A WordSource yields a new word with every call
type WordSource interface {
	Next() bool
	Word() string
	Err() error
}
