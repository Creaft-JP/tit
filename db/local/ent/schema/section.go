package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Section holds the schema definition for the Section entity.
type Section struct {
	ent.Schema
}

// Fields of the Section.
func (Section) Fields() []ent.Field {
	return []ent.Field{
		field.String("slug").NotEmpty(),
		field.String("title").MaxLen(50),
		field.String("overview_sentence"),
		field.Int("number").Positive(),
	}
}

// Edges of the Section.
func (Section) Edges() []ent.Edge {
	return []ent.Edge{edge.From("page", Page.Type).Ref("sections").Unique()}
}
