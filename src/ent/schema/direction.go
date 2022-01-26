package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Direction holds the schema definition for the Direction entity.
type Direction struct {
	ent.Schema
}

// Fields of the Direction.
func (Direction) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("institute_id"),
	}
}

func (Direction) Annotations() []schema.Annotation {
	return []schema.Annotation {
		entsql.Annotation{Table: "Direction"},
	}
}

// Edges of the Direction.
func (Direction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Institute", Institute.Type).
			Ref("Directions").
			Unique().
			Field("institute_id").
			Required(),
		edge.To("Profile", Profile.Type),
	}
}
