// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/0B1t322/Magic-Circle/ent/adjacenttable"
	"github.com/0B1t322/Magic-Circle/ent/predicate"
	"github.com/0B1t322/Magic-Circle/ent/profile"
	"github.com/0B1t322/Magic-Circle/ent/sector"
)

// AdjacentTableUpdate is the builder for updating AdjacentTable entities.
type AdjacentTableUpdate struct {
	config
	hooks    []Hook
	mutation *AdjacentTableMutation
}

// Where appends a list predicates to the AdjacentTableUpdate builder.
func (atu *AdjacentTableUpdate) Where(ps ...predicate.AdjacentTable) *AdjacentTableUpdate {
	atu.mutation.Where(ps...)
	return atu
}

// SetSectorID sets the "sector_id" field.
func (atu *AdjacentTableUpdate) SetSectorID(i int) *AdjacentTableUpdate {
	atu.mutation.SetSectorID(i)
	return atu
}

// SetProfileID sets the "profile_id" field.
func (atu *AdjacentTableUpdate) SetProfileID(i int) *AdjacentTableUpdate {
	atu.mutation.SetProfileID(i)
	return atu
}

// SetProfile sets the "Profile" edge to the Profile entity.
func (atu *AdjacentTableUpdate) SetProfile(p *Profile) *AdjacentTableUpdate {
	return atu.SetProfileID(p.ID)
}

// SetSector sets the "Sector" edge to the Sector entity.
func (atu *AdjacentTableUpdate) SetSector(s *Sector) *AdjacentTableUpdate {
	return atu.SetSectorID(s.ID)
}

// Mutation returns the AdjacentTableMutation object of the builder.
func (atu *AdjacentTableUpdate) Mutation() *AdjacentTableMutation {
	return atu.mutation
}

// ClearProfile clears the "Profile" edge to the Profile entity.
func (atu *AdjacentTableUpdate) ClearProfile() *AdjacentTableUpdate {
	atu.mutation.ClearProfile()
	return atu
}

