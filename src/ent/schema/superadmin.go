package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// SuperAdmin holds the schema definition for the SuperAdmin entity.
type SuperAdmin struct {
	ent.Schema
}

// Fields of the SuperAdmin.
func (SuperAdmin) Fields() []ent.Field {
	return []ent.Field{
		field.String("login").Unique(),
		field.String("password"),
		field.String("email").Optional(),
	}
}

// Edges of the SuperAdmin.
func (SuperAdmin) Edges() []ent.Edge {
	return nil
}

func (SuperAdmin) Annotations() []schema.Annotation {
	return []schema.Annotation {
		entsql.Annotation{Table: "SuperAdmin"},
	}
}