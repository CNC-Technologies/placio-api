package business

import (
	"context"
	"errors"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
	"log"
	"placio-app/domains/events_management"
	"placio-app/domains/places"
	"placio-app/domains/search"
	"placio-app/ent"
	"placio-app/ent/business"
	"placio-app/ent/businessfollowbusiness"
	"placio-app/ent/businessfollowuser"
	"placio-app/ent/predicate"
	"placio-app/ent/user"
	"placio-app/ent/userbusiness"
	"placio-app/utility"
	appErrors "placio-pkg/errors"
	"strings"
	"time"
)

type BusinessAccountService interface {
	CreateBusinessAccount(ctx context.Context, businessData *BusinessDto) (*ent.Business, error)
	GetBusinessAccount(ctx context.Context, businessAccountID string) (*ent.Business, error)
	GetUserBusinessAccounts(ctx context.Context, userId string) ([]*ent.Business, error)
	AssociateUserWithBusinessAccount(ctx context.Context, adminUser, userID, businessAccountID, role string) error
	GetBusinessAccountsForUser(ctx context.Context, userID string) ([]*ent.UserBusiness, error)
	ListBusinessAccounts(ctx context.Context, page, pageSize int, sortBy string, filters ...predicate.Business) ([]*ent.Business, error)
	RemoveUserFromBusinessAccount(ctx context.Context, userID, businessAccountID string) error
	TransferBusinessAccountOwnership(ctx context.Context, currentOwnerID, newOwnerID, businessAccountID string) error
	FollowUser(ctx context.Context, businessID string, userID string) error
	UpdateBusinessAccount(ctx context.Context, businessID string, businessData map[string]interface{}) (*ent.Business, error)
	FollowBusiness(ctx context.Context, followerID string, followedID string) error
	UnfollowUser(ctx context.Context, businessID string, userID string) error
	UnfollowBusiness(ctx context.Context, followerID string, followedID string) error
	GetFollowedContents(ctx context.Context, businessID string) ([]*ent.Post, error)
	GetPlacesAndEventsAssociatedWithBusinessAccount(c context.Context, relatedType string, businessId string) (BusinessAccountPlacesAndEvents, error)
	AddTeamMember(ctx context.Context, adminUser, userID, businessAccountID, role, permissions string) error
	ListTeamMembers(ctx context.Context, businessAccountID string) ([]*ent.UserBusiness, error)
	EditTeamMember(ctx context.Context, userID, businessAccountID, newRole, newPermissions string) error
	RemoveTeamMember(ctx context.Context, userID, businessAccountID string) error
	SearchTeamMembers(ctx context.Context, businessAccountID, searchText string) ([]*ent.UserBusiness, error)
	//CanPerformAction(ctx context.Context, userID, businessAccountID, action string) (bool, error)
}

type BusinessAccountServiceImpl struct {
	store         *ent.Business
	client        *ent.Client
	searchService search.SearchService
	cacheService  *utility.RedisClient
	placesService places.PlaceService
	eventService  events_management.IEventService
}

func NewBusinessAccountService(client *ent.Client, searchService search.SearchService, cache *utility.RedisClient, service places.PlaceService) BusinessAccountService {
	return &BusinessAccountServiceImpl{client: client, store: &ent.Business{}, searchService: searchService, cacheService: cache, placesService: service}
}

func (s *BusinessAccountServiceImpl) FollowUser(ctx context.Context, businessID string, userID string) error {
	_, err := s.client.BusinessFollowUser.
		Create().
		SetBusinessID(businessID).
		SetUserID(userID).
		Save(ctx)
	return err
}

