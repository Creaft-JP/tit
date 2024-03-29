// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Creaft-JP/tit/db/local/ent/stagedfile"
)

// StagedFile is the model entity for the StagedFile schema.
type StagedFile struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Path holds the value of the "path" field.
	Path string `json:"path,omitempty"`
	// Content holds the value of the "content" field.
	Content      string `json:"content,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*StagedFile) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case stagedfile.FieldID:
			values[i] = new(sql.NullInt64)
		case stagedfile.FieldPath, stagedfile.FieldContent:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the StagedFile fields.
func (sf *StagedFile) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case stagedfile.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sf.ID = int(value.Int64)
		case stagedfile.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				sf.Path = value.String
			}
		case stagedfile.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				sf.Content = value.String
			}
		default:
			sf.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the StagedFile.
// This includes values selected through modifiers, order, etc.
func (sf *StagedFile) Value(name string) (ent.Value, error) {
	return sf.selectValues.Get(name)
}

// Update returns a builder for updating this StagedFile.
// Note that you need to call StagedFile.Unwrap() before calling this method if this StagedFile
// was returned from a transaction, and the transaction was committed or rolled back.
func (sf *StagedFile) Update() *StagedFileUpdateOne {
	return NewStagedFileClient(sf.config).UpdateOne(sf)
}

// Unwrap unwraps the StagedFile entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sf *StagedFile) Unwrap() *StagedFile {
	_tx, ok := sf.config.driver.(*txDriver)
	if !ok {
		panic("ent: StagedFile is not a transactional entity")
	}
	sf.config.driver = _tx.drv
	return sf
}

// String implements the fmt.Stringer.
func (sf *StagedFile) String() string {
	var builder strings.Builder
	builder.WriteString("StagedFile(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sf.ID))
	builder.WriteString("path=")
	builder.WriteString(sf.Path)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(sf.Content)
	builder.WriteByte(')')
	return builder.String()
}

// StagedFiles is a parsable slice of StagedFile.
type StagedFiles []*StagedFile
