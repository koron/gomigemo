package migemo

import (
	"regexp"
)

type Matcher struct {
}

type Match struct {
	Start, End int
}

func NewMatcher() (m *Matcher) {
	m = &Matcher{}
	return
}

func (m *Matcher) Add(s string) {
	// TODO:
}

func (m *Matcher) Match(s string) (r *Match) {
	// TODO:
	return
}

func (m *Matcher) Pattern() (s string) {
	// TODO:
	return
}

func (m *Matcher) Regexp() (r *regexp.Regexp) {
	// TODO:
	return
}
