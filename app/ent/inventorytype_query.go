// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"placio-app/ent/inventoryattribute"
	"placio-app/ent/inventorytype"
	"placio-app/ent/placeinventory"
	"placio-app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// InventoryTypeQuery is the builder for querying InventoryType entities.
type InventoryTypeQuery struct {
	config
	ctx                  *QueryContext
	order                []inventorytype.OrderOption
	inters               []Interceptor
	predicates           []predicate.InventoryType
	withAttributes       *InventoryAttributeQuery
	withPlaceInventories *PlaceInventoryQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the InventoryTypeQuery builder.
func (itq *InventoryTypeQuery) Where(ps ...predicate.InventoryType) *InventoryTypeQuery {
	itq.predicates = append(itq.predicates, ps...)
	return itq
}

// Limit the number of records to be returned by this query.
func (itq *InventoryTypeQuery) Limit(limit int) *InventoryTypeQuery {
	itq.ctx.Limit = &limit
	return itq
}

// Offset to start from.
func (itq *InventoryTypeQuery) Offset(offset int) *InventoryTypeQuery {
	itq.ctx.Offset = &offset
	return itq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (itq *InventoryTypeQuery) Unique(unique bool) *InventoryTypeQuery {
	itq.ctx.Unique = &unique
	return itq
}

// Order specifies how the records should be ordered.
func (itq *InventoryTypeQuery) Order(o ...inventorytype.OrderOption) *InventoryTypeQuery {
	itq.order = append(itq.order, o...)
	return itq
}

// QueryAttributes chains the current query on the "attributes" edge.
func (itq *InventoryTypeQuery) QueryAttributes() *InventoryAttributeQuery {
	query := (&InventoryAttributeClient{config: itq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := itq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := itq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(inventorytype.Table, inventorytype.FieldID, selector),
			sqlgraph.To(inventoryattribute.Table, inventoryattribute.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, inventorytype.AttributesTable, inventorytype.AttributesColumn),
		)
		fromU = sqlgraph.SetNeighbors(itq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPlaceInventories chains the current query on the "place_inventories" edge.
func (itq *InventoryTypeQuery) QueryPlaceInventories() *PlaceInventoryQuery {
	query := (&PlaceInventoryClient{config: itq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := itq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := itq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(inventorytype.Table, inventorytype.FieldID, selector),
			sqlgraph.To(placeinventory.Table, placeinventory.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, inventorytype.PlaceInventoriesTable, inventorytype.PlaceInventoriesColumn),
		)
		fromU = sqlgraph.SetNeighbors(itq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first InventoryType entity from the query.
// Returns a *NotFoundError when no InventoryType was found.
func (itq *InventoryTypeQuery) First(ctx context.Context) (*InventoryType, error) {
	nodes, err := itq.Limit(1).All(setContextOp(ctx, itq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{inventorytype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (itq *InventoryTypeQuery) FirstX(ctx context.Context) *InventoryType {
	node, err := itq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first InventoryType ID from the query.
// Returns a *NotFoundError when no InventoryType ID was found.
func (itq *InventoryTypeQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = itq.Limit(1).IDs(setContextOp(ctx, itq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{inventorytype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (itq *InventoryTypeQuery) FirstIDX(ctx context.Context) string {
	id, err := itq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single InventoryType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one InventoryType entity is found.
// Returns a *NotFoundError when no InventoryType entities are found.
func (itq *InventoryTypeQuery) Only(ctx context.Context) (*InventoryType, error) {
	nodes, err := itq.Limit(2).All(setContextOp(ctx, itq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{inventorytype.Label}
	default:
		return nil, &NotSingularError{inventorytype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (itq *InventoryTypeQuery) OnlyX(ctx context.Context) *InventoryType {
	node, err := itq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only InventoryType ID in the query.
// Returns a *NotSingularError when more than one InventoryType ID is found.
// Returns a *NotFoundError when no entities are found.
func (itq *InventoryTypeQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = itq.Limit(2).IDs(setContextOp(ctx, itq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{inventorytype.Label}
	default:
		err = &NotSingularError{inventorytype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (itq *InventoryTypeQuery) OnlyIDX(ctx context.Context) string {
	id, err := itq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of InventoryTypes.
func (itq *InventoryTypeQuery) All(ctx context.Context) ([]*InventoryType, error) {
	ctx = setContextOp(ctx, itq.ctx, "All")
	if err := itq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*InventoryType, *InventoryTypeQuery]()
	return withInterceptors[[]*InventoryType](ctx, itq, qr, itq.inters)
}

// AllX is like All, but panics if an error occurs.
func (itq *InventoryTypeQuery) AllX(ctx context.Context) []*InventoryType {
	nodes, err := itq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of InventoryType IDs.
func (itq *InventoryTypeQuery) IDs(ctx context.Context) (ids []string, err error) {
	if itq.ctx.Unique == nil && itq.path != nil {
		itq.Unique(true)
	}
	ctx = setContextOp(ctx, itq.ctx, "IDs")
	if err = itq.Select(inventorytype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (itq *InventoryTypeQuery) IDsX(ctx context.Context) []string {
	ids, err := itq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (itq *InventoryTypeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, itq.ctx, "Count")
	if err := itq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, itq, querierCount[*InventoryTypeQuery](), itq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (itq *InventoryTypeQuery) CountX(ctx context.Context) int {
	count, err := itq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (itq *InventoryTypeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, itq.ctx, "Exist")
	switch _, err := itq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (itq *InventoryTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := itq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the InventoryTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (itq *InventoryTypeQuery) Clone() *InventoryTypeQuery {
	if itq == nil {
		return nil
	}
	return &InventoryTypeQuery{
		config:               itq.config,
		ctx:                  itq.ctx.Clone(),
		order:                append([]inventorytype.OrderOption{}, itq.order...),
		inters:               append([]Interceptor{}, itq.inters...),
		predicates:           append([]predicate.InventoryType{}, itq.predicates...),
		withAttributes:       itq.withAttributes.Clone(),
		withPlaceInventories: itq.withPlaceInventories.Clone(),
		// clone intermediate query.
		sql:  itq.sql.Clone(),
		path: itq.path,
	}
}

// WithAttributes tells the query-builder to eager-load the nodes that are connected to
// the "attributes" edge. The optional arguments are used to configure the query builder of the edge.
func (itq *InventoryTypeQuery) WithAttributes(opts ...func(*InventoryAttributeQuery)) *InventoryTypeQuery {
	query := (&InventoryAttributeClient{config: itq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	itq.withAttributes = query
	return itq
}

// WithPlaceInventories tells the query-builder to eager-load the nodes that are connected to
// the "place_inventories" edge. The optional arguments are used to configure the query builder of the edge.
func (itq *InventoryTypeQuery) WithPlaceInventories(opts ...func(*PlaceInventoryQuery)) *InventoryTypeQuery {
	query := (&PlaceInventoryClient{config: itq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	itq.withPlaceInventories = query
	return itq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.InventoryType.Query().
//		GroupBy(inventorytype.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (itq *InventoryTypeQuery) GroupBy(field string, fields ...string) *InventoryTypeGroupBy {
	itq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &InventoryTypeGroupBy{build: itq}
	grbuild.flds = &itq.ctx.Fields
	grbuild.label = inventorytype.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.InventoryType.Query().
//		Select(inventorytype.FieldName).
//		Scan(ctx, &v)
func (itq *InventoryTypeQuery) Select(fields ...string) *InventoryTypeSelect {
	itq.ctx.Fields = append(itq.ctx.Fields, fields...)
	sbuild := &InventoryTypeSelect{InventoryTypeQuery: itq}
	sbuild.label = inventorytype.Label
	sbuild.flds, sbuild.scan = &itq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a InventoryTypeSelect configured with the given aggregations.
func (itq *InventoryTypeQuery) Aggregate(fns ...AggregateFunc) *InventoryTypeSelect {
	return itq.Select().Aggregate(fns...)
}

func (itq *InventoryTypeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range itq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, itq); err != nil {
				return err
			}
		}
	}
	for _, f := range itq.ctx.Fields {
		if !inventorytype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if itq.path != nil {
		prev, err := itq.path(ctx)
		if err != nil {
			return err
		}
		itq.sql = prev
	}
	return nil
}

func (itq *InventoryTypeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*InventoryType, error) {
	var (
		nodes       = []*InventoryType{}
		_spec       = itq.querySpec()
		loadedTypes = [2]bool{
			itq.withAttributes != nil,
			itq.withPlaceInventories != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*InventoryType).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &InventoryType{config: itq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, itq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := itq.withAttributes; query != nil {
		if err := itq.loadAttributes(ctx, query, nodes,
			func(n *InventoryType) { n.Edges.Attributes = []*InventoryAttribute{} },
			func(n *InventoryType, e *InventoryAttribute) { n.Edges.Attributes = append(n.Edges.Attributes, e) }); err != nil {
			return nil, err
		}
	}
	if query := itq.withPlaceInventories; query != nil {
		if err := itq.loadPlaceInventories(ctx, query, nodes,
			func(n *InventoryType) { n.Edges.PlaceInventories = []*PlaceInventory{} },
			func(n *InventoryType, e *PlaceInventory) {
				n.Edges.PlaceInventories = append(n.Edges.PlaceInventories, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (itq *InventoryTypeQuery) loadAttributes(ctx context.Context, query *InventoryAttributeQuery, nodes []*InventoryType, init func(*InventoryType), assign func(*InventoryType, *InventoryAttribute)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*InventoryType)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.InventoryAttribute(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(inventorytype.AttributesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.inventory_type_attributes
		if fk == nil {
			return fmt.Errorf(`foreign-key "inventory_type_attributes" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "inventory_type_attributes" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (itq *InventoryTypeQuery) loadPlaceInventories(ctx context.Context, query *PlaceInventoryQuery, nodes []*InventoryType, init func(*InventoryType), assign func(*InventoryType, *PlaceInventory)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*InventoryType)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.PlaceInventory(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(inventorytype.PlaceInventoriesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.inventory_type_place_inventories
		if fk == nil {
			return fmt.Errorf(`foreign-key "inventory_type_place_inventories" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "inventory_type_place_inventories" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (itq *InventoryTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := itq.querySpec()
	_spec.Node.Columns = itq.ctx.Fields
	if len(itq.ctx.Fields) > 0 {
		_spec.Unique = itq.ctx.Unique != nil && *itq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, itq.driver, _spec)
}

func (itq *InventoryTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(inventorytype.Table, inventorytype.Columns, sqlgraph.NewFieldSpec(inventorytype.FieldID, field.TypeString))
	_spec.From = itq.sql
	if unique := itq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if itq.path != nil {
		_spec.Unique = true
	}
	if fields := itq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, inventorytype.FieldID)
		for i := range fields {
			if fields[i] != inventorytype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := itq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := itq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := itq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := itq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (itq *InventoryTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(itq.driver.Dialect())
	t1 := builder.Table(inventorytype.Table)
	columns := itq.ctx.Fields
	if len(columns) == 0 {
		columns = inventorytype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if itq.sql != nil {
		selector = itq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if itq.ctx.Unique != nil && *itq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range itq.predicates {
		p(selector)
	}
	for _, p := range itq.order {
		p(selector)
	}
	if offset := itq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := itq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// InventoryTypeGroupBy is the group-by builder for InventoryType entities.
type InventoryTypeGroupBy struct {
	selector
	build *InventoryTypeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (itgb *InventoryTypeGroupBy) Aggregate(fns ...AggregateFunc) *InventoryTypeGroupBy {
	itgb.fns = append(itgb.fns, fns...)
	return itgb
}

// Scan applies the selector query and scans the result into the given value.
func (itgb *InventoryTypeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, itgb.build.ctx, "GroupBy")
	if err := itgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InventoryTypeQuery, *InventoryTypeGroupBy](ctx, itgb.build, itgb, itgb.build.inters, v)
}

func (itgb *InventoryTypeGroupBy) sqlScan(ctx context.Context, root *InventoryTypeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(itgb.fns))
	for _, fn := range itgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*itgb.flds)+len(itgb.fns))
		for _, f := range *itgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*itgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := itgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// InventoryTypeSelect is the builder for selecting fields of InventoryType entities.
type InventoryTypeSelect struct {
	*InventoryTypeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (its *InventoryTypeSelect) Aggregate(fns ...AggregateFunc) *InventoryTypeSelect {
	its.fns = append(its.fns, fns...)
	return its
}

// Scan applies the selector query and scans the result into the given value.
func (its *InventoryTypeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, its.ctx, "Select")
	if err := its.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InventoryTypeQuery, *InventoryTypeSelect](ctx, its.InventoryTypeQuery, its, its.inters, v)
}

func (its *InventoryTypeSelect) sqlScan(ctx context.Context, root *InventoryTypeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(its.fns))
	for _, fn := range its.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*its.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := its.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
