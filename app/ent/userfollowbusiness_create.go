// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/business"
	"placio-app/ent/user"
	"placio-app/ent/userfollowbusiness"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserFollowBusinessCreate is the builder for creating a UserFollowBusiness entity.
type UserFollowBusinessCreate struct {
	config
	mutation *UserFollowBusinessMutation
	hooks    []Hook
}

// SetCreatedAt sets the "CreatedAt" field.
func (ufbc *UserFollowBusinessCreate) SetCreatedAt(t time.Time) *UserFollowBusinessCreate {
	ufbc.mutation.SetCreatedAt(t)
	return ufbc
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (ufbc *UserFollowBusinessCreate) SetNillableCreatedAt(t *time.Time) *UserFollowBusinessCreate {
	if t != nil {
		ufbc.SetCreatedAt(*t)
	}
	return ufbc
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (ufbc *UserFollowBusinessCreate) SetUpdatedAt(t time.Time) *UserFollowBusinessCreate {
	ufbc.mutation.SetUpdatedAt(t)
	return ufbc
}

// SetID sets the "id" field.
func (ufbc *UserFollowBusinessCreate) SetID(s string) *UserFollowBusinessCreate {
	ufbc.mutation.SetID(s)
	return ufbc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ufbc *UserFollowBusinessCreate) SetUserID(id string) *UserFollowBusinessCreate {
	ufbc.mutation.SetUserID(id)
	return ufbc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (ufbc *UserFollowBusinessCreate) SetNillableUserID(id *string) *UserFollowBusinessCreate {
	if id != nil {
		ufbc = ufbc.SetUserID(*id)
	}
	return ufbc
}

// SetUser sets the "user" edge to the User entity.
func (ufbc *UserFollowBusinessCreate) SetUser(u *User) *UserFollowBusinessCreate {
	return ufbc.SetUserID(u.ID)
}

// SetBusinessID sets the "business" edge to the Business entity by ID.
func (ufbc *UserFollowBusinessCreate) SetBusinessID(id string) *UserFollowBusinessCreate {
	ufbc.mutation.SetBusinessID(id)
	return ufbc
}

// SetNillableBusinessID sets the "business" edge to the Business entity by ID if the given value is not nil.
func (ufbc *UserFollowBusinessCreate) SetNillableBusinessID(id *string) *UserFollowBusinessCreate {
	if id != nil {
		ufbc = ufbc.SetBusinessID(*id)
	}
	return ufbc
}

// SetBusiness sets the "business" edge to the Business entity.
func (ufbc *UserFollowBusinessCreate) SetBusiness(b *Business) *UserFollowBusinessCreate {
	return ufbc.SetBusinessID(b.ID)
}

// Mutation returns the UserFollowBusinessMutation object of the builder.
func (ufbc *UserFollowBusinessCreate) Mutation() *UserFollowBusinessMutation {
	return ufbc.mutation
}

// Save creates the UserFollowBusiness in the database.
func (ufbc *UserFollowBusinessCreate) Save(ctx context.Context) (*UserFollowBusiness, error) {
	ufbc.defaults()
	return withHooks(ctx, ufbc.sqlSave, ufbc.mutation, ufbc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ufbc *UserFollowBusinessCreate) SaveX(ctx context.Context) *UserFollowBusiness {
	v, err := ufbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ufbc *UserFollowBusinessCreate) Exec(ctx context.Context) error {
	_, err := ufbc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ufbc *UserFollowBusinessCreate) ExecX(ctx context.Context) {
	if err := ufbc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ufbc *UserFollowBusinessCreate) defaults() {
	if _, ok := ufbc.mutation.CreatedAt(); !ok {
		v := userfollowbusiness.DefaultCreatedAt()
		ufbc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ufbc *UserFollowBusinessCreate) check() error {
	if _, ok := ufbc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "CreatedAt", err: errors.New(`ent: missing required field "UserFollowBusiness.CreatedAt"`)}
	}
	if _, ok := ufbc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "UpdatedAt", err: errors.New(`ent: missing required field "UserFollowBusiness.UpdatedAt"`)}
	}
	return nil
}

func (ufbc *UserFollowBusinessCreate) sqlSave(ctx context.Context) (*UserFollowBusiness, error) {
	if err := ufbc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ufbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ufbc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected UserFollowBusiness.ID type: %T", _spec.ID.Value)
		}
	}
	ufbc.mutation.id = &_node.ID
	ufbc.mutation.done = true
	return _node, nil
}

func (ufbc *UserFollowBusinessCreate) createSpec() (*UserFollowBusiness, *sqlgraph.CreateSpec) {
	var (
		_node = &UserFollowBusiness{config: ufbc.config}
		_spec = sqlgraph.NewCreateSpec(userfollowbusiness.Table, sqlgraph.NewFieldSpec(userfollowbusiness.FieldID, field.TypeString))
	)
	if id, ok := ufbc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ufbc.mutation.CreatedAt(); ok {
		_spec.SetField(userfollowbusiness.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ufbc.mutation.UpdatedAt(); ok {
		_spec.SetField(userfollowbusiness.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := ufbc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userfollowbusiness.UserTable,
			Columns: []string{userfollowbusiness.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_followed_businesses = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ufbc.mutation.BusinessIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userfollowbusiness.BusinessTable,
			Columns: []string{userfollowbusiness.BusinessColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.business_follower_users = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserFollowBusinessCreateBulk is the builder for creating many UserFollowBusiness entities in bulk.
type UserFollowBusinessCreateBulk struct {
	config
	err      error
	builders []*UserFollowBusinessCreate
}

// Save creates the UserFollowBusiness entities in the database.
func (ufbcb *UserFollowBusinessCreateBulk) Save(ctx context.Context) ([]*UserFollowBusiness, error) {
	if ufbcb.err != nil {
		return nil, ufbcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ufbcb.builders))
	nodes := make([]*UserFollowBusiness, len(ufbcb.builders))
	mutators := make([]Mutator, len(ufbcb.builders))
	for i := range ufbcb.builders {
		func(i int, root context.Context) {
			builder := ufbcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserFollowBusinessMutation)
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
					_, err = mutators[i+1].Mutate(root, ufbcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ufbcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ufbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ufbcb *UserFollowBusinessCreateBulk) SaveX(ctx context.Context) []*UserFollowBusiness {
	v, err := ufbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ufbcb *UserFollowBusinessCreateBulk) Exec(ctx context.Context) error {
	_, err := ufbcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ufbcb *UserFollowBusinessCreateBulk) ExecX(ctx context.Context) {
	if err := ufbcb.Exec(ctx); err != nil {
		panic(err)
	}
}
