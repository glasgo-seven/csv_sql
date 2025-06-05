package csv_sql

import "slices"

type SelectQuery struct {
	selected_fields []string
	as_fields []string
	is_all bool
}

func Select(fields ...string) SelectQuery {
	if DEBUG {
		println("[_DEBUG] Select")
	}

	if slices.Contains(fields, "*") {
		if len(fields) != 1 {
			panic("Can only use \"*\" without other fields.")
		} else {
			return SelectQuery{
				selected_fields: []string {},
				as_fields: []string {},
				is_all: true,
			}
		}
	}

	return SelectQuery{
		selected_fields: fields,
		as_fields: []string {},
		is_all: false,
	}
}

func (q SelectQuery) As(fields ...string) SelectQuery {
	if DEBUG {
		println("[_DEBUG] As")
	}

	if len(fields) > len(q.selected_fields) {
		panic("Number of named field should be less or equal to the number of selected fields.")
	}
	return SelectQuery{
		selected_fields: q.selected_fields,
		as_fields: fields,
		is_all: q.is_all,
	}
}