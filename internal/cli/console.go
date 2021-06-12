package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Console provides basic console for gmigemo.
type Console struct {
	Reader *bufio.Reader
	Writer io.Writer
}

// NewConsole allocates a new console.
func NewConsole() *Console {
	return &Console{
		Reader: bufio.NewReader(os.Stdin),
		Writer: os.Stdout,
	}
}

// GetQuery gets a query from user.
func (c *Console) GetQuery() (string, error) {
	fmt.Print("QUERY: ")
	l, err := c.Reader.ReadString('\n')
	return strings.TrimSpace(l), err
}

// PutPattern puts a regexp pattern as migemo result.
func (c *Console) PutPattern(p string) error {
	_, err := fmt.Fprintf(c.Writer, "PATTERN: %s\n", p)
	return err
}
