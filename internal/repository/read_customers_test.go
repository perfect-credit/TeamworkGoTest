package repository

import (
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
Bonnie,Ortiz,bortiz1@cyberchimps.com,Female,197.54.209.129
Invalid,User,invalid-email,Female,256.256.256.256
Empty,User,,Female,192.168.1.1
John,Doe,john@example.com,Male,10.0.0.1
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
		"example.com":     1,
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
	}
}

func TestValidateEntry(t *testing.T) {
	tests := []struct {
		name    string
		record  []string
		wantErr bool
	}{
		{
			name:    "Valid record",
			record:  []string{"Mildred", "Hernandez", "mhernandez0@github.io", "Female", "38.194.51.128"},
			wantErr: false,
		},
		{
			name:    "Invalid email",
			record:  []string{"Invalid", "User", "invalid-email", "Female", "38.194.51.128"},
			wantErr: true,
		},
		{
			name:    "Invalid gender",
			record:  []string{"John", "Doe", "john@example.com", "Other", "38.194.51.128"},
			wantErr: true,
		},
		{
			name:    "Invalid IP format",
			record:  []string{"John", "Doe", "john@example.com", "Male", "256.256.256.256"},
			wantErr: true,
		},
		{
			name:    "Empty first name",
			record:  []string{"", "Doe", "john@example.com", "Male", "38.194.51.128"},
			wantErr: true,
		},
		{
			name:    "Empty last name",
			record:  []string{"John", "", "john@example.com", "Male", "38.194.51.128"},
			wantErr: true,
		},
		{
			name:    "Too few columns",
			record:  []string{"John", "Doe", "john@example.com"},
			wantErr: true,
		},
		{
			name:    "Empty email",
			record:  []string{"John", "Doe", "", "Male", "38.194.51.128"},
			wantErr: true,
		},
		{
			name:    "Invalid IP octet",
			record:  []string{"John", "Doe", "john@example.com", "Male", "256.1.2.3"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateEntry(tt.record)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateEntry() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExtractDomain(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected string
	}{
		{
			name:     "Valid email",
			email:    "user@example.com",
			expected: "example.com",
		},
		{
			name:     "Invalid email",
			email:    "invalid-email",
			expected: "",
		},
		{
			name:     "Empty email",
			email:    "",
			expected: "",
		},
		{
			name:     "Complex domain",
			email:    "user@sub.example.co.uk",
			expected: "sub.example.co.uk",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractDomain(tt.email)
			if got != tt.expected {
				t.Errorf("ExtractDomain() = %v, want %v", got, tt.expected)
			}
		})
	}
} 