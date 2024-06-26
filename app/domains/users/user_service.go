package users

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/getsentry/sentry-go"
	"io"
	"log"
	"net"
	"net/http"
	"placio-app/domains/search"
	"placio-app/ent"
	"placio-app/ent/business"
	"placio-app/ent/like"
	"placio-app/ent/user"
	"placio-app/ent/userfollowbusiness"
	"placio-app/ent/userfollowuser"
	"placio-app/models"
	"placio-app/utility"
	"placio-pkg/hash"
	"reflect"
	"strings"
	"time"

	"github.com/auth0/go-auth0/management"
)

const (
	auth0TokenCacheKey      = "auth0_mgmt_token"
	auth0TokenEncryptionKey = "eyJhbGciOiJSUzIC"
)

type UserService interface {
	GetUser(ctx context.Context, authOID string) (*ent.User, error)
	GetUserWithoutAuth0Data(ctx context.Context, auth0ID string) (*ent.User, error)
	GetAuth0UserData(userID string) (*management.User, error)
	UpdateAuth0UserMetadata(userID string, userMetaData *models.Metadata) (*management.User, error)
	UpdateAuth0UserInformation(userID string, userData *models.Auth0UserData) (*management.User, error)
	GetUserByUserId(ctx context.Context, userId string) (*ent.User, error)
	UpdateUser(ctx context.Context, userID string, userData map[string]interface{}) (*ent.User, error)
	CheckFollowing(ctx context.Context, userId string, followedID, businessId string) (bool, error)
	// GetAuth0ManagementToken GetAuth0UserMetaData(userID string, IdToken string) (models.Metadata, error)
	//GetAuth0AppMetaData(userID string, IdToken string) (models.AppMetadata, error)
	//GetAuth0UserRoles(userID string, IdToken string) ([]string, error)
	//GetAuth0UserPermissions(userID string, IdToken string) ([]string, error)
	//GetAuth0UserGroups(userID string, IdToken string) ([]string, error)
	//GetAuth0UserRolesPermissions(userID string, IdToken string) ([]string, error)
	//AuthorizeUser(userID string, IdToken string, roles []string, permissions []string, groups []string) error
	//DeAuthorizeUser(userID string, IdToken string, roles []string, permissions []string, groups []string) error
	GetAuth0ManagementToken(ctx context.Context) (string, error)
	UpdateAuth0AppMetadata(userID string, appData *models.AppMetadata) (*management.User, error)
	GetPostsByUser(ctx context.Context, userID string) ([]*ent.Post, error)
	FollowUser(ctx context.Context, followerID string, followedID string) error
	FollowBusiness(ctx context.Context, followerID string, businessID string) error
	UnfollowUser(ctx context.Context, followerID string, followedID string) error
	UnfollowBusiness(ctx context.Context, followerID string, businessID string) error
	GetFollowedContents(ctx context.Context, userID string) ([]*ent.Post, error)
	GetFollowers(ctx context.Context, userID, nextPageToken string, limit int) ([]*ent.User, error)
	GetLikes(ctx context.Context, userID, nextPageToken string, limit int) ([]*ent.Like, error)
}

type UserServiceImpl struct {
	client        *ent.Client
	cache         *utility.RedisClient
	searchService search.SearchService
}

type auth0TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func NewUserService(client *ent.Client, cache *utility.RedisClient, searchService search.SearchService) *UserServiceImpl {
	return &UserServiceImpl{client: client, cache: cache, searchService: searchService}
}

func (s *UserServiceImpl) CheckFollowing(ctx context.Context, userId string, followedID, businessId string) (bool, error) {
	if businessId != "" {
		// Check if the user is following the business
		followedBusiness, err := s.client.UserFollowBusiness.
			Query().
			Where(userfollowbusiness.HasUserWith(user.ID(userId)), userfollowbusiness.HasBusinessWith(business.ID(businessId))).
			Only(ctx)

		if err != nil {
			if ent.IsNotFound(err) {
				return false, nil
			}
			return false, err
		}

		if followedBusiness != nil {
			return true, nil
		}

	} else {
		// Check if the user is following another user
		followedUser, err := s.client.UserFollowUser.
			Query().
			Where(userfollowuser.HasFollowerWith(user.ID(userId)), userfollowuser.HasFollowedWith(user.ID(followedID))).
			Only(ctx)

		if err != nil {
			if ent.IsNotFound(err) {
				return false, nil
			}
			return false, err
		}

		if followedUser != nil {
			return true, nil
		}
	}

	return false, nil
}