func (s *BusinessAccountServiceImpl) GetPlacesAndEventsAssociatedWithBusinessAccount(c context.Context, relatedType string, businessId string) (BusinessAccountPlacesAndEvents, error) {
	if relatedType == "" {
		relatedType = "all"
	}

	var businessAccountPlacesAndEvents BusinessAccountPlacesAndEvents
	businessAccount, err := s.GetBusinessAccount(c, businessId)
	if err != nil {
		return businessAccountPlacesAndEvents, err
	}

	if relatedType == "all" || relatedType == "places" {
		places, err := s.placesService.GetPlacesAssociatedWithBusinessAccount(c, businessAccount.ID)
		if err != nil {
			return businessAccountPlacesAndEvents, err
		}
		businessAccountPlacesAndEvents.Places = places
	}

	if relatedType == "all" || relatedType == "events" {
		//events, err := s.eventService.GetEventsAssociatedWithBusinessAccount(c, businessAccount.ID)
		//if err != nil {
		//	return businessAccountPlacesAndEvents, err
		//}
		//businessAccountPlacesAndEvents.Events = events
	}

	return businessAccountPlacesAndEvents, nil
}

func (s *BusinessAccountServiceImpl) FollowBusiness(ctx context.Context, followerID string, followedID string) error {
	_, err := s.client.BusinessFollowBusiness.
		Create().
		SetFollowerID(followerID).
		SetFollowedID(followedID).
		Save(ctx)
	return err
}

func (s *BusinessAccountServiceImpl) UnfollowUser(ctx context.Context, businessID string, userID string) error {
	bf, err := s.client.BusinessFollowUser.
		Query().
		Where(businessfollowuser.HasBusinessWith(business.ID(businessID)), businessfollowuser.HasUserWith(user.ID(userID))).
		Only(ctx)
	if err != nil {
		return err
	}
	return s.client.BusinessFollowUser.DeleteOne(bf).Exec(ctx)
}

func (s *BusinessAccountServiceImpl) UnfollowBusiness(ctx context.Context, followerID string, followedID string) error {
	bf, err := s.client.BusinessFollowBusiness.
		Query().
		Where(businessfollowbusiness.HasFollowerWith(business.ID(followerID)), businessfollowbusiness.HasFollowedWith(business.ID(followedID))).
		Only(ctx)
	if err != nil {
		return err
	}
	return s.client.BusinessFollowBusiness.DeleteOne(bf).Exec(ctx)
}

func (s *BusinessAccountServiceImpl) GetFollowedContents(ctx context.Context, businessID string) ([]*ent.Post, error) {
	followedUsers, err := s.client.Business.
		Query().
		Where(business.IDEQ(businessID)).
		QueryFollowedUsers().
		All(ctx)
	if err != nil {
		return nil, err
	}

	followedBusinesses, err := s.client.Business.
		Query().
		Where(business.IDEQ(businessID)).
		QueryFollowedBusinesses().
		All(ctx)
	if err != nil {
		return nil, err
	}

	var allPosts []*ent.Post
	for _, followedUser := range followedUsers {
		posts, err := s.client.User.
			Query().
			Where(user.IDEQ(followedUser.ID)).
			QueryPosts().
			All(ctx)
		if err != nil {
			return nil, err
		}
		allPosts = append(allPosts, posts...)
	}

	for _, followedBusiness := range followedBusinesses {
		posts, err := s.client.Business.
			Query().
			Where(business.IDEQ(followedBusiness.ID)).
			QueryPosts().
			All(ctx)
		if err != nil {
			return nil, err
		}
		allPosts = append(allPosts, posts...)
	}

	return allPosts, nil
}

func extractFirstWord(str string) string {
	words := strings.Fields(str)
	if len(words) > 0 {
		return words[0]
	}
	return str
}

