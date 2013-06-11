package migemo

import (
	"bytes"
	"regexp"
)

type Matcher struct {
	trie *TernaryTrie
}

type Match struct {
	Start, End int
}

func NewMatcher() (m *Matcher) {
	m = &Matcher{NewTernaryTrie()}
	return
}

func (m *Matcher) Add(s string) {
	p := &m.trie.root
	var n *TernaryTrieNode
	for _, ch := range s {
		p, n = digRune(p, ch)
		if n.Value != nil {
			return
		}
	}
	n.Value = true
	n.eq = nil
	return
}

func (m *Matcher) Match(s string) (r *Match) {
	// TODO:
	return
}

func (m *Matcher) Pattern() (s string) {
	buf := new(bytes.Buffer)
	// TODO:
	s = buf.String()
	return
}

func (m *Matcher) Regexp() (r *regexp.Regexp) {
	r, err := regexp.Compile(m.Pattern())
	if err != nil {
		r = nil
	}
	return
}
