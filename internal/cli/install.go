package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install [package]",
	Short: "Install a package",
	Long:  `Install a package from the available buckets.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		packageName := args[0]
		fmt.Printf("Simulating installation of package: %s\n", packageName)
		fmt.Printf("✓ Would download and install %s\n", packageName)
		fmt.Printf("✓ Package %s installed successfully!\n", packageName)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
