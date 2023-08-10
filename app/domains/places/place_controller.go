package places

import (
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "placio-app/Dto"
	"placio-app/domains/amenities"
	"placio-app/ent"
	_ "placio-app/ent"
	"placio-app/utility"
	"placio-pkg/errors"
	"strconv"
)

type PlaceController struct {
	placeService PlaceService
	cache        utility.RedisClient
}

func NewPlaceController(placeService PlaceService, cache utility.RedisClient) *PlaceController {
	return &PlaceController{placeService: placeService, cache: cache}
}

func (c *PlaceController) RegisterRoutes(router, routerWithoutAuth *gin.RouterGroup) {
	placeRouter := router.Group("/places")
	placeRouterWithoutAuth := routerWithoutAuth.Group("/places")
	{
		placeRouterWithoutAuth.GET("/:id", utility.Use(c.getPlace))
		placeRouter.POST("/", utility.Use(c.createPlace))
		placeRouterWithoutAuth.GET("/", utility.Use(c.getPlacesByFilters))
		placeRouter.PATCH("/:id", utility.Use(c.updatePlace))
		placeRouter.DELETE("/:id", utility.Use(c.deletePlace))
		placeRouter.POST("/:id/amenities", utility.Use(c.addAmenitiesToPlace))
		placeRouter.PATCH("/:id/media", utility.Use(c.addMediaToAPlace))
		placeRouterWithoutAuth.GET("/all", utility.Use(c.getAllPlaces))
	}
}

// @Summary Get a place
// @Description Get a place by ID
// @Tags Place
// @Accept json
// @Produce json
// @Param id path string true "ID of the place to get"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} ent.Place "Successfully retrieved place"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/places/{id} [get]
func (c *PlaceController) getPlace(ctx *gin.Context) error {
	id := ctx.Param("id")

	// get place from cache
	cacheKey := "place:" + id
	return utility.GetDataFromCache[*ent.Place](ctx, &c.cache, c.placeService.GetPlace, id, cacheKey)
}

// @Summary Create a place
// @Description Create a new place
// @Tags Place
// @Accept json
// @Produce json
// @Param place body Dto.CreatePlaceDTO true "Place to create"
// @Param business_id query string true "ID of the business to create the place for"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} ent.Place "Successfully created place"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/places/ [post]
func (c *PlaceController) createPlace(ctx *gin.Context) error {
	var placeData CreatePlaceDTO
	if err := ctx.ShouldBindJSON(&placeData); err != nil {

		return err
	}

	placeData.BusinessID = ctx.Query("business_id")
	if placeData.BusinessID == "" {
		return errors.IDMissing
	}

	place, err := c.placeService.CreatePlace(ctx, placeData)
	if err != nil {

		return err
	}

	ctx.JSON(http.StatusOK, place)
	return nil
}

// @Summary Get all places
// @Description Get all places
// @Tags Place
// @Accept json
// @Produce json
// @Param nextPageToken query string false "Token for the next page of results"
// @Param limit query int false "Number of results to return"
// @Success 200 {array} []ent.Place "Successfully retrieved all places"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/places/all [get]
func (c *PlaceController) getAllPlaces(ctx *gin.Context) error {
	nextPageToken := ctx.Query("nextPageToken")
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}

	placesData, token, err := c.placeService.GetAllPlaces(ctx, nextPageToken, limit)
	if err != nil {

		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(placesData, "success", "places retrieved successfully", token))
	return nil
}

// @Summary Add amenities to a place
// @Description Add amenities to a place by ID
// @Tags Place
// @Accept json
// @Produce json
// @Param id path string true "ID of the place to add amenities to"
// @Param amenity body Dto.AmenityAdditionDTO true "Amenities to add"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} ent.Place "Successfully added amenities to place"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/places/{id}/amenities [post]
func (c *PlaceController) addAmenitiesToPlace(ctx *gin.Context) error {
	id := ctx.Param("id")

	var amenityDTO amenities.AmenityAdditionDTO
	if err := ctx.ShouldBindJSON(&amenityDTO); err != nil {

		return err
	}

	if err := c.placeService.AddAmenitiesToPlace(ctx, id, amenityDTO.AmenityIDs); err != nil {

		return err
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Amenities added successfully",
	})
	return nil
}

// @Summary Add Media to a place
// @Description Add media to a place by ID
// @Tags Place
// @Accept multi
// @Produce json
// @Param id path string true "ID of the place to add amenities to"
// @Param amenity body Dto.AmenityAdditionDTO true "Amenities to add"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} ent.Place "Successfully added amenities to place"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/places/{id}/amenities [post]
func (c *PlaceController) addMediaToAPlace(ctx *gin.Context) error {
	id := ctx.Param("id")

	var amenityDTO amenities.AmenityAdditionDTO
	if err := ctx.ShouldBindJSON(&amenityDTO); err != nil {

		return err
	}

	go func() {
		if err := c.placeService.AddAmenitiesToPlace(ctx, id, amenityDTO.AmenityIDs); err != nil {
			sentry.CaptureException(err)
		}
	}()

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Media added successfully",
	})
	return nil
}

// @Summary Update a place
// @Description Update a place by ID
// @Tags Place
// @Accept json
// @Produce json
// @Param id path string true "ID of the place to update"
// @Param place body Dto.UpdatePlaceDTO true "Place data to update"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} ent.Place "Successfully updated place"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/places/{id} [patch]
func (c *PlaceController) updatePlace(ctx *gin.Context) error {
	id := ctx.Param("id")
	var placeData UpdatePlaceDTO
	if err := ctx.ShouldBindJSON(&placeData); err != nil {

		return err
	}

	place, err := c.placeService.UpdatePlace(ctx, id, placeData)
	if err != nil {

		return err
	}

	ctx.JSON(http.StatusOK, place)
	return nil
}

// @Summary Delete a place
// @Description Delete a place by ID
// @Tags Place
// @Accept json
// @Produce json
// @Param id path string true "ID of the place to delete"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} ent.Place "Successfully deleted place"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/places/{id} [delete]
func (c *PlaceController) deletePlace(ctx *gin.Context) error {
	id := ctx.Param("id")

	err := c.placeService.DeletePlace(ctx, id)
	if err != nil {

		return err
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Successfully deleted place",
	})
	return nil
}

// GetPlaces godoc
// @Summary Get all Places
// @Description Get Places by applying various filters (ID, Name, Type, Country, City, State, Tags, Features)
// @Tags Place
// @Accept  json
// @Produce  json
// @Param filter query service.PlaceFilter false "Filter"
// @Param page query int false "Page Number"
// @Param pageSize query int false "Page Size"
// @Param Authorization header string true "Bearer Token"
// @Success 200 {array} ent.Place
// @Failure 400 {object} Dto.ErrorDTO "Invalid inputs"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized. Invalid or expired token"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /places [get]
func (c *PlaceController) getPlacesByFilters(ctx *gin.Context) error {
	var filter PlaceFilter
	if err := ctx.ShouldBindQuery(&filter); err != nil {

		return err
	}

	page := ctx.Query("nextPageToken")

	pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if err != nil {

		return err
	}

	places, nextPageToken, err := c.placeService.GetPlaces(ctx, &filter, page, pageSize)
	if err != nil {

		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(places, "success", "places retrieved successfully", nextPageToken))
	return nil
}