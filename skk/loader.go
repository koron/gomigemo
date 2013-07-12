package skk

import (
	"bufio"
	"code.google.com/p/mahonia"
	"errors"
	"io"
	"os"
	"strings"
)

type Loader struct {
	path string
}

type Comment struct {
	Message string
}

type Entry struct {
	Key    string
	Values []string
}

const sep string = "/"

func parseJisyoEntry(line string) (r interface{}, err error) {
	items := strings.SplitN(line, " ", 2)
	if len(items) != 2 {
		err = errors.New("an invalid SKK JISYO entry: " + line)
	} else {
		values := strings.Split(strings.Trim(items[1], sep), sep)
		for i, s := range values {
			n := strings.Index(s, ";")
			if n >= 0 {
				values[i] = s[0:n]
			}
		}
		r = Entry{items[0], values}
	}
	return
}

func parseJisyo(c chan interface{}, r *bufio.Reader) {
	defer close(c)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				c <- err
			}
			return
		}
		if strings.HasPrefix(line, ";;") {
			c <- Comment{strings.TrimSpace(line[2:])}
		} else {
			event, err := parseJisyoEntry(line)
			if err != nil {
				c <- err
				return
			} else {
				c <- event
			}
		}
	}
}

func newJisyoChan(r io.Reader) (c chan interface{}) {
	d := mahonia.NewDecoder("enc-jp")
	if d == nil {
		return
	}
	c = make(chan interface{})
	go parseJisyo(c, bufio.NewReader(d.NewReader(r)))
	return
}

func NewLoader(path string) (l *Loader) {
	l = &Loader{path}
	return
}

func (l *Loader) Load() (c chan interface{}) {
	f, err := os.Open(l.path)
	if err != nil {
		return
	}
	c = newJisyoChan(f)
	return
}
