package GetData

import (
	"fmt"
)

const (
	CONCAT                              = "CONCAT"
	COLUMN_NAME                         = "COLUMN_NAME"
	COLUMN_TYPE                         = "COLUMN_TYPE"
	COLUMN_COMMENT                      = "COLUMN_COMMENT"
	INFORMATION_SCHEMA_TABLES           = "INFORMATION_SCHEMA.`TABLES`"
	INFORMATION_SCHEMA_COLUMNS          = "INFORMATION_SCHEMA.`COLUMNS`"
	INFORMATION_SCHEMA_KEY_COLUMN_USAGE = "INFORMATION_SCHEMA.`KEY_COLUMN_USAGE`"
	REFERENCED_COLUMN_NAME              = "REFERENCED_COLUMN_NAME"
	REFERENCED_TABLE_NAME               = "REFERENCED_TABLE_NAME"
	TABLE_COMMENT                       = "TABLE_COMMENT"
	TABLE_SCHEMA                        = "TABLE_SCHEMA"
	TABLE_NAME                          = "TABLE_NAME"
)

var (
	DataBaseName = ""
	IsTime       = ""
)

func GetDataBaseStruct() string {
	sqlDataBaseStruct := fmt.Sprintf(
		"SELECT %v,%v FROM %v WHERE %v = ?;",
		TABLE_NAME, TABLE_COMMENT,
		INFORMATION_SCHEMA_TABLES,
		TABLE_SCHEMA)
	return sqlDataBaseStruct
}

func GetDataBaseTableStruct() string {
	sqlTableStruct := fmt.Sprintf(
		"SELECT %v ,%v ,%v FROM %v WHERE %v LIKE ? AND %v LIKE ?;",
		COLUMN_NAME, COLUMN_TYPE,
		COLUMN_COMMENT,
		INFORMATION_SCHEMA_COLUMNS,
		TABLE_SCHEMA, TABLE_NAME)
	return sqlTableStruct
}

func GetDataBaseTablePK() string {
	sqlTablePK := fmt.Sprintf(
		"SELECT %v(%v), %v(%v) FROM %v WHERE %v = ? AND %v > %v",
		CONCAT, TABLE_NAME,
		CONCAT, REFERENCED_TABLE_NAME,
		INFORMATION_SCHEMA_KEY_COLUMN_USAGE,
		TABLE_SCHEMA, REFERENCED_TABLE_NAME, "\"\"")
	return sqlTablePK
}
