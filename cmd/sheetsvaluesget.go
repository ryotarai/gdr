package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ryotarai/gdr/client"
	"github.com/spf13/cobra"
	"google.golang.org/api/sheets/v4"
)

func buildSheetGet() *cobra.Command {
	var serviceAccountFile string
	var sheetID string
	var getRange string

	cmd := &cobra.Command{
		Use: "get",
		Run: func(cmd *cobra.Command, args []string) {
			c, err:= client.NewFromServiceAccountFile(serviceAccountFile)
			if err != nil {
				log.Fatalf("Unable to load service account: %v", err)
			}

			srv, err := sheets.New(c)
			if err != nil {
				log.Fatalf("Unable to retrieve Sheets client: %v", err)
			}

			valueRange, err := srv.Spreadsheets.Values.Get(sheetID, getRange).Do()
			if err != nil {
				log.Fatalf("Unable to get values: %v", err)
			}

			b, err := json.Marshal(valueRange)
			if err != nil {
				log.Fatalf("Unable to generate result: %v", err)
			}

			fmt.Printf("%s\n", string(b))
		},
	}

	cmd.Flags().StringVar(&serviceAccountFile, "service-account-file", "", "Path to service account file")
	cmd.MarkFlagRequired("service-account-file")

	cmd.Flags().StringVar(&sheetID, "sheet-id", "", "Spreadsheet ID")
	cmd.MarkFlagRequired("sheet-id")
	cmd.Flags().StringVar(&getRange, "range", "", "Range")
	cmd.MarkFlagRequired("range")

	return cmd
}
