// Code generated by ent, DO NOT EDIT.

package event

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the event type in the database.
	Label = "event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldSearchText holds the string denoting the search_text field in the database.
	FieldSearchText = "search_text"
	// FieldRelevanceScore holds the string denoting the relevance_score field in the database.
	FieldRelevanceScore = "relevance_score"
	// EdgeTickets holds the string denoting the tickets edge name in mutations.
	EdgeTickets = "tickets"
	// EdgeTicketOptions holds the string denoting the ticket_options edge name in mutations.
	EdgeTicketOptions = "ticket_options"
	// Table holds the table name of the event in the database.
	Table = "events"
	// TicketsTable is the table that holds the tickets relation/edge.
	TicketsTable = "tickets"
	// TicketsInverseTable is the table name for the Ticket entity.
	// It exists in this package in order to avoid circular dependency with the "ticket" package.
	TicketsInverseTable = "tickets"
	// TicketsColumn is the table column denoting the tickets relation/edge.
	TicketsColumn = "event_tickets"
	// TicketOptionsTable is the table that holds the ticket_options relation/edge.
	TicketOptionsTable = "ticket_options"
	// TicketOptionsInverseTable is the table name for the TicketOption entity.
	// It exists in this package in order to avoid circular dependency with the "ticketoption" package.
	TicketOptionsInverseTable = "ticket_options"
	// TicketOptionsColumn is the table column denoting the ticket_options relation/edge.
	TicketOptionsColumn = "event_ticket_options"
)

// Columns holds all SQL columns for event fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldSearchText,
	FieldRelevanceScore,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "events"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"place_events",
	"user_events",
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updatedAt" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updatedAt" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the Event queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCreatedAt orders the results by the createdAt field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updatedAt field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// BySearchText orders the results by the search_text field.
func BySearchText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSearchText, opts...).ToFunc()
}

// ByRelevanceScore orders the results by the relevance_score field.
func ByRelevanceScore(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRelevanceScore, opts...).ToFunc()
}

// ByTicketsCount orders the results by tickets count.
func ByTicketsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTicketsStep(), opts...)
	}
}

// ByTickets orders the results by tickets terms.
func ByTickets(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTicketsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTicketOptionsCount orders the results by ticket_options count.
func ByTicketOptionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTicketOptionsStep(), opts...)
	}
}

// ByTicketOptions orders the results by ticket_options terms.
func ByTicketOptions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTicketOptionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTicketsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TicketsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TicketsTable, TicketsColumn),
	)
}
func newTicketOptionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TicketOptionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TicketOptionsTable, TicketOptionsColumn),
	)
}