// CreateBusinessAccount creates a new Business Account and associates it with a user.
func (s *BusinessAccountServiceImpl) CreateBusinessAccount(ctx context.Context, businessData *BusinessDto) (*ent.Business, error) {
	// Validate inputs
	// grab the user id from the context
	userID := ctx.Value("user").(string)

	// get the user
	user, err := s.client.User.
		Query().
		Where(user.IDEQ(userID)).
		Only(ctx)

	if err != nil {
		log.Println("error getting user", err)
		return nil, err
	}

	if businessData.Name == "" {
		return nil, errors.New("business account name cannot be empty")
	}
	//if businessData.Role != "owner" && businessData.Role != "admin" && businessData.Role != "member" {
	//	return nil, errors.New("invalid role")
	//}

	// Create a new transaction
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("error starting transaction: %w", err)
	}

	// Create business account
	businessAccount, err := tx.Business.
		Create().
		SetID(uuid.New().String()).
		SetName(strings.ToLower(businessData.Name)).
		SetDescription(businessData.Description).
		SetWebsite(businessData.Website).
		SetEmail(businessData.Email).
		SetPhone(businessData.Phone).
		SetLocation(businessData.Location).
		SetPicture(businessData.Picture).
		SetCoverImage(businessData.CoverImage).
		Save(ctx)

	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error creating business account: %w", err)
	}

	// add the business account to the search index
	err = s.searchService.CreateOrUpdateBusiness(ctx, businessAccount)
	if err != nil {
		log.Println("error adding buisness account to search index", err)
		log.Println("error adding buisness account to search index", err)
		tx.Rollback()
		return nil, fmt.Errorf("error creating business account: %w", err)
	}

	// Create user-business relationship
	_, err = tx.UserBusiness.
		Create().
		SetID(uuid.New().String()).
		SetUser(user).
		SetBusiness(businessAccount).
		SetRole("admin").
		Save(ctx)

	if err != nil {
		log.Println("error creating user-business relationship", err)
		tx.Rollback()
		return nil, fmt.Errorf("error creating user-business relationship: %w", err)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("error committing transaction: %w", err)
	}

	domainName := extractFirstWord(businessData.Name) + uuid.New().String()[:6]

	// I need to create a business website
	_, err = s.client.Website.
		Create().
		SetID(uuid.New().String()).
		SetBusiness(businessAccount).
		SetDomainName(domainName).
		SetEmail(businessData.Email).
		SetTitle(businessData.Name).
		SetLogo("https://res.cloudinary.com/placio/image/upload/v1698229500/defaults/samantha-gades-fIHozNWfcvs-unsplash_kto0l3.jpg").
		SetBannerTwoSectionBackgroundImage("https://res.cloudinary.com/placio/image/upload/v1698229499/defaults/francesca-saraco-_dS27XGgRyQ-unsplash_jnuvcx.jpg").
		SetBannerSectionBackgroundImage("https://res.cloudinary.com/placio/image/upload/v1698229501/defaults/fernando-alvarez-rodriguez-M7GddPqJowg-unsplash_gchmsc.jpg").
		SetBannerTwoRightSideImage("https://res.cloudinary.com/placio/image/upload/v1698229500/defaults/samantha-gades-fIHozNWfcvs-unsplash_kto0l3.jpg").
		SetLastUpdated(time.Now()).
		Save(ctx)

	if err != nil {
		log.Println("error creating business website", err)
		return nil, fmt.Errorf("error creating business website: %w", err)
	}

	log.Println("businessAccount", businessAccount)

	// Now we need to fetch the created business account with its relationships
	businessAccount, err = s.client.Business.
		Query().
		Where(business.ID(businessAccount.ID)).
		WithUserBusinesses().
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("error fetching created business account: %w", err)
	}

	//// add business account to search index
	//err = s.searchService.CreateOrUpdateBusiness(ctx, businessAccount)
	//if err != nil {
	//	return nil, fmt.Errorf("error adding buisness account to search index: %w", err)
	//}
	return businessAccount, nil
}

func (s *BusinessAccountServiceImpl) CanPerformAction(ctx context.Context, userID, businessAccountID, action string) (bool, error) {
	relationship, err := s.client.UserBusiness.Query().Where(userbusiness.HasUserWith(user.ID(userID)), userbusiness.HasBusinessWith(business.ID(businessAccountID))).Only(ctx)
	if err != nil {
		return false, err
	}

	// Check if the user's role within the business account allows the action
	// This will depend on how you define the capabilities of each role
	if relationship.Role == "admin" && action == "delete_account" {
		return true, nil
	}

	return false, nil
}

