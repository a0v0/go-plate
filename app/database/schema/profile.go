package schema

import (
	"time"

	"go_plate/pkg/id"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Profile holds the schema definition for the Profile entity.
type Profile struct {
	ent.Schema
}

// Fields of the Profile.
func (Profile) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Immutable().
			NotEmpty().
			Unique().
			DefaultFunc(id.NewProfileID),

		field.Text("username").
			Unique().
			Optional(),

		field.Text("display_name").
			Optional(),

		field.Text("bio").
			Optional(),

		field.Text("lang_tag").
			Default("en-US"),

		field.Text("location").
			Optional(),

		field.Text("timezone").
			Optional(),

		field.Time("update_time").
			Default(time.Now).
			UpdateDefault(time.Now),

		field.Time("create_time").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the Profile.
func (Profile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).
			Ref("profile").
			Unique().
			// We add the "Required" method to the builder
			// to make this edge required on entity creation.
			// i.e. Profile cannot be created without its Account.
			Required(),
	}
}
