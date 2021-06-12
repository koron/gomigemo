package readutil

import (
	"container/list"
	"io"
	"strings"
)

// StackableRuneReader provides io.RuneReader which can be pushed back strings.
type StackableRuneReader struct {
	readers *list.List
}

// NewStackabeRuneReader creates a new StackableRuneReader instance.
func NewStackabeRuneReader() *StackableRuneReader {
	return &StackableRuneReader{list.New()}
}

// PushFront pushes back a string.
func (r *StackableRuneReader) PushFront(s string) {
	if len(s) > 0 {
		r.readers.PushFront(strings.NewReader(s))
	}
}

// ReadRune reads a rune, which implements io.RuneReader.
func (r *StackableRuneReader) ReadRune() (ch rune, size int, err error) {
	for r.readers.Len() > 0 {
		front := r.readers.Front()
		curr := front.Value.(*strings.Reader)
		ch, size, err = curr.ReadRune()
		if err != io.EOF {
			return
		}
		r.readers.Remove(front)
	}
	return 0, 0, io.EOF
}
