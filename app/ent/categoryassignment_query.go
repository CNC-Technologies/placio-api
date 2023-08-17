// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"placio-app/ent/business"
	"placio-app/ent/category"
	"placio-app/ent/categoryassignment"
	"placio-app/ent/place"
	"placio-app/ent/predicate"
	"placio-app/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CategoryAssignmentQuery is the builder for querying CategoryAssignment entities.
type CategoryAssignmentQuery struct {
	config
	ctx          *QueryContext
	order        []categoryassignment.OrderOption
	inters       []Interceptor
	predicates   []predicate.CategoryAssignment
	withUser     *UserQuery
	withBusiness *BusinessQuery
	withPlace    *PlaceQuery
	withCategory *CategoryQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CategoryAssignmentQuery builder.
func (caq *CategoryAssignmentQuery) Where(ps ...predicate.CategoryAssignment) *CategoryAssignmentQuery {
	caq.predicates = append(caq.predicates, ps...)
	return caq
}

// Limit the number of records to be returned by this query.
func (caq *CategoryAssignmentQuery) Limit(limit int) *CategoryAssignmentQuery {
	caq.ctx.Limit = &limit
	return caq
}

// Offset to start from.
func (caq *CategoryAssignmentQuery) Offset(offset int) *CategoryAssignmentQuery {
	caq.ctx.Offset = &offset
	return caq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (caq *CategoryAssignmentQuery) Unique(unique bool) *CategoryAssignmentQuery {
	caq.ctx.Unique = &unique
	return caq
}

// Order specifies how the records should be ordered.
func (caq *CategoryAssignmentQuery) Order(o ...categoryassignment.OrderOption) *CategoryAssignmentQuery {
	caq.order = append(caq.order, o...)
	return caq
}

// QueryUser chains the current query on the "user" edge.
func (caq *CategoryAssignmentQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: caq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := caq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := caq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(categoryassignment.Table, categoryassignment.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, categoryassignment.UserTable, categoryassignment.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(caq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBusiness chains the current query on the "business" edge.
func (caq *CategoryAssignmentQuery) QueryBusiness() *BusinessQuery {
	query := (&BusinessClient{config: caq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := caq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := caq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(categoryassignment.Table, categoryassignment.FieldID, selector),
			sqlgraph.To(business.Table, business.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, categoryassignment.BusinessTable, categoryassignment.BusinessColumn),
		)
		fromU = sqlgraph.SetNeighbors(caq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPlace chains the current query on the "place" edge.
func (caq *CategoryAssignmentQuery) QueryPlace() *PlaceQuery {
	query := (&PlaceClient{config: caq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := caq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := caq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(categoryassignment.Table, categoryassignment.FieldID, selector),
			sqlgraph.To(place.Table, place.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, categoryassignment.PlaceTable, categoryassignment.PlaceColumn),
		)
		fromU = sqlgraph.SetNeighbors(caq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCategory chains the current query on the "category" edge.
func (caq *CategoryAssignmentQuery) QueryCategory() *CategoryQuery {
	query := (&CategoryClient{config: caq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := caq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := caq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(categoryassignment.Table, categoryassignment.FieldID, selector),
			sqlgraph.To(category.Table, category.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, categoryassignment.CategoryTable, categoryassignment.CategoryColumn),
		)
		fromU = sqlgraph.SetNeighbors(caq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CategoryAssignment entity from the query.
// Returns a *NotFoundError when no CategoryAssignment was found.
func (caq *CategoryAssignmentQuery) First(ctx context.Context) (*CategoryAssignment, error) {
	nodes, err := caq.Limit(1).All(setContextOp(ctx, caq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{categoryassignment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (caq *CategoryAssignmentQuery) FirstX(ctx context.Context) *CategoryAssignment {
	node, err := caq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CategoryAssignment ID from the query.
// Returns a *NotFoundError when no CategoryAssignment ID was found.
func (caq *CategoryAssignmentQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = caq.Limit(1).IDs(setContextOp(ctx, caq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{categoryassignment.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (caq *CategoryAssignmentQuery) FirstIDX(ctx context.Context) string {
	id, err := caq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CategoryAssignment entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CategoryAssignment entity is found.
// Returns a *NotFoundError when no CategoryAssignment entities are found.
func (caq *CategoryAssignmentQuery) Only(ctx context.Context) (*CategoryAssignment, error) {
	nodes, err := caq.Limit(2).All(setContextOp(ctx, caq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{categoryassignment.Label}
	default:
		return nil, &NotSingularError{categoryassignment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (caq *CategoryAssignmentQuery) OnlyX(ctx context.Context) *CategoryAssignment {
	node, err := caq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CategoryAssignment ID in the query.
// Returns a *NotSingularError when more than one CategoryAssignment ID is found.
// Returns a *NotFoundError when no entities are found.
func (caq *CategoryAssignmentQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = caq.Limit(2).IDs(setContextOp(ctx, caq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{categoryassignment.Label}
	default:
		err = &NotSingularError{categoryassignment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (caq *CategoryAssignmentQuery) OnlyIDX(ctx context.Context) string {
	id, err := caq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CategoryAssignments.
func (caq *CategoryAssignmentQuery) All(ctx context.Context) ([]*CategoryAssignment, error) {
	ctx = setContextOp(ctx, caq.ctx, "All")
	if err := caq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CategoryAssignment, *CategoryAssignmentQuery]()
	return withInterceptors[[]*CategoryAssignment](ctx, caq, qr, caq.inters)
}

// AllX is like All, but panics if an error occurs.
func (caq *CategoryAssignmentQuery) AllX(ctx context.Context) []*CategoryAssignment {
	nodes, err := caq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CategoryAssignment IDs.
func (caq *CategoryAssignmentQuery) IDs(ctx context.Context) (ids []string, err error) {
	if caq.ctx.Unique == nil && caq.path != nil {
		caq.Unique(true)
	}
	ctx = setContextOp(ctx, caq.ctx, "IDs")
	if err = caq.Select(categoryassignment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (caq *CategoryAssignmentQuery) IDsX(ctx context.Context) []string {
	ids, err := caq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (caq *CategoryAssignmentQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, caq.ctx, "Count")
	if err := caq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, caq, querierCount[*CategoryAssignmentQuery](), caq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (caq *CategoryAssignmentQuery) CountX(ctx context.Context) int {
	count, err := caq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (caq *CategoryAssignmentQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, caq.ctx, "Exist")
	switch _, err := caq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (caq *CategoryAssignmentQuery) ExistX(ctx context.Context) bool {
	exist, err := caq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CategoryAssignmentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (caq *CategoryAssignmentQuery) Clone() *CategoryAssignmentQuery {
	if caq == nil {
		return nil
	}
	return &CategoryAssignmentQuery{
		config:       caq.config,
		ctx:          caq.ctx.Clone(),
		order:        append([]categoryassignment.OrderOption{}, caq.order...),
		inters:       append([]Interceptor{}, caq.inters...),
		predicates:   append([]predicate.CategoryAssignment{}, caq.predicates...),
		withUser:     caq.withUser.Clone(),
		withBusiness: caq.withBusiness.Clone(),
		withPlace:    caq.withPlace.Clone(),
		withCategory: caq.withCategory.Clone(),
		// clone intermediate query.
		sql:  caq.sql.Clone(),
		path: caq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (caq *CategoryAssignmentQuery) WithUser(opts ...func(*UserQuery)) *CategoryAssignmentQuery {
	query := (&UserClient{config: caq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	caq.withUser = query
	return caq
}

// WithBusiness tells the query-builder to eager-load the nodes that are connected to
// the "business" edge. The optional arguments are used to configure the query builder of the edge.
func (caq *CategoryAssignmentQuery) WithBusiness(opts ...func(*BusinessQuery)) *CategoryAssignmentQuery {
	query := (&BusinessClient{config: caq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	caq.withBusiness = query
	return caq
}

// WithPlace tells the query-builder to eager-load the nodes that are connected to
// the "place" edge. The optional arguments are used to configure the query builder of the edge.
func (caq *CategoryAssignmentQuery) WithPlace(opts ...func(*PlaceQuery)) *CategoryAssignmentQuery {
	query := (&PlaceClient{config: caq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	caq.withPlace = query
	return caq
}

// WithCategory tells the query-builder to eager-load the nodes that are connected to
// the "category" edge. The optional arguments are used to configure the query builder of the edge.
func (caq *CategoryAssignmentQuery) WithCategory(opts ...func(*CategoryQuery)) *CategoryAssignmentQuery {
	query := (&CategoryClient{config: caq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	caq.withCategory = query
	return caq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		EntityID string `json:"entity_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CategoryAssignment.Query().
//		GroupBy(categoryassignment.FieldEntityID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (caq *CategoryAssignmentQuery) GroupBy(field string, fields ...string) *CategoryAssignmentGroupBy {
	caq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CategoryAssignmentGroupBy{build: caq}
	grbuild.flds = &caq.ctx.Fields
	grbuild.label = categoryassignment.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		EntityID string `json:"entity_id,omitempty"`
//	}
//
//	client.CategoryAssignment.Query().
//		Select(categoryassignment.FieldEntityID).
//		Scan(ctx, &v)
func (caq *CategoryAssignmentQuery) Select(fields ...string) *CategoryAssignmentSelect {
	caq.ctx.Fields = append(caq.ctx.Fields, fields...)
	sbuild := &CategoryAssignmentSelect{CategoryAssignmentQuery: caq}
	sbuild.label = categoryassignment.Label
	sbuild.flds, sbuild.scan = &caq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CategoryAssignmentSelect configured with the given aggregations.
func (caq *CategoryAssignmentQuery) Aggregate(fns ...AggregateFunc) *CategoryAssignmentSelect {
	return caq.Select().Aggregate(fns...)
}

func (caq *CategoryAssignmentQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range caq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, caq); err != nil {
				return err
			}
		}
	}
	for _, f := range caq.ctx.Fields {
		if !categoryassignment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if caq.path != nil {
		prev, err := caq.path(ctx)
		if err != nil {
			return err
		}
		caq.sql = prev
	}
	return nil
}

func (caq *CategoryAssignmentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CategoryAssignment, error) {
	var (
		nodes       = []*CategoryAssignment{}
		withFKs     = caq.withFKs
		_spec       = caq.querySpec()
		loadedTypes = [4]bool{
			caq.withUser != nil,
			caq.withBusiness != nil,
			caq.withPlace != nil,
			caq.withCategory != nil,
		}
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, categoryassignment.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CategoryAssignment).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CategoryAssignment{config: caq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, caq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := caq.withUser; query != nil {
		if err := caq.loadUser(ctx, query, nodes, nil,
			func(n *CategoryAssignment, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := caq.withBusiness; query != nil {
		if err := caq.loadBusiness(ctx, query, nodes, nil,
			func(n *CategoryAssignment, e *Business) { n.Edges.Business = e }); err != nil {
			return nil, err
		}
	}
	if query := caq.withPlace; query != nil {
		if err := caq.loadPlace(ctx, query, nodes, nil,
			func(n *CategoryAssignment, e *Place) { n.Edges.Place = e }); err != nil {
			return nil, err
		}
	}
	if query := caq.withCategory; query != nil {
		if err := caq.loadCategory(ctx, query, nodes, nil,
			func(n *CategoryAssignment, e *Category) { n.Edges.Category = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (caq *CategoryAssignmentQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*CategoryAssignment, init func(*CategoryAssignment), assign func(*CategoryAssignment, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*CategoryAssignment)
	for i := range nodes {
		fk := nodes[i].EntityID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "entity_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (caq *CategoryAssignmentQuery) loadBusiness(ctx context.Context, query *BusinessQuery, nodes []*CategoryAssignment, init func(*CategoryAssignment), assign func(*CategoryAssignment, *Business)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*CategoryAssignment)
	for i := range nodes {
		fk := nodes[i].EntityID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(business.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "entity_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (caq *CategoryAssignmentQuery) loadPlace(ctx context.Context, query *PlaceQuery, nodes []*CategoryAssignment, init func(*CategoryAssignment), assign func(*CategoryAssignment, *Place)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*CategoryAssignment)
	for i := range nodes {
		fk := nodes[i].EntityID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(place.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "entity_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (caq *CategoryAssignmentQuery) loadCategory(ctx context.Context, query *CategoryQuery, nodes []*CategoryAssignment, init func(*CategoryAssignment), assign func(*CategoryAssignment, *Category)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*CategoryAssignment)
	for i := range nodes {
		fk := nodes[i].CategoryID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(category.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "category_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (caq *CategoryAssignmentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := caq.querySpec()
	_spec.Node.Columns = caq.ctx.Fields
	if len(caq.ctx.Fields) > 0 {
		_spec.Unique = caq.ctx.Unique != nil && *caq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, caq.driver, _spec)
}

func (caq *CategoryAssignmentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(categoryassignment.Table, categoryassignment.Columns, sqlgraph.NewFieldSpec(categoryassignment.FieldID, field.TypeString))
	_spec.From = caq.sql
	if unique := caq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if caq.path != nil {
		_spec.Unique = true
	}
	if fields := caq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, categoryassignment.FieldID)
		for i := range fields {
			if fields[i] != categoryassignment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if caq.withUser != nil {
			_spec.Node.AddColumnOnce(categoryassignment.FieldEntityID)
		}
		if caq.withBusiness != nil {
			_spec.Node.AddColumnOnce(categoryassignment.FieldEntityID)
		}
		if caq.withPlace != nil {
			_spec.Node.AddColumnOnce(categoryassignment.FieldEntityID)
		}
		if caq.withCategory != nil {
			_spec.Node.AddColumnOnce(categoryassignment.FieldCategoryID)
		}
	}
	if ps := caq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := caq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := caq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := caq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (caq *CategoryAssignmentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(caq.driver.Dialect())
	t1 := builder.Table(categoryassignment.Table)
	columns := caq.ctx.Fields
	if len(columns) == 0 {
		columns = categoryassignment.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if caq.sql != nil {
		selector = caq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if caq.ctx.Unique != nil && *caq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range caq.predicates {
		p(selector)
	}
	for _, p := range caq.order {
		p(selector)
	}
	if offset := caq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := caq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CategoryAssignmentGroupBy is the group-by builder for CategoryAssignment entities.
type CategoryAssignmentGroupBy struct {
	selector
	build *CategoryAssignmentQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cagb *CategoryAssignmentGroupBy) Aggregate(fns ...AggregateFunc) *CategoryAssignmentGroupBy {
	cagb.fns = append(cagb.fns, fns...)
	return cagb
}

// Scan applies the selector query and scans the result into the given value.
func (cagb *CategoryAssignmentGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cagb.build.ctx, "GroupBy")
	if err := cagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CategoryAssignmentQuery, *CategoryAssignmentGroupBy](ctx, cagb.build, cagb, cagb.build.inters, v)
}

func (cagb *CategoryAssignmentGroupBy) sqlScan(ctx context.Context, root *CategoryAssignmentQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cagb.fns))
	for _, fn := range cagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cagb.flds)+len(cagb.fns))
		for _, f := range *cagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CategoryAssignmentSelect is the builder for selecting fields of CategoryAssignment entities.
type CategoryAssignmentSelect struct {
	*CategoryAssignmentQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cas *CategoryAssignmentSelect) Aggregate(fns ...AggregateFunc) *CategoryAssignmentSelect {
	cas.fns = append(cas.fns, fns...)
	return cas
}

// Scan applies the selector query and scans the result into the given value.
func (cas *CategoryAssignmentSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cas.ctx, "Select")
	if err := cas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CategoryAssignmentQuery, *CategoryAssignmentSelect](ctx, cas.CategoryAssignmentQuery, cas, cas.inters, v)
}

func (cas *CategoryAssignmentSelect) sqlScan(ctx context.Context, root *CategoryAssignmentQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cas.fns))
	for _, fn := range cas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
