package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall [package]",
	Short: "Uninstall a package",
	Long:  `Remove a previously installed package.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		packageName := args[0]
		fmt.Printf("Simulating uninstallation of package: %s\n", packageName)
		fmt.Printf("✓ Would remove package: %s\n", packageName)
		fmt.Printf("✓ Package %s uninstalled successfully!\n", packageName)
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
