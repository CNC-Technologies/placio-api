// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"placio-app/ent/accountsettings"
	"placio-app/ent/accountwallet"
	"placio-app/ent/business"
	"placio-app/ent/website"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Business is the model entity for the Business schema.
type Business struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Picture holds the value of the "picture" field.
	Picture string `json:"picture,omitempty"`
	// CoverImage holds the value of the "cover_image" field.
	CoverImage string `json:"cover_image,omitempty"`
	// Website holds the value of the "website" field.
	Website string `json:"website,omitempty"`
	// Location holds the value of the "location" field.
	Location string `json:"location,omitempty"`
	// Longitude holds the value of the "longitude" field.
	Longitude string `json:"longitude,omitempty"`
	// MapCoordinates holds the value of the "map_coordinates" field.
	MapCoordinates map[string]interface{} `json:"map_coordinates,omitempty"`
	// Latitude holds the value of the "latitude" field.
	Latitude string `json:"latitude,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// BusinessSettings holds the value of the "business_settings" field.
	BusinessSettings map[string]interface{} `json:"business_settings,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// SearchText holds the value of the "search_text" field.
	SearchText string `json:"search_text,omitempty"`
	// RelevanceScore holds the value of the "relevance_score" field.
	RelevanceScore float64 `json:"relevance_score,omitempty"`
	// FollowerCount holds the value of the "follower_count" field.
	FollowerCount int `json:"follower_count,omitempty"`
	// FollowingCount holds the value of the "following_count" field.
	FollowingCount int `json:"following_count,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BusinessQuery when eager-loading is set.
	Edges        BusinessEdges `json:"edges"`
	selectValues sql.SelectValues
}

// BusinessEdges holds the relations/edges for other nodes in the graph.
type BusinessEdges struct {
	// UserBusinesses holds the value of the userBusinesses edge.
	UserBusinesses []*UserBusiness `json:"userBusinesses,omitempty"`
	// BusinessAccountSettings holds the value of the business_account_settings edge.
	BusinessAccountSettings *AccountSettings `json:"business_account_settings,omitempty"`
	// Posts holds the value of the posts edge.
	Posts []*Post `json:"posts,omitempty"`
	// FollowedUsers holds the value of the followedUsers edge.
	FollowedUsers []*BusinessFollowUser `json:"followedUsers,omitempty"`
	// FollowerUsers holds the value of the followerUsers edge.
	FollowerUsers []*UserFollowBusiness `json:"followerUsers,omitempty"`
	// FollowedBusinesses holds the value of the followedBusinesses edge.
	FollowedBusinesses []*BusinessFollowBusiness `json:"followedBusinesses,omitempty"`
	// FollowerBusinesses holds the value of the followerBusinesses edge.
	FollowerBusinesses []*BusinessFollowBusiness `json:"followerBusinesses,omitempty"`
	// Places holds the value of the places edge.
	Places []*Place `json:"places,omitempty"`
	// Categories holds the value of the categories edge.
	Categories []*Category `json:"categories,omitempty"`
	// CategoryAssignments holds the value of the categoryAssignments edge.
	CategoryAssignments []*CategoryAssignment `json:"categoryAssignments,omitempty"`
	// Events holds the value of the events edge.
	Events []*Event `json:"events,omitempty"`
	// BusinessFollowEvents holds the value of the businessFollowEvents edge.
	BusinessFollowEvents []*BusinessFollowEvent `json:"businessFollowEvents,omitempty"`
	// Faqs holds the value of the faqs edge.
	Faqs []*FAQ `json:"faqs,omitempty"`
	// Ratings holds the value of the ratings edge.
	Ratings []*Rating `json:"ratings,omitempty"`
	// PlaceInventories holds the value of the place_inventories edge.
	PlaceInventories []*PlaceInventory `json:"place_inventories,omitempty"`
	// Websites holds the value of the websites edge.
	Websites *Website `json:"websites,omitempty"`
	// Notifications holds the value of the notifications edge.
	Notifications []*Notification `json:"notifications,omitempty"`
	// Wallet holds the value of the wallet edge.
	Wallet *AccountWallet `json:"wallet,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [18]bool
}

// UserBusinessesOrErr returns the UserBusinesses value or an error if the edge
// was not loaded in eager-loading.
func (e BusinessEdges) UserBusinessesOrErr() ([]*UserBusiness, error) {
	if e.loadedTypes[0] {
		return e.UserBusinesses, nil
	}
	return nil, &NotLoadedError{edge: "userBusinesses"}
}

// BusinessAccountSettingsOrErr returns the BusinessAccountSettings value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BusinessEdges) BusinessAccountSettingsOrErr() (*AccountSettings, error) {
	if e.loadedTypes[1] {
		if e.BusinessAccountSettings == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: accountsettings.Label}
		}
		return e.BusinessAccountSettings, nil
	}
	return nil, &NotLoadedError{edge: "business_account_settings"}
}

