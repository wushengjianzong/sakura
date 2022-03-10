package main

import (
	"sakura/pkg/sandbox"

	"github.com/urfave/cli"
)

var InitCommand = cli.Command{
	Name: "init",
	Action: func(ctx *cli.Context) error {
		return sandbox.RunContainerInitProcess()
	},
}
