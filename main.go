package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bvsaur/xml-parser/core"
	"github.com/urfave/cli/v2"
)

func main() {

	var xmlFolder string
	var csvFile string

	app := &cli.App{
		Name:     "XML Parser",
		HelpName: "xml-parser",
		Usage:    "Parse an invoice XML file and export it to an CSV format",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "folder",
				Aliases:     []string{"f"},
				Value:       "files",
				Usage:       "Name of the folder that contains the XML files",
				Destination: &xmlFolder,
			},
			&cli.StringFlag{
				Name:        "file",
				Value:       "invoices",
				Usage:       "Name of the CSV file that will be exported with the invoices result",
				Destination: &csvFile,
			},
		},
		Action: func(cCtx *cli.Context) error {
			results, err := core.ParseXML(xmlFolder)
			if err != nil {
				return err
			}

			err = core.GenerateCSVFile(results, fmt.Sprintf("%s.csv", csvFile))

			if err != nil {
				return err
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