// PostsOrErr returns the Posts value or an error if the edge
// was not loaded in eager-loading.
func (e BusinessEdges) PostsOrErr() ([]*Post, error) {
	if e.loadedTypes[2] {
		return e.Posts, nil
	}
	return nil, &NotLoadedError{edge: "posts"}
}

// FollowedUsersOrErr returns the FollowedUsers value or an error if the edge
// was not loaded in eager-loading.
func (e BusinessEdges) FollowedUsersOrErr() ([]*BusinessFollowUser, error) {
	if e.loadedTypes[3] {
		return e.FollowedUsers, nil
	}
	return nil, &NotLoadedError{edge: "followedUsers"}
}

// FollowerUsersOrErr returns the FollowerUsers value or an error if the edge
// was not loaded in eager-loading.
func (e BusinessEdges) FollowerUsersOrErr() ([]*UserFollowBusiness, error) {
	if e.loadedTypes[4] {
		return e.FollowerUsers, nil
	}
	return nil, &NotLoadedError{edge: "followerUsers"}
}

// FollowedBusinessesOrErr returns the FollowedBusinesses value or an error if the edge
// was not loaded in eager-loading.
func (e BusinessEdges) FollowedBusinessesOrErr() ([]*BusinessFollowBusiness, error) {
	if e.loadedTypes[5] {
		return e.FollowedBusinesses, nil
	}
	return nil, &NotLoadedError{edge: "followedBusinesses"}
}

// FollowerBusinessesOrErr returns the FollowerBusinesses value or an error if the edge
// was not loaded in eager-loading.
func (e BusinessEdges) FollowerBusinessesOrErr() ([]*BusinessFollowBusiness, error) {
	if e.loadedTypes[6] {
		return e.FollowerBusinesses, nil
	}
	return nil, &NotLoadedError{edge: "followerBusinesses"}
}

// PlacesOrErr returns the Places value or an error if the edge
// was not loaded in eager-loading.
func (e BusinessEdges) PlacesOrErr() ([]*Place, error) {
	if e.loadedTypes[7] {
		return e.Places, nil
	}
	return nil, &NotLoadedError{edge: "places"}
}

// CategoriesOrErr returns the Categories value or an error if the edge
// was not loaded in eager-loading.
func (e BusinessEdges) CategoriesOrErr() ([]*Category, error) {
	if e.loadedTypes[8] {
		return e.Categories, nil
	}
	return nil, &NotLoadedError{edge: "categories"}
}

// CategoryAssignmentsOrErr returns the CategoryAssignments value or an error if the edge
// was not loaded in eager-loading.
func (e BusinessEdges) CategoryAssignmentsOrErr() ([]*CategoryAssignment, error) {
	if e.loadedTypes[9] {
		return e.CategoryAssignments, nil
	}
	return nil, &NotLoadedError{edge: "categoryAssignments"}
}

// EventsOrErr returns the Events value or an error if the edge
// was not loaded in eager-loading.
func (e BusinessEdges) EventsOrErr() ([]*Event, error) {
	if e.loadedTypes[10] {
		return e.Events, nil
	}
	return nil, &NotLoadedError{edge: "events"}
}

// BusinessFollowEventsOrErr returns the BusinessFollowEvents value or an error if the edge
// was not loaded in eager-loading.
func (e BusinessEdges) BusinessFollowEventsOrErr() ([]*BusinessFollowEvent, error) {
	if e.loadedTypes[11] {
		return e.BusinessFollowEvents, nil
	}
	return nil, &NotLoadedError{edge: "businessFollowEvents"}
}

// FaqsOrErr returns the Faqs value or an error if the edge
// was not loaded in eager-loading.
func (e BusinessEdges) FaqsOrErr() ([]*FAQ, error) {
	if e.loadedTypes[12] {
		return e.Faqs, nil
	}
	return nil, &NotLoadedError{edge: "faqs"}
}

// RatingsOrErr returns the Ratings value or an error if the edge
// was not loaded in eager-loading.
func (e BusinessEdges) RatingsOrErr() ([]*Rating, error) {
	if e.loadedTypes[13] {
		return e.Ratings, nil
	}
	return nil, &NotLoadedError{edge: "ratings"}
}

// PlaceInventoriesOrErr returns the PlaceInventories value or an error if the edge
// was not loaded in eager-loading.
func (e BusinessEdges) PlaceInventoriesOrErr() ([]*PlaceInventory, error) {
	if e.loadedTypes[14] {
		return e.PlaceInventories, nil
	}
	return nil, &NotLoadedError{edge: "place_inventories"}
}

