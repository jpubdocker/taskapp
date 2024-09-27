package main

import (
	"log"

	"github.com/jpubdocker/taskapp/pkg/app/tools/cmd/mysql"
	"github.com/jpubdocker/taskapp/pkg/cli"
)

func main() {
	c := cli.NewCLI("taskapp-tools", "The utility tools of taskapp")
	c.AddCommands(
		mysql.NewCommand(),
	)
	if err := c.Execute(); err != nil {
		log.Fatal(err)
	}
}
