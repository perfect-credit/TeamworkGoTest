package customerimporter

import (
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

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
			return nil, err
	}

	domainCounts := make(map[string]int)

	for _, record := range records[1:] { // Skip header
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
	var domains []DomainCount
	for domain, count := range domainCounts {
			domains = append(domains, DomainCount{Domain: domain, Count: count})
	}

	sort.Slice(domains, func(i, j int) bool {
			return domains[i].Domain < domains[j].Domain
	})

	return domains
}



 