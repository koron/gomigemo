package main

import (
	"github.com/koron/gomigemo/migemo"
	"log"
)

func main() {
	dict, err := migemo.Load("../dict")
	if err != nil {
		log.Fatal(err)
	}
	re, err := migemo.Compile(dict, "kensaku")
	if err != nil {
		log.Fatal(err)
	}
	if re.MatchString("検索") {
		fmt.Println(`"検索" is matched as "kensaku"`)
	} else {
		fmt.Println(`"検索" isn't matched as "kensaku"`)
	}
}
