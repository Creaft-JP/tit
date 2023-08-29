package schema

import "entgo.io/ent"

// Remote holds the schema definition for the Remote entity.
type Remote struct {
	ent.Schema
}

// Fields of the Remote.
func (Remote) Fields() []ent.Field {
	return nil
}

// Edges of the Remote.
func (Remote) Edges() []ent.Edge {
	return nil
}
