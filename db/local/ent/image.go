// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Creaft-JP/tit/db/local/ent/image"
	"github.com/google/uuid"
)

// Image is the model entity for the Image schema.
type Image struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Extension holds the value of the "extension" field.
	Extension string `json:"extension,omitempty"`
	// Contents holds the value of the "contents" field.
	Contents []byte `json:"contents,omitempty"`
	// Number holds the value of the "number" field.
	Number int `json:"number,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ImageQuery when eager-loading is set.
	Edges        ImageEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ImageEdges holds the relations/edges for other nodes in the graph.
type ImageEdges struct {
	// Commit holds the value of the commit edge.
	Commit []*TitCommit `json:"commit,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CommitOrErr returns the Commit value or an error if the edge
// was not loaded in eager-loading.
func (e ImageEdges) CommitOrErr() ([]*TitCommit, error) {
	if e.loadedTypes[0] {
		return e.Commit, nil
	}
	return nil, &NotLoadedError{edge: "commit"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Image) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case image.FieldContents:
			values[i] = new([]byte)
		case image.FieldNumber:
			values[i] = new(sql.NullInt64)
		case image.FieldExtension, image.FieldDescription:
			values[i] = new(sql.NullString)
		case image.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Image fields.
func (i *Image) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case image.FieldID:
			if value, ok := values[j].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[j])
			} else if value != nil {
				i.ID = *value
			}
		case image.FieldExtension:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field extension", values[j])
			} else if value.Valid {
				i.Extension = value.String
			}
		case image.FieldContents:
			if value, ok := values[j].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field contents", values[j])
			} else if value != nil {
				i.Contents = *value
			}
		case image.FieldNumber:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field number", values[j])
			} else if value.Valid {
				i.Number = int(value.Int64)
			}
		case image.FieldDescription:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[j])
			} else if value.Valid {
				i.Description = value.String
			}
		default:
			i.selectValues.Set(columns[j], values[j])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Image.
// This includes values selected through modifiers, order, etc.
func (i *Image) Value(name string) (ent.Value, error) {
	return i.selectValues.Get(name)
}

// QueryCommit queries the "commit" edge of the Image entity.
func (i *Image) QueryCommit() *TitCommitQuery {
	return NewImageClient(i.config).QueryCommit(i)
}

// Update returns a builder for updating this Image.
// Note that you need to call Image.Unwrap() before calling this method if this Image
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Image) Update() *ImageUpdateOne {
	return NewImageClient(i.config).UpdateOne(i)
}

// Unwrap unwraps the Image entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Image) Unwrap() *Image {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Image is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Image) String() string {
	var builder strings.Builder
	builder.WriteString("Image(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("extension=")
	builder.WriteString(i.Extension)
	builder.WriteString(", ")
	builder.WriteString("contents=")
	builder.WriteString(fmt.Sprintf("%v", i.Contents))
	builder.WriteString(", ")
	builder.WriteString("number=")
	builder.WriteString(fmt.Sprintf("%v", i.Number))
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(i.Description)
	builder.WriteByte(')')
	return builder.String()
}

// Images is a parsable slice of Image.
type Images []*Image
