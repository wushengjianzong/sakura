package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

const (
	AppName = "sakura"
)

func main() {
	app := cli.NewApp()
	app.Name = AppName

	app.Commands = []cli.Command{
		InitCommand,
		RunCommand,
		CommitCommand,
		ListCommand,
		LogCommand,
		ExecCommand,
		StopCommand,
		RemoveCommand,
		NetworkCommand,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
