// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"placio-app/ent/faq"
	"placio-app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FAQDelete is the builder for deleting a FAQ entity.
type FAQDelete struct {
	config
	hooks    []Hook
	mutation *FAQMutation
}

// Where appends a list predicates to the FAQDelete builder.
func (fd *FAQDelete) Where(ps ...predicate.FAQ) *FAQDelete {
	fd.mutation.Where(ps...)
	return fd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fd *FAQDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, fd.sqlExec, fd.mutation, fd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (fd *FAQDelete) ExecX(ctx context.Context) int {
	n, err := fd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fd *FAQDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(faq.Table, sqlgraph.NewFieldSpec(faq.FieldID, field.TypeString))
	if ps := fd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	fd.mutation.done = true
	return affected, err
}

// FAQDeleteOne is the builder for deleting a single FAQ entity.
type FAQDeleteOne struct {
	fd *FAQDelete
}

// Where appends a list predicates to the FAQDelete builder.
func (fdo *FAQDeleteOne) Where(ps ...predicate.FAQ) *FAQDeleteOne {
	fdo.fd.mutation.Where(ps...)
	return fdo
}

// Exec executes the deletion query.
func (fdo *FAQDeleteOne) Exec(ctx context.Context) error {
	n, err := fdo.fd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{faq.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fdo *FAQDeleteOne) ExecX(ctx context.Context) {
	if err := fdo.Exec(ctx); err != nil {
		panic(err)
	}
}
