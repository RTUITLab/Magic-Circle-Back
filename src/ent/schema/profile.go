package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
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
		field.String("name"),
		field.Int("direction_id"),
	}
}

// Edges of the Profile.
func (Profile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Direction", Direction.Type).
			Ref("Profile").
			Unique().
			Field("direction_id").
			Required(),
		edge.To("AdjacentTables", AdjacentTable.Type),
		// edge.To("AdditonalDescriptions", AdditonalDescription.Type),
	}
}

func (Profile) Annotations() []schema.Annotation {
	return []schema.Annotation {
		entsql.Annotation{Table: "Profile"},
	}
}