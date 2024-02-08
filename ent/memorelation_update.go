// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/usememos/memos/ent/memo"
	"github.com/usememos/memos/ent/memorelation"
	"github.com/usememos/memos/ent/predicate"
)

// MemoRelationUpdate is the builder for updating MemoRelation entities.
type MemoRelationUpdate struct {
	config
	hooks    []Hook
	mutation *MemoRelationMutation
}

// Where appends a list predicates to the MemoRelationUpdate builder.
func (mru *MemoRelationUpdate) Where(ps ...predicate.MemoRelation) *MemoRelationUpdate {
	mru.mutation.Where(ps...)
	return mru
}

// SetType sets the "type" field.
func (mru *MemoRelationUpdate) SetType(s string) *MemoRelationUpdate {
	mru.mutation.SetType(s)
	return mru
}

// SetNillableType sets the "type" field if the given value is not nil.
func (mru *MemoRelationUpdate) SetNillableType(s *string) *MemoRelationUpdate {
	if s != nil {
		mru.SetType(*s)
	}
	return mru
}

// SetMemoID sets the "memo_id" field.
func (mru *MemoRelationUpdate) SetMemoID(i int) *MemoRelationUpdate {
	mru.mutation.SetMemoID(i)
	return mru
}

// SetNillableMemoID sets the "memo_id" field if the given value is not nil.
func (mru *MemoRelationUpdate) SetNillableMemoID(i *int) *MemoRelationUpdate {
	if i != nil {
		mru.SetMemoID(*i)
	}
	return mru
}

// SetRelatedMemoID sets the "related_memo_id" field.
func (mru *MemoRelationUpdate) SetRelatedMemoID(i int) *MemoRelationUpdate {
	mru.mutation.SetRelatedMemoID(i)
	return mru
}

// SetNillableRelatedMemoID sets the "related_memo_id" field if the given value is not nil.
func (mru *MemoRelationUpdate) SetNillableRelatedMemoID(i *int) *MemoRelationUpdate {
	if i != nil {
		mru.SetRelatedMemoID(*i)
	}
	return mru
}

// SetMemo sets the "memo" edge to the Memo entity.
func (mru *MemoRelationUpdate) SetMemo(m *Memo) *MemoRelationUpdate {
	return mru.SetMemoID(m.ID)
}

// SetRelatedMemo sets the "related_memo" edge to the Memo entity.
func (mru *MemoRelationUpdate) SetRelatedMemo(m *Memo) *MemoRelationUpdate {
	return mru.SetRelatedMemoID(m.ID)
}

// Mutation returns the MemoRelationMutation object of the builder.
func (mru *MemoRelationUpdate) Mutation() *MemoRelationMutation {
	return mru.mutation
}

// ClearMemo clears the "memo" edge to the Memo entity.
func (mru *MemoRelationUpdate) ClearMemo() *MemoRelationUpdate {
	mru.mutation.ClearMemo()
	return mru
}

// ClearRelatedMemo clears the "related_memo" edge to the Memo entity.
func (mru *MemoRelationUpdate) ClearRelatedMemo() *MemoRelationUpdate {
	mru.mutation.ClearRelatedMemo()
	return mru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mru *MemoRelationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mru.sqlSave, mru.mutation, mru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mru *MemoRelationUpdate) SaveX(ctx context.Context) int {
	affected, err := mru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mru *MemoRelationUpdate) Exec(ctx context.Context) error {
	_, err := mru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mru *MemoRelationUpdate) ExecX(ctx context.Context) {
	if err := mru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mru *MemoRelationUpdate) check() error {
	if _, ok := mru.mutation.MemoID(); mru.mutation.MemoCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "MemoRelation.memo"`)
	}
	if _, ok := mru.mutation.RelatedMemoID(); mru.mutation.RelatedMemoCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "MemoRelation.related_memo"`)
	}
	return nil
}

