// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"placio-app/ent/business"
	"placio-app/ent/user"
	"placio-app/ent/userfollowbusiness"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// UserFollowBusiness is the model entity for the UserFollowBusiness schema.
type UserFollowBusiness struct {
	config
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserFollowBusinessQuery when eager-loading is set.
	Edges                    UserFollowBusinessEdges `json:"edges"`
	business_follower_users  *string
	user_followed_businesses *string
	selectValues             sql.SelectValues
}

// UserFollowBusinessEdges holds the relations/edges for other nodes in the graph.
type UserFollowBusinessEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Business holds the value of the business edge.
	Business *Business `json:"business,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserFollowBusinessEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// BusinessOrErr returns the Business value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserFollowBusinessEdges) BusinessOrErr() (*Business, error) {
	if e.loadedTypes[1] {
		if e.Business == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: business.Label}
		}
		return e.Business, nil
	}
	return nil, &NotLoadedError{edge: "business"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserFollowBusiness) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userfollowbusiness.FieldID:
			values[i] = new(sql.NullString)
		case userfollowbusiness.ForeignKeys[0]: // business_follower_users
			values[i] = new(sql.NullString)
		case userfollowbusiness.ForeignKeys[1]: // user_followed_businesses
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserFollowBusiness fields.
func (ufb *UserFollowBusiness) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userfollowbusiness.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ufb.ID = value.String
			}
		case userfollowbusiness.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field business_follower_users", values[i])
			} else if value.Valid {
				ufb.business_follower_users = new(string)
				*ufb.business_follower_users = value.String
			}
		case userfollowbusiness.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_followed_businesses", values[i])
			} else if value.Valid {
				ufb.user_followed_businesses = new(string)
				*ufb.user_followed_businesses = value.String
			}
		default:
			ufb.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserFollowBusiness.
// This includes values selected through modifiers, order, etc.
func (ufb *UserFollowBusiness) Value(name string) (ent.Value, error) {
	return ufb.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the UserFollowBusiness entity.
func (ufb *UserFollowBusiness) QueryUser() *UserQuery {
	return NewUserFollowBusinessClient(ufb.config).QueryUser(ufb)
}

// QueryBusiness queries the "business" edge of the UserFollowBusiness entity.
func (ufb *UserFollowBusiness) QueryBusiness() *BusinessQuery {
	return NewUserFollowBusinessClient(ufb.config).QueryBusiness(ufb)
}

// Update returns a builder for updating this UserFollowBusiness.
// Note that you need to call UserFollowBusiness.Unwrap() before calling this method if this UserFollowBusiness
// was returned from a transaction, and the transaction was committed or rolled back.
func (ufb *UserFollowBusiness) Update() *UserFollowBusinessUpdateOne {
	return NewUserFollowBusinessClient(ufb.config).UpdateOne(ufb)
}

// Unwrap unwraps the UserFollowBusiness entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ufb *UserFollowBusiness) Unwrap() *UserFollowBusiness {
	_tx, ok := ufb.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserFollowBusiness is not a transactional entity")
	}
	ufb.config.driver = _tx.drv
	return ufb
}

// String implements the fmt.Stringer.
func (ufb *UserFollowBusiness) String() string {
	var builder strings.Builder
	builder.WriteString("UserFollowBusiness(")
	builder.WriteString(fmt.Sprintf("id=%v", ufb.ID))
	builder.WriteByte(')')
	return builder.String()
}

// UserFollowBusinesses is a parsable slice of UserFollowBusiness.
type UserFollowBusinesses []*UserFollowBusiness
