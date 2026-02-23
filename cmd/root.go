package cmd

import (
	"fmt"

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

func init() {
	RootCmd.AddCommand(ListCmd)
}
