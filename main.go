package main

import (
	"github.com/htsee/fzlaunch/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cobra.CheckErr(cmd.RootCmd.Execute())
}
