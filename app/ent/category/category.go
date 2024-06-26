// Code generated by ent, DO NOT EDIT.

package category

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the category type in the database.
	Label = "category"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldIcon holds the string denoting the icon field in the database.
	FieldIcon = "icon"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldParentID holds the string denoting the parent_id field in the database.
	FieldParentID = "parent_id"
	// FieldParentName holds the string denoting the parent_name field in the database.
	FieldParentName = "parent_name"
	// FieldParentImage holds the string denoting the parent_image field in the database.
	FieldParentImage = "parent_image"
	// FieldParentDescription holds the string denoting the parent_description field in the database.
	FieldParentDescription = "parent_description"
	// EdgeCategoryAssignments holds the string denoting the categoryassignments edge name in mutations.
	EdgeCategoryAssignments = "categoryAssignments"
	// EdgePlaceInventories holds the string denoting the place_inventories edge name in mutations.
	EdgePlaceInventories = "place_inventories"
	// EdgeMedia holds the string denoting the media edge name in mutations.
	EdgeMedia = "media"
	// EdgeMenus holds the string denoting the menus edge name in mutations.
	EdgeMenus = "menus"
	// Table holds the table name of the category in the database.
	Table = "categories"
	// CategoryAssignmentsTable is the table that holds the categoryAssignments relation/edge.
	CategoryAssignmentsTable = "category_assignments"
	// CategoryAssignmentsInverseTable is the table name for the CategoryAssignment entity.
	// It exists in this package in order to avoid circular dependency with the "categoryassignment" package.
	CategoryAssignmentsInverseTable = "category_assignments"
	// CategoryAssignmentsColumn is the table column denoting the categoryAssignments relation/edge.
	CategoryAssignmentsColumn = "category_id"
	// PlaceInventoriesTable is the table that holds the place_inventories relation/edge.
	PlaceInventoriesTable = "place_inventories"
	// PlaceInventoriesInverseTable is the table name for the PlaceInventory entity.
	// It exists in this package in order to avoid circular dependency with the "placeinventory" package.
	PlaceInventoriesInverseTable = "place_inventories"
	// PlaceInventoriesColumn is the table column denoting the place_inventories relation/edge.
	PlaceInventoriesColumn = "category_place_inventories"
	// MediaTable is the table that holds the media relation/edge. The primary key declared below.
	MediaTable = "category_media"
	// MediaInverseTable is the table name for the Media entity.
	// It exists in this package in order to avoid circular dependency with the "media" package.
	MediaInverseTable = "media"
	// MenusTable is the table that holds the menus relation/edge. The primary key declared below.
	MenusTable = "menu_categories"
	// MenusInverseTable is the table name for the Menu entity.
	// It exists in this package in order to avoid circular dependency with the "menu" package.
	MenusInverseTable = "menus"
)

// Columns holds all SQL columns for category fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldImage,
	FieldDescription,
	FieldIcon,
	FieldType,
	FieldParentID,
	FieldParentName,
	FieldParentImage,
	FieldParentDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "categories"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"business_categories",
	"event_event_categories",
	"place_categories",
	"post_categories",
	"user_categories",
}

var (
	// MediaPrimaryKey and MediaColumn2 are the table columns denoting the
	// primary key for the media relation (M2M).
	MediaPrimaryKey = []string{"category_id", "media_id"}
	// MenusPrimaryKey and MenusColumn2 are the table columns denoting the
	// primary key for the menus relation (M2M).
	MenusPrimaryKey = []string{"menu_id", "category_id"}
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
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the Category queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByImage orders the results by the image field.
func ByImage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImage, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByIcon orders the results by the icon field.
func ByIcon(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIcon, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByParentID orders the results by the parent_id field.
func ByParentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParentID, opts...).ToFunc()
}

// ByParentName orders the results by the parent_name field.
func ByParentName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParentName, opts...).ToFunc()
}

// ByParentImage orders the results by the parent_image field.
func ByParentImage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParentImage, opts...).ToFunc()
}

// ByParentDescription orders the results by the parent_description field.
func ByParentDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParentDescription, opts...).ToFunc()
}

// ByCategoryAssignmentsCount orders the results by categoryAssignments count.
func ByCategoryAssignmentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCategoryAssignmentsStep(), opts...)
	}
}

// ByCategoryAssignments orders the results by categoryAssignments terms.
func ByCategoryAssignments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCategoryAssignmentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPlaceInventoriesCount orders the results by place_inventories count.
func ByPlaceInventoriesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPlaceInventoriesStep(), opts...)
	}
}

// ByPlaceInventories orders the results by place_inventories terms.
func ByPlaceInventories(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPlaceInventoriesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMediaCount orders the results by media count.
func ByMediaCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMediaStep(), opts...)
	}
}

// ByMedia orders the results by media terms.
func ByMedia(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMediaStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMenusCount orders the results by menus count.
func ByMenusCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMenusStep(), opts...)
	}
}

// ByMenus orders the results by menus terms.
func ByMenus(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMenusStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCategoryAssignmentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CategoryAssignmentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CategoryAssignmentsTable, CategoryAssignmentsColumn),
	)
}
func newPlaceInventoriesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PlaceInventoriesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PlaceInventoriesTable, PlaceInventoriesColumn),
	)
}
func newMediaStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MediaInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, MediaTable, MediaPrimaryKey...),
	)
}
func newMenusStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MenusInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, MenusTable, MenusPrimaryKey...),
	)
}