func (s *UserServiceImpl) FollowUser(ctx context.Context, followerID string, followedID string) error {
	_, err := s.client.UserFollowUser.
		Create().
		SetFollowerID(followerID).
		SetFollowedID(followedID).
		Save(ctx)
	return err
}

func (s *UserServiceImpl) FollowBusiness(ctx context.Context, followerID string, businessID string) error {
	_, err := s.client.UserFollowBusiness.
		Create().
		SetUserID(followerID).
		SetBusinessID(businessID).
		Save(ctx)
	return err
}

func (s *UserServiceImpl) UnfollowUser(ctx context.Context, followerID string, followedID string) error {
	uf, err := s.client.UserFollowUser.
		Query().
		Where(userfollowuser.HasFollowerWith(user.ID(followerID)), userfollowuser.HasFollowedWith(user.ID(followedID))).
		Only(ctx)
	if err != nil {
		return err
	}
	return s.client.UserFollowUser.DeleteOne(uf).Exec(ctx)
}

func (s *UserServiceImpl) UnfollowBusiness(ctx context.Context, followerID string, businessID string) error {
	ub, err := s.client.UserFollowBusiness.
		Query().
		Where(userfollowbusiness.HasUserWith(user.ID(followerID)), userfollowbusiness.HasBusinessWith(business.ID(businessID))).
		Only(ctx)
	if err != nil {
		return err
	}
	return s.client.UserFollowBusiness.DeleteOne(ub).Exec(ctx)
}

func (s *UserServiceImpl) GetFollowedContents(ctx context.Context, userID string) ([]*ent.Post, error) {
	// First, fetch the followed users.
	followedUsers, err := s.client.User.
		Query().
		Where(user.IDEQ(userID)).
		QueryFollowedUsers().
		All(ctx)
	if err != nil {
		return nil, err
	}

	// Then, for each followed user, fetch their posts.
	var allPosts []*ent.Post
	for _, followedUser := range followedUsers {
		posts, err := s.client.User.
			Query().
			Where(user.IDEQ(followedUser.ID)).
			QueryPosts().
			WithMedias().
			WithUser().
			WithComments(func(q *ent.CommentQuery) {
				q.WithUser()
			}).
			All(ctx)
		if err != nil {
			return nil, err
		}
		allPosts = append(allPosts, posts...)
	}

	return allPosts, nil
}

func (s *UserServiceImpl) GetUser(ctx context.Context, auth0ID string) (*ent.User, error) {
	var userId string

	if auth0ID == "" {
		return nil, errors.New("auth0ID is empty")
	}

	if strings.Contains(auth0ID, "|") {
		userId = strings.Split(auth0ID, "|")[1]
	} else {
		userId = auth0ID
	}

	log.Println("GetUser", auth0ID, "userId", userId)

	u, err := s.client.User.
		Query().
		Where(user.IDEQ(userId)).
		WithLikes(func(query *ent.LikeQuery) {
			query.WithPost(func(query *ent.PostQuery) {
				query.WithUser()
				query.WithMedias()
				query.WithBusinessAccount()
				query.WithComments(func(query *ent.CommentQuery) {
					query.WithUser()
				})
				query.WithLikes(func(query *ent.LikeQuery) {
					query.WithUser()
					query.WithPost()
				})
			})
		}).
		WithFollowedUsers(func(query *ent.UserFollowUserQuery) {
			query.WithFollowed()

		}).
		WithFollowerUsers(func(query *ent.UserFollowUserQuery) {
			query.WithFollower()
		}).
		WithReviews(func(query *ent.ReviewQuery) {
			query.WithBusiness()
		}).
		WithFollowedBusinesses(func(query *ent.UserFollowBusinessQuery) {
			query.WithBusiness()
		}).
		WithFollowedPlaces(func(query *ent.UserFollowPlaceQuery) {
			query.WithPlace()
		}).
		WithPosts(func(query *ent.PostQuery) {
			query.WithMedias()
			query.WithUser()
			query.WithBusinessAccount()
			query.WithComments(func(query *ent.CommentQuery) {
				query.WithUser()
				query.WithReplies(func(query *ent.CommentQuery) {
					query.WithUser()
				})
			})
		}).
		WithComments(func(query *ent.CommentQuery) {
			query.WithUser()
		}).
		WithLikedPlaces(func(query *ent.UserLikePlaceQuery) {
			query.WithPlace()
		}).
		First(ctx)

	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, err
		}

		log.Println("GetUser", auth0ID, "user not found, creating new user")

		auth0Data, err := s.getAuth0UserDataWithRetry(auth0ID, 3, 1*time.Second)
		if err != nil {
			return nil, err
		}

		// userId should be the same as auth0ID but without the auth0| prefix
		newUser, err := s.client.User.
			Create().
			SetID(userId).
			SetAuth0ID(auth0ID).
			SetName(strings.ToLower(*auth0Data.Name)).
			SetUsername(strings.ToLower(utility.GenerateRandomUsername())).
			SetPicture(*auth0Data.Picture).
			Save(ctx)

		if err != nil {
			log.Println("GetUser", auth0ID, "error creating new user", err)
			return nil, err
		}

		go func() {
			// add user to search index
			err = s.searchService.CreateOrUpdateUser(ctx, newUser)
			if err != nil {
				log.Println("GetUser", auth0ID, "error adding user to search index", err)
			}
		}()

		//newUser.AppSettings = *auth0Data.AppMetadata
		//newUser.UserSettings = *auth0Data.UserMetadata
		newUser.Auth0Data = auth0Data
		return newUser, nil
	}

	//auth0Data, err := s.getAuth0UserDataWithRetry(auth0ID, 3, 1*time.Second)
	//if err != nil {
	//	return nil, err
	//}

	//u.AppSettings = *auth0Data.AppMetadata
	//u.UserSettings = *auth0Data.UserMetadata
	//u.Auth0Data = auth0Data

	return u, nil
}

