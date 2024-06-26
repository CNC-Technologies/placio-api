// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/category"
	"placio-app/ent/categoryassignment"
	"placio-app/ent/media"
	"placio-app/ent/menu"
	"placio-app/ent/placeinventory"
	"placio-app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CategoryUpdate is the builder for updating Category entities.
type CategoryUpdate struct {
	config
	hooks    []Hook
	mutation *CategoryMutation
}

// Where appends a list predicates to the CategoryUpdate builder.
func (cu *CategoryUpdate) Where(ps ...predicate.Category) *CategoryUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *CategoryUpdate) SetName(s string) *CategoryUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableName(s *string) *CategoryUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// SetImage sets the "image" field.
func (cu *CategoryUpdate) SetImage(s string) *CategoryUpdate {
	cu.mutation.SetImage(s)
	return cu
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableImage(s *string) *CategoryUpdate {
	if s != nil {
		cu.SetImage(*s)
	}
	return cu
}

// ClearImage clears the value of the "image" field.
func (cu *CategoryUpdate) ClearImage() *CategoryUpdate {
	cu.mutation.ClearImage()
	return cu
}

// SetDescription sets the "description" field.
func (cu *CategoryUpdate) SetDescription(s string) *CategoryUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableDescription(s *string) *CategoryUpdate {
	if s != nil {
		cu.SetDescription(*s)
	}
	return cu
}

// ClearDescription clears the value of the "description" field.
func (cu *CategoryUpdate) ClearDescription() *CategoryUpdate {
	cu.mutation.ClearDescription()
	return cu
}

// SetIcon sets the "icon" field.
func (cu *CategoryUpdate) SetIcon(s string) *CategoryUpdate {
	cu.mutation.SetIcon(s)
	return cu
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableIcon(s *string) *CategoryUpdate {
	if s != nil {
		cu.SetIcon(*s)
	}
	return cu
}

// ClearIcon clears the value of the "icon" field.
func (cu *CategoryUpdate) ClearIcon() *CategoryUpdate {
	cu.mutation.ClearIcon()
	return cu
}

// SetType sets the "type" field.
func (cu *CategoryUpdate) SetType(s string) *CategoryUpdate {
	cu.mutation.SetType(s)
	return cu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableType(s *string) *CategoryUpdate {
	if s != nil {
		cu.SetType(*s)
	}
	return cu
}

// ClearType clears the value of the "type" field.
func (cu *CategoryUpdate) ClearType() *CategoryUpdate {
	cu.mutation.ClearType()
	return cu
}

// SetParentID sets the "parent_id" field.
func (cu *CategoryUpdate) SetParentID(s string) *CategoryUpdate {
	cu.mutation.SetParentID(s)
	return cu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableParentID(s *string) *CategoryUpdate {
	if s != nil {
		cu.SetParentID(*s)
	}
	return cu
}

// ClearParentID clears the value of the "parent_id" field.
func (cu *CategoryUpdate) ClearParentID() *CategoryUpdate {
	cu.mutation.ClearParentID()
	return cu
}

// SetParentName sets the "parent_name" field.
func (cu *CategoryUpdate) SetParentName(s string) *CategoryUpdate {
	cu.mutation.SetParentName(s)
	return cu
}

// SetNillableParentName sets the "parent_name" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableParentName(s *string) *CategoryUpdate {
	if s != nil {
		cu.SetParentName(*s)
	}
	return cu
}

// ClearParentName clears the value of the "parent_name" field.
func (cu *CategoryUpdate) ClearParentName() *CategoryUpdate {
	cu.mutation.ClearParentName()
	return cu
}

// SetParentImage sets the "parent_image" field.
func (cu *CategoryUpdate) SetParentImage(s string) *CategoryUpdate {
	cu.mutation.SetParentImage(s)
	return cu
}

// SetNillableParentImage sets the "parent_image" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableParentImage(s *string) *CategoryUpdate {
	if s != nil {
		cu.SetParentImage(*s)
	}
	return cu
}

// ClearParentImage clears the value of the "parent_image" field.
func (cu *CategoryUpdate) ClearParentImage() *CategoryUpdate {
	cu.mutation.ClearParentImage()
	return cu
}

// SetParentDescription sets the "parent_description" field.
func (cu *CategoryUpdate) SetParentDescription(s string) *CategoryUpdate {
	cu.mutation.SetParentDescription(s)
	return cu
}

// SetNillableParentDescription sets the "parent_description" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableParentDescription(s *string) *CategoryUpdate {
	if s != nil {
		cu.SetParentDescription(*s)
	}
	return cu
}

// ClearParentDescription clears the value of the "parent_description" field.
func (cu *CategoryUpdate) ClearParentDescription() *CategoryUpdate {
	cu.mutation.ClearParentDescription()
	return cu
}

// AddCategoryAssignmentIDs adds the "categoryAssignments" edge to the CategoryAssignment entity by IDs.
func (cu *CategoryUpdate) AddCategoryAssignmentIDs(ids ...string) *CategoryUpdate {
	cu.mutation.AddCategoryAssignmentIDs(ids...)
	return cu
}

// AddCategoryAssignments adds the "categoryAssignments" edges to the CategoryAssignment entity.
func (cu *CategoryUpdate) AddCategoryAssignments(c ...*CategoryAssignment) *CategoryUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddCategoryAssignmentIDs(ids...)
}

