// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/amenity"
	"placio-app/ent/place"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AmenityCreate is the builder for creating a Amenity entity.
type AmenityCreate struct {
	config
	mutation *AmenityMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ac *AmenityCreate) SetName(s string) *AmenityCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetIcon sets the "icon" field.
func (ac *AmenityCreate) SetIcon(s string) *AmenityCreate {
	ac.mutation.SetIcon(s)
	return ac
}

// SetID sets the "id" field.
func (ac *AmenityCreate) SetID(s string) *AmenityCreate {
	ac.mutation.SetID(s)
	return ac
}

// AddPlaceIDs adds the "places" edge to the Place entity by IDs.
func (ac *AmenityCreate) AddPlaceIDs(ids ...string) *AmenityCreate {
	ac.mutation.AddPlaceIDs(ids...)
	return ac
}

// AddPlaces adds the "places" edges to the Place entity.
func (ac *AmenityCreate) AddPlaces(p ...*Place) *AmenityCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ac.AddPlaceIDs(ids...)
}

// Mutation returns the AmenityMutation object of the builder.
func (ac *AmenityCreate) Mutation() *AmenityMutation {
	return ac.mutation
}

// Save creates the Amenity in the database.
func (ac *AmenityCreate) Save(ctx context.Context) (*Amenity, error) {
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AmenityCreate) SaveX(ctx context.Context) *Amenity {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AmenityCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AmenityCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AmenityCreate) check() error {
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Amenity.name"`)}
	}
	if _, ok := ac.mutation.Icon(); !ok {
		return &ValidationError{Name: "icon", err: errors.New(`ent: missing required field "Amenity.icon"`)}
	}
	if v, ok := ac.mutation.ID(); ok {
		if err := amenity.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Amenity.id": %w`, err)}
		}
	}
	return nil
}

func (ac *AmenityCreate) sqlSave(ctx context.Context) (*Amenity, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Amenity.ID type: %T", _spec.ID.Value)
		}
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AmenityCreate) createSpec() (*Amenity, *sqlgraph.CreateSpec) {
	var (
		_node = &Amenity{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(amenity.Table, sqlgraph.NewFieldSpec(amenity.FieldID, field.TypeString))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(amenity.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ac.mutation.Icon(); ok {
		_spec.SetField(amenity.FieldIcon, field.TypeString, value)
		_node.Icon = value
	}
	if nodes := ac.mutation.PlacesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   amenity.PlacesTable,
			Columns: amenity.PlacesPrimaryKey,
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

// AmenityCreateBulk is the builder for creating many Amenity entities in bulk.
type AmenityCreateBulk struct {
	config
	builders []*AmenityCreate
}

// Save creates the Amenity entities in the database.
func (acb *AmenityCreateBulk) Save(ctx context.Context) ([]*Amenity, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Amenity, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AmenityMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AmenityCreateBulk) SaveX(ctx context.Context) []*Amenity {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AmenityCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AmenityCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}