// WebsitesOrErr returns the Websites value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BusinessEdges) WebsitesOrErr() (*Website, error) {
	if e.loadedTypes[15] {
		if e.Websites == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: website.Label}
		}
		return e.Websites, nil
	}
	return nil, &NotLoadedError{edge: "websites"}
}

// NotificationsOrErr returns the Notifications value or an error if the edge
// was not loaded in eager-loading.
func (e BusinessEdges) NotificationsOrErr() ([]*Notification, error) {
	if e.loadedTypes[16] {
		return e.Notifications, nil
	}
	return nil, &NotLoadedError{edge: "notifications"}
}

// WalletOrErr returns the Wallet value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BusinessEdges) WalletOrErr() (*AccountWallet, error) {
	if e.loadedTypes[17] {
		if e.Wallet == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: accountwallet.Label}
		}
		return e.Wallet, nil
	}
	return nil, &NotLoadedError{edge: "wallet"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Business) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case business.FieldMapCoordinates, business.FieldBusinessSettings:
			values[i] = new([]byte)
		case business.FieldRelevanceScore:
			values[i] = new(sql.NullFloat64)
		case business.FieldFollowerCount, business.FieldFollowingCount:
			values[i] = new(sql.NullInt64)
		case business.FieldID, business.FieldName, business.FieldDescription, business.FieldPicture, business.FieldCoverImage, business.FieldWebsite, business.FieldLocation, business.FieldLongitude, business.FieldLatitude, business.FieldEmail, business.FieldPhone, business.FieldURL, business.FieldSearchText:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Business fields.
