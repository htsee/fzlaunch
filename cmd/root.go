package cmd

import (
	"fmt"
	"os/exec"
	"slices"
	"strings"

	"github.com/htsee/fzlaunch/internal"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "fzlaunch",
	Short: "CLI app launcher for fuzzy finders",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmd.Help(); err != nil {
			return err
		}
		return nil
	},
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List applications",
	RunE: func(cmd *cobra.Command, args []string) error {
		entries, err := internal.DesktopEntries()
		if err != nil {
			return err
		}
		var keys []string
		for key := range entries {
			keys = append(keys, key)
		}
		slices.Sort(keys)
		for _, name := range keys {
			fmt.Println(name)
		}
		return nil
	},
}

var RunCmd = &cobra.Command{
	Use:   "run [app]",
	Short: "Run application",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		appName := args[0]
		if appName == "" {
			return nil
		}
		entries, err := internal.DesktopEntries()
		if err != nil {
			return err
		}
		entry, exist := entries[appName]
		if !exist {
			fmt.Printf("cannot find application %q", appName)
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
	Use:   "preview [app]",
	Short: "Show information about application",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		appName := args[0]
		if appName == "" {
			return nil
		}
		entries, err := internal.DesktopEntries()
		if err != nil {
			return err
		}
		entry, exist := entries[appName]
		if !exist {
			fmt.Printf("cannot find application %q", appName)
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

var Version = "0.1.0"

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("fzlaunch %s\n", Version)
	},
}

func init() {
	RootCmd.AddCommand(ListCmd)
	RootCmd.AddCommand(RunCmd)
	RootCmd.AddCommand(PreviewCmd)
	RootCmd.AddCommand(VersionCmd)
}
