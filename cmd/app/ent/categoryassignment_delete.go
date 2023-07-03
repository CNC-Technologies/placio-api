// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"placio-app/ent/categoryassignment"
	"placio-app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CategoryAssignmentDelete is the builder for deleting a CategoryAssignment entity.
type CategoryAssignmentDelete struct {
	config
	hooks    []Hook
	mutation *CategoryAssignmentMutation
}

// Where appends a list predicates to the CategoryAssignmentDelete builder.
func (cad *CategoryAssignmentDelete) Where(ps ...predicate.CategoryAssignment) *CategoryAssignmentDelete {
	cad.mutation.Where(ps...)
	return cad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cad *CategoryAssignmentDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cad.sqlExec, cad.mutation, cad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cad *CategoryAssignmentDelete) ExecX(ctx context.Context) int {
	n, err := cad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cad *CategoryAssignmentDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(categoryassignment.Table, sqlgraph.NewFieldSpec(categoryassignment.FieldID, field.TypeString))
	if ps := cad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cad.mutation.done = true
	return affected, err
}

// CategoryAssignmentDeleteOne is the builder for deleting a single CategoryAssignment entity.
type CategoryAssignmentDeleteOne struct {
	cad *CategoryAssignmentDelete
}

// Where appends a list predicates to the CategoryAssignmentDelete builder.
func (cado *CategoryAssignmentDeleteOne) Where(ps ...predicate.CategoryAssignment) *CategoryAssignmentDeleteOne {
	cado.cad.mutation.Where(ps...)
	return cado
}

// Exec executes the deletion query.
func (cado *CategoryAssignmentDeleteOne) Exec(ctx context.Context) error {
	n, err := cado.cad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{categoryassignment.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cado *CategoryAssignmentDeleteOne) ExecX(ctx context.Context) {
	if err := cado.Exec(ctx); err != nil {
		panic(err)
	}
}