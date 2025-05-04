package repository

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

//Total Number of Column
const	COLUMN_NUMBER = 5

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

// ReadCustomers reads customer data from a CSV file and counts email domains.
func ReadCustomers(filename string) (map[string]int, error) {
		var rowNumber int = 0

		file, err := os.Open(filename)

		if err != nil {
				return nil, err
		}
		defer file.Close()

		reader := csv.NewReader(bufio.NewReader(file))
		
		domainCounts := make(map[string]int)
		
		//Name invalid file as invalid.csv
		invalidFile := flag.String("invalid", "./data/invalid.csv", "Invalid CSV file")

		flag.Parse()
		
		//Create invalid.csv
		if *invalidFile != "" {
				invalid_file, _ := os.Create(*invalidFile)
				_, _ = fmt.Fprintf(invalid_file, "row, first_name,last_name,email,gender,ip_address\n")

				for {
						rowNumber++
						record, err := reader.Read()

						//Handle reading errors
						if err != nil {
								if err.Error() == "EOF" {
										break
								}

								fmt.Fprintf(invalid_file, "%d, ",	rowNumber,)

								for i := 0; i < COLUMN_NUMBER - 1; i++ {
										fmt.Fprintf(invalid_file, "%s,",	GetValue(record, i))								
								}
								fmt.Fprintf(invalid_file, "%s\n",	GetValue(record, COLUMN_NUMBER - 1))

								continue
						}

						//Header pass
						if rowNumber == 1 {
								continue
						}

						typeCompare := ValidateEntry(record)
						if typeCompare != nil {
							fmt.Fprintf(invalid_file, "%d, ",	rowNumber,)

							for i:=0; i < COLUMN_NUMBER - 1; i++ {
									fmt.Fprintf(invalid_file, "%s,",	record[i])								
							}
							fmt.Fprintf(invalid_file, "%s\n",	record[COLUMN_NUMBER - 1])				
								
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
