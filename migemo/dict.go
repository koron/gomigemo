package migemo

import (
	"errors"
	"github.com/koron/gomigemo/conv"
	skk "github.com/koron/gomigemo/dict"
	"github.com/koron/gomigemo/inflator"
	"path/filepath"
)

type dict struct {
	path     string
	inflator inflator.Inflatable
}

func (d *dict) Matcher(s string) (Matcher, error) {
	return newMatcher(d, s)
}

func (d *dict) load() error {
	if d.inflator != nil {
		return errors.New("Dictionaries were loaded already.")
	}

	// Load dictionaries.
	k2k, err := skk.LoadSKK(filepath.Join(d.path, "SKK-JISYO.utf-8.L"))
	if err != nil {
		return err
	}
	r2h, err := conv.LoadFile(filepath.Join(d.path, "roma2hira.txt"))
	if err != nil {
		return err
	}

	// Build inflator.
	d.inflator = inflator.Join(
		inflator.Dispatch(inflator.Echo(), r2h),
		inflator.Dispatch(inflator.Echo(), k2k))

	// For debug.
	d.inflator = inflator.Dispatch(inflator.Echo(), r2h)
	d.inflator = r2h

	// FIXME: Make these (loader and builder) flexible.
	return nil
}
