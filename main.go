package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/unidoc/unipdf/v3/model"
)

func showHelp() {
	fmt.Println("PDF Metadata Extractor")
	fmt.Println("======================")
	fmt.Println("A tool to extract and display metadata from PDF files.")
	fmt.Println()
	fmt.Printf("Usage: %s [--help] <pdf-file-path>\n", os.Args[0])
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  --help    Show this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Printf("  %s document.pdf\n", os.Args[0])
	fmt.Printf("  %s \"C:\\Documents\\report.pdf\"\n", os.Args[0])
	fmt.Printf("  %s /home/user/files/presentation.pdf\n", os.Args[0])
	fmt.Println()
	os.Exit(0)
}

func main() {

	// Check for --help flag
	for _, arg := range os.Args {
		if arg == "--help" || arg == "-h" {
			showHelp()
		}
	}

	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <pdf-file-path>", os.Args[0])
		os.Exit(1)
	}

	pdfPath := os.Args[1]

	file, err := os.Open(pdfPath)

	// Check for file extension
	fileExt := strings.ToLower(filepath.Ext(pdfPath))

	if fileExt != ".pdf" {
		log.Fatalf("Error: The file %s is not a PDF file. Right now only pdf files are supported", pdfPath)
		os.Exit(1)
	}

	if err != nil {
		log.Fatalf("Error while opening the pdf file: %s", err)
	}
	defer file.Close()

	// Create a new PDF reader instance
	pdfReader, err := model.NewPdfReader(file)

	if err != nil {
		log.Fatalf("Error while creating PDF reader: %s", err)
	}

	// Extract and display the metadata
	pdfMetadata, err := pdfReader.GetPdfInfo()

	if err != nil {
		log.Fatalf("Error while extracting metadata: %s", err)
	}

	log.Println("PDF Metadata:")
	if pdfMetadata != nil {
		log.Printf("Title: %s\n", pdfMetadata.Title)
		log.Printf("Author: %s\n", pdfMetadata.Author)
		log.Printf("Subject: %s\n", pdfMetadata.Subject)
		log.Printf("Creator: %s\n", pdfMetadata.Creator)
	} else {
		log.Println("No metadata found.")
	}

}
