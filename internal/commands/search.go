package commands

import (
	"github.com/changsun20/ferre/internal/core"
	"github.com/urfave/cli/v3"
)

var SearchCmd = &cli.Command{
	Name:      "search",
	Usage:     "Search for a package",
	ArgsUsage: "[query]",
	Action:    core.SearchAction,
}