func (s *UserServiceImpl) GetUserWithoutAuth0Data(ctx context.Context, auth0ID string) (*ent.User, error) {

	u, err := s.client.User.
		Query().
		Where(user.Auth0IDEQ(auth0ID)).
		First(ctx)

	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, err
		}

		newUser, err := s.client.User.
			Create().
			SetID(utility.GenerateID()).
			SetAuth0ID(auth0ID).
			Save(ctx)

		if err != nil {
			log.Println("GetUser", auth0ID, "error creating new user", err)
			return nil, err
		}
		return newUser, nil
	}

	return u, nil
}

func (s *UserServiceImpl) GetUserByUserId(ctx context.Context, userId string) (*ent.User, error) {

	u, err := s.client.User.
		Query().
		Where(user.IDEQ(userId)).
		WithUserBusinesses(func(query *ent.UserBusinessQuery) {
			query.WithUser()
			query.WithBusiness()
		}).
		WithPlaces(func(query *ent.PlaceQuery) {
			query.WithBusiness()
		}).
		WithLikes(func(query *ent.LikeQuery) {
			query.WithPost(func(query *ent.PostQuery) {
				query.WithUser()
				query.WithMedias()
				query.WithBusinessAccount()
				query.WithComments(func(query *ent.CommentQuery) {
					query.WithUser()
				})
				query.WithLikes(func(query *ent.LikeQuery) {
					query.WithUser()
					query.WithPost()
				})
			})
		}).
		WithLikedPlaces(func(query *ent.UserLikePlaceQuery) {
			query.WithPlace(func(query *ent.PlaceQuery) {
				query.WithBusiness()
			})
		}).
		WithPosts(func(query *ent.PostQuery) {
			query.WithUser()
			query.WithMedias()
		}).
		WithFollowedUsers(func(query *ent.UserFollowUserQuery) {
			query.WithFollowed()

		}).
		WithFollowedBusinesses(func(query *ent.UserFollowBusinessQuery) {
			query.WithBusiness()
			query.WithUser()
		}).
		WithFollowerUsers(func(query *ent.UserFollowUserQuery) {
			query.WithFollower()
		}).
		WithPlaces(func(query *ent.PlaceQuery) {
			query.WithBusiness()
			query.WithLikedByUsers()
			query.WithUsers()
			query.WithMedias()
			query.WithFollowerUsers()
		}).
		First(ctx)

	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, err
		}
		return nil, err
	}

	return u, nil
}

