package main

import (
	"flag"
	"github.com/pritunl/pritunl-auth/cmd"
	"github.com/pritunl/pritunl-auth/requires"
	"github.com/pritunl/pritunl-auth/saml"
)

func main() {
	requires.Init()
	flag.Parse()

	err := saml.InitSignCert()
	if err != nil {
		panic(err)
	}

	switch flag.Arg(0) {
	case "app":
		cmd.App()
	}
}
