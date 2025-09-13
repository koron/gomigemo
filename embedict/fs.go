//go:build go1.16
// +build go1.16

package embedict

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
)

//go:embed _dict
var dictFS embed.FS

var DictFS fs.FS

func init() {
	fs, err := fs.Sub(dictFS, "_dict")
	if err != nil {
		panic(err)
	}
	DictFS = fs
}

func Asset(name string) ([]byte, error) {
	f, err := DictFS.Open(name)
	if err != nil {
		return nil, fmt.Errorf("Asset %q not found: %w", name, err)
	}
	defer f.Close()
	return io.ReadAll(f)
}
