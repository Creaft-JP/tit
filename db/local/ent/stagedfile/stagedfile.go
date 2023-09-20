// Code generated by ent, DO NOT EDIT.

package stagedfile

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the stagedfile type in the database.
	Label = "staged_file"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// Table holds the table name of the stagedfile in the database.
	Table = "staged_files"
)

// Columns holds all SQL columns for stagedfile fields.
var Columns = []string{
	FieldID,
	FieldPath,
	FieldContent,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// PathValidator is a validator for the "path" field. It is called by the builders before save.
	PathValidator func(string) error
)

// OrderOption defines the ordering options for the StagedFile queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPath orders the results by the path field.
func ByPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPath, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
}