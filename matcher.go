package migemo

import (
	"bytes"
	"container/list"
	"regexp"
)

type Matcher struct {
	trie       *TernaryTrie
	opOr       string
	opGroupIn  string
	opGroupOut string
	opClassIn  string
	opClassOut string
	opWSpaces  string
}

type Match struct {
	Start, End int
}

func NewMatcher() (m *Matcher) {
	m = &Matcher{NewTernaryTrie(), "|", "(?:", ")", "[", "]", "\\s*"}
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

func split(n *TernaryTrieNode) (s string, l *list.List) {
	l = list.New()
	buf := new(bytes.Buffer)
	n.EachWidthInOrder(func(node *TernaryTrieNode) {
		if node.eq != nil {
			l.PushBack(node)
		} else {
			buf.WriteRune(node.ch)
		}
	})
	s = buf.String()
	return
}

func (m *Matcher) outputPattern(buf *bytes.Buffer, n *TernaryTrieNode) {
	s, l := split(n)
	g := false
	if len(s)+l.Len() > 1 && l.Len() > 0 {
		g = true
	}
	if g {
		buf.WriteString(m.opGroupIn)
	}
	// Output nodes which doesn't have any children.
	if len(s) > 0 {
		if len(s) > 1 {
			buf.WriteString(m.opClassIn)
			buf.WriteString(s)
			buf.WriteString(m.opClassOut)
		} else {
			buf.WriteString(s)
		}
	}
	// Ouput nodes which have some children.
	if l.Len() > 0 {
		if len(s) > 0 {
			buf.WriteString(m.opOr)
		}
		first := true
		for e := l.Front(); e != nil; e = e.Next() {
			if !first {
				buf.WriteString(m.opOr)
			}
			first = false
			n2 := e.Value.(*TernaryTrieNode)
			buf.WriteRune(n2.ch)
			buf.WriteString(m.opWSpaces)
			m.outputPattern(buf, n2.eq)
		}
	}
	if g {
		buf.WriteString(m.opGroupOut)
	}
}

func (m *Matcher) Pattern() (s string) {
	buf := new(bytes.Buffer)
	m.outputPattern(buf, m.trie.Root())
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
