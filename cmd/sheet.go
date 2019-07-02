package cmd

import (
	"github.com/spf13/cobra"
)

func buildSheet() *cobra.Command {
	cmd := &cobra.Command{
		Use: "sheets",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(buildSheetsValues())
	return cmd
}
