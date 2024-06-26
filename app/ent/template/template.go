// Code generated by ent, DO NOT EDIT.

package template

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the template type in the database.
	Label = "template"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDefaultHTML holds the string denoting the defaulthtml field in the database.
	FieldDefaultHTML = "default_html"
	// FieldDefaultCSS holds the string denoting the defaultcss field in the database.
	FieldDefaultCSS = "default_css"
	// EdgeWebsites holds the string denoting the websites edge name in mutations.
	EdgeWebsites = "websites"
	// Table holds the table name of the template in the database.
	Table = "templates"
	// WebsitesTable is the table that holds the websites relation/edge.
	WebsitesTable = "websites"
	// WebsitesInverseTable is the table name for the Website entity.
	// It exists in this package in order to avoid circular dependency with the "website" package.
	WebsitesInverseTable = "websites"
	// WebsitesColumn is the table column denoting the websites relation/edge.
	WebsitesColumn = "template_websites"
)

// Columns holds all SQL columns for template fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDefaultHTML,
	FieldDefaultCSS,
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

// OrderOption defines the ordering options for the Template queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDefaultHTML orders the results by the defaultHTML field.
func ByDefaultHTML(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDefaultHTML, opts...).ToFunc()
}

// ByDefaultCSS orders the results by the defaultCSS field.
func ByDefaultCSS(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDefaultCSS, opts...).ToFunc()
}

// ByWebsitesCount orders the results by websites count.
func ByWebsitesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWebsitesStep(), opts...)
	}
}

// ByWebsites orders the results by websites terms.
func ByWebsites(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWebsitesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newWebsitesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WebsitesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, WebsitesTable, WebsitesColumn),
	)
}
