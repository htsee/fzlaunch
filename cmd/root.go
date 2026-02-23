package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/htsee/fzlaunch/internal"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "fzlaunch",
	Short: "cli app launcher, can be piped into a fuzzy search menu",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmd.Help(); err != nil {
			return err
		}
		return nil
	},
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "list applications",
	RunE: func(cmd *cobra.Command, args []string) error {
		entries, err := internal.DesktopEntries()
		if err != nil {
			return err
		}
		for name := range entries {
			fmt.Println(name)
		}
		return nil
	},
}

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "run application",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		appName := args[0]
		entries, err := internal.DesktopEntries()
		if err != nil {
			return err
		}
		entry, exist := entries[appName]
		if !exist {
			fmt.Printf("cannot find application %v", appName)
			return nil
		}
		app := exec.Command(entry.Exec, entry.Args...)
		if err := app.Start(); err != nil {
			return err
		}
		return nil
	},
}

var PreviewCmd = &cobra.Command{
	Use:   "preview",
	Short: "show information about application",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		appName := args[0]
		entries, err := internal.DesktopEntries()
		if err != nil {
			return err
		}
		entry, exist := entries[appName]
		if !exist {
			fmt.Printf("cannot find application %v", appName)
			return nil
		}
		preview := appName + "\n\n"
		if entry.GenericName != "" {
			preview += entry.GenericName + "\n"
		}
		if len(entry.Categories) != 0 {
			preview += fmt.Sprintf("Categories: %v", strings.Join(entry.Categories, ", ")) + "\n"
		}
		if len(entry.Keywords) != 0 {
			preview += fmt.Sprintf("Keywords: %v", strings.Join(entry.Keywords, ", ")) + "\n"
		}
		if entry.Comment != "" {
			preview += "\n" + entry.Comment + "\n"
		}
		fmt.Println(preview)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(ListCmd)
	RootCmd.AddCommand(RunCmd)
	RootCmd.AddCommand(PreviewCmd)
}
