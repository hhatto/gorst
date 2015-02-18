package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/hhatto/gorst"
)

var format = flag.String("t", "html", "output format")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [FILE]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	r := os.Stdin
	if flag.NArg() > 0 {
		f, err := os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		r = f
	}

	p := rst.NewParser(nil)
	w := bufio.NewWriter(os.Stdout)
	p.ReStructuredText(r, rst.ToHTML(w))
	w.Flush()
}
