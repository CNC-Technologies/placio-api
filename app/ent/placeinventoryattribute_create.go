// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/inventoryattribute"
	"placio-app/ent/placeinventory"
	"placio-app/ent/placeinventoryattribute"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PlaceInventoryAttributeCreate is the builder for creating a PlaceInventoryAttribute entity.
type PlaceInventoryAttributeCreate struct {
	config
	mutation *PlaceInventoryAttributeMutation
	hooks    []Hook
}

// SetValue sets the "value" field.
func (piac *PlaceInventoryAttributeCreate) SetValue(s string) *PlaceInventoryAttributeCreate {
	piac.mutation.SetValue(s)
	return piac
}

// SetCategorySpecificValue sets the "category_specific_value" field.
func (piac *PlaceInventoryAttributeCreate) SetCategorySpecificValue(m map[string]interface{}) *PlaceInventoryAttributeCreate {
	piac.mutation.SetCategorySpecificValue(m)
	return piac
}

// SetID sets the "id" field.
func (piac *PlaceInventoryAttributeCreate) SetID(s string) *PlaceInventoryAttributeCreate {
	piac.mutation.SetID(s)
	return piac
}

// SetInventoryID sets the "inventory" edge to the PlaceInventory entity by ID.
func (piac *PlaceInventoryAttributeCreate) SetInventoryID(id string) *PlaceInventoryAttributeCreate {
	piac.mutation.SetInventoryID(id)
	return piac
}

// SetNillableInventoryID sets the "inventory" edge to the PlaceInventory entity by ID if the given value is not nil.
func (piac *PlaceInventoryAttributeCreate) SetNillableInventoryID(id *string) *PlaceInventoryAttributeCreate {
	if id != nil {
		piac = piac.SetInventoryID(*id)
	}
	return piac
}

// SetInventory sets the "inventory" edge to the PlaceInventory entity.
func (piac *PlaceInventoryAttributeCreate) SetInventory(p *PlaceInventory) *PlaceInventoryAttributeCreate {
	return piac.SetInventoryID(p.ID)
}

// SetAttributeTypeID sets the "attribute_type" edge to the InventoryAttribute entity by ID.
func (piac *PlaceInventoryAttributeCreate) SetAttributeTypeID(id string) *PlaceInventoryAttributeCreate {
	piac.mutation.SetAttributeTypeID(id)
	return piac
}

// SetNillableAttributeTypeID sets the "attribute_type" edge to the InventoryAttribute entity by ID if the given value is not nil.
func (piac *PlaceInventoryAttributeCreate) SetNillableAttributeTypeID(id *string) *PlaceInventoryAttributeCreate {
	if id != nil {
		piac = piac.SetAttributeTypeID(*id)
	}
	return piac
}

// SetAttributeType sets the "attribute_type" edge to the InventoryAttribute entity.
func (piac *PlaceInventoryAttributeCreate) SetAttributeType(i *InventoryAttribute) *PlaceInventoryAttributeCreate {
	return piac.SetAttributeTypeID(i.ID)
}

// Mutation returns the PlaceInventoryAttributeMutation object of the builder.
func (piac *PlaceInventoryAttributeCreate) Mutation() *PlaceInventoryAttributeMutation {
	return piac.mutation
}

// Save creates the PlaceInventoryAttribute in the database.
func (piac *PlaceInventoryAttributeCreate) Save(ctx context.Context) (*PlaceInventoryAttribute, error) {
	return withHooks(ctx, piac.sqlSave, piac.mutation, piac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (piac *PlaceInventoryAttributeCreate) SaveX(ctx context.Context) *PlaceInventoryAttribute {
	v, err := piac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (piac *PlaceInventoryAttributeCreate) Exec(ctx context.Context) error {
	_, err := piac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (piac *PlaceInventoryAttributeCreate) ExecX(ctx context.Context) {
	if err := piac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (piac *PlaceInventoryAttributeCreate) check() error {
	if _, ok := piac.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "PlaceInventoryAttribute.value"`)}
	}
	if v, ok := piac.mutation.ID(); ok {
		if err := placeinventoryattribute.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "PlaceInventoryAttribute.id": %w`, err)}
		}
	}
	return nil
}

func (piac *PlaceInventoryAttributeCreate) sqlSave(ctx context.Context) (*PlaceInventoryAttribute, error) {
	if err := piac.check(); err != nil {
		return nil, err
	}
	_node, _spec := piac.createSpec()
	if err := sqlgraph.CreateNode(ctx, piac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected PlaceInventoryAttribute.ID type: %T", _spec.ID.Value)
		}
	}
	piac.mutation.id = &_node.ID
	piac.mutation.done = true
	return _node, nil
}

func (piac *PlaceInventoryAttributeCreate) createSpec() (*PlaceInventoryAttribute, *sqlgraph.CreateSpec) {
	var (
		_node = &PlaceInventoryAttribute{config: piac.config}
		_spec = sqlgraph.NewCreateSpec(placeinventoryattribute.Table, sqlgraph.NewFieldSpec(placeinventoryattribute.FieldID, field.TypeString))
	)
	if id, ok := piac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := piac.mutation.Value(); ok {
		_spec.SetField(placeinventoryattribute.FieldValue, field.TypeString, value)
		_node.Value = value
	}
	if value, ok := piac.mutation.CategorySpecificValue(); ok {
		_spec.SetField(placeinventoryattribute.FieldCategorySpecificValue, field.TypeJSON, value)
		_node.CategorySpecificValue = value
	}
	if nodes := piac.mutation.InventoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   placeinventoryattribute.InventoryTable,
			Columns: []string{placeinventoryattribute.InventoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(placeinventory.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.place_inventory_attributes = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := piac.mutation.AttributeTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   placeinventoryattribute.AttributeTypeTable,
			Columns: []string{placeinventoryattribute.AttributeTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(inventoryattribute.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.inventory_attribute_place_inventory_attributes = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PlaceInventoryAttributeCreateBulk is the builder for creating many PlaceInventoryAttribute entities in bulk.
type PlaceInventoryAttributeCreateBulk struct {
	config
	err      error
	builders []*PlaceInventoryAttributeCreate
}

// Save creates the PlaceInventoryAttribute entities in the database.
func (piacb *PlaceInventoryAttributeCreateBulk) Save(ctx context.Context) ([]*PlaceInventoryAttribute, error) {
	if piacb.err != nil {
		return nil, piacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(piacb.builders))
	nodes := make([]*PlaceInventoryAttribute, len(piacb.builders))
	mutators := make([]Mutator, len(piacb.builders))
	for i := range piacb.builders {
		func(i int, root context.Context) {
			builder := piacb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PlaceInventoryAttributeMutation)
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
					_, err = mutators[i+1].Mutate(root, piacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, piacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, piacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (piacb *PlaceInventoryAttributeCreateBulk) SaveX(ctx context.Context) []*PlaceInventoryAttribute {
	v, err := piacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (piacb *PlaceInventoryAttributeCreateBulk) Exec(ctx context.Context) error {
	_, err := piacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (piacb *PlaceInventoryAttributeCreateBulk) ExecX(ctx context.Context) {
	if err := piacb.Exec(ctx); err != nil {
		panic(err)
	}
}
