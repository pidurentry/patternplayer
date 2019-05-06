package main

import "flag"

var config = struct {
	pattern string
}{}

func init() {
	flag.Parse()

	config.pattern = flag.Arg(0)
	if config.pattern == "" {
		panic("missing pattern argument")
	}
}
