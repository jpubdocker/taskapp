package main

import (
	"log"

	"github.com/jpubdocker/taskapp/pkg/app/api/cmd/config"
	"github.com/jpubdocker/taskapp/pkg/app/api/cmd/server"
	"github.com/jpubdocker/taskapp/pkg/cli"
)

func main() {
	c := cli.NewCLI("taskapp-api", "The API application of taskapp")
	c.AddCommands(
		server.NewCommand(),
		config.NewCommand(),
	)
	if err := c.Execute(); err != nil {
		log.Fatal(err)
	}
}
