package cmd

import (
	"fmt"
	"os/exec"

	"github.com/htsee/fzlaunch/internal"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "fzlaunch",
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

func init() {
	RootCmd.AddCommand(ListCmd)
	RootCmd.AddCommand(RunCmd)
}
