package main

import (
//	"fmt"
	"os"
	"flag"
)

// these should probably be randomized by default
var name = flag.String("name", "app", "Subdomain name, defaults to \"app\")")
var port = flag.String("port", "4444", "Listening port, defaults to 4444")
var domain = flag.String("domain", "\\.xenoethics\\.org$", "domain, defaults to xenoethics.org")

func main() {
	flag.Parse()
	template := 
		"\n$HTTP[\"host\"] =~ \"^" + *name + *domain + 
		"\" {\nproxy.server = (\"\" => ((\"host\" => \"127.0.0.1\", \"port\" => " +
		*port + ")))}\n"
	os.Stdout.WriteString(template)
}
// 	$HTTP["host"] =~ "^go\.xenoethics\.org$" {
//   proxy.server = ("" => (
//   (
//   "host" => "127.0.0.1", "port" => 2048)
//   )
//   )
// }