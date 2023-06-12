// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"placio-app/ent/business"
	"placio-app/ent/post"
	"placio-app/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Post is the model entity for the Post schema.
type Post struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Content holds the value of the "Content" field.
	Content string `json:"Content,omitempty"`
	// CreatedAt holds the value of the "CreatedAt" field.
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
	// UpdatedAt holds the value of the "UpdatedAt" field.
	UpdatedAt time.Time `json:"UpdatedAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PostQuery when eager-loading is set.
	Edges          PostEdges `json:"edges"`
	business_posts *int
	user_posts     *int
	selectValues   sql.SelectValues
}

// PostEdges holds the relations/edges for other nodes in the graph.
type PostEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// BusinessAccount holds the value of the business_account edge.
	BusinessAccount *Business `json:"business_account,omitempty"`
	// Medias holds the value of the medias edge.
	Medias []*Media `json:"medias,omitempty"`
	// Comments holds the value of the comments edge.
	Comments []*Comment `json:"comments,omitempty"`
	// Likes holds the value of the likes edge.
	Likes []*Like `json:"likes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PostEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// BusinessAccountOrErr returns the BusinessAccount value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PostEdges) BusinessAccountOrErr() (*Business, error) {
	if e.loadedTypes[1] {
		if e.BusinessAccount == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: business.Label}
		}
		return e.BusinessAccount, nil
	}
	return nil, &NotLoadedError{edge: "business_account"}
}

// MediasOrErr returns the Medias value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) MediasOrErr() ([]*Media, error) {
	if e.loadedTypes[2] {
		return e.Medias, nil
	}
	return nil, &NotLoadedError{edge: "medias"}
}

// CommentsOrErr returns the Comments value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) CommentsOrErr() ([]*Comment, error) {
	if e.loadedTypes[3] {
		return e.Comments, nil
	}
	return nil, &NotLoadedError{edge: "comments"}
}

// LikesOrErr returns the Likes value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) LikesOrErr() ([]*Like, error) {
	if e.loadedTypes[4] {
		return e.Likes, nil
	}
	return nil, &NotLoadedError{edge: "likes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Post) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case post.FieldID:
			values[i] = new(sql.NullInt64)
		case post.FieldContent:
			values[i] = new(sql.NullString)
		case post.FieldCreatedAt, post.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case post.ForeignKeys[0]: // business_posts
			values[i] = new(sql.NullInt64)
		case post.ForeignKeys[1]: // user_posts
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Post fields.
func (po *Post) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case post.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			po.ID = int(value.Int64)
		case post.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Content", values[i])
			} else if value.Valid {
				po.Content = value.String
			}
		case post.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedAt", values[i])
			} else if value.Valid {
				po.CreatedAt = value.Time
			}
		case post.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field UpdatedAt", values[i])
			} else if value.Valid {
				po.UpdatedAt = value.Time
			}
		case post.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field business_posts", value)
			} else if value.Valid {
				po.business_posts = new(int)
				*po.business_posts = int(value.Int64)
			}
		case post.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_posts", value)
			} else if value.Valid {
				po.user_posts = new(int)
				*po.user_posts = int(value.Int64)
			}
		default:
			po.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Post.
// This includes values selected through modifiers, order, etc.
func (po *Post) Value(name string) (ent.Value, error) {
	return po.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Post entity.
func (po *Post) QueryUser() *UserQuery {
	return NewPostClient(po.config).QueryUser(po)
}

// QueryBusinessAccount queries the "business_account" edge of the Post entity.
func (po *Post) QueryBusinessAccount() *BusinessQuery {
	return NewPostClient(po.config).QueryBusinessAccount(po)
}

// QueryMedias queries the "medias" edge of the Post entity.
func (po *Post) QueryMedias() *MediaQuery {
	return NewPostClient(po.config).QueryMedias(po)
}

// QueryComments queries the "comments" edge of the Post entity.
func (po *Post) QueryComments() *CommentQuery {
	return NewPostClient(po.config).QueryComments(po)
}

// QueryLikes queries the "likes" edge of the Post entity.
func (po *Post) QueryLikes() *LikeQuery {
	return NewPostClient(po.config).QueryLikes(po)
}

// Update returns a builder for updating this Post.
// Note that you need to call Post.Unwrap() before calling this method if this Post
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *Post) Update() *PostUpdateOne {
	return NewPostClient(po.config).UpdateOne(po)
}

// Unwrap unwraps the Post entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *Post) Unwrap() *Post {
	_tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: Post is not a transactional entity")
	}
	po.config.driver = _tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *Post) String() string {
	var builder strings.Builder
	builder.WriteString("Post(")
	builder.WriteString(fmt.Sprintf("id=%v, ", po.ID))
	builder.WriteString("Content=")
	builder.WriteString(po.Content)
	builder.WriteString(", ")
	builder.WriteString("CreatedAt=")
	builder.WriteString(po.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("UpdatedAt=")
	builder.WriteString(po.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Posts is a parsable slice of Post.
type Posts []*Post