func (s *BusinessAccountServiceImpl) GetBusinessAccount(ctx context.Context, businessAccountID string) (*ent.Business, error) {
	if businessAccountID == "" {
		return nil, errors.New("businessAccountID cannot be nil")
	}

	businessAccount, err := s.client.Business.
		Query().
		Where(business.IDEQ(businessAccountID)).
		WithFollowerUsers(func(query *ent.UserFollowBusinessQuery) {
			query.WithBusiness()
			query.WithUser()
		}).
		WithFollowedUsers(func(query *ent.BusinessFollowUserQuery) {
			query.WithBusiness()
			query.WithUser()
		}).
		WithFollowerBusinesses(func(query *ent.BusinessFollowBusinessQuery) {
			query.WithFollower()

		}).
		WithFollowedBusinesses(func(query *ent.BusinessFollowBusinessQuery) {
			query.WithFollowed()

		}).
		WithEvents(func(query *ent.EventQuery) {
		}).
		WithPlaces(func(query *ent.PlaceQuery) {
			query.WithMedias()
		}).
		WithUserBusinesses().
		WithPosts(func(query *ent.PostQuery) {
			query.WithUser()
			query.WithBusinessAccount()
			query.WithMedias()
			query.WithLikes(func(query *ent.LikeQuery) {
				query.WithUser()
			})
			query.WithComments(func(query *ent.CommentQuery) {
				query.WithUser()
			})
		}).
		WithWebsites(func(query *ent.WebsiteQuery) {
		}).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return businessAccount, nil
}

func (s *BusinessAccountServiceImpl) AddTeamMember(ctx context.Context, adminUser, userID, businessAccountID, role, permissions string) error {
	// Check if adminUser is indeed an admin of the business account
	adminUserBusiness, err := s.client.UserBusiness.
		Query().
		Where(userbusiness.HasUserWith(user.ID(adminUser))).
		Where(userbusiness.HasBusinessWith(business.ID(businessAccountID))).
		Only(ctx)

	if err != nil {
		// Add proper error handling here
		log.Println(err)
		return err
	}

	if adminUserBusiness.Role != "admin" {
		return appErrors.ErrPermissionDenied
	}

	// Check if the user is already part of the team
	existingRelation, err := s.client.UserBusiness.
		Query().
		Where(userbusiness.HasUserWith(user.ID(userID))).
		Where(userbusiness.HasBusinessWith(business.ID(businessAccountID))).
		Only(ctx)

	if err != nil && !ent.IsNotFound(err) {
		// Add proper error handling here
		log.Println(err)
		return err
	}

	if existingRelation != nil {
		// Handle situation when user is already part of the team. Return an error or ignore based on your requirements
		return appErrors.ErrUserAlreadyInTeam
	}

	log.Println("role", role)

	// Fetch user and business account
	userToAdd, err := s.client.User.Get(ctx, userID)
	if err != nil {
		return err
	}
	log.Println("userToAdd", userToAdd)

	businessAccount, err := s.client.Business.Query().Where(business.ID(businessAccountID)).Only(ctx)
	if err != nil {
		return err
	}

	// Create the relationship with role and permissions
	_, err = s.client.UserBusiness.Create().
		SetID(uuid.New().String()).
		SetUser(userToAdd).
		SetBusiness(businessAccount).
		SetRole(role).
		SetPermissions(permissions).
		Save(ctx)

	return err
}

func (s *BusinessAccountServiceImpl) ListTeamMembers(ctx context.Context, businessAccountID string) ([]*ent.UserBusiness, error) {
	relationships, err := s.client.Business.Query().
		Where(business.ID(businessAccountID)).
		QueryUserBusinesses().
		WithUser().
		All(ctx)
	return relationships, err
}

