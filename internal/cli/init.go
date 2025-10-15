package cli

import (
	"fmt"
	"os"

	"github.com/changsun20/ferre/internal/core"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Ferre on your system",
	Long: `Initialize Ferre by creating the directory structure, 
copying the executable, and adding it to your PATH.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := core.Bootstrap(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
