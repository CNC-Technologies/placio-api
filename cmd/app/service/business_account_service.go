package service

import (
	"context"
	"errors"
	"fmt"
	"placio-app/Dto"
	"placio-app/ent"
	"placio-app/ent/business"
	"placio-app/ent/predicate"
	"placio-app/ent/user"
	"placio-app/ent/userbusiness"
)

type BusinessAccountService interface {
	CreateBusinessAccount(ctx context.Context, businessData *Dto.BusinessDto) (*ent.Business, error)
	GetBusinessAccount(ctx context.Context, businessAccountID string) (*ent.Business, error)
	GetUserBusinessAccounts(ctx context.Context, userID string) ([]*ent.Business, error)
	AssociateUserWithBusinessAccount(ctx context.Context, userID, businessAccountID, role string) error
	GetBusinessAccountsForUser(ctx context.Context, userID string) ([]*ent.UserBusiness, error)
	ListBusinessAccounts(ctx context.Context, page, pageSize int, sortBy string, filters ...predicate.Business) ([]*ent.Business, error)
	UpdateBusinessAccount(ctx context.Context, businessAccount *ent.Business) (*ent.Business, error)
	RemoveUserFromBusinessAccount(ctx context.Context, userID, businessAccountID string) error
	TransferBusinessAccountOwnership(ctx context.Context, currentOwnerID, newOwnerID, businessAccountID string) error
	//CanPerformAction(ctx context.Context, userID, businessAccountID, action string) (bool, error)
}

type BusinessAccountServiceImpl struct {
	store  *ent.Business
	client *ent.Client
}

func NewBusinessAccountService(client *ent.Client) BusinessAccountService {
	return &BusinessAccountServiceImpl{client: client, store: &ent.Business{}}
}

// CreateBusinessAccount creates a new Business Account and associates it with a user.
func (s *BusinessAccountServiceImpl) CreateBusinessAccount(ctx context.Context, businessData *Dto.BusinessDto) (*ent.Business, error) {
	// Validate inputs
	if businessData.UserID == "" {
		return nil, errors.New("user ID cannot be empty")
	}
	if businessData.Name == "" {
		return nil, errors.New("business account name cannot be empty")
	}
	if businessData.Role != "owner" && businessData.Role != "admin" && businessData.Role != "member" {
		return nil, errors.New("invalid role")
	}

	// Create a new transaction
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("error starting transaction: %w", err)
	}

	// Create business account
	businessAccount, err := tx.Business.
		Create().
		SetName(businessData.Name).
		Save(ctx)

	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error creating business account: %w", err)
	}

	// Get the user
	user, err := tx.User.
		Query().
		Where(user.Auth0ID(businessData.UserID)).
		Only(ctx)

	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error fetching user: %w", err)
	}

	// Create user-business relationship
	_, err = tx.UserBusiness.
		Create().
		SetUser(user).
		SetBusiness(businessAccount).
		SetRole(businessData.Role).
		Save(ctx)

	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error creating user-business relationship: %w", err)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("error committing transaction: %w", err)
	}

	// Now we need to fetch the created business account with its relationships
	businessAccount, err = s.client.Business.
		Query().
		Where(business.ID(businessAccount.ID)).
		WithUserBusinesses().
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("error fetching created business account: %w", err)
	}

	return businessAccount, nil
}

func (s *UserServiceImpl) CanPerformAction(ctx context.Context, userID, businessAccountID, action string) (bool, error) {
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

func (bas *BusinessAccountServiceImpl) GetBusinessAccount(ctx context.Context, businessAccountID string) (*ent.Business, error) {
	if businessAccountID == "" {
		return nil, errors.New("businessAccountID cannot be nil")
	}

	businessAccount, err := bas.client.Business.Get(ctx, businessAccountID)
	if err != nil {
		return nil, err
	}

	return businessAccount, nil
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
func (s *BusinessAccountServiceImpl) AssociateUserWithBusinessAccount(ctx context.Context, userID, businessAccountID, role string) error {
	// get the user
	user, err := s.client.User.
		Query().
		Where(user.Auth0ID(userID)).
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
func (s *BusinessAccountServiceImpl) GetUserBusinessAccounts(ctx context.Context, userID string) ([]*ent.Business, error) {
	// Find the user with the provided ID.
	user, err := s.client.User.Get(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	// Query the associated businesses.
	businesses, err := user.QueryUserBusinesses().QueryBusiness().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user's businesses: %w", err)
	}

	return businesses, nil
}

func (bas *BusinessAccountServiceImpl) UpdateBusinessAccount(ctx context.Context, businessAccount *ent.Business) (*ent.Business, error) {
	if businessAccount == nil {
		return nil, errors.New("businessAccount cannot be nil")
	}

	businessAccount, err := bas.client.Business.UpdateOne(businessAccount).Save(ctx)
	if err != nil {
		return nil, err
	}

	return businessAccount, nil
}

func (bas *BusinessAccountServiceImpl) DeleteBusinessAccount(ctx context.Context, businessAccountID string) error {
	if businessAccountID == "" {
		return errors.New("businessAccountID cannot be nil")
	}

	err := bas.client.Business.DeleteOneID(businessAccountID).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (bas *BusinessAccountServiceImpl) ListBusinessAccounts(ctx context.Context, page, pageSize int, sortBy string, filters ...predicate.Business) ([]*ent.Business, error) {
	businesses, err := bas.client.Business.Query().
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
