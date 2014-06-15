package migemo

import (
	"github.com/koron/gelatin/trie"
)

type matcher struct {
	options   MatcherOptions
	trie      *trie.TernaryTrie
	pattern   string
	patterned bool
}

func newMatcher(d *dict, s string) (*matcher, error) {
	m := &matcher{
		options: defaultMatcherOptions,
		trie:    trie.NewTernaryTrie(),
	}
	// TODO: inflate s word, add those to trie.
	m.trie.Balance()
	return m, nil
}

func (m *matcher) Match(s string) (chan Match, error) {
	// TODO:
	return nil, nil
}

func (m *matcher) SetOptions(o MatcherOptions) {
	m.options = o
	return
}

func (m *matcher) GetOptions() MatcherOptions {
	return m.options
}
