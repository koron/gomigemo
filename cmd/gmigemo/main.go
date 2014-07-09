package main

import (
	"github.com/koron/gomigemo/embedict"
	"github.com/koron/gomigemo/migemo"
	"io"
	"log"
)

func adjustMatcher(m migemo.Matcher) {
	o := m.GetOptions()
	o.OpWSpaces = ""
	m.SetOptions(o)
}

func query(d migemo.Dict, s string) (string, error) {
	m, err := d.Matcher(s)
	if err != nil {
		return "", err
	}
	adjustMatcher(m)
	p, err := m.Pattern()
	if err != nil {
		return "", err
	}
	return p, nil
}

func queryLoop(v View, d migemo.Dict) {
	for {
		q, err := v.GetQuery()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Print(err)
			continue
		}
		p, err := query(d, q)
		if err != nil {
			log.Print(err)
			continue
		}
		err = v.PutPattern(p)
		if err != nil {
			log.Print(err)
			continue
		}
	}
}

func main() {
	dict, err := embedict.Load()
	if err != nil {
		log.Fatal(err)
	}
	queryLoop(NewConsole(), dict)
}
