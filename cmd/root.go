package cmd

import (
	"github.com/spf13/cobra"
)
var(
	rootCmd = &cobra.Command{
		Use:"cobra",
		Short:"IPtracker CLI app",
		Long:`IPtracker CLI interface`,
	}
)

func Execute() error{
	return rootCmd.Execute()
}