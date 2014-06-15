package skk

import (
	"github.com/koron/gomigemo/readutil"
)

type Word struct {
	Text string
	Desc string
}

type DictEntry struct {
	Label string
	Words []Word
}

type DictEntryProc func(entry *DictEntry)

func LoadDict(path string, proc DictEntryProc) error {
	return readutil.ReadFileLines(path, func(line string, err error) error {
		entry, err2 := line2entry(line)
		if err2 != nil {
			return err2
		}
		proc(entry)
		return err
	})
}

func line2entry(line string) (entry *DictEntry, err error) {
	// TODO: Convert SKK dict line (string) to entry.
	return nil, nil
}
