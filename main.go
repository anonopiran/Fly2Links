package main

import (
	config "Fly2Links/Config"
	server "Fly2Links/Server"

	"flag"
)

func help() {
	config.Describe()
}
func run() {
	server.Serve()
}
func main() {
	f_help := flag.Bool("env", false, "list available env variables")
	flag.Parse()
	if *f_help {
		help()
	} else {
		run()
	}
}
