package trie

// A ternary trie.
type Trie struct {
	root *Node
}

// NewTrie returns a Trie that represent ternary trie.
func NewTrie() *Trie {
	return &Trie{nil}
}

// Root returns a root node of the trie.
func (t *Trie) Root() *Node {
	return t.root
}

// Dig returns a last node to store provided key.  And creates all intermediate
// nodes if necessary.
func (t *Trie) Dig(key string) (last *Node) {
	p := &t.root
	for _, ch := range key {
		p, last = digRune(p, ch)
	}
	return last
}

// Add key and value.
func (t *Trie) Add(key string, value interface{}) {
	t.Dig(key).Value = value
}

// Find returns a node for provided key.
func (t *Trie) Find(key string) (node *Node) {
	p := t.root
	for _, ch := range key {
		for {
			if p == nil {
				node = nil
				break
			}

			diff := ch - p.ch
			if diff == 0 {
				node = p
				p = p.eq
				break
			} else if diff < 0 {
				p = p.lo
			} else {
				p = p.hi
			}
		}
	}
	return
}

// Balance makes all nodes in trie balanced.
func (t *Trie) Balance() {
	t.root = t.root.balance()
	t.root.Each(func(node *Node) {
		if node.eq != nil {
			node.eq = node.eq.balance()
		}
	})
}
