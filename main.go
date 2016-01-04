// Termcat copies lines from a terminal to stdout. Its intended use
// is as a debug console for hardware development.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pkg/term"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("term:")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: termcat tty\n")
		os.Exit(2)
	}
	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
	}

	t, err := term.Open(flag.Arg(0), term.Speed(57600), term.RawMode)
	if err != nil {
		log.Fatal(err)
	}
	defer t.Close()

	s := bufio.NewScanner(t)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}