func (s *BusinessAccountServiceImpl) EditTeamMember(ctx context.Context, userID, businessAccountID, newRole, newPermissions string) error {
	userBusinessRel, err := s.client.UserBusiness.Query().
		Where(
			userbusiness.HasUserWith(user.ID(userID)),
			userbusiness.HasBusinessWith(business.ID(businessAccountID)),
		).Only(ctx)

	if err != nil {
		return err
	}

	_, err = s.client.UserBusiness.UpdateOneID(userBusinessRel.ID).
		SetRole(newRole).
		SetPermissions(newPermissions).
		Save(ctx)

	return err
}

func (s *BusinessAccountServiceImpl) RemoveTeamMember(ctx context.Context, userID, businessAccountID string) error {
	userBusinessRel, err := s.client.UserBusiness.Query().
		Where(
			userbusiness.HasUserWith(user.ID(userID)),
			userbusiness.HasBusinessWith(business.ID(businessAccountID)),
		).Only(ctx)

	if err != nil {
		return err
	}

	return s.client.UserBusiness.DeleteOne(userBusinessRel).Exec(ctx)
}

func (s *BusinessAccountServiceImpl) SearchTeamMembers(ctx context.Context, businessAccountID, searchText string) ([]*ent.UserBusiness, error) {
	relationships, err := s.client.Business.Query().
		Where(business.ID(businessAccountID)).
		QueryUserBusinesses().
		Where(
			userbusiness.HasUserWith(
				user.Or(
					user.NameContains(searchText),
					user.UsernameContains(searchText),
				),
			),
			userbusiness.Or(
				userbusiness.RoleContains(searchText),
				userbusiness.PermissionsContains(searchText),
			),
		).
		WithUser().
		All(ctx)

	return relationships, err
}

// RemoveUserFromBusinessAccount removes a User's association with a Business Account.
func (s *BusinessAccountServiceImpl) RemoveUserFromBusinessAccount(ctx context.Context, userID, businessAccountID string) error {
	ubr, err := s.client.UserBusiness.Query().Where(userbusiness.HasUserWith(user.ID(userID)), userbusiness.HasBusinessWith(business.ID(businessAccountID))).Only(ctx)
	if err != nil {
		return err
	}

	return s.client.UserBusiness.DeleteOne(ubr).Exec(ctx)
}

func (s *BusinessAccountServiceImpl) TransferBusinessAccountOwnership(ctx context.Context, currentOwnerID, newOwnerID, businessAccountID string) error {
	_, err := s.client.Business.Get(ctx, businessAccountID)
	if err != nil {
		return err
	}
	//if businessAccount.OwnerID != currentOwnerID {
	//	return errors.New("Current user is not the owner of the business account")
	//}
	//_, err = s.client.Business.UpdateOneID(businessAccountID).SetOwnerID(newOwnerID).Save(ctx)
	return err
}

// AssociateUserWithBusinessAccount associates a user with a Business Account.
func (s *BusinessAccountServiceImpl) AssociateUserWithBusinessAccount(ctx context.Context, adminUser, userID, businessAccountID, role string) error {
	// get the user
	user, err := s.client.User.
		Query().
		Where(user.ID(userID)).
		Only(ctx)

	if err != nil {
		return err
	}

	// get the business account
	businessAccount, err := s.client.Business.
		Query().
		Where(business.IDEQ(businessAccountID)).
		Only(ctx)

	if err != nil {
		return err
	}

	// create the relationship
	_, err = s.client.UserBusiness.
		Create().
		SetUser(user).
		SetBusiness(businessAccount).
		SetRole(role).
		Save(ctx)

	if err != nil {
		return err
	}

	return nil
}

// GetBusinessAccountsForUser returns all Business Accounts associated with a user.
func (s *BusinessAccountServiceImpl) GetBusinessAccountsForUser(ctx context.Context, userID string) ([]*ent.UserBusiness, error) {
	relationships, err := s.client.User.Query().
		Where(user.IDEQ(userID)).
		QueryUserBusinesses().
		All(ctx)

	if err != nil {
		return nil, err
	}

	return relationships, nil
}

