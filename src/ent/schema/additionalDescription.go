package schema

// import (
// 	"entgo.io/ent"
// 	"entgo.io/ent/dialect/entsql"
// 	"entgo.io/ent/schema"
// 	"entgo.io/ent/schema/edge"
// 	"entgo.io/ent/schema/field"
// )

// type AdditonalDescription struct {
// 	ent.Schema
// }

// func(AdditonalDescription) Fields() []ent.Field {
// 	return []ent.Field{
// 		field.Text("additionalDescription"),
// 		field.Int("sector_id"),
// 		field.Int("profile_id"),
// 	}
// }

// func (AdditonalDescription) Edges() []ent.Edge {
// 	return []ent.Edge{
// 		edge.From("Profile", Profile.Type).
// 			Ref("AdditonalDescriptions").
// 			Unique().
// 			Field("profile_id").
// 			Required(),
// 		edge.From("Sector", Sector.Type).
// 			Ref("AdditonalDescriptions").
// 			Unique().
// 			Field("sector_id").
// 			Required(),
// 	}
// }

// func (AdditonalDescription) Annotations() []schema.Annotation {
// 	return []schema.Annotation {
// 		entsql.Annotation{Table: "AdditonalDescription"},
// 	}
// }