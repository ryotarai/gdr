package cmd

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"

	"github.com/ryotarai/gdr/client"
	"github.com/spf13/cobra"
	"google.golang.org/api/sheets/v4"
)

func buildSheetUpdate() *cobra.Command {
	var serviceAccountFile string
	var sheetID string
	var sheetRange string
	var valueRangeJSON string
	var valueInputOption string

	cmd := &cobra.Command{
		Use: "update",
		Run: func(cmd *cobra.Command, args []string) {
			c, err := client.NewFromServiceAccountFile(serviceAccountFile)
			if err != nil {
				log.Fatalf("Unable to load service account: %v", err)
			}

			srv, err := sheets.New(c)
			if err != nil {
				log.Fatalf("Unable to retrieve Sheets client: %v", err)
			}

			var r io.Reader
			if valueRangeJSON == "-" {
				r = os.Stdin
			} else {
				r = strings.NewReader(valueRangeJSON)
			}
			valueRange := &sheets.ValueRange{}
			err = json.NewDecoder(r).Decode(&valueRange)
			if err != nil {
				log.Fatalf("Unable to decode value range: %v", err)
			}

			resp, err := srv.Spreadsheets.Values.Update(sheetID, sheetRange, valueRange).ValueInputOption(valueInputOption).Do()
			if err != nil {
				log.Fatalf("Unable to update values: %v", err)
			}

			e := json.NewEncoder(os.Stdout)
			err = e.Encode(resp)
			if err != nil {
				log.Fatalf("Unable to generate result: %v", err)
			}
		},
	}

	cmd.Flags().StringVar(&serviceAccountFile, "service-account-file", "", "Path to service account file")
	cmd.MarkFlagRequired("service-account-file")

	cmd.Flags().StringVar(&sheetID, "sheet-id", "", "Spreadsheet ID")
	cmd.MarkFlagRequired("sheet-id")
	cmd.Flags().StringVar(&sheetRange, "range", "", "Range")
	cmd.MarkFlagRequired("range")

	cmd.Flags().StringVar(&valueRangeJSON, "value-range", "-", `Value range in JSON format ("-" is stdin)`)
	cmd.Flags().StringVar(&valueInputOption, "value-input-option", "USER_ENTERED", "https://developers.google.com/sheets/api/reference/rest/v4/ValueInputOption")

	return cmd
}
