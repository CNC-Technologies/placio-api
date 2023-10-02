// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/accountsettings"
	"placio-app/ent/accountwallet"
	"placio-app/ent/business"
	"placio-app/ent/businessfollowbusiness"
	"placio-app/ent/businessfollowevent"
	"placio-app/ent/businessfollowuser"
	"placio-app/ent/category"
	"placio-app/ent/categoryassignment"
	"placio-app/ent/event"
	"placio-app/ent/faq"
	"placio-app/ent/notification"
	"placio-app/ent/place"
	"placio-app/ent/placeinventory"
	"placio-app/ent/post"
	"placio-app/ent/rating"
	"placio-app/ent/userbusiness"
	"placio-app/ent/userfollowbusiness"
	"placio-app/ent/website"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BusinessCreate is the builder for creating a Business entity.
type BusinessCreate struct {
	config
	mutation *BusinessMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (bc *BusinessCreate) SetName(s string) *BusinessCreate {
	bc.mutation.SetName(s)
	return bc
}

// SetDescription sets the "description" field.
func (bc *BusinessCreate) SetDescription(s string) *BusinessCreate {
	bc.mutation.SetDescription(s)
	return bc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (bc *BusinessCreate) SetNillableDescription(s *string) *BusinessCreate {
	if s != nil {
		bc.SetDescription(*s)
	}
	return bc
}

// SetPicture sets the "picture" field.
func (bc *BusinessCreate) SetPicture(s string) *BusinessCreate {
	bc.mutation.SetPicture(s)
	return bc
}

// SetNillablePicture sets the "picture" field if the given value is not nil.
func (bc *BusinessCreate) SetNillablePicture(s *string) *BusinessCreate {
	if s != nil {
		bc.SetPicture(*s)
	}
	return bc
}

// SetCoverImage sets the "cover_image" field.
func (bc *BusinessCreate) SetCoverImage(s string) *BusinessCreate {
	bc.mutation.SetCoverImage(s)
	return bc
}

// SetNillableCoverImage sets the "cover_image" field if the given value is not nil.
func (bc *BusinessCreate) SetNillableCoverImage(s *string) *BusinessCreate {
	if s != nil {
		bc.SetCoverImage(*s)
	}
	return bc
}

// SetWebsite sets the "website" field.
func (bc *BusinessCreate) SetWebsite(s string) *BusinessCreate {
	bc.mutation.SetWebsite(s)
	return bc
}

// SetNillableWebsite sets the "website" field if the given value is not nil.
func (bc *BusinessCreate) SetNillableWebsite(s *string) *BusinessCreate {
	if s != nil {
		bc.SetWebsite(*s)
	}
	return bc
}

// SetLocation sets the "location" field.
func (bc *BusinessCreate) SetLocation(s string) *BusinessCreate {
	bc.mutation.SetLocation(s)
	return bc
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (bc *BusinessCreate) SetNillableLocation(s *string) *BusinessCreate {
	if s != nil {
		bc.SetLocation(*s)
	}
	return bc
}

// SetLongitude sets the "longitude" field.
func (bc *BusinessCreate) SetLongitude(s string) *BusinessCreate {
	bc.mutation.SetLongitude(s)
	return bc
}

// SetNillableLongitude sets the "longitude" field if the given value is not nil.
func (bc *BusinessCreate) SetNillableLongitude(s *string) *BusinessCreate {
	if s != nil {
		bc.SetLongitude(*s)
	}
	return bc
}

// SetMapCoordinates sets the "map_coordinates" field.
func (bc *BusinessCreate) SetMapCoordinates(m map[string]interface{}) *BusinessCreate {
	bc.mutation.SetMapCoordinates(m)
	return bc
}

// SetLatitude sets the "latitude" field.
func (bc *BusinessCreate) SetLatitude(s string) *BusinessCreate {
	bc.mutation.SetLatitude(s)
	return bc
}

// SetNillableLatitude sets the "latitude" field if the given value is not nil.
func (bc *BusinessCreate) SetNillableLatitude(s *string) *BusinessCreate {
	if s != nil {
		bc.SetLatitude(*s)
	}
	return bc
}

// SetEmail sets the "email" field.
func (bc *BusinessCreate) SetEmail(s string) *BusinessCreate {
	bc.mutation.SetEmail(s)
	return bc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (bc *BusinessCreate) SetNillableEmail(s *string) *BusinessCreate {
	if s != nil {
		bc.SetEmail(*s)
	}
	return bc
}

// SetPhone sets the "phone" field.
func (bc *BusinessCreate) SetPhone(s string) *BusinessCreate {
	bc.mutation.SetPhone(s)
	return bc
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (bc *BusinessCreate) SetNillablePhone(s *string) *BusinessCreate {
	if s != nil {
		bc.SetPhone(*s)
	}
	return bc
}

// SetBusinessSettings sets the "business_settings" field.
func (bc *BusinessCreate) SetBusinessSettings(m map[string]interface{}) *BusinessCreate {
	bc.mutation.SetBusinessSettings(m)
	return bc
}

// SetURL sets the "url" field.
func (bc *BusinessCreate) SetURL(s string) *BusinessCreate {
	bc.mutation.SetURL(s)
	return bc
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (bc *BusinessCreate) SetNillableURL(s *string) *BusinessCreate {
	if s != nil {
		bc.SetURL(*s)
	}
	return bc
}

// SetSearchText sets the "search_text" field.
func (bc *BusinessCreate) SetSearchText(s string) *BusinessCreate {
	bc.mutation.SetSearchText(s)
	return bc
}

// SetNillableSearchText sets the "search_text" field if the given value is not nil.
func (bc *BusinessCreate) SetNillableSearchText(s *string) *BusinessCreate {
	if s != nil {
		bc.SetSearchText(*s)
	}
	return bc
}

// SetRelevanceScore sets the "relevance_score" field.
func (bc *BusinessCreate) SetRelevanceScore(f float64) *BusinessCreate {
	bc.mutation.SetRelevanceScore(f)
	return bc
}

// SetNillableRelevanceScore sets the "relevance_score" field if the given value is not nil.
func (bc *BusinessCreate) SetNillableRelevanceScore(f *float64) *BusinessCreate {
	if f != nil {
		bc.SetRelevanceScore(*f)
	}
	return bc
}

// SetFollowerCount sets the "follower_count" field.
func (bc *BusinessCreate) SetFollowerCount(i int) *BusinessCreate {
	bc.mutation.SetFollowerCount(i)
	return bc
}

// SetNillableFollowerCount sets the "follower_count" field if the given value is not nil.
func (bc *BusinessCreate) SetNillableFollowerCount(i *int) *BusinessCreate {
	if i != nil {
		bc.SetFollowerCount(*i)
	}
	return bc
}

// SetFollowingCount sets the "following_count" field.
func (bc *BusinessCreate) SetFollowingCount(i int) *BusinessCreate {
	bc.mutation.SetFollowingCount(i)
	return bc
}

// SetNillableFollowingCount sets the "following_count" field if the given value is not nil.
func (bc *BusinessCreate) SetNillableFollowingCount(i *int) *BusinessCreate {
	if i != nil {
		bc.SetFollowingCount(*i)
	}
	return bc
}

// SetID sets the "id" field.
func (bc *BusinessCreate) SetID(s string) *BusinessCreate {
	bc.mutation.SetID(s)
	return bc
}

// AddUserBusinessIDs adds the "userBusinesses" edge to the UserBusiness entity by IDs.
func (bc *BusinessCreate) AddUserBusinessIDs(ids ...string) *BusinessCreate {
	bc.mutation.AddUserBusinessIDs(ids...)
	return bc
}

// AddUserBusinesses adds the "userBusinesses" edges to the UserBusiness entity.
func (bc *BusinessCreate) AddUserBusinesses(u ...*UserBusiness) *BusinessCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return bc.AddUserBusinessIDs(ids...)
}

// SetBusinessAccountSettingsID sets the "business_account_settings" edge to the AccountSettings entity by ID.
func (bc *BusinessCreate) SetBusinessAccountSettingsID(id string) *BusinessCreate {
	bc.mutation.SetBusinessAccountSettingsID(id)
	return bc
}

// SetNillableBusinessAccountSettingsID sets the "business_account_settings" edge to the AccountSettings entity by ID if the given value is not nil.
func (bc *BusinessCreate) SetNillableBusinessAccountSettingsID(id *string) *BusinessCreate {
	if id != nil {
		bc = bc.SetBusinessAccountSettingsID(*id)
	}
	return bc
}

// SetBusinessAccountSettings sets the "business_account_settings" edge to the AccountSettings entity.
func (bc *BusinessCreate) SetBusinessAccountSettings(a *AccountSettings) *BusinessCreate {
	return bc.SetBusinessAccountSettingsID(a.ID)
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (bc *BusinessCreate) AddPostIDs(ids ...string) *BusinessCreate {
	bc.mutation.AddPostIDs(ids...)
	return bc
}

// AddPosts adds the "posts" edges to the Post entity.
func (bc *BusinessCreate) AddPosts(p ...*Post) *BusinessCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bc.AddPostIDs(ids...)
}

// AddFollowedUserIDs adds the "followedUsers" edge to the BusinessFollowUser entity by IDs.
func (bc *BusinessCreate) AddFollowedUserIDs(ids ...string) *BusinessCreate {
	bc.mutation.AddFollowedUserIDs(ids...)
	return bc
}

// AddFollowedUsers adds the "followedUsers" edges to the BusinessFollowUser entity.
func (bc *BusinessCreate) AddFollowedUsers(b ...*BusinessFollowUser) *BusinessCreate {
	ids := make([]string, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bc.AddFollowedUserIDs(ids...)
}

// AddFollowerUserIDs adds the "followerUsers" edge to the UserFollowBusiness entity by IDs.
func (bc *BusinessCreate) AddFollowerUserIDs(ids ...string) *BusinessCreate {
	bc.mutation.AddFollowerUserIDs(ids...)
	return bc
}

// AddFollowerUsers adds the "followerUsers" edges to the UserFollowBusiness entity.
func (bc *BusinessCreate) AddFollowerUsers(u ...*UserFollowBusiness) *BusinessCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return bc.AddFollowerUserIDs(ids...)
}

// AddFollowedBusinessIDs adds the "followedBusinesses" edge to the BusinessFollowBusiness entity by IDs.
func (bc *BusinessCreate) AddFollowedBusinessIDs(ids ...string) *BusinessCreate {
	bc.mutation.AddFollowedBusinessIDs(ids...)
	return bc
}

// AddFollowedBusinesses adds the "followedBusinesses" edges to the BusinessFollowBusiness entity.
func (bc *BusinessCreate) AddFollowedBusinesses(b ...*BusinessFollowBusiness) *BusinessCreate {
	ids := make([]string, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bc.AddFollowedBusinessIDs(ids...)
}

// AddFollowerBusinessIDs adds the "followerBusinesses" edge to the BusinessFollowBusiness entity by IDs.
func (bc *BusinessCreate) AddFollowerBusinessIDs(ids ...string) *BusinessCreate {
	bc.mutation.AddFollowerBusinessIDs(ids...)
	return bc
}

// AddFollowerBusinesses adds the "followerBusinesses" edges to the BusinessFollowBusiness entity.
func (bc *BusinessCreate) AddFollowerBusinesses(b ...*BusinessFollowBusiness) *BusinessCreate {
	ids := make([]string, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bc.AddFollowerBusinessIDs(ids...)
}

// AddPlaceIDs adds the "places" edge to the Place entity by IDs.
func (bc *BusinessCreate) AddPlaceIDs(ids ...string) *BusinessCreate {
	bc.mutation.AddPlaceIDs(ids...)
	return bc
}

// AddPlaces adds the "places" edges to the Place entity.
func (bc *BusinessCreate) AddPlaces(p ...*Place) *BusinessCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bc.AddPlaceIDs(ids...)
}

// AddCategoryIDs adds the "categories" edge to the Category entity by IDs.
func (bc *BusinessCreate) AddCategoryIDs(ids ...string) *BusinessCreate {
	bc.mutation.AddCategoryIDs(ids...)
	return bc
}

// AddCategories adds the "categories" edges to the Category entity.
func (bc *BusinessCreate) AddCategories(c ...*Category) *BusinessCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return bc.AddCategoryIDs(ids...)
}

// AddCategoryAssignmentIDs adds the "categoryAssignments" edge to the CategoryAssignment entity by IDs.
func (bc *BusinessCreate) AddCategoryAssignmentIDs(ids ...string) *BusinessCreate {
	bc.mutation.AddCategoryAssignmentIDs(ids...)
	return bc
}

// AddCategoryAssignments adds the "categoryAssignments" edges to the CategoryAssignment entity.
func (bc *BusinessCreate) AddCategoryAssignments(c ...*CategoryAssignment) *BusinessCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return bc.AddCategoryAssignmentIDs(ids...)
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (bc *BusinessCreate) AddEventIDs(ids ...string) *BusinessCreate {
	bc.mutation.AddEventIDs(ids...)
	return bc
}

// AddEvents adds the "events" edges to the Event entity.
func (bc *BusinessCreate) AddEvents(e ...*Event) *BusinessCreate {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return bc.AddEventIDs(ids...)
}

// AddBusinessFollowEventIDs adds the "businessFollowEvents" edge to the BusinessFollowEvent entity by IDs.
func (bc *BusinessCreate) AddBusinessFollowEventIDs(ids ...string) *BusinessCreate {
	bc.mutation.AddBusinessFollowEventIDs(ids...)
	return bc
}

// AddBusinessFollowEvents adds the "businessFollowEvents" edges to the BusinessFollowEvent entity.
func (bc *BusinessCreate) AddBusinessFollowEvents(b ...*BusinessFollowEvent) *BusinessCreate {
	ids := make([]string, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bc.AddBusinessFollowEventIDs(ids...)
}

// AddFaqIDs adds the "faqs" edge to the FAQ entity by IDs.
func (bc *BusinessCreate) AddFaqIDs(ids ...string) *BusinessCreate {
	bc.mutation.AddFaqIDs(ids...)
	return bc
}

// AddFaqs adds the "faqs" edges to the FAQ entity.
func (bc *BusinessCreate) AddFaqs(f ...*FAQ) *BusinessCreate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return bc.AddFaqIDs(ids...)
}

// AddRatingIDs adds the "ratings" edge to the Rating entity by IDs.
func (bc *BusinessCreate) AddRatingIDs(ids ...string) *BusinessCreate {
	bc.mutation.AddRatingIDs(ids...)
	return bc
}

// AddRatings adds the "ratings" edges to the Rating entity.
func (bc *BusinessCreate) AddRatings(r ...*Rating) *BusinessCreate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return bc.AddRatingIDs(ids...)
}

// AddPlaceInventoryIDs adds the "place_inventories" edge to the PlaceInventory entity by IDs.
func (bc *BusinessCreate) AddPlaceInventoryIDs(ids ...string) *BusinessCreate {
	bc.mutation.AddPlaceInventoryIDs(ids...)
	return bc
}

// AddPlaceInventories adds the "place_inventories" edges to the PlaceInventory entity.
func (bc *BusinessCreate) AddPlaceInventories(p ...*PlaceInventory) *BusinessCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bc.AddPlaceInventoryIDs(ids...)
}

// SetWebsitesID sets the "websites" edge to the Website entity by ID.
func (bc *BusinessCreate) SetWebsitesID(id string) *BusinessCreate {
	bc.mutation.SetWebsitesID(id)
	return bc
}

// SetNillableWebsitesID sets the "websites" edge to the Website entity by ID if the given value is not nil.
func (bc *BusinessCreate) SetNillableWebsitesID(id *string) *BusinessCreate {
	if id != nil {
		bc = bc.SetWebsitesID(*id)
	}
	return bc
}

// SetWebsites sets the "websites" edge to the Website entity.
func (bc *BusinessCreate) SetWebsites(w *Website) *BusinessCreate {
	return bc.SetWebsitesID(w.ID)
}

// AddNotificationIDs adds the "notifications" edge to the Notification entity by IDs.
func (bc *BusinessCreate) AddNotificationIDs(ids ...string) *BusinessCreate {
	bc.mutation.AddNotificationIDs(ids...)
	return bc
}

// AddNotifications adds the "notifications" edges to the Notification entity.
func (bc *BusinessCreate) AddNotifications(n ...*Notification) *BusinessCreate {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return bc.AddNotificationIDs(ids...)
}

// SetWalletID sets the "wallet" edge to the AccountWallet entity by ID.
func (bc *BusinessCreate) SetWalletID(id string) *BusinessCreate {
	bc.mutation.SetWalletID(id)
	return bc
}

// SetNillableWalletID sets the "wallet" edge to the AccountWallet entity by ID if the given value is not nil.
func (bc *BusinessCreate) SetNillableWalletID(id *string) *BusinessCreate {
	if id != nil {
		bc = bc.SetWalletID(*id)
	}
	return bc
}

// SetWallet sets the "wallet" edge to the AccountWallet entity.
func (bc *BusinessCreate) SetWallet(a *AccountWallet) *BusinessCreate {
	return bc.SetWalletID(a.ID)
}

// Mutation returns the BusinessMutation object of the builder.
func (bc *BusinessCreate) Mutation() *BusinessMutation {
	return bc.mutation
}

// Save creates the Business in the database.
func (bc *BusinessCreate) Save(ctx context.Context) (*Business, error) {
	bc.defaults()
	return withHooks(ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BusinessCreate) SaveX(ctx context.Context) *Business {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BusinessCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BusinessCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BusinessCreate) defaults() {
	if _, ok := bc.mutation.CoverImage(); !ok {
		v := business.DefaultCoverImage
		bc.mutation.SetCoverImage(v)
	}
	if _, ok := bc.mutation.FollowerCount(); !ok {
		v := business.DefaultFollowerCount
		bc.mutation.SetFollowerCount(v)
	}
	if _, ok := bc.mutation.FollowingCount(); !ok {
		v := business.DefaultFollowingCount
		bc.mutation.SetFollowingCount(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BusinessCreate) check() error {
	if _, ok := bc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Business.name"`)}
	}
	if _, ok := bc.mutation.FollowerCount(); !ok {
		return &ValidationError{Name: "follower_count", err: errors.New(`ent: missing required field "Business.follower_count"`)}
	}
	if _, ok := bc.mutation.FollowingCount(); !ok {
		return &ValidationError{Name: "following_count", err: errors.New(`ent: missing required field "Business.following_count"`)}
	}
	if v, ok := bc.mutation.ID(); ok {
		if err := business.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Business.id": %w`, err)}
		}
	}
	return nil
}

func (bc *BusinessCreate) sqlSave(ctx context.Context) (*Business, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Business.ID type: %T", _spec.ID.Value)
		}
	}
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BusinessCreate) createSpec() (*Business, *sqlgraph.CreateSpec) {
	var (
		_node = &Business{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(business.Table, sqlgraph.NewFieldSpec(business.FieldID, field.TypeString))
	)
	if id, ok := bc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := bc.mutation.Name(); ok {
		_spec.SetField(business.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := bc.mutation.Description(); ok {
		_spec.SetField(business.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := bc.mutation.Picture(); ok {
		_spec.SetField(business.FieldPicture, field.TypeString, value)
		_node.Picture = value
	}
	if value, ok := bc.mutation.CoverImage(); ok {
		_spec.SetField(business.FieldCoverImage, field.TypeString, value)
		_node.CoverImage = value
	}
	if value, ok := bc.mutation.Website(); ok {
		_spec.SetField(business.FieldWebsite, field.TypeString, value)
		_node.Website = value
	}
	if value, ok := bc.mutation.Location(); ok {
		_spec.SetField(business.FieldLocation, field.TypeString, value)
		_node.Location = value
	}
	if value, ok := bc.mutation.Longitude(); ok {
		_spec.SetField(business.FieldLongitude, field.TypeString, value)
		_node.Longitude = value
	}
	if value, ok := bc.mutation.MapCoordinates(); ok {
		_spec.SetField(business.FieldMapCoordinates, field.TypeJSON, value)
		_node.MapCoordinates = value
	}
	if value, ok := bc.mutation.Latitude(); ok {
		_spec.SetField(business.FieldLatitude, field.TypeString, value)
		_node.Latitude = value
	}
	if value, ok := bc.mutation.Email(); ok {
		_spec.SetField(business.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := bc.mutation.Phone(); ok {
		_spec.SetField(business.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if value, ok := bc.mutation.BusinessSettings(); ok {
		_spec.SetField(business.FieldBusinessSettings, field.TypeJSON, value)
		_node.BusinessSettings = value
	}
	if value, ok := bc.mutation.URL(); ok {
		_spec.SetField(business.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := bc.mutation.SearchText(); ok {
		_spec.SetField(business.FieldSearchText, field.TypeString, value)
		_node.SearchText = value
	}
	if value, ok := bc.mutation.RelevanceScore(); ok {
		_spec.SetField(business.FieldRelevanceScore, field.TypeFloat64, value)
		_node.RelevanceScore = value
	}
	if value, ok := bc.mutation.FollowerCount(); ok {
		_spec.SetField(business.FieldFollowerCount, field.TypeInt, value)
		_node.FollowerCount = value
	}
	if value, ok := bc.mutation.FollowingCount(); ok {
		_spec.SetField(business.FieldFollowingCount, field.TypeInt, value)
		_node.FollowingCount = value
	}
	if nodes := bc.mutation.UserBusinessesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   business.UserBusinessesTable,
			Columns: []string{business.UserBusinessesColumn},
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
	if nodes := bc.mutation.BusinessAccountSettingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   business.BusinessAccountSettingsTable,
			Columns: []string{business.BusinessAccountSettingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(accountsettings.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   business.PostsTable,
			Columns: []string{business.PostsColumn},
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
	if nodes := bc.mutation.FollowedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   business.FollowedUsersTable,
			Columns: []string{business.FollowedUsersColumn},
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
	if nodes := bc.mutation.FollowerUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   business.FollowerUsersTable,
			Columns: []string{business.FollowerUsersColumn},
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
	if nodes := bc.mutation.FollowedBusinessesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   business.FollowedBusinessesTable,
			Columns: []string{business.FollowedBusinessesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(businessfollowbusiness.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.FollowerBusinessesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   business.FollowerBusinessesTable,
			Columns: []string{business.FollowerBusinessesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(businessfollowbusiness.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.PlacesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   business.PlacesTable,
			Columns: []string{business.PlacesColumn},
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
	if nodes := bc.mutation.CategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   business.CategoriesTable,
			Columns: []string{business.CategoriesColumn},
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
	if nodes := bc.mutation.CategoryAssignmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   business.CategoryAssignmentsTable,
			Columns: []string{business.CategoryAssignmentsColumn},
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
	if nodes := bc.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   business.EventsTable,
			Columns: []string{business.EventsColumn},
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
	if nodes := bc.mutation.BusinessFollowEventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   business.BusinessFollowEventsTable,
			Columns: []string{business.BusinessFollowEventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(businessfollowevent.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.FaqsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   business.FaqsTable,
			Columns: []string{business.FaqsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(faq.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.RatingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   business.RatingsTable,
			Columns: []string{business.RatingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rating.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.PlaceInventoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   business.PlaceInventoriesTable,
			Columns: []string{business.PlaceInventoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(placeinventory.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.WebsitesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   business.WebsitesTable,
			Columns: []string{business.WebsitesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(website.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.NotificationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   business.NotificationsTable,
			Columns: business.NotificationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notification.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.WalletIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   business.WalletTable,
			Columns: []string{business.WalletColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(accountwallet.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BusinessCreateBulk is the builder for creating many Business entities in bulk.
type BusinessCreateBulk struct {
	config
	err      error
	builders []*BusinessCreate
}

// Save creates the Business entities in the database.
func (bcb *BusinessCreateBulk) Save(ctx context.Context) ([]*Business, error) {
	if bcb.err != nil {
		return nil, bcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Business, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BusinessMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BusinessCreateBulk) SaveX(ctx context.Context) []*Business {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BusinessCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BusinessCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
