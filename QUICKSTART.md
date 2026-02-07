# Quick Start Guide

Get started with Invoice Generator in 3 simple steps!

## Step 1: Build the Tool

```bash
cd invoice-generator
go build -o invoice-generator
```

## Step 2: Create Your Invoice JSON

Create a file called `my-invoice.json`:

```json
{
  "date": "2026-02-08",
  "currency": "USD",
  "company": {
    "name": "Your Company Name",
    "address": "Your Address",
    "email": "your@email.com"
  },
  "client": {
    "name": "Client Name",
    "address": "Client Address"
  },
  "line_items": [
    {
      "description": "Service Description",
      "quantity": 1,
      "unit_price": 1000.00,
      "discount_percent": 0
    }
  ],
  "tax_rate": 10.0,
  "tax_name": "Sales Tax",
  "payment_terms": {
    "due_days": 30
  }
}
```

## Step 3: Generate Your Invoice

```bash
./invoice-generator -input my-invoice.json -output my-invoice.pdf
```

That's it! Your professional PDF invoice is ready.

## Try the Examples

Test with the included examples:

```bash
./invoice-generator -input examples/simple-invoice.json
./invoice-generator -input examples/sample-invoice.json
./invoice-generator -input examples/invoice-with-logo.json
```

## Common Customizations

### Add a Logo

```json
{
  "company": {
    "name": "Your Company",
    "logo_path": "./path/to/logo.png",
    ...
  }
}
```

### Add Discounts

```json
{
  "line_items": [
    {
      "description": "Item",
      "quantity": 10,
      "unit_price": 100.00,
      "discount_percent": 15
    }
  ]
}
```

### Change Currency

```json
{
  "currency": "EUR",
  ...
}
```

Supported: USD, EUR, GBP, JPY, INR, AUD, CAD, CHF, CNY

### Add Payment Details

```json
{
  "payment_terms": {
    "due_days": 30,
    "payment_methods": "Bank Transfer, PayPal",
    "bank_name": "Your Bank",
    "account_number": "123456789",
    "iban": "GB12 1234 5678 9012 3456 78"
  }
}
```

### Add Notes

```json
{
  "notes": "Thank you for your business! Payment due within 30 days."
}
```

## Need Help?

- See the full [README.md](README.md) for complete documentation
- Check the `examples/` folder for more samples
- Run `./invoice-generator -help` for command-line options

## Tips

1. Invoice numbers are auto-generated using UUIDs
2. Tax is calculated on the amount after discounts
3. Dates can be in YYYY-MM-DD format (recommended)
4. Logo images should be PNG, JPG, or GIF format
5. All monetary amounts use 2 decimal places

Happy invoicing!
