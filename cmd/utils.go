package csv_sql

var DEBUG bool = false

func print_array(array []string) {
	print("[")
	for i, value := range array {
		if i == 0 {
			print(value)
		} else {
			print(", ", value)
		}
		
	}
	print("]\n")
}

func print_map(dict []map[string]string) {
	print("[\n")
	for _, row := range dict{
		print("\t{\n")
		for key, value := range row {
			print("\t\t", key, ": ", value, ",\n")
		}
		print("\t},\n")
	}
	print("]\n")
}