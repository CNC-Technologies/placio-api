// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"placio-app/ent/business"
	"placio-app/ent/customblock"
	"placio-app/ent/media"
	"placio-app/ent/predicate"
	"placio-app/ent/template"
	"placio-app/ent/website"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WebsiteQuery is the builder for querying Website entities.
type WebsiteQuery struct {
	config
	ctx              *QueryContext
	order            []website.OrderOption
	inters           []Interceptor
	predicates       []predicate.Website
	withBusiness     *BusinessQuery
	withTemplate     *TemplateQuery
	withCustomBlocks *CustomBlockQuery
	withAssets       *MediaQuery
	withFKs          bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WebsiteQuery builder.
func (wq *WebsiteQuery) Where(ps ...predicate.Website) *WebsiteQuery {
	wq.predicates = append(wq.predicates, ps...)
	return wq
}

// Limit the number of records to be returned by this query.
func (wq *WebsiteQuery) Limit(limit int) *WebsiteQuery {
	wq.ctx.Limit = &limit
	return wq
}

// Offset to start from.
func (wq *WebsiteQuery) Offset(offset int) *WebsiteQuery {
	wq.ctx.Offset = &offset
	return wq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wq *WebsiteQuery) Unique(unique bool) *WebsiteQuery {
	wq.ctx.Unique = &unique
	return wq
}

// Order specifies how the records should be ordered.
func (wq *WebsiteQuery) Order(o ...website.OrderOption) *WebsiteQuery {
	wq.order = append(wq.order, o...)
	return wq
}

// QueryBusiness chains the current query on the "business" edge.
func (wq *WebsiteQuery) QueryBusiness() *BusinessQuery {
	query := (&BusinessClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(website.Table, website.FieldID, selector),
			sqlgraph.To(business.Table, business.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, website.BusinessTable, website.BusinessColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTemplate chains the current query on the "template" edge.
func (wq *WebsiteQuery) QueryTemplate() *TemplateQuery {
	query := (&TemplateClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(website.Table, website.FieldID, selector),
			sqlgraph.To(template.Table, template.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, website.TemplateTable, website.TemplateColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCustomBlocks chains the current query on the "customBlocks" edge.
func (wq *WebsiteQuery) QueryCustomBlocks() *CustomBlockQuery {
	query := (&CustomBlockClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(website.Table, website.FieldID, selector),
			sqlgraph.To(customblock.Table, customblock.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, website.CustomBlocksTable, website.CustomBlocksColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAssets chains the current query on the "assets" edge.
func (wq *WebsiteQuery) QueryAssets() *MediaQuery {
	query := (&MediaClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(website.Table, website.FieldID, selector),
			sqlgraph.To(media.Table, media.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, website.AssetsTable, website.AssetsColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Website entity from the query.
// Returns a *NotFoundError when no Website was found.
func (wq *WebsiteQuery) First(ctx context.Context) (*Website, error) {
	nodes, err := wq.Limit(1).All(setContextOp(ctx, wq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{website.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wq *WebsiteQuery) FirstX(ctx context.Context) *Website {
	node, err := wq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Website ID from the query.
// Returns a *NotFoundError when no Website ID was found.
func (wq *WebsiteQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = wq.Limit(1).IDs(setContextOp(ctx, wq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{website.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wq *WebsiteQuery) FirstIDX(ctx context.Context) string {
	id, err := wq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Website entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Website entity is found.
// Returns a *NotFoundError when no Website entities are found.
func (wq *WebsiteQuery) Only(ctx context.Context) (*Website, error) {
	nodes, err := wq.Limit(2).All(setContextOp(ctx, wq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{website.Label}
	default:
		return nil, &NotSingularError{website.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wq *WebsiteQuery) OnlyX(ctx context.Context) *Website {
	node, err := wq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Website ID in the query.
// Returns a *NotSingularError when more than one Website ID is found.
// Returns a *NotFoundError when no entities are found.
func (wq *WebsiteQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = wq.Limit(2).IDs(setContextOp(ctx, wq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{website.Label}
	default:
		err = &NotSingularError{website.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wq *WebsiteQuery) OnlyIDX(ctx context.Context) string {
	id, err := wq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Websites.
func (wq *WebsiteQuery) All(ctx context.Context) ([]*Website, error) {
	ctx = setContextOp(ctx, wq.ctx, "All")
	if err := wq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Website, *WebsiteQuery]()
	return withInterceptors[[]*Website](ctx, wq, qr, wq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wq *WebsiteQuery) AllX(ctx context.Context) []*Website {
	nodes, err := wq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Website IDs.
func (wq *WebsiteQuery) IDs(ctx context.Context) (ids []string, err error) {
	if wq.ctx.Unique == nil && wq.path != nil {
		wq.Unique(true)
	}
	ctx = setContextOp(ctx, wq.ctx, "IDs")
	if err = wq.Select(website.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wq *WebsiteQuery) IDsX(ctx context.Context) []string {
	ids, err := wq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wq *WebsiteQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wq.ctx, "Count")
	if err := wq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wq, querierCount[*WebsiteQuery](), wq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wq *WebsiteQuery) CountX(ctx context.Context) int {
	count, err := wq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wq *WebsiteQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wq.ctx, "Exist")
	switch _, err := wq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wq *WebsiteQuery) ExistX(ctx context.Context) bool {
	exist, err := wq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WebsiteQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wq *WebsiteQuery) Clone() *WebsiteQuery {
	if wq == nil {
		return nil
	}
	return &WebsiteQuery{
		config:           wq.config,
		ctx:              wq.ctx.Clone(),
		order:            append([]website.OrderOption{}, wq.order...),
		inters:           append([]Interceptor{}, wq.inters...),
		predicates:       append([]predicate.Website{}, wq.predicates...),
		withBusiness:     wq.withBusiness.Clone(),
		withTemplate:     wq.withTemplate.Clone(),
		withCustomBlocks: wq.withCustomBlocks.Clone(),
		withAssets:       wq.withAssets.Clone(),
		// clone intermediate query.
		sql:  wq.sql.Clone(),
		path: wq.path,
	}
}

// WithBusiness tells the query-builder to eager-load the nodes that are connected to
// the "business" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WebsiteQuery) WithBusiness(opts ...func(*BusinessQuery)) *WebsiteQuery {
	query := (&BusinessClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withBusiness = query
	return wq
}

// WithTemplate tells the query-builder to eager-load the nodes that are connected to
// the "template" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WebsiteQuery) WithTemplate(opts ...func(*TemplateQuery)) *WebsiteQuery {
	query := (&TemplateClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withTemplate = query
	return wq
}

// WithCustomBlocks tells the query-builder to eager-load the nodes that are connected to
// the "customBlocks" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WebsiteQuery) WithCustomBlocks(opts ...func(*CustomBlockQuery)) *WebsiteQuery {
	query := (&CustomBlockClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withCustomBlocks = query
	return wq
}

// WithAssets tells the query-builder to eager-load the nodes that are connected to
// the "assets" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WebsiteQuery) WithAssets(opts ...func(*MediaQuery)) *WebsiteQuery {
	query := (&MediaClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withAssets = query
	return wq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DomainName string `json:"domainName,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Website.Query().
//		GroupBy(website.FieldDomainName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (wq *WebsiteQuery) GroupBy(field string, fields ...string) *WebsiteGroupBy {
	wq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WebsiteGroupBy{build: wq}
	grbuild.flds = &wq.ctx.Fields
	grbuild.label = website.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		DomainName string `json:"domainName,omitempty"`
//	}
//
//	client.Website.Query().
//		Select(website.FieldDomainName).
//		Scan(ctx, &v)
func (wq *WebsiteQuery) Select(fields ...string) *WebsiteSelect {
	wq.ctx.Fields = append(wq.ctx.Fields, fields...)
	sbuild := &WebsiteSelect{WebsiteQuery: wq}
	sbuild.label = website.Label
	sbuild.flds, sbuild.scan = &wq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WebsiteSelect configured with the given aggregations.
func (wq *WebsiteQuery) Aggregate(fns ...AggregateFunc) *WebsiteSelect {
	return wq.Select().Aggregate(fns...)
}

func (wq *WebsiteQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wq); err != nil {
				return err
			}
		}
	}
	for _, f := range wq.ctx.Fields {
		if !website.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wq.path != nil {
		prev, err := wq.path(ctx)
		if err != nil {
			return err
		}
		wq.sql = prev
	}
	return nil
}

func (wq *WebsiteQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Website, error) {
	var (
		nodes       = []*Website{}
		withFKs     = wq.withFKs
		_spec       = wq.querySpec()
		loadedTypes = [4]bool{
			wq.withBusiness != nil,
			wq.withTemplate != nil,
			wq.withCustomBlocks != nil,
			wq.withAssets != nil,
		}
	)
	if wq.withBusiness != nil || wq.withTemplate != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, website.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Website).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Website{config: wq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := wq.withBusiness; query != nil {
		if err := wq.loadBusiness(ctx, query, nodes, nil,
			func(n *Website, e *Business) { n.Edges.Business = e }); err != nil {
			return nil, err
		}
	}
	if query := wq.withTemplate; query != nil {
		if err := wq.loadTemplate(ctx, query, nodes, nil,
			func(n *Website, e *Template) { n.Edges.Template = e }); err != nil {
			return nil, err
		}
	}
	if query := wq.withCustomBlocks; query != nil {
		if err := wq.loadCustomBlocks(ctx, query, nodes,
			func(n *Website) { n.Edges.CustomBlocks = []*CustomBlock{} },
			func(n *Website, e *CustomBlock) { n.Edges.CustomBlocks = append(n.Edges.CustomBlocks, e) }); err != nil {
			return nil, err
		}
	}
	if query := wq.withAssets; query != nil {
		if err := wq.loadAssets(ctx, query, nodes,
			func(n *Website) { n.Edges.Assets = []*Media{} },
			func(n *Website, e *Media) { n.Edges.Assets = append(n.Edges.Assets, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (wq *WebsiteQuery) loadBusiness(ctx context.Context, query *BusinessQuery, nodes []*Website, init func(*Website), assign func(*Website, *Business)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Website)
	for i := range nodes {
		if nodes[i].business_websites == nil {
			continue
		}
		fk := *nodes[i].business_websites
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
			return fmt.Errorf(`unexpected foreign-key "business_websites" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (wq *WebsiteQuery) loadTemplate(ctx context.Context, query *TemplateQuery, nodes []*Website, init func(*Website), assign func(*Website, *Template)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Website)
	for i := range nodes {
		if nodes[i].template_websites == nil {
			continue
		}
		fk := *nodes[i].template_websites
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(template.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "template_websites" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (wq *WebsiteQuery) loadCustomBlocks(ctx context.Context, query *CustomBlockQuery, nodes []*Website, init func(*Website), assign func(*Website, *CustomBlock)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Website)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.CustomBlock(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(website.CustomBlocksColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.website_custom_blocks
		if fk == nil {
			return fmt.Errorf(`foreign-key "website_custom_blocks" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "website_custom_blocks" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (wq *WebsiteQuery) loadAssets(ctx context.Context, query *MediaQuery, nodes []*Website, init func(*Website), assign func(*Website, *Media)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Website)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Media(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(website.AssetsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.website_assets
		if fk == nil {
			return fmt.Errorf(`foreign-key "website_assets" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "website_assets" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (wq *WebsiteQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wq.querySpec()
	_spec.Node.Columns = wq.ctx.Fields
	if len(wq.ctx.Fields) > 0 {
		_spec.Unique = wq.ctx.Unique != nil && *wq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wq.driver, _spec)
}

func (wq *WebsiteQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(website.Table, website.Columns, sqlgraph.NewFieldSpec(website.FieldID, field.TypeString))
	_spec.From = wq.sql
	if unique := wq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wq.path != nil {
		_spec.Unique = true
	}
	if fields := wq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, website.FieldID)
		for i := range fields {
			if fields[i] != website.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wq *WebsiteQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wq.driver.Dialect())
	t1 := builder.Table(website.Table)
	columns := wq.ctx.Fields
	if len(columns) == 0 {
		columns = website.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wq.sql != nil {
		selector = wq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wq.ctx.Unique != nil && *wq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range wq.predicates {
		p(selector)
	}
	for _, p := range wq.order {
		p(selector)
	}
	if offset := wq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WebsiteGroupBy is the group-by builder for Website entities.
type WebsiteGroupBy struct {
	selector
	build *WebsiteQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wgb *WebsiteGroupBy) Aggregate(fns ...AggregateFunc) *WebsiteGroupBy {
	wgb.fns = append(wgb.fns, fns...)
	return wgb
}

// Scan applies the selector query and scans the result into the given value.
func (wgb *WebsiteGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wgb.build.ctx, "GroupBy")
	if err := wgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WebsiteQuery, *WebsiteGroupBy](ctx, wgb.build, wgb, wgb.build.inters, v)
}

func (wgb *WebsiteGroupBy) sqlScan(ctx context.Context, root *WebsiteQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wgb.fns))
	for _, fn := range wgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wgb.flds)+len(wgb.fns))
		for _, f := range *wgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WebsiteSelect is the builder for selecting fields of Website entities.
type WebsiteSelect struct {
	*WebsiteQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ws *WebsiteSelect) Aggregate(fns ...AggregateFunc) *WebsiteSelect {
	ws.fns = append(ws.fns, fns...)
	return ws
}

// Scan applies the selector query and scans the result into the given value.
func (ws *WebsiteSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ws.ctx, "Select")
	if err := ws.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WebsiteQuery, *WebsiteSelect](ctx, ws.WebsiteQuery, ws, ws.inters, v)
}

func (ws *WebsiteSelect) sqlScan(ctx context.Context, root *WebsiteQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ws.fns))
	for _, fn := range ws.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ws.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
