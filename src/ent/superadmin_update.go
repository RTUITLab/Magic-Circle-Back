// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/0B1t322/Magic-Circle/ent/predicate"
	"github.com/0B1t322/Magic-Circle/ent/superadmin"
)

// SuperAdminUpdate is the builder for updating SuperAdmin entities.
type SuperAdminUpdate struct {
	config
	hooks    []Hook
	mutation *SuperAdminMutation
}

// Where appends a list predicates to the SuperAdminUpdate builder.
func (sau *SuperAdminUpdate) Where(ps ...predicate.SuperAdmin) *SuperAdminUpdate {
	sau.mutation.Where(ps...)
	return sau
}

// SetLogin sets the "login" field.
func (sau *SuperAdminUpdate) SetLogin(s string) *SuperAdminUpdate {
	sau.mutation.SetLogin(s)
	return sau
}

// SetPassword sets the "password" field.
func (sau *SuperAdminUpdate) SetPassword(s string) *SuperAdminUpdate {
	sau.mutation.SetPassword(s)
	return sau
}

// SetEmail sets the "email" field.
func (sau *SuperAdminUpdate) SetEmail(s string) *SuperAdminUpdate {
	sau.mutation.SetEmail(s)
	return sau
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (sau *SuperAdminUpdate) SetNillableEmail(s *string) *SuperAdminUpdate {
	if s != nil {
		sau.SetEmail(*s)
	}
	return sau
}

// ClearEmail clears the value of the "email" field.
func (sau *SuperAdminUpdate) ClearEmail() *SuperAdminUpdate {
	sau.mutation.ClearEmail()
	return sau
}

// Mutation returns the SuperAdminMutation object of the builder.
func (sau *SuperAdminUpdate) Mutation() *SuperAdminMutation {
	return sau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sau *SuperAdminUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sau.hooks) == 0 {
		affected, err = sau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SuperAdminMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sau.mutation = mutation
			affected, err = sau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sau.hooks) - 1; i >= 0; i-- {
			if sau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sau *SuperAdminUpdate) SaveX(ctx context.Context) int {
	affected, err := sau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sau *SuperAdminUpdate) Exec(ctx context.Context) error {
	_, err := sau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sau *SuperAdminUpdate) ExecX(ctx context.Context) {
	if err := sau.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sau *SuperAdminUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   superadmin.Table,
			Columns: superadmin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: superadmin.FieldID,
			},
		},
	}
	if ps := sau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sau.mutation.Login(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: superadmin.FieldLogin,
		})
	}
	if value, ok := sau.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: superadmin.FieldPassword,
		})
	}
	if value, ok := sau.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: superadmin.FieldEmail,
		})
	}
	if sau.mutation.EmailCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: superadmin.FieldEmail,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{superadmin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// SuperAdminUpdateOne is the builder for updating a single SuperAdmin entity.
type SuperAdminUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SuperAdminMutation
}

// SetLogin sets the "login" field.
func (sauo *SuperAdminUpdateOne) SetLogin(s string) *SuperAdminUpdateOne {
	sauo.mutation.SetLogin(s)
	return sauo
}

// SetPassword sets the "password" field.
func (sauo *SuperAdminUpdateOne) SetPassword(s string) *SuperAdminUpdateOne {
	sauo.mutation.SetPassword(s)
	return sauo
}

// SetEmail sets the "email" field.
func (sauo *SuperAdminUpdateOne) SetEmail(s string) *SuperAdminUpdateOne {
	sauo.mutation.SetEmail(s)
	return sauo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (sauo *SuperAdminUpdateOne) SetNillableEmail(s *string) *SuperAdminUpdateOne {
	if s != nil {
		sauo.SetEmail(*s)
	}
	return sauo
}

// ClearEmail clears the value of the "email" field.
func (sauo *SuperAdminUpdateOne) ClearEmail() *SuperAdminUpdateOne {
	sauo.mutation.ClearEmail()
	return sauo
}

// Mutation returns the SuperAdminMutation object of the builder.
func (sauo *SuperAdminUpdateOne) Mutation() *SuperAdminMutation {
	return sauo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sauo *SuperAdminUpdateOne) Select(field string, fields ...string) *SuperAdminUpdateOne {
	sauo.fields = append([]string{field}, fields...)
	return sauo
}

// Save executes the query and returns the updated SuperAdmin entity.
func (sauo *SuperAdminUpdateOne) Save(ctx context.Context) (*SuperAdmin, error) {
	var (
		err  error
		node *SuperAdmin
	)
	if len(sauo.hooks) == 0 {
		node, err = sauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SuperAdminMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sauo.mutation = mutation
			node, err = sauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sauo.hooks) - 1; i >= 0; i-- {
			if sauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sauo *SuperAdminUpdateOne) SaveX(ctx context.Context) *SuperAdmin {
	node, err := sauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sauo *SuperAdminUpdateOne) Exec(ctx context.Context) error {
	_, err := sauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sauo *SuperAdminUpdateOne) ExecX(ctx context.Context) {
	if err := sauo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sauo *SuperAdminUpdateOne) sqlSave(ctx context.Context) (_node *SuperAdmin, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   superadmin.Table,
			Columns: superadmin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: superadmin.FieldID,
			},
		},
	}
	id, ok := sauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing SuperAdmin.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := sauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, superadmin.FieldID)
		for _, f := range fields {
			if !superadmin.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != superadmin.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sauo.mutation.Login(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: superadmin.FieldLogin,
		})
	}
	if value, ok := sauo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: superadmin.FieldPassword,
		})
	}
	if value, ok := sauo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: superadmin.FieldEmail,
		})
	}
	if sauo.mutation.EmailCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: superadmin.FieldEmail,
		})
	}
	_node = &SuperAdmin{config: sauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{superadmin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