func (mru *MemoRelationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(memorelation.Table, memorelation.Columns, sqlgraph.NewFieldSpec(memorelation.FieldID, field.TypeInt))
	if ps := mru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mru.mutation.GetType(); ok {
		_spec.SetField(memorelation.FieldType, field.TypeString, value)
	}
	if mru.mutation.MemoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   memorelation.MemoTable,
			Columns: []string{memorelation.MemoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memo.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mru.mutation.MemoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   memorelation.MemoTable,
			Columns: []string{memorelation.MemoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mru.mutation.RelatedMemoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   memorelation.RelatedMemoTable,
			Columns: []string{memorelation.RelatedMemoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memo.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mru.mutation.RelatedMemoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   memorelation.RelatedMemoTable,
			Columns: []string{memorelation.RelatedMemoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{memorelation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mru.mutation.done = true
	return n, nil
}

// MemoRelationUpdateOne is the builder for updating a single MemoRelation entity.
type MemoRelationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MemoRelationMutation
}

// SetType sets the "type" field.
func (mruo *MemoRelationUpdateOne) SetType(s string) *MemoRelationUpdateOne {
	mruo.mutation.SetType(s)
	return mruo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (mruo *MemoRelationUpdateOne) SetNillableType(s *string) *MemoRelationUpdateOne {
	if s != nil {
		mruo.SetType(*s)
	}
	return mruo
}

// SetMemoID sets the "memo_id" field.
func (mruo *MemoRelationUpdateOne) SetMemoID(i int) *MemoRelationUpdateOne {
	mruo.mutation.SetMemoID(i)
	return mruo
}

// SetNillableMemoID sets the "memo_id" field if the given value is not nil.
func (mruo *MemoRelationUpdateOne) SetNillableMemoID(i *int) *MemoRelationUpdateOne {
	if i != nil {
		mruo.SetMemoID(*i)
	}
	return mruo
}

// SetRelatedMemoID sets the "related_memo_id" field.
func (mruo *MemoRelationUpdateOne) SetRelatedMemoID(i int) *MemoRelationUpdateOne {
	mruo.mutation.SetRelatedMemoID(i)
	return mruo
}

// SetNillableRelatedMemoID sets the "related_memo_id" field if the given value is not nil.
func (mruo *MemoRelationUpdateOne) SetNillableRelatedMemoID(i *int) *MemoRelationUpdateOne {
	if i != nil {
		mruo.SetRelatedMemoID(*i)
	}
	return mruo
}

// SetMemo sets the "memo" edge to the Memo entity.
func (mruo *MemoRelationUpdateOne) SetMemo(m *Memo) *MemoRelationUpdateOne {
	return mruo.SetMemoID(m.ID)
}

// SetRelatedMemo sets the "related_memo" edge to the Memo entity.
func (mruo *MemoRelationUpdateOne) SetRelatedMemo(m *Memo) *MemoRelationUpdateOne {
	return mruo.SetRelatedMemoID(m.ID)
}

// Mutation returns the MemoRelationMutation object of the builder.
func (mruo *MemoRelationUpdateOne) Mutation() *MemoRelationMutation {
	return mruo.mutation
}

// ClearMemo clears the "memo" edge to the Memo entity.
func (mruo *MemoRelationUpdateOne) ClearMemo() *MemoRelationUpdateOne {
	mruo.mutation.ClearMemo()
	return mruo
}

// ClearRelatedMemo clears the "related_memo" edge to the Memo entity.
func (mruo *MemoRelationUpdateOne) ClearRelatedMemo() *MemoRelationUpdateOne {
	mruo.mutation.ClearRelatedMemo()
	return mruo
}

// Where appends a list predicates to the MemoRelationUpdate builder.
func (mruo *MemoRelationUpdateOne) Where(ps ...predicate.MemoRelation) *MemoRelationUpdateOne {
	mruo.mutation.Where(ps...)
	return mruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mruo *MemoRelationUpdateOne) Select(field string, fields ...string) *MemoRelationUpdateOne {
	mruo.fields = append([]string{field}, fields...)
	return mruo
}

// Save executes the query and returns the updated MemoRelation entity.
func (mruo *MemoRelationUpdateOne) Save(ctx context.Context) (*MemoRelation, error) {
	return withHooks(ctx, mruo.sqlSave, mruo.mutation, mruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mruo *MemoRelationUpdateOne) SaveX(ctx context.Context) *MemoRelation {
	node, err := mruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mruo *MemoRelationUpdateOne) Exec(ctx context.Context) error {
	_, err := mruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mruo *MemoRelationUpdateOne) ExecX(ctx context.Context) {
	if err := mruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mruo *MemoRelationUpdateOne) check() error {
	if _, ok := mruo.mutation.MemoID(); mruo.mutation.MemoCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "MemoRelation.memo"`)
	}
	if _, ok := mruo.mutation.RelatedMemoID(); mruo.mutation.RelatedMemoCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "MemoRelation.related_memo"`)
	}
	return nil
}

func (mruo *MemoRelationUpdateOne) sqlSave(ctx context.Context) (_node *MemoRelation, err error) {
	if err := mruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(memorelation.Table, memorelation.Columns, sqlgraph.NewFieldSpec(memorelation.FieldID, field.TypeInt))
	id, ok := mruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MemoRelation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, memorelation.FieldID)
		for _, f := range fields {
			if !memorelation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != memorelation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mruo.mutation.GetType(); ok {
		_spec.SetField(memorelation.FieldType, field.TypeString, value)
	}
	if mruo.mutation.MemoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   memorelation.MemoTable,
			Columns: []string{memorelation.MemoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memo.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mruo.mutation.MemoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   memorelation.MemoTable,
			Columns: []string{memorelation.MemoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mruo.mutation.RelatedMemoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   memorelation.RelatedMemoTable,
			Columns: []string{memorelation.RelatedMemoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memo.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mruo.mutation.RelatedMemoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   memorelation.RelatedMemoTable,
			Columns: []string{memorelation.RelatedMemoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &MemoRelation{config: mruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{memorelation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mruo.mutation.done = true
	return _node, nil
}
