package csv_sql

import (
	
)

type FromQuery struct {
	parent SelectQuery
	table Table
	Result Table
}

func (q SelectQuery) From(table Table) FromQuery {
	if _DEBUG {
		println("[_DEBUG] From")
	}
	
	return FromQuery {
		parent: q,
		table: table,
		Result: Table{
			data: formResult(table.data, q.selected_fields, q.as_fields, -1),
			header: q.as_fields,
		},
	}
}

func (q FromQuery) Limit(limit int) FromQuery {
	if _DEBUG {
		println("[_DEBUG] Limit")
	}
	
	return FromQuery {
		parent: q.parent,
		Result: Table{
			data: formResult(q.table.data, q.parent.selected_fields, q.parent.as_fields, limit),
			header: q.parent.as_fields,
		},
	}
}