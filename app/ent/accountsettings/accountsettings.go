// Code generated by ent, DO NOT EDIT.

package accountsettings

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the accountsettings type in the db.
	Label = "account_settings"
	// FieldID holds the string denoting the id field in the db.
	FieldID = "id"
	// FieldTwoFactorAuthentication holds the string denoting the twofactorauthentication field in the db.
	FieldTwoFactorAuthentication = "two_factor_authentication"
	// FieldBlockedUsers holds the string denoting the blockedusers field in the db.
	FieldBlockedUsers = "blocked_users"
	// FieldMutedUsers holds the string denoting the mutedusers field in the db.
	FieldMutedUsers = "muted_users"
	// EdgeBusinessAccount holds the string denoting the business_account edge name in mutations.
	EdgeBusinessAccount = "business_account"
	// Table holds the table name of the accountsettings in the db.
	Table = "account_settings"
	// BusinessAccountTable is the table that holds the business_account relation/edge.
	BusinessAccountTable = "account_settings"
	// BusinessAccountInverseTable is the table name for the Business entity.
	// It exists in this package in order to avoid circular dependency with the "business" package.
	BusinessAccountInverseTable = "businesses"
	// BusinessAccountColumn is the table column denoting the business_account relation/edge.
	BusinessAccountColumn = "business_business_account_settings"
)

// Columns holds all SQL columns for accountsettings fields.
var Columns = []string{
	FieldID,
	FieldTwoFactorAuthentication,
	FieldBlockedUsers,
	FieldMutedUsers,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "account_settings"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"business_business_account_settings",
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

// OrderOption defines the ordering options for the AccountSettings queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTwoFactorAuthentication orders the results by the TwoFactorAuthentication field.
func ByTwoFactorAuthentication(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTwoFactorAuthentication, opts...).ToFunc()
}

// ByBusinessAccountField orders the results by business_account field.
func ByBusinessAccountField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBusinessAccountStep(), sql.OrderByField(field, opts...))
	}
}
func newBusinessAccountStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BusinessAccountInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, BusinessAccountTable, BusinessAccountColumn),
	)
}