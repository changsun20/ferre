package commands

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

var InstallCmd = &cli.Command{
	Name:      "install",
	Usage:     "Install a package",
	ArgsUsage: "[package]",
	Action: func(ctx context.Context, c *cli.Command) error {
		fmt.Printf("Simulating installation of package: %s\n", c.Args().First())
		return nil
	},
}
