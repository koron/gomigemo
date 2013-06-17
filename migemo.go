package migemo

import (
	"regexp"
)

type Migemo struct {
	dictpath string
}

func New(dictpath string) (g *Migemo) {
	g = &Migemo{dictpath}
	return
}

func (g *Migemo) load() (err error) {
	// TODO:
	return
}

func (g *Migemo) Matcher(s string) (m *Matcher) {
	err := g.load()
	if err != nil {
		return
	}
	m = NewMatcher()
	// TODO:
	return
}

func (g *Migemo) Pattern(s string) (p string) {
	m := g.Matcher(s)
	if m != nil {
		p = m.Pattern()
	}
	return
}

func (g *Migemo) Regexp(s string) (r *regexp.Regexp) {
	m := g.Matcher(s)
	if m != nil {
		r = m.Regexp()
	}
	return
}
