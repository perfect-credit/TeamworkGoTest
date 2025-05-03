package customerimporter

import (
	"os"
	"testing"
)

func TestReadCustomers(t *testing.T) {
	// Create a temporary CSV file for testing
	testCSV := "test_customers.csv"
	file, err := os.Create(testCSV)
	if err != nil {
		t.Fatalf("Failed to create test CSV file: %v", err)
	}
	defer os.Remove(testCSV) // Clean up after test

	// Write test data to the CSV file
	data := `first_name,last_name,email,gender,ip_address
Mildred,Hernandez,mhernandez0@github.io,Female,38.194.51.128
Bonnie,Ortiz,bortiz1@cyberchimps.com,Female,197.54.209.129
Invalid,User,invalid-email,Female,256.256.256.256
`
	_, err = file.WriteString(data)
	if err != nil {
		t.Fatalf("Failed to write to test CSV file: %v", err)
	}
	file.Close()

	// Call ReadCustomers
	domainCounts, err := ReadCustomers(testCSV)
	if err != nil {
		t.Fatalf("Error reading customers: %v", err)
	}

	// Check expected domain counts
	expectedCounts := map[string]int{
		"github.io":      1,
		"cyberchimps.com": 1,
	}

	for domain, count := range expectedCounts {
		if domainCounts[domain] != count {
			t.Errorf("Expected %d for domain %s, got %d", count, domain, domainCounts[domain])
		}
	}

	// Check for invalid entries in invalid.csv
	invalidFile := "invalid.csv"
	if _, err := os.Stat(invalidFile); os.IsNotExist(err) {
		t.Errorf("Invalid file not created")
	}
}

func TestValidateEntry(t *testing.T) {
	validRecord := []string{"Mildred", "Hernandez", "mhernandez0@github.io", "Female", "38.194.51.128"}
	invalidRecord := []string{"Invalid", "User", "invalid-email", "Unknown", "256.256.256.256"}

	if err := ValidateEntry(validRecord); err != nil {
		t.Errorf("Valid record failed validation: %v", err)
	}

	if err := ValidateEntry(invalidRecord); err == nil {
		t.Error("Invalid record passed validation, but it should have failed")
	}
}

func TestSortDomains(t *testing.T) {
	domainCounts := map[string]int{
		"github.io":      3,
		"cyberchimps.com": 2,
		"example.com":    1,
	}

	sortedDomains := SortDomains(domainCounts)

	if len(sortedDomains) != len(domainCounts) {
		t.Errorf("Expected %d domains, got %d", len(domainCounts), len(sortedDomains))
	}

	if sortedDomains[0].Domain != "cyberchimps.com" {
		t.Errorf("Expected first domain to be 'cyberchimps.com', got '%s'", sortedDomains[0].Domain)
	}
}
