// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"placio-app/ent/template"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Template is the model entity for the Template schema.
type Template struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// DefaultHTML holds the value of the "defaultHTML" field.
	DefaultHTML string `json:"defaultHTML,omitempty"`
	// DefaultCSS holds the value of the "defaultCSS" field.
	DefaultCSS string `json:"defaultCSS,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TemplateQuery when eager-loading is set.
	Edges        TemplateEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TemplateEdges holds the relations/edges for other nodes in the graph.
type TemplateEdges struct {
	// Websites holds the value of the websites edge.
	Websites []*Website `json:"websites,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// WebsitesOrErr returns the Websites value or an error if the edge
// was not loaded in eager-loading.
func (e TemplateEdges) WebsitesOrErr() ([]*Website, error) {
	if e.loadedTypes[0] {
		return e.Websites, nil
	}
	return nil, &NotLoadedError{edge: "websites"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Template) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case template.FieldID, template.FieldName, template.FieldDefaultHTML, template.FieldDefaultCSS:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Template fields.
func (t *Template) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case template.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				t.ID = value.String
			}
		case template.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case template.FieldDefaultHTML:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field defaultHTML", values[i])
			} else if value.Valid {
				t.DefaultHTML = value.String
			}
		case template.FieldDefaultCSS:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field defaultCSS", values[i])
			} else if value.Valid {
				t.DefaultCSS = value.String
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Template.
// This includes values selected through modifiers, order, etc.
func (t *Template) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryWebsites queries the "websites" edge of the Template entity.
func (t *Template) QueryWebsites() *WebsiteQuery {
	return NewTemplateClient(t.config).QueryWebsites(t)
}

// Update returns a builder for updating this Template.
// Note that you need to call Template.Unwrap() before calling this method if this Template
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Template) Update() *TemplateUpdateOne {
	return NewTemplateClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Template entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Template) Unwrap() *Template {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Template is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Template) String() string {
	var builder strings.Builder
	builder.WriteString("Template(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteString(", ")
	builder.WriteString("defaultHTML=")
	builder.WriteString(t.DefaultHTML)
	builder.WriteString(", ")
	builder.WriteString("defaultCSS=")
	builder.WriteString(t.DefaultCSS)
	builder.WriteByte(')')
	return builder.String()
}

// Templates is a parsable slice of Template.
type Templates []*Template