func (s *UserServiceImpl) getAuth0UserDataWithRetry(auth0ID string, maxRetries int, retryDelay time.Duration) (*management.User, error) {
	// Replace "YourServiceType" with the actual type of your service

	log.Println("GetUser", auth0ID, "attempting to get Auth0 user data")
	var auth0Data *management.User
	retryCount := 0

	for {
		auth0DataAttempt, err := s.GetAuth0UserData(auth0ID)
		if err != nil {
			if specificErr, ok := err.(net.Error); ok {
				if specificErr.Temporary() {
					// handle temporary error
				} else if specificErr.Timeout() {
					// handle timeout
				}
				// handle specific error
				log.Println("Specific error occurred when getting Auth0 user data", specificErr)
				return nil, specificErr
			} else if retryCount < maxRetries {
				// wait for a while before retrying
				time.Sleep(retryDelay)
				retryCount++
				continue
			} else {
				// max retries exceeded, return the error
				log.Println("Error getting Auth0 user data", err)
				return nil, err
			}
		}

		auth0Data = auth0DataAttempt
		break
	}

	return auth0Data, nil
}

func (s *UserServiceImpl) mergeAuth0DataIntoUser(user models.User, auth0Data *management.User) models.User {
	user.Auth0Data = auth0Data
	return user
}

func (s *UserServiceImpl) UpdateAuth0UserInformation(userID string, userData *models.Auth0UserData) (*management.User, error) {
	//mergedData, err := s.prepareUserData(userID, userData)
	mergedData, err := utility.StructToMap(&userData)
	if err != nil {
		return nil, err
	}

	// Iterate over the map and delete keys with zero values
	for key, value := range mergedData {
		v := reflect.ValueOf(value)
		if (v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface) && v.IsNil() {
			delete(mergedData, key)
		} else if v.Kind() != reflect.Ptr && v.IsZero() {
			delete(mergedData, key)
		}
	}

	return s.updateAuth0Data(userID, mergedData, "user information")
}

func (s *UserServiceImpl) UpdateAuth0UserMetadata(userID string, userMetaData *models.Metadata) (*management.User, error) {
	newUserData, err := utility.StructToMap(&userMetaData)
	// TODO: Merge the new data with the existing data need to figure out how to do this
	//mergedData, err := s.prepareUserData(userID, userMetaData, "user_metadata")
	if err != nil {
		return nil, err
	}
	return s.updateAuth0Data(userID, newUserData, "user_metadata")
}

func (s *UserServiceImpl) UpdateAuth0AppMetadata(userID string, appData *models.AppMetadata) (*management.User, error) {
	log.Println("Updating app metadata")
	newAppData, err := utility.StructToMap(&appData)
	// TODO: This is not working
	//mergedData, err := s.prepareUserData(userID, appData, "app_metadata")
	if err != nil {
		return nil, err
	}
	return s.updateAuth0Data(userID, newAppData, "app_metadata")
}

func (s *UserServiceImpl) prepareUserData(userID string, data interface{}, dataType string) (map[string]interface{}, error) {
	currUserData, err := s.GetAuth0UserData(userID)
	if err != nil {
		log.Println("Error getting current user data", err)
		return nil, err
	}

	currUserDataMap, err := utility.StructToMap(&currUserData)
	if err != nil {
		return nil, err
	}

	newDataMap, err := utility.StructToMap(data)
	if err != nil {
		return nil, err
	}

	if dataType != "user information" {
		// Get the current data for this type
		currTypeData, ok := currUserDataMap[dataType].(map[string]interface{})
		if !ok {
			// If it doesn't exist or is not a map, initialize it
			currTypeData = make(map[string]interface{})
		}

		// Merge the new data into the current data for this type
		newTypeData := utility.MergeMaps(currTypeData, newDataMap)

		// Put the merged data back into the current user data
		currUserDataMap[dataType] = newTypeData
	} else {
		// For "user information", merge the new data into the whole current user data
		currUserDataMap = utility.MergeMaps(currUserDataMap, newDataMap)
	}

	return currUserDataMap, nil
}

