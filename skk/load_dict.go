package skk

import (
	"fmt"
	"github.com/koron/gomigemo/readutil"
	"strings"
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
		if entry != nil {
			proc(entry)
		}
		return err
	})
}

func line2entry(line string) (entry *DictEntry, err error) {
	// Convert SKK dict line (string) to entry.
	if strings.HasPrefix(line, ";;") {
		return nil, nil
	}
	line = strings.TrimRight(line, " \t\r\n")
	items := strings.SplitN(line, " ", 2)
	if items == nil || len(items) != 2 {
		return nil, fmt.Errorf("Invalid format")
	}
	label := items[0]
	values := strings.Split(strings.Trim(items[1], "/"), "/")
	words := make([]Word, len(values))
	for i, v := range values {
		words[i] = value2word(v)
	}
	return &DictEntry{
		Label: label,
		Words: words,
	}, nil
}

func value2word(v string) Word {
	n := strings.Index(v, ";")
	if n < 0 {
		return Word{Text: v}
	}
	return Word{
		Text: v[0:n],
		Desc: v[n+1:],
	}
}