// GetUsersForBusinessAccount returns all Users associated with a Business Account.
func (s *BusinessAccountServiceImpl) GetUsersForBusinessAccount(ctx context.Context, businessAccountID string) ([]*ent.UserBusiness, error) {
	relationships, err := s.client.Business.Query().
		Where(business.IDEQ(businessAccountID)).QueryUserBusinesses().
		All(ctx)

	if err != nil {
		return nil, err
	}

	return relationships, nil
}

// GetUserBusinessAccounts retrieves all business accounts associated with a specific user.
func (s *BusinessAccountServiceImpl) GetUserBusinessAccounts(ctx context.Context, userId string) ([]*ent.Business, error) {
	// Find the user with the provided ID.
	businesses, err := s.client.User.Query().
		Where(user.IDEQ(userId)).
		QueryUserBusinesses().
		QueryBusiness().
		//WithFollowedBusinesses().
		//WithFollowerUsers().
		//WithFollowedUsers().
		//WithFollowerBusinesses().
		//WithPlaces(func(query *ent.PlaceQuery) {
		//	query.WithMedias()
		//	query.WithLikedByUsers()
		//	query.WithUsers()
		//}).
		//WithPlaceInventories().
		All(ctx)

	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	return businesses, nil
}

func (s *BusinessAccountServiceImpl) UpdateBusinessAccount(ctx context.Context, businessID string, businessData map[string]interface{}) (*ent.Business, error) {
	// Check if business exists
	business, err := s.client.Business.Get(ctx, businessID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("business does not exist")
		}
		return nil, fmt.Errorf("failed checking business existence: %w", err)
	}

	// Get an updater for the business
	upd := s.client.Business.UpdateOne(business)

	// Update fields
	if v, ok := businessData["name"]; ok {
		upd.SetName(v.(string))
	}
	if v, ok := businessData["cover_image"]; ok {
		upd.SetCoverImage(v.(string))
	}
	if v, ok := businessData["description"]; ok {
		upd.SetDescription(v.(string))
	}
	if v, ok := businessData["location"]; ok {
		upd.SetLocation(v.(string))
	}
	if v, ok := businessData["website"]; ok {
		upd.SetWebsite(v.(string))
	}
	if v, ok := businessData["picture"]; ok {
		upd.SetPicture(v.(string))
	}
	if v, ok := businessData["phone"]; ok {
		upd.SetPhone(v.(string))
	}
	if v, ok := businessData["email"]; ok {
		upd.SetEmail(v.(string))
	}

	// Update business settings
	if v, ok := businessData["business_settings"]; ok {
		// Merge existing and new settings
		newSettings := v.(map[string]interface{})
		for k, value := range business.BusinessSettings {
			if _, exists := newSettings[k]; !exists {
				newSettings[k] = value
			}
		}
		upd.SetBusinessSettings(newSettings)
	}

	// Save the updates
	business, err = upd.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed updating business: %w", err)
	}

	// update elasticsearch
	if err := s.searchService.CreateOrUpdateBusiness(ctx, business); err != nil {
		return nil, fmt.Errorf("failed updating business in elasticsearch: %w", err)
	}

	return business, nil
}

func (s *BusinessAccountServiceImpl) DeleteBusinessAccount(ctx context.Context, businessAccountID string) error {
	if businessAccountID == "" {
		return errors.New("businessAccountID cannot be nil")
	}

	err := s.client.Business.DeleteOneID(businessAccountID).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *BusinessAccountServiceImpl) ListBusinessAccounts(ctx context.Context, page, pageSize int, sortBy string, filters ...predicate.Business) ([]*ent.Business, error) {
	businesses, err := s.client.Business.Query().
		Where(filters...).
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Order(ent.Asc(sortBy)).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return businesses, nil
}
