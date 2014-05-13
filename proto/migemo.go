package migemo

import (
	"regexp"
)

type Migemo interface {
	Matcher(string) (Matcher, error)
}

type Matcher interface {
	Match(string) (chan Match, error)
	Pattern() (string, error)
	SetOptions(MatcherOptions)
	GetOptions() MatcherOptions
}

type MatcherOptions struct {
}

type Match struct {
	Start, End int
}

func Load(path string) (Migemo, error) {
	// TODO:
	return nil, nil
}

func Compile(g Migemo, s string) (*regexp.Regexp, error) {
	m, err := g.Matcher(s)
	if err != nil {
		return nil, err
	}
	return NewRegexp(m)
}

func NewRegexp(m Matcher) (*regexp.Regexp, error) {
	p, err := m.Pattern()
	if err != nil {
		return nil, err
	}
	return regexp.Compile(p)
}
