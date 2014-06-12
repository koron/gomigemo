/*
Package trie provides ternary trie data structure.

	import "github.com/koron/gomigemo/trie"

	t := trie.NewTrie()
	t.Add("foo", 111)
	t.Add("bar", 222)
	t.Add("baz", 333)

	fmt.Println(t.Find("bar").Value) // 222
*/
package trie
