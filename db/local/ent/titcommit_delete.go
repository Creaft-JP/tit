// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Creaft-JP/tit/db/local/ent/predicate"
	"github.com/Creaft-JP/tit/db/local/ent/titcommit"
)

// TitCommitDelete is the builder for deleting a TitCommit entity.
type TitCommitDelete struct {
	config
	hooks    []Hook
	mutation *TitCommitMutation
}

// Where appends a list predicates to the TitCommitDelete builder.
func (tcd *TitCommitDelete) Where(ps ...predicate.TitCommit) *TitCommitDelete {
	tcd.mutation.Where(ps...)
	return tcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tcd *TitCommitDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, tcd.sqlExec, tcd.mutation, tcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tcd *TitCommitDelete) ExecX(ctx context.Context) int {
	n, err := tcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tcd *TitCommitDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(titcommit.Table, sqlgraph.NewFieldSpec(titcommit.FieldID, field.TypeInt))
	if ps := tcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tcd.mutation.done = true
	return affected, err
}

// TitCommitDeleteOne is the builder for deleting a single TitCommit entity.
type TitCommitDeleteOne struct {
	tcd *TitCommitDelete
}

// Where appends a list predicates to the TitCommitDelete builder.
func (tcdo *TitCommitDeleteOne) Where(ps ...predicate.TitCommit) *TitCommitDeleteOne {
	tcdo.tcd.mutation.Where(ps...)
	return tcdo
}

// Exec executes the deletion query.
func (tcdo *TitCommitDeleteOne) Exec(ctx context.Context) error {
	n, err := tcdo.tcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{titcommit.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tcdo *TitCommitDeleteOne) ExecX(ctx context.Context) {
	if err := tcdo.Exec(ctx); err != nil {
		panic(err)
	}
}
