package main

import (
	"io"
	"log"

	"github.com/koron/gomigemo/rpc/client"
)

func query(c *client.Client, s string) (string, error) {
	rx, err := c.Compile(s)
	if err != nil {
		return "", err
	}
	return rx.String(), nil
}

func queryLoop(v View, c *client.Client) {
	for {
		q, err := v.GetQuery()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Print(err)
			continue
		}
		p, err := query(c, q)
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
	c, err := client.Connect()
	if err != nil {
		log.Fatal(err)
	}
	queryLoop(NewConsole(), c)
}
