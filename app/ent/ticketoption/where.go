// Code generated by ent, DO NOT EDIT.

package ticketoption

import (
	"placio-app/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldContainsFold(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldEQ(FieldDescription, v))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v float64) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldEQ(FieldPrice, v))
}

// QuantityAvailable applies equality check predicate on the "quantityAvailable" field. It's identical to QuantityAvailableEQ.
func QuantityAvailable(v int) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldEQ(FieldQuantityAvailable, v))
}

// QuantitySold applies equality check predicate on the "quantitySold" field. It's identical to QuantitySoldEQ.
func QuantitySold(v int) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldEQ(FieldQuantitySold, v))
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldEQ(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldContainsFold(FieldDescription, v))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v float64) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v float64) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...float64) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...float64) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v float64) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v float64) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v float64) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v float64) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldLTE(FieldPrice, v))
}

// QuantityAvailableEQ applies the EQ predicate on the "quantityAvailable" field.
func QuantityAvailableEQ(v int) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldEQ(FieldQuantityAvailable, v))
}

// QuantityAvailableNEQ applies the NEQ predicate on the "quantityAvailable" field.
func QuantityAvailableNEQ(v int) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldNEQ(FieldQuantityAvailable, v))
}

// QuantityAvailableIn applies the In predicate on the "quantityAvailable" field.
func QuantityAvailableIn(vs ...int) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldIn(FieldQuantityAvailable, vs...))
}

// QuantityAvailableNotIn applies the NotIn predicate on the "quantityAvailable" field.
func QuantityAvailableNotIn(vs ...int) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldNotIn(FieldQuantityAvailable, vs...))
}

// QuantityAvailableGT applies the GT predicate on the "quantityAvailable" field.
func QuantityAvailableGT(v int) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldGT(FieldQuantityAvailable, v))
}

// QuantityAvailableGTE applies the GTE predicate on the "quantityAvailable" field.
func QuantityAvailableGTE(v int) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldGTE(FieldQuantityAvailable, v))
}

// QuantityAvailableLT applies the LT predicate on the "quantityAvailable" field.
func QuantityAvailableLT(v int) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldLT(FieldQuantityAvailable, v))
}

// QuantityAvailableLTE applies the LTE predicate on the "quantityAvailable" field.
func QuantityAvailableLTE(v int) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldLTE(FieldQuantityAvailable, v))
}

// QuantitySoldEQ applies the EQ predicate on the "quantitySold" field.
func QuantitySoldEQ(v int) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldEQ(FieldQuantitySold, v))
}

// QuantitySoldNEQ applies the NEQ predicate on the "quantitySold" field.
func QuantitySoldNEQ(v int) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldNEQ(FieldQuantitySold, v))
}

// QuantitySoldIn applies the In predicate on the "quantitySold" field.
func QuantitySoldIn(vs ...int) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldIn(FieldQuantitySold, vs...))
}

// QuantitySoldNotIn applies the NotIn predicate on the "quantitySold" field.
func QuantitySoldNotIn(vs ...int) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldNotIn(FieldQuantitySold, vs...))
}

// QuantitySoldGT applies the GT predicate on the "quantitySold" field.
func QuantitySoldGT(v int) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldGT(FieldQuantitySold, v))
}

// QuantitySoldGTE applies the GTE predicate on the "quantitySold" field.
func QuantitySoldGTE(v int) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldGTE(FieldQuantitySold, v))
}

// QuantitySoldLT applies the LT predicate on the "quantitySold" field.
func QuantitySoldLT(v int) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldLT(FieldQuantitySold, v))
}

// QuantitySoldLTE applies the LTE predicate on the "quantitySold" field.
func QuantitySoldLTE(v int) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldLTE(FieldQuantitySold, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldNotIn(FieldStatus, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updatedAt" field.
func UpdatedAtGT(v time.Time) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updatedAt" field.
func UpdatedAtLT(v time.Time) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.TicketOption {
	return predicate.TicketOption(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasEvent applies the HasEdge predicate on the "event" edge.
func HasEvent() predicate.TicketOption {
	return predicate.TicketOption(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EventTable, EventColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventWith applies the HasEdge predicate on the "event" edge with a given conditions (other predicates).
func HasEventWith(preds ...predicate.Event) predicate.TicketOption {
	return predicate.TicketOption(func(s *sql.Selector) {
		step := newEventStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTickets applies the HasEdge predicate on the "tickets" edge.
func HasTickets() predicate.TicketOption {
	return predicate.TicketOption(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TicketsTable, TicketsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTicketsWith applies the HasEdge predicate on the "tickets" edge with a given conditions (other predicates).
func HasTicketsWith(preds ...predicate.Ticket) predicate.TicketOption {
	return predicate.TicketOption(func(s *sql.Selector) {
		step := newTicketsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TicketOption) predicate.TicketOption {
	return predicate.TicketOption(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TicketOption) predicate.TicketOption {
	return predicate.TicketOption(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TicketOption) predicate.TicketOption {
	return predicate.TicketOption(sql.NotPredicates(p))
}
