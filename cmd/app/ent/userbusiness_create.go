// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/business"
	"placio-app/ent/user"
	"placio-app/ent/userbusiness"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserBusinessCreate is the builder for creating a UserBusiness entity.
type UserBusinessCreate struct {
	config
	mutation *UserBusinessMutation
	hooks    []Hook
}

// SetRole sets the "role" field.
func (ubc *UserBusinessCreate) SetRole(s string) *UserBusinessCreate {
	ubc.mutation.SetRole(s)
	return ubc
}

// SetID sets the "id" field.
func (ubc *UserBusinessCreate) SetID(s string) *UserBusinessCreate {
	ubc.mutation.SetID(s)
	return ubc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ubc *UserBusinessCreate) SetUserID(id string) *UserBusinessCreate {
	ubc.mutation.SetUserID(id)
	return ubc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (ubc *UserBusinessCreate) SetNillableUserID(id *string) *UserBusinessCreate {
	if id != nil {
		ubc = ubc.SetUserID(*id)
	}
	return ubc
}

// SetUser sets the "user" edge to the User entity.
func (ubc *UserBusinessCreate) SetUser(u *User) *UserBusinessCreate {
	return ubc.SetUserID(u.ID)
}

// SetBusinessID sets the "business" edge to the Business entity by ID.
func (ubc *UserBusinessCreate) SetBusinessID(id string) *UserBusinessCreate {
	ubc.mutation.SetBusinessID(id)
	return ubc
}

// SetNillableBusinessID sets the "business" edge to the Business entity by ID if the given value is not nil.
func (ubc *UserBusinessCreate) SetNillableBusinessID(id *string) *UserBusinessCreate {
	if id != nil {
		ubc = ubc.SetBusinessID(*id)
	}
	return ubc
}

// SetBusiness sets the "business" edge to the Business entity.
func (ubc *UserBusinessCreate) SetBusiness(b *Business) *UserBusinessCreate {
	return ubc.SetBusinessID(b.ID)
}

// Mutation returns the UserBusinessMutation object of the builder.
func (ubc *UserBusinessCreate) Mutation() *UserBusinessMutation {
	return ubc.mutation
}

// Save creates the UserBusiness in the database.
func (ubc *UserBusinessCreate) Save(ctx context.Context) (*UserBusiness, error) {
	return withHooks(ctx, ubc.sqlSave, ubc.mutation, ubc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ubc *UserBusinessCreate) SaveX(ctx context.Context) *UserBusiness {
	v, err := ubc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ubc *UserBusinessCreate) Exec(ctx context.Context) error {
	_, err := ubc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ubc *UserBusinessCreate) ExecX(ctx context.Context) {
	if err := ubc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ubc *UserBusinessCreate) check() error {
	if _, ok := ubc.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`ent: missing required field "UserBusiness.role"`)}
	}
	if v, ok := ubc.mutation.ID(); ok {
		if err := userbusiness.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "UserBusiness.id": %w`, err)}
		}
	}
	return nil
}

func (ubc *UserBusinessCreate) sqlSave(ctx context.Context) (*UserBusiness, error) {
	if err := ubc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ubc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ubc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected UserBusiness.ID type: %T", _spec.ID.Value)
		}
	}
	ubc.mutation.id = &_node.ID
	ubc.mutation.done = true
	return _node, nil
}

func (ubc *UserBusinessCreate) createSpec() (*UserBusiness, *sqlgraph.CreateSpec) {
	var (
		_node = &UserBusiness{config: ubc.config}
		_spec = sqlgraph.NewCreateSpec(userbusiness.Table, sqlgraph.NewFieldSpec(userbusiness.FieldID, field.TypeString))
	)
	if id, ok := ubc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ubc.mutation.Role(); ok {
		_spec.SetField(userbusiness.FieldRole, field.TypeString, value)
		_node.Role = value
	}
	if nodes := ubc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userbusiness.UserTable,
			Columns: []string{userbusiness.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_user_businesses = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ubc.mutation.BusinessIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userbusiness.BusinessTable,
			Columns: []string{userbusiness.BusinessColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.business_user_businesses = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserBusinessCreateBulk is the builder for creating many UserBusiness entities in bulk.
type UserBusinessCreateBulk struct {
	config
	builders []*UserBusinessCreate
}

// Save creates the UserBusiness entities in the database.
func (ubcb *UserBusinessCreateBulk) Save(ctx context.Context) ([]*UserBusiness, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ubcb.builders))
	nodes := make([]*UserBusiness, len(ubcb.builders))
	mutators := make([]Mutator, len(ubcb.builders))
	for i := range ubcb.builders {
		func(i int, root context.Context) {
			builder := ubcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserBusinessMutation)
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
					_, err = mutators[i+1].Mutate(root, ubcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ubcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ubcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ubcb *UserBusinessCreateBulk) SaveX(ctx context.Context) []*UserBusiness {
	v, err := ubcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ubcb *UserBusinessCreateBulk) Exec(ctx context.Context) error {
	_, err := ubcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ubcb *UserBusinessCreateBulk) ExecX(ctx context.Context) {
	if err := ubcb.Exec(ctx); err != nil {
		panic(err)
	}
}
