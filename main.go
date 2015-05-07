package main

import (
	"flag"
	"github.com/pritunl/pritunl-auth/cmd"
	"github.com/pritunl/pritunl-auth/requires"
)

func main() {
	requires.Init()
	flag.Parse()

	switch flag.Arg(0) {
	case "app":
		cmd.App()
	}
}
