// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"placio-app/ent/media"
	"placio-app/ent/post"
	"placio-app/ent/review"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Media is the model entity for the Media schema.
type Media struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// URL holds the value of the "URL" field.
	URL string `json:"URL,omitempty"`
	// image, gif, or video
	MediaType string `json:"MediaType,omitempty"`
	// CreatedAt holds the value of the "CreatedAt" field.
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
	// UpdatedAt holds the value of the "UpdatedAt" field.
	UpdatedAt time.Time `json:"UpdatedAt,omitempty"`
	// Count of likes for this media.
	LikeCount int `json:"likeCount,omitempty"`
	// Count of dislikes for this media.
	DislikeCount int `json:"dislikeCount,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MediaQuery when eager-loading is set.
	Edges           MediaEdges `json:"edges"`
	menu_item_media *string
	post_medias     *string
	review_medias   *string
	website_assets  *string
	selectValues    sql.SelectValues
}

// MediaEdges holds the relations/edges for other nodes in the graph.
type MediaEdges struct {
	// Post holds the value of the post edge.
	Post *Post `json:"post,omitempty"`
	// Review holds the value of the review edge.
	Review *Review `json:"review,omitempty"`
	// Categories holds the value of the categories edge.
	Categories []*Category `json:"categories,omitempty"`
	// Place holds the value of the place edge.
	Place []*Place `json:"place,omitempty"`
	// PlaceInventory holds the value of the place_inventory edge.
	PlaceInventory []*PlaceInventory `json:"place_inventory,omitempty"`
	// Menu holds the value of the menu edge.
	Menu []*Menu `json:"menu,omitempty"`
	// RoomCategory holds the value of the room_category edge.
	RoomCategory []*RoomCategory `json:"room_category,omitempty"`
	// Room holds the value of the room edge.
	Room []*Room `json:"room,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [8]bool
}

// PostOrErr returns the Post value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MediaEdges) PostOrErr() (*Post, error) {
	if e.loadedTypes[0] {
		if e.Post == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: post.Label}
		}
		return e.Post, nil
	}
	return nil, &NotLoadedError{edge: "post"}
}

// ReviewOrErr returns the Review value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MediaEdges) ReviewOrErr() (*Review, error) {
	if e.loadedTypes[1] {
		if e.Review == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: review.Label}
		}
		return e.Review, nil
	}
	return nil, &NotLoadedError{edge: "review"}
}

// CategoriesOrErr returns the Categories value or an error if the edge
// was not loaded in eager-loading.
func (e MediaEdges) CategoriesOrErr() ([]*Category, error) {
	if e.loadedTypes[2] {
		return e.Categories, nil
	}
	return nil, &NotLoadedError{edge: "categories"}
}

// PlaceOrErr returns the Place value or an error if the edge
// was not loaded in eager-loading.
func (e MediaEdges) PlaceOrErr() ([]*Place, error) {
	if e.loadedTypes[3] {
		return e.Place, nil
	}
	return nil, &NotLoadedError{edge: "place"}
}

// PlaceInventoryOrErr returns the PlaceInventory value or an error if the edge
// was not loaded in eager-loading.
func (e MediaEdges) PlaceInventoryOrErr() ([]*PlaceInventory, error) {
	if e.loadedTypes[4] {
		return e.PlaceInventory, nil
	}
	return nil, &NotLoadedError{edge: "place_inventory"}
}

// MenuOrErr returns the Menu value or an error if the edge
// was not loaded in eager-loading.
func (e MediaEdges) MenuOrErr() ([]*Menu, error) {
	if e.loadedTypes[5] {
		return e.Menu, nil
	}
	return nil, &NotLoadedError{edge: "menu"}
}

// RoomCategoryOrErr returns the RoomCategory value or an error if the edge
// was not loaded in eager-loading.
func (e MediaEdges) RoomCategoryOrErr() ([]*RoomCategory, error) {
	if e.loadedTypes[6] {
		return e.RoomCategory, nil
	}
	return nil, &NotLoadedError{edge: "room_category"}
}

