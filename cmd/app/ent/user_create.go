// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/booking"
	"placio-app/ent/businessfollowuser"
	"placio-app/ent/category"
	"placio-app/ent/categoryassignment"
	"placio-app/ent/comment"
	"placio-app/ent/event"
	"placio-app/ent/help"
	"placio-app/ent/like"
	"placio-app/ent/place"
	"placio-app/ent/post"
	"placio-app/ent/reservation"
	"placio-app/ent/review"
	"placio-app/ent/user"
	"placio-app/ent/userbusiness"
	"placio-app/ent/userfollowbusiness"
	"placio-app/ent/userfollowuser"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/auth0/go-auth0/management"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetAuth0ID sets the "auth0_id" field.
func (uc *UserCreate) SetAuth0ID(s string) *UserCreate {
	uc.mutation.SetAuth0ID(s)
	return uc
}

// SetName sets the "name" field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uc *UserCreate) SetNillableName(s *string) *UserCreate {
	if s != nil {
		uc.SetName(*s)
	}
	return uc
}

// SetPicture sets the "picture" field.
func (uc *UserCreate) SetPicture(s string) *UserCreate {
	uc.mutation.SetPicture(s)
	return uc
}

// SetNillablePicture sets the "picture" field if the given value is not nil.
func (uc *UserCreate) SetNillablePicture(s *string) *UserCreate {
	if s != nil {
		uc.SetPicture(*s)
	}
	return uc
}

// SetCoverImage sets the "cover_image" field.
func (uc *UserCreate) SetCoverImage(s string) *UserCreate {
	uc.mutation.SetCoverImage(s)
	return uc
}

// SetNillableCoverImage sets the "cover_image" field if the given value is not nil.
func (uc *UserCreate) SetNillableCoverImage(s *string) *UserCreate {
	if s != nil {
		uc.SetCoverImage(*s)
	}
	return uc
}

// SetUsername sets the "username" field.
func (uc *UserCreate) SetUsername(s string) *UserCreate {
	uc.mutation.SetUsername(s)
	return uc
}

// SetWebsite sets the "website" field.
func (uc *UserCreate) SetWebsite(s string) *UserCreate {
	uc.mutation.SetWebsite(s)
	return uc
}

// SetNillableWebsite sets the "website" field if the given value is not nil.
func (uc *UserCreate) SetNillableWebsite(s *string) *UserCreate {
	if s != nil {
		uc.SetWebsite(*s)
	}
	return uc
}

