package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/perfect-credit/teamworkGoTest/internal/repository"
)

func main() {
		inputFile := flag.String("input", "", "Input CSV file")
		outputFile := flag.String("output", "", "Output CSV file")
		flag.Parse()

		domainCounts, err := repository.ReadCustomers(*inputFile)
		if err != nil {
				log.Fatalf("Error reading customers: %v", err)
		}

		sortedDomains := repository.SortDomains(domainCounts)
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
