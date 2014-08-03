package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Console struct {
	Reader *bufio.Reader
	Writer io.Writer
}

func NewConsole() *Console {
	return &Console{
		Reader: bufio.NewReader(os.Stdin),
		Writer: os.Stdout,
	}
}

func (c *Console) GetQuery() (string, error) {
	fmt.Print("QUERY: ")
	l, err := c.Reader.ReadString('\n')
	return strings.TrimSpace(l), err
}

func (c *Console) PutPattern(p string) error {
	_, err := fmt.Fprintf(c.Writer, "PATTERN: %s\n", p)
	return err
}
