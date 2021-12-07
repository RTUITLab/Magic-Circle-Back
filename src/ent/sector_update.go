// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/0B1t322/Magic-Circle/ent/adjacenttable"
	"github.com/0B1t322/Magic-Circle/ent/predicate"
	"github.com/0B1t322/Magic-Circle/ent/sector"
)

// SectorUpdate is the builder for updating Sector entities.
type SectorUpdate struct {
	config
	hooks    []Hook
	mutation *SectorMutation
}

// Where appends a list predicates to the SectorUpdate builder.
func (su *SectorUpdate) Where(ps ...predicate.Sector) *SectorUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetCoords sets the "coords" field.
func (su *SectorUpdate) SetCoords(s string) *SectorUpdate {
	su.mutation.SetCoords(s)
	return su
}

// SetDescription sets the "description" field.
func (su *SectorUpdate) SetDescription(s string) *SectorUpdate {
	su.mutation.SetDescription(s)
	return su
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (su *SectorUpdate) SetNillableDescription(s *string) *SectorUpdate {
	if s != nil {
		su.SetDescription(*s)
	}
	return su
}

// ClearDescription clears the value of the "description" field.
func (su *SectorUpdate) ClearDescription() *SectorUpdate {
	su.mutation.ClearDescription()
	return su
}

// AddAdjacentTableIDs adds the "AdjacentTables" edge to the AdjacentTable entity by IDs.
func (su *SectorUpdate) AddAdjacentTableIDs(ids ...int) *SectorUpdate {
	su.mutation.AddAdjacentTableIDs(ids...)
	return su
}

// AddAdjacentTables adds the "AdjacentTables" edges to the AdjacentTable entity.
func (su *SectorUpdate) AddAdjacentTables(a ...*AdjacentTable) *SectorUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return su.AddAdjacentTableIDs(ids...)
}

// Mutation returns the SectorMutation object of the builder.
func (su *SectorUpdate) Mutation() *SectorMutation {
	return su.mutation
}

// ClearAdjacentTables clears all "AdjacentTables" edges to the AdjacentTable entity.
func (su *SectorUpdate) ClearAdjacentTables() *SectorUpdate {
	su.mutation.ClearAdjacentTables()
	return su
}

// RemoveAdjacentTableIDs removes the "AdjacentTables" edge to AdjacentTable entities by IDs.
func (su *SectorUpdate) RemoveAdjacentTableIDs(ids ...int) *SectorUpdate {
	su.mutation.RemoveAdjacentTableIDs(ids...)
	return su
}

// RemoveAdjacentTables removes "AdjacentTables" edges to AdjacentTable entities.
func (su *SectorUpdate) RemoveAdjacentTables(a ...*AdjacentTable) *SectorUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return su.RemoveAdjacentTableIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SectorUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SectorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SectorUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SectorUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SectorUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SectorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sector.Table,
			Columns: sector.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sector.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Coords(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sector.FieldCoords,
		})
	}
	if value, ok := su.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sector.FieldDescription,
		})
	}
	if su.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sector.FieldDescription,
		})
	}
	if su.mutation.AdjacentTablesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sector.AdjacentTablesTable,
			Columns: []string{sector.AdjacentTablesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: adjacenttable.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedAdjacentTablesIDs(); len(nodes) > 0 && !su.mutation.AdjacentTablesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sector.AdjacentTablesTable,
			Columns: []string{sector.AdjacentTablesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: adjacenttable.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.AdjacentTablesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sector.AdjacentTablesTable,
			Columns: []string{sector.AdjacentTablesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: adjacenttable.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sector.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// SectorUpdateOne is the builder for updating a single Sector entity.
type SectorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SectorMutation
}

// SetCoords sets the "coords" field.
func (suo *SectorUpdateOne) SetCoords(s string) *SectorUpdateOne {
	suo.mutation.SetCoords(s)
	return suo
}

// SetDescription sets the "description" field.
func (suo *SectorUpdateOne) SetDescription(s string) *SectorUpdateOne {
	suo.mutation.SetDescription(s)
	return suo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (suo *SectorUpdateOne) SetNillableDescription(s *string) *SectorUpdateOne {
	if s != nil {
		suo.SetDescription(*s)
	}
	return suo
}

// ClearDescription clears the value of the "description" field.
func (suo *SectorUpdateOne) ClearDescription() *SectorUpdateOne {
	suo.mutation.ClearDescription()
	return suo
}

// AddAdjacentTableIDs adds the "AdjacentTables" edge to the AdjacentTable entity by IDs.
func (suo *SectorUpdateOne) AddAdjacentTableIDs(ids ...int) *SectorUpdateOne {
	suo.mutation.AddAdjacentTableIDs(ids...)
	return suo
}

// AddAdjacentTables adds the "AdjacentTables" edges to the AdjacentTable entity.
func (suo *SectorUpdateOne) AddAdjacentTables(a ...*AdjacentTable) *SectorUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return suo.AddAdjacentTableIDs(ids...)
}

// Mutation returns the SectorMutation object of the builder.
func (suo *SectorUpdateOne) Mutation() *SectorMutation {
	return suo.mutation
}

// ClearAdjacentTables clears all "AdjacentTables" edges to the AdjacentTable entity.
func (suo *SectorUpdateOne) ClearAdjacentTables() *SectorUpdateOne {
	suo.mutation.ClearAdjacentTables()
	return suo
}

// RemoveAdjacentTableIDs removes the "AdjacentTables" edge to AdjacentTable entities by IDs.
func (suo *SectorUpdateOne) RemoveAdjacentTableIDs(ids ...int) *SectorUpdateOne {
	suo.mutation.RemoveAdjacentTableIDs(ids...)
	return suo
}

// RemoveAdjacentTables removes "AdjacentTables" edges to AdjacentTable entities.
func (suo *SectorUpdateOne) RemoveAdjacentTables(a ...*AdjacentTable) *SectorUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return suo.RemoveAdjacentTableIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SectorUpdateOne) Select(field string, fields ...string) *SectorUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Sector entity.
func (suo *SectorUpdateOne) Save(ctx context.Context) (*Sector, error) {
	var (
		err  error
		node *Sector
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SectorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SectorUpdateOne) SaveX(ctx context.Context) *Sector {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SectorUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SectorUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SectorUpdateOne) sqlSave(ctx context.Context) (_node *Sector, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sector.Table,
			Columns: sector.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sector.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Sector.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sector.FieldID)
		for _, f := range fields {
			if !sector.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sector.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Coords(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sector.FieldCoords,
		})
	}
	if value, ok := suo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sector.FieldDescription,
		})
	}
	if suo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sector.FieldDescription,
		})
	}
	if suo.mutation.AdjacentTablesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sector.AdjacentTablesTable,
			Columns: []string{sector.AdjacentTablesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: adjacenttable.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedAdjacentTablesIDs(); len(nodes) > 0 && !suo.mutation.AdjacentTablesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sector.AdjacentTablesTable,
			Columns: []string{sector.AdjacentTablesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: adjacenttable.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.AdjacentTablesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sector.AdjacentTablesTable,
			Columns: []string{sector.AdjacentTablesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: adjacenttable.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Sector{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sector.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
