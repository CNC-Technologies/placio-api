// Code generated by ent, DO NOT EDIT.

package accountsettings

import (
	"placio-app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.AccountSettings {
	return predicate.AccountSettings(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.AccountSettings {
	return predicate.AccountSettings(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.AccountSettings {
	return predicate.AccountSettings(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.AccountSettings {
	return predicate.AccountSettings(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.AccountSettings {
	return predicate.AccountSettings(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.AccountSettings {
	return predicate.AccountSettings(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.AccountSettings {
	return predicate.AccountSettings(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.AccountSettings {
	return predicate.AccountSettings(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.AccountSettings {
	return predicate.AccountSettings(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.AccountSettings {
	return predicate.AccountSettings(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.AccountSettings {
	return predicate.AccountSettings(sql.FieldContainsFold(FieldID, id))
}

// TwoFactorAuthentication applies equality check predicate on the "TwoFactorAuthentication" field. It's identical to TwoFactorAuthenticationEQ.
func TwoFactorAuthentication(v bool) predicate.AccountSettings {
	return predicate.AccountSettings(sql.FieldEQ(FieldTwoFactorAuthentication, v))
}

// TwoFactorAuthenticationEQ applies the EQ predicate on the "TwoFactorAuthentication" field.
func TwoFactorAuthenticationEQ(v bool) predicate.AccountSettings {
	return predicate.AccountSettings(sql.FieldEQ(FieldTwoFactorAuthentication, v))
}

// TwoFactorAuthenticationNEQ applies the NEQ predicate on the "TwoFactorAuthentication" field.
func TwoFactorAuthenticationNEQ(v bool) predicate.AccountSettings {
	return predicate.AccountSettings(sql.FieldNEQ(FieldTwoFactorAuthentication, v))
}

// HasBusinessAccount applies the HasEdge predicate on the "business_account" edge.
func HasBusinessAccount() predicate.AccountSettings {
	return predicate.AccountSettings(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, BusinessAccountTable, BusinessAccountColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessAccountWith applies the HasEdge predicate on the "business_account" edge with a given conditions (other predicates).
func HasBusinessAccountWith(preds ...predicate.Business) predicate.AccountSettings {
	return predicate.AccountSettings(func(s *sql.Selector) {
		step := newBusinessAccountStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AccountSettings) predicate.AccountSettings {
	return predicate.AccountSettings(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AccountSettings) predicate.AccountSettings {
	return predicate.AccountSettings(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AccountSettings) predicate.AccountSettings {
	return predicate.AccountSettings(sql.NotPredicates(p))
}
