package repository

import (
	"testing"
		"fmt"
)

func TestValidateEntry(t *testing.T) {
	testCases := []struct {
		name     string
		record   []string
		expected error
	}{
		{
			name:     "valid entry",
			record:   []string{"John", "Doe", "john.doe@example.com", "Male", "192.168.1.100"},
			expected: nil,
		},
		{
			name:     "invalid number of columns",
			record:   []string{"John", "Doe", "john.doe@example.com", "Male"},
			expected: fmt.Errorf("invalid number of columns: got 4, want 5"),
		},
		{
			name:     "empty first name",
			record:   []string{"", "Doe", "john.doe@example.com", "Male", "192.168.1.100"},
			expected: fmt.Errorf("column 1: first name cannot be empty"),
		},
		{
			name:     "empty last name",
			record:   []string{"John", "", "john.doe@example.com", "Male", "192.168.1.100"},
			expected: fmt.Errorf("column 2: last name cannot be empty"),
		},
		{
			name:     "invalid email",
			record:   []string{"John", "Doe", "john.doe", "Male", "192.168.1.100"},
			expected: fmt.Errorf("column 3: invalid email format: john.doe"),
		},
		{
			name:     "invalid gender",
			record:   []string{"John", "Doe", "john.doe@example.com", "Other", "192.168.1.100"},
			expected: fmt.Errorf("column 4: invalid gender: Other, must be 'Male' or 'Female'"),
		},
		{
			name:     "invalid IP address",
			record:   []string{"John", "Doe", "john.doe@example.com", "Male", "192.168.1.300"},
			expected: fmt.Errorf("column 5: invalid IP octet value: 300"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateEntry(tc.record)
			if (err == nil && tc.expected != nil) || (err != nil && tc.expected == nil) || (err != nil && tc.expected != nil && err.Error() != tc.expected.Error()) {
				t.Errorf("ValidateEntry(%v) = %v, want %v", tc.record, err, tc.expected)
			}
		})
	}
}