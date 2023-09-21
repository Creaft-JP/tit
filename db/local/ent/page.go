// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Creaft-JP/tit/db/local/ent/page"
)

// Page is the model entity for the Page schema.
type Page struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Pathname holds the value of the "pathname" field.
	Pathname string `json:"pathname,omitempty"`
	// OrderWithinSiblings holds the value of the "order_within_siblings" field.
	OrderWithinSiblings int `json:"order_within_siblings,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// OverviewSentence holds the value of the "overview_sentence" field.
	OverviewSentence string `json:"overview_sentence,omitempty"`
	selectValues     sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Page) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case page.FieldID, page.FieldOrderWithinSiblings:
			values[i] = new(sql.NullInt64)
		case page.FieldPathname, page.FieldTitle, page.FieldOverviewSentence:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Page fields.
func (pa *Page) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case page.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pa.ID = int(value.Int64)
		case page.FieldPathname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pathname", values[i])
			} else if value.Valid {
				pa.Pathname = value.String
			}
		case page.FieldOrderWithinSiblings:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_within_siblings", values[i])
			} else if value.Valid {
				pa.OrderWithinSiblings = int(value.Int64)
			}
		case page.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				pa.Title = value.String
			}
		case page.FieldOverviewSentence:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field overview_sentence", values[i])
			} else if value.Valid {
				pa.OverviewSentence = value.String
			}
		default:
			pa.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Page.
// This includes values selected through modifiers, order, etc.
func (pa *Page) Value(name string) (ent.Value, error) {
	return pa.selectValues.Get(name)
}

// Update returns a builder for updating this Page.
// Note that you need to call Page.Unwrap() before calling this method if this Page
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Page) Update() *PageUpdateOne {
	return NewPageClient(pa.config).UpdateOne(pa)
}

// Unwrap unwraps the Page entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pa *Page) Unwrap() *Page {
	_tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Page is not a transactional entity")
	}
	pa.config.driver = _tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Page) String() string {
	var builder strings.Builder
	builder.WriteString("Page(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pa.ID))
	builder.WriteString("pathname=")
	builder.WriteString(pa.Pathname)
	builder.WriteString(", ")
	builder.WriteString("order_within_siblings=")
	builder.WriteString(fmt.Sprintf("%v", pa.OrderWithinSiblings))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(pa.Title)
	builder.WriteString(", ")
	builder.WriteString("overview_sentence=")
	builder.WriteString(pa.OverviewSentence)
	builder.WriteByte(')')
	return builder.String()
}

// Pages is a parsable slice of Page.
type Pages []*Page
