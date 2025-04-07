package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

var pathcheckCmd = &cobra.Command{
	Use:   "pathcheck",
	Short: "Check for writable directories in $PATH",
	Long: `This command checks the system's $PATH environment variable for any directories 
that are writable by the current user. Writable directories in $PATH can allow privilege escalation 
via malicious binary injection.`,
	Run: func(cmd *cobra.Command, args []string) {
		checkWritablePathDirs()
	},
}

func init() {
	rootCmd.AddCommand(pathcheckCmd)
}

func checkWritablePathDirs() {
	path := os.Getenv("PATH")
	dirs := strings.Split(path, ":")

	currentUID := os.Geteuid()

	fmt.Println("Scanning $PATH for writable directories...")
	safe := true
	for _, dir := range dirs {
		absDir, err := filepath.Abs(dir)
		if err != nil {
			continue
		}

		info, err := os.Stat(absDir)
		if err != nil || !info.IsDir() {
			continue
		}

		mode := info.Mode()
		stat := info.Sys().(*syscall.Stat_t)

		if currentUID == int(stat.Uid) && mode&0200 != 0 {
			color.Yellow("[MODERATE] Writable directory  (owner): %s\n ", absDir)
			safe = false
			continue
		}

		if currentUID != int(stat.Uid) && mode&0020 != 0 {
			color.Red("[CRITICAL] Writable directory (group): %s\n", absDir)
			safe = false
			continue
		}

		if mode&0002 != 0 {
			color.Red("[CRITICAL] Writable directory (others): %s\n", absDir)
			safe = false
		}
	}
	if safe {
		color.Green("Completed scanning $PATH for writable directories. None found.")
	}

}
