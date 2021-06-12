package dict

import (
	"io"

	"github.com/koron/go-skkdict"
)

func addDictEntry(d *Dict, entry *skkdict.Entry) {
	words := make([]string, len(entry.Words))
	for i, w := range entry.Words {
		words[i] = w.Text
	}
	d.Add(entry.Label, words)
}

// ReadSKK reads a SKK dictionary from io.Reader.
func ReadSKK(rd io.Reader) (d *Dict, err error) {
	d = New()
	r := skkdict.NewReader(rd)
	for {
		n, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		if n != nil {
			addDictEntry(d, n)
		}
	}
	return d, nil
}
