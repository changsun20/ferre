package commands

import (
	"github.com/changsun20/ferre/internal/core"
	"github.com/urfave/cli/v3"
)

var BucketCommand = &cli.Command{
	Name:  "bucket",
	Usage: "Manage bucket operations",
	Commands: []*cli.Command{
		BucketAddCmd,
		BucketRemoveCmd,
	},
}

var BucketAddCmd = &cli.Command{
	Name:      "add",
	Usage:     "Add a new bucket",
	ArgsUsage: "<name> [<url>]",
	Action:    core.BucketAddAction,
}

var BucketRemoveCmd = &cli.Command{
	Name:      "rm",
	Usage:     "Remove an existing bucket",
	ArgsUsage: "<name>",
	Action:    core.BucketRemoveAction,
}
