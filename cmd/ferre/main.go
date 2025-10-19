package main

import (
	"context"
	"os"

	"github.com/changsun20/ferre/internal/commands"
	"github.com/fatih/color"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:    "ferre",
		Usage:   "A Fast Windows CLI installer",
		Version: "v0.1.0",
		Commands: []*cli.Command{
			commands.InstallCmd,
			commands.UninstallCmd,
			commands.InitCmd,
			commands.BucketCommand,
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		color.Red(err.Error())
	}
}
