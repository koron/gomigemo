package main

import (
	"fmt"
	"github.com/koron/gomigemo/migemo"
	"log"
)

func main() {
	dict, err := migemo.Load("./_dict")
	if err != nil {
		log.Fatal(err)
	}
	pat, err := migemo.Pattern(dict, "ai")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("pattern=%s\n", pat)
}
