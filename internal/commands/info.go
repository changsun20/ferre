package commands

import (
	"github.com/changsun20/ferre/internal/core"
	"github.com/urfave/cli/v3"
)

var InfoCmd = &cli.Command{
	Name:   "info",
	Usage:  "Show information about applications in the buckets",
	Action: core.InfoAction,
}
