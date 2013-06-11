package migemo

import (
    "container/list"
    "io"
    "strings"
)

type StackedReader struct {
    readers *list.List
}

func NewStackedReader() (sr *StackedReader) {
    sr = &StackedReader { list.New() }
    return
}

func (sr *StackedReader) PushFront(s string) {
    if len(s) > 0 {
        sr.readers.PushFront(strings.NewReader(s))
    }
}

func (sr *StackedReader) ReadRune() (ch rune, size int, err error) {
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
