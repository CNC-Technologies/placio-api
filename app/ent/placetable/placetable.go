// Code generated by ent, DO NOT EDIT.

package placetable

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the placetable type in the database.
	Label = "place_table"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNumber holds the string denoting the number field in the database.
	FieldNumber = "number"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCapacity holds the string denoting the capacity field in the database.
	FieldCapacity = "capacity"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldIsDeleted holds the string denoting the is_deleted field in the database.
	FieldIsDeleted = "is_deleted"
	// FieldQrCode holds the string denoting the qr_code field in the database.
	FieldQrCode = "qr_code"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldIsActive holds the string denoting the is_active field in the database.
	FieldIsActive = "is_active"
	// FieldIsReserved holds the string denoting the is_reserved field in the database.
	FieldIsReserved = "is_reserved"
	// FieldIsVip holds the string denoting the is_vip field in the database.
	FieldIsVip = "is_vip"
	// FieldIsPremium holds the string denoting the is_premium field in the database.
	FieldIsPremium = "is_premium"
	// FieldLocationDescription holds the string denoting the location_description field in the database.
	FieldLocationDescription = "location_description"
	// FieldMinimumSpend holds the string denoting the minimum_spend field in the database.
	FieldMinimumSpend = "minimum_spend"
	// FieldReservationTime holds the string denoting the reservation_time field in the database.
	FieldReservationTime = "reservation_time"
	// FieldNextAvailableTime holds the string denoting the next_available_time field in the database.
	FieldNextAvailableTime = "next_available_time"
	// FieldSpecialRequirements holds the string denoting the special_requirements field in the database.
	FieldSpecialRequirements = "special_requirements"
	// FieldLayout holds the string denoting the layout field in the database.
	FieldLayout = "layout"
	// FieldServiceArea holds the string denoting the service_area field in the database.
	FieldServiceArea = "service_area"
	// FieldAmbient holds the string denoting the ambient field in the database.
	FieldAmbient = "ambient"
	// FieldImageURL holds the string denoting the image_url field in the database.
	FieldImageURL = "image_url"
	// FieldRating holds the string denoting the rating field in the database.
	FieldRating = "rating"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// FieldMetadata holds the string denoting the metadata field in the database.
	FieldMetadata = "metadata"
	// EdgePlace holds the string denoting the place edge name in mutations.
	EdgePlace = "place"
	// EdgeCreatedBy holds the string denoting the created_by edge name in mutations.
	EdgeCreatedBy = "created_by"
	// EdgeUpdatedBy holds the string denoting the updated_by edge name in mutations.
	EdgeUpdatedBy = "updated_by"
	// EdgeDeletedBy holds the string denoting the deleted_by edge name in mutations.
	EdgeDeletedBy = "deleted_by"
	// EdgeReservedBy holds the string denoting the reserved_by edge name in mutations.
	EdgeReservedBy = "reserved_by"
	// EdgeWaiter holds the string denoting the waiter edge name in mutations.
	EdgeWaiter = "waiter"
	// EdgeOrders holds the string denoting the orders edge name in mutations.
	EdgeOrders = "orders"
	// Table holds the table name of the placetable in the database.
	Table = "place_tables"
	// PlaceTable is the table that holds the place relation/edge.
	PlaceTable = "place_tables"
	// PlaceInverseTable is the table name for the Place entity.
	// It exists in this package in order to avoid circular dependency with the "place" package.
	PlaceInverseTable = "places"
	// PlaceColumn is the table column denoting the place relation/edge.
	PlaceColumn = "place_tables"
	// CreatedByTable is the table that holds the created_by relation/edge.
	CreatedByTable = "place_tables"
	// CreatedByInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	CreatedByInverseTable = "users"
	// CreatedByColumn is the table column denoting the created_by relation/edge.
	CreatedByColumn = "user_tables_created"
	// UpdatedByTable is the table that holds the updated_by relation/edge.
	UpdatedByTable = "place_tables"
	// UpdatedByInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UpdatedByInverseTable = "users"
	// UpdatedByColumn is the table column denoting the updated_by relation/edge.
	UpdatedByColumn = "user_tables_updated"
	// DeletedByTable is the table that holds the deleted_by relation/edge.
	DeletedByTable = "place_tables"
	// DeletedByInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	DeletedByInverseTable = "users"
	// DeletedByColumn is the table column denoting the deleted_by relation/edge.
	DeletedByColumn = "user_tables_deleted"
	// ReservedByTable is the table that holds the reserved_by relation/edge.
	ReservedByTable = "place_tables"
	// ReservedByInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	ReservedByInverseTable = "users"
	// ReservedByColumn is the table column denoting the reserved_by relation/edge.
	ReservedByColumn = "user_tables_reserved"
	// WaiterTable is the table that holds the waiter relation/edge.
	WaiterTable = "place_tables"
	// WaiterInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	WaiterInverseTable = "users"
	// WaiterColumn is the table column denoting the waiter relation/edge.
	WaiterColumn = "user_tables_waited"
	// OrdersTable is the table that holds the orders relation/edge. The primary key declared below.
	OrdersTable = "place_table_orders"
	// OrdersInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	OrdersInverseTable = "orders"
)

