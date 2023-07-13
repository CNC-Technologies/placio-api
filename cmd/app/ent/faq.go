// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"placio-app/ent/business"
	"placio-app/ent/faq"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// FAQ is the model entity for the FAQ schema.
type FAQ struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Question holds the value of the "question" field.
	Question string `json:"question,omitempty"`
	// Answer holds the value of the "answer" field.
	Answer string `json:"answer,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FAQQuery when eager-loading is set.
	Edges         FAQEdges `json:"edges"`
	business_faqs *string
	selectValues  sql.SelectValues
}

// FAQEdges holds the relations/edges for other nodes in the graph.
type FAQEdges struct {
	// Business holds the value of the business edge.
	Business *Business `json:"business,omitempty"`
	// Place holds the value of the place edge.
	Place []*Place `json:"place,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// BusinessOrErr returns the Business value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FAQEdges) BusinessOrErr() (*Business, error) {
	if e.loadedTypes[0] {
		if e.Business == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: business.Label}
		}
		return e.Business, nil
	}
	return nil, &NotLoadedError{edge: "business"}
}

// PlaceOrErr returns the Place value or an error if the edge
// was not loaded in eager-loading.
func (e FAQEdges) PlaceOrErr() ([]*Place, error) {
	if e.loadedTypes[1] {
		return e.Place, nil
	}
	return nil, &NotLoadedError{edge: "place"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FAQ) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case faq.FieldID, faq.FieldQuestion, faq.FieldAnswer:
			values[i] = new(sql.NullString)
		case faq.ForeignKeys[0]: // business_faqs
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FAQ fields.
func (f *FAQ) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case faq.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				f.ID = value.String
			}
		case faq.FieldQuestion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field question", values[i])
			} else if value.Valid {
				f.Question = value.String
			}
		case faq.FieldAnswer:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field answer", values[i])
			} else if value.Valid {
				f.Answer = value.String
			}
		case faq.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field business_faqs", values[i])
			} else if value.Valid {
				f.business_faqs = new(string)
				*f.business_faqs = value.String
			}
		default:
			f.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the FAQ.
// This includes values selected through modifiers, order, etc.
func (f *FAQ) Value(name string) (ent.Value, error) {
	return f.selectValues.Get(name)
}

// QueryBusiness queries the "business" edge of the FAQ entity.
func (f *FAQ) QueryBusiness() *BusinessQuery {
	return NewFAQClient(f.config).QueryBusiness(f)
}

// QueryPlace queries the "place" edge of the FAQ entity.
func (f *FAQ) QueryPlace() *PlaceQuery {
	return NewFAQClient(f.config).QueryPlace(f)
}

// Update returns a builder for updating this FAQ.
// Note that you need to call FAQ.Unwrap() before calling this method if this FAQ
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *FAQ) Update() *FAQUpdateOne {
	return NewFAQClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the FAQ entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *FAQ) Unwrap() *FAQ {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: FAQ is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *FAQ) String() string {
	var builder strings.Builder
	builder.WriteString("FAQ(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("question=")
	builder.WriteString(f.Question)
	builder.WriteString(", ")
	builder.WriteString("answer=")
	builder.WriteString(f.Answer)
	builder.WriteByte(')')
	return builder.String()
}

// FAQs is a parsable slice of FAQ.
type FAQs []*FAQ
