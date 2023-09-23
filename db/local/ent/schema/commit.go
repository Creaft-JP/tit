package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Commit holds the schema definition for the Commit entity.
type Commit struct {
	ent.Schema
}

// Fields of the Commit.
func (Commit) Fields() []ent.Field {
	return []ent.Field{field.Int("number").Positive(), field.String("message").NotEmpty()}
}

// Edges of the Commit.
func (Commit) Edges() []ent.Edge {
	return []ent.Edge{edge.To("files", CommittedFile.Type)}
}
