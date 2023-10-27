// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/menuitem"
	"placio-app/ent/order"
	"placio-app/ent/orderitem"
	"placio-app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// OrderItemUpdate is the builder for updating OrderItem entities.
type OrderItemUpdate struct {
	config
	hooks    []Hook
	mutation *OrderItemMutation
}

// Where appends a list predicates to the OrderItemUpdate builder.
func (oiu *OrderItemUpdate) Where(ps ...predicate.OrderItem) *OrderItemUpdate {
	oiu.mutation.Where(ps...)
	return oiu
}

// SetQuantity sets the "quantity" field.
func (oiu *OrderItemUpdate) SetQuantity(i int) *OrderItemUpdate {
	oiu.mutation.ResetQuantity()
	oiu.mutation.SetQuantity(i)
	return oiu
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillableQuantity(i *int) *OrderItemUpdate {
	if i != nil {
		oiu.SetQuantity(*i)
	}
	return oiu
}

// AddQuantity adds i to the "quantity" field.
func (oiu *OrderItemUpdate) AddQuantity(i int) *OrderItemUpdate {
	oiu.mutation.AddQuantity(i)
	return oiu
}

// SetPricePerItem sets the "price_per_item" field.
func (oiu *OrderItemUpdate) SetPricePerItem(f float64) *OrderItemUpdate {
	oiu.mutation.ResetPricePerItem()
	oiu.mutation.SetPricePerItem(f)
	return oiu
}

// AddPricePerItem adds f to the "price_per_item" field.
func (oiu *OrderItemUpdate) AddPricePerItem(f float64) *OrderItemUpdate {
	oiu.mutation.AddPricePerItem(f)
	return oiu
}

// SetOptions sets the "options" field.
func (oiu *OrderItemUpdate) SetOptions(s []string) *OrderItemUpdate {
	oiu.mutation.SetOptions(s)
	return oiu
}

// AppendOptions appends s to the "options" field.
func (oiu *OrderItemUpdate) AppendOptions(s []string) *OrderItemUpdate {
	oiu.mutation.AppendOptions(s)
	return oiu
}

// ClearOptions clears the value of the "options" field.
func (oiu *OrderItemUpdate) ClearOptions() *OrderItemUpdate {
	oiu.mutation.ClearOptions()
	return oiu
}

// AddOrderIDs adds the "order" edge to the Order entity by IDs.
func (oiu *OrderItemUpdate) AddOrderIDs(ids ...string) *OrderItemUpdate {
	oiu.mutation.AddOrderIDs(ids...)
	return oiu
}

// AddOrder adds the "order" edges to the Order entity.
func (oiu *OrderItemUpdate) AddOrder(o ...*Order) *OrderItemUpdate {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oiu.AddOrderIDs(ids...)
}

// AddMenuItemIDs adds the "menu_item" edge to the MenuItem entity by IDs.
func (oiu *OrderItemUpdate) AddMenuItemIDs(ids ...string) *OrderItemUpdate {
	oiu.mutation.AddMenuItemIDs(ids...)
	return oiu
}

// AddMenuItem adds the "menu_item" edges to the MenuItem entity.
func (oiu *OrderItemUpdate) AddMenuItem(m ...*MenuItem) *OrderItemUpdate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return oiu.AddMenuItemIDs(ids...)
}

// Mutation returns the OrderItemMutation object of the builder.
func (oiu *OrderItemUpdate) Mutation() *OrderItemMutation {
	return oiu.mutation
}

// ClearOrder clears all "order" edges to the Order entity.
func (oiu *OrderItemUpdate) ClearOrder() *OrderItemUpdate {
	oiu.mutation.ClearOrder()
	return oiu
}

// RemoveOrderIDs removes the "order" edge to Order entities by IDs.
func (oiu *OrderItemUpdate) RemoveOrderIDs(ids ...string) *OrderItemUpdate {
	oiu.mutation.RemoveOrderIDs(ids...)
	return oiu
}

// RemoveOrder removes "order" edges to Order entities.
func (oiu *OrderItemUpdate) RemoveOrder(o ...*Order) *OrderItemUpdate {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oiu.RemoveOrderIDs(ids...)
}

// ClearMenuItem clears all "menu_item" edges to the MenuItem entity.
func (oiu *OrderItemUpdate) ClearMenuItem() *OrderItemUpdate {
	oiu.mutation.ClearMenuItem()
	return oiu
}

// RemoveMenuItemIDs removes the "menu_item" edge to MenuItem entities by IDs.
func (oiu *OrderItemUpdate) RemoveMenuItemIDs(ids ...string) *OrderItemUpdate {
	oiu.mutation.RemoveMenuItemIDs(ids...)
	return oiu
}

