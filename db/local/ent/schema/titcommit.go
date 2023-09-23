package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TitCommit holds the schema definition for the TitCommit entity.
type TitCommit struct {
	ent.Schema
}

// Fields of the TitCommit.
func (TitCommit) Fields() []ent.Field {
	return []ent.Field{field.Int("number").Positive(), field.String("message").NotEmpty()}
}

// Edges of the TitCommit.
func (TitCommit) Edges() []ent.Edge {
	return []ent.Edge{edge.To("files", CommittedFile.Type)}
}
