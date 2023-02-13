package cmd

import (
	"fmt"
	"os"

	"github.com/alist-org/alist/v3/cmd/flags"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "alist",
	Short: "A file list program that supports multiple storage.",
	Long: `A file list program that supports multiple storage,
built with love by Xhofe and friends in Go/Solid.js.
Complete documentation is available at https://alist.nn.ci/`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVar(&flags.DataDir, "data", "data", "config file")
}