// AddPlaceInventoryIDs adds the "place_inventories" edge to the PlaceInventory entity by IDs.
func (cu *CategoryUpdate) AddPlaceInventoryIDs(ids ...string) *CategoryUpdate {
	cu.mutation.AddPlaceInventoryIDs(ids...)
	return cu
}

// AddPlaceInventories adds the "place_inventories" edges to the PlaceInventory entity.
func (cu *CategoryUpdate) AddPlaceInventories(p ...*PlaceInventory) *CategoryUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.AddPlaceInventoryIDs(ids...)
}

// AddMediumIDs adds the "media" edge to the Media entity by IDs.
func (cu *CategoryUpdate) AddMediumIDs(ids ...string) *CategoryUpdate {
	cu.mutation.AddMediumIDs(ids...)
	return cu
}

// AddMedia adds the "media" edges to the Media entity.
func (cu *CategoryUpdate) AddMedia(m ...*Media) *CategoryUpdate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cu.AddMediumIDs(ids...)
}

// AddMenuIDs adds the "menus" edge to the Menu entity by IDs.
func (cu *CategoryUpdate) AddMenuIDs(ids ...string) *CategoryUpdate {
	cu.mutation.AddMenuIDs(ids...)
	return cu
}

// AddMenus adds the "menus" edges to the Menu entity.
func (cu *CategoryUpdate) AddMenus(m ...*Menu) *CategoryUpdate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cu.AddMenuIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cu *CategoryUpdate) Mutation() *CategoryMutation {
	return cu.mutation
}

// ClearCategoryAssignments clears all "categoryAssignments" edges to the CategoryAssignment entity.
func (cu *CategoryUpdate) ClearCategoryAssignments() *CategoryUpdate {
	cu.mutation.ClearCategoryAssignments()
	return cu
}

// RemoveCategoryAssignmentIDs removes the "categoryAssignments" edge to CategoryAssignment entities by IDs.
func (cu *CategoryUpdate) RemoveCategoryAssignmentIDs(ids ...string) *CategoryUpdate {
	cu.mutation.RemoveCategoryAssignmentIDs(ids...)
	return cu
}

// RemoveCategoryAssignments removes "categoryAssignments" edges to CategoryAssignment entities.
func (cu *CategoryUpdate) RemoveCategoryAssignments(c ...*CategoryAssignment) *CategoryUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveCategoryAssignmentIDs(ids...)
}

// ClearPlaceInventories clears all "place_inventories" edges to the PlaceInventory entity.
func (cu *CategoryUpdate) ClearPlaceInventories() *CategoryUpdate {
	cu.mutation.ClearPlaceInventories()
	return cu
}

// RemovePlaceInventoryIDs removes the "place_inventories" edge to PlaceInventory entities by IDs.
func (cu *CategoryUpdate) RemovePlaceInventoryIDs(ids ...string) *CategoryUpdate {
	cu.mutation.RemovePlaceInventoryIDs(ids...)
	return cu
}

// RemovePlaceInventories removes "place_inventories" edges to PlaceInventory entities.
func (cu *CategoryUpdate) RemovePlaceInventories(p ...*PlaceInventory) *CategoryUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.RemovePlaceInventoryIDs(ids...)
}

// ClearMedia clears all "media" edges to the Media entity.
func (cu *CategoryUpdate) ClearMedia() *CategoryUpdate {
	cu.mutation.ClearMedia()
	return cu
}

