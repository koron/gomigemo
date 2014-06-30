package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/koron/gomigemo/migemo"
)

var dictdir = flag.String("d", migemo.DefaultDictdir(), "Location to dictionary")

func grep(r io.Reader, re *regexp.Regexp) error {
	buf := bufio.NewReader(r)
	for {
		b, _, err := buf.ReadLine()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		line := string(b)
		if re.MatchString(line) {
			fmt.Println(line)
		}
	}
	return nil
}

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage of main: [pattern] [files...]")
	}

	flag.Parse()

	if flag.NArg() != 1 && flag.NArg() != 2 {
		flag.Usage()
		flag.PrintDefaults()
		os.Exit(1)
	}

	dict, err := migemo.Load(*dictdir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	re, err := migemo.Compile(dict, flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if flag.NArg() == 1 {
		if err = grep(os.Stdin, re); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	} else {
		for _, arg := range flag.Args()[1:] {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			defer f.Close()
			if err = grep(f, re); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
	}
}
