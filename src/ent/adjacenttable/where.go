// Code generated by entc, DO NOT EDIT.

package adjacenttable

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/0B1t322/Magic-Circle/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.AdjacentTable {
	return predicate.AdjacentTable(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.AdjacentTable {
	return predicate.AdjacentTable(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.AdjacentTable {
	return predicate.AdjacentTable(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.AdjacentTable {
	return predicate.AdjacentTable(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.AdjacentTable {
	return predicate.AdjacentTable(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.AdjacentTable {
	return predicate.AdjacentTable(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.AdjacentTable {
	return predicate.AdjacentTable(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.AdjacentTable {
	return predicate.AdjacentTable(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.AdjacentTable {
	return predicate.AdjacentTable(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// SectorID applies equality check predicate on the "sector_id" field. It's identical to SectorIDEQ.
func SectorID(v int) predicate.AdjacentTable {
	return predicate.AdjacentTable(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSectorID), v))
	})
}

// VariantID applies equality check predicate on the "variant_id" field. It's identical to VariantIDEQ.
func VariantID(v int) predicate.AdjacentTable {
	return predicate.AdjacentTable(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVariantID), v))
	})
}

// SectorIDEQ applies the EQ predicate on the "sector_id" field.
func SectorIDEQ(v int) predicate.AdjacentTable {
	return predicate.AdjacentTable(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSectorID), v))
	})
}

// SectorIDNEQ applies the NEQ predicate on the "sector_id" field.
func SectorIDNEQ(v int) predicate.AdjacentTable {
	return predicate.AdjacentTable(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSectorID), v))
	})
}

// SectorIDIn applies the In predicate on the "sector_id" field.
func SectorIDIn(vs ...int) predicate.AdjacentTable {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdjacentTable(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSectorID), v...))
	})
}

// SectorIDNotIn applies the NotIn predicate on the "sector_id" field.
func SectorIDNotIn(vs ...int) predicate.AdjacentTable {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdjacentTable(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSectorID), v...))
	})
}

// VariantIDEQ applies the EQ predicate on the "variant_id" field.
func VariantIDEQ(v int) predicate.AdjacentTable {
	return predicate.AdjacentTable(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVariantID), v))
	})
}

// VariantIDNEQ applies the NEQ predicate on the "variant_id" field.
func VariantIDNEQ(v int) predicate.AdjacentTable {
	return predicate.AdjacentTable(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVariantID), v))
	})
}

// VariantIDIn applies the In predicate on the "variant_id" field.
func VariantIDIn(vs ...int) predicate.AdjacentTable {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdjacentTable(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldVariantID), v...))
	})
}

// VariantIDNotIn applies the NotIn predicate on the "variant_id" field.
func VariantIDNotIn(vs ...int) predicate.AdjacentTable {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdjacentTable(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldVariantID), v...))
	})
}

// HasVariant applies the HasEdge predicate on the "Variant" edge.
func HasVariant() predicate.AdjacentTable {
	return predicate.AdjacentTable(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VariantTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, VariantTable, VariantColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVariantWith applies the HasEdge predicate on the "Variant" edge with a given conditions (other predicates).
func HasVariantWith(preds ...predicate.Variant) predicate.AdjacentTable {
	return predicate.AdjacentTable(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VariantInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, VariantTable, VariantColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSector applies the HasEdge predicate on the "Sector" edge.
func HasSector() predicate.AdjacentTable {
	return predicate.AdjacentTable(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SectorTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SectorTable, SectorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSectorWith applies the HasEdge predicate on the "Sector" edge with a given conditions (other predicates).
func HasSectorWith(preds ...predicate.Sector) predicate.AdjacentTable {
	return predicate.AdjacentTable(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SectorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SectorTable, SectorColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AdjacentTable) predicate.AdjacentTable {
	return predicate.AdjacentTable(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AdjacentTable) predicate.AdjacentTable {
	return predicate.AdjacentTable(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AdjacentTable) predicate.AdjacentTable {
	return predicate.AdjacentTable(func(s *sql.Selector) {
		p(s.Not())
	})
}
