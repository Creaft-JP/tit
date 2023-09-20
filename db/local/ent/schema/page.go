package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Page holds the schema definition for the Page entity.
type Page struct {
	ent.Schema
}

// Fields of the Page.
func (Page) Fields() []ent.Field {
	return []ent.Field{
		field.Strings("path"),
		field.Int("order_within_siblings").Positive(),
		field.String("title"),
		field.String("overview_sentence"),
	}
}

// Edges of the Page.
func (Page) Edges() []ent.Edge {
	return nil
}
