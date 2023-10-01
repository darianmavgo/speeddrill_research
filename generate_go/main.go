package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type TableInfo struct {
	TableName string
	Columns   []ColumnInfo
}

type ColumnInfo struct {
	ColumnName string
	ColumnType string
	Default    string // Default value for the column
	Pk         string
	Dontcare   string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <sqlite3_file>")
		return
	}

	dbFile := os.Args[1]

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatalf("Failed to open database: %v\n", err)
	}
	defer db.Close()

	tables, err := getTables(db)
	if err != nil {
		log.Fatalf("Failed to get table information: %v\n", err)
	}

	for _, table := range tables {
		fmt.Printf("type %s struct {\n", toCamelCase(table.TableName))
		for _, column := range table.Columns {
			fmt.Printf("\t%s %s\n", toCamelCase(column.ColumnName), mapSQLiteTypeToGoType(column.ColumnType))
		}
		fmt.Println("}")
	}
}

func getTables(db *sql.DB) ([]TableInfo, error) {
	rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []TableInfo

	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			return nil, err
		}

		columns, err := getColumns(db, tableName)
		if err != nil {
			return nil, err
		}

		tableInfo := TableInfo{
			TableName: tableName,
			Columns:   columns,
		}

		tables = append(tables, tableInfo)
	}

	return tables, nil
}

func getColumns(db *sql.DB, tableName string) ([]ColumnInfo, error) {
	rows, err := db.Query(fmt.Sprintf("PRAGMA table_info(%s)", tableName))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var columns []ColumnInfo

	for rows.Next() {
		var cid int
		var columnName string
		var columnType string
		var notNull int
		var pK string
		// var dontcare string
		var defaultValue sql.NullString
		err := rows.Scan(&cid, &columnName, &columnType, &notNull, &defaultValue, &pK)
		if err != nil {
			return nil, err
		}

		defaultValueStr := defaultValue.String
		if !defaultValue.Valid {
			defaultValueStr = "NULL"
		}

		columnInfo := ColumnInfo{
			ColumnName: columnName,
			ColumnType: columnType,
			Default:    defaultValueStr,
			// Pk:         pK,
			// Dontcare:   dontcare,
		}

		columns = append(columns, columnInfo)
	}

	return columns, nil
}

func mapSQLiteTypeToGoType(sqliteType string) string {
	switch strings.ToLower(sqliteType) {
	case "integer":
		return "int"
	case "text":
		return "string"
	case "real":
		return "float64"
	case "blob":
		return "[]byte"
	default:
		return "interface{}"
	}
}

func toCamelCase(input string) string {
	parts := strings.Split(input, "_")
	for i := 1; i < len(parts); i++ {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}

func printTableInfo(tables []TableInfo) {
	for _, table := range tables {
		fmt.Printf("Table: %s\n", table.TableName)
		fmt.Println("Columns:")
		for _, column := range table.Columns {
			fmt.Printf("\t%s: %s\n", column.ColumnName, column.ColumnType)
		}
		fmt.Println()
	}
}