func (s *UserServiceImpl) updateAuth0Data(userID string, mergedUserData map[string]interface{}, dataType string) (*management.User, error) {
	log.Println("Updating auth0 data", dataType)
	client := &http.Client{}

	if dataType != "user information" {
		mergedUserData = map[string]interface{}{
			dataType: mergedUserData,
		}
	}

	jsonPayload, err := json.Marshal(mergedUserData)
	if err != nil {
		return nil, err
	}
	log.Println("JSON payload", string(jsonPayload))

	req, err := http.NewRequest("PATCH", fmt.Sprintf("https://auth.placio.io/api/v2/users/%s", userID), bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Println("Error creating request", err)
		return nil, err
	}

	managementToken, err := s.GetAuth0ManagementToken(context.Background())
	if err != nil {
		log.Println("Error getting management token", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", managementToken))

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request", err)
		return nil, err
	}
	defer resp.Body.Close()

	log.Println("Response status code", resp.StatusCode)
	log.Println("Response status", resp.Status)

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	log.Println("Successfully updated", dataType, "for user", userID)
	// unmarshal the response body into management.User
	var user management.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserServiceImpl) UpdateUser(ctx context.Context, userID string, userData map[string]interface{}) (*ent.User, error) {
	// Check if user exists
	userToUpdate, err := s.client.User.Get(ctx, userID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("user does not exist")
		}
		return nil, fmt.Errorf("failed checking user existence: %w", err)
	}

	// Get a updater for the user
	upd := s.client.User.UpdateOne(userToUpdate)

	// Update fields
	if v, ok := userData["name"]; ok {
		upd.SetName(v.(string))
	}
	if v, ok := userData["cover_image"]; ok {
		upd.SetCoverImage(v.(string))
	}
	if v, ok := userData["bio"]; ok {
		upd.SetBio(v.(string))
	}
	if v, ok := userData["location"]; ok {
		upd.SetLocation(v.(string))
	}
	if v, ok := userData["website"]; ok {
		upd.SetWebsite(v.(string))
	}
	if v, ok := userData["username"]; ok {
		upd.SetUsername(v.(string))
	}
	if v, ok := userData["picture"]; ok {
		upd.SetPicture(v.(string))
	}

	// Update app settings
	if v, ok := userData["app_settings"]; ok {
		// Merge existing and new settings
		newSettings := v.(map[string]interface{})
		for k, value := range userToUpdate.AppSettings {
			if _, exists := newSettings[k]; !exists {
				newSettings[k] = value
			}
		}
		upd.SetAppSettings(newSettings)
	}

	// Update user settings
	if v, ok := userData["user_settings"]; ok {
		// Merge existing and new settings
		newSettings := v.(map[string]interface{})
		for k, value := range userToUpdate.UserSettings {
			if _, exists := newSettings[k]; !exists {
				newSettings[k] = value
			}
		}
		upd.SetUserSettings(newSettings)
	}

	// Save the updates
	userToUpdate, err = upd.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed updating user: %w", err)
	}

	// update elasticsearch
	if err := s.searchService.CreateOrUpdateUser(ctx, userToUpdate); err != nil {
		return nil, fmt.Errorf("failed updating user in elasticsearch: %w", err)
	}

	go func() {
		cacheKey := "user:" + userToUpdate.ID
		userToCache, _ := s.client.User.Query().
			Where(user.ID(userToUpdate.ID)).
			WithFollowedUsers().
			WithFollowerUsers().WithReviews().First(ctx)
		if err := s.cache.SetCache(ctx, cacheKey, userToCache); err != nil {
			sentry.CaptureException(err)
			return
		}
	}()

	return userToUpdate, nil
}

// GetAuth0UserData retrieves the current user data from Auth0.
func (s *UserServiceImpl) GetAuth0UserData(userID string) (*management.User, error) {
	log.Println("Getting Auth0 user data", userID)
	// Create an HTTP client
	client := &http.Client{}

	// Create the request
	req, err := http.NewRequest("GET", fmt.Sprintf("https://auth.placio.io/api/v2/users/%s", userID), nil)
	if err != nil {
		log.Println("Error creating request", err)
		return &management.User{}, err
	}

	//Get the token
	managementToken, err := s.GetAuth0ManagementToken(context.Background())
	if err != nil {
		log.Println("Error getting management token", err)
		return &management.User{}, err
	}
	// Set the headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", managementToken))

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request", err)
		return &management.User{}, err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return &management.User{}, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	// Parse the response body
	var userData management.User
	err = json.NewDecoder(resp.Body).Decode(&userData)
	if err != nil {
		return &management.User{}, err
	}

	return &userData, nil
}

