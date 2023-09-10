package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// StagedFile holds the schema definition for the StagedFile entity.
type StagedFile struct {
	ent.Schema
}

// Fields of the StagedFile.
func (StagedFile) Fields() []ent.Field {
	return []ent.Field{field.String("path").NotEmpty().Unique(), field.String("content").NotEmpty()}
}

// Edges of the StagedFile.
func (StagedFile) Edges() []ent.Edge {
	return nil
}
