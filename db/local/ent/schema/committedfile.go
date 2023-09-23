package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CommittedFile holds the schema definition for the CommittedFile entity.
type CommittedFile struct {
	ent.Schema
}

// Fields of the CommittedFile.
func (CommittedFile) Fields() []ent.Field {
	return []ent.Field{field.String("path").NotEmpty(), field.String("content")}
}

// Edges of the CommittedFile.
func (CommittedFile) Edges() []ent.Edge {
	return []ent.Edge{edge.From("commit", Commit.Type).Ref("files").Unique()}
}
