package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Sector holds the schema definition for the Sector entity.
type Sector struct {
	ent.Schema
}

// Fields of the Sector.
func (Sector) Fields() []ent.Field {
	return []ent.Field{
		field.String("coords").
			Unique(),
	}
}

func (Sector) Annotations() []schema.Annotation {
	return []schema.Annotation {
		entsql.Annotation{Table: "Sector"},
	}
}

// Edges of the Sector.
func (Sector) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("AdjacentTables", AdjacentTable.Type),
	}
}
