package migemo

import (
	"io"
	"os"
	"path/filepath"
)

// AssetProc is a function to proceed asset's io.Reader.
type AssetProc func(io.Reader) error

// Assets provides assets collection, which can be obtained by Get.
type Assets interface {
	Get(name string, proc AssetProc) error
}

// PathAssets is an implementation of Assets interface with physical file
// system (path).
type PathAssets struct {
	root string
}

// Get obtains an asset by name and proceed it with proc.
func (a *PathAssets) Get(name string, proc AssetProc) error {
	path := filepath.Join(a.root, name)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return proc(file)
}
