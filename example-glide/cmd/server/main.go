package main

import (
	"os"

	"github.com/amsokol/openshift-golang-template/example-glide/pkg/cmd/cli"
	"github.com/amsokol/openshift-golang-template/example-glide/pkg/cmd/server"
)

func main() {
	cli.EchoArgs()
	if err := server.Start(cli.GetPort()); err != nil {
		os.Exit(1)
	}
}