// RemoveMediumIDs removes the "media" edge to Media entities by IDs.
func (cu *CategoryUpdate) RemoveMediumIDs(ids ...string) *CategoryUpdate {
	cu.mutation.RemoveMediumIDs(ids...)
	return cu
}

// RemoveMedia removes "media" edges to Media entities.
func (cu *CategoryUpdate) RemoveMedia(m ...*Media) *CategoryUpdate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cu.RemoveMediumIDs(ids...)
}

// ClearMenus clears all "menus" edges to the Menu entity.
func (cu *CategoryUpdate) ClearMenus() *CategoryUpdate {
	cu.mutation.ClearMenus()
	return cu
}

// RemoveMenuIDs removes the "menus" edge to Menu entities by IDs.
func (cu *CategoryUpdate) RemoveMenuIDs(ids ...string) *CategoryUpdate {
	cu.mutation.RemoveMenuIDs(ids...)
	return cu
}

// RemoveMenus removes "menus" edges to Menu entities.
func (cu *CategoryUpdate) RemoveMenus(m ...*Menu) *CategoryUpdate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cu.RemoveMenuIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CategoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CategoryUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CategoryUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(category.Table, category.Columns, sqlgraph.NewFieldSpec(category.FieldID, field.TypeString))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(category.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.Image(); ok {
		_spec.SetField(category.FieldImage, field.TypeString, value)
	}
	if cu.mutation.ImageCleared() {
		_spec.ClearField(category.FieldImage, field.TypeString)
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.SetField(category.FieldDescription, field.TypeString, value)
	}
	if cu.mutation.DescriptionCleared() {
		_spec.ClearField(category.FieldDescription, field.TypeString)
	}
	if value, ok := cu.mutation.Icon(); ok {
		_spec.SetField(category.FieldIcon, field.TypeString, value)
	}
	if cu.mutation.IconCleared() {
		_spec.ClearField(category.FieldIcon, field.TypeString)
	}
	if value, ok := cu.mutation.GetType(); ok {
		_spec.SetField(category.FieldType, field.TypeString, value)
	}
	if cu.mutation.TypeCleared() {
		_spec.ClearField(category.FieldType, field.TypeString)
	}
	if value, ok := cu.mutation.ParentID(); ok {
		_spec.SetField(category.FieldParentID, field.TypeString, value)
	}
	if cu.mutation.ParentIDCleared() {
		_spec.ClearField(category.FieldParentID, field.TypeString)
	}
	if value, ok := cu.mutation.ParentName(); ok {
		_spec.SetField(category.FieldParentName, field.TypeString, value)
	}
	if cu.mutation.ParentNameCleared() {
		_spec.ClearField(category.FieldParentName, field.TypeString)
	}
	if value, ok := cu.mutation.ParentImage(); ok {
		_spec.SetField(category.FieldParentImage, field.TypeString, value)
	}
	if cu.mutation.ParentImageCleared() {
		_spec.ClearField(category.FieldParentImage, field.TypeString)
	}
	if value, ok := cu.mutation.ParentDescription(); ok {
		_spec.SetField(category.FieldParentDescription, field.TypeString, value)
	}
	if cu.mutation.ParentDescriptionCleared() {
		_spec.ClearField(category.FieldParentDescription, field.TypeString)
	}
	if cu.mutation.CategoryAssignmentsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedCategoryAssignmentsIDs(); len(nodes) > 0 && !cu.mutation.CategoryAssignmentsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CategoryAssignmentsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.PlaceInventoriesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedPlaceInventoriesIDs(); len(nodes) > 0 && !cu.mutation.PlaceInventoriesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.PlaceInventoriesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.MediaCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedMediaIDs(); len(nodes) > 0 && !cu.mutation.MediaCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.MediaIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.MenusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.MenusTable,
			Columns: category.MenusPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedMenusIDs(); len(nodes) > 0 && !cu.mutation.MenusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.MenusTable,
			Columns: category.MenusPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.MenusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.MenusTable,
			Columns: category.MenusPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{category.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CategoryUpdateOne is the builder for updating a single Category entity.
type CategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CategoryMutation
}

// SetName sets the "name" field.
func (cuo *CategoryUpdateOne) SetName(s string) *CategoryUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableName(s *string) *CategoryUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// SetImage sets the "image" field.
func (cuo *CategoryUpdateOne) SetImage(s string) *CategoryUpdateOne {
	cuo.mutation.SetImage(s)
	return cuo
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableImage(s *string) *CategoryUpdateOne {
	if s != nil {
		cuo.SetImage(*s)
	}
	return cuo
}

// ClearImage clears the value of the "image" field.
func (cuo *CategoryUpdateOne) ClearImage() *CategoryUpdateOne {
	cuo.mutation.ClearImage()
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *CategoryUpdateOne) SetDescription(s string) *CategoryUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableDescription(s *string) *CategoryUpdateOne {
	if s != nil {
		cuo.SetDescription(*s)
	}
	return cuo
}

// ClearDescription clears the value of the "description" field.
func (cuo *CategoryUpdateOne) ClearDescription() *CategoryUpdateOne {
	cuo.mutation.ClearDescription()
	return cuo
}

// SetIcon sets the "icon" field.
func (cuo *CategoryUpdateOne) SetIcon(s string) *CategoryUpdateOne {
	cuo.mutation.SetIcon(s)
	return cuo
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableIcon(s *string) *CategoryUpdateOne {
	if s != nil {
		cuo.SetIcon(*s)
	}
	return cuo
}

// ClearIcon clears the value of the "icon" field.
func (cuo *CategoryUpdateOne) ClearIcon() *CategoryUpdateOne {
	cuo.mutation.ClearIcon()
	return cuo
}

// SetType sets the "type" field.
func (cuo *CategoryUpdateOne) SetType(s string) *CategoryUpdateOne {
	cuo.mutation.SetType(s)
	return cuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableType(s *string) *CategoryUpdateOne {
	if s != nil {
		cuo.SetType(*s)
	}
	return cuo
}

// ClearType clears the value of the "type" field.
func (cuo *CategoryUpdateOne) ClearType() *CategoryUpdateOne {
	cuo.mutation.ClearType()
	return cuo
}

// SetParentID sets the "parent_id" field.
func (cuo *CategoryUpdateOne) SetParentID(s string) *CategoryUpdateOne {
	cuo.mutation.SetParentID(s)
	return cuo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableParentID(s *string) *CategoryUpdateOne {
	if s != nil {
		cuo.SetParentID(*s)
	}
	return cuo
}

// ClearParentID clears the value of the "parent_id" field.
func (cuo *CategoryUpdateOne) ClearParentID() *CategoryUpdateOne {
	cuo.mutation.ClearParentID()
	return cuo
}

// SetParentName sets the "parent_name" field.
func (cuo *CategoryUpdateOne) SetParentName(s string) *CategoryUpdateOne {
	cuo.mutation.SetParentName(s)
	return cuo
}

// SetNillableParentName sets the "parent_name" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableParentName(s *string) *CategoryUpdateOne {
	if s != nil {
		cuo.SetParentName(*s)
	}
	return cuo
}

// ClearParentName clears the value of the "parent_name" field.
func (cuo *CategoryUpdateOne) ClearParentName() *CategoryUpdateOne {
	cuo.mutation.ClearParentName()
	return cuo
}

// SetParentImage sets the "parent_image" field.
func (cuo *CategoryUpdateOne) SetParentImage(s string) *CategoryUpdateOne {
	cuo.mutation.SetParentImage(s)
	return cuo
}

// SetNillableParentImage sets the "parent_image" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableParentImage(s *string) *CategoryUpdateOne {
	if s != nil {
		cuo.SetParentImage(*s)
	}
	return cuo
}

// ClearParentImage clears the value of the "parent_image" field.
func (cuo *CategoryUpdateOne) ClearParentImage() *CategoryUpdateOne {
	cuo.mutation.ClearParentImage()
	return cuo
}

// SetParentDescription sets the "parent_description" field.
func (cuo *CategoryUpdateOne) SetParentDescription(s string) *CategoryUpdateOne {
	cuo.mutation.SetParentDescription(s)
	return cuo
}

// SetNillableParentDescription sets the "parent_description" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableParentDescription(s *string) *CategoryUpdateOne {
	if s != nil {
		cuo.SetParentDescription(*s)
	}
	return cuo
}

// ClearParentDescription clears the value of the "parent_description" field.
func (cuo *CategoryUpdateOne) ClearParentDescription() *CategoryUpdateOne {
	cuo.mutation.ClearParentDescription()
	return cuo
}

// AddCategoryAssignmentIDs adds the "categoryAssignments" edge to the CategoryAssignment entity by IDs.
func (cuo *CategoryUpdateOne) AddCategoryAssignmentIDs(ids ...string) *CategoryUpdateOne {
	cuo.mutation.AddCategoryAssignmentIDs(ids...)
	return cuo
}

// AddCategoryAssignments adds the "categoryAssignments" edges to the CategoryAssignment entity.
func (cuo *CategoryUpdateOne) AddCategoryAssignments(c ...*CategoryAssignment) *CategoryUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddCategoryAssignmentIDs(ids...)
}

// AddPlaceInventoryIDs adds the "place_inventories" edge to the PlaceInventory entity by IDs.
func (cuo *CategoryUpdateOne) AddPlaceInventoryIDs(ids ...string) *CategoryUpdateOne {
	cuo.mutation.AddPlaceInventoryIDs(ids...)
	return cuo
}

// AddPlaceInventories adds the "place_inventories" edges to the PlaceInventory entity.
func (cuo *CategoryUpdateOne) AddPlaceInventories(p ...*PlaceInventory) *CategoryUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.AddPlaceInventoryIDs(ids...)
}

