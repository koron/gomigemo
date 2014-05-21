package strconv

import (
	"github.com/koron/gomigemo/inflator"
)

func (c *Converter) Inflate(s string) <-chan string {
	return inflator.Start(func(ch chan<- string) {
		v, err := c.Convert(s)
		if err != nil {
			return
		}
		ch <- v
	})
}
