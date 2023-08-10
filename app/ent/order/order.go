// Code generated by ent, DO NOT EDIT.

package order

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the order type in the db.
	Label = "order"
	// FieldID holds the string denoting the id field in the db.
	FieldID = "id"
	// Table holds the table name of the order in the db.
	Table = "orders"
)

// Columns holds all SQL columns for order fields.
var Columns = []string{
	FieldID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Order queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}