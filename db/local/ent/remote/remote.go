// Code generated by ent, DO NOT EDIT.

package remote

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the remote type in the database.
	Label = "remote"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// Table holds the table name of the remote in the database.
	Table = "remotes"
)

// Columns holds all SQL columns for remote fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldURL,
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
	// URLValidator is a validator for the "url" field. It is called by the builders before save.
	URLValidator func(string) error
)

// OrderOption defines the ordering options for the Remote queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByURL orders the results by the url field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}
