package commands

import (
	"github.com/changsun20/ferre/internal/core"
	"github.com/urfave/cli/v3"
)

var InstallCmd = &cli.Command{
	Name:      "install",
	Usage:     "Install a package",
	ArgsUsage: "[package]",
	Action:    core.InstallAction,
}