// SetLocation sets the "location" field.
func (uc *UserCreate) SetLocation(s string) *UserCreate {
	uc.mutation.SetLocation(s)
	return uc
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (uc *UserCreate) SetNillableLocation(s *string) *UserCreate {
	if s != nil {
		uc.SetLocation(*s)
	}
	return uc
}

// SetBio sets the "bio" field.
func (uc *UserCreate) SetBio(s string) *UserCreate {
	uc.mutation.SetBio(s)
	return uc
}

// SetNillableBio sets the "bio" field if the given value is not nil.
func (uc *UserCreate) SetNillableBio(s *string) *UserCreate {
	if s != nil {
		uc.SetBio(*s)
	}
	return uc
}

// SetAuth0Data sets the "auth0_data" field.
func (uc *UserCreate) SetAuth0Data(m *management.User) *UserCreate {
	uc.mutation.SetAuth0Data(m)
	return uc
}

// SetAppSettings sets the "app_settings" field.
func (uc *UserCreate) SetAppSettings(m map[string]interface{}) *UserCreate {
	uc.mutation.SetAppSettings(m)
	return uc
}

// SetUserSettings sets the "user_settings" field.
func (uc *UserCreate) SetUserSettings(m map[string]interface{}) *UserCreate {
	uc.mutation.SetUserSettings(m)
	return uc
}

// SetSearchText sets the "search_text" field.
func (uc *UserCreate) SetSearchText(s string) *UserCreate {
	uc.mutation.SetSearchText(s)
	return uc
}

// SetNillableSearchText sets the "search_text" field if the given value is not nil.
func (uc *UserCreate) SetNillableSearchText(s *string) *UserCreate {
	if s != nil {
		uc.SetSearchText(*s)
	}
	return uc
}

// SetRelevanceScore sets the "relevance_score" field.
func (uc *UserCreate) SetRelevanceScore(f float64) *UserCreate {
	uc.mutation.SetRelevanceScore(f)
	return uc
}

// SetNillableRelevanceScore sets the "relevance_score" field if the given value is not nil.
func (uc *UserCreate) SetNillableRelevanceScore(f *float64) *UserCreate {
	if f != nil {
		uc.SetRelevanceScore(*f)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(s string) *UserCreate {
	uc.mutation.SetID(s)
	return uc
}

// AddUserBusinessIDs adds the "userBusinesses" edge to the UserBusiness entity by IDs.
func (uc *UserCreate) AddUserBusinessIDs(ids ...string) *UserCreate {
	uc.mutation.AddUserBusinessIDs(ids...)
	return uc
}

// AddUserBusinesses adds the "userBusinesses" edges to the UserBusiness entity.
func (uc *UserCreate) AddUserBusinesses(u ...*UserBusiness) *UserCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddUserBusinessIDs(ids...)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (uc *UserCreate) AddCommentIDs(ids ...string) *UserCreate {
	uc.mutation.AddCommentIDs(ids...)
	return uc
}

// AddComments adds the "comments" edges to the Comment entity.
func (uc *UserCreate) AddComments(c ...*Comment) *UserCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uc.AddCommentIDs(ids...)
}

// AddLikeIDs adds the "likes" edge to the Like entity by IDs.
func (uc *UserCreate) AddLikeIDs(ids ...string) *UserCreate {
	uc.mutation.AddLikeIDs(ids...)
	return uc
}

// AddLikes adds the "likes" edges to the Like entity.
func (uc *UserCreate) AddLikes(l ...*Like) *UserCreate {
	ids := make([]string, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return uc.AddLikeIDs(ids...)
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (uc *UserCreate) AddPostIDs(ids ...string) *UserCreate {
	uc.mutation.AddPostIDs(ids...)
	return uc
}

// AddPosts adds the "posts" edges to the Post entity.
func (uc *UserCreate) AddPosts(p ...*Post) *UserCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uc.AddPostIDs(ids...)
}

// AddFollowedUserIDs adds the "followedUsers" edge to the UserFollowUser entity by IDs.
func (uc *UserCreate) AddFollowedUserIDs(ids ...string) *UserCreate {
	uc.mutation.AddFollowedUserIDs(ids...)
	return uc
}

// AddFollowedUsers adds the "followedUsers" edges to the UserFollowUser entity.
func (uc *UserCreate) AddFollowedUsers(u ...*UserFollowUser) *UserCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddFollowedUserIDs(ids...)
}

// AddFollowerUserIDs adds the "followerUsers" edge to the UserFollowUser entity by IDs.
func (uc *UserCreate) AddFollowerUserIDs(ids ...string) *UserCreate {
	uc.mutation.AddFollowerUserIDs(ids...)
	return uc
}

// AddFollowerUsers adds the "followerUsers" edges to the UserFollowUser entity.
func (uc *UserCreate) AddFollowerUsers(u ...*UserFollowUser) *UserCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddFollowerUserIDs(ids...)
}

// AddFollowedBusinessIDs adds the "followedBusinesses" edge to the UserFollowBusiness entity by IDs.
func (uc *UserCreate) AddFollowedBusinessIDs(ids ...string) *UserCreate {
	uc.mutation.AddFollowedBusinessIDs(ids...)
	return uc
}

// AddFollowedBusinesses adds the "followedBusinesses" edges to the UserFollowBusiness entity.
func (uc *UserCreate) AddFollowedBusinesses(u ...*UserFollowBusiness) *UserCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddFollowedBusinessIDs(ids...)
}

