package repository

import (
	"fmt"
	"os"
	"strings"
)

// logInvalidRow logs invalid rows to the invalid CSV file.
func logInvalidRow(file *os.File, rowNumber int, record []string) {
	fields := make([]string, 0, COLUMN_NUMBER+1)
	fields = append(fields, fmt.Sprintf("%d", rowNumber))
	for i := 0; i < COLUMN_NUMBER; i++ {
		fields = append(fields, GetValue(record, i))
	}
	fmt.Fprintln(file, strings.Join(fields, ","))
}

// GetValue retrieves a value from the record at the specified index.
func GetValue(record []string, index int) string {
	if index < len(record) && record[index] != "" {
		return record[index]
	}
	return "null!"
}