// RoomOrErr returns the Room value or an error if the edge
// was not loaded in eager-loading.
func (e MediaEdges) RoomOrErr() ([]*Room, error) {
	if e.loadedTypes[7] {
		return e.Room, nil
	}
	return nil, &NotLoadedError{edge: "room"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Media) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case media.FieldLikeCount, media.FieldDislikeCount:
			values[i] = new(sql.NullInt64)
		case media.FieldID, media.FieldURL, media.FieldMediaType:
			values[i] = new(sql.NullString)
		case media.FieldCreatedAt, media.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case media.ForeignKeys[0]: // menu_item_media
			values[i] = new(sql.NullString)
		case media.ForeignKeys[1]: // post_medias
			values[i] = new(sql.NullString)
		case media.ForeignKeys[2]: // review_medias
			values[i] = new(sql.NullString)
		case media.ForeignKeys[3]: // website_assets
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Media fields.
func (m *Media) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case media.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				m.ID = value.String
			}
		case media.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field URL", values[i])
			} else if value.Valid {
				m.URL = value.String
			}
		case media.FieldMediaType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field MediaType", values[i])
			} else if value.Valid {
				m.MediaType = value.String
			}
		case media.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedAt", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		case media.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field UpdatedAt", values[i])
			} else if value.Valid {
				m.UpdatedAt = value.Time
			}
		case media.FieldLikeCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field likeCount", values[i])
			} else if value.Valid {
				m.LikeCount = int(value.Int64)
			}
		case media.FieldDislikeCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field dislikeCount", values[i])
			} else if value.Valid {
				m.DislikeCount = int(value.Int64)
			}
		case media.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field menu_item_media", values[i])
			} else if value.Valid {
				m.menu_item_media = new(string)
				*m.menu_item_media = value.String
			}
		case media.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field post_medias", values[i])
			} else if value.Valid {
				m.post_medias = new(string)
				*m.post_medias = value.String
			}
		case media.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field review_medias", values[i])
			} else if value.Valid {
				m.review_medias = new(string)
				*m.review_medias = value.String
			}
		case media.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field website_assets", values[i])
			} else if value.Valid {
				m.website_assets = new(string)
				*m.website_assets = value.String
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Media.
// This includes values selected through modifiers, order, etc.
func (m *Media) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryPost queries the "post" edge of the Media entity.
func (m *Media) QueryPost() *PostQuery {
	return NewMediaClient(m.config).QueryPost(m)
}

// QueryReview queries the "review" edge of the Media entity.
func (m *Media) QueryReview() *ReviewQuery {
	return NewMediaClient(m.config).QueryReview(m)
}

// QueryCategories queries the "categories" edge of the Media entity.
func (m *Media) QueryCategories() *CategoryQuery {
	return NewMediaClient(m.config).QueryCategories(m)
}

// QueryPlace queries the "place" edge of the Media entity.
func (m *Media) QueryPlace() *PlaceQuery {
	return NewMediaClient(m.config).QueryPlace(m)
}

// QueryPlaceInventory queries the "place_inventory" edge of the Media entity.
func (m *Media) QueryPlaceInventory() *PlaceInventoryQuery {
	return NewMediaClient(m.config).QueryPlaceInventory(m)
}

// QueryMenu queries the "menu" edge of the Media entity.
func (m *Media) QueryMenu() *MenuQuery {
	return NewMediaClient(m.config).QueryMenu(m)
}

// QueryRoomCategory queries the "room_category" edge of the Media entity.
func (m *Media) QueryRoomCategory() *RoomCategoryQuery {
	return NewMediaClient(m.config).QueryRoomCategory(m)
}

// QueryRoom queries the "room" edge of the Media entity.
func (m *Media) QueryRoom() *RoomQuery {
	return NewMediaClient(m.config).QueryRoom(m)
}

// Update returns a builder for updating this Media.
// Note that you need to call Media.Unwrap() before calling this method if this Media
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Media) Update() *MediaUpdateOne {
	return NewMediaClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Media entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Media) Unwrap() *Media {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Media is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Media) String() string {
	var builder strings.Builder
	builder.WriteString("Media(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("URL=")
	builder.WriteString(m.URL)
	builder.WriteString(", ")
	builder.WriteString("MediaType=")
	builder.WriteString(m.MediaType)
	builder.WriteString(", ")
	builder.WriteString("CreatedAt=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("UpdatedAt=")
	builder.WriteString(m.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("likeCount=")
	builder.WriteString(fmt.Sprintf("%v", m.LikeCount))
	builder.WriteString(", ")
	builder.WriteString("dislikeCount=")
	builder.WriteString(fmt.Sprintf("%v", m.DislikeCount))
	builder.WriteByte(')')
	return builder.String()
}

// MediaSlice is a parsable slice of Media.
type MediaSlice []*Media
