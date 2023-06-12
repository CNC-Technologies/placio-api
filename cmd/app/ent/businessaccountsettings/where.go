// Code generated by ent, DO NOT EDIT.

package businessaccountsettings

import (
	"placio-app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldLTE(FieldID, id))
}

// BusinessAccountSettingsID applies equality check predicate on the "BusinessAccountSettingsID" field. It's identical to BusinessAccountSettingsIDEQ.
func BusinessAccountSettingsID(v string) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldEQ(FieldBusinessAccountSettingsID, v))
}

// BusinessAccountID applies equality check predicate on the "BusinessAccountID" field. It's identical to BusinessAccountIDEQ.
func BusinessAccountID(v string) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldEQ(FieldBusinessAccountID, v))
}

// TwoFactorAuthentication applies equality check predicate on the "TwoFactorAuthentication" field. It's identical to TwoFactorAuthenticationEQ.
func TwoFactorAuthentication(v bool) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldEQ(FieldTwoFactorAuthentication, v))
}

// BusinessAccountSettingsIDEQ applies the EQ predicate on the "BusinessAccountSettingsID" field.
func BusinessAccountSettingsIDEQ(v string) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldEQ(FieldBusinessAccountSettingsID, v))
}

// BusinessAccountSettingsIDNEQ applies the NEQ predicate on the "BusinessAccountSettingsID" field.
func BusinessAccountSettingsIDNEQ(v string) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldNEQ(FieldBusinessAccountSettingsID, v))
}

// BusinessAccountSettingsIDIn applies the In predicate on the "BusinessAccountSettingsID" field.
func BusinessAccountSettingsIDIn(vs ...string) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldIn(FieldBusinessAccountSettingsID, vs...))
}

// BusinessAccountSettingsIDNotIn applies the NotIn predicate on the "BusinessAccountSettingsID" field.
func BusinessAccountSettingsIDNotIn(vs ...string) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldNotIn(FieldBusinessAccountSettingsID, vs...))
}

// BusinessAccountSettingsIDGT applies the GT predicate on the "BusinessAccountSettingsID" field.
func BusinessAccountSettingsIDGT(v string) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldGT(FieldBusinessAccountSettingsID, v))
}

// BusinessAccountSettingsIDGTE applies the GTE predicate on the "BusinessAccountSettingsID" field.
func BusinessAccountSettingsIDGTE(v string) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldGTE(FieldBusinessAccountSettingsID, v))
}

// BusinessAccountSettingsIDLT applies the LT predicate on the "BusinessAccountSettingsID" field.
func BusinessAccountSettingsIDLT(v string) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldLT(FieldBusinessAccountSettingsID, v))
}

// BusinessAccountSettingsIDLTE applies the LTE predicate on the "BusinessAccountSettingsID" field.
func BusinessAccountSettingsIDLTE(v string) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldLTE(FieldBusinessAccountSettingsID, v))
}

// BusinessAccountSettingsIDContains applies the Contains predicate on the "BusinessAccountSettingsID" field.
func BusinessAccountSettingsIDContains(v string) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldContains(FieldBusinessAccountSettingsID, v))
}

// BusinessAccountSettingsIDHasPrefix applies the HasPrefix predicate on the "BusinessAccountSettingsID" field.
func BusinessAccountSettingsIDHasPrefix(v string) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldHasPrefix(FieldBusinessAccountSettingsID, v))
}

// BusinessAccountSettingsIDHasSuffix applies the HasSuffix predicate on the "BusinessAccountSettingsID" field.
func BusinessAccountSettingsIDHasSuffix(v string) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldHasSuffix(FieldBusinessAccountSettingsID, v))
}

// BusinessAccountSettingsIDEqualFold applies the EqualFold predicate on the "BusinessAccountSettingsID" field.
func BusinessAccountSettingsIDEqualFold(v string) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldEqualFold(FieldBusinessAccountSettingsID, v))
}

