// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Creaft-JP/tit/db/local/ent/page"
	"github.com/Creaft-JP/tit/db/local/ent/section"
)

// Section is the model entity for the Section schema.
type Section struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// OverviewSentence holds the value of the "overview_sentence" field.
	OverviewSentence string `json:"overview_sentence,omitempty"`
	// Number holds the value of the "number" field.
	Number int `json:"number,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SectionQuery when eager-loading is set.
	Edges         SectionEdges `json:"edges"`
	page_sections *int
	selectValues  sql.SelectValues
}

// SectionEdges holds the relations/edges for other nodes in the graph.
type SectionEdges struct {
	// Page holds the value of the page edge.
	Page *Page `json:"page,omitempty"`
	// Commits holds the value of the commits edge.
	Commits []*TitCommit `json:"commits,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PageOrErr returns the Page value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SectionEdges) PageOrErr() (*Page, error) {
	if e.loadedTypes[0] {
		if e.Page == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: page.Label}
		}
		return e.Page, nil
	}
	return nil, &NotLoadedError{edge: "page"}
}

// CommitsOrErr returns the Commits value or an error if the edge
// was not loaded in eager-loading.
func (e SectionEdges) CommitsOrErr() ([]*TitCommit, error) {
	if e.loadedTypes[1] {
		return e.Commits, nil
	}
	return nil, &NotLoadedError{edge: "commits"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Section) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case section.FieldID, section.FieldNumber:
			values[i] = new(sql.NullInt64)
		case section.FieldSlug, section.FieldTitle, section.FieldOverviewSentence:
			values[i] = new(sql.NullString)
		case section.ForeignKeys[0]: // page_sections
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Section fields.
func (s *Section) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case section.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case section.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				s.Slug = value.String
			}
		case section.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				s.Title = value.String
			}
		case section.FieldOverviewSentence:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field overview_sentence", values[i])
			} else if value.Valid {
				s.OverviewSentence = value.String
			}
		case section.FieldNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field number", values[i])
			} else if value.Valid {
				s.Number = int(value.Int64)
			}
		case section.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field page_sections", value)
			} else if value.Valid {
				s.page_sections = new(int)
				*s.page_sections = int(value.Int64)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Section.
// This includes values selected through modifiers, order, etc.
func (s *Section) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryPage queries the "page" edge of the Section entity.
func (s *Section) QueryPage() *PageQuery {
	return NewSectionClient(s.config).QueryPage(s)
}

// QueryCommits queries the "commits" edge of the Section entity.
func (s *Section) QueryCommits() *TitCommitQuery {
	return NewSectionClient(s.config).QueryCommits(s)
}

// Update returns a builder for updating this Section.
// Note that you need to call Section.Unwrap() before calling this method if this Section
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Section) Update() *SectionUpdateOne {
	return NewSectionClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Section entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Section) Unwrap() *Section {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Section is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Section) String() string {
	var builder strings.Builder
	builder.WriteString("Section(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("slug=")
	builder.WriteString(s.Slug)
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(s.Title)
	builder.WriteString(", ")
	builder.WriteString("overview_sentence=")
	builder.WriteString(s.OverviewSentence)
	builder.WriteString(", ")
	builder.WriteString("number=")
	builder.WriteString(fmt.Sprintf("%v", s.Number))
	builder.WriteByte(')')
	return builder.String()
}

// Sections is a parsable slice of Section.
type Sections []*Section
