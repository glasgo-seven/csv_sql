package main

import (
	"csv_sql/cmd"
)


func main() {
	table := csv_sql.ParseCSV("./rsc/happiness.csv", ",", true)
	result := csv_sql.Select("Country or region", "Score").As("Country", "AAA").From(table).Limit(5).Result

	result.Print()

	// csv_viewer.MainGUI()
}