// AddFollowerBusinessIDs adds the "followerBusinesses" edge to the BusinessFollowUser entity by IDs.
func (uc *UserCreate) AddFollowerBusinessIDs(ids ...string) *UserCreate {
	uc.mutation.AddFollowerBusinessIDs(ids...)
	return uc
}

// AddFollowerBusinesses adds the "followerBusinesses" edges to the BusinessFollowUser entity.
func (uc *UserCreate) AddFollowerBusinesses(b ...*BusinessFollowUser) *UserCreate {
	ids := make([]string, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return uc.AddFollowerBusinessIDs(ids...)
}

// AddReviewIDs adds the "reviews" edge to the Review entity by IDs.
func (uc *UserCreate) AddReviewIDs(ids ...string) *UserCreate {
	uc.mutation.AddReviewIDs(ids...)
	return uc
}

// AddReviews adds the "reviews" edges to the Review entity.
func (uc *UserCreate) AddReviews(r ...*Review) *UserCreate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uc.AddReviewIDs(ids...)
}

// AddBookingIDs adds the "bookings" edge to the Booking entity by IDs.
func (uc *UserCreate) AddBookingIDs(ids ...string) *UserCreate {
	uc.mutation.AddBookingIDs(ids...)
	return uc
}

// AddBookings adds the "bookings" edges to the Booking entity.
func (uc *UserCreate) AddBookings(b ...*Booking) *UserCreate {
	ids := make([]string, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return uc.AddBookingIDs(ids...)
}

// AddReservationIDs adds the "reservations" edge to the Reservation entity by IDs.
func (uc *UserCreate) AddReservationIDs(ids ...string) *UserCreate {
	uc.mutation.AddReservationIDs(ids...)
	return uc
}

// AddReservations adds the "reservations" edges to the Reservation entity.
func (uc *UserCreate) AddReservations(r ...*Reservation) *UserCreate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uc.AddReservationIDs(ids...)
}

// AddHelpIDs adds the "helps" edge to the Help entity by IDs.
func (uc *UserCreate) AddHelpIDs(ids ...string) *UserCreate {
	uc.mutation.AddHelpIDs(ids...)
	return uc
}

// AddHelps adds the "helps" edges to the Help entity.
func (uc *UserCreate) AddHelps(h ...*Help) *UserCreate {
	ids := make([]string, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return uc.AddHelpIDs(ids...)
}

// AddCategoryIDs adds the "categories" edge to the Category entity by IDs.
func (uc *UserCreate) AddCategoryIDs(ids ...string) *UserCreate {
	uc.mutation.AddCategoryIDs(ids...)
	return uc
}

// AddCategories adds the "categories" edges to the Category entity.
func (uc *UserCreate) AddCategories(c ...*Category) *UserCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uc.AddCategoryIDs(ids...)
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (uc *UserCreate) AddEventIDs(ids ...string) *UserCreate {
	uc.mutation.AddEventIDs(ids...)
	return uc
}

// AddEvents adds the "events" edges to the Event entity.
func (uc *UserCreate) AddEvents(e ...*Event) *UserCreate {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uc.AddEventIDs(ids...)
}

// AddPlaceIDs adds the "places" edge to the Place entity by IDs.
func (uc *UserCreate) AddPlaceIDs(ids ...string) *UserCreate {
	uc.mutation.AddPlaceIDs(ids...)
	return uc
}

// AddPlaces adds the "places" edges to the Place entity.
func (uc *UserCreate) AddPlaces(p ...*Place) *UserCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uc.AddPlaceIDs(ids...)
}

// AddCategoryAssignmentIDs adds the "categoryAssignments" edge to the CategoryAssignment entity by IDs.
func (uc *UserCreate) AddCategoryAssignmentIDs(ids ...string) *UserCreate {
	uc.mutation.AddCategoryAssignmentIDs(ids...)
	return uc
}

