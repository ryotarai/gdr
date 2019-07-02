package cmd

import "github.com/spf13/cobra"

func Execute() error {
	return buildRoot().Execute()
}

func buildRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gdr",
		Short: "gdr is command line tool for Google Drive",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(buildSheet())
	cmd.AddCommand(versionCmd)
	return cmd
}
