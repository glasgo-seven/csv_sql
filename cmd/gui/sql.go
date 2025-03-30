package csv_viewer

import (
	"strings"
	// "unicode"
)

type SQL_Result struct {
	err string
	result []map[string]interface{}
}

// select 'Overall rank' from table where 'Overall rank' = 1;
// select'Overallrank'fromtablewhere'Overallrank'=1;


func stripCharacter(str string, c rune) string {
	// return str
	i_left := 0
	i_right := len(str)-1

	for rune(str[i_left]) == c {
		i_left++
	}
	for rune(str[i_right]) == c {
		i_right--
	}

	return str[i_left:i_right+1]
}


func parseQuery() []string {
	if len(fileContent) == 0 && fileName == "" {
		return []string{"NO FILE IS OPENNED"}
	} else if len(fileContent) == 0 {
		return []string{"FILE IS EMPTY"}
	}

	query := query_editor_input.Text
	QUERY := strings.ToUpper(query)

	SEMICOLON_index := strings.Index(QUERY, ";")
	if SEMICOLON_index == -1 {
		println("NO SEMICOLON PRESENT")
		return []string{"NO SEMICOLON PRESENT"}
	} else if SEMICOLON_index != len(query)-1 {
		println("SEMICOLON WRONGLY PLACED")
		return []string{"SEMICOLON WRONGLY PLACED"}
	}

	SELECT_index := strings.Index(QUERY, "SELECT ")
	if SELECT_index == -1 {
		println("NO SELECT PRESENT")
		return []string{"NO SELECT PRESENT"}
	} else if SELECT_index != 0  {
		println("SELECT WRONGLY PLACED")
		return []string{"SELECT WRONGLY PLACED"}
	}
	
	FROM_index := strings.Index(QUERY, " FROM ")
	if FROM_index == -1 {
		println("NO FROM COMMAND")
		return []string{"NO FROM COMMAND"}
	}

	WHERE_index := strings.Index(QUERY, " WHERE ")
	GROUPBY_index := strings.Index(QUERY, " GROUP BY ")
	ORDERBY_index := strings.Index(QUERY, " ORDER BY ")

	_case := ""
	if WHERE_index == -1 {
		_case += "0"
	} else {
		_case += "1"
	}
	if GROUPBY_index == -1 {
		_case += "0"
	} else {
		_case += "1"
	}
	if ORDERBY_index == -1 {
		_case += "0"
	} else {
		_case += "1"
	}

	SELECT	:= query[7:FROM_index]
	FROM	:= ""
	WHERE	:= "NONE"
	GROUPBY	:= "NONE"
	ORDERBY	:= "NONE"

	switch _case {
	//   W G O
	case "000":
		FROM	= query[FROM_index+6:SEMICOLON_index]
	case "001":
		FROM	= query[FROM_index+6:ORDERBY_index]
		ORDERBY	= query[ORDERBY_index+10:SEMICOLON_index]
	case "010":
		FROM	= query[FROM_index+6:GROUPBY_index]
		GROUPBY	= query[GROUPBY_index+10:SEMICOLON_index]
	case "011":
		FROM	= query[FROM_index+6:GROUPBY_index]
		GROUPBY	= query[GROUPBY_index+10:ORDERBY_index]
		ORDERBY	= query[ORDERBY_index+10:SEMICOLON_index]
	case "100":
		FROM	= query[FROM_index+6:WHERE_index]
		WHERE	= query[WHERE_index+7:SEMICOLON_index]
	case "101":
		FROM	= query[FROM_index+6:WHERE_index]
		WHERE	= query[WHERE_index+7:ORDERBY_index]
		ORDERBY	= query[ORDERBY_index+10:SEMICOLON_index]
	case "110":
		FROM	= query[FROM_index+6:WHERE_index]
		WHERE	= query[WHERE_index+7:GROUPBY_index]
		GROUPBY	= query[GROUPBY_index+10:SEMICOLON_index]
	case "111":
		FROM	= query[FROM_index+6:WHERE_index]
		WHERE	= query[WHERE_index+7:GROUPBY_index]
		GROUPBY	= query[GROUPBY_index+10:ORDERBY_index]
		ORDERBY	= query[ORDERBY_index+10:SEMICOLON_index]
	}

	print("SELECT [", SELECT, "] FROM [", FROM, "] WHERE [", WHERE, "] GROUP BY [", GROUPBY, "] ORDER BY [", ORDERBY, "] ;")

	return []string {SELECT, FROM, WHERE, GROUPBY, ORDERBY}
}



func executeQuery() SQL_Result {
	
	parsedQuery := parseQuery()

	if len(parsedQuery) == 1 {
		return SQL_Result{parsedQuery[0], nil }
	}

	// println(query_editor_input.Text)
	// for _, block := range parsedQuery {
	// 	println(block)
	// }

	// * SELECT fields
	query_fields := []string {}
	query_result := []map[string]interface{} {}

	if parsedQuery[0] == "*" {
		query_fields = header
	} else {
		for _, field := range strings.Split(parsedQuery[0], ",") {
			query_fields = append(query_fields, stripCharacter( stripCharacter(field, '\''), ' ') )
		}
	}

	println("\nSELECT")
	for _, field := range query_fields {
		println("\t", field)
	}

	// * WHERE conditions
	// TODO - Brackets support
	where := parsedQuery[2]
	where_conditions := map[int]string {}
	i := 0
	and := strings.Index(strings.ToUpper(where), "AND")
	for and != -1 {
		where_conditions[i] = stripCharacter(where[:and], ' ')
		where = where[and+3:]
		and = strings.Index(strings.ToUpper(where), "AND")
		i++
	}
	where_conditions[i] = stripCharacter(where, ' ')
	
	println("WHERE")
	for i, condition := range where_conditions {
		println("\t", i, condition)
	}





	return SQL_Result{"SUCCESS1", query_result}
}

// select 'id', value from table where value > 300 and value < 500;

func conditionsEvaluator(condition string) []int {
	and := strings.Index(strings.ToUpper(condition), "AND")
	or := strings.Index(strings.ToUpper(condition), "OR")

	if and != -1 {
		evaluateConditionsAnd(condition[:and-1], condition[and+4:])
	} else 
}