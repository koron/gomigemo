package conv

import (
	"github.com/koron/gelatin/trie"
	"github.com/koron/gomigemo/inflator"
)

// Inflate inflates all convesion patterns.
func (c *Converter) Inflate(s string) <-chan string {
	return inflator.Start(func(ch chan<- string) {
		c.convert2(s, func(core, remain string, n trie.Node) {
			extend := false
			if n != c.trie.Root() {
				recursiveEach(n, func(m trie.Node) {
					if e, ok := m.Value().(*entry); ok && e.output != "" {
						ch <- core + e.output
						extend = true
					}
				})
			}
			if !extend {
				ch <- core
			}
		})
	})
}

func recursiveEach(n trie.Node, proc func(trie.Node)) {
	n.Each(func(m trie.Node) bool {
		proc(m)
		recursiveEach(m, proc)
		return true
	})
}
