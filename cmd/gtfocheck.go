package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/meethereum/rogue/internal"
)

type BinaryStatus struct {
	Name       string
	Exists     bool
	IsSUID     bool
	CanSudo    bool
	GTFOBinned bool
	Path       string
}

func CheckBinaryStatus(bin string) BinaryStatus {
	status := BinaryStatus{Name: bin, GTFOBinned: true}

	fullPath, err := exec.LookPath(bin)
	if err == nil {
		status.Exists = true
		status.Path = fullPath

	
		info, err := os.Stat(fullPath)
		if err == nil {
			mode := info.Mode()
			if mode&os.ModeSetuid != 0 || mode.Perm()&04000 != 0 {
				status.IsSUID = true
			}
		}


		out, err := exec.Command("sudo", "-n", "-l").CombinedOutput()
		if err == nil && strings.Contains(string(out), bin) {
			status.CanSudo = true
		}
	}
	return status
}

var gtfocheckCmd = &cobra.Command{
	Use:   "gtfocheck",
	Short: "Check GTFOBins binaries for SUID or sudo misconfigurations",
	Run: func(cmd *cobra.Command, args []string) {
		color.Magenta("Scraping GTFOBins...")
		gtfoBins, err := internal.GetAllGTFOBins()
		if err != nil {
			fmt.Println("Failed to fetch GTFOBins:", err)
			return
		}
		color.Magenta("[*] %d binaries found on GTFOBins\n", len(gtfoBins))

		for _, bin := range gtfoBins {
			status := CheckBinaryStatus(bin)

			if status.Exists && (status.IsSUID || status.CanSudo) {
				color.Red("\n Critically vulnerable binary: %s\n", status.Name)
				color.Red("    → Path: %s\n", status.Path)
				if status.IsSUID {
					color.Red("SUID bit is set")
				}
				if status.CanSudo {
					color.Red("Can be run via sudo")
				}
				color.Cyan("    Listed on GTFOBins: https://gtfobins.github.io/gtfobins/%s/\n", status.Name)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(gtfocheckCmd)
}
