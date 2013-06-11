package migemo

import (
    "bytes"
    "io"
)

// SConv - String Converter

type SConv struct {
    trie *TernaryTrie
}

type SConvEntry struct {
    Output string
    Remain string
}

func NewSConv() (c *SConv) {
    c = &SConv{ NewTernaryTrie() }
    return
}

func (c *SConv) AddEntry(k string, e *SConvEntry) {
    c.trie.Add(k, e)
    return
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
            ch, _, err := pending.ReadRune()
            if err != nil {
                if err != io.EOF {
                    break
                }
            } else {
                out.WriteRune(ch)
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
        } else {
            pending.WriteRune(ch)
        }
        out.WriteRune(ch)
    }
    if pending.Len() > 0 {
        out.Write(pending.Bytes())
    }
    d = out.String()
    return
}
