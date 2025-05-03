package customerimporter

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

//Total Number of Column
const	columnNum = 5

//Number of row
var rowNo int = 0

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
		
		domainCounts := make(map[string]int)
		
		//Name invalid file as invalid.csv
		invalidFile := flag.String("invalid", "invalid.csv", "Invalid CSV file")

		flag.Parse()
		
		//Create invalid.csv
		if *invalidFile != "" {
				invalid_file, _ := os.Create(*invalidFile)
				_, _ = fmt.Fprintf(invalid_file, "row, first_name,last_name,email,gender,ip_address\n")

				for {
						rowNo++
						record, err := reader.Read()

						//Handle reading errors
						if err != nil {
								if err.Error() == "EOF" {
										break
								}

								fmt.Fprintf(invalid_file, "%d, ",	rowNo,)

								for i := 0; i < columnNum - 1; i++ {
										fmt.Fprintf(invalid_file, "%s,",	GetValue(record, i))								
								}
								fmt.Fprintf(invalid_file, "%s\n",	GetValue(record, columnNum - 1))

								continue
						}

						//Header pass
						if rowNo == 1 {
								continue
						}

						typeCompare := ValidateEntry(record)
						if typeCompare != nil {
							fmt.Fprintf(invalid_file, "%d, ",	rowNo,)

							for i:=0; i < columnNum - 1; i++ {
									fmt.Fprintf(invalid_file, "%s,",	record[i])								
							}
							fmt.Fprintf(invalid_file, "%s\n",	record[columnNum - 1])				
								
							continue
						}	

						email := record[2]
						domain := ExtractDomain(email)

						if domain != "" {
								domainCounts[domain]++
						}
				}
		}

		return domainCounts, nil
}

//Handle null value in row
func GetValue(record []string, index int) string {
		if index < len(record) && record[index] != "" {
				return record[index]
		}

		return "null!"
}

// ValidateEntry checks if the entry is valid according to the CSV structure.
func ValidateEntry(record []string) error {	
		if err := ValidateFirstName(record[0]); err != nil {
				return err
		}

		if err := ValidateLastName(record[1]); err != nil {
				return err
		}

		if err := ValidateEmail(record[2]); err != nil {
				return err
		}

		if err := ValidateGender(record[3]); err != nil {
				return err
		}

		if err := ValidateIPAddress(record[4]); err != nil {
				return err
		}

		return nil
}

// ValidateFirstName checks if the first name is valid (non-empty).
func ValidateFirstName(firstName string) error {
		if firstName == "" {
				return fmt.Errorf("first name cannot be empty")
		}

		return nil
}

// ValidateLastName checks if the last name is valid (non-empty).
func ValidateLastName(lastName string) error {
		if lastName == "" {
				return fmt.Errorf("last name cannot be empty")
		}

		return nil
}

// ValidateEmail checks if the email is in a valid format.
func ValidateEmail(email string) error {
		const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		re := regexp.MustCompile(emailRegex)
		if !re.MatchString(email) {
				return fmt.Errorf("invalid email format: %s", email)
		}

		return nil
}

// ValidateGender checks if the gender is either "Male" or "Female".
func ValidateGender(gender string) error {
		if gender != "Male" && gender != "Female" {
				return fmt.Errorf("invalid gender: %s, must be 'Male' or 'Female'", gender)
		}

		return nil
}

// ValidateIPAddress checks if the IP address is in a valid format.
func ValidateIPAddress(ip string) error {
		const ipRegex = `^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`
		re := regexp.MustCompile(ipRegex)

		if !re.MatchString(ip) {
				return fmt.Errorf("invalid IP address format: %s", ip)
		}

		return nil
}

//ExtractDomain extracts the domain from an email address.
func ExtractDomain(email string) (string) {
		// Regular expression for validating an email
		const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

		// Compile the regex
		re := regexp.MustCompile(emailRegex)
		atIndex := strings.Index(email, "@")

		// Check if the email matches the regex
		if re.MatchString(email) {
				return email[atIndex+1:] // Return the email if valid
		}	
		
		return ""
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



 