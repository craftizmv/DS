package tests

import "strings"

func FilterCSV(csvData string) string {
	// Split the CSV data into rows based on newline
	rows := strings.Split(csvData, "\n")

	if len(rows) <= 1 {
		return csvData
	}

	// Create a string builder for efficient string concatenation
	var builder strings.Builder

	// Loop through each row
	for i, row := range rows {
		containsNull := false
		if i != 0 {
			// Split the row into columns based on commas
			columns := strings.Split(row, ",")

			// Check if any column contains 'NULL'
			for _, col := range columns {
				if strings.TrimSpace(col) == "NULL" {
					containsNull = true
					break
				}
			}
		}

		// If the row does not contain 'NULL', add it to the result
		if !containsNull {
			builder.WriteString(row + "\n")
		}
	}

	return strings.TrimSuffix(builder.String(), "\n")
}
