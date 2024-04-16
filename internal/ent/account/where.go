// Code generated by ent, DO NOT EDIT.

package account

import (
	"frisbane/internal/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldUserID, v))
}

// DisableTime applies equality check predicate on the "disable_time" field. It's identical to DisableTimeEQ.
func DisableTime(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldDisableTime, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.Account {
	return predicate.Account(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldUserID, v))
}

// DisableTimeEQ applies the EQ predicate on the "disable_time" field.
func DisableTimeEQ(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldDisableTime, v))
}

// DisableTimeNEQ applies the NEQ predicate on the "disable_time" field.
func DisableTimeNEQ(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldDisableTime, v))
}

// DisableTimeIn applies the In predicate on the "disable_time" field.
func DisableTimeIn(vs ...time.Time) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldDisableTime, vs...))
}

// DisableTimeNotIn applies the NotIn predicate on the "disable_time" field.
func DisableTimeNotIn(vs ...time.Time) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldDisableTime, vs...))
}

// DisableTimeGT applies the GT predicate on the "disable_time" field.
func DisableTimeGT(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldDisableTime, v))
}

// DisableTimeGTE applies the GTE predicate on the "disable_time" field.
func DisableTimeGTE(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldDisableTime, v))
}

// DisableTimeLT applies the LT predicate on the "disable_time" field.
func DisableTimeLT(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldDisableTime, v))
}

// DisableTimeLTE applies the LTE predicate on the "disable_time" field.
func DisableTimeLTE(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldDisableTime, v))
}

// DisableTimeIsNil applies the IsNil predicate on the "disable_time" field.
func DisableTimeIsNil() predicate.Account {
	return predicate.Account(sql.FieldIsNull(FieldDisableTime))
}

// DisableTimeNotNil applies the NotNil predicate on the "disable_time" field.
func DisableTimeNotNil() predicate.Account {
	return predicate.Account(sql.FieldNotNull(FieldDisableTime))
}

// HasProfile applies the HasEdge predicate on the "profile" edge.
func HasProfile() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ProfileTable, ProfileColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProfileWith applies the HasEdge predicate on the "profile" edge with a given conditions (other predicates).
func HasProfileWith(preds ...predicate.Profile) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProfileInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ProfileTable, ProfileColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Account) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Account) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Account) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		p(s.Not())
	})
}
