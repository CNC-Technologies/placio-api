// Code generated by ent, DO NOT EDIT.

package placeinventoryattribute

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the placeinventoryattribute type in the database.
	Label = "place_inventory_attribute"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldCategorySpecificValue holds the string denoting the category_specific_value field in the database.
	FieldCategorySpecificValue = "category_specific_value"
	// EdgeInventory holds the string denoting the inventory edge name in mutations.
	EdgeInventory = "inventory"
	// EdgeAttributeType holds the string denoting the attribute_type edge name in mutations.
	EdgeAttributeType = "attribute_type"
	// Table holds the table name of the placeinventoryattribute in the database.
	Table = "place_inventory_attributes"
	// InventoryTable is the table that holds the inventory relation/edge.
	InventoryTable = "place_inventory_attributes"
	// InventoryInverseTable is the table name for the PlaceInventory entity.
	// It exists in this package in order to avoid circular dependency with the "placeinventory" package.
	InventoryInverseTable = "place_inventories"
	// InventoryColumn is the table column denoting the inventory relation/edge.
	InventoryColumn = "place_inventory_attributes"
	// AttributeTypeTable is the table that holds the attribute_type relation/edge.
	AttributeTypeTable = "place_inventory_attributes"
	// AttributeTypeInverseTable is the table name for the InventoryAttribute entity.
	// It exists in this package in order to avoid circular dependency with the "inventoryattribute" package.
	AttributeTypeInverseTable = "inventory_attributes"
	// AttributeTypeColumn is the table column denoting the attribute_type relation/edge.
	AttributeTypeColumn = "inventory_attribute_place_inventory_attributes"
)

// Columns holds all SQL columns for placeinventoryattribute fields.
var Columns = []string{
	FieldID,
	FieldValue,
	FieldCategorySpecificValue,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "place_inventory_attributes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"inventory_attribute_place_inventory_attributes",
	"place_inventory_attributes",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the PlaceInventoryAttribute queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByValue orders the results by the value field.
func ByValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValue, opts...).ToFunc()
}

// ByInventoryField orders the results by inventory field.
func ByInventoryField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInventoryStep(), sql.OrderByField(field, opts...))
	}
}

// ByAttributeTypeField orders the results by attribute_type field.
func ByAttributeTypeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAttributeTypeStep(), sql.OrderByField(field, opts...))
	}
}
func newInventoryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InventoryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, InventoryTable, InventoryColumn),
	)
}
func newAttributeTypeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AttributeTypeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AttributeTypeTable, AttributeTypeColumn),
	)
}
