package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)


// AdjacentTable holds the schema definition for the AdjacentTable entity.
type AdjacentTable struct {
	ent.Schema
}

// Fields of the AdjacentTable.
func (AdjacentTable) Fields() []ent.Field {
	return []ent.Field{
		field.Int("sector_id"),
		field.Int("profile_id"),
		field.Text("additionalDescription").Optional(),
	}
}

func (AdjacentTable) Annotations() []schema.Annotation {
	return []schema.Annotation {
		entsql.Annotation{Table: "AdjacentTable"},
	}
}

// Edges of the AdjacentTable.
func (AdjacentTable) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Profile", Profile.Type).
			Ref("AdjacentTables").
			Unique().
			Field("profile_id").
			Required(),
		edge.From("Sector", Sector.Type).
			Ref("AdjacentTables").
			Unique().
			Field("sector_id").
			Required(),
	}
}
