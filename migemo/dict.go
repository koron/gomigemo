package migemo

import (
	"errors"
	"github.com/koron/gomigemo/conv"
	skkdict "github.com/koron/gomigemo/dict"
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
	skk, err := skkdict.LoadSKK(filepath.Join(d.path, "SKK-JISYO.utf-8.L"))
	if err != nil {
		return err
	}
	roma2hira, err := conv.LoadFile(filepath.Join(d.path, "roma2hira.txt"))
	if err != nil {
		return err
	}
	hira2kata, err := conv.LoadFile(filepath.Join(d.path, "hira2kata.txt"))
	if err != nil {
		return err
	}
	wide2narrow, err := conv.LoadFile(filepath.Join(d.path, "wide2narrow.txt"))
	if err != nil {
		return err
	}

	// Build inflator.
	d.inflator = inflator.Join(
		inflator.DispatchEcho(
			inflator.Join(
				roma2hira,
				inflator.DispatchEcho(inflator.Join(
					hira2kata,
					inflator.DispatchEcho(wide2narrow),
				)),
			),
		),
		inflator.DispatchEcho(skk),
	)

	// FIXME: Make these (loader and builder) flexible.
	return nil
}
