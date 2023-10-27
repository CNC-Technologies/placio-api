// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/media"
	"placio-app/ent/menu"
	"placio-app/ent/menuitem"
	"placio-app/ent/orderitem"
	"placio-app/ent/placeinventory"
	"placio-app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// MenuItemUpdate is the builder for updating MenuItem entities.
type MenuItemUpdate struct {
	config
	hooks    []Hook
	mutation *MenuItemMutation
}

// Where appends a list predicates to the MenuItemUpdate builder.
func (miu *MenuItemUpdate) Where(ps ...predicate.MenuItem) *MenuItemUpdate {
	miu.mutation.Where(ps...)
	return miu
}

// SetName sets the "name" field.
func (miu *MenuItemUpdate) SetName(s string) *MenuItemUpdate {
	miu.mutation.SetName(s)
	return miu
}

// SetDescription sets the "description" field.
func (miu *MenuItemUpdate) SetDescription(s string) *MenuItemUpdate {
	miu.mutation.SetDescription(s)
	return miu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (miu *MenuItemUpdate) SetNillableDescription(s *string) *MenuItemUpdate {
	if s != nil {
		miu.SetDescription(*s)
	}
	return miu
}

// ClearDescription clears the value of the "description" field.
func (miu *MenuItemUpdate) ClearDescription() *MenuItemUpdate {
	miu.mutation.ClearDescription()
	return miu
}

// SetPrice sets the "price" field.
func (miu *MenuItemUpdate) SetPrice(f float64) *MenuItemUpdate {
	miu.mutation.ResetPrice()
	miu.mutation.SetPrice(f)
	return miu
}

// AddPrice adds f to the "price" field.
func (miu *MenuItemUpdate) AddPrice(f float64) *MenuItemUpdate {
	miu.mutation.AddPrice(f)
	return miu
}

// SetPreparationTime sets the "preparation_time" field.
func (miu *MenuItemUpdate) SetPreparationTime(i int) *MenuItemUpdate {
	miu.mutation.ResetPreparationTime()
	miu.mutation.SetPreparationTime(i)
	return miu
}

// SetNillablePreparationTime sets the "preparation_time" field if the given value is not nil.
func (miu *MenuItemUpdate) SetNillablePreparationTime(i *int) *MenuItemUpdate {
	if i != nil {
		miu.SetPreparationTime(*i)
	}
	return miu
}

// AddPreparationTime adds i to the "preparation_time" field.
func (miu *MenuItemUpdate) AddPreparationTime(i int) *MenuItemUpdate {
	miu.mutation.AddPreparationTime(i)
	return miu
}

// ClearPreparationTime clears the value of the "preparation_time" field.
func (miu *MenuItemUpdate) ClearPreparationTime() *MenuItemUpdate {
	miu.mutation.ClearPreparationTime()
	return miu
}

// SetOptions sets the "options" field.
func (miu *MenuItemUpdate) SetOptions(s []string) *MenuItemUpdate {
	miu.mutation.SetOptions(s)
	return miu
}

// AppendOptions appends s to the "options" field.
func (miu *MenuItemUpdate) AppendOptions(s []string) *MenuItemUpdate {
	miu.mutation.AppendOptions(s)
	return miu
}

// ClearOptions clears the value of the "options" field.
func (miu *MenuItemUpdate) ClearOptions() *MenuItemUpdate {
	miu.mutation.ClearOptions()
	return miu
}

// SetDeletedAt sets the "deleted_at" field.
func (miu *MenuItemUpdate) SetDeletedAt(s string) *MenuItemUpdate {
	miu.mutation.SetDeletedAt(s)
	return miu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (miu *MenuItemUpdate) SetNillableDeletedAt(s *string) *MenuItemUpdate {
	if s != nil {
		miu.SetDeletedAt(*s)
	}
	return miu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (miu *MenuItemUpdate) ClearDeletedAt() *MenuItemUpdate {
	miu.mutation.ClearDeletedAt()
	return miu
}

// SetIsDeleted sets the "is_deleted" field.
func (miu *MenuItemUpdate) SetIsDeleted(b bool) *MenuItemUpdate {
	miu.mutation.SetIsDeleted(b)
	return miu
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (miu *MenuItemUpdate) SetNillableIsDeleted(b *bool) *MenuItemUpdate {
	if b != nil {
		miu.SetIsDeleted(*b)
	}
	return miu
}

// AddMenuIDs adds the "menu" edge to the Menu entity by IDs.
func (miu *MenuItemUpdate) AddMenuIDs(ids ...string) *MenuItemUpdate {
	miu.mutation.AddMenuIDs(ids...)
	return miu
}

// AddMenu adds the "menu" edges to the Menu entity.
func (miu *MenuItemUpdate) AddMenu(m ...*Menu) *MenuItemUpdate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return miu.AddMenuIDs(ids...)
}

// SetInventoryID sets the "inventory" edge to the PlaceInventory entity by ID.
func (miu *MenuItemUpdate) SetInventoryID(id string) *MenuItemUpdate {
	miu.mutation.SetInventoryID(id)
	return miu
}

// SetNillableInventoryID sets the "inventory" edge to the PlaceInventory entity by ID if the given value is not nil.
func (miu *MenuItemUpdate) SetNillableInventoryID(id *string) *MenuItemUpdate {
	if id != nil {
		miu = miu.SetInventoryID(*id)
	}
	return miu
}

// SetInventory sets the "inventory" edge to the PlaceInventory entity.
func (miu *MenuItemUpdate) SetInventory(p *PlaceInventory) *MenuItemUpdate {
	return miu.SetInventoryID(p.ID)
}

// AddMediumIDs adds the "media" edge to the Media entity by IDs.
func (miu *MenuItemUpdate) AddMediumIDs(ids ...string) *MenuItemUpdate {
	miu.mutation.AddMediumIDs(ids...)
	return miu
}

// AddMedia adds the "media" edges to the Media entity.
func (miu *MenuItemUpdate) AddMedia(m ...*Media) *MenuItemUpdate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return miu.AddMediumIDs(ids...)
}

// AddOrderItemIDs adds the "order_items" edge to the OrderItem entity by IDs.
func (miu *MenuItemUpdate) AddOrderItemIDs(ids ...string) *MenuItemUpdate {
	miu.mutation.AddOrderItemIDs(ids...)
	return miu
}

// AddOrderItems adds the "order_items" edges to the OrderItem entity.
func (miu *MenuItemUpdate) AddOrderItems(o ...*OrderItem) *MenuItemUpdate {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return miu.AddOrderItemIDs(ids...)
}

// Mutation returns the MenuItemMutation object of the builder.
func (miu *MenuItemUpdate) Mutation() *MenuItemMutation {
	return miu.mutation
}

// ClearMenu clears all "menu" edges to the Menu entity.
func (miu *MenuItemUpdate) ClearMenu() *MenuItemUpdate {
	miu.mutation.ClearMenu()
	return miu
}

// RemoveMenuIDs removes the "menu" edge to Menu entities by IDs.
func (miu *MenuItemUpdate) RemoveMenuIDs(ids ...string) *MenuItemUpdate {
	miu.mutation.RemoveMenuIDs(ids...)
	return miu
}

// RemoveMenu removes "menu" edges to Menu entities.
func (miu *MenuItemUpdate) RemoveMenu(m ...*Menu) *MenuItemUpdate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return miu.RemoveMenuIDs(ids...)
}

// ClearInventory clears the "inventory" edge to the PlaceInventory entity.
func (miu *MenuItemUpdate) ClearInventory() *MenuItemUpdate {
	miu.mutation.ClearInventory()
	return miu
}

// ClearMedia clears all "media" edges to the Media entity.
func (miu *MenuItemUpdate) ClearMedia() *MenuItemUpdate {
	miu.mutation.ClearMedia()
	return miu
}

// RemoveMediumIDs removes the "media" edge to Media entities by IDs.
func (miu *MenuItemUpdate) RemoveMediumIDs(ids ...string) *MenuItemUpdate {
	miu.mutation.RemoveMediumIDs(ids...)
	return miu
}

// RemoveMedia removes "media" edges to Media entities.
func (miu *MenuItemUpdate) RemoveMedia(m ...*Media) *MenuItemUpdate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return miu.RemoveMediumIDs(ids...)
}

