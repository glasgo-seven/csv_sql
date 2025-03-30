package csv_sql

import (
	"bufio"
	"os"
	"strings"
	"slices"
)


type Table struct {
	header []string
	data []map[string]string
}
func (t Table) Print() {
	if _DEBUG {
		println("[_DEBUG] Print\n")
	}
	
	print("_id\t")
	for field := range t.header {
		print(t.header[field], "\t")
	}
	print("\n")

	for i, object := range t.data {
		print(i+1, "\t")
		// for key, value := range object {
		// 	println("\t", key, "\t", value)
		// }
		for field := range t.header {
			print(object[t.header[field]], "\t")
		}
		print("\n")
	}
}


func ParseCSV(fileName string, delimiter string, hasHeader bool) Table {
	if _DEBUG {
		println("[_DEBUG] ParseCSV")
	}

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close() 

	var content []map[string]string

	scanner := bufio.NewScanner(file)
	var header []string
	if scanner.Scan() {
		line := strings.Split(scanner.Text(), delimiter)
		if hasHeader {
			header = line
		} else {
			row := map[string]string {}
			line := strings.Split(scanner.Text(), delimiter)
			for i := 0; i < len(header); i++ {
				header = append(header, "field")
				row[header[i]] = line[i]
			}
			content = append(content, row)
		}
	}
	

	for scanner.Scan() {
		row := map[string]string {}
		line := strings.Split(scanner.Text(), delimiter)
		for i := 0; i < len(header); i++ {
			row[header[i]] = line[i]
		}
		content = append(content, row)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return Table {
		header: header,
		data: content,
	}
}


func formResult(data []map[string]string, _select []string, _as []string, _limit int) []map[string]string {
	if _DEBUG {
		println("[_DEBUG] formResult")
	}

	result := []map[string]string {}

	for i, row := range data {
		if i == _limit {
			break
		}

		select_fields := map[string]string {}

		for key, value := range row {
			if slices.Contains(_select, key) {
				as_index := slices.Index(_select, key)
				if len(_as) == 0 || as_index >= len(_as) {
					select_fields[key] = value
				} else {
					select_fields[_as[as_index]] = value
				}
			}
		}

		if len(select_fields) != 0 {
			result = append(result, select_fields)
		}
	}

	return result
}
