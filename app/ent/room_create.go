// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/amenity"
	"placio-app/ent/booking"
	"placio-app/ent/media"
	"placio-app/ent/place"
	"placio-app/ent/reservation"
	"placio-app/ent/room"
	"placio-app/ent/roomcategory"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoomCreate is the builder for creating a Room entity.
type RoomCreate struct {
	config
	mutation *RoomMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (rc *RoomCreate) SetName(s string) *RoomCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (rc *RoomCreate) SetNillableName(s *string) *RoomCreate {
	if s != nil {
		rc.SetName(*s)
	}
	return rc
}

// SetRoomNumber sets the "room_number" field.
func (rc *RoomCreate) SetRoomNumber(i int) *RoomCreate {
	rc.mutation.SetRoomNumber(i)
	return rc
}

// SetNillableRoomNumber sets the "room_number" field if the given value is not nil.
func (rc *RoomCreate) SetNillableRoomNumber(i *int) *RoomCreate {
	if i != nil {
		rc.SetRoomNumber(*i)
	}
	return rc
}

// SetRoomType sets the "room_type" field.
func (rc *RoomCreate) SetRoomType(s string) *RoomCreate {
	rc.mutation.SetRoomType(s)
	return rc
}

// SetNillableRoomType sets the "room_type" field if the given value is not nil.
func (rc *RoomCreate) SetNillableRoomType(s *string) *RoomCreate {
	if s != nil {
		rc.SetRoomType(*s)
	}
	return rc
}

// SetRoomStatus sets the "room_status" field.
func (rc *RoomCreate) SetRoomStatus(s string) *RoomCreate {
	rc.mutation.SetRoomStatus(s)
	return rc
}

// SetNillableRoomStatus sets the "room_status" field if the given value is not nil.
func (rc *RoomCreate) SetNillableRoomStatus(s *string) *RoomCreate {
	if s != nil {
		rc.SetRoomStatus(*s)
	}
	return rc
}

// SetRoomRating sets the "room_rating" field.
func (rc *RoomCreate) SetRoomRating(s string) *RoomCreate {
	rc.mutation.SetRoomRating(s)
	return rc
}

// SetNillableRoomRating sets the "room_rating" field if the given value is not nil.
func (rc *RoomCreate) SetNillableRoomRating(s *string) *RoomCreate {
	if s != nil {
		rc.SetRoomRating(*s)
	}
	return rc
}

// SetRoomPrice sets the "room_price" field.
func (rc *RoomCreate) SetRoomPrice(f float64) *RoomCreate {
	rc.mutation.SetRoomPrice(f)
	return rc
}

// SetNillableRoomPrice sets the "room_price" field if the given value is not nil.
func (rc *RoomCreate) SetNillableRoomPrice(f *float64) *RoomCreate {
	if f != nil {
		rc.SetRoomPrice(*f)
	}
	return rc
}

// SetGuestCapacity sets the "guest_capacity" field.
func (rc *RoomCreate) SetGuestCapacity(s string) *RoomCreate {
	rc.mutation.SetGuestCapacity(s)
	return rc
}

// SetNillableGuestCapacity sets the "guest_capacity" field if the given value is not nil.
func (rc *RoomCreate) SetNillableGuestCapacity(s *string) *RoomCreate {
	if s != nil {
		rc.SetGuestCapacity(*s)
	}
	return rc
}

// SetQrCode sets the "qr_code" field.
func (rc *RoomCreate) SetQrCode(s string) *RoomCreate {
	rc.mutation.SetQrCode(s)
	return rc
}

// SetNillableQrCode sets the "qr_code" field if the given value is not nil.
func (rc *RoomCreate) SetNillableQrCode(s *string) *RoomCreate {
	if s != nil {
		rc.SetQrCode(*s)
	}
	return rc
}

// SetStatus sets the "status" field.
func (rc *RoomCreate) SetStatus(r room.Status) *RoomCreate {
	rc.mutation.SetStatus(r)
	return rc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rc *RoomCreate) SetNillableStatus(r *room.Status) *RoomCreate {
	if r != nil {
		rc.SetStatus(*r)
	}
	return rc
}

// SetExtras sets the "extras" field.
func (rc *RoomCreate) SetExtras(m map[string]interface{}) *RoomCreate {
	rc.mutation.SetExtras(m)
	return rc
}

// SetDescription sets the "description" field.
func (rc *RoomCreate) SetDescription(s string) *RoomCreate {
	rc.mutation.SetDescription(s)
	return rc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (rc *RoomCreate) SetNillableDescription(s *string) *RoomCreate {
	if s != nil {
		rc.SetDescription(*s)
	}
	return rc
}

// SetAvailability sets the "availability" field.
func (rc *RoomCreate) SetAvailability(b bool) *RoomCreate {
	rc.mutation.SetAvailability(b)
	return rc
}

// SetNillableAvailability sets the "availability" field if the given value is not nil.
func (rc *RoomCreate) SetNillableAvailability(b *bool) *RoomCreate {
	if b != nil {
		rc.SetAvailability(*b)
	}
	return rc
}

// SetImage sets the "image" field.
func (rc *RoomCreate) SetImage(s string) *RoomCreate {
	rc.mutation.SetImage(s)
	return rc
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (rc *RoomCreate) SetNillableImage(s *string) *RoomCreate {
	if s != nil {
		rc.SetImage(*s)
	}
	return rc
}

// SetID sets the "id" field.
func (rc *RoomCreate) SetID(s string) *RoomCreate {
	rc.mutation.SetID(s)
	return rc
}

// AddPlaceIDs adds the "place" edge to the Place entity by IDs.
func (rc *RoomCreate) AddPlaceIDs(ids ...string) *RoomCreate {
	rc.mutation.AddPlaceIDs(ids...)
	return rc
}

// AddPlace adds the "place" edges to the Place entity.
func (rc *RoomCreate) AddPlace(p ...*Place) *RoomCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return rc.AddPlaceIDs(ids...)
}

// AddRoomCategoryIDs adds the "room_category" edge to the RoomCategory entity by IDs.
func (rc *RoomCreate) AddRoomCategoryIDs(ids ...string) *RoomCreate {
	rc.mutation.AddRoomCategoryIDs(ids...)
	return rc
}

// AddRoomCategory adds the "room_category" edges to the RoomCategory entity.
func (rc *RoomCreate) AddRoomCategory(r ...*RoomCategory) *RoomCreate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rc.AddRoomCategoryIDs(ids...)
}

// AddBookingIDs adds the "bookings" edge to the Booking entity by IDs.
func (rc *RoomCreate) AddBookingIDs(ids ...string) *RoomCreate {
	rc.mutation.AddBookingIDs(ids...)
	return rc
}

// AddBookings adds the "bookings" edges to the Booking entity.
func (rc *RoomCreate) AddBookings(b ...*Booking) *RoomCreate {
	ids := make([]string, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return rc.AddBookingIDs(ids...)
}

// AddAmenityIDs adds the "amenities" edge to the Amenity entity by IDs.
func (rc *RoomCreate) AddAmenityIDs(ids ...string) *RoomCreate {
	rc.mutation.AddAmenityIDs(ids...)
	return rc
}

// AddAmenities adds the "amenities" edges to the Amenity entity.
func (rc *RoomCreate) AddAmenities(a ...*Amenity) *RoomCreate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return rc.AddAmenityIDs(ids...)
}

// AddMediumIDs adds the "media" edge to the Media entity by IDs.
func (rc *RoomCreate) AddMediumIDs(ids ...string) *RoomCreate {
	rc.mutation.AddMediumIDs(ids...)
	return rc
}

// AddMedia adds the "media" edges to the Media entity.
func (rc *RoomCreate) AddMedia(m ...*Media) *RoomCreate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return rc.AddMediumIDs(ids...)
}

