// Code generated by ent, DO NOT EDIT.

package orderitem

import (
	"placio-app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldContainsFold(FieldID, id))
}

// Quantity applies equality check predicate on the "quantity" field. It's identical to QuantityEQ.
func Quantity(v int) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldQuantity, v))
}

// PricePerItem applies equality check predicate on the "price_per_item" field. It's identical to PricePerItemEQ.
func PricePerItem(v float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldPricePerItem, v))
}

// TotalPrice applies equality check predicate on the "total_price" field. It's identical to TotalPriceEQ.
func TotalPrice(v float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldTotalPrice, v))
}

// QuantityEQ applies the EQ predicate on the "quantity" field.
func QuantityEQ(v int) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldQuantity, v))
}

// QuantityNEQ applies the NEQ predicate on the "quantity" field.
func QuantityNEQ(v int) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNEQ(FieldQuantity, v))
}

// QuantityIn applies the In predicate on the "quantity" field.
func QuantityIn(vs ...int) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldIn(FieldQuantity, vs...))
}

// QuantityNotIn applies the NotIn predicate on the "quantity" field.
func QuantityNotIn(vs ...int) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNotIn(FieldQuantity, vs...))
}

// QuantityGT applies the GT predicate on the "quantity" field.
func QuantityGT(v int) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGT(FieldQuantity, v))
}

// QuantityGTE applies the GTE predicate on the "quantity" field.
func QuantityGTE(v int) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGTE(FieldQuantity, v))
}

// QuantityLT applies the LT predicate on the "quantity" field.
func QuantityLT(v int) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLT(FieldQuantity, v))
}

// QuantityLTE applies the LTE predicate on the "quantity" field.
func QuantityLTE(v int) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLTE(FieldQuantity, v))
}

// PricePerItemEQ applies the EQ predicate on the "price_per_item" field.
func PricePerItemEQ(v float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldPricePerItem, v))
}

// PricePerItemNEQ applies the NEQ predicate on the "price_per_item" field.
func PricePerItemNEQ(v float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNEQ(FieldPricePerItem, v))
}

// PricePerItemIn applies the In predicate on the "price_per_item" field.
func PricePerItemIn(vs ...float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldIn(FieldPricePerItem, vs...))
}

// PricePerItemNotIn applies the NotIn predicate on the "price_per_item" field.
func PricePerItemNotIn(vs ...float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNotIn(FieldPricePerItem, vs...))
}

// PricePerItemGT applies the GT predicate on the "price_per_item" field.
func PricePerItemGT(v float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGT(FieldPricePerItem, v))
}

// PricePerItemGTE applies the GTE predicate on the "price_per_item" field.
func PricePerItemGTE(v float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGTE(FieldPricePerItem, v))
}

// PricePerItemLT applies the LT predicate on the "price_per_item" field.
func PricePerItemLT(v float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLT(FieldPricePerItem, v))
}

// PricePerItemLTE applies the LTE predicate on the "price_per_item" field.
func PricePerItemLTE(v float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLTE(FieldPricePerItem, v))
}

// TotalPriceEQ applies the EQ predicate on the "total_price" field.
func TotalPriceEQ(v float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldTotalPrice, v))
}

// TotalPriceNEQ applies the NEQ predicate on the "total_price" field.
func TotalPriceNEQ(v float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNEQ(FieldTotalPrice, v))
}

// TotalPriceIn applies the In predicate on the "total_price" field.
func TotalPriceIn(vs ...float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldIn(FieldTotalPrice, vs...))
}

// TotalPriceNotIn applies the NotIn predicate on the "total_price" field.
func TotalPriceNotIn(vs ...float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNotIn(FieldTotalPrice, vs...))
}

// TotalPriceGT applies the GT predicate on the "total_price" field.
func TotalPriceGT(v float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGT(FieldTotalPrice, v))
}

// TotalPriceGTE applies the GTE predicate on the "total_price" field.
func TotalPriceGTE(v float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGTE(FieldTotalPrice, v))
}

// TotalPriceLT applies the LT predicate on the "total_price" field.
func TotalPriceLT(v float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLT(FieldTotalPrice, v))
}

// TotalPriceLTE applies the LTE predicate on the "total_price" field.
func TotalPriceLTE(v float64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLTE(FieldTotalPrice, v))
}

// OptionsIsNil applies the IsNil predicate on the "options" field.
func OptionsIsNil() predicate.OrderItem {
	return predicate.OrderItem(sql.FieldIsNull(FieldOptions))
}

// OptionsNotNil applies the NotNil predicate on the "options" field.
func OptionsNotNil() predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNotNull(FieldOptions))
}

// HasOrder applies the HasEdge predicate on the "order" edge.
func HasOrder() predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, OrderTable, OrderPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderWith applies the HasEdge predicate on the "order" edge with a given conditions (other predicates).
func HasOrderWith(preds ...predicate.Order) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		step := newOrderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMenuItem applies the HasEdge predicate on the "menu_item" edge.
func HasMenuItem() predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, MenuItemTable, MenuItemPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMenuItemWith applies the HasEdge predicate on the "menu_item" edge with a given conditions (other predicates).
func HasMenuItemWith(preds ...predicate.MenuItem) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		step := newMenuItemStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrderItem) predicate.OrderItem {
	return predicate.OrderItem(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrderItem) predicate.OrderItem {
	return predicate.OrderItem(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.OrderItem) predicate.OrderItem {
	return predicate.OrderItem(sql.NotPredicates(p))
}