// AddCategoryAssignments adds the "categoryAssignments" edges to the CategoryAssignment entity.
func (uc *UserCreate) AddCategoryAssignments(c ...*CategoryAssignment) *UserCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uc.AddCategoryAssignmentIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.CoverImage(); !ok {
		v := user.DefaultCoverImage
		uc.mutation.SetCoverImage(v)
	}
	if _, ok := uc.mutation.Bio(); !ok {
		v := user.DefaultBio
		uc.mutation.SetBio(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.Auth0ID(); !ok {
		return &ValidationError{Name: "auth0_id", err: errors.New(`ent: missing required field "User.auth0_id"`)}
	}
	if _, ok := uc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "User.username"`)}
	}
	if v, ok := uc.mutation.ID(); ok {
		if err := user.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "User.id": %w`, err)}
		}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected User.ID type: %T", _spec.ID.Value)
		}
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeString))
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.Auth0ID(); ok {
		_spec.SetField(user.FieldAuth0ID, field.TypeString, value)
		_node.Auth0ID = value
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uc.mutation.Picture(); ok {
		_spec.SetField(user.FieldPicture, field.TypeString, value)
		_node.Picture = value
	}
	if value, ok := uc.mutation.CoverImage(); ok {
		_spec.SetField(user.FieldCoverImage, field.TypeString, value)
		_node.CoverImage = value
	}
	if value, ok := uc.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := uc.mutation.Website(); ok {
		_spec.SetField(user.FieldWebsite, field.TypeString, value)
		_node.Website = value
	}
	if value, ok := uc.mutation.Location(); ok {
		_spec.SetField(user.FieldLocation, field.TypeString, value)
		_node.Location = value
	}
	if value, ok := uc.mutation.Bio(); ok {
		_spec.SetField(user.FieldBio, field.TypeString, value)
		_node.Bio = value
	}
	if value, ok := uc.mutation.Auth0Data(); ok {
		_spec.SetField(user.FieldAuth0Data, field.TypeJSON, value)
		_node.Auth0Data = value
	}
	if value, ok := uc.mutation.AppSettings(); ok {
		_spec.SetField(user.FieldAppSettings, field.TypeJSON, value)
		_node.AppSettings = value
	}
	if value, ok := uc.mutation.UserSettings(); ok {
		_spec.SetField(user.FieldUserSettings, field.TypeJSON, value)
		_node.UserSettings = value
	}
	if value, ok := uc.mutation.SearchText(); ok {
		_spec.SetField(user.FieldSearchText, field.TypeString, value)
		_node.SearchText = value
	}
	if value, ok := uc.mutation.RelevanceScore(); ok {
		_spec.SetField(user.FieldRelevanceScore, field.TypeFloat64, value)
		_node.RelevanceScore = value
	}
	if nodes := uc.mutation.UserBusinessesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserBusinessesTable,
			Columns: []string{user.UserBusinessesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userbusiness.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CommentsTable,
			Columns: []string{user.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.LikesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.LikesTable,
			Columns: []string{user.LikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(like.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PostsTable,
			Columns: []string{user.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.FollowedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.FollowedUsersTable,
			Columns: []string{user.FollowedUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userfollowuser.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.FollowerUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.FollowerUsersTable,
			Columns: []string{user.FollowerUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userfollowuser.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.FollowedBusinessesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.FollowedBusinessesTable,
			Columns: []string{user.FollowedBusinessesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userfollowbusiness.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.FollowerBusinessesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.FollowerBusinessesTable,
			Columns: []string{user.FollowerBusinessesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(businessfollowuser.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.ReviewsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ReviewsTable,
			Columns: []string{user.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(review.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.BookingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.BookingsTable,
			Columns: []string{user.BookingsColumn},
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
	if nodes := uc.mutation.ReservationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ReservationsTable,
			Columns: []string{user.ReservationsColumn},
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
	if nodes := uc.mutation.HelpsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.HelpsTable,
			Columns: []string{user.HelpsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(help.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.CategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CategoriesTable,
			Columns: []string{user.CategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.EventsTable,
			Columns: []string{user.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.PlacesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PlacesTable,
			Columns: []string{user.PlacesColumn},
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
	if nodes := uc.mutation.CategoryAssignmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CategoryAssignmentsTable,
			Columns: []string{user.CategoryAssignmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(categoryassignment.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
