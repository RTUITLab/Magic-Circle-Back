package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Variant holds the schema definition for the Variant entity.
type Variant struct {
	ent.Schema
}

// Fields of the Variant.
func (Variant) Fields() []ent.Field {
	return []ent.Field{
		field.Int("direction_id"),
		field.Int("profile_id"),
		field.Int("insitute_id"),
	}
}

func (Variant) Annotations() []schema.Annotation {
	return []schema.Annotation {
		entsql.Annotation{Table: "Variant"},
	}
}

// Edges of the Variant.
func (Variant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Insitute", Institute.Type).
			Ref("Variants").
			Unique().
			Field("insitute_id").
			Required(),
		edge.From("Direction", Direction.Type).
			Ref("Variants").
			Unique().
			Field("direction_id").
			Required(),
		edge.From("Profile", Profile.Type).
			Ref("Variants").
			Unique().
			Field("profile_id").
			Required(),
		edge.To("AdjacentTables", AdjacentTable.Type),
	}
}
