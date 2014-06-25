package main

import (
	"fmt"
	"github.com/koron/gomigemo/migemo"
	"log"
)

func dictdir() string {
	return "./_dict"
}

func query(d migemo.Dict, s string) (string, error) {
	return migemo.Pattern(d, s)
}

func main() {
	dict, err := migemo.Load(dictdir())
	if err != nil {
		log.Fatal(err)
	}
	for q := range cl {
		p, err := query(dict, q)
		if err != nil {
			log.Error(err)
			continue
		}
		fmt.Println("PATTERN: %s\n", p)
	}
}
