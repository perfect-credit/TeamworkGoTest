package repository

import (
	// "encoding/csv"
	"os"
	"path/filepath"
	"testing"
)

func TestReadCustomers(t *testing.T) {
	// Create data directory if it doesn't exist
	dataDir := "./data"
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		t.Fatalf("Failed to create data directory: %v", err)
	}
	defer os.RemoveAll(dataDir) // Clean up data directory after test

	// Create a temporary CSV file for testing
	testCSV := filepath.Join(dataDir, "test_customers.csv")
	file, err := os.Create(testCSV)
	if err != nil {
		t.Fatalf("Failed to create test CSV file: %v", err)
	}
	defer os.Remove(testCSV) // Clean up test file

	// Write test data to the CSV file
	data := `first_name,last_name,email,gender,ip_address
Mildred,Hernandez,mhernandez0@github.io,Female,38.194.51.128
Bad,Example1,invalid-email.com,Female,8.8.8.8
Good,Example2,example@example.com,Male,192.168.1.1
Bad,Example3,@invalid-email2.com,Female,1.1.1.1
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
		"github.io": 1,
		"example.com": 1,
	}

	for domain, count := range expectedCounts {
		if domainCounts[domain] != count {
			t.Errorf("Expected %d for domain %s, got %d", count, domain, domainCounts[domain])
		}
	}

	// Check for invalid entries in invalid.csv
	invalidFile := filepath.Join(dataDir, "invalid.csv")
	if _, err := os.Stat(invalidFile); os.IsNotExist(err) {
		t.Errorf("Invalid file not created")
	} else {
		// Check the content of invalid.csv
		invalidRecords, err := os.ReadFile(invalidFile)
		if err != nil {
			t.Fatalf("Failed to read invalid CSV file: %v", err)
		}

		expectedInvalidData := `row,first_name,last_name,email,gender,ip_address
3,Bad,Example1,invalid-email.com,Female,8.8.8.8
5,Bad,Example3,@invalid-email2.com,Female,1.1.1.1
`
		if string(invalidRecords) != expectedInvalidData {
			t.Errorf("Invalid CSV content does not match expected content.\nExpected:\n%s\nGot:\n%s", expectedInvalidData, string(invalidRecords))
		}
	}
}
