package service

import (
	"sort"
)

// DomainCount holds the domain and its count.
type DomainCount struct {
	Domain string
	Count  int
}

// SortByDomain returns a sorted slice of DomainCount.
func SortByDomain(domainCounts map[string]int) []DomainCount {
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

// SortByCount returns a slice of DomainCount sorted by count in descending order.
func SortByCount(domainCounts map[string]int) []DomainCount {
	// Create a slice of DomainCount to hold all entries
	sortedDomainCounts := make([]DomainCount, 0, len(domainCounts))
	
	// Populate the slice with all domain counts
	for domain, count := range domainCounts {
		sortedDomainCounts = append(sortedDomainCounts, DomainCount{Domain: domain, Count: count})
	}

	// Sort by count in descending order
	sort.Slice(sortedDomainCounts, func(i, j int) bool {
		if sortedDomainCounts[i].Count == sortedDomainCounts[j].Count {
			// If counts are equal, sort by domain name
			return sortedDomainCounts[i].Domain < sortedDomainCounts[j].Domain
		}
		return sortedDomainCounts[i].Count > sortedDomainCounts[j].Count
	})

	return sortedDomainCounts
}