// AddReservationIDs adds the "reservations" edge to the Reservation entity by IDs.
func (rc *RoomCreate) AddReservationIDs(ids ...string) *RoomCreate {
	rc.mutation.AddReservationIDs(ids...)
	return rc
}

// AddReservations adds the "reservations" edges to the Reservation entity.
func (rc *RoomCreate) AddReservations(r ...*Reservation) *RoomCreate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rc.AddReservationIDs(ids...)
}

// Mutation returns the RoomMutation object of the builder.
func (rc *RoomCreate) Mutation() *RoomMutation {
	return rc.mutation
}

// Save creates the Room in the database.
func (rc *RoomCreate) Save(ctx context.Context) (*Room, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RoomCreate) SaveX(ctx context.Context) *Room {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RoomCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RoomCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RoomCreate) defaults() {
	if _, ok := rc.mutation.Status(); !ok {
		v := room.DefaultStatus
		rc.mutation.SetStatus(v)
	}
	if _, ok := rc.mutation.Availability(); !ok {
		v := room.DefaultAvailability
		rc.mutation.SetAvailability(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RoomCreate) check() error {
	if _, ok := rc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Room.status"`)}
	}
	if v, ok := rc.mutation.Status(); ok {
		if err := room.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Room.status": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Availability(); !ok {
		return &ValidationError{Name: "availability", err: errors.New(`ent: missing required field "Room.availability"`)}
	}
	if v, ok := rc.mutation.ID(); ok {
		if err := room.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Room.id": %w`, err)}
		}
	}
	return nil
}

