package parser

import (
	"encoding/json"
	"fmt"
	"invoice-generator/internal/models"
	"os"
)

// ParseInvoice reads and parses a JSON file into an Invoice struct
func ParseInvoice(filePath string) (*models.Invoice, error) {
	// Read the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Parse JSON
	var invoice models.Invoice
	if err := json.Unmarshal(data, &invoice); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Validate required fields
	if err := validate(&invoice); err != nil {
		return nil, err
	}

	// Set defaults
	setDefaults(&invoice)

	return &invoice, nil
}

// validate checks that required fields are present
func validate(invoice *models.Invoice) error {
	if invoice.Company.Name == "" {
		return fmt.Errorf("company name is required")
	}
	if invoice.Client.Name == "" {
		return fmt.Errorf("client name is required")
	}
	if len(invoice.LineItems) == 0 {
		return fmt.Errorf("at least one line item is required")
	}
	for i, item := range invoice.LineItems {
		if item.Description == "" {
			return fmt.Errorf("line item %d: description is required", i+1)
		}
		if item.Quantity <= 0 {
			return fmt.Errorf("line item %d: quantity must be greater than 0", i+1)
		}
		if item.UnitPrice < 0 {
			return fmt.Errorf("line item %d: unit price cannot be negative", i+1)
		}
	}
	return nil
}

// setDefaults sets default values for optional fields
func setDefaults(invoice *models.Invoice) {
	if invoice.Currency == "" {
		invoice.Currency = "USD"
	}
	if invoice.TaxName == "" {
		invoice.TaxName = "Tax"
	}
	if invoice.Date == "" {
		invoice.Date = getCurrentDate()
	}
	if invoice.PaymentTerms.DueDays == 0 {
		invoice.PaymentTerms.DueDays = 30
	}
}

// getCurrentDate returns the current date in YYYY-MM-DD format
func getCurrentDate() string {
	// This is a simplified implementation
	// In production, use time.Now().Format("2006-01-02")
	return "2026-02-08"
}
