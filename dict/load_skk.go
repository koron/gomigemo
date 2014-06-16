package dict

import (
	"github.com/koron/gomigemo/skk"
)

func (d *Dict) LoadSKK(path string) (count int, err error) {
	err = skk.LoadDict(path, func(entry *skk.DictEntry) {
		words := make([]string, len(entry.Words))
		for i, w := range entry.Words {
			words[i] = w.Text
		}
		d.Add(entry.Label, words)
		count++
	})
	return count, err
}

func LoadSKK(path string) (*Dict, error) {
	d := New()
	_, err := d.LoadSKK(path)
	if err != nil {
		return nil, err
	}
	return d, nil
}
