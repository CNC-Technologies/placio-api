package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "placio-app/Dto"
	_ "placio-app/ent"
	"placio-app/service"
	"placio-app/utility"
)

type FollowController struct {
	service service.IFollowService
}

func NewFollowController(service service.IFollowService) *FollowController {
	return &FollowController{service: service}
}

func (fc *FollowController) RegisterRoutes(router *gin.RouterGroup) {
	followRouter := router.Group("/follow")
	{
		followRouter.POST("business/:businessID", utility.Use(fc.followUserToBusiness))
		followRouter.DELETE("business/:businessID", utility.Use(fc.unfollowUserToBusiness))
		followRouter.GET("business", utility.Use(fc.getFollowedBusinessesByUser))

		followRouter.POST("user/:userID", utility.Use(fc.followUserToUser))
		followRouter.DELETE("user/:userID", utility.Use(fc.unfollowUserToUser))
		followRouter.GET("user", utility.Use(fc.getFollowedUsersByUser))

		followRouter.POST("place/:placeID", utility.Use(fc.followUserToPlace))
		followRouter.DELETE("place/:placeID", utility.Use(fc.unfollowUserToPlace))
		followRouter.GET("place", utility.Use(fc.getFollowedPlacesByUser))

		followRouter.POST("event/:eventID", utility.Use(fc.followUserToEvent))
		followRouter.DELETE("event/:eventID", utility.Use(fc.unfollowUserToEvent))
		followRouter.GET("event", utility.Use(fc.getFollowedEventsByUser))
	}
}

// @Summary Follow a business
// @ID user-follow-business
// @Tags Follow
// @Produce json
// @Param businessID path string true "Business ID"
// @Security Bearer
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Follow a specific Business by the user
// @Success 201 {object} string "User successfully followed the business"
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /follow/business/{businessID} [post]
func (fc *FollowController) followUserToBusiness(c *gin.Context) error {
	businessID := c.Param("businessID")
	userID := c.GetString("user")

	err := fc.service.FollowUserToBusiness(c, userID, businessID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utility.ProcessResponse(nil, "failed", err.Error()))
		return nil
	}

	c.JSON(http.StatusCreated, utility.ProcessResponse(nil, "success", "User successfully followed the business"))
	return nil
}

// @Summary Unfollow a business
// @ID user-unfollow-business
// @Tags Follow
// @Produce json
// @Param businessID path string true "Business ID"
// @Security Bearer
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Unfollow a specific Business by the user
// @Success 200 {object} string "User successfully unfollowed the business"
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /unfollow/business/{businessID} [delete]
func (fc *FollowController) unfollowUserToBusiness(c *gin.Context) error {
	businessID := c.Param("businessID")
	userID := c.GetString("user")

	err := fc.service.UnfollowUserToBusiness(c, userID, businessID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utility.ProcessResponse(nil, "failed", err.Error()))
		return nil
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(nil, "success", "User successfully unfollowed the business"))
	return nil
}

// @Summary Get followed businesses by a user
// @ID get-followed-businesses
// @Tags Follow
// @Produce json
// @Security Bearer
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Get all Businesses followed by the user
// @Success 200 {array} ent.Business
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /followed/businesses [get]
func (fc *FollowController) getFollowedBusinessesByUser(c *gin.Context) error {
	userID := c.GetString("user")

	businesses, err := fc.service.GetFollowedBusinessesByUser(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utility.ProcessResponse(nil, "failed", err.Error()))
		return nil
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(businesses, "success", "Successfully retrieved followed businesses"))
	return nil
}

// @Summary Follow a user
// @ID user-follow-user
// @Tags Follow
// @Produce json
// @Param userID path string true "User ID"
// @Security Bearer
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Follow a specific User
// @Success 201 {object} string "User successfully followed the user"
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /follow/user/{userID} [post]
func (fc *FollowController) followUserToUser(c *gin.Context) error {
	followedID := c.Param("userID")
	followerID := c.GetString("user")

	err := fc.service.FollowUserToUser(c, followerID, followedID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utility.ProcessResponse(nil, "failed", err.Error()))
		return nil
	}

	c.JSON(http.StatusCreated, utility.ProcessResponse(nil, "success", "User successfully followed the user"))
	return nil
}

// @Summary Unfollow a user
// @ID user-unfollow-user
// @Tags Follow
// @Produce json
// @Param userID path string true "User ID"
// @Security Bearer
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Unfollow a specific User
// @Success 200 {object} string "User successfully unfollowed the user"
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /unfollow/user/{userID} [delete]
func (fc *FollowController) unfollowUserToUser(c *gin.Context) error {
	followedID := c.Param("userID")
	followerID := c.GetString("user")

	err := fc.service.UnfollowUserToUser(c, followerID, followedID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utility.ProcessResponse(nil, "failed", err.Error()))
		return nil
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(nil, "success", "User successfully unfollowed the user"))
	return nil
}

