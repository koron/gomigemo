package strconv

import (
	"bytes"
	"github.com/koron/gelatin/trie"
	"github.com/koron/gomigemo/readutil"
	"io"
)

type Converter struct {
	trie     *trie.TernaryTrie
	balanced bool
}

type entry struct {
	output, remain string
}

func New() *Converter {
	return &Converter{
		trie:     trie.NewTernaryTrie(),
		balanced: false,
	}
}

func (c *Converter) Add(key, output, remain string) {
	c.trie.Put(key, &entry{output, remain})
	c.balanced = false
}

func (c *Converter) Convert(s string) (string, error) {
	if !c.balanced {
		c.balance()
	}

	var out, pending bytes.Buffer
	r := readutil.NewStackabeRuneReader()
	r.PushFront(s)
	n := c.trie.Root()

	for {
		ch, _, err := r.ReadRune()
		if err == io.EOF {
			err = nil
			break
		} else if err != nil {
			return "", err
		}

		n = n.Get(ch)
		if n == nil {
			pending.WriteRune(ch)
			ch2, _, err := pending.ReadRune()
			if err == nil {
				out.WriteRune(ch2)
				r.PushFront(pending.String())
				pending.Reset()
				n = c.trie.Root()
			} else if err != io.EOF {
				return "", err
			}
		} else if e, ok := n.Value().(*entry); ok {
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
		}
	}

	if pending.Len() > 0 {
		out.WriteString(pending.String())
	}
	return out.String(), nil
}

func (c *Converter) balance() {
	c.trie.Balance()
	c.balanced = true
}
