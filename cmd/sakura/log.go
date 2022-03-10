package main

import "github.com/urfave/cli"

var LogCommand = cli.Command{
	Name: "logs",
	Action: func(ctx *cli.Context) error {
		return nil
	},
}
