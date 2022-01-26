package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Institute holds the schema definition for the Institute entity.
type Institute struct {
	ent.Schema
}

// Fields of the Institute.
func (Institute) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
	}
}

func (Institute) Annotations() []schema.Annotation {
	return []schema.Annotation {
		entsql.Annotation{Table: "Institute"},
	}
}

// Edges of the Institute.
func (Institute) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Directions", Direction.Type),
	}
}