func (b *Business) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case business.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				b.ID = value.String
			}
		case business.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				b.Name = value.String
			}
		case business.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				b.Description = value.String
			}
		case business.FieldPicture:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field picture", values[i])
			} else if value.Valid {
				b.Picture = value.String
			}
		case business.FieldCoverImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cover_image", values[i])
			} else if value.Valid {
				b.CoverImage = value.String
			}
		case business.FieldWebsite:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field website", values[i])
			} else if value.Valid {
				b.Website = value.String
			}
		case business.FieldLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location", values[i])
			} else if value.Valid {
				b.Location = value.String
			}
		case business.FieldLongitude:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field longitude", values[i])
			} else if value.Valid {
				b.Longitude = value.String
			}
		case business.FieldMapCoordinates:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field map_coordinates", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &b.MapCoordinates); err != nil {
					return fmt.Errorf("unmarshal field map_coordinates: %w", err)
				}
			}
		case business.FieldLatitude:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field latitude", values[i])
			} else if value.Valid {
				b.Latitude = value.String
			}
		case business.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				b.Email = value.String
			}
		case business.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				b.Phone = value.String
			}
		case business.FieldBusinessSettings:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field business_settings", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &b.BusinessSettings); err != nil {
					return fmt.Errorf("unmarshal field business_settings: %w", err)
				}
			}
		case business.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				b.URL = value.String
			}
		case business.FieldSearchText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field search_text", values[i])
			} else if value.Valid {
				b.SearchText = value.String
			}
		case business.FieldRelevanceScore:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field relevance_score", values[i])
			} else if value.Valid {
				b.RelevanceScore = value.Float64
			}
		case business.FieldFollowerCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field follower_count", values[i])
			} else if value.Valid {
				b.FollowerCount = int(value.Int64)
			}
		case business.FieldFollowingCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field following_count", values[i])
			} else if value.Valid {
				b.FollowingCount = int(value.Int64)
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Business.
// This includes values selected through modifiers, order, etc.
func (b *Business) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// QueryUserBusinesses queries the "userBusinesses" edge of the Business entity.
func (b *Business) QueryUserBusinesses() *UserBusinessQuery {
	return NewBusinessClient(b.config).QueryUserBusinesses(b)
}

// QueryBusinessAccountSettings queries the "business_account_settings" edge of the Business entity.
func (b *Business) QueryBusinessAccountSettings() *AccountSettingsQuery {
	return NewBusinessClient(b.config).QueryBusinessAccountSettings(b)
}

// QueryPosts queries the "posts" edge of the Business entity.
func (b *Business) QueryPosts() *PostQuery {
	return NewBusinessClient(b.config).QueryPosts(b)
}

// QueryFollowedUsers queries the "followedUsers" edge of the Business entity.
func (b *Business) QueryFollowedUsers() *BusinessFollowUserQuery {
	return NewBusinessClient(b.config).QueryFollowedUsers(b)
}

// QueryFollowerUsers queries the "followerUsers" edge of the Business entity.
func (b *Business) QueryFollowerUsers() *UserFollowBusinessQuery {
	return NewBusinessClient(b.config).QueryFollowerUsers(b)
}

// QueryFollowedBusinesses queries the "followedBusinesses" edge of the Business entity.
func (b *Business) QueryFollowedBusinesses() *BusinessFollowBusinessQuery {
	return NewBusinessClient(b.config).QueryFollowedBusinesses(b)
}

// QueryFollowerBusinesses queries the "followerBusinesses" edge of the Business entity.
func (b *Business) QueryFollowerBusinesses() *BusinessFollowBusinessQuery {
	return NewBusinessClient(b.config).QueryFollowerBusinesses(b)
}

// QueryPlaces queries the "places" edge of the Business entity.
func (b *Business) QueryPlaces() *PlaceQuery {
	return NewBusinessClient(b.config).QueryPlaces(b)
}

// QueryCategories queries the "categories" edge of the Business entity.
func (b *Business) QueryCategories() *CategoryQuery {
	return NewBusinessClient(b.config).QueryCategories(b)
}

// QueryCategoryAssignments queries the "categoryAssignments" edge of the Business entity.
func (b *Business) QueryCategoryAssignments() *CategoryAssignmentQuery {
	return NewBusinessClient(b.config).QueryCategoryAssignments(b)
}

// QueryEvents queries the "events" edge of the Business entity.
func (b *Business) QueryEvents() *EventQuery {
	return NewBusinessClient(b.config).QueryEvents(b)
}

// QueryBusinessFollowEvents queries the "businessFollowEvents" edge of the Business entity.
func (b *Business) QueryBusinessFollowEvents() *BusinessFollowEventQuery {
	return NewBusinessClient(b.config).QueryBusinessFollowEvents(b)
}

// QueryFaqs queries the "faqs" edge of the Business entity.
func (b *Business) QueryFaqs() *FAQQuery {
	return NewBusinessClient(b.config).QueryFaqs(b)
}

// QueryRatings queries the "ratings" edge of the Business entity.
func (b *Business) QueryRatings() *RatingQuery {
	return NewBusinessClient(b.config).QueryRatings(b)
}

// QueryPlaceInventories queries the "place_inventories" edge of the Business entity.
func (b *Business) QueryPlaceInventories() *PlaceInventoryQuery {
	return NewBusinessClient(b.config).QueryPlaceInventories(b)
}

// QueryWebsites queries the "websites" edge of the Business entity.
func (b *Business) QueryWebsites() *WebsiteQuery {
	return NewBusinessClient(b.config).QueryWebsites(b)
}

// QueryNotifications queries the "notifications" edge of the Business entity.
func (b *Business) QueryNotifications() *NotificationQuery {
	return NewBusinessClient(b.config).QueryNotifications(b)
}

// QueryWallet queries the "wallet" edge of the Business entity.
func (b *Business) QueryWallet() *AccountWalletQuery {
	return NewBusinessClient(b.config).QueryWallet(b)
}

// Update returns a builder for updating this Business.
// Note that you need to call Business.Unwrap() before calling this method if this Business
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Business) Update() *BusinessUpdateOne {
	return NewBusinessClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Business entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Business) Unwrap() *Business {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Business is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Business) String() string {
	var builder strings.Builder
	builder.WriteString("Business(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("name=")
	builder.WriteString(b.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(b.Description)
	builder.WriteString(", ")
	builder.WriteString("picture=")
	builder.WriteString(b.Picture)
	builder.WriteString(", ")
	builder.WriteString("cover_image=")
	builder.WriteString(b.CoverImage)
	builder.WriteString(", ")
	builder.WriteString("website=")
	builder.WriteString(b.Website)
	builder.WriteString(", ")
	builder.WriteString("location=")
	builder.WriteString(b.Location)
	builder.WriteString(", ")
	builder.WriteString("longitude=")
	builder.WriteString(b.Longitude)
	builder.WriteString(", ")
	builder.WriteString("map_coordinates=")
	builder.WriteString(fmt.Sprintf("%v", b.MapCoordinates))
	builder.WriteString(", ")
	builder.WriteString("latitude=")
	builder.WriteString(b.Latitude)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(b.Email)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(b.Phone)
	builder.WriteString(", ")
	builder.WriteString("business_settings=")
	builder.WriteString(fmt.Sprintf("%v", b.BusinessSettings))
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(b.URL)
	builder.WriteString(", ")
	builder.WriteString("search_text=")
	builder.WriteString(b.SearchText)
	builder.WriteString(", ")
	builder.WriteString("relevance_score=")
	builder.WriteString(fmt.Sprintf("%v", b.RelevanceScore))
	builder.WriteString(", ")
	builder.WriteString("follower_count=")
	builder.WriteString(fmt.Sprintf("%v", b.FollowerCount))
	builder.WriteString(", ")
	builder.WriteString("following_count=")
	builder.WriteString(fmt.Sprintf("%v", b.FollowingCount))
	builder.WriteByte(')')
	return builder.String()
}

// Businesses is a parsable slice of Business.
type Businesses []*Business
