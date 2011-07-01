package main

import (
	"os"
	"flag"
)

var ( // Process flags, save logic state.
	printNewlineAtTheEnd = flag.Bool("n", true, "Print final newline")

	)

const (
	Space = " "
	Newline = "\n"
	Tab = "\t"
)

func main() {
	// check ~~flag args~~ state variables.
	flag.Parse()
	// initialize buffer
	var m string = ""
	// iterate through non-flag args (words)
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			//if not at the very beginning, add separator
			m += Space
		}
		//append each to buffer
		m += flag.Arg(i)
	}


	// either print newline or don't
	if *printNewlineAtTheEnd {
		m += Newline
		}
	os.Stdout.WriteString(m)
/* Do not return. Main function doesn't return anything, use os.Exit(errnum) */
}