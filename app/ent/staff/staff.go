// Code generated by ent, DO NOT EDIT.

package staff

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the staff type in the database.
	Label = "staff"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPosition holds the string denoting the position field in the database.
	FieldPosition = "position"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgePlace holds the string denoting the place edge name in mutations.
	EdgePlace = "place"
	// EdgePermissions holds the string denoting the permissions edge name in mutations.
	EdgePermissions = "permissions"
	// EdgeBusiness holds the string denoting the business edge name in mutations.
	EdgeBusiness = "business"
	// Table holds the table name of the staff in the database.
	Table = "staffs"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "staffs"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_staffs"
	// PlaceTable is the table that holds the place relation/edge. The primary key declared below.
	PlaceTable = "place_staffs"
	// PlaceInverseTable is the table name for the Place entity.
	// It exists in this package in order to avoid circular dependency with the "place" package.
	PlaceInverseTable = "places"
	// PermissionsTable is the table that holds the permissions relation/edge. The primary key declared below.
	PermissionsTable = "staff_permissions"
	// PermissionsInverseTable is the table name for the Permission entity.
	// It exists in this package in order to avoid circular dependency with the "permission" package.
	PermissionsInverseTable = "permissions"
	// BusinessTable is the table that holds the business relation/edge. The primary key declared below.
	BusinessTable = "business_staffs"
	// BusinessInverseTable is the table name for the Business entity.
	// It exists in this package in order to avoid circular dependency with the "business" package.
	BusinessInverseTable = "businesses"
)

// Columns holds all SQL columns for staff fields.
var Columns = []string{
	FieldID,
	FieldPosition,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "staffs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_staffs",
}

var (
	// PlacePrimaryKey and PlaceColumn2 are the table columns denoting the
	// primary key for the place relation (M2M).
	PlacePrimaryKey = []string{"place_id", "staff_id"}
	// PermissionsPrimaryKey and PermissionsColumn2 are the table columns denoting the
	// primary key for the permissions relation (M2M).
	PermissionsPrimaryKey = []string{"staff_id", "permission_id"}
	// BusinessPrimaryKey and BusinessColumn2 are the table columns denoting the
	// primary key for the business relation (M2M).
	BusinessPrimaryKey = []string{"business_id", "staff_id"}
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
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the Staff queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPosition orders the results by the position field.
func ByPosition(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPosition, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByPlaceCount orders the results by place count.
func ByPlaceCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPlaceStep(), opts...)
	}
}

// ByPlace orders the results by place terms.
func ByPlace(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPlaceStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPermissionsCount orders the results by permissions count.
func ByPermissionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPermissionsStep(), opts...)
	}
}

// ByPermissions orders the results by permissions terms.
func ByPermissions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPermissionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByBusinessCount orders the results by business count.
func ByBusinessCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBusinessStep(), opts...)
	}
}

// ByBusiness orders the results by business terms.
func ByBusiness(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBusinessStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newPlaceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PlaceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, PlaceTable, PlacePrimaryKey...),
	)
}
func newPermissionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PermissionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, PermissionsTable, PermissionsPrimaryKey...),
	)
}
func newBusinessStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BusinessInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, BusinessTable, BusinessPrimaryKey...),
	)
}