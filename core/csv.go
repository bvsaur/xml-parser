package core

import (
	"encoding/csv"
	"os"

	"github.com/bvsaur/xml-parser/models"
)

var headers = []string{
	"ID",
	"IssueDate",
	"RUC",
	"BusinessName",
	"SaleValue",
	"IGV",
	"Total",
	"Currency",
}

func GenerateCSVFile(invoiceResults []models.Result, fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	records := parseCSVRecords(invoiceResults)

	w := csv.NewWriter(f)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			return err
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		return err
	}

	return nil
}

func parseCSVRecords(invoiceResults []models.Result) [][]string {
	var records [][]string

	records = append(records, headers)

	for _, result := range invoiceResults {
		record := []string{
			result.ID,
			result.IssueDate,
			result.RUC,
			result.BusinessName,
			result.SaleValue,
			result.IGV,
			result.Total,
			result.Currency,
		}

		records = append(records, record)
	}

	return records
}
