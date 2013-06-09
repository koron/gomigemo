package migemo

import (
    "fmt"
    "testing"
)

func checkNode(t *testing.T, n *TernaryTrieNode, ch rune, value int) {
    if (n == nil) {
        t.Fatal("Node is null")
    }
    if (n.Ch() != ch) {
        t.Errorf("Node.Ch() expected:'%c' actual:'%c'", ch, n.Ch())
    }
    if (n.Value.(int) != value) {
        t.Errorf("Node.Value expected:%d actual:%d", value, n.Value)
    }
}

func assertNil(t *testing.T, node *TernaryTrieNode) {
    if node != nil {
        t.Errorf("Not nil: %s", &node)
    }
}

func assertNilBoth(t *testing.T, node *TernaryTrieNode) {
    assertNil(t, node.lo)
    assertNil(t, node.hi)
}

func TestTernaryTrie(t *testing.T) {
    trie := NewTernaryTrie()
    trie.Add("1", 111)
    trie.Add("2", 222)
    trie.Add("3", 333)
    trie.Add("4", 444)
    trie.Add("5", 555)

    n1 := trie.Root()
    checkNode(t, n1, '1', 111)
    n2 := n1.Hi()
    checkNode(t, n2, '2', 222)
    n3 := n2.Hi()
    checkNode(t, n3, '3', 333)
    n4 := n3.Hi()
    checkNode(t, n4, '4', 444)
    n5 := n4.Hi()
    checkNode(t, n5, '5', 555)

    if (n1.Width() != 5) {
        t.Errorf("n1.Width expected:%d actual:%d", 5, n1.Width())
    }
}

func newBalance1(t *testing.T) (trie *TernaryTrie) {
    trie = NewTernaryTrie()
    for i, ch := range "123456789ABCDEF" {
        trie.Add(fmt.Sprintf("%c", ch), i)
    }
    if (trie.Root().Width() != 15) {
        t.Fatalf("Width() expected:15 actual:%d", trie.Root().Width())
    }
    return
}

func checkBalance1(t *testing.T, top *TernaryTrieNode) {
    n8 := top
    checkNode(t, n8, '8', 7)
    n4 := n8.Lo()
    checkNode(t, n4, '4', 3)
    n12 := n8.Hi()
    checkNode(t, n12, 'C', 11)
    n2 := n4.Lo()
    checkNode(t, n2, '2', 1)
    n6 := n4.Hi()
    checkNode(t, n6, '6', 5)
    n10 := n12.Lo()
    checkNode(t, n10, 'A', 9)
    n14 := n12.Hi()
    checkNode(t, n14, 'E', 13)
    n1 := n2.Lo()
    checkNode(t, n1, '1', 0)
    n3 := n2.Hi()
    checkNode(t, n3, '3', 2)
    n5 := n6.Lo()
    checkNode(t, n5, '5', 4)
    n7 := n6.Hi()
    checkNode(t, n7, '7', 6)
    n9 := n10.Lo()
    checkNode(t, n9, '9', 8)
    n11 := n10.Hi()
    checkNode(t, n11, 'B', 10)
    n13 := n14.Lo()
    checkNode(t, n13, 'D', 12)
    n15 := n14.Hi()
    checkNode(t, n15, 'F', 14)
    assertNilBoth(t, n1)
    assertNilBoth(t, n3)
    assertNilBoth(t, n5)
    assertNilBoth(t, n7)
    assertNilBoth(t, n9)
    assertNilBoth(t, n11)
    assertNilBoth(t, n13)
    assertNilBoth(t, n15)
}

func TestTernaryNodeBalance(t *testing.T) {
    trie := newBalance1(t)
    checkBalance1(t, trie.root.balance())
}

func TestTernaryTreeBalance(t *testing.T) {
    // FIXME: use more good data.
    trie := newBalance1(t)
    trie.Balance()
    checkBalance1(t, trie.Root())
}