// ClearOrderItems clears all "order_items" edges to the OrderItem entity.
func (miu *MenuItemUpdate) ClearOrderItems() *MenuItemUpdate {
	miu.mutation.ClearOrderItems()
	return miu
}

// RemoveOrderItemIDs removes the "order_items" edge to OrderItem entities by IDs.
func (miu *MenuItemUpdate) RemoveOrderItemIDs(ids ...string) *MenuItemUpdate {
	miu.mutation.RemoveOrderItemIDs(ids...)
	return miu
}

// RemoveOrderItems removes "order_items" edges to OrderItem entities.
func (miu *MenuItemUpdate) RemoveOrderItems(o ...*OrderItem) *MenuItemUpdate {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return miu.RemoveOrderItemIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (miu *MenuItemUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, miu.sqlSave, miu.mutation, miu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (miu *MenuItemUpdate) SaveX(ctx context.Context) int {
	affected, err := miu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (miu *MenuItemUpdate) Exec(ctx context.Context) error {
	_, err := miu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (miu *MenuItemUpdate) ExecX(ctx context.Context) {
	if err := miu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (miu *MenuItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(menuitem.Table, menuitem.Columns, sqlgraph.NewFieldSpec(menuitem.FieldID, field.TypeString))
	if ps := miu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := miu.mutation.Name(); ok {
		_spec.SetField(menuitem.FieldName, field.TypeString, value)
	}
	if value, ok := miu.mutation.Description(); ok {
		_spec.SetField(menuitem.FieldDescription, field.TypeString, value)
	}
	if miu.mutation.DescriptionCleared() {
		_spec.ClearField(menuitem.FieldDescription, field.TypeString)
	}
	if value, ok := miu.mutation.Price(); ok {
		_spec.SetField(menuitem.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := miu.mutation.AddedPrice(); ok {
		_spec.AddField(menuitem.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := miu.mutation.PreparationTime(); ok {
		_spec.SetField(menuitem.FieldPreparationTime, field.TypeInt, value)
	}
	if value, ok := miu.mutation.AddedPreparationTime(); ok {
		_spec.AddField(menuitem.FieldPreparationTime, field.TypeInt, value)
	}
	if miu.mutation.PreparationTimeCleared() {
		_spec.ClearField(menuitem.FieldPreparationTime, field.TypeInt)
	}
	if value, ok := miu.mutation.Options(); ok {
		_spec.SetField(menuitem.FieldOptions, field.TypeJSON, value)
	}
	if value, ok := miu.mutation.AppendedOptions(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, menuitem.FieldOptions, value)
		})
	}
	if miu.mutation.OptionsCleared() {
		_spec.ClearField(menuitem.FieldOptions, field.TypeJSON)
	}
	if value, ok := miu.mutation.DeletedAt(); ok {
		_spec.SetField(menuitem.FieldDeletedAt, field.TypeString, value)
	}
	if miu.mutation.DeletedAtCleared() {
		_spec.ClearField(menuitem.FieldDeletedAt, field.TypeString)
	}
	if value, ok := miu.mutation.IsDeleted(); ok {
		_spec.SetField(menuitem.FieldIsDeleted, field.TypeBool, value)
	}
	if miu.mutation.MenuCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   menuitem.MenuTable,
			Columns: menuitem.MenuPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := miu.mutation.RemovedMenuIDs(); len(nodes) > 0 && !miu.mutation.MenuCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   menuitem.MenuTable,
			Columns: menuitem.MenuPrimaryKey,
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
	if nodes := miu.mutation.MenuIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   menuitem.MenuTable,
			Columns: menuitem.MenuPrimaryKey,
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
	if miu.mutation.InventoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   menuitem.InventoryTable,
			Columns: []string{menuitem.InventoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(placeinventory.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := miu.mutation.InventoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   menuitem.InventoryTable,
			Columns: []string{menuitem.InventoryColumn},
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
	if miu.mutation.MediaCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menuitem.MediaTable,
			Columns: []string{menuitem.MediaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(media.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := miu.mutation.RemovedMediaIDs(); len(nodes) > 0 && !miu.mutation.MediaCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menuitem.MediaTable,
			Columns: []string{menuitem.MediaColumn},
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
	if nodes := miu.mutation.MediaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menuitem.MediaTable,
			Columns: []string{menuitem.MediaColumn},
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
	if miu.mutation.OrderItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   menuitem.OrderItemsTable,
			Columns: menuitem.OrderItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := miu.mutation.RemovedOrderItemsIDs(); len(nodes) > 0 && !miu.mutation.OrderItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   menuitem.OrderItemsTable,
			Columns: menuitem.OrderItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := miu.mutation.OrderItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   menuitem.OrderItemsTable,
			Columns: menuitem.OrderItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, miu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{menuitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	miu.mutation.done = true
	return n, nil
}

// MenuItemUpdateOne is the builder for updating a single MenuItem entity.
type MenuItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MenuItemMutation
}

// SetName sets the "name" field.
func (miuo *MenuItemUpdateOne) SetName(s string) *MenuItemUpdateOne {
	miuo.mutation.SetName(s)
	return miuo
}

// SetDescription sets the "description" field.
func (miuo *MenuItemUpdateOne) SetDescription(s string) *MenuItemUpdateOne {
	miuo.mutation.SetDescription(s)
	return miuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (miuo *MenuItemUpdateOne) SetNillableDescription(s *string) *MenuItemUpdateOne {
	if s != nil {
		miuo.SetDescription(*s)
	}
	return miuo
}

// ClearDescription clears the value of the "description" field.
func (miuo *MenuItemUpdateOne) ClearDescription() *MenuItemUpdateOne {
	miuo.mutation.ClearDescription()
	return miuo
}

// SetPrice sets the "price" field.
func (miuo *MenuItemUpdateOne) SetPrice(f float64) *MenuItemUpdateOne {
	miuo.mutation.ResetPrice()
	miuo.mutation.SetPrice(f)
	return miuo
}

// AddPrice adds f to the "price" field.
func (miuo *MenuItemUpdateOne) AddPrice(f float64) *MenuItemUpdateOne {
	miuo.mutation.AddPrice(f)
	return miuo
}

// SetPreparationTime sets the "preparation_time" field.
func (miuo *MenuItemUpdateOne) SetPreparationTime(i int) *MenuItemUpdateOne {
	miuo.mutation.ResetPreparationTime()
	miuo.mutation.SetPreparationTime(i)
	return miuo
}

// SetNillablePreparationTime sets the "preparation_time" field if the given value is not nil.
func (miuo *MenuItemUpdateOne) SetNillablePreparationTime(i *int) *MenuItemUpdateOne {
	if i != nil {
		miuo.SetPreparationTime(*i)
	}
	return miuo
}

// AddPreparationTime adds i to the "preparation_time" field.
func (miuo *MenuItemUpdateOne) AddPreparationTime(i int) *MenuItemUpdateOne {
	miuo.mutation.AddPreparationTime(i)
	return miuo
}

// ClearPreparationTime clears the value of the "preparation_time" field.
func (miuo *MenuItemUpdateOne) ClearPreparationTime() *MenuItemUpdateOne {
	miuo.mutation.ClearPreparationTime()
	return miuo
}

// SetOptions sets the "options" field.
func (miuo *MenuItemUpdateOne) SetOptions(s []string) *MenuItemUpdateOne {
	miuo.mutation.SetOptions(s)
	return miuo
}

// AppendOptions appends s to the "options" field.
func (miuo *MenuItemUpdateOne) AppendOptions(s []string) *MenuItemUpdateOne {
	miuo.mutation.AppendOptions(s)
	return miuo
}

// ClearOptions clears the value of the "options" field.
func (miuo *MenuItemUpdateOne) ClearOptions() *MenuItemUpdateOne {
	miuo.mutation.ClearOptions()
	return miuo
}

// SetDeletedAt sets the "deleted_at" field.
func (miuo *MenuItemUpdateOne) SetDeletedAt(s string) *MenuItemUpdateOne {
	miuo.mutation.SetDeletedAt(s)
	return miuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (miuo *MenuItemUpdateOne) SetNillableDeletedAt(s *string) *MenuItemUpdateOne {
	if s != nil {
		miuo.SetDeletedAt(*s)
	}
	return miuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (miuo *MenuItemUpdateOne) ClearDeletedAt() *MenuItemUpdateOne {
	miuo.mutation.ClearDeletedAt()
	return miuo
}

// SetIsDeleted sets the "is_deleted" field.
func (miuo *MenuItemUpdateOne) SetIsDeleted(b bool) *MenuItemUpdateOne {
	miuo.mutation.SetIsDeleted(b)
	return miuo
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (miuo *MenuItemUpdateOne) SetNillableIsDeleted(b *bool) *MenuItemUpdateOne {
	if b != nil {
		miuo.SetIsDeleted(*b)
	}
	return miuo
}

// AddMenuIDs adds the "menu" edge to the Menu entity by IDs.
func (miuo *MenuItemUpdateOne) AddMenuIDs(ids ...string) *MenuItemUpdateOne {
	miuo.mutation.AddMenuIDs(ids...)
	return miuo
}

// AddMenu adds the "menu" edges to the Menu entity.
func (miuo *MenuItemUpdateOne) AddMenu(m ...*Menu) *MenuItemUpdateOne {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return miuo.AddMenuIDs(ids...)
}

// SetInventoryID sets the "inventory" edge to the PlaceInventory entity by ID.
func (miuo *MenuItemUpdateOne) SetInventoryID(id string) *MenuItemUpdateOne {
	miuo.mutation.SetInventoryID(id)
	return miuo
}

// SetNillableInventoryID sets the "inventory" edge to the PlaceInventory entity by ID if the given value is not nil.
func (miuo *MenuItemUpdateOne) SetNillableInventoryID(id *string) *MenuItemUpdateOne {
	if id != nil {
		miuo = miuo.SetInventoryID(*id)
	}
	return miuo
}

// SetInventory sets the "inventory" edge to the PlaceInventory entity.
func (miuo *MenuItemUpdateOne) SetInventory(p *PlaceInventory) *MenuItemUpdateOne {
	return miuo.SetInventoryID(p.ID)
}

// AddMediumIDs adds the "media" edge to the Media entity by IDs.
func (miuo *MenuItemUpdateOne) AddMediumIDs(ids ...string) *MenuItemUpdateOne {
	miuo.mutation.AddMediumIDs(ids...)
	return miuo
}

// AddMedia adds the "media" edges to the Media entity.
func (miuo *MenuItemUpdateOne) AddMedia(m ...*Media) *MenuItemUpdateOne {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return miuo.AddMediumIDs(ids...)
}

// AddOrderItemIDs adds the "order_items" edge to the OrderItem entity by IDs.
func (miuo *MenuItemUpdateOne) AddOrderItemIDs(ids ...string) *MenuItemUpdateOne {
	miuo.mutation.AddOrderItemIDs(ids...)
	return miuo
}

// AddOrderItems adds the "order_items" edges to the OrderItem entity.
func (miuo *MenuItemUpdateOne) AddOrderItems(o ...*OrderItem) *MenuItemUpdateOne {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return miuo.AddOrderItemIDs(ids...)
}

// Mutation returns the MenuItemMutation object of the builder.
func (miuo *MenuItemUpdateOne) Mutation() *MenuItemMutation {
	return miuo.mutation
}

// ClearMenu clears all "menu" edges to the Menu entity.
func (miuo *MenuItemUpdateOne) ClearMenu() *MenuItemUpdateOne {
	miuo.mutation.ClearMenu()
	return miuo
}

// RemoveMenuIDs removes the "menu" edge to Menu entities by IDs.
func (miuo *MenuItemUpdateOne) RemoveMenuIDs(ids ...string) *MenuItemUpdateOne {
	miuo.mutation.RemoveMenuIDs(ids...)
	return miuo
}

// RemoveMenu removes "menu" edges to Menu entities.
func (miuo *MenuItemUpdateOne) RemoveMenu(m ...*Menu) *MenuItemUpdateOne {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return miuo.RemoveMenuIDs(ids...)
}

// ClearInventory clears the "inventory" edge to the PlaceInventory entity.
func (miuo *MenuItemUpdateOne) ClearInventory() *MenuItemUpdateOne {
	miuo.mutation.ClearInventory()
	return miuo
}

// ClearMedia clears all "media" edges to the Media entity.
func (miuo *MenuItemUpdateOne) ClearMedia() *MenuItemUpdateOne {
	miuo.mutation.ClearMedia()
	return miuo
}

// RemoveMediumIDs removes the "media" edge to Media entities by IDs.
func (miuo *MenuItemUpdateOne) RemoveMediumIDs(ids ...string) *MenuItemUpdateOne {
	miuo.mutation.RemoveMediumIDs(ids...)
	return miuo
}

// RemoveMedia removes "media" edges to Media entities.
func (miuo *MenuItemUpdateOne) RemoveMedia(m ...*Media) *MenuItemUpdateOne {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return miuo.RemoveMediumIDs(ids...)
}

// ClearOrderItems clears all "order_items" edges to the OrderItem entity.
func (miuo *MenuItemUpdateOne) ClearOrderItems() *MenuItemUpdateOne {
	miuo.mutation.ClearOrderItems()
	return miuo
}

// RemoveOrderItemIDs removes the "order_items" edge to OrderItem entities by IDs.
func (miuo *MenuItemUpdateOne) RemoveOrderItemIDs(ids ...string) *MenuItemUpdateOne {
	miuo.mutation.RemoveOrderItemIDs(ids...)
	return miuo
}

// RemoveOrderItems removes "order_items" edges to OrderItem entities.
func (miuo *MenuItemUpdateOne) RemoveOrderItems(o ...*OrderItem) *MenuItemUpdateOne {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return miuo.RemoveOrderItemIDs(ids...)
}

// Where appends a list predicates to the MenuItemUpdate builder.
func (miuo *MenuItemUpdateOne) Where(ps ...predicate.MenuItem) *MenuItemUpdateOne {
	miuo.mutation.Where(ps...)
	return miuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (miuo *MenuItemUpdateOne) Select(field string, fields ...string) *MenuItemUpdateOne {
	miuo.fields = append([]string{field}, fields...)
	return miuo
}

// Save executes the query and returns the updated MenuItem entity.
func (miuo *MenuItemUpdateOne) Save(ctx context.Context) (*MenuItem, error) {
	return withHooks(ctx, miuo.sqlSave, miuo.mutation, miuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (miuo *MenuItemUpdateOne) SaveX(ctx context.Context) *MenuItem {
	node, err := miuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (miuo *MenuItemUpdateOne) Exec(ctx context.Context) error {
	_, err := miuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (miuo *MenuItemUpdateOne) ExecX(ctx context.Context) {
	if err := miuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (miuo *MenuItemUpdateOne) sqlSave(ctx context.Context) (_node *MenuItem, err error) {
	_spec := sqlgraph.NewUpdateSpec(menuitem.Table, menuitem.Columns, sqlgraph.NewFieldSpec(menuitem.FieldID, field.TypeString))
	id, ok := miuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MenuItem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := miuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, menuitem.FieldID)
		for _, f := range fields {
			if !menuitem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != menuitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := miuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := miuo.mutation.Name(); ok {
		_spec.SetField(menuitem.FieldName, field.TypeString, value)
	}
	if value, ok := miuo.mutation.Description(); ok {
		_spec.SetField(menuitem.FieldDescription, field.TypeString, value)
	}
	if miuo.mutation.DescriptionCleared() {
		_spec.ClearField(menuitem.FieldDescription, field.TypeString)
	}
	if value, ok := miuo.mutation.Price(); ok {
		_spec.SetField(menuitem.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := miuo.mutation.AddedPrice(); ok {
		_spec.AddField(menuitem.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := miuo.mutation.PreparationTime(); ok {
		_spec.SetField(menuitem.FieldPreparationTime, field.TypeInt, value)
	}
	if value, ok := miuo.mutation.AddedPreparationTime(); ok {
		_spec.AddField(menuitem.FieldPreparationTime, field.TypeInt, value)
	}
	if miuo.mutation.PreparationTimeCleared() {
		_spec.ClearField(menuitem.FieldPreparationTime, field.TypeInt)
	}
	if value, ok := miuo.mutation.Options(); ok {
		_spec.SetField(menuitem.FieldOptions, field.TypeJSON, value)
	}
	if value, ok := miuo.mutation.AppendedOptions(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, menuitem.FieldOptions, value)
		})
	}
	if miuo.mutation.OptionsCleared() {
		_spec.ClearField(menuitem.FieldOptions, field.TypeJSON)
	}
	if value, ok := miuo.mutation.DeletedAt(); ok {
		_spec.SetField(menuitem.FieldDeletedAt, field.TypeString, value)
	}
	if miuo.mutation.DeletedAtCleared() {
		_spec.ClearField(menuitem.FieldDeletedAt, field.TypeString)
	}
	if value, ok := miuo.mutation.IsDeleted(); ok {
		_spec.SetField(menuitem.FieldIsDeleted, field.TypeBool, value)
	}
	if miuo.mutation.MenuCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   menuitem.MenuTable,
			Columns: menuitem.MenuPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := miuo.mutation.RemovedMenuIDs(); len(nodes) > 0 && !miuo.mutation.MenuCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   menuitem.MenuTable,
			Columns: menuitem.MenuPrimaryKey,
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
	if nodes := miuo.mutation.MenuIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   menuitem.MenuTable,
			Columns: menuitem.MenuPrimaryKey,
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
	if miuo.mutation.InventoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   menuitem.InventoryTable,
			Columns: []string{menuitem.InventoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(placeinventory.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := miuo.mutation.InventoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   menuitem.InventoryTable,
			Columns: []string{menuitem.InventoryColumn},
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
	if miuo.mutation.MediaCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menuitem.MediaTable,
			Columns: []string{menuitem.MediaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(media.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := miuo.mutation.RemovedMediaIDs(); len(nodes) > 0 && !miuo.mutation.MediaCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menuitem.MediaTable,
			Columns: []string{menuitem.MediaColumn},
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
	if nodes := miuo.mutation.MediaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menuitem.MediaTable,
			Columns: []string{menuitem.MediaColumn},
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
	if miuo.mutation.OrderItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   menuitem.OrderItemsTable,
			Columns: menuitem.OrderItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := miuo.mutation.RemovedOrderItemsIDs(); len(nodes) > 0 && !miuo.mutation.OrderItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   menuitem.OrderItemsTable,
			Columns: menuitem.OrderItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := miuo.mutation.OrderItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   menuitem.OrderItemsTable,
			Columns: menuitem.OrderItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &MenuItem{config: miuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, miuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{menuitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	miuo.mutation.done = true
	return _node, nil
}