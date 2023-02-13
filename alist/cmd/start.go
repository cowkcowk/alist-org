package cmd

import "github.com/spf13/cobra"

var StartCmd = &cobra.Command{
	Use: "start",
	Short: "Silent start alist server with `--force-bin-dir`",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func start() {
	
}
