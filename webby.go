package main

import (
	"fmt"
	"http"
	"os"
	"flag"
)

var port = flag.String("port", "2048", "TCP port for inbound connections")
var quiet = flag.Bool("q", false, "Be quiet, don't print anything")

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	flag.Parse()
	http.HandleFunc("/", handler)
	if !*quiet {
		os.Stdout.WriteString("Listening on port " + *port + "\n")
	}
	http.ListenAndServe(":" + *port, nil)
}