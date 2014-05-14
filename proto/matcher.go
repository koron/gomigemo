package migemo

type matcher struct {
	options MatcherOptions
}

func newMatcher(d *dict, s string) (*matcher, error) {
	m := &matcher{defaultMatcherOptions}
	// TODO:
	return m, nil
}

func (m *matcher) Match(s string) (chan Match, error) {
	// TODO:
	return nil, nil
}

func (m *matcher) Pattern() (string, error) {
	// TODO:
	return "", nil
}

func (m *matcher) SetOptions(o MatcherOptions) {
	m.options = o
	return
}

func (m *matcher) GetOptions() MatcherOptions {
	return m.options
}
