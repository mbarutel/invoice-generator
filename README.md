# Invoice Generator

A professional command-line invoice generator that creates PDF invoices from JSON input files. Built with Go and featuring automatic calculations, tax support, itemized discounts, and customizable payment terms.

## Features

- **PDF Generation**: Creates professional, formatted PDF invoices
- **UUID Invoice Numbers**: Automatically generates unique invoice identifiers
- **Tax Calculations**: Automatic calculation of subtotals, taxes, and totals
- **Itemized Discounts**: Support for per-item discount percentages
- **Payment Terms**: Include due dates, payment methods, and bank details
- **Logo Support**: Embed company logos in invoices
- **Multiple Currencies**: Support for USD, EUR, GBP, JPY, INR, and more
- **Clean CLI**: Simple command-line interface

## Installation

### Prerequisites

- Go 1.16 or higher

### Build from Source

```bash
cd invoice-generator
go build -o invoice-generator
```

This will create an executable named `invoice-generator` in the current directory.

## Usage

### Basic Usage

```bash
./invoice-generator -input examples/sample-invoice.json -output output/my-invoice.pdf
```

### Command-Line Options

- `-input <file>` - Path to input JSON file (required)
- `-output <file>` - Path to output PDF file (default: output/invoice.pdf)
- `-help` - Show help message

### Examples

```bash
# Generate invoice with default output location
./invoice-generator -input examples/simple-invoice.json

# Generate invoice with custom output name
./invoice-generator -input examples/sample-invoice.json -output invoices/client-feb-2026.pdf

# Show help
./invoice-generator -help
```

## JSON Input Format

The invoice data is defined in JSON format with the following structure:

### Complete Example

```json
{
  "date": "2026-02-08",
  "currency": "USD",
  "company": {
    "name": "Your Company Name",
    "logo_path": "./logo.png",
    "address": "123 Business Street",
    "city": "San Francisco",
    "state": "CA",
    "zip_code": "94102",
    "country": "United States",
    "email": "billing@yourcompany.com",
    "phone": "+1 (555) 123-4567",
    "website": "www.yourcompany.com"
  },
  "client": {
    "name": "Client Company Name",
    "address": "456 Client Avenue",
    "city": "New York",
    "state": "NY",
    "zip_code": "10001",
    "country": "United States",
    "email": "accounts@client.com",
    "phone": "+1 (555) 987-6543"
  },
  "line_items": [
    {
      "description": "Service or Product Description",
      "quantity": 10,
      "unit_price": 100.00,
      "discount_percent": 5
    }
  ],
  "tax_rate": 8.5,
  "tax_name": "Sales Tax",
  "payment_terms": {
    "due_days": 30,
    "payment_methods": "Bank Transfer, Credit Card, PayPal",
    "bank_name": "First National Bank",
    "account_number": "1234567890",
    "routing_number": "021000021",
    "iban": "US12 1234 5678 9012 3456 78",
    "swift": "FNBAUS44"
  },
  "notes": "Additional notes or payment instructions"
}
```

### Required Fields

- `company.name` - Your company name
- `company.address` - Your company address
- `company.email` - Your company email
- `client.name` - Client/customer name
- `client.address` - Client address
- `line_items` - Array of at least one line item
  - `description` - Item description
  - `quantity` - Quantity (must be > 0)
  - `unit_price` - Price per unit

### Optional Fields

- `invoice_number` - Custom invoice number (auto-generated if omitted)
- `date` - Invoice date (defaults to current date)
- `currency` - Currency code (defaults to "USD")
- `company.logo_path` - Path to company logo image (PNG, JPG, GIF)
- `company.city`, `state`, `zip_code`, `country` - Additional address fields
- `company.phone`, `website` - Contact information
- `client.city`, `state`, `zip_code`, `country`, `email`, `phone` - Client details
- `line_items[].discount_percent` - Discount percentage (0-100)
- `tax_rate` - Tax percentage (defaults to 0)
- `tax_name` - Tax label (defaults to "Tax")
- `payment_terms.due_days` - Days until payment due (defaults to 30)
- `payment_terms.payment_methods` - Accepted payment methods
- `payment_terms.bank_name`, `account_number`, etc. - Bank details
- `notes` - Additional notes or terms

## Supported Currencies

The following currency codes are supported with their symbols:

- **USD** - $ (US Dollar)
- **EUR** - € (Euro)
- **GBP** - £ (British Pound)
- **JPY** - ¥ (Japanese Yen)
- **INR** - ₹ (Indian Rupee)
- **AUD** - A$ (Australian Dollar)
- **CAD** - C$ (Canadian Dollar)
- **CHF** - Fr (Swiss Franc)
- **CNY** - ¥ (Chinese Yuan)

Other currency codes can be used and will be displayed as text.

## How It Works

1. **Parse JSON**: Reads and validates the input JSON file
2. **Generate UUID**: Creates a unique invoice number if not provided
3. **Calculate Totals**: Automatically computes:
   - Line item totals with discounts
   - Subtotal (sum of all line items)
   - Total discount amount
   - Taxable amount (subtotal - discounts)
   - Tax amount (based on tax rate)
   - Grand total
   - Due date (invoice date + due days)
4. **Create PDF**: Generates a professional PDF with:
   - Company logo (if provided)
   - Company and client information
   - Invoice number, date, and due date
   - Itemized table with quantities, prices, and discounts
   - Subtotals, tax, and grand total
   - Payment terms and bank details
   - Additional notes

## Project Structure

```
invoice-generator/
├── main.go                      # CLI entry point
├── go.mod                       # Go module dependencies
├── internal/
│   ├── models/
│   │   └── invoice.go           # Invoice data structures
│   ├── generator/
│   │   └── pdf.go               # PDF generation logic
│   ├── calculator/
│   │   └── taxes.go             # Tax and discount calculations
│   └── parser/
│       └── json.go              # JSON input parsing
├── examples/
│   ├── sample-invoice.json      # Comprehensive example
│   └── simple-invoice.json      # Simple example
├── output/                      # Generated PDFs (created automatically)
└── README.md                    # This file
```

## Development

### Dependencies

- [gofpdf](https://github.com/jung-kurt/gofpdf) - PDF generation library
- [uuid](https://github.com/google/uuid) - UUID generation

### Building

```bash
go build -o invoice-generator
```

### Testing

Try the included examples:

```bash
./invoice-generator -input examples/simple-invoice.json -output test1.pdf
./invoice-generator -input examples/sample-invoice.json -output test2.pdf
```

## Tips

1. **Logo Images**: Supported formats are PNG, JPG, and GIF. For best results, use a PNG with transparent background, approximately 300x100 pixels.

2. **Date Format**: Dates can be in various formats (YYYY-MM-DD, MM/DD/YYYY, etc.). The YYYY-MM-DD format is recommended.

3. **Discount Calculations**: Discounts are applied per line item before calculating subtotals and taxes.

4. **Tax Calculations**: Tax is calculated on the subtotal after discounts have been applied.

5. **Invoice Numbers**: If you don't specify an invoice number, a UUID will be generated automatically. For sequential numbering, specify your own invoice numbers in the JSON.

## License

MIT License - feel free to use this project for personal or commercial purposes.

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests.

## Support

For questions or issues, please open an issue on the project repository.
