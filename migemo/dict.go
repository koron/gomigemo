package migemo

import (
	"github.com/koron/gomigemo/inflator"
)

type dict struct {
	path     string
	inflator inflator.Inflatable
}

func (d *dict) Matcher(s string) (Matcher, error) {
	return newMatcher(d, s)
}

func (d *dict) load() error {
	// TODO: Load dictonaries and build inflator.
	return nil
}