// AddMediumIDs adds the "media" edge to the Media entity by IDs.
func (cuo *CategoryUpdateOne) AddMediumIDs(ids ...string) *CategoryUpdateOne {
	cuo.mutation.AddMediumIDs(ids...)
	return cuo
}

// AddMedia adds the "media" edges to the Media entity.
func (cuo *CategoryUpdateOne) AddMedia(m ...*Media) *CategoryUpdateOne {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cuo.AddMediumIDs(ids...)
}

// AddMenuIDs adds the "menus" edge to the Menu entity by IDs.
func (cuo *CategoryUpdateOne) AddMenuIDs(ids ...string) *CategoryUpdateOne {
	cuo.mutation.AddMenuIDs(ids...)
	return cuo
}

// AddMenus adds the "menus" edges to the Menu entity.
func (cuo *CategoryUpdateOne) AddMenus(m ...*Menu) *CategoryUpdateOne {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cuo.AddMenuIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cuo *CategoryUpdateOne) Mutation() *CategoryMutation {
	return cuo.mutation
}

// ClearCategoryAssignments clears all "categoryAssignments" edges to the CategoryAssignment entity.
func (cuo *CategoryUpdateOne) ClearCategoryAssignments() *CategoryUpdateOne {
	cuo.mutation.ClearCategoryAssignments()
	return cuo
}

