package csv_viewer

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

var (
	fileContent []map[string]interface{}
	header []string
	fileName string
)


/*
	Gets the contents of _fileName as array of lines
*/
func readFile(_fileName string) []string {
	// Open file
	file, err := os.Open(_fileName)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return []string{}
	}
	defer file.Close()

	// Creates scanner to read file line by line
	scanner := bufio.NewScanner(file)
	var content []string
	for scanner.Scan() {
		text := scanner.Text()
		content = append(content, text)
		// fmt.Println(text)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения:", err)
	}

	// Returns array of lines from file
	return content
}

/*
	Parses _data as array of strings, where useful blocks separated by _delimiter-s
*/
func parseData(_data []string, _delimiter string) [][]string {
	var parsed_content [][]string

	for _, line := range _data {
		parsed_content = append(parsed_content, strings.Split(line, _delimiter))
	}

	return parsed_content
}

/*
	Prepares CSV data for a coherent read later
*/
func readCSV(_filename string, _delimiter string) []map[string]interface{} {
	fileContent = []map[string]interface{} {}
	fileContent_ := parseData(readFile(_filename), _delimiter)

	header = fileContent_[0]

	for _, row := range fileContent_ {
		dict := make(map[string]interface{})
		// dict["ID"] = i
		// THERE SHOULD BE TYPE CONVERSIONS TO INT - FLOAT - STRING
		// println(row[len(row)-1])
		for j, cell := range row {
			dict[header[j]] = cell
		}
		fileContent = append(fileContent, dict)
	}

	fileName = _filename

	return fileContent
}
