// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/usememos/memos/ent/memorelation"
	"github.com/usememos/memos/ent/predicate"
)

// MemoRelationDelete is the builder for deleting a MemoRelation entity.
type MemoRelationDelete struct {
	config
	hooks    []Hook
	mutation *MemoRelationMutation
}

// Where appends a list predicates to the MemoRelationDelete builder.
func (mrd *MemoRelationDelete) Where(ps ...predicate.MemoRelation) *MemoRelationDelete {
	mrd.mutation.Where(ps...)
	return mrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mrd *MemoRelationDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, mrd.sqlExec, mrd.mutation, mrd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (mrd *MemoRelationDelete) ExecX(ctx context.Context) int {
	n, err := mrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mrd *MemoRelationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(memorelation.Table, sqlgraph.NewFieldSpec(memorelation.FieldID, field.TypeInt))
	if ps := mrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, mrd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	mrd.mutation.done = true
	return affected, err
}

// MemoRelationDeleteOne is the builder for deleting a single MemoRelation entity.
type MemoRelationDeleteOne struct {
	mrd *MemoRelationDelete
}

// Where appends a list predicates to the MemoRelationDelete builder.
func (mrdo *MemoRelationDeleteOne) Where(ps ...predicate.MemoRelation) *MemoRelationDeleteOne {
	mrdo.mrd.mutation.Where(ps...)
	return mrdo
}

// Exec executes the deletion query.
func (mrdo *MemoRelationDeleteOne) Exec(ctx context.Context) error {
	n, err := mrdo.mrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{memorelation.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mrdo *MemoRelationDeleteOne) ExecX(ctx context.Context) {
	if err := mrdo.Exec(ctx); err != nil {
		panic(err)
	}
}
