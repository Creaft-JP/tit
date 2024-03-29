// Code generated by ent, DO NOT EDIT.

package section

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the section type in the database.
	Label = "section"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSlug holds the string denoting the slug field in the database.
	FieldSlug = "slug"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldOverviewSentence holds the string denoting the overview_sentence field in the database.
	FieldOverviewSentence = "overview_sentence"
	// FieldNumber holds the string denoting the number field in the database.
	FieldNumber = "number"
	// EdgePage holds the string denoting the page edge name in mutations.
	EdgePage = "page"
	// EdgeCommits holds the string denoting the commits edge name in mutations.
	EdgeCommits = "commits"
	// Table holds the table name of the section in the database.
	Table = "sections"
	// PageTable is the table that holds the page relation/edge.
	PageTable = "sections"
	// PageInverseTable is the table name for the Page entity.
	// It exists in this package in order to avoid circular dependency with the "page" package.
	PageInverseTable = "pages"
	// PageColumn is the table column denoting the page relation/edge.
	PageColumn = "page_sections"
	// CommitsTable is the table that holds the commits relation/edge.
	CommitsTable = "tit_commits"
	// CommitsInverseTable is the table name for the TitCommit entity.
	// It exists in this package in order to avoid circular dependency with the "titcommit" package.
	CommitsInverseTable = "tit_commits"
	// CommitsColumn is the table column denoting the commits relation/edge.
	CommitsColumn = "section_commits"
)

// Columns holds all SQL columns for section fields.
var Columns = []string{
	FieldID,
	FieldSlug,
	FieldTitle,
	FieldOverviewSentence,
	FieldNumber,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "sections"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"page_sections",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	SlugValidator func(string) error
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// NumberValidator is a validator for the "number" field. It is called by the builders before save.
	NumberValidator func(int) error
)

// OrderOption defines the ordering options for the Section queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySlug orders the results by the slug field.
func BySlug(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSlug, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByOverviewSentence orders the results by the overview_sentence field.
func ByOverviewSentence(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOverviewSentence, opts...).ToFunc()
}

// ByNumber orders the results by the number field.
func ByNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNumber, opts...).ToFunc()
}

// ByPageField orders the results by page field.
func ByPageField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPageStep(), sql.OrderByField(field, opts...))
	}
}

// ByCommitsCount orders the results by commits count.
func ByCommitsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCommitsStep(), opts...)
	}
}

// ByCommits orders the results by commits terms.
func ByCommits(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCommitsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPageStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PageInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PageTable, PageColumn),
	)
}
func newCommitsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CommitsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CommitsTable, CommitsColumn),
	)
}
