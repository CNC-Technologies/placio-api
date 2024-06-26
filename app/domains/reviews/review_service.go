package reviews

import (
	"context"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
	"log"
	"mime/multipart"
	"placio-app/domains/media"
	"placio-app/ent"
	"placio-app/ent/business"
	"placio-app/ent/event"
	"placio-app/ent/place"
	"placio-app/ent/review"
	"placio-app/ent/user"
	"placio-pkg/errors"
	"strconv"
	"time"
)

type ReviewStats struct {
	ScoreCounts map[string]int `json:"score_counts"`
	Total       int            `json:"total"`
}

// ReviewService represents the contract for your review related operations.
type ReviewService interface {
	RatePlace(placeID, userID string, score float64, content string, files []*multipart.FileHeader) (*ent.Review, error)
	RateEvent(eventID, userID string, score float64, content string, files []*multipart.FileHeader) (*ent.Review, error)
	RateBusiness(businessID, userID string, score float64, content string, files []*multipart.FileHeader) (*ent.Review, error)
	RemoveReview(reviewID, userID string) error
	GetReviewByID(reviewID string) (*ent.Review, error)
	GetReviewsByUserID(userID string) ([]*ent.Review, error)
	GetReviewsByPlaceID(placeID string) ([]*ent.Review, error)
	GetReviewsByEventID(eventID string) ([]*ent.Review, error)
	GetReviewsByBusinessID(businessID string) ([]*ent.Review, error)
	GetReviewsByScore(score float64) ([]*ent.Review, error)
	GetReviewsInDateRange(startDate, endDate time.Time) ([]*ent.Review, error)
	LikeReview(reviewID, userID string) error
	DislikeReview(reviewID, userID string) error
	UpdateReviewContent(reviewID, userID, newContent, score string) error
	AddMediaToReview(reviewID string, media []*ent.Media) error
	// GetReviewsByLikeCount FlagReview(reviewID, userID string) error
	//AddResponseToReview(reviewID, userID, response string) error
	GetReviewsByLikeCount() ([]*ent.Review, error)
	GetReviewsByDislikeCount() ([]*ent.Review, error)
	GetReviewByIDTypeID(typeId, typeToReview, lastId string, limit int) ([]*ent.Review, string, ReviewStats, error)
}

type Reviewable interface {
	GetID() string
}

type ReviewablePlace struct {
	*ent.Place
}

func (rp ReviewablePlace) GetID() string {
	return rp.ID
}

type ReviewableEvent struct {
	*ent.Event
}

func (re ReviewableEvent) GetID() string {
	return re.ID
}

type ReviewableBusiness struct {
	*ent.Business
}

func (rb ReviewableBusiness) GetID() string {
	return rb.ID
}

type ReviewServiceImpl struct {
	client       *ent.Client
	mediaService media.MediaService
}

func NewReviewService(client *ent.Client, mediaService media.MediaService) ReviewService {
	return &ReviewServiceImpl{client: client, mediaService: mediaService}
}

func (rs *ReviewServiceImpl) rateItem(item Reviewable, userID string, score float64, content string, files []*multipart.FileHeader) (*ent.Review, error) {
	user, err := rs.client.User.Get(context.Background(), userID)
	if err != nil {
		return nil, err
	}
	fmt.Println("rateItem", item.GetID(), userID, score, content)

	// Create a new review
	reviewCreate := rs.client.Review.Create().
		SetID(uuid.New().String()).
		SetUser(user).
		SetScore(score).
		SetContent(content).
		SetCreatedAt(time.Now())

	// Use type assertions to determine the type of the reviewable and set the correct field
	switch v := item.(type) {
	case ReviewablePlace:
		reviewCreate.SetPlace(v.Place)
	case ReviewableEvent:
		reviewCreate.SetEvent(v.Event)
	case ReviewableBusiness:
		reviewCreate.SetBusiness(v.Business)
	default:
		return nil, errors.New("invalid reviewable type")
	}

	// Save the review
	reviewResp, err := reviewCreate.Save(context.Background())
	log.Println("reviewResp", reviewResp)

	if err != nil {
		log.Println("reviewResp error", err.Error())
		return nil, err
	}

	if len(files) > 0 {
		go func() {
			media, err := rs.mediaService.UploadAndCreateMedia(context.Background(), files)
			if err != nil {
				log.Println("error uploading media", err.Error())
				return
			}
			err = rs.AddMediaToReview(reviewResp.ID, media)
			if err != nil {
				log.Println("error adding media to review", err.Error())
				return
			}
		}()
	}

	return reviewResp, nil
}

