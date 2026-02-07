package generator

import (
	"fmt"
	"invoice-generator/internal/calculator"
	"invoice-generator/internal/models"
	"os"

	"github.com/google/uuid"
	"github.com/jung-kurt/gofpdf"
)

const (
	// Page dimensions and margins
	pageWidth   = 210.0 // A4 width in mm
	marginLeft  = 20.0
	marginTop   = 20.0
	marginRight = 20.0
)

// GeneratePDF creates a PDF invoice from the given invoice data
func GeneratePDF(invoice *models.Invoice, outputPath string) error {
	// Generate invoice number if not provided
	if invoice.InvoiceNumber == "" {
		invoice.InvoiceNumber = generateInvoiceNumber()
	}

	// Calculate invoice totals
	calc := calculator.Calculate(invoice)

	// Create PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Set font
	pdf.SetFont("Arial", "", 10)

	// Draw invoice header
	drawHeader(pdf, invoice)

	// Draw company and client details
	yPos := drawCompanyAndClient(pdf, invoice)

	// Draw invoice details (date, invoice number, due date)
	yPos = drawInvoiceDetails(pdf, invoice, calc, yPos)

	// Draw line items table
	yPos = drawLineItemsTable(pdf, invoice, calc, yPos)

	// Draw totals section
	yPos = drawTotals(pdf, invoice, calc, yPos)

	// Draw payment terms
	drawPaymentTerms(pdf, invoice, calc, yPos)

	// Save PDF
	return pdf.OutputFileAndClose(outputPath)
}

// generateInvoiceNumber creates a UUID-based invoice number
func generateInvoiceNumber() string {
	id := uuid.New()
	return id.String()
}

// drawHeader draws the invoice title and logo
func drawHeader(pdf *gofpdf.Fpdf, invoice *models.Invoice) {
	// Draw logo if provided
	if invoice.Company.LogoPath != "" {
		if _, err := os.Stat(invoice.Company.LogoPath); err == nil {
			// Determine image type
			imgType := getImageType(invoice.Company.LogoPath)
			if imgType != "" {
				pdf.ImageOptions(invoice.Company.LogoPath, marginLeft, marginTop, 30, 0, false,
					gofpdf.ImageOptions{ImageType: imgType, ReadDpi: true}, 0, "")
			}
		}
	}

	// Draw "INVOICE" title on the right
	pdf.SetFont("Arial", "B", 24)
	pdf.SetXY(pageWidth-marginRight-60, marginTop)
	pdf.CellFormat(60, 10, "INVOICE", "", 0, "R", false, 0, "")
}

// drawCompanyAndClient draws company and client information
func drawCompanyAndClient(pdf *gofpdf.Fpdf, invoice *models.Invoice) float64 {
	yPos := marginTop + 35.0

	// Company details (left side)
	pdf.SetFont("Arial", "B", 10)
	pdf.SetXY(marginLeft, yPos)
	pdf.Cell(40, 5, "From:")

	pdf.SetFont("Arial", "", 10)
	yPos += 6
	pdf.SetXY(marginLeft, yPos)
	pdf.Cell(80, 5, invoice.Company.Name)

	yPos += 5
	pdf.SetXY(marginLeft, yPos)
	pdf.Cell(80, 5, invoice.Company.Address)

	if invoice.Company.City != "" {
		yPos += 5
		cityLine := invoice.Company.City
		if invoice.Company.State != "" {
			cityLine += ", " + invoice.Company.State
		}
		if invoice.Company.ZipCode != "" {
			cityLine += " " + invoice.Company.ZipCode
		}
		pdf.SetXY(marginLeft, yPos)
		pdf.Cell(80, 5, cityLine)
	}

	if invoice.Company.Country != "" {
		yPos += 5
		pdf.SetXY(marginLeft, yPos)
		pdf.Cell(80, 5, invoice.Company.Country)
	}

	if invoice.Company.Email != "" {
		yPos += 5
		pdf.SetXY(marginLeft, yPos)
		pdf.Cell(80, 5, invoice.Company.Email)
	}

	if invoice.Company.Phone != "" {
		yPos += 5
		pdf.SetXY(marginLeft, yPos)
		pdf.Cell(80, 5, invoice.Company.Phone)
	}

	// Client details (right side)
	yPosClient := marginTop + 35.0
	pdf.SetFont("Arial", "B", 10)
	pdf.SetXY(pageWidth-marginRight-80, yPosClient)
	pdf.Cell(40, 5, "Bill To:")

	pdf.SetFont("Arial", "", 10)
	yPosClient += 6
	pdf.SetXY(pageWidth-marginRight-80, yPosClient)
	pdf.Cell(80, 5, invoice.Client.Name)

	yPosClient += 5
	pdf.SetXY(pageWidth-marginRight-80, yPosClient)
	pdf.Cell(80, 5, invoice.Client.Address)

	if invoice.Client.City != "" {
		yPosClient += 5
		cityLine := invoice.Client.City
		if invoice.Client.State != "" {
			cityLine += ", " + invoice.Client.State
		}
		if invoice.Client.ZipCode != "" {
			cityLine += " " + invoice.Client.ZipCode
		}
		pdf.SetXY(pageWidth-marginRight-80, yPosClient)
		pdf.Cell(80, 5, cityLine)
	}

	if invoice.Client.Country != "" {
		yPosClient += 5
		pdf.SetXY(pageWidth-marginRight-80, yPosClient)
		pdf.Cell(80, 5, invoice.Client.Country)
	}

	if invoice.Client.Email != "" {
		yPosClient += 5
		pdf.SetXY(pageWidth-marginRight-80, yPosClient)
		pdf.Cell(80, 5, invoice.Client.Email)
	}

	// Return the max Y position
	if yPos > yPosClient {
		return yPos + 10
	}
	return yPosClient + 10
}