// ClearSector clears the "Sector" edge to the Sector entity.
func (atu *AdjacentTableUpdate) ClearSector() *AdjacentTableUpdate {
	atu.mutation.ClearSector()
	return atu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (atu *AdjacentTableUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(atu.hooks) == 0 {
		if err = atu.check(); err != nil {
			return 0, err
		}
		affected, err = atu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdjacentTableMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = atu.check(); err != nil {
				return 0, err
			}
			atu.mutation = mutation
			affected, err = atu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(atu.hooks) - 1; i >= 0; i-- {
			if atu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = atu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (atu *AdjacentTableUpdate) SaveX(ctx context.Context) int {
	affected, err := atu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (atu *AdjacentTableUpdate) Exec(ctx context.Context) error {
	_, err := atu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atu *AdjacentTableUpdate) ExecX(ctx context.Context) {
	if err := atu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atu *AdjacentTableUpdate) check() error {
	if _, ok := atu.mutation.ProfileID(); atu.mutation.ProfileCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"Profile\"")
	}
	if _, ok := atu.mutation.SectorID(); atu.mutation.SectorCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"Sector\"")
	}
	return nil
}

func (atu *AdjacentTableUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   adjacenttable.Table,
			Columns: adjacenttable.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: adjacenttable.FieldID,
			},
		},
	}
	if ps := atu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if atu.mutation.ProfileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   adjacenttable.ProfileTable,
			Columns: []string{adjacenttable.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: profile.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.ProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   adjacenttable.ProfileTable,
			Columns: []string{adjacenttable.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: profile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if atu.mutation.SectorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   adjacenttable.SectorTable,
			Columns: []string{adjacenttable.SectorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: sector.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.SectorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   adjacenttable.SectorTable,
			Columns: []string{adjacenttable.SectorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: sector.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, atu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{adjacenttable.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AdjacentTableUpdateOne is the builder for updating a single AdjacentTable entity.
type AdjacentTableUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AdjacentTableMutation
}

// SetSectorID sets the "sector_id" field.
func (atuo *AdjacentTableUpdateOne) SetSectorID(i int) *AdjacentTableUpdateOne {
	atuo.mutation.SetSectorID(i)
	return atuo
}

// SetProfileID sets the "profile_id" field.
func (atuo *AdjacentTableUpdateOne) SetProfileID(i int) *AdjacentTableUpdateOne {
	atuo.mutation.SetProfileID(i)
	return atuo
}

// SetProfile sets the "Profile" edge to the Profile entity.
func (atuo *AdjacentTableUpdateOne) SetProfile(p *Profile) *AdjacentTableUpdateOne {
	return atuo.SetProfileID(p.ID)
}

// SetSector sets the "Sector" edge to the Sector entity.
func (atuo *AdjacentTableUpdateOne) SetSector(s *Sector) *AdjacentTableUpdateOne {
	return atuo.SetSectorID(s.ID)
}

// Mutation returns the AdjacentTableMutation object of the builder.
func (atuo *AdjacentTableUpdateOne) Mutation() *AdjacentTableMutation {
	return atuo.mutation
}

// ClearProfile clears the "Profile" edge to the Profile entity.
func (atuo *AdjacentTableUpdateOne) ClearProfile() *AdjacentTableUpdateOne {
	atuo.mutation.ClearProfile()
	return atuo
}

// ClearSector clears the "Sector" edge to the Sector entity.
func (atuo *AdjacentTableUpdateOne) ClearSector() *AdjacentTableUpdateOne {
	atuo.mutation.ClearSector()
	return atuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (atuo *AdjacentTableUpdateOne) Select(field string, fields ...string) *AdjacentTableUpdateOne {
	atuo.fields = append([]string{field}, fields...)
	return atuo
}

// Save executes the query and returns the updated AdjacentTable entity.
func (atuo *AdjacentTableUpdateOne) Save(ctx context.Context) (*AdjacentTable, error) {
	var (
		err  error
		node *AdjacentTable
	)
	if len(atuo.hooks) == 0 {
		if err = atuo.check(); err != nil {
			return nil, err
		}
		node, err = atuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdjacentTableMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = atuo.check(); err != nil {
				return nil, err
			}
			atuo.mutation = mutation
			node, err = atuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(atuo.hooks) - 1; i >= 0; i-- {
			if atuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = atuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (atuo *AdjacentTableUpdateOne) SaveX(ctx context.Context) *AdjacentTable {
	node, err := atuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (atuo *AdjacentTableUpdateOne) Exec(ctx context.Context) error {
	_, err := atuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atuo *AdjacentTableUpdateOne) ExecX(ctx context.Context) {
	if err := atuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atuo *AdjacentTableUpdateOne) check() error {
	if _, ok := atuo.mutation.ProfileID(); atuo.mutation.ProfileCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"Profile\"")
	}
	if _, ok := atuo.mutation.SectorID(); atuo.mutation.SectorCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"Sector\"")
	}
	return nil
}

func (atuo *AdjacentTableUpdateOne) sqlSave(ctx context.Context) (_node *AdjacentTable, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   adjacenttable.Table,
			Columns: adjacenttable.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: adjacenttable.FieldID,
			},
		},
	}
	id, ok := atuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing AdjacentTable.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := atuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, adjacenttable.FieldID)
		for _, f := range fields {
			if !adjacenttable.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != adjacenttable.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := atuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if atuo.mutation.ProfileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   adjacenttable.ProfileTable,
			Columns: []string{adjacenttable.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: profile.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.ProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   adjacenttable.ProfileTable,
			Columns: []string{adjacenttable.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: profile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if atuo.mutation.SectorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   adjacenttable.SectorTable,
			Columns: []string{adjacenttable.SectorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: sector.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.SectorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   adjacenttable.SectorTable,
			Columns: []string{adjacenttable.SectorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: sector.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AdjacentTable{config: atuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, atuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{adjacenttable.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
