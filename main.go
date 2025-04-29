package main

import (
	"customerimporter/customerimporter"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	inputFile := flag.String("input", "customers.csv", "Input CSV file")
	outputFile := flag.String("output", "", "Output CSV file")
	flag.Parse()

	domainCounts, err := customerimporter.ReadCustomers(*inputFile)
	if err != nil {
			log.Fatalf("Error reading customers: %v", err)
	}

	sortedDomains := customerimporter.SortDomains(domainCounts)

	if *outputFile != "" {
			file, err := os.Create(*outputFile)
			if err != nil {
					log.Fatalf("Error creating output file: %v", err)
			}
			defer file.Close()

			for _, domainCount := range sortedDomains {
					_, err := fmt.Fprintf(file, "%s: %d\n", domainCount.Domain, domainCount.Count)
					if err != nil {
							log.Fatalf("Error writing to output file: %v", err)
					}
			}
	} else {
			for _, domainCount := range sortedDomains {
					fmt.Printf("%s: %d\n", domainCount.Domain, domainCount.Count)
			}
	}
}
