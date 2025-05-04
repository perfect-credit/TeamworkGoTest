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

// GetValue retrieves a value from the record at the specified index.
func GetValue(record []string, index int) string {
    if index < len(record) && record[index] != "" {
        return record[index]
    }
    return "null!"
}

// ValidateEntry checks if the entry is valid according to the CSV structure.
func ValidateEntry(record []string) error {
    if len(record) < COLUMN_NUMBER {
        return fmt.Errorf("invalid number of columns: got %d, want %d", len(record), COLUMN_NUMBER)
    }

    validators := []func(string) error{
        ValidateFirstName,
        ValidateLastName,
        ValidateEmail,
        ValidateGender,
        ValidateIPAddress,
    }

    for i, validator := range validators {
        if err := validator(record[i]); err != nil {
            return fmt.Errorf("column %d: %w", i+1, err)
        }
    }
    return nil
}

// Validation functions for each field
func ValidateFirstName(firstName string) error {
    if firstName == "" {
        return fmt.Errorf("first name cannot be empty")
    }
    return nil
}

func ValidateLastName(lastName string) error {
    if lastName == "" {
        return fmt.Errorf("last name cannot be empty")
    }
    return nil
}

func ValidateEmail(email string) error {
    const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
    if !regexp.MustCompile(emailRegex).MatchString(email) {
        return fmt.Errorf("invalid email format: %s", email)
    }
    return nil
}

func ValidateGender(gender string) error {
    if gender != "Male" && gender != "Female" {
        return fmt.Errorf("invalid gender: %s, must be 'Male' or 'Female'", gender)
    }
    return nil
}

func ValidateIPAddress(ip string) error {
    const ipRegex = `^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`
    if !regexp.MustCompile(ipRegex).MatchString(ip) {
        return fmt.Errorf("invalid IP address format: %s", ip)
    }

    // Split IP into octets and validate each
    octets := strings.Split(ip, ".")
    for _, octet := range octets {
        num := 0
        fmt.Sscanf(octet, "%d", &num)
        if num < 0 || num > 255 {
            return fmt.Errorf("invalid IP octet value: %s", octet)
        }
    }
    return nil
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

// logInvalidRow logs invalid rows to the invalid CSV file.
func logInvalidRow(file *os.File, rowNumber int, record []string) {
    fmt.Fprintf(file, "%d,", rowNumber)
    for i := 0; i < COLUMN_NUMBER; i++ {
        fmt.Fprintf(file, "%s,", GetValue(record, i))
    }
    fmt.Fprintln(file)
}