package models

import "encoding/xml"

type Invoice struct {
	XMLName       xml.Name `xml:"Invoice"`
	Text          string   `xml:",chardata"`
	Xmlns         string   `xml:"xmlns,attr"`
	Cac           string   `xml:"cac,attr"`
	Cbc           string   `xml:"cbc,attr"`
	Ccts          string   `xml:"ccts,attr"`
	Ds            string   `xml:"ds,attr"`
	Ext           string   `xml:"ext,attr"`
	Qdt           string   `xml:"qdt,attr"`
	Sac           string   `xml:"sac,attr"`
	Udt           string   `xml:"udt,attr"`
	Xsi           string   `xml:"xsi,attr"`
	UBLExtensions struct {
		Text         string `xml:",chardata"`
		UBLExtension struct {
			Text             string `xml:",chardata"`
			ExtensionContent struct {
				Text      string `xml:",chardata"`
				Signature struct {
					Text       string `xml:",chardata"`
					ID         string `xml:"Id,attr"`
					Xmlns      string `xml:"xmlns,attr"`
					SignedInfo struct {
						Text                   string `xml:",chardata"`
						CanonicalizationMethod struct {
							Text      string `xml:",chardata"`
							Algorithm string `xml:"Algorithm,attr"`
						} `xml:"CanonicalizationMethod"`
						SignatureMethod struct {
							Text      string `xml:",chardata"`
							Algorithm string `xml:"Algorithm,attr"`
						} `xml:"SignatureMethod"`
						Reference struct {
							Text       string `xml:",chardata"`
							URI        string `xml:"URI,attr"`
							Transforms struct {
								Text      string `xml:",chardata"`
								Transform struct {
									Text      string `xml:",chardata"`
									Algorithm string `xml:"Algorithm,attr"`
								} `xml:"Transform"`
							} `xml:"Transforms"`
							DigestMethod struct {
								Text      string `xml:",chardata"`
								Algorithm string `xml:"Algorithm,attr"`
							} `xml:"DigestMethod"`
							DigestValue string `xml:"DigestValue"`
						} `xml:"Reference"`
					} `xml:"SignedInfo"`
					SignatureValue string `xml:"SignatureValue"`
					KeyInfo        struct {
						Text     string `xml:",chardata"`
						X509Data struct {
							Text            string `xml:",chardata"`
							X509SubjectName string `xml:"X509SubjectName"`
							X509Certificate string `xml:"X509Certificate"`
						} `xml:"X509Data"`
					} `xml:"KeyInfo"`
				} `xml:"Signature"`
			} `xml:"ExtensionContent"`
		} `xml:"UBLExtension"`
	} `xml:"UBLExtensions"`
	UBLVersionID    string `xml:"UBLVersionID"`
	CustomizationID string `xml:"CustomizationID"`
	ID              string `xml:"ID"`
	IssueDate       string `xml:"IssueDate"`
	IssueTime       string `xml:"IssueTime"`
	DueDate         string `xml:"DueDate"`
	InvoiceTypeCode struct {
		Text           string `xml:",chardata"`
		ListName       string `xml:"listName,attr"`
		ListID         string `xml:"listID,attr"`
		Name           string `xml:"name,attr"`
		ListSchemeURI  string `xml:"listSchemeURI,attr"`
		ListAgencyName string `xml:"listAgencyName,attr"`
		ListURI        string `xml:"listURI,attr"`
	} `xml:"InvoiceTypeCode"`
	Note []struct {
		Text             string `xml:",chardata"`
		LanguageLocaleID string `xml:"languageLocaleID,attr"`
	} `xml:"Note"`
	DocumentCurrencyCode struct {
		Text           string `xml:",chardata"`
		ListID         string `xml:"listID,attr"`
		ListName       string `xml:"listName,attr"`
		ListAgencyName string `xml:"listAgencyName,attr"`
	} `xml:"DocumentCurrencyCode"`
	LineCountNumeric          string `xml:"LineCountNumeric"`
	DespatchDocumentReference struct {
		Text             string `xml:",chardata"`
		ID               string `xml:"ID"`
		DocumentTypeCode struct {
			Text           string `xml:",chardata"`
			ListAgencyName string `xml:"listAgencyName,attr"`
			ListName       string `xml:"listName,attr"`
			ListURI        string `xml:"listURI,attr"`
		} `xml:"DocumentTypeCode"`
	} `xml:"DespatchDocumentReference"`
	Signature struct {
		Text           string `xml:",chardata"`
		ID             string `xml:"ID"`
		SignatoryParty struct {
			Text                string `xml:",chardata"`
			PartyIdentification struct {
				Text string `xml:",chardata"`
				ID   string `xml:"ID"`
			} `xml:"PartyIdentification"`
			PartyName struct {
				Text string `xml:",chardata"`
				Name string `xml:"Name"`
			} `xml:"PartyName"`
		} `xml:"SignatoryParty"`
		DigitalSignatureAttachment struct {
			Text              string `xml:",chardata"`
			ExternalReference struct {
				Text string `xml:",chardata"`
				URI  string `xml:"URI"`
			} `xml:"ExternalReference"`
		} `xml:"DigitalSignatureAttachment"`
	} `xml:"Signature"`
	AccountingSupplierParty struct {
		Text  string `xml:",chardata"`
		Party struct {
			Text                string `xml:",chardata"`
			PartyIdentification struct {
				Text string `xml:",chardata"`
				ID   struct {
					Text             string `xml:",chardata"`
					SchemeID         string `xml:"schemeID,attr"`
					SchemeName       string `xml:"schemeName,attr"`
					SchemeAgencyName string `xml:"schemeAgencyName,attr"`
					SchemeURI        string `xml:"schemeURI,attr"`
				} `xml:"ID"`
			} `xml:"PartyIdentification"`
			PartyName struct {
				Text string `xml:",chardata"`
				Name string `xml:"Name"`
			} `xml:"PartyName"`
			PartyLegalEntity struct {
				Text                string `xml:",chardata"`
				RegistrationName    string `xml:"RegistrationName"`
				RegistrationAddress struct {
					Text            string `xml:",chardata"`
					AddressTypeCode struct {
						Text           string `xml:",chardata"`
						ListAgencyName string `xml:"listAgencyName,attr"`
						ListName       string `xml:"listName,attr"`
					} `xml:"AddressTypeCode"`
					CityName         string `xml:"CityName"`
					CountrySubentity string `xml:"CountrySubentity"`
					District         string `xml:"District"`
					AddressLine      struct {
						Text string `xml:",chardata"`
						Line string `xml:"Line"`
					} `xml:"AddressLine"`
					Country struct {
						Text               string `xml:",chardata"`
						IdentificationCode struct {
							Text           string `xml:",chardata"`
							ListID         string `xml:"listID,attr"`
							ListAgencyName string `xml:"listAgencyName,attr"`
							ListName       string `xml:"listName,attr"`
						} `xml:"IdentificationCode"`
					} `xml:"Country"`
				} `xml:"RegistrationAddress"`
			} `xml:"PartyLegalEntity"`
		} `xml:"Party"`
	} `xml:"AccountingSupplierParty"`
	AccountingCustomerParty struct {
		Text  string `xml:",chardata"`
		Party struct {
			Text                string `xml:",chardata"`
			PartyIdentification struct {
				Text string `xml:",chardata"`
				ID   struct {
					Text             string `xml:",chardata"`
					SchemeID         string `xml:"schemeID,attr"`
					SchemeName       string `xml:"schemeName,attr"`
					SchemeAgencyName string `xml:"schemeAgencyName,attr"`
					SchemeURI        string `xml:"schemeURI,attr"`
				} `xml:"ID"`
			} `xml:"PartyIdentification"`
			PartyLegalEntity struct {
				Text                string `xml:",chardata"`
				RegistrationName    string `xml:"RegistrationName"`
				RegistrationAddress struct {
					Text             string `xml:",chardata"`
					CountrySubentity string `xml:"CountrySubentity"`
					District         string `xml:"District"`
					AddressLine      struct {
						Text string `xml:",chardata"`
						Line string `xml:"Line"`
					} `xml:"AddressLine"`
				} `xml:"RegistrationAddress"`
			} `xml:"PartyLegalEntity"`
		} `xml:"Party"`
	} `xml:"AccountingCustomerParty"`
	PaymentTerms struct {
		Text           string `xml:",chardata"`
		ID             string `xml:"ID"`
		PaymentMeansID string `xml:"PaymentMeansID"`
	} `xml:"PaymentTerms"`
	TaxTotal struct {
		Text      string `xml:",chardata"`
		TaxAmount struct {
			Text       string `xml:",chardata"`
			CurrencyID string `xml:"currencyID,attr"`
		} `xml:"TaxAmount"`
		TaxSubtotal struct {
			Text          string `xml:",chardata"`
			TaxableAmount struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"TaxableAmount"`
			TaxAmount struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"TaxAmount"`
			TaxCategory struct {
				Text      string `xml:",chardata"`
				TaxScheme struct {
					Text string `xml:",chardata"`
					ID   struct {
						Text             string `xml:",chardata"`
						SchemeName       string `xml:"schemeName,attr"`
						SchemeAgencyName string `xml:"schemeAgencyName,attr"`
						SchemeURI        string `xml:"schemeURI,attr"`
					} `xml:"ID"`
					Name        string `xml:"Name"`
					TaxTypeCode string `xml:"TaxTypeCode"`
				} `xml:"TaxScheme"`
			} `xml:"TaxCategory"`
		} `xml:"TaxSubtotal"`
	} `xml:"TaxTotal"`
	LegalMonetaryTotal struct {
		Text                string `xml:",chardata"`
		LineExtensionAmount struct {
			Text       string `xml:",chardata"`
			CurrencyID string `xml:"currencyID,attr"`
		} `xml:"LineExtensionAmount"`
		TaxInclusiveAmount struct {
			Text       string `xml:",chardata"`
			CurrencyID string `xml:"currencyID,attr"`
		} `xml:"TaxInclusiveAmount"`
		PayableAmount struct {
			Text       string `xml:",chardata"`
			CurrencyID string `xml:"currencyID,attr"`
		} `xml:"PayableAmount"`
	} `xml:"LegalMonetaryTotal"`
	InvoiceLine []struct {
		Text             string `xml:",chardata"`
		ID               string `xml:"ID"`
		InvoicedQuantity struct {
			Text                   string `xml:",chardata"`
			UnitCode               string `xml:"unitCode,attr"`
			UnitCodeListID         string `xml:"unitCodeListID,attr"`
			UnitCodeListAgencyName string `xml:"unitCodeListAgencyName,attr"`
		} `xml:"InvoicedQuantity"`
		LineExtensionAmount struct {
			Text       string `xml:",chardata"`
			CurrencyID string `xml:"currencyID,attr"`
		} `xml:"LineExtensionAmount"`
		PricingReference struct {
			Text                      string `xml:",chardata"`
			AlternativeConditionPrice struct {
				Text        string `xml:",chardata"`
				PriceAmount struct {
					Text       string `xml:",chardata"`
					CurrencyID string `xml:"currencyID,attr"`
				} `xml:"PriceAmount"`
				PriceTypeCode struct {
					Text           string `xml:",chardata"`
					ListAgencyName string `xml:"listAgencyName,attr"`
					ListName       string `xml:"listName,attr"`
					ListURI        string `xml:"listURI,attr"`
				} `xml:"PriceTypeCode"`
			} `xml:"AlternativeConditionPrice"`
		} `xml:"PricingReference"`
		AllowanceCharge struct {
			Text                      string `xml:",chardata"`
			ChargeIndicator           string `xml:"ChargeIndicator"`
			AllowanceChargeReasonCode struct {
				Text           string `xml:",chardata"`
				ListAgencyName string `xml:"listAgencyName,attr"`
				ListName       string `xml:"listName,attr"`
				ListURI        string `xml:"listURI,attr"`
			} `xml:"AllowanceChargeReasonCode"`
			MultiplierFactorNumeric string `xml:"MultiplierFactorNumeric"`
			Amount                  struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"Amount"`
			BaseAmount struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"BaseAmount"`
		} `xml:"AllowanceCharge"`
		TaxTotal struct {
			Text      string `xml:",chardata"`
			TaxAmount struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"TaxAmount"`
			TaxSubtotal struct {
				Text          string `xml:",chardata"`
				TaxableAmount struct {
					Text       string `xml:",chardata"`
					CurrencyID string `xml:"currencyID,attr"`
				} `xml:"TaxableAmount"`
				TaxAmount struct {
					Text       string `xml:",chardata"`
					CurrencyID string `xml:"currencyID,attr"`
				} `xml:"TaxAmount"`
				TaxCategory struct {
					Text                   string `xml:",chardata"`
					Percent                string `xml:"Percent"`
					TaxExemptionReasonCode struct {
						Text           string `xml:",chardata"`
						ListAgencyName string `xml:"listAgencyName,attr"`
						ListName       string `xml:"listName,attr"`
						ListURI        string `xml:"listURI,attr"`
					} `xml:"TaxExemptionReasonCode"`
					TaxScheme struct {
						Text string `xml:",chardata"`
						ID   struct {
							Text             string `xml:",chardata"`
							SchemeName       string `xml:"schemeName,attr"`
							SchemeAgencyName string `xml:"schemeAgencyName,attr"`
							SchemeURI        string `xml:"schemeURI,attr"`
						} `xml:"ID"`
						Name        string `xml:"Name"`
						TaxTypeCode string `xml:"TaxTypeCode"`
					} `xml:"TaxScheme"`
				} `xml:"TaxCategory"`
			} `xml:"TaxSubtotal"`
		} `xml:"TaxTotal"`
		Item struct {
			Text                      string `xml:",chardata"`
			Description               string `xml:"Description"`
			SellersItemIdentification struct {
				Text string `xml:",chardata"`
				ID   string `xml:"ID"`
			} `xml:"SellersItemIdentification"`
			CommodityClassification struct {
				Text                   string `xml:",chardata"`
				ItemClassificationCode struct {
					Text           string `xml:",chardata"`
					ListID         string `xml:"listID,attr"`
					ListAgencyName string `xml:"listAgencyName,attr"`
					ListName       string `xml:"listName,attr"`
				} `xml:"ItemClassificationCode"`
			} `xml:"CommodityClassification"`
		} `xml:"Item"`
		Price struct {
			Text        string `xml:",chardata"`
			PriceAmount struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"PriceAmount"`
		} `xml:"Price"`
	} `xml:"InvoiceLine"`
}
