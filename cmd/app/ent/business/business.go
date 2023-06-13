// Code generated by ent, DO NOT EDIT.

package business

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the business type in the database.
	Label = "business"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeUserBusinesses holds the string denoting the userbusinesses edge name in mutations.
	EdgeUserBusinesses = "userBusinesses"
	// EdgeBusinessAccountSettings holds the string denoting the business_account_settings edge name in mutations.
	EdgeBusinessAccountSettings = "business_account_settings"
	// EdgePosts holds the string denoting the posts edge name in mutations.
	EdgePosts = "posts"
	// Table holds the table name of the business in the database.
	Table = "businesses"
	// UserBusinessesTable is the table that holds the userBusinesses relation/edge.
	UserBusinessesTable = "user_businesses"
	// UserBusinessesInverseTable is the table name for the UserBusiness entity.
	// It exists in this package in order to avoid circular dependency with the "userbusiness" package.
	UserBusinessesInverseTable = "user_businesses"
	// UserBusinessesColumn is the table column denoting the userBusinesses relation/edge.
	UserBusinessesColumn = "business_user_businesses"
	// BusinessAccountSettingsTable is the table that holds the business_account_settings relation/edge.
	BusinessAccountSettingsTable = "account_settings"
	// BusinessAccountSettingsInverseTable is the table name for the AccountSettings entity.
	// It exists in this package in order to avoid circular dependency with the "accountsettings" package.
	BusinessAccountSettingsInverseTable = "account_settings"
	// BusinessAccountSettingsColumn is the table column denoting the business_account_settings relation/edge.
	BusinessAccountSettingsColumn = "business_business_account_settings"
	// PostsTable is the table that holds the posts relation/edge.
	PostsTable = "posts"
	// PostsInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostsInverseTable = "posts"
	// PostsColumn is the table column denoting the posts relation/edge.
	PostsColumn = "business_posts"
)

// Columns holds all SQL columns for business fields.
var Columns = []string{
	FieldID,
	FieldName,
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

var (
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the Business queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByUserBusinessesCount orders the results by userBusinesses count.
func ByUserBusinessesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserBusinessesStep(), opts...)
	}
}

// ByUserBusinesses orders the results by userBusinesses terms.
func ByUserBusinesses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserBusinessesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByBusinessAccountSettingsField orders the results by business_account_settings field.
func ByBusinessAccountSettingsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBusinessAccountSettingsStep(), sql.OrderByField(field, opts...))
	}
}

// ByPostsCount orders the results by posts count.
func ByPostsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPostsStep(), opts...)
	}
}

// ByPosts orders the results by posts terms.
func ByPosts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPostsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUserBusinessesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserBusinessesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UserBusinessesTable, UserBusinessesColumn),
	)
}
func newBusinessAccountSettingsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BusinessAccountSettingsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, BusinessAccountSettingsTable, BusinessAccountSettingsColumn),
	)
}
func newPostsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PostsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PostsTable, PostsColumn),
	)
}