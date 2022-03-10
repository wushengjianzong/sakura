package main

import "github.com/urfave/cli"

var ExecCommand = cli.Command{
	Name: "exec",
	Action: func(ctx *cli.Context) error {
		return nil
	},
}
