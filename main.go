package main

import (
	"flag"
	"github.com/pritunl/pritunl-link/cmd"
	"github.com/pritunl/pritunl-link/logger"
	"github.com/pritunl/pritunl-link/requires"
)

func main() {
	flag.Parse()

	requires.Init()
	logger.Init()

	switch flag.Arg(0) {
	case "start":
		err := cmd.Start()
		if err != nil {
			panic(err)
		}
		break
	case "add":
		err := cmd.Add(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		break
	case "remove":
		err := cmd.Remove(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		break
	case "clear":
		err := cmd.Clear()
		if err != nil {
			panic(err)
		}
		break
	}
}
