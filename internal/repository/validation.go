package repository

import (
	"fmt"
	"regexp"
	"strings"
)

// ValidateEntry checks if the entry is valid according to the CSV structure.
func ValidateEntry(record []string) error {
	if len(record) != COLUMN_NUMBER {
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
