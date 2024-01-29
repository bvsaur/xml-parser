package core

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/bvsaur/xml-parser/models"
)

func ParseXML(folder string) ([]models.Result, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	files, err := os.ReadDir(path.Join(wd, folder))
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, fmt.Errorf("no invoices on the selected folder")
	}

	var fileNames []string

	for _, v := range files {
		fileNames = append(fileNames, v.Name())
	}

	var invoices []models.Result

	for _, file := range fileNames {
		if !validateXML(file) {
			log.Printf("File %s is not an XML, skipping.", file)
			continue
		}

		xmlBuffer, err := os.ReadFile(path.Join(wd, folder, file))
		if err != nil {
			log.Fatal("cannot read xml file:", err)
		}

		var invoice models.Invoice

		if err = xml.Unmarshal(xmlBuffer, &invoice); err != nil {
			log.Fatal("cannot parse xml file:", err)
		}

		invoices = append(invoices, models.Result{
			ID:           invoice.ID,
			IssueDate:    invoice.IssueDate,
			RUC:          invoice.AccountingCustomerParty.Party.PartyIdentification.ID.Text,
			BusinessName: invoice.AccountingCustomerParty.Party.PartyLegalEntity.RegistrationName,
			SaleValue:    invoice.LegalMonetaryTotal.LineExtensionAmount.Text,
			IGV:          invoice.TaxTotal.TaxAmount.Text,
			Total:        invoice.LegalMonetaryTotal.PayableAmount.Text,
			Currency:     invoice.DocumentCurrencyCode.Text,
		})
	}

	return invoices, nil
}

func validateXML(fileName string) bool {
	fileFragments := strings.Split(fileName, ".")
	fileExtension := fileFragments[len(fileFragments)-1]

	return fileExtension == "xml"
}
