package main

import (
	"bufio"
	"fmt"
	"github.com/koron/gomigemo/migemo"
	"log"
	"os"
	"strings"
)

func dictdir() string {
	return "./_dict"
}

func setupOption(m migemo.Matcher) {
	o := m.GetOptions()
	o.OpWSpaces = ""
	m.SetOptions(o)
}

func query(d migemo.Dict, s string) (string, error) {
	m, err := d.Matcher(s)
	if err != nil {
		return "", err
	}
	setupOption(m)
	p, err := m.Pattern()
	if err != nil {
		return "", err
	}
	return p, nil
}

func readInput(ch chan<- string) {
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("QUERY: ")
		l, err := r.ReadString('\n')
		if l != "" {
			ch <- strings.TrimSpace(l)
		}
		if err != nil {
			return
		}
	}
}

func openQueries() <-chan string {
	ch := make(chan string, 1)
	go func() {
		readInput(ch)
		close(ch)
	}()
	return ch
}

func main() {
	dict, err := migemo.Load(dictdir())
	if err != nil {
		log.Fatal(err)
	}
	ch := openQueries()
	for q := range ch {
		p, err := query(dict, q)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("PATTERN: %s\n", p)
	}
}