func (rc *RoomCreate) sqlSave(ctx context.Context) (*Room, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Room.ID type: %T", _spec.ID.Value)
		}
	}
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RoomCreate) createSpec() (*Room, *sqlgraph.CreateSpec) {
	var (
		_node = &Room{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(room.Table, sqlgraph.NewFieldSpec(room.FieldID, field.TypeString))
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.Name(); ok {
		_spec.SetField(room.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := rc.mutation.RoomNumber(); ok {
		_spec.SetField(room.FieldRoomNumber, field.TypeInt, value)
		_node.RoomNumber = value
	}
	if value, ok := rc.mutation.RoomType(); ok {
		_spec.SetField(room.FieldRoomType, field.TypeString, value)
		_node.RoomType = value
	}
	if value, ok := rc.mutation.RoomStatus(); ok {
		_spec.SetField(room.FieldRoomStatus, field.TypeString, value)
		_node.RoomStatus = value
	}
	if value, ok := rc.mutation.RoomRating(); ok {
		_spec.SetField(room.FieldRoomRating, field.TypeString, value)
		_node.RoomRating = value
	}
	if value, ok := rc.mutation.RoomPrice(); ok {
		_spec.SetField(room.FieldRoomPrice, field.TypeFloat64, value)
		_node.RoomPrice = value
	}
	if value, ok := rc.mutation.GuestCapacity(); ok {
		_spec.SetField(room.FieldGuestCapacity, field.TypeString, value)
		_node.GuestCapacity = value
	}
	if value, ok := rc.mutation.QrCode(); ok {
		_spec.SetField(room.FieldQrCode, field.TypeString, value)
		_node.QrCode = value
	}
	if value, ok := rc.mutation.Status(); ok {
		_spec.SetField(room.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := rc.mutation.Extras(); ok {
		_spec.SetField(room.FieldExtras, field.TypeJSON, value)
		_node.Extras = value
	}
	if value, ok := rc.mutation.Description(); ok {
		_spec.SetField(room.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := rc.mutation.Availability(); ok {
		_spec.SetField(room.FieldAvailability, field.TypeBool, value)
		_node.Availability = value
	}
	if value, ok := rc.mutation.Image(); ok {
		_spec.SetField(room.FieldImage, field.TypeString, value)
		_node.Image = value
	}
	if nodes := rc.mutation.PlaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   room.PlaceTable,
			Columns: room.PlacePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(place.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.RoomCategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   room.RoomCategoryTable,
			Columns: room.RoomCategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(roomcategory.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.BookingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.BookingsTable,
			Columns: []string{room.BookingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(booking.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.AmenitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   room.AmenitiesTable,
			Columns: room.AmenitiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(amenity.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.MediaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   room.MediaTable,
			Columns: room.MediaPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(media.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.ReservationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.ReservationsTable,
			Columns: []string{room.ReservationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reservation.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RoomCreateBulk is the builder for creating many Room entities in bulk.
type RoomCreateBulk struct {
	config
	err      error
	builders []*RoomCreate
}

// Save creates the Room entities in the database.
func (rcb *RoomCreateBulk) Save(ctx context.Context) ([]*Room, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Room, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoomMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RoomCreateBulk) SaveX(ctx context.Context) []*Room {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RoomCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RoomCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
