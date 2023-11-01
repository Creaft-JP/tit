package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Image holds the schema definition for the Image entity.
type Image struct {
	ent.Schema
}

// Fields of the Image.
func (Image) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Unique().Default(uuid.New),
		field.String("extension").NotEmpty(),
		field.Bytes("contents"),
		field.String("description"),
	}
}

// Edges of the Image.
func (Image) Edges() []ent.Edge {
	return []ent.Edge{edge.From("commit", TitCommit.Type).Ref("images")}
}
