package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"placio-app/ent"
	_ "placio-app/ent"
	appErr "placio-app/errors"
	"placio-app/service"
	"placio-app/utility"
	"strconv"
)

type ReviewController struct {
	reviewService service.ReviewService
	mediaService  service.MediaService
}

func NewReviewController(reviewService service.ReviewService, mediaService service.MediaService) *ReviewController {
	return &ReviewController{reviewService: reviewService, mediaService: mediaService}
}

func (rc *ReviewController) RegisterRoutes(router *gin.RouterGroup) {
	reviewRouter := router.Group("/reviews")
	{
		reviewRouter.POST("/", utility.Use(rc.rateItem))
		reviewRouter.DELETE("/:reviewID", utility.Use(rc.removeReview))
		reviewRouter.GET("/:reviewID", utility.Use(rc.getReviewByID))
		reviewRouter.PUT("/:reviewID", utility.Use(rc.updateReviewContent))
		reviewRouter.POST("/:reviewID/addMedia", utility.Use(rc.addMediaToReview))
		reviewRouter.GET("/", utility.Use(rc.getReviewsByQuery))
		reviewRouter.POST("/:reviewID/like", utility.Use(rc.likeReview))
		reviewRouter.POST("/:reviewID/dislike", utility.Use(rc.dislikeReview))
	}
}

// @Summary Review a Place, Event, or Business
// @Description Rate a place, event, or business
// @Tags Review
// @Accept  json
// @Produce  json
// @Param type query string true "Item Type (place, event, business)"
// @Param id query string true "Item ID (placeID, eventID, businessID)"
// @Param score body number true "Rating Score"
// @Param content body string true "Review Content"
// @Param Authorization header string true "JWT Token"
// @Success 200 {string} string "Successfully rated [itemType]"
// @Failure 400 {string} string "Invalid item type"
// @Failure 500 {string} string "Error message"
// @Router /reviews/ [post]
func (rc *ReviewController) rateItem(ctx *gin.Context) error {
	itemType := ctx.Query("type")
	if itemType != "place" && itemType != "event" && itemType != "business" {
		return appErr.InvalidItemType
	}
	itemID := ctx.Query("id")
	if itemID == "" {
		return appErr.IDMissing
	}
	userID := ctx.MustGet("user").(string)
	if userID == "" {
		return appErr.ErrUnauthorized
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return err
	}

	files, _ := form.File["files"]

	var score float64
	scoreStr, ok := form.Value["score"]
	if ok && len(scoreStr) > 0 {
		score, err = strconv.ParseFloat(scoreStr[0], 64)
		if err != nil || score < 1 || score > 5 {
			return errors.New("invalid rating score: must be between 1 and 5")
		}
	}

	content := form.Value["content"][0]

	var review *ent.Review

	switch itemType {
	case "place":
		log.Println("rate place")
		review, err = rc.reviewService.RatePlace(itemID, userID, score, content, files)
	case "event":
		review, err = rc.reviewService.RateEvent(itemID, userID, score, content, files)
	case "business":
		review, err = rc.reviewService.RateBusiness(itemID, userID, score, content, files)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item type"})
		return nil
	}

	if err != nil {

		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(review, "success", "Successfully rated "+itemType))
	return nil
}

