package strconv

import (
	"fmt"
	"github.com/koron/gomigemo/readutil"
	"io"
	"os"
	"strings"
)

func (c *Converter) LoadFile(path string) (count int, err error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return c.load(file, path)
}

func (c *Converter) load(rd io.Reader, name string) (count int, err error) {
	lnum := 0
	err = readutil.ReadLines(rd, func(line string, err error) error {
		lnum++
		line = strings.TrimRight(line, " \t\r\n")
		if len(line) == 0 || line[0] == '#' {
			return err
		}
		parts := strings.SplitN(line, "\t", 3)
		if parts == nil || len(parts) < 2 {
			return fmt.Errorf("Invalid format in file %s at line %d",
				name, lnum)
		}
		key, emit := parts[0], parts[1]
		var remain string
		if len(parts) >= 3 {
			remain = parts[2]
		}
		c.Add(key, emit, remain)
		count++
		return err
	})
	return count, err
}
