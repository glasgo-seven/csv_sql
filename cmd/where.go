package csv_sql

import (
	"strings"
)


type WhereQuery struct {
	parent FromQuery
	Result []map[string]string
}


func (q FromQuery) Where(condition string) WhereQuery {
	// where_data := []map[string]string {}

	if strings.Contains(condition, ">") {
		condition = "GT"
	} else if strings.Contains(condition, ">=") {
		condition = "GTE"
	} else if strings.Contains(condition, "<") {
		condition = "LT"
	} else if strings.Contains(condition, "<=") {
		condition = "LTE"
	} else if strings.Contains(condition, "==") {
		condition = "E"
	} else if strings.Contains(condition, "!=") {
		condition = "NE"
	} else {
		condition = ""
		panic("unknown comparison operator")
	}

	// for _, row := range q.data {
	// 	if 

	// 	where_data = append(where_data, select_fields)
	// }

	var result []map[string]string

	return WhereQuery{
		parent: q,
		Result: result,
	}
}

