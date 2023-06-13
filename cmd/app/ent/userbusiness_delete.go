// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"placio-app/ent/predicate"
	"placio-app/ent/userbusiness"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserBusinessDelete is the builder for deleting a UserBusiness entity.
type UserBusinessDelete struct {
	config
	hooks    []Hook
	mutation *UserBusinessMutation
}

// Where appends a list predicates to the UserBusinessDelete builder.
func (ubd *UserBusinessDelete) Where(ps ...predicate.UserBusiness) *UserBusinessDelete {
	ubd.mutation.Where(ps...)
	return ubd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ubd *UserBusinessDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ubd.sqlExec, ubd.mutation, ubd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ubd *UserBusinessDelete) ExecX(ctx context.Context) int {
	n, err := ubd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ubd *UserBusinessDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(userbusiness.Table, sqlgraph.NewFieldSpec(userbusiness.FieldID, field.TypeString))
	if ps := ubd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ubd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ubd.mutation.done = true
	return affected, err
}

// UserBusinessDeleteOne is the builder for deleting a single UserBusiness entity.
type UserBusinessDeleteOne struct {
	ubd *UserBusinessDelete
}

// Where appends a list predicates to the UserBusinessDelete builder.
func (ubdo *UserBusinessDeleteOne) Where(ps ...predicate.UserBusiness) *UserBusinessDeleteOne {
	ubdo.ubd.mutation.Where(ps...)
	return ubdo
}

// Exec executes the deletion query.
func (ubdo *UserBusinessDeleteOne) Exec(ctx context.Context) error {
	n, err := ubdo.ubd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userbusiness.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ubdo *UserBusinessDeleteOne) ExecX(ctx context.Context) {
	if err := ubdo.Exec(ctx); err != nil {
		panic(err)
	}
}
