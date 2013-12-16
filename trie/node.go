package trie

import (
	"container/list"
)

// A node of ternary trie.
type Node struct {
	// Rune of this node.
	ch rune

	// Lower, equality, higher nodes.
	lo, eq, hi *Node

	// User value container.
	Value interface{}
}

// Get a rune of this node.
func (n *Node) Ch() rune {
	return n.ch
}

// Get a lower node.
func (n *Node) Lo() *Node {
	return n.lo
}

// Get an equality node.
func (n *Node) Eq() *Node {
	return n.eq
}

// Get a higher node.
func (n *Node) Hi() *Node {
	return n.hi
}

// Enumerate nodes for direction of width.
func (n *Node) EachWidth(proc func(*Node)) {
	queue := list.New()
	queue.PushBack(n)
	for queue.Len() != 0 {
		front := queue.Front()
		target := (front.Value).(*Node)
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

// Enumerate nodes for direction of width in ascent order.
func (n *Node) EachWidthInOrder(proc func(*Node)) {
	if n != nil {
		n.lo.EachWidthInOrder(proc)
		proc(n)
		n.hi.EachWidthInOrder(proc)
	}
}

// Count sibling nodes.
func (n *Node) Width() (count int) {
	count = 0
	n.EachWidth(func(node *Node) {
		count++
	})
	return
}

func balance(array []*Node, s int, e int) (r *Node) {
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

func (oldtop *Node) balance() (newtop *Node) {
	array := make([]*Node, oldtop.Width())
	idx := 0
	oldtop.EachWidthInOrder(func(node *Node) {
		array[idx] = node
		idx++
	})
	for _, node := range array {
		node.lo = nil
		node.hi = nil
	}
	return balance(array, 0, len(array))
}

// Enumerate all (sibilings and descendant) nodes.
func (node *Node) Each(proc func(*Node)) {
	if node != nil {
		node.lo.Each(proc)
		proc(node)
		node.eq.Each(proc)
		node.hi.Each(proc)
	}
}

// Find a node which have the rune.
func (n *Node) Find(ch rune) (o *Node) {
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

// digRune returns a node, for provided rune.
func digRune(p **Node, ch rune) (nextp **Node, node *Node) {
	for {
		if *p == nil {
			*p = &Node{ch, nil, nil, nil, nil}
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
