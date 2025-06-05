package csv_sql

import (
	
)

type FromQuery struct {
	parent SelectQuery
	table Table
	Result Table
}

func (q SelectQuery) From(table Table) FromQuery {
	if DEBUG {
		println("[_DEBUG] From")
	}
	
	return FromQuery {
		parent: q,
		table: table,
		Result: formResult(table.data, q.selected_fields, q.as_fields, -1),
	}
}

func (q FromQuery) Limit(limit int) FromQuery {
	if DEBUG {
		println("[_DEBUG] Limit")
	}
	
	return FromQuery {
		parent: q.parent,
		Result: formResult(q.table.data, q.parent.selected_fields, q.parent.as_fields, limit),
	}
}