package customerimporter

import (
	"bufio"
	"encoding/csv"
	"os"
	"sort"
	"strings"
)

// DomainCount holds the domain and its count.
type DomainCount struct {
	Domain string
	Count  int
}

// ReadCustomers reads customer data from a CSV file and counts email domains.
func ReadCustomers(filename string) (map[string]int, error) {
	file, err := os.Open(filename)
	if err != nil {
			return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	// Skip the header
	if _, err := reader.Read(); err != nil {
			return nil, err
	}

	domainCounts := make(map[string]int)

	for {
			record, err := reader.Read()
			if err != nil {
					break // End of file or an error
			}

			if len(record) < 3 {
					continue // Skip records that don't have enough fields
			}

			email := record[2]
			domain := extractDomain(email)
			if domain != "" {
					domainCounts[domain]++
			}
	}

	return domainCounts, nil
}

// extractDomain extracts the domain from an email address.
func extractDomain(email string) string {
	atIndex := strings.Index(email, "@")
	if atIndex == -1 {
			return ""
	}
	return email[atIndex+1:]
}

// SortDomains returns a sorted slice of DomainCount.
func SortDomains(domainCounts map[string]int) []DomainCount {
	// Create a slice to hold the keys (domains)
	domains := make([]string, 0, len(domainCounts))
	
	// Populate the slice with keys from the map
	for domain := range domainCounts {
			domains = append(domains, domain)
	}

	// Sort the slice of domains
	sort.Strings(domains)

	// Create a slice of DomainCount to hold the results
	sortedDomainCounts := make([]DomainCount, len(domains))
	for i, domain := range domains {
			sortedDomainCounts[i] = DomainCount{Domain: domain, Count: domainCounts[domain]}
	}

	return sortedDomainCounts
}



 