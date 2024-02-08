// Code generated by ent, DO NOT EDIT.

package memorelation

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/usememos/memos/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldLTE(FieldID, id))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldEQ(FieldType, v))
}

// MemoID applies equality check predicate on the "memo_id" field. It's identical to MemoIDEQ.
func MemoID(v int) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldEQ(FieldMemoID, v))
}

// RelatedMemoID applies equality check predicate on the "related_memo_id" field. It's identical to RelatedMemoIDEQ.
func RelatedMemoID(v int) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldEQ(FieldRelatedMemoID, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldContainsFold(FieldType, v))
}

// MemoIDEQ applies the EQ predicate on the "memo_id" field.
func MemoIDEQ(v int) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldEQ(FieldMemoID, v))
}

// MemoIDNEQ applies the NEQ predicate on the "memo_id" field.
func MemoIDNEQ(v int) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldNEQ(FieldMemoID, v))
}

// MemoIDIn applies the In predicate on the "memo_id" field.
func MemoIDIn(vs ...int) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldIn(FieldMemoID, vs...))
}

// MemoIDNotIn applies the NotIn predicate on the "memo_id" field.
func MemoIDNotIn(vs ...int) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldNotIn(FieldMemoID, vs...))
}

// RelatedMemoIDEQ applies the EQ predicate on the "related_memo_id" field.
func RelatedMemoIDEQ(v int) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldEQ(FieldRelatedMemoID, v))
}

// RelatedMemoIDNEQ applies the NEQ predicate on the "related_memo_id" field.
func RelatedMemoIDNEQ(v int) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldNEQ(FieldRelatedMemoID, v))
}

// RelatedMemoIDIn applies the In predicate on the "related_memo_id" field.
func RelatedMemoIDIn(vs ...int) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldIn(FieldRelatedMemoID, vs...))
}

// RelatedMemoIDNotIn applies the NotIn predicate on the "related_memo_id" field.
func RelatedMemoIDNotIn(vs ...int) predicate.MemoRelation {
	return predicate.MemoRelation(sql.FieldNotIn(FieldRelatedMemoID, vs...))
}

// HasMemo applies the HasEdge predicate on the "memo" edge.
func HasMemo() predicate.MemoRelation {
	return predicate.MemoRelation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, MemoTable, MemoColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMemoWith applies the HasEdge predicate on the "memo" edge with a given conditions (other predicates).
func HasMemoWith(preds ...predicate.Memo) predicate.MemoRelation {
	return predicate.MemoRelation(func(s *sql.Selector) {
		step := newMemoStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRelatedMemo applies the HasEdge predicate on the "related_memo" edge.
func HasRelatedMemo() predicate.MemoRelation {
	return predicate.MemoRelation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, RelatedMemoTable, RelatedMemoColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRelatedMemoWith applies the HasEdge predicate on the "related_memo" edge with a given conditions (other predicates).
func HasRelatedMemoWith(preds ...predicate.Memo) predicate.MemoRelation {
	return predicate.MemoRelation(func(s *sql.Selector) {
		step := newRelatedMemoStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MemoRelation) predicate.MemoRelation {
	return predicate.MemoRelation(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MemoRelation) predicate.MemoRelation {
	return predicate.MemoRelation(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.MemoRelation) predicate.MemoRelation {
	return predicate.MemoRelation(sql.NotPredicates(p))
}
