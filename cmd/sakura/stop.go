package main

import "github.com/urfave/cli"

var StopCommand = cli.Command{
	Name: "stop",
	Action: func(ctx *cli.Context) error {
		return nil
	},
}
