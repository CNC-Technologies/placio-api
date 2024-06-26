// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"placio-app/ent/comment"
	"placio-app/ent/post"
	"placio-app/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Comment is the model entity for the Comment schema.
type Comment struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Content holds the value of the "Content" field.
	Content string `json:"Content,omitempty"`
	// CreatedAt holds the value of the "CreatedAt" field.
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
	// UpdatedAt holds the value of the "UpdatedAt" field.
	UpdatedAt time.Time `json:"UpdatedAt,omitempty"`
	// ParentCommentID holds the value of the "parentCommentID" field.
	ParentCommentID *string `json:"parentCommentID,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CommentQuery when eager-loading is set.
	Edges                CommentEdges `json:"edges"`
	event_event_comments *string
	post_comments        *string
	review_comments      *string
	user_comments        *string
	selectValues         sql.SelectValues
}

// CommentEdges holds the relations/edges for other nodes in the graph.
type CommentEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Post holds the value of the post edge.
	Post *Post `json:"post,omitempty"`
	// ParentComment holds the value of the parentComment edge.
	ParentComment *Comment `json:"parentComment,omitempty"`
	// Replies holds the value of the replies edge.
	Replies []*Comment `json:"replies,omitempty"`
	// Notifications holds the value of the notifications edge.
	Notifications []*Notification `json:"notifications,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CommentEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// PostOrErr returns the Post value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CommentEdges) PostOrErr() (*Post, error) {
	if e.loadedTypes[1] {
		if e.Post == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: post.Label}
		}
		return e.Post, nil
	}
	return nil, &NotLoadedError{edge: "post"}
}

// ParentCommentOrErr returns the ParentComment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CommentEdges) ParentCommentOrErr() (*Comment, error) {
	if e.loadedTypes[2] {
		if e.ParentComment == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: comment.Label}
		}
		return e.ParentComment, nil
	}
	return nil, &NotLoadedError{edge: "parentComment"}
}

// RepliesOrErr returns the Replies value or an error if the edge
// was not loaded in eager-loading.
func (e CommentEdges) RepliesOrErr() ([]*Comment, error) {
	if e.loadedTypes[3] {
		return e.Replies, nil
	}
	return nil, &NotLoadedError{edge: "replies"}
}

// NotificationsOrErr returns the Notifications value or an error if the edge
// was not loaded in eager-loading.
func (e CommentEdges) NotificationsOrErr() ([]*Notification, error) {
	if e.loadedTypes[4] {
		return e.Notifications, nil
	}
	return nil, &NotLoadedError{edge: "notifications"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Comment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case comment.FieldID, comment.FieldContent, comment.FieldParentCommentID:
			values[i] = new(sql.NullString)
		case comment.FieldCreatedAt, comment.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case comment.ForeignKeys[0]: // event_event_comments
			values[i] = new(sql.NullString)
		case comment.ForeignKeys[1]: // post_comments
			values[i] = new(sql.NullString)
		case comment.ForeignKeys[2]: // review_comments
			values[i] = new(sql.NullString)
		case comment.ForeignKeys[3]: // user_comments
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Comment fields.
func (c *Comment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case comment.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				c.ID = value.String
			}
		case comment.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Content", values[i])
			} else if value.Valid {
				c.Content = value.String
			}
		case comment.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedAt", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case comment.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field UpdatedAt", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case comment.FieldParentCommentID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field parentCommentID", values[i])
			} else if value.Valid {
				c.ParentCommentID = new(string)
				*c.ParentCommentID = value.String
			}
		case comment.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field event_event_comments", values[i])
			} else if value.Valid {
				c.event_event_comments = new(string)
				*c.event_event_comments = value.String
			}
		case comment.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field post_comments", values[i])
			} else if value.Valid {
				c.post_comments = new(string)
				*c.post_comments = value.String
			}
		case comment.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field review_comments", values[i])
			} else if value.Valid {
				c.review_comments = new(string)
				*c.review_comments = value.String
			}
		case comment.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_comments", values[i])
			} else if value.Valid {
				c.user_comments = new(string)
				*c.user_comments = value.String
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Comment.
// This includes values selected through modifiers, order, etc.
func (c *Comment) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Comment entity.
func (c *Comment) QueryUser() *UserQuery {
	return NewCommentClient(c.config).QueryUser(c)
}

// QueryPost queries the "post" edge of the Comment entity.
func (c *Comment) QueryPost() *PostQuery {
	return NewCommentClient(c.config).QueryPost(c)
}

// QueryParentComment queries the "parentComment" edge of the Comment entity.
func (c *Comment) QueryParentComment() *CommentQuery {
	return NewCommentClient(c.config).QueryParentComment(c)
}

// QueryReplies queries the "replies" edge of the Comment entity.
func (c *Comment) QueryReplies() *CommentQuery {
	return NewCommentClient(c.config).QueryReplies(c)
}

// QueryNotifications queries the "notifications" edge of the Comment entity.
func (c *Comment) QueryNotifications() *NotificationQuery {
	return NewCommentClient(c.config).QueryNotifications(c)
}

// Update returns a builder for updating this Comment.
// Note that you need to call Comment.Unwrap() before calling this method if this Comment
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Comment) Update() *CommentUpdateOne {
	return NewCommentClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Comment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Comment) Unwrap() *Comment {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Comment is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Comment) String() string {
	var builder strings.Builder
	builder.WriteString("Comment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("Content=")
	builder.WriteString(c.Content)
	builder.WriteString(", ")
	builder.WriteString("CreatedAt=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("UpdatedAt=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := c.ParentCommentID; v != nil {
		builder.WriteString("parentCommentID=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Comments is a parsable slice of Comment.
type Comments []*Comment