// RemoveCategoryAssignmentIDs removes the "categoryAssignments" edge to CategoryAssignment entities by IDs.
func (cuo *CategoryUpdateOne) RemoveCategoryAssignmentIDs(ids ...string) *CategoryUpdateOne {
	cuo.mutation.RemoveCategoryAssignmentIDs(ids...)
	return cuo
}

// RemoveCategoryAssignments removes "categoryAssignments" edges to CategoryAssignment entities.
func (cuo *CategoryUpdateOne) RemoveCategoryAssignments(c ...*CategoryAssignment) *CategoryUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveCategoryAssignmentIDs(ids...)
}

// ClearPlaceInventories clears all "place_inventories" edges to the PlaceInventory entity.
func (cuo *CategoryUpdateOne) ClearPlaceInventories() *CategoryUpdateOne {
	cuo.mutation.ClearPlaceInventories()
	return cuo
}

// RemovePlaceInventoryIDs removes the "place_inventories" edge to PlaceInventory entities by IDs.
func (cuo *CategoryUpdateOne) RemovePlaceInventoryIDs(ids ...string) *CategoryUpdateOne {
	cuo.mutation.RemovePlaceInventoryIDs(ids...)
	return cuo
}

// RemovePlaceInventories removes "place_inventories" edges to PlaceInventory entities.
func (cuo *CategoryUpdateOne) RemovePlaceInventories(p ...*PlaceInventory) *CategoryUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.RemovePlaceInventoryIDs(ids...)
}

// ClearMedia clears all "media" edges to the Media entity.
func (cuo *CategoryUpdateOne) ClearMedia() *CategoryUpdateOne {
	cuo.mutation.ClearMedia()
	return cuo
}

// RemoveMediumIDs removes the "media" edge to Media entities by IDs.
func (cuo *CategoryUpdateOne) RemoveMediumIDs(ids ...string) *CategoryUpdateOne {
	cuo.mutation.RemoveMediumIDs(ids...)
	return cuo
}

