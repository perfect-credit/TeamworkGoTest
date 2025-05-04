package service

import (
	"sort"
)

// DomainCount holds the domain and its count.
type DomainCount struct {
	Domain string
	Count  int
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



 