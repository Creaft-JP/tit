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
	return []ent.Field{field.String("key").NotEmpty().Unique(), field.String("value").NotEmpty()}
}

// Edges of the GlobalConfig.
func (GlobalConfig) Edges() []ent.Edge {
	return nil
}