// RemoveMedia removes "media" edges to Media entities.
func (cuo *CategoryUpdateOne) RemoveMedia(m ...*Media) *CategoryUpdateOne {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cuo.RemoveMediumIDs(ids...)
}

// ClearMenus clears all "menus" edges to the Menu entity.
func (cuo *CategoryUpdateOne) ClearMenus() *CategoryUpdateOne {
	cuo.mutation.ClearMenus()
	return cuo
}

// RemoveMenuIDs removes the "menus" edge to Menu entities by IDs.
func (cuo *CategoryUpdateOne) RemoveMenuIDs(ids ...string) *CategoryUpdateOne {
	cuo.mutation.RemoveMenuIDs(ids...)
	return cuo
}

// RemoveMenus removes "menus" edges to Menu entities.
func (cuo *CategoryUpdateOne) RemoveMenus(m ...*Menu) *CategoryUpdateOne {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cuo.RemoveMenuIDs(ids...)
}

// Where appends a list predicates to the CategoryUpdate builder.
func (cuo *CategoryUpdateOne) Where(ps ...predicate.Category) *CategoryUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CategoryUpdateOne) Select(field string, fields ...string) *CategoryUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Category entity.
func (cuo *CategoryUpdateOne) Save(ctx context.Context) (*Category, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CategoryUpdateOne) SaveX(ctx context.Context) *Category {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CategoryUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CategoryUpdateOne) sqlSave(ctx context.Context) (_node *Category, err error) {
	_spec := sqlgraph.NewUpdateSpec(category.Table, category.Columns, sqlgraph.NewFieldSpec(category.FieldID, field.TypeString))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Category.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, category.FieldID)
		for _, f := range fields {
			if !category.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != category.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(category.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Image(); ok {
		_spec.SetField(category.FieldImage, field.TypeString, value)
	}
	if cuo.mutation.ImageCleared() {
		_spec.ClearField(category.FieldImage, field.TypeString)
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.SetField(category.FieldDescription, field.TypeString, value)
	}
	if cuo.mutation.DescriptionCleared() {
		_spec.ClearField(category.FieldDescription, field.TypeString)
	}
	if value, ok := cuo.mutation.Icon(); ok {
		_spec.SetField(category.FieldIcon, field.TypeString, value)
	}
	if cuo.mutation.IconCleared() {
		_spec.ClearField(category.FieldIcon, field.TypeString)
	}
	if value, ok := cuo.mutation.GetType(); ok {
		_spec.SetField(category.FieldType, field.TypeString, value)
	}
	if cuo.mutation.TypeCleared() {
		_spec.ClearField(category.FieldType, field.TypeString)
	}
	if value, ok := cuo.mutation.ParentID(); ok {
		_spec.SetField(category.FieldParentID, field.TypeString, value)
	}
	if cuo.mutation.ParentIDCleared() {
		_spec.ClearField(category.FieldParentID, field.TypeString)
	}
	if value, ok := cuo.mutation.ParentName(); ok {
		_spec.SetField(category.FieldParentName, field.TypeString, value)
	}
	if cuo.mutation.ParentNameCleared() {
		_spec.ClearField(category.FieldParentName, field.TypeString)
	}
	if value, ok := cuo.mutation.ParentImage(); ok {
		_spec.SetField(category.FieldParentImage, field.TypeString, value)
	}
	if cuo.mutation.ParentImageCleared() {
		_spec.ClearField(category.FieldParentImage, field.TypeString)
	}
	if value, ok := cuo.mutation.ParentDescription(); ok {
		_spec.SetField(category.FieldParentDescription, field.TypeString, value)
	}
	if cuo.mutation.ParentDescriptionCleared() {
		_spec.ClearField(category.FieldParentDescription, field.TypeString)
	}
	if cuo.mutation.CategoryAssignmentsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedCategoryAssignmentsIDs(); len(nodes) > 0 && !cuo.mutation.CategoryAssignmentsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CategoryAssignmentsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.PlaceInventoriesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedPlaceInventoriesIDs(); len(nodes) > 0 && !cuo.mutation.PlaceInventoriesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.PlaceInventoriesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.MediaCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedMediaIDs(); len(nodes) > 0 && !cuo.mutation.MediaCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.MediaIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.MenusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.MenusTable,
			Columns: category.MenusPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedMenusIDs(); len(nodes) > 0 && !cuo.mutation.MenusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.MenusTable,
			Columns: category.MenusPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.MenusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.MenusTable,
			Columns: category.MenusPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Category{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{category.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
