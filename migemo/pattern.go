package migemo

import (
	"bytes"
	"container/list"
	"regexp"
	"unicode/utf8"

	"github.com/koron/gelatin/trie"
)

func (m *matcher) Pattern() (pattern string, err error) {
	if m.patterned {
		return m.pattern, nil
	}
	b := new(bytes.Buffer)
	err = m.writePattern(b, m.trie.Root())
	if err != nil {
		return "", err
	}
	m.pattern = b.String()
	m.patterned = true
	return m.pattern, nil
}

func (m *matcher) writePattern(b *bytes.Buffer, n trie.Node) error {
	labels, chlidNodes := m.splitLabels(n)
	// Output group in.
	grouped := false
	c0 := utf8.RuneCountInString(labels)
	c1 := chlidNodes.Len()
	if c0+c1 > 1 && c1 > 0 {
		grouped = true
		b.WriteString(m.options.OpGroupIn)
	}
	// Output nodes which doesn't have any children.
	if c0 > 0 {
		if c0 > 1 {
			b.WriteString(m.options.OpClassIn)
			b.WriteString(m.quoteMeta(labels))
			b.WriteString(m.options.OpClassOut)
		} else {
			b.WriteString(m.quoteMeta(labels))
		}
	}
	// Output nodes which have some children.
	if c1 > 0 {
		first := c0 == 0
		for e := chlidNodes.Front(); e != nil; e = e.Next() {
			if !first {
				b.WriteString(m.options.OpOr)
			} else {
				first = false
			}
			child := e.Value.(*trie.TernaryNode)
			b.WriteString(m.quoteMeta(string(child.Label())))
			b.WriteString(m.options.OpWSpaces)
			m.writePattern(b, child)
		}
	}
	// Output group out.
	if grouped {
		b.WriteString(m.options.OpGroupOut)
	}
	return nil
}

// splitLabels split children which have children or not.
func (m *matcher) splitLabels(n trie.Node) (label string, nodes *list.List) {
	l := list.New()
	b := new(bytes.Buffer)
	n.Each(func(t trie.Node) bool {
		if t.HasChildren() {
			l.PushBack(t)
		} else {
			b.WriteRune(t.Label())
		}
		return true
	})
	return b.String(), l
}

func (m *matcher) quoteMeta(s string) string {
	// Quote regexp meta chars.
	return regexp.QuoteMeta(s)
}
