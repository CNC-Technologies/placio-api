// Code generated by ent, DO NOT EDIT.

package room

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the room type in the db.
	Label = "room"
	// FieldID holds the string denoting the id field in the db.
	FieldID = "id"
	// FieldNumber holds the string denoting the number field in the db.
	FieldNumber = "number"
	// FieldType holds the string denoting the type field in the db.
	FieldType = "type"
	// FieldPrice holds the string denoting the price field in the db.
	FieldPrice = "price"
	// FieldDescription holds the string denoting the description field in the db.
	FieldDescription = "description"
	// FieldAvailability holds the string denoting the availability field in the db.
	FieldAvailability = "availability"
	// FieldImage holds the string denoting the image field in the db.
	FieldImage = "image"
	// EdgePlace holds the string denoting the place edge name in mutations.
	EdgePlace = "place"
	// EdgeBookings holds the string denoting the bookings edge name in mutations.
	EdgeBookings = "bookings"
	// Table holds the table name of the room in the db.
	Table = "rooms"
	// PlaceTable is the table that holds the place relation/edge.
	PlaceTable = "rooms"
	// PlaceInverseTable is the table name for the Place entity.
	// It exists in this package in order to avoid circular dependency with the "place" package.
	PlaceInverseTable = "places"
	// PlaceColumn is the table column denoting the place relation/edge.
	PlaceColumn = "place_rooms"
	// BookingsTable is the table that holds the bookings relation/edge.
	BookingsTable = "bookings"
	// BookingsInverseTable is the table name for the Booking entity.
	// It exists in this package in order to avoid circular dependency with the "booking" package.
	BookingsInverseTable = "bookings"
	// BookingsColumn is the table column denoting the bookings relation/edge.
	BookingsColumn = "room_bookings"
)

// Columns holds all SQL columns for room fields.
var Columns = []string{
	FieldID,
	FieldNumber,
	FieldType,
	FieldPrice,
	FieldDescription,
	FieldAvailability,
	FieldImage,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "rooms"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"place_rooms",
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

// OrderOption defines the ordering options for the Room queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByNumber orders the results by the number field.
func ByNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNumber, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByAvailability orders the results by the availability field.
func ByAvailability(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvailability, opts...).ToFunc()
}

// ByImage orders the results by the image field.
func ByImage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImage, opts...).ToFunc()
}

// ByPlaceField orders the results by place field.
func ByPlaceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPlaceStep(), sql.OrderByField(field, opts...))
	}
}

// ByBookingsCount orders the results by bookings count.
func ByBookingsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBookingsStep(), opts...)
	}
}

// ByBookings orders the results by bookings terms.
func ByBookings(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBookingsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPlaceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PlaceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PlaceTable, PlaceColumn),
	)
}
func newBookingsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BookingsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, BookingsTable, BookingsColumn),
	)
}