package migemo

import (
    "bytes"
    "fmt"
    "io"
)

// SConv - String Converter

type SConv struct {
    debug bool
    trie *TernaryTrie
}

type SConvEntry struct {
    Output string
    Remain string
}

func NewSConv() (c *SConv) {
    c = &SConv{ false, NewTernaryTrie() }
    return
}

func (c *SConv) debugf(format string, a ...interface{}) {
    // FIXME: support vargs.
    if c.debug {
        fmt.Printf(format, a)
    }
}

func (c *SConv) debugln(a ...interface{}) {
    // FIXME: support vargs.
    if c.debug {
        fmt.Println(a)
    }
}

func (c *SConv) Add(k string, output string, remain string) {
    c.trie.Add(k, &SConvEntry { output, remain })
}

func (c *SConv) Convert(s string) (d string, err error) {
    out := new(bytes.Buffer)
    reader := NewStackedReader()
    reader.PushFront(s)
    pending := new (bytes.Buffer)
    node := c.trie.Root()
    for {
        ch, _, err := reader.ReadRune()
        if err != nil {
            if err == io.EOF {
                err = nil
            }
            break
        }

        node = node.Find(ch)
        if node == nil {
            pending.WriteRune(ch)
            ch2, _, err := pending.ReadRune()
            if err != nil {
                if err != io.EOF {
                    break
                }
            } else {
                out.WriteRune(ch2)
                reader.PushFront(pending.String())
                pending.Reset()
            }
            node = c.trie.Root()
        } else if node.Value != nil {
            e := node.Value.(*SConvEntry)
            if len(e.Output) > 0 {
                out.WriteString(e.Output)
            }
            if len(e.Remain) > 0 {
                reader.PushFront(e.Remain)
            }
            pending.Reset()
            node = c.trie.Root()
        } else {
            pending.WriteRune(ch)
            node = node.Eq()
        }
    }
    if pending.Len() > 0 {
        s := pending.String()
        out.WriteString(s)
    }
    d = out.String()
    return
}
