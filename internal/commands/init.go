package commands

import (
	"context"
	"fmt"

	"github.com/changsun20/ferre/internal/core"
	"github.com/urfave/cli/v3"
)

var InitCmd = &cli.Command{
	Name:  "init",
	Usage: "Initialize Ferre on your system",
	Action: func(ctx context.Context, c *cli.Command) error {
		if err := core.Bootstrap(); err != nil {
			return fmt.Errorf("Error: %v\n", err)
		}
		return nil
	},
}
