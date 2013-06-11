package migemo

import (
	"container/list"
)

////////////////////////////////////////////////////////////////////////////
// Ternary Trie Node

type TernaryTrieNode struct {
	ch         rune
	lo, eq, hi *TernaryTrieNode
	Value      interface{}
}

func (n *TernaryTrieNode) Ch() rune {
	return n.ch
}

func (n *TernaryTrieNode) Lo() *TernaryTrieNode {
	return n.lo
}

func (n *TernaryTrieNode) Eq() *TernaryTrieNode {
	return n.eq
}

func (n *TernaryTrieNode) Hi() *TernaryTrieNode {
	return n.hi
}

func (n *TernaryTrieNode) EachWidth(proc func(*TernaryTrieNode)) {
	queue := list.New()
	queue.PushBack(n)
	for queue.Len() != 0 {
		front := queue.Front()
		target := (front.Value).(*TernaryTrieNode)
		queue.Remove(front)
		if target.lo != nil {
			queue.PushBack(target.lo)
		}
		if target.hi != nil {
			queue.PushBack(target.hi)
		}
		proc(target)
	}
	return
}

func (n *TernaryTrieNode) EachWidthInOrder(proc func(*TernaryTrieNode)) {
	if n.lo != nil {
		n.lo.EachWidthInOrder(proc)
	}
	proc(n)
	if n.hi != nil {
		n.hi.EachWidthInOrder(proc)
	}
}

func (n *TernaryTrieNode) Width() (count int) {
	count = 0
	n.EachWidth(func(node *TernaryTrieNode) {
		count++
	})
	return
}

func balance(array []*TernaryTrieNode, s int, e int) (r *TernaryTrieNode) {
	count := e - s
	if count <= 0 {
		r = nil
	} else if count == 1 {
		r = array[s]
	} else if count == 2 {
		array[s].hi = array[s+1]
		r = array[s]
	} else {
		mid := (s + e) / 2
		r = array[mid]
		array[mid].lo = balance(array, s, mid)
		array[mid].hi = balance(array, mid+1, e)
	}
	return
}

func (oldtop *TernaryTrieNode) balance() (newtop *TernaryTrieNode) {
	array := make([]*TernaryTrieNode, oldtop.Width())
	idx := 0
	oldtop.EachWidthInOrder(func(node *TernaryTrieNode) {
		array[idx] = node
		idx++
	})
	for _, node := range array {
		node.lo = nil
		node.hi = nil
	}
	return balance(array, 0, len(array))
}

func (node *TernaryTrieNode) Each(proc func(*TernaryTrieNode)) {
	if node != nil {
		node.lo.Each(proc)
		proc(node)
		node.eq.Each(proc)
		node.hi.Each(proc)
	}
}

func (n *TernaryTrieNode) Find(ch rune) (o *TernaryTrieNode) {
	for n != nil {
		if ch < n.ch {
			n = n.lo
		} else if ch > n.ch {
			n = n.hi
		} else {
			o = n
			break
		}
	}
	return
}

func digRune(p **TernaryTrieNode, ch rune) (nextp **TernaryTrieNode, node *TernaryTrieNode) {
	for {
		if *p == nil {
			*p = &TernaryTrieNode{ch, nil, nil, nil, nil}
		}

		diff := ch - (*p).ch
		if diff == 0 {
			nextp = &(*p).eq
			node = *p
			return
		} else if diff < 0 {
			p = &(*p).lo
		} else {
			p = &(*p).hi
		}
	}
}

////////////////////////////////////////////////////////////////////////////
// Ternary Trie

type TernaryTrie struct {
	root *TernaryTrieNode
}

func NewTernaryTrie() *TernaryTrie {
	return &TernaryTrie{nil}
}

func (t *TernaryTrie) Root() *TernaryTrieNode {
	return t.root
}

func (t *TernaryTrie) Dig(key string) (last *TernaryTrieNode) {
	p := &t.root
	for _, ch := range key {
		p, last = digRune(p, ch)
	}
	return last
}

func (t *TernaryTrie) Add(key string, value interface{}) {
	t.Dig(key).Value = value
}

func (t *TernaryTrie) Find(key string) (node *TernaryTrieNode) {
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

func (t *TernaryTrie) Balance() {
	t.root = t.root.balance()
	t.root.Each(func(node *TernaryTrieNode) {
		if node.eq != nil {
			node.eq = node.eq.balance()
		}
	})
}
