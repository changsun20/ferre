package commands

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

var UninstallCmd = &cli.Command{
	Name:      "uninstall",
	Usage:     "Uninstall a package",
	ArgsUsage: "[package]",
	Action: func(ctx context.Context, c *cli.Command) error {
		fmt.Printf("Simulating uninstallation of package: %s\n", c.Args().First())
		return nil
	},
}
