package main

import (
	"csv_sql/cmd"
)


func main() {
	// Set debug mode
	csv_sql.DEBUG = true

	// Read CSV File
	table := csv_sql.ParseCSV("./rsc/file.csv", ",", true)

	// Use SQL syntax formed as a chain of methods
	result := csv_sql.Select("Score").As("Value").
				From(table).
				Limit(5).
				Result
	// .Where("Field > 10 AND Field < 20").OrderBy("Field DESC", "Field_A ASC").Result

	// Print result of the query
	result.Print()
}
