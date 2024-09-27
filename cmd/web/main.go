package main

import (
	"log"

	"github.com/jpubdocker/taskapp/pkg/app/web/cmd/server"
	"github.com/jpubdocker/taskapp/pkg/cli"
)

func main() {
	c := cli.NewCLI("taskapp-web", "The web application of taskapp")
	c.AddCommands(
		server.NewCommand(),
	)
	if err := c.Execute(); err != nil {
		log.Fatal(err)
	}
}
