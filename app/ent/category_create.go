// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/category"
	"placio-app/ent/categoryassignment"
	"placio-app/ent/media"
	"placio-app/ent/placeinventory"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CategoryCreate is the builder for creating a Category entity.
type CategoryCreate struct {
	config
	mutation *CategoryMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cc *CategoryCreate) SetName(s string) *CategoryCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetImage sets the "image" field.
func (cc *CategoryCreate) SetImage(s string) *CategoryCreate {
	cc.mutation.SetImage(s)
	return cc
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableImage(s *string) *CategoryCreate {
	if s != nil {
		cc.SetImage(*s)
	}
	return cc
}

// SetDescription sets the "description" field.
func (cc *CategoryCreate) SetDescription(s string) *CategoryCreate {
	cc.mutation.SetDescription(s)
	return cc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableDescription(s *string) *CategoryCreate {
	if s != nil {
		cc.SetDescription(*s)
	}
	return cc
}

// SetIcon sets the "icon" field.
func (cc *CategoryCreate) SetIcon(s string) *CategoryCreate {
	cc.mutation.SetIcon(s)
	return cc
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableIcon(s *string) *CategoryCreate {
	if s != nil {
		cc.SetIcon(*s)
	}
	return cc
}

// SetType sets the "type" field.
func (cc *CategoryCreate) SetType(s string) *CategoryCreate {
	cc.mutation.SetType(s)
	return cc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableType(s *string) *CategoryCreate {
	if s != nil {
		cc.SetType(*s)
	}
	return cc
}

// SetParentID sets the "parent_id" field.
func (cc *CategoryCreate) SetParentID(s string) *CategoryCreate {
	cc.mutation.SetParentID(s)
	return cc
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableParentID(s *string) *CategoryCreate {
	if s != nil {
		cc.SetParentID(*s)
	}
	return cc
}

// SetParentName sets the "parent_name" field.
func (cc *CategoryCreate) SetParentName(s string) *CategoryCreate {
	cc.mutation.SetParentName(s)
	return cc
}

// SetNillableParentName sets the "parent_name" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableParentName(s *string) *CategoryCreate {
	if s != nil {
		cc.SetParentName(*s)
	}
	return cc
}

// SetParentImage sets the "parent_image" field.
func (cc *CategoryCreate) SetParentImage(s string) *CategoryCreate {
	cc.mutation.SetParentImage(s)
	return cc
}

// SetNillableParentImage sets the "parent_image" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableParentImage(s *string) *CategoryCreate {
	if s != nil {
		cc.SetParentImage(*s)
	}
	return cc
}

// SetParentDescription sets the "parent_description" field.
func (cc *CategoryCreate) SetParentDescription(s string) *CategoryCreate {
	cc.mutation.SetParentDescription(s)
	return cc
}

// SetNillableParentDescription sets the "parent_description" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableParentDescription(s *string) *CategoryCreate {
	if s != nil {
		cc.SetParentDescription(*s)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CategoryCreate) SetID(s string) *CategoryCreate {
	cc.mutation.SetID(s)
	return cc
}

// AddCategoryAssignmentIDs adds the "categoryAssignments" edge to the CategoryAssignment entity by IDs.
func (cc *CategoryCreate) AddCategoryAssignmentIDs(ids ...string) *CategoryCreate {
	cc.mutation.AddCategoryAssignmentIDs(ids...)
	return cc
}

// AddCategoryAssignments adds the "categoryAssignments" edges to the CategoryAssignment entity.
func (cc *CategoryCreate) AddCategoryAssignments(c ...*CategoryAssignment) *CategoryCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddCategoryAssignmentIDs(ids...)
}

// AddPlaceInventoryIDs adds the "place_inventories" edge to the PlaceInventory entity by IDs.
func (cc *CategoryCreate) AddPlaceInventoryIDs(ids ...string) *CategoryCreate {
	cc.mutation.AddPlaceInventoryIDs(ids...)
	return cc
}

// AddPlaceInventories adds the "place_inventories" edges to the PlaceInventory entity.
func (cc *CategoryCreate) AddPlaceInventories(p ...*PlaceInventory) *CategoryCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cc.AddPlaceInventoryIDs(ids...)
}

// AddMediumIDs adds the "media" edge to the Media entity by IDs.
func (cc *CategoryCreate) AddMediumIDs(ids ...string) *CategoryCreate {
	cc.mutation.AddMediumIDs(ids...)
	return cc
}

// AddMedia adds the "media" edges to the Media entity.
func (cc *CategoryCreate) AddMedia(m ...*Media) *CategoryCreate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cc.AddMediumIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cc *CategoryCreate) Mutation() *CategoryMutation {
	return cc.mutation
}

// Save creates the Category in the database.
func (cc *CategoryCreate) Save(ctx context.Context) (*Category, error) {
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CategoryCreate) SaveX(ctx context.Context) *Category {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CategoryCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CategoryCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CategoryCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Category.name"`)}
	}
	if v, ok := cc.mutation.ID(); ok {
		if err := category.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Category.id": %w`, err)}
		}
	}
	return nil
}

func (cc *CategoryCreate) sqlSave(ctx context.Context) (*Category, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Category.ID type: %T", _spec.ID.Value)
		}
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CategoryCreate) createSpec() (*Category, *sqlgraph.CreateSpec) {
	var (
		_node = &Category{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(category.Table, sqlgraph.NewFieldSpec(category.FieldID, field.TypeString))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(category.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.Image(); ok {
		_spec.SetField(category.FieldImage, field.TypeString, value)
		_node.Image = value
	}
	if value, ok := cc.mutation.Description(); ok {
		_spec.SetField(category.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := cc.mutation.Icon(); ok {
		_spec.SetField(category.FieldIcon, field.TypeString, value)
		_node.Icon = value
	}
	if value, ok := cc.mutation.GetType(); ok {
		_spec.SetField(category.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := cc.mutation.ParentID(); ok {
		_spec.SetField(category.FieldParentID, field.TypeString, value)
		_node.ParentID = value
	}
	if value, ok := cc.mutation.ParentName(); ok {
		_spec.SetField(category.FieldParentName, field.TypeString, value)
		_node.ParentName = value
	}
	if value, ok := cc.mutation.ParentImage(); ok {
		_spec.SetField(category.FieldParentImage, field.TypeString, value)
		_node.ParentImage = value
	}
	if value, ok := cc.mutation.ParentDescription(); ok {
		_spec.SetField(category.FieldParentDescription, field.TypeString, value)
		_node.ParentDescription = value
	}
	if nodes := cc.mutation.CategoryAssignmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.CategoryAssignmentsTable,
			Columns: []string{category.CategoryAssignmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(categoryassignment.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.PlaceInventoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.PlaceInventoriesTable,
			Columns: []string{category.PlaceInventoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(placeinventory.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.MediaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   category.MediaTable,
			Columns: category.MediaPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(media.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CategoryCreateBulk is the builder for creating many Category entities in bulk.
type CategoryCreateBulk struct {
	config
	err      error
	builders []*CategoryCreate
}

// Save creates the Category entities in the database.
func (ccb *CategoryCreateBulk) Save(ctx context.Context) ([]*Category, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Category, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CategoryMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CategoryCreateBulk) SaveX(ctx context.Context) []*Category {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CategoryCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CategoryCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
