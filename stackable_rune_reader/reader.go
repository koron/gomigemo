package stackable_rune_reader

import (
	"container/list"
	"io"
	"strings"
)

type StackableRuneReader struct {
	readers *list.List
}

func New() (sr *StackableRuneReader) {
	sr = &StackableRuneReader{list.New()}
	return
}

func (sr *StackableRuneReader) PushFront(s string) {
	if len(s) > 0 {
		sr.readers.PushFront(strings.NewReader(s))
	}
}

func (sr *StackableRuneReader) ReadRune() (ch rune, size int, err error) {
	for sr.readers.Len() > 0 {
		f := sr.readers.Front()
		r := f.Value.(*strings.Reader)
		ch, size, err = r.ReadRune()
		if err != io.EOF {
			return
		}
		sr.readers.Remove(f)
	}
	err = io.EOF
	return
}