// Columns holds all SQL columns for placetable fields.
var Columns = []string{
	FieldID,
	FieldNumber,
	FieldName,
	FieldCapacity,
	FieldDeletedAt,
	FieldIsDeleted,
	FieldQrCode,
	FieldDescription,
	FieldStatus,
	FieldType,
	FieldIsActive,
	FieldIsReserved,
	FieldIsVip,
	FieldIsPremium,
	FieldLocationDescription,
	FieldMinimumSpend,
	FieldReservationTime,
	FieldNextAvailableTime,
	FieldSpecialRequirements,
	FieldLayout,
	FieldServiceArea,
	FieldAmbient,
	FieldImageURL,
	FieldRating,
	FieldTags,
	FieldMetadata,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "place_tables"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"place_tables",
	"user_tables_created",
	"user_tables_updated",
	"user_tables_deleted",
	"user_tables_reserved",
	"user_tables_waited",
}

var (
	// OrdersPrimaryKey and OrdersColumn2 are the table columns denoting the
	// primary key for the orders relation (M2M).
	OrdersPrimaryKey = []string{"place_table_id", "order_id"}
)

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
	// DefaultCapacity holds the default value on creation for the "capacity" field.
	DefaultCapacity int
	// DefaultIsDeleted holds the default value on creation for the "is_deleted" field.
	DefaultIsDeleted bool
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus string
	// DefaultIsActive holds the default value on creation for the "is_active" field.
	DefaultIsActive bool
	// DefaultIsReserved holds the default value on creation for the "is_reserved" field.
	DefaultIsReserved bool
	// DefaultIsVip holds the default value on creation for the "is_vip" field.
	DefaultIsVip bool
	// DefaultIsPremium holds the default value on creation for the "is_premium" field.
	DefaultIsPremium bool
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeRegular Type = "regular"
	TypeVip     Type = "vip"
	TypePremium Type = "premium"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeRegular, TypeVip, TypePremium:
		return nil
	default:
		return fmt.Errorf("placetable: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the PlaceTable queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByNumber orders the results by the number field.
func ByNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNumber, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCapacity orders the results by the capacity field.
func ByCapacity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCapacity, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByIsDeleted orders the results by the is_deleted field.
func ByIsDeleted(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsDeleted, opts...).ToFunc()
}

// ByQrCode orders the results by the qr_code field.
func ByQrCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQrCode, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByIsActive orders the results by the is_active field.
func ByIsActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsActive, opts...).ToFunc()
}

// ByIsReserved orders the results by the is_reserved field.
func ByIsReserved(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsReserved, opts...).ToFunc()
}

// ByIsVip orders the results by the is_vip field.
func ByIsVip(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsVip, opts...).ToFunc()
}

// ByIsPremium orders the results by the is_premium field.
func ByIsPremium(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsPremium, opts...).ToFunc()
}

// ByLocationDescription orders the results by the location_description field.
func ByLocationDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocationDescription, opts...).ToFunc()
}

// ByMinimumSpend orders the results by the minimum_spend field.
func ByMinimumSpend(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMinimumSpend, opts...).ToFunc()
}

// ByReservationTime orders the results by the reservation_time field.
func ByReservationTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReservationTime, opts...).ToFunc()
}

// ByNextAvailableTime orders the results by the next_available_time field.
func ByNextAvailableTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNextAvailableTime, opts...).ToFunc()
}

// ByLayout orders the results by the layout field.
func ByLayout(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLayout, opts...).ToFunc()
}

// ByServiceArea orders the results by the service_area field.
func ByServiceArea(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldServiceArea, opts...).ToFunc()
}

// ByAmbient orders the results by the ambient field.
func ByAmbient(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmbient, opts...).ToFunc()
}

// ByImageURL orders the results by the image_url field.
func ByImageURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImageURL, opts...).ToFunc()
}

// ByRating orders the results by the rating field.
func ByRating(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRating, opts...).ToFunc()
}

// ByPlaceField orders the results by place field.
func ByPlaceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPlaceStep(), sql.OrderByField(field, opts...))
	}
}

// ByCreatedByField orders the results by created_by field.
func ByCreatedByField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCreatedByStep(), sql.OrderByField(field, opts...))
	}
}

// ByUpdatedByField orders the results by updated_by field.
func ByUpdatedByField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUpdatedByStep(), sql.OrderByField(field, opts...))
	}
}

// ByDeletedByField orders the results by deleted_by field.
func ByDeletedByField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDeletedByStep(), sql.OrderByField(field, opts...))
	}
}

// ByReservedByField orders the results by reserved_by field.
func ByReservedByField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReservedByStep(), sql.OrderByField(field, opts...))
	}
}

// ByWaiterField orders the results by waiter field.
func ByWaiterField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWaiterStep(), sql.OrderByField(field, opts...))
	}
}

// ByOrdersCount orders the results by orders count.
func ByOrdersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOrdersStep(), opts...)
	}
}

// ByOrders orders the results by orders terms.
func ByOrders(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrdersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPlaceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PlaceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PlaceTable, PlaceColumn),
	)
}
func newCreatedByStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CreatedByInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CreatedByTable, CreatedByColumn),
	)
}
func newUpdatedByStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UpdatedByInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UpdatedByTable, UpdatedByColumn),
	)
}
func newDeletedByStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DeletedByInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, DeletedByTable, DeletedByColumn),
	)
}
func newReservedByStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReservedByInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ReservedByTable, ReservedByColumn),
	)
}
func newWaiterStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WaiterInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, WaiterTable, WaiterColumn),
	)
}
func newOrdersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrdersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, OrdersTable, OrdersPrimaryKey...),
	)
}
