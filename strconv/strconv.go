package strconv

import (
	"bytes"
	"github.com/koron/gomigemo/runereader"
	"github.com/koron/gomigemo/tree"
	"io"
)

type Converter struct {
	trie *tree.Trie
}

type entry struct {
	output, remain string
}

func New() *Converter {
	return &Converter{tree.NewTrie()}
}

func (c *Converter) Add(key, output, remain string) {
	c.trie.Put(key, &entry{output, remain})
}

func (c *Converter) Convert(s string) (string, error) {
	var out, pending bytes.Buffer
	r := runereader.New()
	r.PushFront(s)
	n := c.trie.Root()

	for {
		ch, _, err := r.ReadRune()
		if err == io.EOF {
			err = nil
			break
		} else if err != nil {
			return "", nil
		}

		n = n.Find(ch)
		if n == nil {
			pending.WriteRune(ch)
			ch2, _, err := pending.ReadRune()
			if err == nil {
				out.WriteRune(ch2)
				r.PushFront(pending.String())
				pending.Reset()
			} else if err != io.EOF {
				return "", err
			}
		} else if n.Value != nil {
			e := n.Value.(*entry)
			if len(e.output) > 0 {
				out.WriteString(e.output)
			}
			if len(e.remain) > 0 {
				r.PushFront(e.remain)
			}
			pending.Reset()
			n = c.trie.Root()
		} else {
			pending.WriteRune(ch)
			n = n.Eq()
		}
	}

	if pending.Len() > 0 {
		out.WriteString(pending.String())
	}
	return out.String(), nil
}
