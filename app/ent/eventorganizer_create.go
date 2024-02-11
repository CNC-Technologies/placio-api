// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/event"
	"placio-app/ent/eventorganizer"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EventOrganizerCreate is the builder for creating a EventOrganizer entity.
type EventOrganizerCreate struct {
	config
	mutation *EventOrganizerMutation
	hooks    []Hook
}

// SetOrganizerID sets the "organizerID" field.
func (eoc *EventOrganizerCreate) SetOrganizerID(s string) *EventOrganizerCreate {
	eoc.mutation.SetOrganizerID(s)
	return eoc
}

// SetOrganizerType sets the "organizerType" field.
func (eoc *EventOrganizerCreate) SetOrganizerType(s string) *EventOrganizerCreate {
	eoc.mutation.SetOrganizerType(s)
	return eoc
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (eoc *EventOrganizerCreate) SetEventID(id string) *EventOrganizerCreate {
	eoc.mutation.SetEventID(id)
	return eoc
}

// SetEvent sets the "event" edge to the Event entity.
func (eoc *EventOrganizerCreate) SetEvent(e *Event) *EventOrganizerCreate {
	return eoc.SetEventID(e.ID)
}

// Mutation returns the EventOrganizerMutation object of the builder.
func (eoc *EventOrganizerCreate) Mutation() *EventOrganizerMutation {
	return eoc.mutation
}

// Save creates the EventOrganizer in the database.
func (eoc *EventOrganizerCreate) Save(ctx context.Context) (*EventOrganizer, error) {
	return withHooks(ctx, eoc.sqlSave, eoc.mutation, eoc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (eoc *EventOrganizerCreate) SaveX(ctx context.Context) *EventOrganizer {
	v, err := eoc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eoc *EventOrganizerCreate) Exec(ctx context.Context) error {
	_, err := eoc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eoc *EventOrganizerCreate) ExecX(ctx context.Context) {
	if err := eoc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eoc *EventOrganizerCreate) check() error {
	if _, ok := eoc.mutation.OrganizerID(); !ok {
		return &ValidationError{Name: "organizerID", err: errors.New(`ent: missing required field "EventOrganizer.organizerID"`)}
	}
	if v, ok := eoc.mutation.OrganizerID(); ok {
		if err := eventorganizer.OrganizerIDValidator(v); err != nil {
			return &ValidationError{Name: "organizerID", err: fmt.Errorf(`ent: validator failed for field "EventOrganizer.organizerID": %w`, err)}
		}
	}
	if _, ok := eoc.mutation.OrganizerType(); !ok {
		return &ValidationError{Name: "organizerType", err: errors.New(`ent: missing required field "EventOrganizer.organizerType"`)}
	}
	if v, ok := eoc.mutation.OrganizerType(); ok {
		if err := eventorganizer.OrganizerTypeValidator(v); err != nil {
			return &ValidationError{Name: "organizerType", err: fmt.Errorf(`ent: validator failed for field "EventOrganizer.organizerType": %w`, err)}
		}
	}
	if _, ok := eoc.mutation.EventID(); !ok {
		return &ValidationError{Name: "event", err: errors.New(`ent: missing required edge "EventOrganizer.event"`)}
	}
	return nil
}

func (eoc *EventOrganizerCreate) sqlSave(ctx context.Context) (*EventOrganizer, error) {
	if err := eoc.check(); err != nil {
		return nil, err
	}
	_node, _spec := eoc.createSpec()
	if err := sqlgraph.CreateNode(ctx, eoc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected EventOrganizer.ID type: %T", _spec.ID.Value)
		}
	}
	eoc.mutation.id = &_node.ID
	eoc.mutation.done = true
	return _node, nil
}

func (eoc *EventOrganizerCreate) createSpec() (*EventOrganizer, *sqlgraph.CreateSpec) {
	var (
		_node = &EventOrganizer{config: eoc.config}
		_spec = sqlgraph.NewCreateSpec(eventorganizer.Table, sqlgraph.NewFieldSpec(eventorganizer.FieldID, field.TypeString))
	)
	if value, ok := eoc.mutation.OrganizerID(); ok {
		_spec.SetField(eventorganizer.FieldOrganizerID, field.TypeString, value)
		_node.OrganizerID = value
	}
	if value, ok := eoc.mutation.OrganizerType(); ok {
		_spec.SetField(eventorganizer.FieldOrganizerType, field.TypeString, value)
		_node.OrganizerType = value
	}
	if nodes := eoc.mutation.EventIDs(); len(nodes) > 0 {
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
		_node.event_event_organizers = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EventOrganizerCreateBulk is the builder for creating many EventOrganizer entities in bulk.
type EventOrganizerCreateBulk struct {
	config
	err      error
	builders []*EventOrganizerCreate
}

// Save creates the EventOrganizer entities in the database.
func (eocb *EventOrganizerCreateBulk) Save(ctx context.Context) ([]*EventOrganizer, error) {
	if eocb.err != nil {
		return nil, eocb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(eocb.builders))
	nodes := make([]*EventOrganizer, len(eocb.builders))
	mutators := make([]Mutator, len(eocb.builders))
	for i := range eocb.builders {
		func(i int, root context.Context) {
			builder := eocb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EventOrganizerMutation)
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
					_, err = mutators[i+1].Mutate(root, eocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, eocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, eocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (eocb *EventOrganizerCreateBulk) SaveX(ctx context.Context) []*EventOrganizer {
	v, err := eocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eocb *EventOrganizerCreateBulk) Exec(ctx context.Context) error {
	_, err := eocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eocb *EventOrganizerCreateBulk) ExecX(ctx context.Context) {
	if err := eocb.Exec(ctx); err != nil {
		panic(err)
	}
}