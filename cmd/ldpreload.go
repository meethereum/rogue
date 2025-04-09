package cmd

import (
	_"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/fatih/color"

)

var ldpreloadCmd = &cobra.Command{
	Use:   "ldpreload",
	Short: "Check for LD_PRELOAD file vulnerability",
	Run: func(cmd *cobra.Command, args []string) {
		checkLDPreload()
	},
}

func checkLDPreload() {
	color.Magenta("Checking for LD_PRELOAD vulnerability");
	path := "/etc/ld.so.preload"
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		color.Green("/etc/ld.so.preload does not exist. Looks safe.")
		return
	}

	
	color.Red("Potential critical vulneraability: /etc/ld.so.preload exists")

	if info.Mode().Perm()&0o002 != 0 {
		color.Red("Potential critical vulnerability: File is world-writable.")
	} else if unixAccessWrite(path) {
		color.Yellow("Potential medium vulnerability: File is user-writable.")
	} else {
		color.Green("File is not writable. Seems safe.")
	}
}

// unixAccessWrite checks write access based on real user ID
func unixAccessWrite(file string) bool {
	f, err := os.OpenFile(file, os.O_WRONLY, 0600)
	if err != nil {
		return false
	}
	f.Close()
	return true
}

func init() {
	rootCmd.AddCommand(ldpreloadCmd)
}
