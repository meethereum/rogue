package cmd

import (
	"os/exec"
	"strings"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Check if user is in the docker group",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Checking if user is in the docker group...")
		output, err := exec.Command("groups").Output()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		if strings.Contains(string(output), "docker") {
			color.Red("[CRITICAL] User is in the docker group.")
			fmt.Println("Try mounting / as a container volume to access sensitive files like /etc/shadow to exploit.")
		} else {
			color.Green("User is NOT in the docker group.")
		}
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
}
