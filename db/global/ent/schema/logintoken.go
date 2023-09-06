package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// LoginToken holds the schema definition for the LoginToken entity.
type LoginToken struct {
	ent.Schema
}

// Fields of the LoginToken.
func (LoginToken) Fields() []ent.Field {
	return []ent.Field{field.String("sign_in_user_slug").NotEmpty().Unique(), field.String("cli_login_token").NotEmpty()}
}

// Edges of the LoginToken.
func (LoginToken) Edges() []ent.Edge {
	return nil
}
