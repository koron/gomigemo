package migemo

import (
	"regexp"
)

type Migemo interface {
	Matcher(string) (Matcher, error)
	Pattern(string) (string, error)
	Regexp(string) (*regexp.Regexp, error)
}

type Matcher interface {
	Match(string) (chan Match, error)
	Pattern() (string, error)
	Regexp() (*regexp.Regexp, error)
}

type Match struct {
	Start, End int
}

func Load(path string) (Migemo, error) {
	// TODO:
	return nil, nil
}
