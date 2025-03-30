<style>
:root {
	--color-optional: rgb(235, 200, 104);
}
</style>

# "Opal" - CSV Viewer
---

## Features
- Use SQL queries in a form of chain of methods, ex.:
```go
table	:= CSV_SQL.ParseCSV("file_name.csv")
result	:= CSV_SQL.Select("value").As("field").From(table).Result
```
---

### TODO

- pretty print
- if as is shorted then select - only as fields will be shown
- where
- group by
- sort by
<br>
- top
- queries that start not with select
<br>
- csv_viewer app GUI