package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"github.com/spf13/cobra"
	"github.com/meethereum/rogue/internal" 
	"github.com/fatih/color"
)

var kernelCheckCmd = &cobra.Command{
	Use:   "kernelcheck",
	Short: "Check kernel version and search for possible exploits",
	Run: func(cmd *cobra.Command, args []string) {
		runKernelCheck()
	},
}

func init() {
	rootCmd.AddCommand(kernelCheckCmd)
}

func runKernelCheck() {
	osName, osVersion, kernel, err := getKernelAndOSInfo()
	if err != nil {
		fmt.Println("Failed to get OS info:", err)
		return
	}

	color.Magenta("Detected OS: %s %s\n", osName, osVersion)
	color.Magenta("Kernel Version: %s\n", kernel)

	exploits, err := internal.ScrapeKernelExploits(kernel)
	if err != nil {
		fmt.Println("Failed to search Exploit-DB:", err)
		return
	}

	if len(exploits) == 0 {
		color.Green("No known kernel exploits found.")
	} else {
		color.Red("Potential kernel exploits from Exploit-DB:")
		for _, link := range exploits {
			color.Cyan("   →", link)
		}
	}
}

func getKernelAndOSInfo() (osName, version, kernel string, err error) {
	osData, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return "", "", "", err
	}

	lines := strings.Split(string(osData), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "NAME=") {
			osName = strings.Trim(line[5:], "\"")
		}
		if strings.HasPrefix(line, "VERSION_ID=") {
			version = strings.Trim(line[11:], "\"")
		}
	}

	kernelBytes, err := exec.Command("uname", "-r").Output()
	if err != nil {
		return "", "", "", err
	}
	kernel = strings.TrimSpace(string(kernelBytes))
	return
}