// drawInvoiceDetails draws invoice number, date, and due date
func drawInvoiceDetails(
	pdf *gofpdf.Fpdf,
	invoice *models.Invoice,
	calc models.Calculations,
	yPos float64,
) float64 {
	pdf.SetFont("Arial", "B", 10)

	// Invoice Number
	pdf.SetXY(marginLeft, yPos)
	pdf.Cell(40, 6, "Invoice Number:")
	pdf.SetFont("Arial", "", 10)
	pdf.SetXY(marginLeft+45, yPos)
	pdf.Cell(60, 6, invoice.InvoiceNumber)

	// Invoice Date
	yPos += 6
	pdf.SetFont("Arial", "B", 10)
	pdf.SetXY(marginLeft, yPos)
	pdf.Cell(40, 6, "Invoice Date:")
	pdf.SetFont("Arial", "", 10)
	pdf.SetXY(marginLeft+45, yPos)
	pdf.Cell(60, 6, invoice.Date)

	// Due Date
	yPos += 6
	pdf.SetFont("Arial", "B", 10)
	pdf.SetXY(marginLeft, yPos)
	pdf.Cell(40, 6, "Due Date:")
	pdf.SetFont("Arial", "", 10)
	pdf.SetXY(marginLeft+45, yPos)
	pdf.Cell(60, 6, calc.DueDate.Format("2006-01-02"))

	return yPos + 15
}

// drawLineItemsTable draws the table of line items
func drawLineItemsTable(
	pdf *gofpdf.Fpdf,
	invoice *models.Invoice,
	calc models.Calculations,
	yPos float64,
) float64 {
	// Table header
	pdf.SetFillColor(230, 230, 230)
	pdf.SetFont("Arial", "B", 10)

	colWidths := []float64{70, 20, 25, 25, 30}
	headers := []string{"Description", "Qty", "Unit Price", "Discount", "Amount"}

	pdf.SetXY(marginLeft, yPos)
	for i, header := range headers {
		pdf.CellFormat(colWidths[i], 8, header, "1", 0, "C", true, 0, "")
	}

	yPos += 8

	// Table rows
	pdf.SetFont("Arial", "", 9)
	for _, item := range invoice.LineItems {
		pdf.SetXY(marginLeft, yPos)

		// Description
		pdf.CellFormat(colWidths[0], 7, item.Description, "1", 0, "L", false, 0, "")

		// Quantity
		qtyStr := fmt.Sprintf("%.2f", item.Quantity)
		pdf.CellFormat(colWidths[1], 7, qtyStr, "1", 0, "C", false, 0, "")

		// Unit Price
		priceStr := fmt.Sprintf("%.2f", item.UnitPrice)
		pdf.CellFormat(colWidths[2], 7, priceStr, "1", 0, "R", false, 0, "")

		// Discount
		discountStr := fmt.Sprintf("%.1f%%", item.DiscountPercent)
		pdf.CellFormat(colWidths[3], 7, discountStr, "1", 0, "R", false, 0, "")

		// Amount (after discount)
		amount := item.Quantity * item.UnitPrice * (1 - item.DiscountPercent/100.0)
		amountStr := fmt.Sprintf("%.2f", amount)
		pdf.CellFormat(colWidths[4], 7, amountStr, "1", 0, "R", false, 0, "")

		yPos += 7
	}

	return yPos + 5
}

