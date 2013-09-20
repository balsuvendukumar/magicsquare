// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.

// A command line tool to generate magic squares of odd size.
package main

import (
	"flag"
	"fmt"
	"github.com/jmcvetta/magicsquare"
	"github.com/kr/pretty"
	"os"
	"strconv"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s size\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "    Must specify an odd integer for size\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		return
	}
	size, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		flag.Usage()
		return
	}
	square, err := magicsquare.MagicSquare(size)
	if err != nil {
		flag.Usage()
		return
	}
	pretty.Println(square)

}