// @Summary Remove a review
// @Description Remove a review based on reviewID
// @Tags Review
// @Accept  json
// @Produce  json
// @Param reviewID path string true "Review ID"
// @Param Authorization header string true "JWT Token"
// @Success 200 {string} string "Successfully deleted review"
// @Failure 500 {string} string "Error message"
// @Router /reviews/{reviewID} [delete]
func (rc *ReviewController) removeReview(ctx *gin.Context) error {
	reviewID := ctx.Param("reviewID")
	userID := ctx.MustGet("user").(string) // query parameter: userID

	err := rc.reviewService.RemoveReview(reviewID, userID)
	if err != nil {

		return err
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully deleted review"})
	return nil
}

// @Summary Get review by ID
// @Description Retrieve a review using its ID
// @Tags Review
// @Accept  json
// @Produce  json
// @Param reviewID path string true "Review ID"
// @Param Authorization header string true "JWT Token"
// @Success 200 {object} ent.Review "Review data"
// @Failure 500 {string} string "Error message"
// @Router /reviews/{reviewID} [get]
func (rc *ReviewController) getReviewByID(ctx *gin.Context) error {
	reviewID := ctx.Param("reviewID")

	review, err := rc.reviewService.GetReviewByID(reviewID)
	if err != nil {

		return err
	}

	ctx.JSON(http.StatusOK, review)
	return nil
}

// @Summary Update review content
// @Description Update the content of a review based on reviewID
// @Tags Review
// @Accept  json
// @Produce  json
// @Param reviewID path string true "Review ID"
// @Param Authorization header string true "JWT Token"
// @Param content body string true "New Review Content"
// @Success 200 {string} string "Successfully updated review"
// @Failure 500 {string} string "Error message"
// @Router /reviews/{reviewID} [put]
func (rc *ReviewController) updateReviewContent(ctx *gin.Context) error {
	reviewID := ctx.Param("reviewID")
	userID := ctx.MustGet("user").(string) // query parameter: userID
	content := ctx.PostForm("content")

	err := rc.reviewService.UpdateReviewContent(reviewID, userID, content)
	if err != nil {

		return err
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully updated review"})
	return nil
}

// @Summary Add media to a review
// @Description Add media to a review based on reviewID
// @Tags Review
// @Accept  json
// @Produce  json
// @Param reviewID path string true "Review ID"
// @Param string body string true "Media File"
// @Success 200 {string} string "Successfully added media to review"
// @Failure 500 {string} string "Error message"
// @Router /reviews/{reviewID}/addMedia [post]
func (rc *ReviewController) addMediaToReview(ctx *gin.Context) error {
	reviewID := ctx.Param("reviewID")
	//mediaFile := ctx.PostForm("media")

	files, _ := ctx.MultipartForm()
	fileArray := files.File["media"]

	mediaUploaded, err := rc.mediaService.UploadAndCreateMedia(ctx, fileArray)

	err = rc.reviewService.AddMediaToReview(reviewID, mediaUploaded)
	if err != nil {

		return err
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully added media to review"})
	return nil
}

// @Summary Get reviews by query
// @Description Get reviews based on query parameters
// @Tags Review
// @Accept  json
// @Produce  json
// @Success 200 {object} []ent.Review "Array of Reviews"
// @Failure 500 {string} string "Error message"
// @Router /reviews/ [get]
func (rc *ReviewController) getReviewsByQuery(ctx *gin.Context) error {
	// Add logic here to get reviews based on various query parameters
	// You would use something like ctx.Query("parameterName") to get the query parameters
	return nil
}

// @Summary Like a review
// @Description Like a review based on reviewID
// @Tags Review
// @Accept  json
// @Produce  json
// @Param reviewID path string true "Review ID"
// @Param Authorization header string true "JWT Token"
// @Success 200 {string} string "Successfully liked review"
// @Failure 500 {string} string "Error message"
// @Router /reviews/{reviewID}/like [post]
func (rc *ReviewController) likeReview(ctx *gin.Context) error {
	reviewID := ctx.Param("reviewID")
	userID := ctx.MustGet("user").(string) // query parameter: userID

	err := rc.reviewService.LikeReview(reviewID, userID)
	if err != nil {

		return err
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully liked review"})
	return nil
}

// @Summary Dislike a review
// @Description Dislike a review based on reviewID
// @Tags Review
// @Accept  json
// @Produce  json
// @Param reviewID path string true "Review ID"
// @Param Authorization header string true "JWT Token"
// @Success 200 {string} string "Successfully disliked review"
// @Failure 500 {string} string "Error message"
// @Router /reviews/{reviewID}/dislike [post]
func (rc *ReviewController) dislikeReview(ctx *gin.Context) error {
	reviewID := ctx.Param("reviewID")
	userID := ctx.MustGet("user").(string) // query parameter: userID

	err := rc.reviewService.DislikeReview(reviewID, userID)
	if err != nil {

		return err
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully disliked review"})
	return nil
}
