package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"TeamworkGoTest/internal/repository"
	"TeamworkGoTest/internal/service"
)

func main() {
		inputFile := flag.String("input", "", "Input CSV file")
		outputFile := flag.String("output", "", "Output CSV file")
		sortBy := flag.String("sort", "domain", "sort by 'domain' or 'count'")
		flag.Parse()

		domainCounts, err := repository.ReadCustomers(*inputFile)
		
		if err != nil {
				log.Fatalf("Error reading customers: %v", err)
		}

		var sortedDomains []service.DomainCount
		switch *sortBy {
		case "domain":
				sortedDomains = service.SortByDomain(domainCounts)
				fmt.Println("Sorted by domain:")
		case "count":
				sortedDomains = service.SortByCount(domainCounts)
				fmt.Println("Sorted by Count:")
		default:
				fmt.Println("Invalid sort option. Use 'domain' or 'count'.")
				return
		}
		
		if *outputFile != "" {			
				file, err := os.Create(*outputFile)
				if err != nil {
						log.Fatalf("Error creating output file: %v", err)
				}
				defer file.Close()

				_, _ = fmt.Fprintf(file, "domain, count\n")
				for _, domainCount := range sortedDomains {
						_, err := fmt.Fprintf(file, "%s, %d\n", domainCount.Domain, domainCount.Count)
						if err != nil {
								log.Fatalf("Error writing to output file: %v", err)
						}
				}
		} else {
				_, _ = fmt.Printf("domain: count\n")
				for _, domainCount := range sortedDomains {
						fmt.Printf("%s: %d\n", domainCount.Domain, domainCount.Count)
				}
		}
}
