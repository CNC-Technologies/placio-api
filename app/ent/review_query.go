// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"placio-app/ent/business"
	"placio-app/ent/comment"
	"placio-app/ent/event"
	"placio-app/ent/like"
	"placio-app/ent/media"
	"placio-app/ent/place"
	"placio-app/ent/predicate"
	"placio-app/ent/review"
	"placio-app/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ReviewQuery is the builder for querying Review entities.
type ReviewQuery struct {
	config
	ctx          *QueryContext
	order        []review.OrderOption
	inters       []Interceptor
	predicates   []predicate.Review
	withUser     *UserQuery
	withBusiness *BusinessQuery
	withPlace    *PlaceQuery
	withEvent    *EventQuery
	withMedias   *MediaQuery
	withComments *CommentQuery
	withLikes    *LikeQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ReviewQuery builder.
func (rq *ReviewQuery) Where(ps ...predicate.Review) *ReviewQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit the number of records to be returned by this query.
func (rq *ReviewQuery) Limit(limit int) *ReviewQuery {
	rq.ctx.Limit = &limit
	return rq
}

// Offset to start from.
func (rq *ReviewQuery) Offset(offset int) *ReviewQuery {
	rq.ctx.Offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *ReviewQuery) Unique(unique bool) *ReviewQuery {
	rq.ctx.Unique = &unique
	return rq
}

// Order specifies how the records should be ordered.
func (rq *ReviewQuery) Order(o ...review.OrderOption) *ReviewQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryUser chains the current query on the "user" edge.
func (rq *ReviewQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(review.Table, review.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, review.UserTable, review.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBusiness chains the current query on the "business" edge.
func (rq *ReviewQuery) QueryBusiness() *BusinessQuery {
	query := (&BusinessClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(review.Table, review.FieldID, selector),
			sqlgraph.To(business.Table, business.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, review.BusinessTable, review.BusinessColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPlace chains the current query on the "place" edge.
func (rq *ReviewQuery) QueryPlace() *PlaceQuery {
	query := (&PlaceClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(review.Table, review.FieldID, selector),
			sqlgraph.To(place.Table, place.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, review.PlaceTable, review.PlaceColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEvent chains the current query on the "event" edge.
func (rq *ReviewQuery) QueryEvent() *EventQuery {
	query := (&EventClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(review.Table, review.FieldID, selector),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, review.EventTable, review.EventColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMedias chains the current query on the "medias" edge.
func (rq *ReviewQuery) QueryMedias() *MediaQuery {
	query := (&MediaClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(review.Table, review.FieldID, selector),
			sqlgraph.To(media.Table, media.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, review.MediasTable, review.MediasColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryComments chains the current query on the "comments" edge.
func (rq *ReviewQuery) QueryComments() *CommentQuery {
	query := (&CommentClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(review.Table, review.FieldID, selector),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, review.CommentsTable, review.CommentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLikes chains the current query on the "likes" edge.
func (rq *ReviewQuery) QueryLikes() *LikeQuery {
	query := (&LikeClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(review.Table, review.FieldID, selector),
			sqlgraph.To(like.Table, like.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, review.LikesTable, review.LikesColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Review entity from the query.
// Returns a *NotFoundError when no Review was found.
func (rq *ReviewQuery) First(ctx context.Context) (*Review, error) {
	nodes, err := rq.Limit(1).All(setContextOp(ctx, rq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{review.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *ReviewQuery) FirstX(ctx context.Context) *Review {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Review ID from the query.
// Returns a *NotFoundError when no Review ID was found.
func (rq *ReviewQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = rq.Limit(1).IDs(setContextOp(ctx, rq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{review.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *ReviewQuery) FirstIDX(ctx context.Context) string {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Review entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Review entity is found.
// Returns a *NotFoundError when no Review entities are found.
func (rq *ReviewQuery) Only(ctx context.Context) (*Review, error) {
	nodes, err := rq.Limit(2).All(setContextOp(ctx, rq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{review.Label}
	default:
		return nil, &NotSingularError{review.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *ReviewQuery) OnlyX(ctx context.Context) *Review {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Review ID in the query.
// Returns a *NotSingularError when more than one Review ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *ReviewQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = rq.Limit(2).IDs(setContextOp(ctx, rq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{review.Label}
	default:
		err = &NotSingularError{review.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *ReviewQuery) OnlyIDX(ctx context.Context) string {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Reviews.
func (rq *ReviewQuery) All(ctx context.Context) ([]*Review, error) {
	ctx = setContextOp(ctx, rq.ctx, "All")
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Review, *ReviewQuery]()
	return withInterceptors[[]*Review](ctx, rq, qr, rq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rq *ReviewQuery) AllX(ctx context.Context) []*Review {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Review IDs.
func (rq *ReviewQuery) IDs(ctx context.Context) (ids []string, err error) {
	if rq.ctx.Unique == nil && rq.path != nil {
		rq.Unique(true)
	}
	ctx = setContextOp(ctx, rq.ctx, "IDs")
	if err = rq.Select(review.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *ReviewQuery) IDsX(ctx context.Context) []string {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *ReviewQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rq.ctx, "Count")
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rq, querierCount[*ReviewQuery](), rq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rq *ReviewQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *ReviewQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rq.ctx, "Exist")
	switch _, err := rq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *ReviewQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ReviewQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *ReviewQuery) Clone() *ReviewQuery {
	if rq == nil {
		return nil
	}
	return &ReviewQuery{
		config:       rq.config,
		ctx:          rq.ctx.Clone(),
		order:        append([]review.OrderOption{}, rq.order...),
		inters:       append([]Interceptor{}, rq.inters...),
		predicates:   append([]predicate.Review{}, rq.predicates...),
		withUser:     rq.withUser.Clone(),
		withBusiness: rq.withBusiness.Clone(),
		withPlace:    rq.withPlace.Clone(),
		withEvent:    rq.withEvent.Clone(),
		withMedias:   rq.withMedias.Clone(),
		withComments: rq.withComments.Clone(),
		withLikes:    rq.withLikes.Clone(),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *ReviewQuery) WithUser(opts ...func(*UserQuery)) *ReviewQuery {
	query := (&UserClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withUser = query
	return rq
}

// WithBusiness tells the query-builder to eager-load the nodes that are connected to
// the "business" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *ReviewQuery) WithBusiness(opts ...func(*BusinessQuery)) *ReviewQuery {
	query := (&BusinessClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withBusiness = query
	return rq
}

// WithPlace tells the query-builder to eager-load the nodes that are connected to
// the "place" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *ReviewQuery) WithPlace(opts ...func(*PlaceQuery)) *ReviewQuery {
	query := (&PlaceClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withPlace = query
	return rq
}

// WithEvent tells the query-builder to eager-load the nodes that are connected to
// the "event" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *ReviewQuery) WithEvent(opts ...func(*EventQuery)) *ReviewQuery {
	query := (&EventClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withEvent = query
	return rq
}

// WithMedias tells the query-builder to eager-load the nodes that are connected to
// the "medias" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *ReviewQuery) WithMedias(opts ...func(*MediaQuery)) *ReviewQuery {
	query := (&MediaClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withMedias = query
	return rq
}

// WithComments tells the query-builder to eager-load the nodes that are connected to
// the "comments" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *ReviewQuery) WithComments(opts ...func(*CommentQuery)) *ReviewQuery {
	query := (&CommentClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withComments = query
	return rq
}

// WithLikes tells the query-builder to eager-load the nodes that are connected to
// the "likes" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *ReviewQuery) WithLikes(opts ...func(*LikeQuery)) *ReviewQuery {
	query := (&LikeClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withLikes = query
	return rq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Score float64 `json:"score,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Review.Query().
//		GroupBy(review.FieldScore).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rq *ReviewQuery) GroupBy(field string, fields ...string) *ReviewGroupBy {
	rq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ReviewGroupBy{build: rq}
	grbuild.flds = &rq.ctx.Fields
	grbuild.label = review.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Score float64 `json:"score,omitempty"`
//	}
//
//	client.Review.Query().
//		Select(review.FieldScore).
//		Scan(ctx, &v)
func (rq *ReviewQuery) Select(fields ...string) *ReviewSelect {
	rq.ctx.Fields = append(rq.ctx.Fields, fields...)
	sbuild := &ReviewSelect{ReviewQuery: rq}
	sbuild.label = review.Label
	sbuild.flds, sbuild.scan = &rq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ReviewSelect configured with the given aggregations.
func (rq *ReviewQuery) Aggregate(fns ...AggregateFunc) *ReviewSelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *ReviewQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rq); err != nil {
				return err
			}
		}
	}
	for _, f := range rq.ctx.Fields {
		if !review.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *ReviewQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Review, error) {
	var (
		nodes       = []*Review{}
		withFKs     = rq.withFKs
		_spec       = rq.querySpec()
		loadedTypes = [7]bool{
			rq.withUser != nil,
			rq.withBusiness != nil,
			rq.withPlace != nil,
			rq.withEvent != nil,
			rq.withMedias != nil,
			rq.withComments != nil,
			rq.withLikes != nil,
		}
	)
	if rq.withUser != nil || rq.withBusiness != nil || rq.withPlace != nil || rq.withEvent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, review.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Review).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Review{config: rq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rq.withUser; query != nil {
		if err := rq.loadUser(ctx, query, nodes, nil,
			func(n *Review, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := rq.withBusiness; query != nil {
		if err := rq.loadBusiness(ctx, query, nodes, nil,
			func(n *Review, e *Business) { n.Edges.Business = e }); err != nil {
			return nil, err
		}
	}
	if query := rq.withPlace; query != nil {
		if err := rq.loadPlace(ctx, query, nodes, nil,
			func(n *Review, e *Place) { n.Edges.Place = e }); err != nil {
			return nil, err
		}
	}
	if query := rq.withEvent; query != nil {
		if err := rq.loadEvent(ctx, query, nodes, nil,
			func(n *Review, e *Event) { n.Edges.Event = e }); err != nil {
			return nil, err
		}
	}
	if query := rq.withMedias; query != nil {
		if err := rq.loadMedias(ctx, query, nodes,
			func(n *Review) { n.Edges.Medias = []*Media{} },
			func(n *Review, e *Media) { n.Edges.Medias = append(n.Edges.Medias, e) }); err != nil {
			return nil, err
		}
	}
	if query := rq.withComments; query != nil {
		if err := rq.loadComments(ctx, query, nodes,
			func(n *Review) { n.Edges.Comments = []*Comment{} },
			func(n *Review, e *Comment) { n.Edges.Comments = append(n.Edges.Comments, e) }); err != nil {
			return nil, err
		}
	}
	if query := rq.withLikes; query != nil {
		if err := rq.loadLikes(ctx, query, nodes,
			func(n *Review) { n.Edges.Likes = []*Like{} },
			func(n *Review, e *Like) { n.Edges.Likes = append(n.Edges.Likes, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rq *ReviewQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Review, init func(*Review), assign func(*Review, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Review)
	for i := range nodes {
		if nodes[i].user_reviews == nil {
			continue
		}
		fk := *nodes[i].user_reviews
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
			return fmt.Errorf(`unexpected foreign-key "user_reviews" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (rq *ReviewQuery) loadBusiness(ctx context.Context, query *BusinessQuery, nodes []*Review, init func(*Review), assign func(*Review, *Business)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Review)
	for i := range nodes {
		if nodes[i].review_business == nil {
			continue
		}
		fk := *nodes[i].review_business
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
			return fmt.Errorf(`unexpected foreign-key "review_business" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (rq *ReviewQuery) loadPlace(ctx context.Context, query *PlaceQuery, nodes []*Review, init func(*Review), assign func(*Review, *Place)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Review)
	for i := range nodes {
		if nodes[i].review_place == nil {
			continue
		}
		fk := *nodes[i].review_place
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
			return fmt.Errorf(`unexpected foreign-key "review_place" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (rq *ReviewQuery) loadEvent(ctx context.Context, query *EventQuery, nodes []*Review, init func(*Review), assign func(*Review, *Event)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Review)
	for i := range nodes {
		if nodes[i].review_event == nil {
			continue
		}
		fk := *nodes[i].review_event
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(event.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "review_event" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (rq *ReviewQuery) loadMedias(ctx context.Context, query *MediaQuery, nodes []*Review, init func(*Review), assign func(*Review, *Media)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Review)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Media(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(review.MediasColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.review_medias
		if fk == nil {
			return fmt.Errorf(`foreign-key "review_medias" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "review_medias" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (rq *ReviewQuery) loadComments(ctx context.Context, query *CommentQuery, nodes []*Review, init func(*Review), assign func(*Review, *Comment)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Review)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(review.CommentsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.review_comments
		if fk == nil {
			return fmt.Errorf(`foreign-key "review_comments" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "review_comments" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (rq *ReviewQuery) loadLikes(ctx context.Context, query *LikeQuery, nodes []*Review, init func(*Review), assign func(*Review, *Like)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Review)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Like(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(review.LikesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.review_likes
		if fk == nil {
			return fmt.Errorf(`foreign-key "review_likes" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "review_likes" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (rq *ReviewQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	_spec.Node.Columns = rq.ctx.Fields
	if len(rq.ctx.Fields) > 0 {
		_spec.Unique = rq.ctx.Unique != nil && *rq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *ReviewQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(review.Table, review.Columns, sqlgraph.NewFieldSpec(review.FieldID, field.TypeString))
	_spec.From = rq.sql
	if unique := rq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rq.path != nil {
		_spec.Unique = true
	}
	if fields := rq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, review.FieldID)
		for i := range fields {
			if fields[i] != review.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *ReviewQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(review.Table)
	columns := rq.ctx.Fields
	if len(columns) == 0 {
		columns = review.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.ctx.Unique != nil && *rq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ReviewGroupBy is the group-by builder for Review entities.
type ReviewGroupBy struct {
	selector
	build *ReviewQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *ReviewGroupBy) Aggregate(fns ...AggregateFunc) *ReviewGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the selector query and scans the result into the given value.
func (rgb *ReviewGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgb.build.ctx, "GroupBy")
	if err := rgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ReviewQuery, *ReviewGroupBy](ctx, rgb.build, rgb, rgb.build.inters, v)
}

func (rgb *ReviewGroupBy) sqlScan(ctx context.Context, root *ReviewQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rgb.flds)+len(rgb.fns))
		for _, f := range *rgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ReviewSelect is the builder for selecting fields of Review entities.
type ReviewSelect struct {
	*ReviewQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *ReviewSelect) Aggregate(fns ...AggregateFunc) *ReviewSelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *ReviewSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rs.ctx, "Select")
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ReviewQuery, *ReviewSelect](ctx, rs.ReviewQuery, rs, rs.inters, v)
}

func (rs *ReviewSelect) sqlScan(ctx context.Context, root *ReviewQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
