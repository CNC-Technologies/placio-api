// Code generated by ent, DO NOT EDIT.

package like

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the like type in the database.
	Label = "like"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgePost holds the string denoting the post edge name in mutations.
	EdgePost = "post"
	// Table holds the table name of the like in the database.
	Table = "likes"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "likes"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_likes"
	// PostTable is the table that holds the post relation/edge.
	PostTable = "likes"
	// PostInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostInverseTable = "posts"
	// PostColumn is the table column denoting the post relation/edge.
	PostColumn = "like_post"
)

// Columns holds all SQL columns for like fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "likes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"like_post",
	"post_likes",
	"user_likes",
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
	// DefaultCreatedAt holds the default value on creation for the "CreatedAt" field.
	DefaultCreatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "UpdatedAt" field.
	UpdateDefaultUpdatedAt func() time.Time
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the Like queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the CreatedAt field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the UpdatedAt field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByPostField orders the results by post field.
func ByPostField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPostStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newPostStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PostInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, PostTable, PostColumn),
	)
}
