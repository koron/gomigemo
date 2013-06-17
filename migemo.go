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

func (g *Migemo) source(s string) (t chan string) {
	err := g.load()
	if err != nil {
		return
	}
	t = make(chan string)
	// TODO: start go routine.
	return
}

func (g *Migemo) Matcher(s string) (m *Matcher) {
	t := g.source(s)
	if t == nil {
		return
	}
	m = NewMatcher()
	for i := range t {
		m.Add(i)
	}
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
