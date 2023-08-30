package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Remote holds the schema definition for the Remote entity.
type Remote struct {
	ent.Schema
}

// Fields of the Remote.
func (Remote) Fields() []ent.Field {
	return []ent.Field{field.String("name").NotEmpty().Unique(), field.String("url").NotEmpty()}
}

// Edges of the Remote.
func (Remote) Edges() []ent.Edge {
	return nil
}
