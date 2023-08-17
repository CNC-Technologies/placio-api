// Code generated by ent, DO NOT EDIT.

package media

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the media type in the database.
	Label = "media"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldMediaType holds the string denoting the mediatype field in the database.
	FieldMediaType = "media_type"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldLikeCount holds the string denoting the likecount field in the database.
	FieldLikeCount = "like_count"
	// FieldDislikeCount holds the string denoting the dislikecount field in the database.
	FieldDislikeCount = "dislike_count"
	// EdgePost holds the string denoting the post edge name in mutations.
	EdgePost = "post"
	// EdgeReview holds the string denoting the review edge name in mutations.
	EdgeReview = "review"
	// EdgeCategories holds the string denoting the categories edge name in mutations.
	EdgeCategories = "categories"
	// EdgePlace holds the string denoting the place edge name in mutations.
	EdgePlace = "place"
	// EdgePlaceInventory holds the string denoting the place_inventory edge name in mutations.
	EdgePlaceInventory = "place_inventory"
	// Table holds the table name of the media in the database.
	Table = "media"
	// PostTable is the table that holds the post relation/edge.
	PostTable = "media"
	// PostInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostInverseTable = "posts"
	// PostColumn is the table column denoting the post relation/edge.
	PostColumn = "post_medias"
	// ReviewTable is the table that holds the review relation/edge.
	ReviewTable = "media"
	// ReviewInverseTable is the table name for the Review entity.
	// It exists in this package in order to avoid circular dependency with the "review" package.
	ReviewInverseTable = "reviews"
	// ReviewColumn is the table column denoting the review relation/edge.
	ReviewColumn = "review_medias"
	// CategoriesTable is the table that holds the categories relation/edge.
	CategoriesTable = "categories"
	// CategoriesInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	CategoriesInverseTable = "categories"
	// CategoriesColumn is the table column denoting the categories relation/edge.
	CategoriesColumn = "media_categories"
	// PlaceTable is the table that holds the place relation/edge. The primary key declared below.
	PlaceTable = "place_medias"
	// PlaceInverseTable is the table name for the Place entity.
	// It exists in this package in order to avoid circular dependency with the "place" package.
	PlaceInverseTable = "places"
	// PlaceInventoryTable is the table that holds the place_inventory relation/edge. The primary key declared below.
	PlaceInventoryTable = "place_inventory_media"
	// PlaceInventoryInverseTable is the table name for the PlaceInventory entity.
	// It exists in this package in order to avoid circular dependency with the "placeinventory" package.
	PlaceInventoryInverseTable = "place_inventories"
)

// Columns holds all SQL columns for media fields.
var Columns = []string{
	FieldID,
	FieldURL,
	FieldMediaType,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldLikeCount,
	FieldDislikeCount,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "media"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"post_medias",
	"review_medias",
}

var (
	// PlacePrimaryKey and PlaceColumn2 are the table columns denoting the
	// primary key for the place relation (M2M).
	PlacePrimaryKey = []string{"place_id", "media_id"}
	// PlaceInventoryPrimaryKey and PlaceInventoryColumn2 are the table columns denoting the
	// primary key for the place_inventory relation (M2M).
	PlaceInventoryPrimaryKey = []string{"place_inventory_id", "media_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "CreatedAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "UpdatedAt" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "UpdatedAt" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultLikeCount holds the default value on creation for the "likeCount" field.
	DefaultLikeCount int
	// DefaultDislikeCount holds the default value on creation for the "dislikeCount" field.
	DefaultDislikeCount int
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the Media queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByURL orders the results by the URL field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}

// ByMediaType orders the results by the MediaType field.
func ByMediaType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMediaType, opts...).ToFunc()
}

// ByCreatedAt orders the results by the CreatedAt field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the UpdatedAt field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByLikeCount orders the results by the likeCount field.
func ByLikeCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLikeCount, opts...).ToFunc()
}

// ByDislikeCount orders the results by the dislikeCount field.
func ByDislikeCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDislikeCount, opts...).ToFunc()
}

// ByPostField orders the results by post field.
func ByPostField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPostStep(), sql.OrderByField(field, opts...))
	}
}

// ByReviewField orders the results by review field.
func ByReviewField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReviewStep(), sql.OrderByField(field, opts...))
	}
}

// ByCategoriesCount orders the results by categories count.
func ByCategoriesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCategoriesStep(), opts...)
	}
}

// ByCategories orders the results by categories terms.
func ByCategories(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCategoriesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPlaceCount orders the results by place count.
func ByPlaceCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPlaceStep(), opts...)
	}
}

// ByPlace orders the results by place terms.
func ByPlace(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPlaceStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPlaceInventoryCount orders the results by place_inventory count.
func ByPlaceInventoryCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPlaceInventoryStep(), opts...)
	}
}

// ByPlaceInventory orders the results by place_inventory terms.
func ByPlaceInventory(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPlaceInventoryStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPostStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PostInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PostTable, PostColumn),
	)
}
func newReviewStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReviewInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ReviewTable, ReviewColumn),
	)
}
func newCategoriesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CategoriesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CategoriesTable, CategoriesColumn),
	)
}
func newPlaceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PlaceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, PlaceTable, PlacePrimaryKey...),
	)
}
func newPlaceInventoryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PlaceInventoryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, PlaceInventoryTable, PlaceInventoryPrimaryKey...),
	)
}
