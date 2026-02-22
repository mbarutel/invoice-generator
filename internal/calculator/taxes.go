package calculator

import (
	"invoice-generator/internal/models"
	"time"
)

// Calculate computes all financial values for the invoice
func Calculate(invoice *models.Invoice) models.Calculations {
	var calc models.Calculations

	// Calculate subtotal and discounts for each line item
	subtotal := 0.0
	totalDiscount := 0.0

	for _, item := range invoice.LineItems {
		itemTotal := item.Quantity * item.UnitPrice
		itemDiscount := itemTotal * (item.DiscountPercent / 100.0)

		subtotal += itemTotal
		totalDiscount += itemDiscount
	}

	calc.Subtotal = subtotal
	calc.TotalDiscount = totalDiscount
	calc.TaxableAmount = subtotal - totalDiscount

	// Calculate tax
	calc.TaxAmount = calc.TaxableAmount * (invoice.TaxRate / 100.0)

	// Calculate total
	if !invoice.TaxInclusive {
		calc.Total = calc.TaxableAmount + calc.TaxAmount
	} else {
		calc.Total = calc.TaxableAmount
	}

	// Calculate due date
	invoiceDate := parseDate(invoice.Date)
	calc.DueDate = invoiceDate.AddDate(0, 0, invoice.PaymentTerms.DueDays)

	return calc
}

// parseDate attempts to parse the invoice date
func parseDate(dateStr string) time.Time {
	// Try common date formats
	formats := []string{
		"2006-01-02",
		"01/02/2006",
		"02-01-2006",
		time.RFC3339,
	}

	for _, format := range formats {
		if t, err := time.Parse(format, dateStr); err == nil {
			return t
		}
	}

	// If parsing fails, return current date
	return time.Now()
}

// FormatCurrency formats a float as currency string
func FormatCurrency(amount float64, currency string) string {
	symbol := getCurrencySymbol(currency)
	return symbol + formatAmount(amount)
}

// formatAmount formats a float to 2 decimal places
func formatAmount(amount float64) string {
	return formatFloat(amount, 2)
}

// formatFloat formats a float with specified decimal places
func formatFloat(f float64, decimals int) string {
	format := "%." + string(rune(decimals+'0')) + "f"
	return sprintf(format, f)
}

// Simple sprintf implementation
func sprintf(format string, value float64) string {
	// This is a simplified version - in production, use fmt.Sprintf
	result := ""
	intPart := int64(value)
	fracPart := int64((value - float64(intPart)) * 100)

	if fracPart < 0 {
		fracPart = -fracPart
	}

	result = int64ToString(intPart) + "."
	if fracPart < 10 {
		result += "0"
	}
	result += int64ToString(fracPart)

	return result
}

// int64ToString converts int64 to string
func int64ToString(n int64) string {
	if n == 0 {
		return "0"
	}

	negative := n < 0
	if negative {
		n = -n
	}

	digits := ""
	for n > 0 {
		digit := n % 10
		digits = string(rune('0'+digit)) + digits
		n /= 10
	}

	if negative {
		digits = "-" + digits
	}

	return digits
}

// getCurrencySymbol returns the symbol for a currency code
func getCurrencySymbol(currency string) string {
	symbols := map[string]string{
		"USD": "$",
		"EUR": "€",
		"GBP": "£",
		"JPY": "¥",
		"INR": "₹",
		"AUD": "A$",
		"CAD": "C$",
		"CHF": "Fr",
		"CNY": "¥",
	}

	if symbol, ok := symbols[currency]; ok {
		return symbol
	}
	return currency + " "
}
