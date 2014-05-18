package tree

// A ternary trie.
type Trie struct {
	// root node.
	root *TrieNode
}

// NewTrie returns a Trie that represent ternary trie.
func NewTrie() *Trie {
	return &Trie{nil}
}

// Root returns a root node of the trie.
func (t *Trie) Root() *TrieNode {
	return t.root
}

// Dig return a node to store provided key.  check is callbacked when digged
// and forwarded a node, when it returns false, digging is terminated.
func (t *Trie) Dig(key string, check func(*TrieNode) bool) (*TrieNode, bool) {
	pnode := &t.root
	var node *TrieNode
	for _, ch := range key {
		node = dig(pnode, ch)
		if check != nil && !check(node) {
			return node, false
		}
		pnode = &node.eq
	}
	return node, true
}

// dig returns a node to store a rune.  This create a new node if needed.
func dig(pnode **TrieNode, ch rune) *TrieNode {
	for {
		if *pnode == nil {
			*pnode = &TrieNode{ch, nil, nil, nil, nil}
		}
		diff := ch - (*pnode).ch
		switch {
		case diff == 0:
			return *pnode
		case diff < 0:
			pnode = &(*pnode).lo
		default:
			pnode = &(*pnode).hi
		}
	}
}

// Put key and value.  If already exists, old value is overridden.
func (t *Trie) Put(key string, value interface{}) *TrieNode {
	n, _ := t.Dig(key, nil)
	if n != nil {
		n.Value = value
	}
	return n
}

// Get returns a node for key.
func (t *Trie) Get(key string) (r *TrieNode) {
	n := t.root
	for _, ch := range key {
		for {
			if n == nil {
				return nil
			}
			diff := ch - n.ch
			switch {
			case diff == 0:
				r = n
				n = n.eq
				break
			case diff < 0:
				n = n.lo
			default:
				n = n.hi
			}
		}
	}
	return r
}

// Balance makes all nodes in trie are balanced.
func (t *Trie) Balance() {
	t.root = t.root.balance()
}

// A node of ternary trie.
type TrieNode struct {
	// Rune of this node.
	ch rune

	// Lower, equality, higher nodes.
	lo, eq, hi *TrieNode

	// User value container.
	Value interface{}
}

// Get a rune of this node.
func (n *TrieNode) Ch() rune {
	return n.ch
}

// Get a lower node.
func (n *TrieNode) Lo() *TrieNode {
	return n.lo
}

// Get an equality node.
func (n *TrieNode) Eq() *TrieNode {
	return n.eq
}

// Get a higher node.
func (n *TrieNode) Hi() *TrieNode {
	return n.hi
}

// Enumerate sibling nodes ascent order.
func (n *TrieNode) Each(proc func(*TrieNode)) {
	if n != nil {
		n.lo.Each(proc)
		proc(n)
		n.hi.Each(proc)
	}
}

func (n *TrieNode) Width() int {
	count := 0
	n.Each(func(*TrieNode) {
		count += 1
	})
	return count
}

func (n *TrieNode) siblings() []*TrieNode {
	nodes := make([]*TrieNode, n.Width())
	idx := 0
	n.Each(func(m *TrieNode) {
		nodes[idx] = m
		idx += 1
	})
	return nodes
}

func (n *TrieNode) balance() *TrieNode {
	nodes := n.siblings()
	for _, m := range nodes {
		m.lo = nil
		m.hi = nil
		if m.eq != nil {
			m.eq = m.eq.balance()
		}
	}
	return balance(nodes, 0, len(nodes))
}

func balance(nodes []*TrieNode, s int, e int) *TrieNode {
	count := e - s
	switch {
	case count <= 0:
		return nil
	case count == 1:
		return nodes[s]
	case count == 2:
		nodes[s].hi = nodes[s+1]
		return nodes[s]
	default:
		mid := (s + e) / 2
		nodes[mid].lo = balance(nodes, s, mid)
		nodes[mid].hi = balance(nodes, mid+1, e)
		return nodes[mid]
	}
}
