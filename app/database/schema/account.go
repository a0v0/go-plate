package schema

import (
	"go_plate/pkg/id"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Immutable().
			NotEmpty().
			Unique().
			DefaultFunc(id.NewAccountID),

		field.String("user_id").
			NotEmpty().
			Immutable().
			Unique(),

		field.Time("disable_time").
			Optional(),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile", Profile.Type).
			Unique(),
	}
}
