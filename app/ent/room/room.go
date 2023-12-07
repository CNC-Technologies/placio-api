// Code generated by ent, DO NOT EDIT.

package room

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the room type in the database.
	Label = "room"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRoomNumber holds the string denoting the room_number field in the database.
	FieldRoomNumber = "room_number"
	// FieldRoomType holds the string denoting the room_type field in the database.
	FieldRoomType = "room_type"
	// FieldRoomStatus holds the string denoting the room_status field in the database.
	FieldRoomStatus = "room_status"
	// FieldRoomRating holds the string denoting the room_rating field in the database.
	FieldRoomRating = "room_rating"
	// FieldRoomPrice holds the string denoting the room_price field in the database.
	FieldRoomPrice = "room_price"
	// FieldQrCode holds the string denoting the qr_code field in the database.
	FieldQrCode = "qr_code"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldExtras holds the string denoting the extras field in the database.
	FieldExtras = "extras"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldAvailability holds the string denoting the availability field in the database.
	FieldAvailability = "availability"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// EdgePlace holds the string denoting the place edge name in mutations.
	EdgePlace = "place"
	// EdgeRoomCategory holds the string denoting the room_category edge name in mutations.
	EdgeRoomCategory = "room_category"
	// EdgeBookings holds the string denoting the bookings edge name in mutations.
	EdgeBookings = "bookings"
	// EdgeAmenities holds the string denoting the amenities edge name in mutations.
	EdgeAmenities = "amenities"
	// EdgeMedia holds the string denoting the media edge name in mutations.
	EdgeMedia = "media"
	// EdgeReservations holds the string denoting the reservations edge name in mutations.
	EdgeReservations = "reservations"
	// Table holds the table name of the room in the database.
	Table = "rooms"
	// PlaceTable is the table that holds the place relation/edge. The primary key declared below.
	PlaceTable = "place_rooms"
	// PlaceInverseTable is the table name for the Place entity.
	// It exists in this package in order to avoid circular dependency with the "place" package.
	PlaceInverseTable = "places"
	// RoomCategoryTable is the table that holds the room_category relation/edge. The primary key declared below.
	RoomCategoryTable = "room_category_rooms"
	// RoomCategoryInverseTable is the table name for the RoomCategory entity.
	// It exists in this package in order to avoid circular dependency with the "roomcategory" package.
	RoomCategoryInverseTable = "room_categories"
	// BookingsTable is the table that holds the bookings relation/edge.
	BookingsTable = "bookings"
	// BookingsInverseTable is the table name for the Booking entity.
	// It exists in this package in order to avoid circular dependency with the "booking" package.
	BookingsInverseTable = "bookings"
	// BookingsColumn is the table column denoting the bookings relation/edge.
	BookingsColumn = "room_bookings"
	// AmenitiesTable is the table that holds the amenities relation/edge. The primary key declared below.
	AmenitiesTable = "room_amenities"
	// AmenitiesInverseTable is the table name for the Amenity entity.
	// It exists in this package in order to avoid circular dependency with the "amenity" package.
	AmenitiesInverseTable = "amenities"
	// MediaTable is the table that holds the media relation/edge. The primary key declared below.
	MediaTable = "room_media"
	// MediaInverseTable is the table name for the Media entity.
	// It exists in this package in order to avoid circular dependency with the "media" package.
	MediaInverseTable = "media"
	// ReservationsTable is the table that holds the reservations relation/edge.
	ReservationsTable = "reservations"
	// ReservationsInverseTable is the table name for the Reservation entity.
	// It exists in this package in order to avoid circular dependency with the "reservation" package.
	ReservationsInverseTable = "reservations"
	// ReservationsColumn is the table column denoting the reservations relation/edge.
	ReservationsColumn = "room_reservations"
)

// Columns holds all SQL columns for room fields.
var Columns = []string{
	FieldID,
	FieldRoomNumber,
	FieldRoomType,
	FieldRoomStatus,
	FieldRoomRating,
	FieldRoomPrice,
	FieldQrCode,
	FieldStatus,
	FieldExtras,
	FieldDescription,
	FieldAvailability,
	FieldImage,
}

var (
	// PlacePrimaryKey and PlaceColumn2 are the table columns denoting the
	// primary key for the place relation (M2M).
	PlacePrimaryKey = []string{"place_id", "room_id"}
	// RoomCategoryPrimaryKey and RoomCategoryColumn2 are the table columns denoting the
	// primary key for the room_category relation (M2M).
	RoomCategoryPrimaryKey = []string{"room_category_id", "room_id"}
	// AmenitiesPrimaryKey and AmenitiesColumn2 are the table columns denoting the
	// primary key for the amenities relation (M2M).
	AmenitiesPrimaryKey = []string{"room_id", "amenity_id"}
	// MediaPrimaryKey and MediaColumn2 are the table columns denoting the
	// primary key for the media relation (M2M).
	MediaPrimaryKey = []string{"room_id", "media_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultAvailability holds the default value on creation for the "availability" field.
	DefaultAvailability bool
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusAvailable   Status = "available"
	StatusUnavailable Status = "unavailable"
	StatusMaintenance Status = "maintenance"
	StatusReserved    Status = "reserved"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusAvailable, StatusUnavailable, StatusMaintenance, StatusReserved:
		return nil
	default:
		return fmt.Errorf("room: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Room queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByRoomNumber orders the results by the room_number field.
func ByRoomNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRoomNumber, opts...).ToFunc()
}

// ByRoomType orders the results by the room_type field.
func ByRoomType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRoomType, opts...).ToFunc()
}

// ByRoomStatus orders the results by the room_status field.
func ByRoomStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRoomStatus, opts...).ToFunc()
}

// ByRoomRating orders the results by the room_rating field.
func ByRoomRating(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRoomRating, opts...).ToFunc()
}

// ByRoomPrice orders the results by the room_price field.
func ByRoomPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRoomPrice, opts...).ToFunc()
}

// ByQrCode orders the results by the qr_code field.
func ByQrCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQrCode, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByAvailability orders the results by the availability field.
func ByAvailability(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvailability, opts...).ToFunc()
}

// ByImage orders the results by the image field.
func ByImage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImage, opts...).ToFunc()
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

// ByRoomCategoryCount orders the results by room_category count.
func ByRoomCategoryCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRoomCategoryStep(), opts...)
	}
}

// ByRoomCategory orders the results by room_category terms.
func ByRoomCategory(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRoomCategoryStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByBookingsCount orders the results by bookings count.
func ByBookingsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBookingsStep(), opts...)
	}
}

// ByBookings orders the results by bookings terms.
func ByBookings(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBookingsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAmenitiesCount orders the results by amenities count.
func ByAmenitiesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAmenitiesStep(), opts...)
	}
}

// ByAmenities orders the results by amenities terms.
func ByAmenities(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAmenitiesStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByReservationsCount orders the results by reservations count.
func ByReservationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newReservationsStep(), opts...)
	}
}

// ByReservations orders the results by reservations terms.
func ByReservations(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReservationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPlaceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PlaceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, PlaceTable, PlacePrimaryKey...),
	)
}
func newRoomCategoryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoomCategoryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, RoomCategoryTable, RoomCategoryPrimaryKey...),
	)
}
func newBookingsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BookingsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, BookingsTable, BookingsColumn),
	)
}
func newAmenitiesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AmenitiesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, AmenitiesTable, AmenitiesPrimaryKey...),
	)
}
func newMediaStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MediaInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, MediaTable, MediaPrimaryKey...),
	)
}
func newReservationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReservationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ReservationsTable, ReservationsColumn),
	)
}
