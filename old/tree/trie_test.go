package tree

import (
	"fmt"
	"testing"
)

func checkTrieNode(t *testing.T, n *TrieNode, ch rune, value int) {
	if n == nil {
		t.Fatal("TrieNode is null")
	}
	if n.Ch() != ch {
		t.Errorf("TrieNode.Ch() expected:'%c' actual:'%c'", ch, n.Ch())
	}
	if n.Value.(int) != value {
		t.Errorf("TrieNode.Value expected:%d actual:%d", value, n.Value)
	}
}

func assertNilBoth(t *testing.T, node *TrieNode) {
	if node.lo != nil {
		t.Errorf("lo node has value", &node.lo)
	}
	if node.hi != nil {
		t.Errorf("hi node has value", &node.hi)
	}
}

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Put("1", 111)
	trie.Put("2", 222)
	trie.Put("3", 333)
	trie.Put("4", 444)
	trie.Put("5", 555)

	n1 := trie.Root()
	checkTrieNode(t, n1, '1', 111)
	n2 := n1.Hi()
	checkTrieNode(t, n2, '2', 222)
	n3 := n2.Hi()
	checkTrieNode(t, n3, '3', 333)
	n4 := n3.Hi()
	checkTrieNode(t, n4, '4', 444)
	n5 := n4.Hi()
	checkTrieNode(t, n5, '5', 555)

	if n1.Width() != 5 {
		t.Errorf("n1.Width expected:%d actual:%d", 5, n1.Width())
	}
}

func TestNotFound(t *testing.T) {
	trie := NewTrie()
	if trie.Get("not_exist") != nil {
		t.Errorf("found 'not_exist' in empty trie")
	}
}

func TestBalance(t *testing.T) {
	trie := NewTrie()
	for i, ch := range "123456789ABCDEF" {
		trie.Put(fmt.Sprintf("%c", ch), i)
	}
	if trie.Root().Width() != 15 {
		t.Fatalf("Width() expected:15 actual:%d", trie.Root().Width())
	}

	trie.Balance()

	n8 := trie.Root()
	checkTrieNode(t, n8, '8', 7)
	n4 := n8.Lo()
	checkTrieNode(t, n4, '4', 3)
	n12 := n8.Hi()
	checkTrieNode(t, n12, 'C', 11)
	n2 := n4.Lo()
	checkTrieNode(t, n2, '2', 1)
	n6 := n4.Hi()
	checkTrieNode(t, n6, '6', 5)
	n10 := n12.Lo()
	checkTrieNode(t, n10, 'A', 9)
	n14 := n12.Hi()
	checkTrieNode(t, n14, 'E', 13)
	n1 := n2.Lo()
	checkTrieNode(t, n1, '1', 0)
	n3 := n2.Hi()
	checkTrieNode(t, n3, '3', 2)
	n5 := n6.Lo()
	checkTrieNode(t, n5, '5', 4)
	n7 := n6.Hi()
	checkTrieNode(t, n7, '7', 6)
	n9 := n10.Lo()
	checkTrieNode(t, n9, '9', 8)
	n11 := n10.Hi()
	checkTrieNode(t, n11, 'B', 10)
	n13 := n14.Lo()
	checkTrieNode(t, n13, 'D', 12)
	n15 := n14.Hi()
	checkTrieNode(t, n15, 'F', 14)
	assertNilBoth(t, n1)
	assertNilBoth(t, n3)
	assertNilBoth(t, n5)
	assertNilBoth(t, n7)
	assertNilBoth(t, n9)
	assertNilBoth(t, n11)
	assertNilBoth(t, n13)
	assertNilBoth(t, n15)
}
