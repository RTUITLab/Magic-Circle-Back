package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Admin holds the schema definition for the Admin entity.
type Admin struct {
	ent.Schema
}

// Fields of the Admin.
func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.String("login").Unique(),
		field.String("password"),
		field.Int("institute_id"),
	}
}

// Edges of the Admin.
func (Admin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Institute", Institute.Type).
			Ref("Admins").
			Unique().
			Field("institute_id").
			Required(),
	}
}

func (Admin) Annotations() []schema.Annotation {
	return []schema.Annotation {
		entsql.Annotation{Table: "Admin"},
	}
}