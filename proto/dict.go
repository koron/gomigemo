package migemo

type dict struct {
	path string
}

func (d *dict) Matcher(s string) (Matcher, error) {
	return newMatcher(d, s);
}

func (d *dict) load() error {
	// TODO:
	return nil
}
