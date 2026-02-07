package main

import (
	"flag"
	"fmt"
	"invoice-generator/internal/generator"
	"invoice-generator/internal/parser"
	"os"
	"path/filepath"
)

func main() {
	// Define CLI flags
	inputFile := flag.String("input", "", "Path to input JSON file (required)")
	outputFile := flag.String("output", "", "Path to output PDF file (default: output/invoice.pdf)")
	help := flag.Bool("help", false, "Show help message")

	flag.Parse()

	// Show help if requested or no input provided
	if *help || *inputFile == "" {
		printHelp()
		return
	}

	// Set default output path if not provided
	if *outputFile == "" {
		*outputFile = "output/invoice.pdf"
	}

	// Ensure output directory exists
	outputDir := filepath.Dir(*outputFile)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	// Parse the invoice JSON
	fmt.Printf("Parsing invoice from: %s\n", *inputFile)
	invoice, err := parser.ParseInvoice(*inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing invoice: %v\n", err)
		os.Exit(1)
	}

	// Generate the PDF
	fmt.Printf("Generating PDF invoice...\n")
	if err := generator.GeneratePDF(invoice, *outputFile); err != nil {
		fmt.Fprintf(os.Stderr, "Error generating PDF: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ Invoice generated successfully!\n")
	fmt.Printf("  Invoice Number: %s\n", invoice.InvoiceNumber)
	fmt.Printf("  Output: %s\n", *outputFile)
}

func printHelp() {
	fmt.Println("Invoice Generator - Create professional PDF invoices from JSON files")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  invoice-generator -input <json-file> [-output <pdf-file>]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -input string")
	fmt.Println("        Path to input JSON file (required)")
	fmt.Println("  -output string")
	fmt.Println("        Path to output PDF file (default: output/invoice.pdf)")
	fmt.Println("  -help")
	fmt.Println("        Show this help message")
	fmt.Println()
	fmt.Println("Example:")
	fmt.Println("  invoice-generator -input examples/sample-invoice.json -output my-invoice.pdf")
	fmt.Println()
	fmt.Println("For more information, see the README.md file.")
}
