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
		for _, entry := range entries {
			fmt.Println(entry.Name)
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
		for _, entry := range entries {
			if appName == entry.Name {
				cmd := exec.Command(entry.Exec)
				if err := cmd.Start(); err != nil {
					return err
				}
				return nil
			}
		}
		fmt.Printf("cannot find application %v", appName)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(ListCmd)
	RootCmd.AddCommand(RunCmd)
}