// RemoveMenuItem removes "menu_item" edges to MenuItem entities.
func (oiu *OrderItemUpdate) RemoveMenuItem(m ...*MenuItem) *OrderItemUpdate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return oiu.RemoveMenuItemIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (oiu *OrderItemUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, oiu.sqlSave, oiu.mutation, oiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oiu *OrderItemUpdate) SaveX(ctx context.Context) int {
	affected, err := oiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (oiu *OrderItemUpdate) Exec(ctx context.Context) error {
	_, err := oiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oiu *OrderItemUpdate) ExecX(ctx context.Context) {
	if err := oiu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (oiu *OrderItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(orderitem.Table, orderitem.Columns, sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeString))
	if ps := oiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oiu.mutation.Quantity(); ok {
		_spec.SetField(orderitem.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := oiu.mutation.AddedQuantity(); ok {
		_spec.AddField(orderitem.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := oiu.mutation.PricePerItem(); ok {
		_spec.SetField(orderitem.FieldPricePerItem, field.TypeFloat64, value)
	}
	if value, ok := oiu.mutation.AddedPricePerItem(); ok {
		_spec.AddField(orderitem.FieldPricePerItem, field.TypeFloat64, value)
	}
	if value, ok := oiu.mutation.Options(); ok {
		_spec.SetField(orderitem.FieldOptions, field.TypeJSON, value)
	}
	if value, ok := oiu.mutation.AppendedOptions(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, orderitem.FieldOptions, value)
		})
	}
	if oiu.mutation.OptionsCleared() {
		_spec.ClearField(orderitem.FieldOptions, field.TypeJSON)
	}
	if oiu.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   orderitem.OrderTable,
			Columns: orderitem.OrderPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oiu.mutation.RemovedOrderIDs(); len(nodes) > 0 && !oiu.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   orderitem.OrderTable,
			Columns: orderitem.OrderPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oiu.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   orderitem.OrderTable,
			Columns: orderitem.OrderPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if oiu.mutation.MenuItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   orderitem.MenuItemTable,
			Columns: orderitem.MenuItemPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menuitem.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oiu.mutation.RemovedMenuItemIDs(); len(nodes) > 0 && !oiu.mutation.MenuItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   orderitem.MenuItemTable,
			Columns: orderitem.MenuItemPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menuitem.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oiu.mutation.MenuItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   orderitem.MenuItemTable,
			Columns: orderitem.MenuItemPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menuitem.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, oiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	oiu.mutation.done = true
	return n, nil
}

// OrderItemUpdateOne is the builder for updating a single OrderItem entity.
type OrderItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderItemMutation
}

// SetQuantity sets the "quantity" field.
func (oiuo *OrderItemUpdateOne) SetQuantity(i int) *OrderItemUpdateOne {
	oiuo.mutation.ResetQuantity()
	oiuo.mutation.SetQuantity(i)
	return oiuo
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillableQuantity(i *int) *OrderItemUpdateOne {
	if i != nil {
		oiuo.SetQuantity(*i)
	}
	return oiuo
}

// AddQuantity adds i to the "quantity" field.
func (oiuo *OrderItemUpdateOne) AddQuantity(i int) *OrderItemUpdateOne {
	oiuo.mutation.AddQuantity(i)
	return oiuo
}

// SetPricePerItem sets the "price_per_item" field.
func (oiuo *OrderItemUpdateOne) SetPricePerItem(f float64) *OrderItemUpdateOne {
	oiuo.mutation.ResetPricePerItem()
	oiuo.mutation.SetPricePerItem(f)
	return oiuo
}

// AddPricePerItem adds f to the "price_per_item" field.
func (oiuo *OrderItemUpdateOne) AddPricePerItem(f float64) *OrderItemUpdateOne {
	oiuo.mutation.AddPricePerItem(f)
	return oiuo
}

// SetOptions sets the "options" field.
func (oiuo *OrderItemUpdateOne) SetOptions(s []string) *OrderItemUpdateOne {
	oiuo.mutation.SetOptions(s)
	return oiuo
}

// AppendOptions appends s to the "options" field.
func (oiuo *OrderItemUpdateOne) AppendOptions(s []string) *OrderItemUpdateOne {
	oiuo.mutation.AppendOptions(s)
	return oiuo
}

// ClearOptions clears the value of the "options" field.
func (oiuo *OrderItemUpdateOne) ClearOptions() *OrderItemUpdateOne {
	oiuo.mutation.ClearOptions()
	return oiuo
}

// AddOrderIDs adds the "order" edge to the Order entity by IDs.
func (oiuo *OrderItemUpdateOne) AddOrderIDs(ids ...string) *OrderItemUpdateOne {
	oiuo.mutation.AddOrderIDs(ids...)
	return oiuo
}

// AddOrder adds the "order" edges to the Order entity.
func (oiuo *OrderItemUpdateOne) AddOrder(o ...*Order) *OrderItemUpdateOne {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oiuo.AddOrderIDs(ids...)
}

// AddMenuItemIDs adds the "menu_item" edge to the MenuItem entity by IDs.
func (oiuo *OrderItemUpdateOne) AddMenuItemIDs(ids ...string) *OrderItemUpdateOne {
	oiuo.mutation.AddMenuItemIDs(ids...)
	return oiuo
}

// AddMenuItem adds the "menu_item" edges to the MenuItem entity.
func (oiuo *OrderItemUpdateOne) AddMenuItem(m ...*MenuItem) *OrderItemUpdateOne {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return oiuo.AddMenuItemIDs(ids...)
}

// Mutation returns the OrderItemMutation object of the builder.
func (oiuo *OrderItemUpdateOne) Mutation() *OrderItemMutation {
	return oiuo.mutation
}

// ClearOrder clears all "order" edges to the Order entity.
func (oiuo *OrderItemUpdateOne) ClearOrder() *OrderItemUpdateOne {
	oiuo.mutation.ClearOrder()
	return oiuo
}

// RemoveOrderIDs removes the "order" edge to Order entities by IDs.
func (oiuo *OrderItemUpdateOne) RemoveOrderIDs(ids ...string) *OrderItemUpdateOne {
	oiuo.mutation.RemoveOrderIDs(ids...)
	return oiuo
}

// RemoveOrder removes "order" edges to Order entities.
func (oiuo *OrderItemUpdateOne) RemoveOrder(o ...*Order) *OrderItemUpdateOne {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oiuo.RemoveOrderIDs(ids...)
}

// ClearMenuItem clears all "menu_item" edges to the MenuItem entity.
func (oiuo *OrderItemUpdateOne) ClearMenuItem() *OrderItemUpdateOne {
	oiuo.mutation.ClearMenuItem()
	return oiuo
}

// RemoveMenuItemIDs removes the "menu_item" edge to MenuItem entities by IDs.
func (oiuo *OrderItemUpdateOne) RemoveMenuItemIDs(ids ...string) *OrderItemUpdateOne {
	oiuo.mutation.RemoveMenuItemIDs(ids...)
	return oiuo
}

// RemoveMenuItem removes "menu_item" edges to MenuItem entities.
func (oiuo *OrderItemUpdateOne) RemoveMenuItem(m ...*MenuItem) *OrderItemUpdateOne {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return oiuo.RemoveMenuItemIDs(ids...)
}

// Where appends a list predicates to the OrderItemUpdate builder.
func (oiuo *OrderItemUpdateOne) Where(ps ...predicate.OrderItem) *OrderItemUpdateOne {
	oiuo.mutation.Where(ps...)
	return oiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oiuo *OrderItemUpdateOne) Select(field string, fields ...string) *OrderItemUpdateOne {
	oiuo.fields = append([]string{field}, fields...)
	return oiuo
}

// Save executes the query and returns the updated OrderItem entity.
func (oiuo *OrderItemUpdateOne) Save(ctx context.Context) (*OrderItem, error) {
	return withHooks(ctx, oiuo.sqlSave, oiuo.mutation, oiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oiuo *OrderItemUpdateOne) SaveX(ctx context.Context) *OrderItem {
	node, err := oiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oiuo *OrderItemUpdateOne) Exec(ctx context.Context) error {
	_, err := oiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oiuo *OrderItemUpdateOne) ExecX(ctx context.Context) {
	if err := oiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (oiuo *OrderItemUpdateOne) sqlSave(ctx context.Context) (_node *OrderItem, err error) {
	_spec := sqlgraph.NewUpdateSpec(orderitem.Table, orderitem.Columns, sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeString))
	id, ok := oiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrderItem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderitem.FieldID)
		for _, f := range fields {
			if !orderitem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != orderitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oiuo.mutation.Quantity(); ok {
		_spec.SetField(orderitem.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := oiuo.mutation.AddedQuantity(); ok {
		_spec.AddField(orderitem.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := oiuo.mutation.PricePerItem(); ok {
		_spec.SetField(orderitem.FieldPricePerItem, field.TypeFloat64, value)
	}
	if value, ok := oiuo.mutation.AddedPricePerItem(); ok {
		_spec.AddField(orderitem.FieldPricePerItem, field.TypeFloat64, value)
	}
	if value, ok := oiuo.mutation.Options(); ok {
		_spec.SetField(orderitem.FieldOptions, field.TypeJSON, value)
	}
	if value, ok := oiuo.mutation.AppendedOptions(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, orderitem.FieldOptions, value)
		})
	}
	if oiuo.mutation.OptionsCleared() {
		_spec.ClearField(orderitem.FieldOptions, field.TypeJSON)
	}
	if oiuo.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   orderitem.OrderTable,
			Columns: orderitem.OrderPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oiuo.mutation.RemovedOrderIDs(); len(nodes) > 0 && !oiuo.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   orderitem.OrderTable,
			Columns: orderitem.OrderPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oiuo.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   orderitem.OrderTable,
			Columns: orderitem.OrderPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if oiuo.mutation.MenuItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   orderitem.MenuItemTable,
			Columns: orderitem.MenuItemPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menuitem.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oiuo.mutation.RemovedMenuItemIDs(); len(nodes) > 0 && !oiuo.mutation.MenuItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   orderitem.MenuItemTable,
			Columns: orderitem.MenuItemPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menuitem.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oiuo.mutation.MenuItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   orderitem.MenuItemTable,
			Columns: orderitem.MenuItemPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menuitem.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OrderItem{config: oiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	oiuo.mutation.done = true
	return _node, nil
}