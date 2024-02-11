// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/event"
	"placio-app/ent/eventorganizer"
	"placio-app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EventOrganizerUpdate is the builder for updating EventOrganizer entities.
type EventOrganizerUpdate struct {
	config
	hooks    []Hook
	mutation *EventOrganizerMutation
}

// Where appends a list predicates to the EventOrganizerUpdate builder.
func (eou *EventOrganizerUpdate) Where(ps ...predicate.EventOrganizer) *EventOrganizerUpdate {
	eou.mutation.Where(ps...)
	return eou
}

// SetOrganizerID sets the "organizerID" field.
func (eou *EventOrganizerUpdate) SetOrganizerID(s string) *EventOrganizerUpdate {
	eou.mutation.SetOrganizerID(s)
	return eou
}

// SetNillableOrganizerID sets the "organizerID" field if the given value is not nil.
func (eou *EventOrganizerUpdate) SetNillableOrganizerID(s *string) *EventOrganizerUpdate {
	if s != nil {
		eou.SetOrganizerID(*s)
	}
	return eou
}

// SetOrganizerType sets the "organizerType" field.
func (eou *EventOrganizerUpdate) SetOrganizerType(s string) *EventOrganizerUpdate {
	eou.mutation.SetOrganizerType(s)
	return eou
}

