// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/customblock"
	"placio-app/ent/website"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CustomBlockCreate is the builder for creating a CustomBlock entity.
type CustomBlockCreate struct {
	config
	mutation *CustomBlockMutation
	hooks    []Hook
}

// SetContent sets the "content" field.
func (cbc *CustomBlockCreate) SetContent(s string) *CustomBlockCreate {
	cbc.mutation.SetContent(s)
	return cbc
}

// SetWebsiteID sets the "website" edge to the Website entity by ID.
func (cbc *CustomBlockCreate) SetWebsiteID(id string) *CustomBlockCreate {
	cbc.mutation.SetWebsiteID(id)
	return cbc
}

// SetWebsite sets the "website" edge to the Website entity.
func (cbc *CustomBlockCreate) SetWebsite(w *Website) *CustomBlockCreate {
	return cbc.SetWebsiteID(w.ID)
}

// Mutation returns the CustomBlockMutation object of the builder.
func (cbc *CustomBlockCreate) Mutation() *CustomBlockMutation {
	return cbc.mutation
}

// Save creates the CustomBlock in the database.
func (cbc *CustomBlockCreate) Save(ctx context.Context) (*CustomBlock, error) {
	return withHooks(ctx, cbc.sqlSave, cbc.mutation, cbc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cbc *CustomBlockCreate) SaveX(ctx context.Context) *CustomBlock {
	v, err := cbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cbc *CustomBlockCreate) Exec(ctx context.Context) error {
	_, err := cbc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cbc *CustomBlockCreate) ExecX(ctx context.Context) {
	if err := cbc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cbc *CustomBlockCreate) check() error {
	if _, ok := cbc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "CustomBlock.content"`)}
	}
	if _, ok := cbc.mutation.WebsiteID(); !ok {
		return &ValidationError{Name: "website", err: errors.New(`ent: missing required edge "CustomBlock.website"`)}
	}
	return nil
}

func (cbc *CustomBlockCreate) sqlSave(ctx context.Context) (*CustomBlock, error) {
	if err := cbc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cbc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected CustomBlock.ID type: %T", _spec.ID.Value)
		}
	}
	cbc.mutation.id = &_node.ID
	cbc.mutation.done = true
	return _node, nil
}

func (cbc *CustomBlockCreate) createSpec() (*CustomBlock, *sqlgraph.CreateSpec) {
	var (
		_node = &CustomBlock{config: cbc.config}
		_spec = sqlgraph.NewCreateSpec(customblock.Table, sqlgraph.NewFieldSpec(customblock.FieldID, field.TypeString))
	)
	if value, ok := cbc.mutation.Content(); ok {
		_spec.SetField(customblock.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if nodes := cbc.mutation.WebsiteIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customblock.WebsiteTable,
			Columns: []string{customblock.WebsiteColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(website.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.website_custom_blocks = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CustomBlockCreateBulk is the builder for creating many CustomBlock entities in bulk.
type CustomBlockCreateBulk struct {
	config
	err      error
	builders []*CustomBlockCreate
}

// Save creates the CustomBlock entities in the database.
func (cbcb *CustomBlockCreateBulk) Save(ctx context.Context) ([]*CustomBlock, error) {
	if cbcb.err != nil {
		return nil, cbcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(cbcb.builders))
	nodes := make([]*CustomBlock, len(cbcb.builders))
	mutators := make([]Mutator, len(cbcb.builders))
	for i := range cbcb.builders {
		func(i int, root context.Context) {
			builder := cbcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CustomBlockMutation)
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
					_, err = mutators[i+1].Mutate(root, cbcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cbcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cbcb *CustomBlockCreateBulk) SaveX(ctx context.Context) []*CustomBlock {
	v, err := cbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cbcb *CustomBlockCreateBulk) Exec(ctx context.Context) error {
	_, err := cbcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cbcb *CustomBlockCreateBulk) ExecX(ctx context.Context) {
	if err := cbcb.Exec(ctx); err != nil {
		panic(err)
	}
}
