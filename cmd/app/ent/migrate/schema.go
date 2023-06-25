// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountSettingsColumns holds the columns for the "account_settings" table.
	AccountSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 36},
		{Name: "two_factor_authentication", Type: field.TypeBool},
		{Name: "blocked_users", Type: field.TypeJSON},
		{Name: "muted_users", Type: field.TypeJSON},
		{Name: "business_business_account_settings", Type: field.TypeString, Unique: true, Size: 36},
	}
	// AccountSettingsTable holds the schema information for the "account_settings" table.
	AccountSettingsTable = &schema.Table{
		Name:       "account_settings",
		Columns:    AccountSettingsColumns,
		PrimaryKey: []*schema.Column{AccountSettingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "account_settings_businesses_business_account_settings",
				Columns:    []*schema.Column{AccountSettingsColumns[4]},
				RefColumns: []*schema.Column{BusinessesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// AmenitiesColumns holds the columns for the "amenities" table.
	AmenitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 36},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "icon", Type: field.TypeString},
	}
	// AmenitiesTable holds the schema information for the "amenities" table.
	AmenitiesTable = &schema.Table{
		Name:       "amenities",
		Columns:    AmenitiesColumns,
		PrimaryKey: []*schema.Column{AmenitiesColumns[0]},
	}
	// BookingsColumns holds the columns for the "bookings" table.
	BookingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 36},
		{Name: "start_date", Type: field.TypeTime},
		{Name: "end_date", Type: field.TypeTime},
		{Name: "status", Type: field.TypeString},
		{Name: "booking_date", Type: field.TypeTime},
		{Name: "place_bookings", Type: field.TypeString, Nullable: true, Size: 36},
		{Name: "room_bookings", Type: field.TypeString, Nullable: true, Size: 36},
		{Name: "user_bookings", Type: field.TypeString, Nullable: true, Size: 36},
	}
	// BookingsTable holds the schema information for the "bookings" table.
	BookingsTable = &schema.Table{
		Name:       "bookings",
		Columns:    BookingsColumns,
		PrimaryKey: []*schema.Column{BookingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "bookings_places_bookings",
				Columns:    []*schema.Column{BookingsColumns[5]},
				RefColumns: []*schema.Column{PlacesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "bookings_rooms_bookings",
				Columns:    []*schema.Column{BookingsColumns[6]},
				RefColumns: []*schema.Column{RoomsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "bookings_users_bookings",
				Columns:    []*schema.Column{BookingsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BusinessesColumns holds the columns for the "businesses" table.
	BusinessesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 36},
		{Name: "name", Type: field.TypeString},
		{Name: "search_text", Type: field.TypeString, Nullable: true},
		{Name: "relevance_score", Type: field.TypeFloat64, Nullable: true},
	}
	// BusinessesTable holds the schema information for the "businesses" table.
	BusinessesTable = &schema.Table{
		Name:       "businesses",
		Columns:    BusinessesColumns,
		PrimaryKey: []*schema.Column{BusinessesColumns[0]},
	}
	// BusinessFollowBusinessesColumns holds the columns for the "business_follow_businesses" table.
	BusinessFollowBusinessesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "business_followed_businesses", Type: field.TypeString, Nullable: true, Size: 36},
		{Name: "business_follower_businesses", Type: field.TypeString, Nullable: true, Size: 36},
	}
	// BusinessFollowBusinessesTable holds the schema information for the "business_follow_businesses" table.
	BusinessFollowBusinessesTable = &schema.Table{
		Name:       "business_follow_businesses",
		Columns:    BusinessFollowBusinessesColumns,
		PrimaryKey: []*schema.Column{BusinessFollowBusinessesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "business_follow_businesses_businesses_followedBusinesses",
				Columns:    []*schema.Column{BusinessFollowBusinessesColumns[1]},
				RefColumns: []*schema.Column{BusinessesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "business_follow_businesses_businesses_followerBusinesses",
				Columns:    []*schema.Column{BusinessFollowBusinessesColumns[2]},
				RefColumns: []*schema.Column{BusinessesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BusinessFollowUsersColumns holds the columns for the "business_follow_users" table.
	BusinessFollowUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "business_followed_users", Type: field.TypeString, Nullable: true, Size: 36},
		{Name: "user_follower_businesses", Type: field.TypeString, Nullable: true, Size: 36},
	}
	// BusinessFollowUsersTable holds the schema information for the "business_follow_users" table.
	BusinessFollowUsersTable = &schema.Table{
		Name:       "business_follow_users",
		Columns:    BusinessFollowUsersColumns,
		PrimaryKey: []*schema.Column{BusinessFollowUsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "business_follow_users_businesses_followedUsers",
				Columns:    []*schema.Column{BusinessFollowUsersColumns[1]},
				RefColumns: []*schema.Column{BusinessesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "business_follow_users_users_followerBusinesses",
				Columns:    []*schema.Column{BusinessFollowUsersColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 36},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "image", Type: field.TypeString, Nullable: true},
		{Name: "business_categories", Type: field.TypeString, Nullable: true, Size: 36},
		{Name: "media_categories", Type: field.TypeString, Nullable: true, Size: 36},
		{Name: "menu_categories", Type: field.TypeString, Nullable: true, Size: 36},
		{Name: "place_categories", Type: field.TypeString, Nullable: true, Size: 36},
		{Name: "post_categories", Type: field.TypeString, Nullable: true, Size: 36},
		{Name: "user_categories", Type: field.TypeString, Nullable: true, Size: 36},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "categories_businesses_categories",
				Columns:    []*schema.Column{CategoriesColumns[3]},
				RefColumns: []*schema.Column{BusinessesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "categories_media_categories",
				Columns:    []*schema.Column{CategoriesColumns[4]},
				RefColumns: []*schema.Column{MediaColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "categories_menus_categories",
				Columns:    []*schema.Column{CategoriesColumns[5]},
				RefColumns: []*schema.Column{MenusColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "categories_places_categories",
				Columns:    []*schema.Column{CategoriesColumns[6]},
				RefColumns: []*schema.Column{PlacesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "categories_posts_categories",
				Columns:    []*schema.Column{CategoriesColumns[7]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "categories_users_categories",
				Columns:    []*schema.Column{CategoriesColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CategoryAssignmentsColumns holds the columns for the "category_assignments" table.
	CategoryAssignmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "entity_type", Type: field.TypeString, Nullable: true},
		{Name: "entity_id", Type: field.TypeString, Nullable: true, Size: 36},
		{Name: "category_id", Type: field.TypeString, Nullable: true, Size: 36},
	}
	// CategoryAssignmentsTable holds the schema information for the "category_assignments" table.
	CategoryAssignmentsTable = &schema.Table{
		Name:       "category_assignments",
		Columns:    CategoryAssignmentsColumns,
		PrimaryKey: []*schema.Column{CategoryAssignmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "category_assignments_businesses_categoryAssignments",
				Columns:    []*schema.Column{CategoryAssignmentsColumns[2]},
				RefColumns: []*schema.Column{BusinessesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "category_assignments_categories_categoryAssignments",
				Columns:    []*schema.Column{CategoryAssignmentsColumns[3]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "category_assignments_places_categoryAssignments",
				Columns:    []*schema.Column{CategoryAssignmentsColumns[2]},
				RefColumns: []*schema.Column{PlacesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "category_assignments_users_categoryAssignments",
				Columns:    []*schema.Column{CategoryAssignmentsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ChatsColumns holds the columns for the "chats" table.
	ChatsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
	}
	// ChatsTable holds the schema information for the "chats" table.
	ChatsTable = &schema.Table{
		Name:       "chats",
		Columns:    ChatsColumns,
		PrimaryKey: []*schema.Column{ChatsColumns[0]},
	}
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 36},
		{Name: "content", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "post_comments", Type: field.TypeString, Nullable: true, Size: 36},
		{Name: "user_comments", Type: field.TypeString, Nullable: true, Size: 36},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:       "comments",
		Columns:    CommentsColumns,
		PrimaryKey: []*schema.Column{CommentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comments_posts_comments",
				Columns:    []*schema.Column{CommentsColumns[4]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "comments_users_comments",
				Columns:    []*schema.Column{CommentsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// EventsColumns holds the columns for the "events" table.
	EventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "place_events", Type: field.TypeString, Nullable: true, Size: 36},
		{Name: "user_events", Type: field.TypeString, Nullable: true, Size: 36},
	}
	// EventsTable holds the schema information for the "events" table.
	EventsTable = &schema.Table{
		Name:       "events",
		Columns:    EventsColumns,
		PrimaryKey: []*schema.Column{EventsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "events_places_events",
				Columns:    []*schema.Column{EventsColumns[4]},
				RefColumns: []*schema.Column{PlacesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "events_users_events",
				Columns:    []*schema.Column{EventsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// HelpsColumns holds the columns for the "helps" table.
	HelpsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "category", Type: field.TypeString},
		{Name: "subject", Type: field.TypeString},
		{Name: "body", Type: field.TypeString, Size: 2147483647},
		{Name: "media", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeString, Default: "open"},
		{Name: "user_id", Type: field.TypeString, Size: 36},
	}
	// HelpsTable holds the schema information for the "helps" table.
	HelpsTable = &schema.Table{
		Name:       "helps",
		Columns:    HelpsColumns,
		PrimaryKey: []*schema.Column{HelpsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "helps_users_helps",
				Columns:    []*schema.Column{HelpsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// LikesColumns holds the columns for the "likes" table.
	LikesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 36},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "like_post", Type: field.TypeString, Nullable: true, Size: 36},
		{Name: "post_likes", Type: field.TypeString, Nullable: true, Size: 36},
		{Name: "user_likes", Type: field.TypeString, Nullable: true, Size: 36},
	}
	// LikesTable holds the schema information for the "likes" table.
	LikesTable = &schema.Table{
		Name:       "likes",
		Columns:    LikesColumns,
		PrimaryKey: []*schema.Column{LikesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "likes_posts_post",
				Columns:    []*schema.Column{LikesColumns[3]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "likes_posts_likes",
				Columns:    []*schema.Column{LikesColumns[4]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "likes_users_likes",
				Columns:    []*schema.Column{LikesColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MediaColumns holds the columns for the "media" table.
	MediaColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 36},
		{Name: "url", Type: field.TypeString},
		{Name: "media_type", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "post_medias", Type: field.TypeString, Nullable: true, Size: 36},
	}
	// MediaTable holds the schema information for the "media" table.
	MediaTable = &schema.Table{
		Name:       "media",
		Columns:    MediaColumns,
		PrimaryKey: []*schema.Column{MediaColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "media_posts_medias",
				Columns:    []*schema.Column{MediaColumns[5]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MenusColumns holds the columns for the "menus" table.
	MenusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 36},
		{Name: "place_menus", Type: field.TypeString, Nullable: true, Size: 36},
	}
	// MenusTable holds the schema information for the "menus" table.
	MenusTable = &schema.Table{
		Name:       "menus",
		Columns:    MenusColumns,
		PrimaryKey: []*schema.Column{MenusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "menus_places_menus",
				Columns:    []*schema.Column{MenusColumns[1]},
				RefColumns: []*schema.Column{PlacesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OrdersColumns holds the columns for the "orders" table.
	OrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
	}
	// OrdersTable holds the schema information for the "orders" table.
	OrdersTable = &schema.Table{
		Name:       "orders",
		Columns:    OrdersColumns,
		PrimaryKey: []*schema.Column{OrdersColumns[0]},
	}
	// PaymentsColumns holds the columns for the "payments" table.
	PaymentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
	}
	// PaymentsTable holds the schema information for the "payments" table.
	PaymentsTable = &schema.Table{
		Name:       "payments",
		Columns:    PaymentsColumns,
		PrimaryKey: []*schema.Column{PaymentsColumns[0]},
	}
	// PlacesColumns holds the columns for the "places" table.
	PlacesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 36},
		{Name: "name", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "location", Type: field.TypeString},
		{Name: "images", Type: field.TypeJSON, Nullable: true},
		{Name: "availability", Type: field.TypeJSON, Nullable: true},
		{Name: "special_offers", Type: field.TypeString, Nullable: true},
		{Name: "sustainability_score", Type: field.TypeFloat64, Nullable: true},
		{Name: "map_coordinates", Type: field.TypeJSON},
		{Name: "search_text", Type: field.TypeString, Nullable: true},
		{Name: "relevance_score", Type: field.TypeFloat64, Nullable: true},
		{Name: "business_places", Type: field.TypeString, Nullable: true, Size: 36},
		{Name: "user_places", Type: field.TypeString, Nullable: true, Size: 36},
	}
	// PlacesTable holds the schema information for the "places" table.
	PlacesTable = &schema.Table{
		Name:       "places",
		Columns:    PlacesColumns,
		PrimaryKey: []*schema.Column{PlacesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "places_businesses_places",
				Columns:    []*schema.Column{PlacesColumns[12]},
				RefColumns: []*schema.Column{BusinessesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "places_users_places",
				Columns:    []*schema.Column{PlacesColumns[13]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 36},
		{Name: "content", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "privacy", Type: field.TypeEnum, Enums: []string{"Public", "FollowersOnly", "OnlyMe"}, Default: "Public"},
		{Name: "business_posts", Type: field.TypeString, Nullable: true, Size: 36},
		{Name: "user_posts", Type: field.TypeString, Nullable: true, Size: 36},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "posts_businesses_posts",
				Columns:    []*schema.Column{PostsColumns[5]},
				RefColumns: []*schema.Column{BusinessesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "posts_users_posts",
				Columns:    []*schema.Column{PostsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RatingsColumns holds the columns for the "ratings" table.
	RatingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
	}
	// RatingsTable holds the schema information for the "ratings" table.
	RatingsTable = &schema.Table{
		Name:       "ratings",
		Columns:    RatingsColumns,
		PrimaryKey: []*schema.Column{RatingsColumns[0]},
	}
	// ReactionsColumns holds the columns for the "reactions" table.
	ReactionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
	}
	// ReactionsTable holds the schema information for the "reactions" table.
	ReactionsTable = &schema.Table{
		Name:       "reactions",
		Columns:    ReactionsColumns,
		PrimaryKey: []*schema.Column{ReactionsColumns[0]},
	}
	// ReservationsColumns holds the columns for the "reservations" table.
	ReservationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 36},
		{Name: "date", Type: field.TypeTime},
		{Name: "time", Type: field.TypeTime},
		{Name: "number_of_people", Type: field.TypeInt},
		{Name: "status", Type: field.TypeString},
		{Name: "place_reservations", Type: field.TypeString, Nullable: true, Size: 36},
		{Name: "user_reservations", Type: field.TypeString, Nullable: true, Size: 36},
	}
	// ReservationsTable holds the schema information for the "reservations" table.
	ReservationsTable = &schema.Table{
		Name:       "reservations",
		Columns:    ReservationsColumns,
		PrimaryKey: []*schema.Column{ReservationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "reservations_places_reservations",
				Columns:    []*schema.Column{ReservationsColumns[5]},
				RefColumns: []*schema.Column{PlacesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "reservations_users_reservations",
				Columns:    []*schema.Column{ReservationsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ReviewsColumns holds the columns for the "reviews" table.
	ReviewsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 36},
		{Name: "rating", Type: field.TypeFloat64},
		{Name: "comment", Type: field.TypeString, Nullable: true},
		{Name: "images_videos", Type: field.TypeJSON, Nullable: true},
		{Name: "timestamp", Type: field.TypeTime},
		{Name: "place_reviews", Type: field.TypeString, Nullable: true, Size: 36},
		{Name: "user_reviews", Type: field.TypeString, Nullable: true, Size: 36},
	}
	// ReviewsTable holds the schema information for the "reviews" table.
	ReviewsTable = &schema.Table{
		Name:       "reviews",
		Columns:    ReviewsColumns,
		PrimaryKey: []*schema.Column{ReviewsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "reviews_places_reviews",
				Columns:    []*schema.Column{ReviewsColumns[5]},
				RefColumns: []*schema.Column{PlacesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "reviews_users_reviews",
				Columns:    []*schema.Column{ReviewsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RoomsColumns holds the columns for the "rooms" table.
	RoomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 36},
		{Name: "number", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "availability", Type: field.TypeBool},
		{Name: "image", Type: field.TypeString, Nullable: true},
		{Name: "place_rooms", Type: field.TypeString, Nullable: true, Size: 36},
	}
	// RoomsTable holds the schema information for the "rooms" table.
	RoomsTable = &schema.Table{
		Name:       "rooms",
		Columns:    RoomsColumns,
		PrimaryKey: []*schema.Column{RoomsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "rooms_places_rooms",
				Columns:    []*schema.Column{RoomsColumns[7]},
				RefColumns: []*schema.Column{PlacesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TicketsColumns holds the columns for the "tickets" table.
	TicketsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "event_tickets", Type: field.TypeString, Nullable: true},
	}
	// TicketsTable holds the schema information for the "tickets" table.
	TicketsTable = &schema.Table{
		Name:       "tickets",
		Columns:    TicketsColumns,
		PrimaryKey: []*schema.Column{TicketsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tickets_events_tickets",
				Columns:    []*schema.Column{TicketsColumns[3]},
				RefColumns: []*schema.Column{EventsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TicketOptionsColumns holds the columns for the "ticket_options" table.
	TicketOptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "event_ticket_options", Type: field.TypeString, Nullable: true},
		{Name: "ticket_ticket_options", Type: field.TypeString, Nullable: true},
	}
	// TicketOptionsTable holds the schema information for the "ticket_options" table.
	TicketOptionsTable = &schema.Table{
		Name:       "ticket_options",
		Columns:    TicketOptionsColumns,
		PrimaryKey: []*schema.Column{TicketOptionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "ticket_options_events_ticket_options",
				Columns:    []*schema.Column{TicketOptionsColumns[3]},
				RefColumns: []*schema.Column{EventsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "ticket_options_tickets_ticket_options",
				Columns:    []*schema.Column{TicketOptionsColumns[4]},
				RefColumns: []*schema.Column{TicketsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 36},
		{Name: "auth0_id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "picture", Type: field.TypeString, Nullable: true},
		{Name: "cover_image", Type: field.TypeString, Nullable: true, Default: "https://res.cloudinary.com/placio/image/upload/v1686842319/mjl8stmbn5xmfsm50vbg.jpg"},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "website", Type: field.TypeString, Nullable: true},
		{Name: "location", Type: field.TypeString, Nullable: true},
		{Name: "bio", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: "Add a bio to your profile"},
		{Name: "auth0_data", Type: field.TypeJSON, Nullable: true},
		{Name: "app_settings", Type: field.TypeJSON, Nullable: true},
		{Name: "user_settings", Type: field.TypeJSON, Nullable: true},
		{Name: "search_text", Type: field.TypeString, Nullable: true},
		{Name: "relevance_score", Type: field.TypeFloat64, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserBusinessesColumns holds the columns for the "user_businesses" table.
	UserBusinessesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 36},
		{Name: "role", Type: field.TypeString},
		{Name: "business_user_businesses", Type: field.TypeString, Nullable: true, Size: 36},
		{Name: "user_user_businesses", Type: field.TypeString, Nullable: true, Size: 36},
	}
	// UserBusinessesTable holds the schema information for the "user_businesses" table.
	UserBusinessesTable = &schema.Table{
		Name:       "user_businesses",
		Columns:    UserBusinessesColumns,
		PrimaryKey: []*schema.Column{UserBusinessesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_businesses_businesses_userBusinesses",
				Columns:    []*schema.Column{UserBusinessesColumns[2]},
				RefColumns: []*schema.Column{BusinessesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "user_businesses_users_userBusinesses",
				Columns:    []*schema.Column{UserBusinessesColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserFollowBusinessesColumns holds the columns for the "user_follow_businesses" table.
	UserFollowBusinessesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "business_follower_users", Type: field.TypeString, Nullable: true, Size: 36},
		{Name: "user_followed_businesses", Type: field.TypeString, Nullable: true, Size: 36},
	}
	// UserFollowBusinessesTable holds the schema information for the "user_follow_businesses" table.
	UserFollowBusinessesTable = &schema.Table{
		Name:       "user_follow_businesses",
		Columns:    UserFollowBusinessesColumns,
		PrimaryKey: []*schema.Column{UserFollowBusinessesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_follow_businesses_businesses_followerUsers",
				Columns:    []*schema.Column{UserFollowBusinessesColumns[1]},
				RefColumns: []*schema.Column{BusinessesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "user_follow_businesses_users_followedBusinesses",
				Columns:    []*schema.Column{UserFollowBusinessesColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserFollowUsersColumns holds the columns for the "user_follow_users" table.
	UserFollowUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "user_followed_users", Type: field.TypeString, Nullable: true, Size: 36},
		{Name: "user_follower_users", Type: field.TypeString, Nullable: true, Size: 36},
	}
	// UserFollowUsersTable holds the schema information for the "user_follow_users" table.
	UserFollowUsersTable = &schema.Table{
		Name:       "user_follow_users",
		Columns:    UserFollowUsersColumns,
		PrimaryKey: []*schema.Column{UserFollowUsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_follow_users_users_followedUsers",
				Columns:    []*schema.Column{UserFollowUsersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "user_follow_users_users_followerUsers",
				Columns:    []*schema.Column{UserFollowUsersColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// AmenityPlacesColumns holds the columns for the "amenity_places" table.
	AmenityPlacesColumns = []*schema.Column{
		{Name: "amenity_id", Type: field.TypeString, Size: 36},
		{Name: "place_id", Type: field.TypeString, Size: 36},
	}
	// AmenityPlacesTable holds the schema information for the "amenity_places" table.
	AmenityPlacesTable = &schema.Table{
		Name:       "amenity_places",
		Columns:    AmenityPlacesColumns,
		PrimaryKey: []*schema.Column{AmenityPlacesColumns[0], AmenityPlacesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "amenity_places_amenity_id",
				Columns:    []*schema.Column{AmenityPlacesColumns[0]},
				RefColumns: []*schema.Column{AmenitiesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "amenity_places_place_id",
				Columns:    []*schema.Column{AmenityPlacesColumns[1]},
				RefColumns: []*schema.Column{PlacesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountSettingsTable,
		AmenitiesTable,
		BookingsTable,
		BusinessesTable,
		BusinessFollowBusinessesTable,
		BusinessFollowUsersTable,
		CategoriesTable,
		CategoryAssignmentsTable,
		ChatsTable,
		CommentsTable,
		EventsTable,
		HelpsTable,
		LikesTable,
		MediaTable,
		MenusTable,
		OrdersTable,
		PaymentsTable,
		PlacesTable,
		PostsTable,
		RatingsTable,
		ReactionsTable,
		ReservationsTable,
		ReviewsTable,
		RoomsTable,
		TicketsTable,
		TicketOptionsTable,
		UsersTable,
		UserBusinessesTable,
		UserFollowBusinessesTable,
		UserFollowUsersTable,
		AmenityPlacesTable,
	}
)

func init() {
	AccountSettingsTable.ForeignKeys[0].RefTable = BusinessesTable
	BookingsTable.ForeignKeys[0].RefTable = PlacesTable
	BookingsTable.ForeignKeys[1].RefTable = RoomsTable
	BookingsTable.ForeignKeys[2].RefTable = UsersTable
	BusinessFollowBusinessesTable.ForeignKeys[0].RefTable = BusinessesTable
	BusinessFollowBusinessesTable.ForeignKeys[1].RefTable = BusinessesTable
	BusinessFollowUsersTable.ForeignKeys[0].RefTable = BusinessesTable
	BusinessFollowUsersTable.ForeignKeys[1].RefTable = UsersTable
	CategoriesTable.ForeignKeys[0].RefTable = BusinessesTable
	CategoriesTable.ForeignKeys[1].RefTable = MediaTable
	CategoriesTable.ForeignKeys[2].RefTable = MenusTable
	CategoriesTable.ForeignKeys[3].RefTable = PlacesTable
	CategoriesTable.ForeignKeys[4].RefTable = PostsTable
	CategoriesTable.ForeignKeys[5].RefTable = UsersTable
	CategoryAssignmentsTable.ForeignKeys[0].RefTable = BusinessesTable
	CategoryAssignmentsTable.ForeignKeys[1].RefTable = CategoriesTable
	CategoryAssignmentsTable.ForeignKeys[2].RefTable = PlacesTable
	CategoryAssignmentsTable.ForeignKeys[3].RefTable = UsersTable
	CommentsTable.ForeignKeys[0].RefTable = PostsTable
	CommentsTable.ForeignKeys[1].RefTable = UsersTable
	EventsTable.ForeignKeys[0].RefTable = PlacesTable
	EventsTable.ForeignKeys[1].RefTable = UsersTable
	HelpsTable.ForeignKeys[0].RefTable = UsersTable
	LikesTable.ForeignKeys[0].RefTable = PostsTable
	LikesTable.ForeignKeys[1].RefTable = PostsTable
	LikesTable.ForeignKeys[2].RefTable = UsersTable
	MediaTable.ForeignKeys[0].RefTable = PostsTable
	MenusTable.ForeignKeys[0].RefTable = PlacesTable
	PlacesTable.ForeignKeys[0].RefTable = BusinessesTable
	PlacesTable.ForeignKeys[1].RefTable = UsersTable
	PostsTable.ForeignKeys[0].RefTable = BusinessesTable
	PostsTable.ForeignKeys[1].RefTable = UsersTable
	ReservationsTable.ForeignKeys[0].RefTable = PlacesTable
	ReservationsTable.ForeignKeys[1].RefTable = UsersTable
	ReviewsTable.ForeignKeys[0].RefTable = PlacesTable
	ReviewsTable.ForeignKeys[1].RefTable = UsersTable
	RoomsTable.ForeignKeys[0].RefTable = PlacesTable
	TicketsTable.ForeignKeys[0].RefTable = EventsTable
	TicketOptionsTable.ForeignKeys[0].RefTable = EventsTable
	TicketOptionsTable.ForeignKeys[1].RefTable = TicketsTable
	UserBusinessesTable.ForeignKeys[0].RefTable = BusinessesTable
	UserBusinessesTable.ForeignKeys[1].RefTable = UsersTable
	UserFollowBusinessesTable.ForeignKeys[0].RefTable = BusinessesTable
	UserFollowBusinessesTable.ForeignKeys[1].RefTable = UsersTable
	UserFollowUsersTable.ForeignKeys[0].RefTable = UsersTable
	UserFollowUsersTable.ForeignKeys[1].RefTable = UsersTable
	AmenityPlacesTable.ForeignKeys[0].RefTable = AmenitiesTable
	AmenityPlacesTable.ForeignKeys[1].RefTable = PlacesTable
}
