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
		output, err := exec.Command("groups").Output()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		if strings.Contains(string(output), "docker") {
			color.Green("[+] Potential vulnerability: User is in the docker group.")
			fmt.Println("Try mounting / as a container volume to access sensitive files like /etc/shadow.")
		} else {
			color.Red("[-] User is NOT in the docker group.")
		}
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
}
