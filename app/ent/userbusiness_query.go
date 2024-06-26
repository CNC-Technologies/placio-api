// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"placio-app/ent/business"
	"placio-app/ent/predicate"
	"placio-app/ent/user"
	"placio-app/ent/userbusiness"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserBusinessQuery is the builder for querying UserBusiness entities.
type UserBusinessQuery struct {
	config
	ctx          *QueryContext
	order        []userbusiness.OrderOption
	inters       []Interceptor
	predicates   []predicate.UserBusiness
	withUser     *UserQuery
	withBusiness *BusinessQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserBusinessQuery builder.
func (ubq *UserBusinessQuery) Where(ps ...predicate.UserBusiness) *UserBusinessQuery {
	ubq.predicates = append(ubq.predicates, ps...)
	return ubq
}

// Limit the number of records to be returned by this query.
func (ubq *UserBusinessQuery) Limit(limit int) *UserBusinessQuery {
	ubq.ctx.Limit = &limit
	return ubq
}

// Offset to start from.
func (ubq *UserBusinessQuery) Offset(offset int) *UserBusinessQuery {
	ubq.ctx.Offset = &offset
	return ubq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ubq *UserBusinessQuery) Unique(unique bool) *UserBusinessQuery {
	ubq.ctx.Unique = &unique
	return ubq
}

// Order specifies how the records should be ordered.
func (ubq *UserBusinessQuery) Order(o ...userbusiness.OrderOption) *UserBusinessQuery {
	ubq.order = append(ubq.order, o...)
	return ubq
}

// QueryUser chains the current query on the "user" edge.
func (ubq *UserBusinessQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: ubq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ubq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ubq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userbusiness.Table, userbusiness.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, userbusiness.UserTable, userbusiness.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(ubq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBusiness chains the current query on the "business" edge.
func (ubq *UserBusinessQuery) QueryBusiness() *BusinessQuery {
	query := (&BusinessClient{config: ubq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ubq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ubq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userbusiness.Table, userbusiness.FieldID, selector),
			sqlgraph.To(business.Table, business.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, userbusiness.BusinessTable, userbusiness.BusinessColumn),
		)
		fromU = sqlgraph.SetNeighbors(ubq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserBusiness entity from the query.
// Returns a *NotFoundError when no UserBusiness was found.
func (ubq *UserBusinessQuery) First(ctx context.Context) (*UserBusiness, error) {
	nodes, err := ubq.Limit(1).All(setContextOp(ctx, ubq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userbusiness.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ubq *UserBusinessQuery) FirstX(ctx context.Context) *UserBusiness {
	node, err := ubq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserBusiness ID from the query.
// Returns a *NotFoundError when no UserBusiness ID was found.
func (ubq *UserBusinessQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ubq.Limit(1).IDs(setContextOp(ctx, ubq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userbusiness.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ubq *UserBusinessQuery) FirstIDX(ctx context.Context) string {
	id, err := ubq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserBusiness entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserBusiness entity is found.
// Returns a *NotFoundError when no UserBusiness entities are found.
func (ubq *UserBusinessQuery) Only(ctx context.Context) (*UserBusiness, error) {
	nodes, err := ubq.Limit(2).All(setContextOp(ctx, ubq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userbusiness.Label}
	default:
		return nil, &NotSingularError{userbusiness.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ubq *UserBusinessQuery) OnlyX(ctx context.Context) *UserBusiness {
	node, err := ubq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserBusiness ID in the query.
// Returns a *NotSingularError when more than one UserBusiness ID is found.
// Returns a *NotFoundError when no entities are found.
func (ubq *UserBusinessQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ubq.Limit(2).IDs(setContextOp(ctx, ubq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userbusiness.Label}
	default:
		err = &NotSingularError{userbusiness.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ubq *UserBusinessQuery) OnlyIDX(ctx context.Context) string {
	id, err := ubq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserBusinesses.
func (ubq *UserBusinessQuery) All(ctx context.Context) ([]*UserBusiness, error) {
	ctx = setContextOp(ctx, ubq.ctx, "All")
	if err := ubq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserBusiness, *UserBusinessQuery]()
	return withInterceptors[[]*UserBusiness](ctx, ubq, qr, ubq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ubq *UserBusinessQuery) AllX(ctx context.Context) []*UserBusiness {
	nodes, err := ubq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserBusiness IDs.
func (ubq *UserBusinessQuery) IDs(ctx context.Context) (ids []string, err error) {
	if ubq.ctx.Unique == nil && ubq.path != nil {
		ubq.Unique(true)
	}
	ctx = setContextOp(ctx, ubq.ctx, "IDs")
	if err = ubq.Select(userbusiness.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ubq *UserBusinessQuery) IDsX(ctx context.Context) []string {
	ids, err := ubq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ubq *UserBusinessQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ubq.ctx, "Count")
	if err := ubq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ubq, querierCount[*UserBusinessQuery](), ubq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ubq *UserBusinessQuery) CountX(ctx context.Context) int {
	count, err := ubq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ubq *UserBusinessQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ubq.ctx, "Exist")
	switch _, err := ubq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ubq *UserBusinessQuery) ExistX(ctx context.Context) bool {
	exist, err := ubq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserBusinessQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ubq *UserBusinessQuery) Clone() *UserBusinessQuery {
	if ubq == nil {
		return nil
	}
	return &UserBusinessQuery{
		config:       ubq.config,
		ctx:          ubq.ctx.Clone(),
		order:        append([]userbusiness.OrderOption{}, ubq.order...),
		inters:       append([]Interceptor{}, ubq.inters...),
		predicates:   append([]predicate.UserBusiness{}, ubq.predicates...),
		withUser:     ubq.withUser.Clone(),
		withBusiness: ubq.withBusiness.Clone(),
		// clone intermediate query.
		sql:  ubq.sql.Clone(),
		path: ubq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (ubq *UserBusinessQuery) WithUser(opts ...func(*UserQuery)) *UserBusinessQuery {
	query := (&UserClient{config: ubq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ubq.withUser = query
	return ubq
}

// WithBusiness tells the query-builder to eager-load the nodes that are connected to
// the "business" edge. The optional arguments are used to configure the query builder of the edge.
func (ubq *UserBusinessQuery) WithBusiness(opts ...func(*BusinessQuery)) *UserBusinessQuery {
	query := (&BusinessClient{config: ubq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ubq.withBusiness = query
	return ubq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Role string `json:"role,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserBusiness.Query().
//		GroupBy(userbusiness.FieldRole).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ubq *UserBusinessQuery) GroupBy(field string, fields ...string) *UserBusinessGroupBy {
	ubq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserBusinessGroupBy{build: ubq}
	grbuild.flds = &ubq.ctx.Fields
	grbuild.label = userbusiness.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Role string `json:"role,omitempty"`
//	}
//
//	client.UserBusiness.Query().
//		Select(userbusiness.FieldRole).
//		Scan(ctx, &v)
func (ubq *UserBusinessQuery) Select(fields ...string) *UserBusinessSelect {
	ubq.ctx.Fields = append(ubq.ctx.Fields, fields...)
	sbuild := &UserBusinessSelect{UserBusinessQuery: ubq}
	sbuild.label = userbusiness.Label
	sbuild.flds, sbuild.scan = &ubq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserBusinessSelect configured with the given aggregations.
func (ubq *UserBusinessQuery) Aggregate(fns ...AggregateFunc) *UserBusinessSelect {
	return ubq.Select().Aggregate(fns...)
}

func (ubq *UserBusinessQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ubq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ubq); err != nil {
				return err
			}
		}
	}
	for _, f := range ubq.ctx.Fields {
		if !userbusiness.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ubq.path != nil {
		prev, err := ubq.path(ctx)
		if err != nil {
			return err
		}
		ubq.sql = prev
	}
	return nil
}

func (ubq *UserBusinessQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserBusiness, error) {
	var (
		nodes       = []*UserBusiness{}
		withFKs     = ubq.withFKs
		_spec       = ubq.querySpec()
		loadedTypes = [2]bool{
			ubq.withUser != nil,
			ubq.withBusiness != nil,
		}
	)
	if ubq.withUser != nil || ubq.withBusiness != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, userbusiness.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserBusiness).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserBusiness{config: ubq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ubq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ubq.withUser; query != nil {
		if err := ubq.loadUser(ctx, query, nodes, nil,
			func(n *UserBusiness, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := ubq.withBusiness; query != nil {
		if err := ubq.loadBusiness(ctx, query, nodes, nil,
			func(n *UserBusiness, e *Business) { n.Edges.Business = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ubq *UserBusinessQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*UserBusiness, init func(*UserBusiness), assign func(*UserBusiness, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*UserBusiness)
	for i := range nodes {
		if nodes[i].user_user_businesses == nil {
			continue
		}
		fk := *nodes[i].user_user_businesses
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
			return fmt.Errorf(`unexpected foreign-key "user_user_businesses" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ubq *UserBusinessQuery) loadBusiness(ctx context.Context, query *BusinessQuery, nodes []*UserBusiness, init func(*UserBusiness), assign func(*UserBusiness, *Business)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*UserBusiness)
	for i := range nodes {
		if nodes[i].business_user_businesses == nil {
			continue
		}
		fk := *nodes[i].business_user_businesses
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
			return fmt.Errorf(`unexpected foreign-key "business_user_businesses" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ubq *UserBusinessQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ubq.querySpec()
	_spec.Node.Columns = ubq.ctx.Fields
	if len(ubq.ctx.Fields) > 0 {
		_spec.Unique = ubq.ctx.Unique != nil && *ubq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ubq.driver, _spec)
}

func (ubq *UserBusinessQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(userbusiness.Table, userbusiness.Columns, sqlgraph.NewFieldSpec(userbusiness.FieldID, field.TypeString))
	_spec.From = ubq.sql
	if unique := ubq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ubq.path != nil {
		_spec.Unique = true
	}
	if fields := ubq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userbusiness.FieldID)
		for i := range fields {
			if fields[i] != userbusiness.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ubq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ubq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ubq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ubq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ubq *UserBusinessQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ubq.driver.Dialect())
	t1 := builder.Table(userbusiness.Table)
	columns := ubq.ctx.Fields
	if len(columns) == 0 {
		columns = userbusiness.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ubq.sql != nil {
		selector = ubq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ubq.ctx.Unique != nil && *ubq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ubq.predicates {
		p(selector)
	}
	for _, p := range ubq.order {
		p(selector)
	}
	if offset := ubq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ubq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserBusinessGroupBy is the group-by builder for UserBusiness entities.
type UserBusinessGroupBy struct {
	selector
	build *UserBusinessQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ubgb *UserBusinessGroupBy) Aggregate(fns ...AggregateFunc) *UserBusinessGroupBy {
	ubgb.fns = append(ubgb.fns, fns...)
	return ubgb
}

// Scan applies the selector query and scans the result into the given value.
func (ubgb *UserBusinessGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ubgb.build.ctx, "GroupBy")
	if err := ubgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserBusinessQuery, *UserBusinessGroupBy](ctx, ubgb.build, ubgb, ubgb.build.inters, v)
}

func (ubgb *UserBusinessGroupBy) sqlScan(ctx context.Context, root *UserBusinessQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ubgb.fns))
	for _, fn := range ubgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ubgb.flds)+len(ubgb.fns))
		for _, f := range *ubgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ubgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ubgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserBusinessSelect is the builder for selecting fields of UserBusiness entities.
type UserBusinessSelect struct {
	*UserBusinessQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ubs *UserBusinessSelect) Aggregate(fns ...AggregateFunc) *UserBusinessSelect {
	ubs.fns = append(ubs.fns, fns...)
	return ubs
}

// Scan applies the selector query and scans the result into the given value.
func (ubs *UserBusinessSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ubs.ctx, "Select")
	if err := ubs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserBusinessQuery, *UserBusinessSelect](ctx, ubs.UserBusinessQuery, ubs, ubs.inters, v)
}

func (ubs *UserBusinessSelect) sqlScan(ctx context.Context, root *UserBusinessQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ubs.fns))
	for _, fn := range ubs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ubs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ubs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