// RatePlace allows a user to rate and review a place.
func (rs *ReviewServiceImpl) RatePlace(placeID, userID string, score float64, content string, files []*multipart.FileHeader) (*ent.Review, error) {
	fmt.Println("RatePlace")
	place, err := rs.client.Place.Get(context.Background(), placeID)
	if err != nil {
		return nil, err
	}
	fmt.Println("RatePlace2")
	return rs.rateItem(ReviewablePlace{place}, userID, score, content, files)
}

func (rs *ReviewServiceImpl) GetReviewByIDTypeID(typeId, typeToReview, lastId string, limit int) ([]*ent.Review, string, ReviewStats, error) {

	// Define a struct to hold all data
	type result struct {
		Reviews []*ent.Review
		Stats   ReviewStats
		NextID  string
		Err     error
	}

	resultCh := make(chan result)

	go func() {
		var r result

		r.Stats.ScoreCounts = make(map[string]int)

		baseQuery := rs.client.Review.Query().Limit(limit + 1)
		if lastId != "" {
			baseQuery = baseQuery.Where(review.IDGT(lastId))
		}

		switch typeToReview {
		case "place":
			r.Reviews, r.Err = baseQuery.
				Where(review.HasPlaceWith(place.ID(typeId))).
				WithLikes().
				WithMedias().
				WithUser().
				All(context.Background())
		case "event":
			r.Reviews, r.Err = baseQuery.
				Where(review.HasEventWith(event.ID(typeId))).
				WithLikes().
				WithMedias().
				WithUser().
				All(context.Background())
		case "business":
			r.Reviews, r.Err = baseQuery.
				Where(review.HasBusinessWith(business.ID(typeId))).
				WithLikes().
				WithMedias().
				WithUser().
				All(context.Background())
		default:
			r.Err = errors.New("invalid typeToReview")
		}

		if r.Err == nil && len(r.Reviews) == limit+1 {
			r.NextID = r.Reviews[len(r.Reviews)-1].ID
			r.Reviews = r.Reviews[:limit]
		}

		for _, review := range r.Reviews {
			scoreStr := fmt.Sprintf("%f", review.Score)
			r.Stats.ScoreCounts[scoreStr]++
		}
		r.Stats.Total = len(r.Reviews)

		resultCh <- r
	}()

	r := <-resultCh

	return r.Reviews, r.NextID, r.Stats, r.Err
}

// RemoveReview allows a user to remove a review.
func (rs *ReviewServiceImpl) RemoveReview(reviewID, userID string) error {
	review, err := rs.client.Review.Get(context.Background(), reviewID)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}
	//if review.Edges.User.ID != userID {
	//	return errors.New("user does not have permission to delete this review")
	//}
	return rs.client.Review.DeleteOne(review).Exec(context.Background())
}

// GetReviewByID retrieves a review by its ID.
func (rs *ReviewServiceImpl) GetReviewByID(reviewID string) (*ent.Review, error) {
	reviews, err := rs.client.Review.
		Query().
		Where(review.ID(reviewID)).
		WithUser().
		WithLikes().
		WithMedias().
		WithPlace().
		First(context.Background())
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	return reviews, nil
}

// UpdateReviewContent allows a user to update the content of their review.
func (rs *ReviewServiceImpl) UpdateReviewContent(reviewID, userID, newContent, score string) error {
	review, err := rs.client.Review.Get(context.Background(), reviewID)
	if err != nil {
		return err
	}
	//if review.Edges.User.ID != userID {
	//	return errors.New("user does not have permission to update this review")
	//}
	scoreToUpdate, _ := strconv.ParseFloat(score, 64)
	log.Println("updating review", review, newContent)
	_, err = rs.client.Review.
		UpdateOne(review).
		SetContent(newContent).
		SetScore(scoreToUpdate).
		Save(context.Background())
	return err
}

