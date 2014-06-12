package migemo

import (
	"bytes"
	"container/list"
	"github.com/koron/gomigemo/trie"
	"regexp"
)

type Matcher struct {
	trie       *trie.Trie
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
	m = &Matcher{trie.NewTrie(), "|", "(?:", ")", "[", "]", "\\s*"}
	return
}

func (m *Matcher) Add(s string) {
	n := m.trie.Dig2(s, func(m *trie.Node) (r bool) {
		if m.Value == nil {
			r = true
		}
		return
	})
	if n != nil {
		n.Value = true
		n.ResetEq()
	}
}

func (m *Matcher) Match(s string) (r *Match) {
	// TODO: make direct match.
	return
}

func split(n *trie.Node) (s string, l *list.List) {
	l = list.New()
	buf := new(bytes.Buffer)
	n.EachWidthInOrder(func(node *trie.Node) {
		if node.Eq() != nil {
			l.PushBack(node)
		} else {
			buf.WriteRune(node.Ch())
		}
	})
	s = buf.String()
	return
}

func (m *Matcher) outputPattern(buf *bytes.Buffer, n *trie.Node) {
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
			n2 := e.Value.(*trie.Node)
			buf.WriteRune(n2.Ch())
			buf.WriteString(m.opWSpaces)
			m.outputPattern(buf, n2.Eq())
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
