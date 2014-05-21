package strconv

import (
	"io"
	"os"
)

func (c *Converter) LoadFile(path string) (count int, err error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return c.Load(file)
}

func (c *Converter) Load(rd io.Reader) (count int, err error) {
	return 0, nil
}
