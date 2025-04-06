package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rogue",
	Short: "A recon and privilege escalation automation tool for Linux environments",
	Long: `rogue is a CLI tool that automates parts of the Linux privilege escalation methodology.
It runs various enumeration checks and outputs findings in a structured way.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use `rogue help` to view available commands.")
	},
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
