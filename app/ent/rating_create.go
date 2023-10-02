// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/business"
	"placio-app/ent/event"
	"placio-app/ent/place"
	"placio-app/ent/rating"
	"placio-app/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RatingCreate is the builder for creating a Rating entity.
type RatingCreate struct {
	config
	mutation *RatingMutation
	hooks    []Hook
}

// SetScore sets the "score" field.
func (rc *RatingCreate) SetScore(i int) *RatingCreate {
	rc.mutation.SetScore(i)
	return rc
}

// SetReview sets the "review" field.
func (rc *RatingCreate) SetReview(s string) *RatingCreate {
	rc.mutation.SetReview(s)
	return rc
}

// SetNillableReview sets the "review" field if the given value is not nil.
func (rc *RatingCreate) SetNillableReview(s *string) *RatingCreate {
	if s != nil {
		rc.SetReview(*s)
	}
	return rc
}

// SetRatedAt sets the "ratedAt" field.
func (rc *RatingCreate) SetRatedAt(t time.Time) *RatingCreate {
	rc.mutation.SetRatedAt(t)
	return rc
}

// SetNillableRatedAt sets the "ratedAt" field if the given value is not nil.
func (rc *RatingCreate) SetNillableRatedAt(t *time.Time) *RatingCreate {
	if t != nil {
		rc.SetRatedAt(*t)
	}
	return rc
}

// SetID sets the "id" field.
func (rc *RatingCreate) SetID(s string) *RatingCreate {
	rc.mutation.SetID(s)
	return rc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (rc *RatingCreate) SetUserID(id string) *RatingCreate {
	rc.mutation.SetUserID(id)
	return rc
}

// SetUser sets the "user" edge to the User entity.
func (rc *RatingCreate) SetUser(u *User) *RatingCreate {
	return rc.SetUserID(u.ID)
}

// SetBusinessID sets the "business" edge to the Business entity by ID.
func (rc *RatingCreate) SetBusinessID(id string) *RatingCreate {
	rc.mutation.SetBusinessID(id)
	return rc
}

// SetNillableBusinessID sets the "business" edge to the Business entity by ID if the given value is not nil.
func (rc *RatingCreate) SetNillableBusinessID(id *string) *RatingCreate {
	if id != nil {
		rc = rc.SetBusinessID(*id)
	}
	return rc
}

// SetBusiness sets the "business" edge to the Business entity.
func (rc *RatingCreate) SetBusiness(b *Business) *RatingCreate {
	return rc.SetBusinessID(b.ID)
}

// SetPlaceID sets the "place" edge to the Place entity by ID.
func (rc *RatingCreate) SetPlaceID(id string) *RatingCreate {
	rc.mutation.SetPlaceID(id)
	return rc
}

// SetNillablePlaceID sets the "place" edge to the Place entity by ID if the given value is not nil.
func (rc *RatingCreate) SetNillablePlaceID(id *string) *RatingCreate {
	if id != nil {
		rc = rc.SetPlaceID(*id)
	}
	return rc
}

// SetPlace sets the "place" edge to the Place entity.
func (rc *RatingCreate) SetPlace(p *Place) *RatingCreate {
	return rc.SetPlaceID(p.ID)
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (rc *RatingCreate) SetEventID(id string) *RatingCreate {
	rc.mutation.SetEventID(id)
	return rc
}

// SetNillableEventID sets the "event" edge to the Event entity by ID if the given value is not nil.
func (rc *RatingCreate) SetNillableEventID(id *string) *RatingCreate {
	if id != nil {
		rc = rc.SetEventID(*id)
	}
	return rc
}

// SetEvent sets the "event" edge to the Event entity.
func (rc *RatingCreate) SetEvent(e *Event) *RatingCreate {
	return rc.SetEventID(e.ID)
}

// Mutation returns the RatingMutation object of the builder.
func (rc *RatingCreate) Mutation() *RatingMutation {
	return rc.mutation
}

// Save creates the Rating in the database.
func (rc *RatingCreate) Save(ctx context.Context) (*Rating, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RatingCreate) SaveX(ctx context.Context) *Rating {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RatingCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RatingCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RatingCreate) defaults() {
	if _, ok := rc.mutation.RatedAt(); !ok {
		v := rating.DefaultRatedAt()
		rc.mutation.SetRatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RatingCreate) check() error {
	if _, ok := rc.mutation.Score(); !ok {
		return &ValidationError{Name: "score", err: errors.New(`ent: missing required field "Rating.score"`)}
	}
	if v, ok := rc.mutation.Score(); ok {
		if err := rating.ScoreValidator(v); err != nil {
			return &ValidationError{Name: "score", err: fmt.Errorf(`ent: validator failed for field "Rating.score": %w`, err)}
		}
	}
	if _, ok := rc.mutation.RatedAt(); !ok {
		return &ValidationError{Name: "ratedAt", err: errors.New(`ent: missing required field "Rating.ratedAt"`)}
	}
	if _, ok := rc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Rating.user"`)}
	}
	return nil
}

func (rc *RatingCreate) sqlSave(ctx context.Context) (*Rating, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Rating.ID type: %T", _spec.ID.Value)
		}
	}
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RatingCreate) createSpec() (*Rating, *sqlgraph.CreateSpec) {
	var (
		_node = &Rating{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(rating.Table, sqlgraph.NewFieldSpec(rating.FieldID, field.TypeString))
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.Score(); ok {
		_spec.SetField(rating.FieldScore, field.TypeInt, value)
		_node.Score = value
	}
	if value, ok := rc.mutation.Review(); ok {
		_spec.SetField(rating.FieldReview, field.TypeString, value)
		_node.Review = value
	}
	if value, ok := rc.mutation.RatedAt(); ok {
		_spec.SetField(rating.FieldRatedAt, field.TypeTime, value)
		_node.RatedAt = value
	}
	if nodes := rc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rating.UserTable,
			Columns: []string{rating.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_ratings = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.BusinessIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   rating.BusinessTable,
			Columns: []string{rating.BusinessColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.rating_business = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.PlaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   rating.PlaceTable,
			Columns: []string{rating.PlaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(place.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.rating_place = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   rating.EventTable,
			Columns: []string{rating.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.rating_event = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RatingCreateBulk is the builder for creating many Rating entities in bulk.
type RatingCreateBulk struct {
	config
	err      error
	builders []*RatingCreate
}

// Save creates the Rating entities in the database.
func (rcb *RatingCreateBulk) Save(ctx context.Context) ([]*Rating, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Rating, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RatingMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RatingCreateBulk) SaveX(ctx context.Context) []*Rating {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RatingCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RatingCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