func (s *UserServiceImpl) GetAuth0ManagementToken(ctx context.Context) (string, error) {

	// Check if token is in cache
	encryptedTokenBytes, err := s.cache.GetCache(ctx, auth0TokenCacheKey)
	if err != nil {
		log.Println("Error retrieving token from cache:", err)
	} else {
		// If token is in cache, decrypt and return
		encryptedToken := string(encryptedTokenBytes)
		encryptedToken = strings.Trim(encryptedToken, "\"") // Remove quotes from the string
		//log.Println("Retrieved encrypted token from cache:", encryptedToken) // Added log line
		token, err := hash.DecryptString(encryptedToken, auth0TokenEncryptionKey)
		//log.Println("Decrypted token:", token) // Added log line
		if err != nil {
			log.Println("Error decrypting token:", err)
		} else {
			return token, nil
		}
	}

	log.Println("Token not in cache, retrieving new token")

	// If token is not in cache or there was an error, retrieve a new one
	token, err := s.retrieveAuth0Token(ctx) // Replace with actual function to retrieve token
	if err != nil {
		return "", err
	}

	// Encrypt and cache the new token
	encryptedToken, err := hash.EncryptString(token, auth0TokenEncryptionKey)
	if err != nil {
		return "", err
	}

	log.Println("Caching encrypted token:", encryptedToken) // Added log line
	err = s.cache.SetCache(ctx, auth0TokenCacheKey, encryptedToken)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserServiceImpl) retrieveAuth0Token(ctx context.Context) (string, error) {

	payload := strings.NewReader(`{
		"client_id": "wORDxmfRFTkBvoSVU06Af4HFQAo25gUI",
		"client_secret": "_tbGyuaBZ7j5zp659MU5AYqkZessCVeNs2bv8Yl1Hp6XUj_hUdQAW9a5zw8hIA3F",
		"audience": "https://dev-qlb0lv3d.us.auth0.com/api/v2/",
		"grant_type": "client_credentials"
	}`)

	log.Println("payload", payload)
	req, _ := http.NewRequest("POST", "https://dev-qlb0lv3d.us.auth0.com/oauth/token", payload)
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error retrieving token from Auth0:", err)
		return "", err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", errors.New("auth0 request not successful")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return "", err
	}

	var tokenResponse auth0TokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		log.Println("Error unmarshalling response body:", err)
		return "", err
	}

	return tokenResponse.AccessToken, nil
}

func (s *UserServiceImpl) GetPostsByUser(ctx context.Context, userID string) ([]*ent.Post, error) {
	// Check if user exists
	user, err := s.client.User.Get(ctx, userID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("user does not exist")
		}
		return nil, fmt.Errorf("failed checking user existence: %w", err)
	}

	// Query posts by the user
	posts, err := user.QueryPosts().
		WithUser().
		WithBusinessAccount().
		WithMedias().
		WithLikes().
		WithComments(func(query *ent.CommentQuery) {
			query.WithUser()

		}).
		All(ctx)
	if err != nil {
		// Handle possible db errors
		return nil, fmt.Errorf("failed querying posts: %w", err)
	}

	// Return fetched posts
	return posts, nil
}

func (s *UserServiceImpl) GetFollowers(ctx context.Context, userID, nextPageToken string, limit int) ([]*ent.User, error) {
	// Check if user exists
	userData, err := s.client.User.Get(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed checking user existence: %w", err)
	}

	// Query followed users
	q := userData.QueryFollowedUsers().
		QueryFollower().
		Limit(limit) // Limit the number of users fetched

	// If a nextPageToken (UUID) is provided, we cmd the query after this user
	if nextPageToken != "" {
		q = q.Where(user.IDGT(nextPageToken)) // users after the given UUID
	}

	// Execute the query
	followedUsers, err := q.All(ctx)
	if err != nil {
		// Handle possible db errors
		return nil, fmt.Errorf("failed querying followed users: %w", err)
	}

	// Return fetched followed users
	return followedUsers, nil
}

func (s *UserServiceImpl) GetLikes(ctx context.Context, userID, nextPageToken string, limit int) ([]*ent.Like, error) {
	// Check if user exists
	userData, err := s.client.User.Get(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed checking user existence: %w", err)
	}

	// Query followed users
	q := userData.QueryLikes().
		Limit(limit) // Limit the number of users fetched

	// If a nextPageToken (UUID) is provided, we cmd the query after this user
	if nextPageToken != "" {
		q = q.Where(like.IDGT(nextPageToken)) // users after the given UUID
	}

	// Execute the query
	followedUsers, err := q.All(ctx)
	if err != nil {
		// Handle possible db errors
		return nil, fmt.Errorf("failed querying followed users: %w", err)
	}

	// Return fetched followed users
	return followedUsers, nil
}

//func (s *UserServiceImpl) RejectInvitation(invitationID uint) error {
//	// Implementation goes here
//	return nil
//}
//
//func (s *UserServiceImpl) AcceptInvitation(invitationID uint) error {
//	// Implementation goes here
//	return nil
//}
//
//func (s *UserServiceImpl) InviteUserToBusinessAccount(email string, businessAccountID uint, role string) (*models.Invitation, error) {
//	// Implementation goes here
//	return nil, nil
//}
