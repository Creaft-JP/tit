package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// GlobalConfig holds the schema definition for the GlobalConfig entity.
type GlobalConfig struct {
	ent.Schema
}

// Fields of the GlobalConfig.
func (GlobalConfig) Fields() []ent.Field {
	return []ent.Field{field.String("key").Unique(), field.String("value")}
}

// Edges of the GlobalConfig.
func (GlobalConfig) Edges() []ent.Edge {
	return nil
}
