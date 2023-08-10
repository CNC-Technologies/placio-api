// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"placio-app/ent/booking"
	"placio-app/ent/room"
	"placio-app/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Booking is the model entity for the Booking schema.
type Booking struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// StartDate holds the value of the "startDate" field.
	StartDate time.Time `json:"startDate,omitempty"`
	// EndDate holds the value of the "endDate" field.
	EndDate time.Time `json:"endDate,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// BookingDate holds the value of the "bookingDate" field.
	BookingDate time.Time `json:"bookingDate,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BookingQuery when eager-loading is set.
	Edges          BookingEdges `json:"edges"`
	place_bookings *string
	room_bookings  *string
	user_bookings  *string
	selectValues   sql.SelectValues
}

// BookingEdges holds the relations/edges for other nodes in the graph.
type BookingEdges struct {
	// Room holds the value of the room edge.
	Room *Room `json:"room,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// RoomOrErr returns the Room value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookingEdges) RoomOrErr() (*Room, error) {
	if e.loadedTypes[0] {
		if e.Room == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: room.Label}
		}
		return e.Room, nil
	}
	return nil, &NotLoadedError{edge: "room"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookingEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Booking) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case booking.FieldID, booking.FieldStatus:
			values[i] = new(sql.NullString)
		case booking.FieldStartDate, booking.FieldEndDate, booking.FieldBookingDate:
			values[i] = new(sql.NullTime)
		case booking.ForeignKeys[0]: // place_bookings
			values[i] = new(sql.NullString)
		case booking.ForeignKeys[1]: // room_bookings
			values[i] = new(sql.NullString)
		case booking.ForeignKeys[2]: // user_bookings
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Booking fields.
func (b *Booking) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case booking.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				b.ID = value.String
			}
		case booking.FieldStartDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field startDate", values[i])
			} else if value.Valid {
				b.StartDate = value.Time
			}
		case booking.FieldEndDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field endDate", values[i])
			} else if value.Valid {
				b.EndDate = value.Time
			}
		case booking.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				b.Status = value.String
			}
		case booking.FieldBookingDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field bookingDate", values[i])
			} else if value.Valid {
				b.BookingDate = value.Time
			}
		case booking.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field place_bookings", values[i])
			} else if value.Valid {
				b.place_bookings = new(string)
				*b.place_bookings = value.String
			}
		case booking.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field room_bookings", values[i])
			} else if value.Valid {
				b.room_bookings = new(string)
				*b.room_bookings = value.String
			}
		case booking.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_bookings", values[i])
			} else if value.Valid {
				b.user_bookings = new(string)
				*b.user_bookings = value.String
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Booking.
// This includes values selected through modifiers, order, etc.
func (b *Booking) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// QueryRoom queries the "room" edge of the Booking entity.
func (b *Booking) QueryRoom() *RoomQuery {
	return NewBookingClient(b.config).QueryRoom(b)
}

// QueryUser queries the "user" edge of the Booking entity.
func (b *Booking) QueryUser() *UserQuery {
	return NewBookingClient(b.config).QueryUser(b)
}

// Update returns a builder for updating this Booking.
// Note that you need to call Booking.Unwrap() before calling this method if this Booking
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Booking) Update() *BookingUpdateOne {
	return NewBookingClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Booking entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Booking) Unwrap() *Booking {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Booking is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Booking) String() string {
	var builder strings.Builder
	builder.WriteString("Booking(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("startDate=")
	builder.WriteString(b.StartDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("endDate=")
	builder.WriteString(b.EndDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(b.Status)
	builder.WriteString(", ")
	builder.WriteString("bookingDate=")
	builder.WriteString(b.BookingDate.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Bookings is a parsable slice of Booking.
type Bookings []*Booking