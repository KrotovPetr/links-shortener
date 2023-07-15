package main

import (
	"flag"
)

var flagRunAddr string

func parseFlags() {
	flag.StringVar(&flagRunAddr, "a", ":8080", "address and port to run server")
	//flag.StringVar(&flagRunAddr, "b", "http://localhost", "host for keeping base url")
	flag.Parse()
}