// drawTotals draws the totals section
func drawTotals(
	pdf *gofpdf.Fpdf,
	invoice *models.Invoice,
	calc models.Calculations,
	yPos float64,
) float64 {
	rightColX := pageWidth - marginRight - 50
	leftColX := rightColX - 40

	pdf.SetFont("Arial", "", 10)

	// Subtotal
	pdf.SetXY(leftColX, yPos)
	pdf.Cell(40, 6, "Subtotal:")
	pdf.SetXY(rightColX, yPos)
	pdf.CellFormat(50, 6, fmt.Sprintf("%.2f", calc.Subtotal), "", 0, "R", false, 0, "")

	// Total Discount
	if calc.TotalDiscount > 0 {
		yPos += 6
		pdf.SetXY(leftColX, yPos)
		pdf.Cell(40, 6, "Discount:")
		pdf.SetXY(rightColX, yPos)
		pdf.CellFormat(50, 6, fmt.Sprintf("-%.2f", calc.TotalDiscount), "", 0, "R", false, 0, "")
	}

	// Tax
	yPos += 6
	taxLabel := fmt.Sprintf("%s (%.1f%%):", invoice.TaxName, invoice.TaxRate)
	pdf.SetXY(leftColX, yPos)
	pdf.Cell(40, 6, taxLabel)
	pdf.SetXY(rightColX, yPos)
	pdf.CellFormat(50, 6, fmt.Sprintf("%.2f", calc.TaxAmount), "", 0, "R", false, 0, "")

	// Total
	yPos += 8
	pdf.SetFont("Arial", "B", 12)
	pdf.SetXY(leftColX, yPos)
	pdf.Cell(40, 8, "Total:")
	pdf.SetXY(rightColX, yPos)
	totalStr := fmt.Sprintf("%s %.2f", getCurrencySymbol(invoice.Currency), calc.Total)
	pdf.CellFormat(50, 8, totalStr, "", 0, "R", false, 0, "")

	return yPos + 15
}

// drawPaymentTerms draws payment information
func drawPaymentTerms(
	pdf *gofpdf.Fpdf,
	invoice *models.Invoice,
	calc models.Calculations,
	yPos float64,
) {
	pdf.SetFont("Arial", "B", 10)
	pdf.SetXY(marginLeft, yPos)
	pdf.Cell(40, 6, "Payment Information")

	yPos += 8
	pdf.SetFont("Arial", "", 9)

	if invoice.PaymentTerms.PaymentMethods != "" {
		pdf.SetXY(marginLeft, yPos)
		pdf.Cell(40, 5, "Payment Methods: "+invoice.PaymentTerms.PaymentMethods)
		yPos += 5
	}

	if invoice.PaymentTerms.BankName != "" {
		pdf.SetXY(marginLeft, yPos)
		pdf.Cell(40, 5, "Bank: "+invoice.PaymentTerms.BankName)
		yPos += 5
	}

	if invoice.PaymentTerms.BsbNumber != "" {
		pdf.SetXY(marginLeft, yPos)
		pdf.Cell(40, 5, "BSB: "+invoice.PaymentTerms.BsbNumber)
		yPos += 5
	}

	if invoice.PaymentTerms.AccountNumber != "" {
		pdf.SetXY(marginLeft, yPos)
		pdf.Cell(40, 5, "Account Number: "+invoice.PaymentTerms.AccountNumber)
		yPos += 5
	}

	if invoice.PaymentTerms.IBAN != "" {
		pdf.SetXY(marginLeft, yPos)
		pdf.Cell(40, 5, "IBAN: "+invoice.PaymentTerms.IBAN)
		yPos += 5
	}

	if invoice.PaymentTerms.SWIFT != "" {
		pdf.SetXY(marginLeft, yPos)
		pdf.Cell(40, 5, "SWIFT: "+invoice.PaymentTerms.SWIFT)
		yPos += 5
	}

	if invoice.Notes != "" {
		yPos += 5
		pdf.SetFont("Arial", "B", 10)
		pdf.SetXY(marginLeft, yPos)
		pdf.Cell(40, 6, "Notes")
		yPos += 6
		pdf.SetFont("Arial", "", 9)
		pdf.SetXY(marginLeft, yPos)
		pdf.MultiCell(pageWidth-marginLeft-marginRight, 5, invoice.Notes, "", "L", false)
	}
}

// getImageType determines the image type from file extension
func getImageType(path string) string {
	if len(path) < 4 {
		return ""
	}
	ext := path[len(path)-4:]
	switch ext {
	case ".jpg", ".JPG":
		return "JPG"
	case "jpeg", "JPEG":
		return "JPEG"
	case ".png", ".PNG":
		return "PNG"
	case ".gif", ".GIF":
		return "GIF"
	default:
		return ""
	}
}

// getCurrencySymbol returns the currency symbol
func getCurrencySymbol(currency string) string {
	symbols := map[string]string{
		"USD": "$",
		"EUR": "€",
		"GBP": "£",
		"JPY": "¥",
		"INR": "₹",
		"AUD": "A$",
		"CAD": "C$",
	}

	if symbol, ok := symbols[currency]; ok {
		return symbol
	}
	return currency
}
