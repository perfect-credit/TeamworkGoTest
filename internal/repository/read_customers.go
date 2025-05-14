package repository

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Total Number of Columns
const COLUMN_NUMBER = 5

// ReadCustomers reads customer data from a CSV file and counts email domains.
func ReadCustomers(filename string) (map[string]int, error) {
	var rowNumber int
	domainCounts := make(map[string]int)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	invalidFile := flag.String("invalid", "./data/invalid.csv", "Invalid CSV file")
	flag.Parse()

	invalidFileHandler, err := os.Create(*invalidFile)
	if err != nil {
		return nil, fmt.Errorf("error creating invalid file: %w", err)
	}
	defer invalidFileHandler.Close()

	// Write header to invalid file
	_, _ = fmt.Fprintf(invalidFileHandler, "row,first_name,last_name,email,gender,ip_address\n")

	for {
		rowNumber++
		record, err := reader.Read()

		// Handle reading errors
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			logInvalidRow(invalidFileHandler, rowNumber, record)
			continue
		}

		// Skip header row
		if rowNumber == 1 {
			continue
		}

		// Validate entry
		if err := ValidateEntry(record); err != nil {
			logInvalidRow(invalidFileHandler, rowNumber, record)
			continue
		}

		email := record[2]
		domain := ExtractDomain(email)
		if domain != "" {
			domainCounts[domain]++
		}
	}
	return domainCounts, nil
}

// ExtractDomain extracts the domain from an email address.
func ExtractDomain(email string) string {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if regexp.MustCompile(emailRegex).MatchString(email) {
		atIndex := strings.Index(email, "@")
		return email[atIndex+1:]
	}
	return ""
}
