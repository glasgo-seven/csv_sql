# csv_sql / Documentation
---

## Contents
1. [Register Table](#registerTable)
   1.1 [LoadCSV](#LoadCSV)
   1.2 [SetHeader](#SetHeader)
   1.3 [SetTypes](#SetTypes)
2. [Functions](#Functions)
3. [SELECT](#SELECT)
   2.1 [Select](#Select)
   2.2 [As](#As)
4. [FROM](#FROM)
	3.1 [From](#From)
	3.2 [Limit](#Limit)
5. [WHERE](#WHERE)
   4.1 [Where](#Where)
----

## 1. Register Table<a name="registerTable"></a>
### 1.1 LoadCSV<a name="LoadCSV"></a>
Used for registering a .csv file as a table.
```go
func LoadCSV(fileName string, delimiter string, tableName string) {}
```
Example:
```go
sql_csv.LoadCSV("./rsc/file.csv", ",", "new_table")
```
Recommended:
- use [SetHeader](#SetHeader) after.
	*By default table is header-less*.
- use [SetTypes](#SetTypes) after.
	*By default all values in a table are of type "string"*.

### 1.2 SetHeader<a name="SetHeader"></a>
Sets a header of a table.
```go
func SetHeader(tableName string, header []string) {}
```
Example:
```go
sql_csv.SetHeader("new_table", []string {"Field_1", "Field_2"} ) {}
```

### 1.3 SetTypes<a name="SetTypes"></a>
Sets types for columns of a table.
```go
func SetTypes(tableName string, types []string) {}
```
Example:
```go
sql_csv.SetTypes("new_table", []string {"string", "int"} ) {}
```
---

## 2. Functions<a name="Functions"></a>
	MIN  returns the smallest value within the selected column
	MAX  returns the largest value within the selected column
	COUNT  returns the number of rows in a set
	SUM  returns the total sum of a numerical column
	AVG  returns the average value of a numerical column
---

## 3. SELECT<a name="SELECT"></a>
### Select
Select only the specified table columns, or all if "*" is passed.
```go
func Select(fields ...string)  {}
```
Example:
```go
//	SELECT * ...
sql_csv.Select("*")

//	SELECT Field_1, Field_2 ...
sql_csv.Select("Field_1", "Field_2")
```

### As
Sets a name alias for selected columns of a table.
Number of arguments must be less or equal to the number of [Select()]() arguments.
If column should have no alias - use *nil* as argument.
If [Select()]() has only one arguments - "*" (all columns) - As() cannot be used.
```go
func As(fields ...string) {}
```
Example:
```go
//	SELECT Field as Value ...
sql_csv.Select("Field").As("Value")

//	SELECT Field_1, Field_2 as Value ...
sql_csv.Select("Field_1", "Field_2").As(nil, "Value2")
```
---

## 4. FROM<a name="FROM"></a>
### 4.1 From<a name="From"></a>
Selects a table for query to be executed on.
```go
func From(table string) {}
```
Example:
```go
//	SELECT * FROM new_table
sql_csv.Select("*").From("new_table")

//	SELECT Field_1 FROM new_table
sql_csv.Select("Field_1").From("new_table")

//	SELECT Field_1 AS Value FROM new_table
sql_csv.Select("Field_1").As("Value").From("new_table")
```

### 4.2 Limit<a name="Limit"></a>
Used for limiting the number of result rows.
```go
func Limit(limit int) {}
```
Example:
```go
//	SELECT * FROM new_table LIMIT 5
sql_csv.Select("*").From("new_table").Limit(5)
```
---

## 5. WHERE<a name="WHERE"></a>
```json
"(field_a > X OR field_b < Y) AND NOT field_c = Z"
{
	"T0" : ("GT", "field_a", "X"),
	"T1" : ("LT", "field_b", "Y"),
	"L0" : ("OR", "T0", "T1"),
	"T2" : ("EQ", "field_c", "Z"),
	"L1" : ("NOT", "T2"),
	"L2" : ("AND", "L0", "L1")
}
```
---


