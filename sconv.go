package migemo

import (
    "bytes"
    "io"
    "strings"
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

func (c *SConv) Add(k string, e *SConvEntry) {
    c.trie.Add(k, e)
    return
}

func (c *SConv) Convert(s string) (d string) {
    out := new(bytes.Buffer)
    reader := strings.NewReader(s)
    remain := new (bytes.Buffer)
    //node := nil
    for {
        ch, _, err := reader.ReadRune()
        if (err == io.EOF) {
            break
        }
        out.WriteRune(ch)
    }
    if remain.Len() > 0 {
        out.Write(remain.Bytes())
    }
    d = out.String()
    return
}
