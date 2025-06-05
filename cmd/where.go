package csv_sql

import (
	// "strings"
)

type LogicOperator int
const (
	NOT LogicOperator = iota
	AND
	OR
)

type LogicExpression struct {
	op LogicOperator
	field_a string
	field_b string
}


type WhereQuery struct {
	parent FromQuery
	Result []map[string]string
}

func parseLogic(condition string) []map[string]LogicExpression {
	dict := []map[string]LogicExpression {}

	// bracket_open := strings.Index(condition, "(")
	// bracket_close := strings.Index(condition, ")")

	// if bracket_open != -1 && bracket_close != -1 && bracket_open < bracket_close {
	// 	brackets_dict := parseLogic(condition[bracket_open:bracket_close])
	// } else if bracket_open == -1 && bracket_close == -1 {
		
	// } else {
	// 	panic("Brackets error")
	// }

	return dict
}


func (q FromQuery) Where(condition string) WhereQuery {
	// where_data := []map[string]string {}

	// * ORDER:
	// () NOT AND OR >< math

	// dict := parseLogic(condition)

	var result []map[string]string

	return WhereQuery{
		parent: q,
		Result: result,
	}
}

