// Code generated by ent, DO NOT EDIT.

package accountwallet

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the accountwallet type in the database.
	Label = "account_wallet"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBalance holds the string denoting the balance field in the database.
	FieldBalance = "balance"
	// FieldTotalDeposited holds the string denoting the total_deposited field in the database.
	FieldTotalDeposited = "total_deposited"
	// FieldTotalWithdrawn holds the string denoting the total_withdrawn field in the database.
	FieldTotalWithdrawn = "total_withdrawn"
	// FieldTotalEarned holds the string denoting the total_earned field in the database.
	FieldTotalEarned = "total_earned"
	// FieldTotalSpent holds the string denoting the total_spent field in the database.
	FieldTotalSpent = "total_spent"
	// FieldTotalRefunded holds the string denoting the total_refunded field in the database.
	FieldTotalRefunded = "total_refunded"
	// FieldTotalFees holds the string denoting the total_fees field in the database.
	FieldTotalFees = "total_fees"
	// FieldTotalTax holds the string denoting the total_tax field in the database.
	FieldTotalTax = "total_tax"
	// FieldTotalDiscount holds the string denoting the total_discount field in the database.
	FieldTotalDiscount = "total_discount"
	// FieldTotalRevenue holds the string denoting the total_revenue field in the database.
	FieldTotalRevenue = "total_revenue"
	// FieldTotalExpenses holds the string denoting the total_expenses field in the database.
	FieldTotalExpenses = "total_expenses"
	// FieldTotalProfit holds the string denoting the total_profit field in the database.
	FieldTotalProfit = "total_profit"
	// FieldTotalLoss holds the string denoting the total_loss field in the database.
	FieldTotalLoss = "total_loss"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeBusiness holds the string denoting the business edge name in mutations.
	EdgeBusiness = "business"
	// Table holds the table name of the accountwallet in the database.
	Table = "account_wallets"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "account_wallets"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_wallet"
	// BusinessTable is the table that holds the business relation/edge.
	BusinessTable = "account_wallets"
	// BusinessInverseTable is the table name for the Business entity.
	// It exists in this package in order to avoid circular dependency with the "business" package.
	BusinessInverseTable = "businesses"
	// BusinessColumn is the table column denoting the business relation/edge.
	BusinessColumn = "business_wallet"
)

// Columns holds all SQL columns for accountwallet fields.
var Columns = []string{
	FieldID,
	FieldBalance,
	FieldTotalDeposited,
	FieldTotalWithdrawn,
	FieldTotalEarned,
	FieldTotalSpent,
	FieldTotalRefunded,
	FieldTotalFees,
	FieldTotalTax,
	FieldTotalDiscount,
	FieldTotalRevenue,
	FieldTotalExpenses,
	FieldTotalProfit,
	FieldTotalLoss,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "account_wallets"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"business_wallet",
	"user_wallet",
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
	// DefaultBalance holds the default value on creation for the "balance" field.
	DefaultBalance float64
	// DefaultTotalDeposited holds the default value on creation for the "total_deposited" field.
	DefaultTotalDeposited float64
	// DefaultTotalWithdrawn holds the default value on creation for the "total_withdrawn" field.
	DefaultTotalWithdrawn float64
	// DefaultTotalEarned holds the default value on creation for the "total_earned" field.
	DefaultTotalEarned float64
	// DefaultTotalSpent holds the default value on creation for the "total_spent" field.
	DefaultTotalSpent float64
	// DefaultTotalRefunded holds the default value on creation for the "total_refunded" field.
	DefaultTotalRefunded float64
	// DefaultTotalFees holds the default value on creation for the "total_fees" field.
	DefaultTotalFees float64
	// DefaultTotalTax holds the default value on creation for the "total_tax" field.
	DefaultTotalTax float64
	// DefaultTotalDiscount holds the default value on creation for the "total_discount" field.
	DefaultTotalDiscount float64
	// DefaultTotalRevenue holds the default value on creation for the "total_revenue" field.
	DefaultTotalRevenue float64
	// DefaultTotalExpenses holds the default value on creation for the "total_expenses" field.
	DefaultTotalExpenses float64
	// DefaultTotalProfit holds the default value on creation for the "total_profit" field.
	DefaultTotalProfit float64
	// DefaultTotalLoss holds the default value on creation for the "total_loss" field.
	DefaultTotalLoss float64
)

// OrderOption defines the ordering options for the AccountWallet queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByBalance orders the results by the balance field.
func ByBalance(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBalance, opts...).ToFunc()
}

// ByTotalDeposited orders the results by the total_deposited field.
func ByTotalDeposited(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalDeposited, opts...).ToFunc()
}

// ByTotalWithdrawn orders the results by the total_withdrawn field.
func ByTotalWithdrawn(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalWithdrawn, opts...).ToFunc()
}

// ByTotalEarned orders the results by the total_earned field.
func ByTotalEarned(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalEarned, opts...).ToFunc()
}

// ByTotalSpent orders the results by the total_spent field.
func ByTotalSpent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalSpent, opts...).ToFunc()
}

// ByTotalRefunded orders the results by the total_refunded field.
func ByTotalRefunded(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalRefunded, opts...).ToFunc()
}

// ByTotalFees orders the results by the total_fees field.
func ByTotalFees(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalFees, opts...).ToFunc()
}

// ByTotalTax orders the results by the total_tax field.
func ByTotalTax(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalTax, opts...).ToFunc()
}

// ByTotalDiscount orders the results by the total_discount field.
func ByTotalDiscount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalDiscount, opts...).ToFunc()
}

// ByTotalRevenue orders the results by the total_revenue field.
func ByTotalRevenue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalRevenue, opts...).ToFunc()
}

// ByTotalExpenses orders the results by the total_expenses field.
func ByTotalExpenses(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalExpenses, opts...).ToFunc()
}

// ByTotalProfit orders the results by the total_profit field.
func ByTotalProfit(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalProfit, opts...).ToFunc()
}

// ByTotalLoss orders the results by the total_loss field.
func ByTotalLoss(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalLoss, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByBusinessField orders the results by business field.
func ByBusinessField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBusinessStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
	)
}
func newBusinessStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BusinessInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, BusinessTable, BusinessColumn),
	)
}