// @Summary Get followed users by a user
// @ID get-followed-users
// @Tags Follow
// @Produce json
// @Security Bearer
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Get all Users followed by the user
// @Success 200 {array} ent.User
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /followed/users [get]
func (fc *FollowController) getFollowedUsersByUser(c *gin.Context) error {
	userID := c.GetString("userID")

	users, err := fc.service.GetFollowedUsersByUser(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utility.ProcessResponse(nil, "failed", err.Error()))
		return nil
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(users, "success", "Successfully retrieved followed users"))
	return nil
}

// @Summary Follow a place
// @ID follow-place
// @Tags Follow
// @Produce json
// @Param placeID path string true "Place ID"
// @Security Bearer
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Follow a specific Place by the user
// @Success 201 {object} string "User successfully followed the place"
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /follow/place/{placeID} [post]
func (fc *FollowController) followUserToPlace(c *gin.Context) error {
	placeID := c.Param("placeID")
	userID := c.GetString("userID")

	err := fc.service.FollowUserToPlace(c, userID, placeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utility.ProcessResponse(nil, "failed", err.Error()))
		return nil
	}

	c.JSON(http.StatusCreated, utility.ProcessResponse(nil, "success", "User successfully followed the place"))
	return nil
}

// @Summary Unfollow a place
// @ID unfollow-place
// @Tags Follow
// @Produce json
// @Param placeID path string true "Place ID"
// @Security Bearer
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Unfollow a specific Place by the user
// @Success 200 {object} string "User successfully unfollowed the place"
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /unfollow/place/{placeID} [delete]
func (fc *FollowController) unfollowUserToPlace(c *gin.Context) error {
	placeID := c.Param("placeID")
	userID := c.GetString("user")

	err := fc.service.UnfollowUserToPlace(c, userID, placeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utility.ProcessResponse(nil, "failed", err.Error()))
		return nil
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(nil, "success", "User successfully unfollowed the place"))
	return nil
}

// @Summary Get followed places by a user
// @ID get-followed-places
// @Tags Follow
// @Produce json
// @Security Bearer
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Get all Places followed by the user
// @Success 200 {array} ent.Place
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /followed/places [get]
func (fc *FollowController) getFollowedPlacesByUser(c *gin.Context) error {
	userID := c.GetString("user")

	places, err := fc.service.GetFollowedPlacesByUser(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utility.ProcessResponse(nil, "failed", err.Error()))
		return nil
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(places, "success", "Successfully retrieved followed places"))
	return nil
}

// @Summary Follow an event
// @ID follow-event
// @Tags Follow
// @Produce json
// @Param eventID path string true "Event ID"
// @Security Bearer
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Follow a specific Event by the user
// @Success 201 {object} string "User successfully followed the event"
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /follow/event/{eventID} [post]
func (fc *FollowController) followUserToEvent(c *gin.Context) error {
	eventID := c.Param("eventID")
	userID := c.GetString("user")

	err := fc.service.FollowUserToEvent(c, userID, eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utility.ProcessResponse(nil, "failed", err.Error()))
		return nil
	}

	c.JSON(http.StatusCreated, utility.ProcessResponse(nil, "success", "User successfully followed the event"))
	return nil
}

// @Summary Unfollow an event
// @ID unfollow-event
// @Tags Follow
// @Produce json
// @Param eventID path string true "Event ID"
// @Security Bearer
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Unfollow a specific Event by the user
// @Success 200 {object} string "User successfully unfollowed the event"
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /unfollow/event/{eventID} [delete]
func (fc *FollowController) unfollowUserToEvent(c *gin.Context) error {
	eventID := c.Param("eventID")
	userID := c.GetString("user")

	err := fc.service.UnfollowUserToEvent(c, userID, eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utility.ProcessResponse(nil, "failed", err.Error()))
		return nil
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(nil, "success", "User successfully unfollowed the event"))
	return nil
}

// @Summary Get followed events by a user
// @ID get-followed-events
// @Tags Follow
// @Produce json
// @Security Bearer
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Get all Events followed by the user
// @Success 200 {array} ent.Event
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /followed/events [get]
func (fc *FollowController) getFollowedEventsByUser(c *gin.Context) error {
	userID := c.GetString("user")

	events, err := fc.service.GetFollowedEventsByUser(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utility.ProcessResponse(nil, "failed", err.Error()))
		return nil
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(events, "success", "Successfully retrieved followed events"))
	return nil
}