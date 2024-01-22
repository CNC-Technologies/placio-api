// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/place"
	"placio-app/ent/trainer"
	"placio-app/ent/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TrainerCreate is the builder for creating a Trainer entity.
type TrainerCreate struct {
	config
	mutation *TrainerMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (tc *TrainerCreate) SetName(s string) *TrainerCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetEmail sets the "email" field.
func (tc *TrainerCreate) SetEmail(s string) *TrainerCreate {
	tc.mutation.SetEmail(s)
	return tc
}

// SetPhone sets the "phone" field.
func (tc *TrainerCreate) SetPhone(s string) *TrainerCreate {
	tc.mutation.SetPhone(s)
	return tc
}

// SetID sets the "id" field.
func (tc *TrainerCreate) SetID(s string) *TrainerCreate {
	tc.mutation.SetID(s)
	return tc
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (tc *TrainerCreate) AddUserIDs(ids ...string) *TrainerCreate {
	tc.mutation.AddUserIDs(ids...)
	return tc
}

// AddUser adds the "user" edges to the User entity.
func (tc *TrainerCreate) AddUser(u ...*User) *TrainerCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tc.AddUserIDs(ids...)
}

// AddPlaceIDs adds the "place" edge to the Place entity by IDs.
func (tc *TrainerCreate) AddPlaceIDs(ids ...string) *TrainerCreate {
	tc.mutation.AddPlaceIDs(ids...)
	return tc
}

// AddPlace adds the "place" edges to the Place entity.
func (tc *TrainerCreate) AddPlace(p ...*Place) *TrainerCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tc.AddPlaceIDs(ids...)
}

// Mutation returns the TrainerMutation object of the builder.
func (tc *TrainerCreate) Mutation() *TrainerMutation {
	return tc.mutation
}

// Save creates the Trainer in the database.
func (tc *TrainerCreate) Save(ctx context.Context) (*Trainer, error) {
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TrainerCreate) SaveX(ctx context.Context) *Trainer {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TrainerCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TrainerCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TrainerCreate) check() error {
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Trainer.name"`)}
	}
	if _, ok := tc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Trainer.email"`)}
	}
	if _, ok := tc.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New(`ent: missing required field "Trainer.phone"`)}
	}
	if v, ok := tc.mutation.ID(); ok {
		if err := trainer.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Trainer.id": %w`, err)}
		}
	}
	return nil
}

func (tc *TrainerCreate) sqlSave(ctx context.Context) (*Trainer, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Trainer.ID type: %T", _spec.ID.Value)
		}
	}
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TrainerCreate) createSpec() (*Trainer, *sqlgraph.CreateSpec) {
	var (
		_node = &Trainer{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(trainer.Table, sqlgraph.NewFieldSpec(trainer.FieldID, field.TypeString))
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.SetField(trainer.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := tc.mutation.Email(); ok {
		_spec.SetField(trainer.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := tc.mutation.Phone(); ok {
		_spec.SetField(trainer.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if nodes := tc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   trainer.UserTable,
			Columns: trainer.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.PlaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   trainer.PlaceTable,
			Columns: trainer.PlacePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(place.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TrainerCreateBulk is the builder for creating many Trainer entities in bulk.
type TrainerCreateBulk struct {
	config
	err      error
	builders []*TrainerCreate
}

// Save creates the Trainer entities in the database.
func (tcb *TrainerCreateBulk) Save(ctx context.Context) ([]*Trainer, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Trainer, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TrainerMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TrainerCreateBulk) SaveX(ctx context.Context) []*Trainer {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TrainerCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TrainerCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}