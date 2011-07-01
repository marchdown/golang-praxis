package main

import (
       "os"
       "flag"
)

var omitNewline = flag.Bool("n", false, "don't print final newline")
var outputTwice = flag.Bool("2", false, "echo two time in succession")

const (
      Space = " "
      Newline = "\n"
)

func main() {
     flag.Parse()
     var s string = ""
     for i := 0; i < flag.NArg(); i++ {
	if i > 0 {
	    s += Space
	    }
	    s += flag.Arg(i)
	 }
	 if *outputTwice {
		s += Space
		s += s
	}
	 if !*omitNewline {
	    s += Newline
	 }
	 os.Stdout.WriteString(s)
}
