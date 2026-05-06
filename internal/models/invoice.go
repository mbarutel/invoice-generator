package models

import "time"

// Invoice represents the complete invoice structure
type Invoice struct {
	InvoiceNumber string       `json:"invoice_number,omitempty"`
	Date          string       `json:"date"`
	Currency      string       `json:"currency"`
	Company       Company      `json:"company"`
	Client        Client       `json:"client"`
	LineItems     []LineItem   `json:"line_items"`
	TaxRate       float64      `json:"tax_rate"`
	TaxName       string       `json:"tax_name"`
	TaxInclusive  bool         `json:"tax_inclusive"`
	PaymentTerms  PaymentTerms `json:"payment_terms"`
	Notes         string       `json:"notes,omitempty"`
}

// Company represents the seller/service provider details
type Company struct {
	Name     string `json:"name"`
	LogoPath string `json:"logo_path,omitempty"`
	Address  string `json:"address"`
	City     string `json:"city,omitempty"`
	State    string `json:"state,omitempty"`
	ZipCode  string `json:"zip_code,omitempty"`
	Country  string `json:"country,omitempty"`
	Email    string `json:"email"`
	Phone    string `json:"phone,omitempty"`
	Website  string `json:"website,omitempty"`
	ABN      string `json:"abn,omitempty"`
}

// Client represents the customer/buyer details
type Client struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	City    string `json:"city,omitempty"`
	State   string `json:"state,omitempty"`
	ZipCode string `json:"zip_code,omitempty"`
	Country string `json:"country,omitempty"`
	Email   string `json:"email,omitempty"`
	Phone   string `json:"phone,omitempty"`
	ABN     string `json:"abn,omitempty"`
}

// LineItem represents a single item/service on the invoice
type LineItem struct {
	Description     string  `json:"description"`
	Quantity        float64 `json:"quantity"`
	UnitPrice       float64 `json:"unit_price"`
	DiscountPercent float64 `json:"discount_percent,omitempty"`
}

// PaymentTerms contains payment-related information
type PaymentTerms struct {
	DueDays        int    `json:"due_days"`
	PaymentMethods string `json:"payment_methods,omitempty"`
	BankName       string `json:"bank_name,omitempty"`
	AccountNumber  string `json:"account_number,omitempty"`
	BsbNumber      string `json:"bank_state_branch,omitempty"`
	RoutingNumber  string `json:"routing_number,omitempty"`
	IBAN           string `json:"iban,omitempty"`
	SWIFT          string `json:"swift,omitempty"`
}

// Calculations represents calculated values for the invoice
type Calculations struct {
	Subtotal      float64
	TotalDiscount float64
	TaxableAmount float64
	TaxAmount     float64
	Total         float64
	DueDate       time.Time
}