// BusinessAccountSettingsIDContainsFold applies the ContainsFold predicate on the "BusinessAccountSettingsID" field.
func BusinessAccountSettingsIDContainsFold(v string) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldContainsFold(FieldBusinessAccountSettingsID, v))
}

// BusinessAccountIDEQ applies the EQ predicate on the "BusinessAccountID" field.
func BusinessAccountIDEQ(v string) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldEQ(FieldBusinessAccountID, v))
}

// BusinessAccountIDNEQ applies the NEQ predicate on the "BusinessAccountID" field.
func BusinessAccountIDNEQ(v string) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldNEQ(FieldBusinessAccountID, v))
}

// BusinessAccountIDIn applies the In predicate on the "BusinessAccountID" field.
func BusinessAccountIDIn(vs ...string) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldIn(FieldBusinessAccountID, vs...))
}

// BusinessAccountIDNotIn applies the NotIn predicate on the "BusinessAccountID" field.
func BusinessAccountIDNotIn(vs ...string) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldNotIn(FieldBusinessAccountID, vs...))
}

// BusinessAccountIDGT applies the GT predicate on the "BusinessAccountID" field.
func BusinessAccountIDGT(v string) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldGT(FieldBusinessAccountID, v))
}

// BusinessAccountIDGTE applies the GTE predicate on the "BusinessAccountID" field.
func BusinessAccountIDGTE(v string) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldGTE(FieldBusinessAccountID, v))
}

// BusinessAccountIDLT applies the LT predicate on the "BusinessAccountID" field.
func BusinessAccountIDLT(v string) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldLT(FieldBusinessAccountID, v))
}

// BusinessAccountIDLTE applies the LTE predicate on the "BusinessAccountID" field.
func BusinessAccountIDLTE(v string) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldLTE(FieldBusinessAccountID, v))
}

// BusinessAccountIDContains applies the Contains predicate on the "BusinessAccountID" field.
func BusinessAccountIDContains(v string) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldContains(FieldBusinessAccountID, v))
}

// BusinessAccountIDHasPrefix applies the HasPrefix predicate on the "BusinessAccountID" field.
func BusinessAccountIDHasPrefix(v string) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldHasPrefix(FieldBusinessAccountID, v))
}

// BusinessAccountIDHasSuffix applies the HasSuffix predicate on the "BusinessAccountID" field.
func BusinessAccountIDHasSuffix(v string) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldHasSuffix(FieldBusinessAccountID, v))
}

// BusinessAccountIDEqualFold applies the EqualFold predicate on the "BusinessAccountID" field.
func BusinessAccountIDEqualFold(v string) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldEqualFold(FieldBusinessAccountID, v))
}

// BusinessAccountIDContainsFold applies the ContainsFold predicate on the "BusinessAccountID" field.
func BusinessAccountIDContainsFold(v string) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldContainsFold(FieldBusinessAccountID, v))
}

// TwoFactorAuthenticationEQ applies the EQ predicate on the "TwoFactorAuthentication" field.
func TwoFactorAuthenticationEQ(v bool) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldEQ(FieldTwoFactorAuthentication, v))
}

// TwoFactorAuthenticationNEQ applies the NEQ predicate on the "TwoFactorAuthentication" field.
func TwoFactorAuthenticationNEQ(v bool) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(sql.FieldNEQ(FieldTwoFactorAuthentication, v))
}

// HasBusinessAccount applies the HasEdge predicate on the "business_account" edge.
func HasBusinessAccount() predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, BusinessAccountTable, BusinessAccountColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessAccountWith applies the HasEdge predicate on the "business_account" edge with a given conditions (other predicates).
func HasBusinessAccountWith(preds ...predicate.Business) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(func(s *sql.Selector) {
		step := newBusinessAccountStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BusinessAccountSettings) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BusinessAccountSettings) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(func(s *sql.Selector) {
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
func Not(p predicate.BusinessAccountSettings) predicate.BusinessAccountSettings {
	return predicate.BusinessAccountSettings(func(s *sql.Selector) {
		p(s.Not())
	})
}
