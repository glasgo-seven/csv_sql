package csv_sql

import (
	"bufio"
	"os"
	"strings"
	"slices"
	"fmt"
)

const (
	PRINT_ID_CELL_FORMAT_D string = "%-6d"
	PRINT_ID_CELL_FORMAT_S string = "%-6s"
	PRINT_FIELD_CELL_FORMAT string = "%-16s"
)


type Table struct {
	header []string
	data []map[string]string
}
func (t Table) Print() {
	if DEBUG {
		println("[_DEBUG] Print\n")
	}
	
	fmt.Printf(PRINT_ID_CELL_FORMAT_S, "_id")
	for field := range t.header {
		fmt.Printf(PRINT_FIELD_CELL_FORMAT, t.header[field])
	}
	print("\n")

	for i, object := range t.data {
		fmt.Printf(PRINT_ID_CELL_FORMAT_D, i+1)
		// for key, value := range object {
		// 	println("\t", key, "\t", value)
		// }
		for field := range t.header {
			fmt.Printf(PRINT_FIELD_CELL_FORMAT, object[t.header[field]])
		}
		print("\n")
	}
}


func ParseCSV(fileName string, delimiter string, hasHeader bool) Table {
	if DEBUG {
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


func formResult(data []map[string]string, _select []string, _as []string, _limit int) Table {
	if DEBUG {
		println("[_DEBUG] formResult")
		print("SELECT ")
		print_array(_select)
		print("AS ")
		print_array(_as)
	}

	result_data := []map[string]string {}

	for i, row := range data {
		if i == _limit {
			break
		}

		select_fields := map[string]string {}

		for key, value := range row {
			as_index := slices.Index(_select, key)
			
			if as_index != -1 {
				if len(_as) == 0 || as_index >= len(_as) {
					select_fields[key] = value
					// result_header = append(result_header, key)
				} else {
					select_fields[_as[as_index]] = value
					// result_header = append(result_header, _as[as_index])
				}
			}
		}

		if len(select_fields) != 0 {
			result_data = append(result_data, select_fields)
		}
	}

	result_header := slices.Concat(_as, _select[len(_as):])

	return Table{
		data: result_data,
		header: result_header,
	}
}