// SetNillableOrganizerType sets the "organizerType" field if the given value is not nil.
func (eou *EventOrganizerUpdate) SetNillableOrganizerType(s *string) *EventOrganizerUpdate {
	if s != nil {
		eou.SetOrganizerType(*s)
	}
	return eou
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (eou *EventOrganizerUpdate) SetEventID(id string) *EventOrganizerUpdate {
	eou.mutation.SetEventID(id)
	return eou
}

// SetEvent sets the "event" edge to the Event entity.
func (eou *EventOrganizerUpdate) SetEvent(e *Event) *EventOrganizerUpdate {
	return eou.SetEventID(e.ID)
}

// Mutation returns the EventOrganizerMutation object of the builder.
func (eou *EventOrganizerUpdate) Mutation() *EventOrganizerMutation {
	return eou.mutation
}

// ClearEvent clears the "event" edge to the Event entity.
func (eou *EventOrganizerUpdate) ClearEvent() *EventOrganizerUpdate {
	eou.mutation.ClearEvent()
	return eou
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eou *EventOrganizerUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, eou.sqlSave, eou.mutation, eou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eou *EventOrganizerUpdate) SaveX(ctx context.Context) int {
	affected, err := eou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eou *EventOrganizerUpdate) Exec(ctx context.Context) error {
	_, err := eou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eou *EventOrganizerUpdate) ExecX(ctx context.Context) {
	if err := eou.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eou *EventOrganizerUpdate) check() error {
	if v, ok := eou.mutation.OrganizerID(); ok {
		if err := eventorganizer.OrganizerIDValidator(v); err != nil {
			return &ValidationError{Name: "organizerID", err: fmt.Errorf(`ent: validator failed for field "EventOrganizer.organizerID": %w`, err)}
		}
	}
	if v, ok := eou.mutation.OrganizerType(); ok {
		if err := eventorganizer.OrganizerTypeValidator(v); err != nil {
			return &ValidationError{Name: "organizerType", err: fmt.Errorf(`ent: validator failed for field "EventOrganizer.organizerType": %w`, err)}
		}
	}
	if _, ok := eou.mutation.EventID(); eou.mutation.EventCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "EventOrganizer.event"`)
	}
	return nil
}

func (eou *EventOrganizerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := eou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(eventorganizer.Table, eventorganizer.Columns, sqlgraph.NewFieldSpec(eventorganizer.FieldID, field.TypeString))
	if ps := eou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eou.mutation.OrganizerID(); ok {
		_spec.SetField(eventorganizer.FieldOrganizerID, field.TypeString, value)
	}
	if value, ok := eou.mutation.OrganizerType(); ok {
		_spec.SetField(eventorganizer.FieldOrganizerType, field.TypeString, value)
	}
	if eou.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventorganizer.EventTable,
			Columns: []string{eventorganizer.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eou.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventorganizer.EventTable,
			Columns: []string{eventorganizer.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{eventorganizer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eou.mutation.done = true
	return n, nil
}

// EventOrganizerUpdateOne is the builder for updating a single EventOrganizer entity.
type EventOrganizerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EventOrganizerMutation
}

// SetOrganizerID sets the "organizerID" field.
func (eouo *EventOrganizerUpdateOne) SetOrganizerID(s string) *EventOrganizerUpdateOne {
	eouo.mutation.SetOrganizerID(s)
	return eouo
}

// SetNillableOrganizerID sets the "organizerID" field if the given value is not nil.
func (eouo *EventOrganizerUpdateOne) SetNillableOrganizerID(s *string) *EventOrganizerUpdateOne {
	if s != nil {
		eouo.SetOrganizerID(*s)
	}
	return eouo
}

// SetOrganizerType sets the "organizerType" field.
func (eouo *EventOrganizerUpdateOne) SetOrganizerType(s string) *EventOrganizerUpdateOne {
	eouo.mutation.SetOrganizerType(s)
	return eouo
}

// SetNillableOrganizerType sets the "organizerType" field if the given value is not nil.
func (eouo *EventOrganizerUpdateOne) SetNillableOrganizerType(s *string) *EventOrganizerUpdateOne {
	if s != nil {
		eouo.SetOrganizerType(*s)
	}
	return eouo
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (eouo *EventOrganizerUpdateOne) SetEventID(id string) *EventOrganizerUpdateOne {
	eouo.mutation.SetEventID(id)
	return eouo
}

// SetEvent sets the "event" edge to the Event entity.
func (eouo *EventOrganizerUpdateOne) SetEvent(e *Event) *EventOrganizerUpdateOne {
	return eouo.SetEventID(e.ID)
}

// Mutation returns the EventOrganizerMutation object of the builder.
func (eouo *EventOrganizerUpdateOne) Mutation() *EventOrganizerMutation {
	return eouo.mutation
}

// ClearEvent clears the "event" edge to the Event entity.
func (eouo *EventOrganizerUpdateOne) ClearEvent() *EventOrganizerUpdateOne {
	eouo.mutation.ClearEvent()
	return eouo
}

// Where appends a list predicates to the EventOrganizerUpdate builder.
func (eouo *EventOrganizerUpdateOne) Where(ps ...predicate.EventOrganizer) *EventOrganizerUpdateOne {
	eouo.mutation.Where(ps...)
	return eouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (eouo *EventOrganizerUpdateOne) Select(field string, fields ...string) *EventOrganizerUpdateOne {
	eouo.fields = append([]string{field}, fields...)
	return eouo
}

// Save executes the query and returns the updated EventOrganizer entity.
func (eouo *EventOrganizerUpdateOne) Save(ctx context.Context) (*EventOrganizer, error) {
	return withHooks(ctx, eouo.sqlSave, eouo.mutation, eouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eouo *EventOrganizerUpdateOne) SaveX(ctx context.Context) *EventOrganizer {
	node, err := eouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (eouo *EventOrganizerUpdateOne) Exec(ctx context.Context) error {
	_, err := eouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eouo *EventOrganizerUpdateOne) ExecX(ctx context.Context) {
	if err := eouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eouo *EventOrganizerUpdateOne) check() error {
	if v, ok := eouo.mutation.OrganizerID(); ok {
		if err := eventorganizer.OrganizerIDValidator(v); err != nil {
			return &ValidationError{Name: "organizerID", err: fmt.Errorf(`ent: validator failed for field "EventOrganizer.organizerID": %w`, err)}
		}
	}
	if v, ok := eouo.mutation.OrganizerType(); ok {
		if err := eventorganizer.OrganizerTypeValidator(v); err != nil {
			return &ValidationError{Name: "organizerType", err: fmt.Errorf(`ent: validator failed for field "EventOrganizer.organizerType": %w`, err)}
		}
	}
	if _, ok := eouo.mutation.EventID(); eouo.mutation.EventCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "EventOrganizer.event"`)
	}
	return nil
}

func (eouo *EventOrganizerUpdateOne) sqlSave(ctx context.Context) (_node *EventOrganizer, err error) {
	if err := eouo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(eventorganizer.Table, eventorganizer.Columns, sqlgraph.NewFieldSpec(eventorganizer.FieldID, field.TypeString))
	id, ok := eouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EventOrganizer.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := eouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, eventorganizer.FieldID)
		for _, f := range fields {
			if !eventorganizer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != eventorganizer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := eouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eouo.mutation.OrganizerID(); ok {
		_spec.SetField(eventorganizer.FieldOrganizerID, field.TypeString, value)
	}
	if value, ok := eouo.mutation.OrganizerType(); ok {
		_spec.SetField(eventorganizer.FieldOrganizerType, field.TypeString, value)
	}
	if eouo.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventorganizer.EventTable,
			Columns: []string{eventorganizer.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eouo.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventorganizer.EventTable,
			Columns: []string{eventorganizer.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &EventOrganizer{config: eouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, eouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{eventorganizer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	eouo.mutation.done = true
	return _node, nil
}