// AddMediaToReview allows a user to add media to their review.
func (rs *ReviewServiceImpl) AddMediaToReview(reviewID string, media []*ent.Media) error {
	_, err := rs.client.Review.UpdateOneID(reviewID).AddMedias(media...).Save(context.Background())
	return err
}

// GetReviewsByUserID retrieves all reviews by a user.
func (rs *ReviewServiceImpl) GetReviewsByUserID(userID string) ([]*ent.Review, error) {
	return rs.client.Review.Query().Where(review.HasUserWith(user.ID(userID))).All(context.Background())
}

// GetReviewsByPlaceID retrieves all reviews for a place.
func (rs *ReviewServiceImpl) GetReviewsByPlaceID(placeID string) ([]*ent.Review, error) {
	return rs.client.Review.Query().Where(review.HasPlaceWith(place.ID(placeID))).All(context.Background())
}

// GetReviewsByEventID retrieves all reviews for an event.
func (rs *ReviewServiceImpl) GetReviewsByEventID(eventID string) ([]*ent.Review, error) {
	return rs.client.Review.Query().Where(review.HasEventWith(event.ID(eventID))).All(context.Background())
}

// GetReviewsByBusinessID retrieves all reviews for a business.
func (rs *ReviewServiceImpl) GetReviewsByBusinessID(businessID string) ([]*ent.Review, error) {
	return rs.client.Review.Query().Where(review.HasBusinessWith(business.ID(businessID))).All(context.Background())
}

// GetReviewsByScore retrieves all reviews with a given score.
func (rs *ReviewServiceImpl) GetReviewsByScore(score float64) ([]*ent.Review, error) {
	return rs.client.Review.Query().Where(review.ScoreEQ(score)).All(context.Background())
}

// GetReviewsInDateRange retrieves all reviews in a given date range.
func (rs *ReviewServiceImpl) GetReviewsInDateRange(startDate, endDate time.Time) ([]*ent.Review, error) {
	return rs.client.Review.Query().Where(review.CreatedAtGTE(startDate), review.CreatedAtLTE(endDate)).All(context.Background())
}

// GetReviewsByLikeCount retrieves all reviews by like count.
func (rs *ReviewServiceImpl) GetReviewsByLikeCount() ([]*ent.Review, error) {
	return rs.client.Review.Query().Order(ent.Desc(review.FieldLikeCount)).All(context.Background())
}

// GetReviewsByDislikeCount retrieves all reviews by dislike count.
func (rs *ReviewServiceImpl) GetReviewsByDislikeCount() ([]*ent.Review, error) {
	return rs.client.Review.Query().Order(ent.Desc(review.FieldDislikeCount)).All(context.Background())
}

// LikeReview // LikeReview allows a user to like a review.
func (rs *ReviewServiceImpl) LikeReview(reviewID, userID string) error {
	_, err := rs.client.Review.UpdateOneID(reviewID).AddLikeCount(1).Save(context.Background())
	return err
}

// DislikeReview // DislikeReview allows a user to dislike a review.
func (rs *ReviewServiceImpl) DislikeReview(reviewID, userID string) error {
	_, err := rs.client.Review.UpdateOneID(reviewID).AddDislikeCount(1).Save(context.Background())
	// get a user and add the review to their disliked reviews
	//user, err := rs.client.User.Get(context.Background(), userID)
	//if err != nil {
	//	return err
	//}
	return err
}

// RateEvent allows a user to rate and review an event.
func (rs *ReviewServiceImpl) RateEvent(eventID, userID string, score float64, content string, files []*multipart.FileHeader) (*ent.Review, error) {
	event, err := rs.client.Event.Get(context.Background(), eventID)
	if err != nil {
		return nil, err
	}

	reviewable := ReviewableEvent{Event: event}

	return rs.rateItem(reviewable, userID, score, content, files)
}

// RateBusiness allows a user to rate and review a business.
func (rs *ReviewServiceImpl) RateBusiness(businessID, userID string, score float64, content string, files []*multipart.FileHeader) (*ent.Review, error) {
	business, err := rs.client.Business.Get(context.Background(), businessID)
	if err != nil {
		return nil, err
	}

	reviewable := ReviewableBusiness{Business: business}

	return rs.rateItem(reviewable, userID, score, content, files)
}
