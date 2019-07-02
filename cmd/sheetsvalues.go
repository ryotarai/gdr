package cmd

import (
	"github.com/spf13/cobra"
)

func buildSheetsValues() *cobra.Command {
	cmd := &cobra.Command{
		Use: "values",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(buildSheetGet())
	cmd.AddCommand(buildSheetUpdate())
	return cmd
}
