package customerimporter

import (
	"testing"
)

// TestExtractDomain tests the extractDomain function.
func TestExtractDomain(t *testing.T) {
    tests := []struct {
        email    string
        expected string
    }{
        {"user@example.com", "example.com"},
        {"admin@domain.org", "domain.org"},
        {"invalid-email", ""},
        {"user@sub.domain.com", "sub.domain.com"},
        {"@missingusername.com", "missingusername.com"},
    }

    for _, test := range tests {
        got := extractDomain(test.email)
        if got != test.expected {
            t.Errorf("extractDomain(%q) = %q; want %q", test.email, got, test.expected)
        }
    }
}

// TestSortDomains tests the SortDomains function.
func TestSortDomains(t *testing.T) {
    domainCounts := map[string]int{
        "example.com": 4,
        "domain.com":  2,
        "test.com":    3,
    }

    expectedSorted := []DomainCount{
        {"domain.com", 2},
        {"example.com", 4},
        {"test.com", 3},
    }

    sortedDomains := SortDomains(domainCounts)

    for i, domainCount := range sortedDomains {
        if domainCount != expectedSorted[i] {
            t.Errorf("SortDomains() = %v; want %v", sortedDomains, expectedSorted)
            break
        }
    }
}