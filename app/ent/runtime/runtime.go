// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"placio-app/ent/accountsettings"
	"placio-app/ent/amenity"
	"placio-app/ent/booking"
	"placio-app/ent/business"
	"placio-app/ent/businessfollowbusiness"
	"placio-app/ent/businessfollowevent"
	"placio-app/ent/businessfollowuser"
	"placio-app/ent/category"
	"placio-app/ent/comment"
	"placio-app/ent/event"
	"placio-app/ent/featurerelease"
	"placio-app/ent/help"
	"placio-app/ent/inventoryattribute"
	"placio-app/ent/inventorytype"
	"placio-app/ent/like"
	"placio-app/ent/media"
	"placio-app/ent/menu"
	"placio-app/ent/place"
	"placio-app/ent/placeinventory"
	"placio-app/ent/placeinventoryattribute"
	"placio-app/ent/post"
	"placio-app/ent/rating"
	"placio-app/ent/reservation"
	"placio-app/ent/review"
	"placio-app/ent/room"
	"placio-app/ent/schema"
	"placio-app/ent/ticket"
	"placio-app/ent/ticketoption"
	"placio-app/ent/transactionhistory"
	"placio-app/ent/user"
	"placio-app/ent/userbusiness"
	"placio-app/ent/userfollowbusiness"
	"placio-app/ent/userfollowevent"
	"placio-app/ent/userfollowplace"
	"placio-app/ent/userfollowuser"
	"placio-app/ent/userlikeplace"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accountsettingsFields := schema.AccountSettings{}.Fields()
	_ = accountsettingsFields
	// accountsettingsDescID is the schema descriptor for id field.
	accountsettingsDescID := accountsettingsFields[0].Descriptor()
	// accountsettings.IDValidator is a validator for the "id" field. It is called by the builders before save.
	accountsettings.IDValidator = accountsettingsDescID.Validators[0].(func(string) error)
	amenityFields := schema.Amenity{}.Fields()
	_ = amenityFields
	// amenityDescID is the schema descriptor for id field.
	amenityDescID := amenityFields[0].Descriptor()
	// amenity.IDValidator is a validator for the "id" field. It is called by the builders before save.
	amenity.IDValidator = amenityDescID.Validators[0].(func(string) error)
	bookingFields := schema.Booking{}.Fields()
	_ = bookingFields
	// bookingDescID is the schema descriptor for id field.
	bookingDescID := bookingFields[0].Descriptor()
	// booking.IDValidator is a validator for the "id" field. It is called by the builders before save.
	booking.IDValidator = bookingDescID.Validators[0].(func(string) error)
	businessHooks := schema.Business{}.Hooks()
	business.Hooks[0] = businessHooks[0]
	businessFields := schema.Business{}.Fields()
	_ = businessFields
	// businessDescCoverImage is the schema descriptor for cover_image field.
	businessDescCoverImage := businessFields[4].Descriptor()
	// business.DefaultCoverImage holds the default value on creation for the cover_image field.
	business.DefaultCoverImage = businessDescCoverImage.Default.(string)
	// businessDescFollowerCount is the schema descriptor for follower_count field.
	businessDescFollowerCount := businessFields[16].Descriptor()
	// business.DefaultFollowerCount holds the default value on creation for the follower_count field.
	business.DefaultFollowerCount = businessDescFollowerCount.Default.(int)
	// businessDescFollowingCount is the schema descriptor for following_count field.
	businessDescFollowingCount := businessFields[17].Descriptor()
	// business.DefaultFollowingCount holds the default value on creation for the following_count field.
	business.DefaultFollowingCount = businessDescFollowingCount.Default.(int)
	// businessDescID is the schema descriptor for id field.
	businessDescID := businessFields[0].Descriptor()
	// business.IDValidator is a validator for the "id" field. It is called by the builders before save.
	business.IDValidator = businessDescID.Validators[0].(func(string) error)
	businessfollowbusinessFields := schema.BusinessFollowBusiness{}.Fields()
	_ = businessfollowbusinessFields
	// businessfollowbusinessDescCreatedAt is the schema descriptor for CreatedAt field.
	businessfollowbusinessDescCreatedAt := businessfollowbusinessFields[1].Descriptor()
	// businessfollowbusiness.DefaultCreatedAt holds the default value on creation for the CreatedAt field.
	businessfollowbusiness.DefaultCreatedAt = businessfollowbusinessDescCreatedAt.Default.(func() time.Time)
	// businessfollowbusinessDescUpdatedAt is the schema descriptor for UpdatedAt field.
	businessfollowbusinessDescUpdatedAt := businessfollowbusinessFields[2].Descriptor()
	// businessfollowbusiness.UpdateDefaultUpdatedAt holds the default value on update for the UpdatedAt field.
	businessfollowbusiness.UpdateDefaultUpdatedAt = businessfollowbusinessDescUpdatedAt.UpdateDefault.(func() time.Time)
	businessfolloweventFields := schema.BusinessFollowEvent{}.Fields()
	_ = businessfolloweventFields
	// businessfolloweventDescCreatedAt is the schema descriptor for createdAt field.
	businessfolloweventDescCreatedAt := businessfolloweventFields[1].Descriptor()
	// businessfollowevent.DefaultCreatedAt holds the default value on creation for the createdAt field.
	businessfollowevent.DefaultCreatedAt = businessfolloweventDescCreatedAt.Default.(func() time.Time)
	// businessfolloweventDescUpdatedAt is the schema descriptor for updatedAt field.
	businessfolloweventDescUpdatedAt := businessfolloweventFields[2].Descriptor()
	// businessfollowevent.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	businessfollowevent.DefaultUpdatedAt = businessfolloweventDescUpdatedAt.Default.(func() time.Time)
	// businessfollowevent.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	businessfollowevent.UpdateDefaultUpdatedAt = businessfolloweventDescUpdatedAt.UpdateDefault.(func() time.Time)
	businessfollowuserFields := schema.BusinessFollowUser{}.Fields()
	_ = businessfollowuserFields
	// businessfollowuserDescCreatedAt is the schema descriptor for CreatedAt field.
	businessfollowuserDescCreatedAt := businessfollowuserFields[1].Descriptor()
	// businessfollowuser.DefaultCreatedAt holds the default value on creation for the CreatedAt field.
	businessfollowuser.DefaultCreatedAt = businessfollowuserDescCreatedAt.Default.(func() time.Time)
	// businessfollowuserDescUpdatedAt is the schema descriptor for UpdatedAt field.
	businessfollowuserDescUpdatedAt := businessfollowuserFields[2].Descriptor()
	// businessfollowuser.UpdateDefaultUpdatedAt holds the default value on update for the UpdatedAt field.
	businessfollowuser.UpdateDefaultUpdatedAt = businessfollowuserDescUpdatedAt.UpdateDefault.(func() time.Time)
	categoryFields := schema.Category{}.Fields()
	_ = categoryFields
	// categoryDescID is the schema descriptor for id field.
	categoryDescID := categoryFields[0].Descriptor()
	// category.IDValidator is a validator for the "id" field. It is called by the builders before save.
	category.IDValidator = categoryDescID.Validators[0].(func(string) error)
	commentFields := schema.Comment{}.Fields()
	_ = commentFields
	// commentDescContent is the schema descriptor for Content field.
	commentDescContent := commentFields[1].Descriptor()
	// comment.ContentValidator is a validator for the "Content" field. It is called by the builders before save.
	comment.ContentValidator = commentDescContent.Validators[0].(func(string) error)
	// commentDescCreatedAt is the schema descriptor for CreatedAt field.
	commentDescCreatedAt := commentFields[2].Descriptor()
	// comment.DefaultCreatedAt holds the default value on creation for the CreatedAt field.
	comment.DefaultCreatedAt = commentDescCreatedAt.Default.(func() time.Time)
	// commentDescUpdatedAt is the schema descriptor for UpdatedAt field.
	commentDescUpdatedAt := commentFields[3].Descriptor()
	// comment.UpdateDefaultUpdatedAt holds the default value on update for the UpdatedAt field.
	comment.UpdateDefaultUpdatedAt = commentDescUpdatedAt.UpdateDefault.(func() time.Time)
	// commentDescID is the schema descriptor for id field.
	commentDescID := commentFields[0].Descriptor()
	// comment.IDValidator is a validator for the "id" field. It is called by the builders before save.
	comment.IDValidator = commentDescID.Validators[0].(func(string) error)
	eventHooks := schema.Event{}.Hooks()
	event.Hooks[0] = eventHooks[0]
	eventFields := schema.Event{}.Fields()
	_ = eventFields
	// eventDescCoverImage is the schema descriptor for cover_image field.
	eventDescCoverImage := eventFields[32].Descriptor()
	// event.DefaultCoverImage holds the default value on creation for the cover_image field.
	event.DefaultCoverImage = eventDescCoverImage.Default.(string)
	// eventDescCreatedAt is the schema descriptor for createdAt field.
	eventDescCreatedAt := eventFields[33].Descriptor()
	// event.DefaultCreatedAt holds the default value on creation for the createdAt field.
	event.DefaultCreatedAt = eventDescCreatedAt.Default.(func() time.Time)
	// eventDescUpdatedAt is the schema descriptor for updatedAt field.
	eventDescUpdatedAt := eventFields[34].Descriptor()
	// event.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	event.DefaultUpdatedAt = eventDescUpdatedAt.Default.(func() time.Time)
	// event.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	event.UpdateDefaultUpdatedAt = eventDescUpdatedAt.UpdateDefault.(func() time.Time)
	// eventDescFollowerCount is the schema descriptor for follower_count field.
	eventDescFollowerCount := eventFields[40].Descriptor()
	// event.DefaultFollowerCount holds the default value on creation for the follower_count field.
	event.DefaultFollowerCount = eventDescFollowerCount.Default.(int)
	// eventDescFollowingCount is the schema descriptor for following_count field.
	eventDescFollowingCount := eventFields[41].Descriptor()
	// event.DefaultFollowingCount holds the default value on creation for the following_count field.
	event.DefaultFollowingCount = eventDescFollowingCount.Default.(int)
	// eventDescIsPremium is the schema descriptor for is_Premium field.
	eventDescIsPremium := eventFields[42].Descriptor()
	// event.DefaultIsPremium holds the default value on creation for the is_Premium field.
	event.DefaultIsPremium = eventDescIsPremium.Default.(bool)
	// eventDescIsPublished is the schema descriptor for is_published field.
	eventDescIsPublished := eventFields[43].Descriptor()
	// event.DefaultIsPublished holds the default value on creation for the is_published field.
	event.DefaultIsPublished = eventDescIsPublished.Default.(bool)
	// eventDescIsOnline is the schema descriptor for is_Online field.
	eventDescIsOnline := eventFields[44].Descriptor()
	// event.DefaultIsOnline holds the default value on creation for the is_Online field.
	event.DefaultIsOnline = eventDescIsOnline.Default.(bool)
	// eventDescIsFree is the schema descriptor for is_Free field.
	eventDescIsFree := eventFields[45].Descriptor()
	// event.DefaultIsFree holds the default value on creation for the is_Free field.
	event.DefaultIsFree = eventDescIsFree.Default.(bool)
	// eventDescIsPaid is the schema descriptor for is_Paid field.
	eventDescIsPaid := eventFields[46].Descriptor()
	// event.DefaultIsPaid holds the default value on creation for the is_Paid field.
	event.DefaultIsPaid = eventDescIsPaid.Default.(bool)
	// eventDescIsOnlineOnly is the schema descriptor for is_Online_Only field.
	eventDescIsOnlineOnly := eventFields[47].Descriptor()
	// event.DefaultIsOnlineOnly holds the default value on creation for the is_Online_Only field.
	event.DefaultIsOnlineOnly = eventDescIsOnlineOnly.Default.(bool)
	// eventDescIsInPersonOnly is the schema descriptor for is_In_Person_Only field.
	eventDescIsInPersonOnly := eventFields[48].Descriptor()
	// event.DefaultIsInPersonOnly holds the default value on creation for the is_In_Person_Only field.
	event.DefaultIsInPersonOnly = eventDescIsInPersonOnly.Default.(bool)
	// eventDescIsHybrid is the schema descriptor for is_Hybrid field.
	eventDescIsHybrid := eventFields[49].Descriptor()
	// event.DefaultIsHybrid holds the default value on creation for the is_Hybrid field.
	event.DefaultIsHybrid = eventDescIsHybrid.Default.(bool)
	// eventDescIsOnlineAndInPerson is the schema descriptor for is_Online_And_In_Person field.
	eventDescIsOnlineAndInPerson := eventFields[50].Descriptor()
	// event.DefaultIsOnlineAndInPerson holds the default value on creation for the is_Online_And_In_Person field.
	event.DefaultIsOnlineAndInPerson = eventDescIsOnlineAndInPerson.Default.(bool)
	// eventDescIsOnlineAndInPersonOnly is the schema descriptor for is_Online_And_In_Person_Only field.
	eventDescIsOnlineAndInPersonOnly := eventFields[51].Descriptor()
	// event.DefaultIsOnlineAndInPersonOnly holds the default value on creation for the is_Online_And_In_Person_Only field.
	event.DefaultIsOnlineAndInPersonOnly = eventDescIsOnlineAndInPersonOnly.Default.(bool)
	// eventDescIsOnlineAndInPersonOrHybrid is the schema descriptor for is_Online_And_In_Person_Or_Hybrid field.
	eventDescIsOnlineAndInPersonOrHybrid := eventFields[52].Descriptor()
	// event.DefaultIsOnlineAndInPersonOrHybrid holds the default value on creation for the is_Online_And_In_Person_Or_Hybrid field.
	event.DefaultIsOnlineAndInPersonOrHybrid = eventDescIsOnlineAndInPersonOrHybrid.Default.(bool)
	// eventDescLikedByCurrentUser is the schema descriptor for likedByCurrentUser field.
	eventDescLikedByCurrentUser := eventFields[53].Descriptor()
	// event.DefaultLikedByCurrentUser holds the default value on creation for the likedByCurrentUser field.
	event.DefaultLikedByCurrentUser = eventDescLikedByCurrentUser.Default.(bool)
	// eventDescFollowedByCurrentUser is the schema descriptor for followedByCurrentUser field.
	eventDescFollowedByCurrentUser := eventFields[54].Descriptor()
	// event.DefaultFollowedByCurrentUser holds the default value on creation for the followedByCurrentUser field.
	event.DefaultFollowedByCurrentUser = eventDescFollowedByCurrentUser.Default.(bool)
	featurereleaseFields := schema.FeatureRelease{}.Fields()
	_ = featurereleaseFields
	// featurereleaseDescReleaseDate is the schema descriptor for release_date field.
	featurereleaseDescReleaseDate := featurereleaseFields[4].Descriptor()
	// featurerelease.DefaultReleaseDate holds the default value on creation for the release_date field.
	featurerelease.DefaultReleaseDate = featurereleaseDescReleaseDate.Default.(func() time.Time)
	// featurereleaseDescID is the schema descriptor for id field.
	featurereleaseDescID := featurereleaseFields[0].Descriptor()
	// featurerelease.IDValidator is a validator for the "id" field. It is called by the builders before save.
	featurerelease.IDValidator = featurereleaseDescID.Validators[0].(func(string) error)
	helpFields := schema.Help{}.Fields()
	_ = helpFields
	// helpDescStatus is the schema descriptor for status field.
	helpDescStatus := helpFields[5].Descriptor()
	// help.DefaultStatus holds the default value on creation for the status field.
	help.DefaultStatus = helpDescStatus.Default.(string)
	// helpDescUserID is the schema descriptor for user_id field.
	helpDescUserID := helpFields[6].Descriptor()
	// help.UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	help.UserIDValidator = helpDescUserID.Validators[0].(func(string) error)
	inventoryattributeFields := schema.InventoryAttribute{}.Fields()
	_ = inventoryattributeFields
	// inventoryattributeDescIsMandatory is the schema descriptor for is_mandatory field.
	inventoryattributeDescIsMandatory := inventoryattributeFields[2].Descriptor()
	// inventoryattribute.DefaultIsMandatory holds the default value on creation for the is_mandatory field.
	inventoryattribute.DefaultIsMandatory = inventoryattributeDescIsMandatory.Default.(bool)
	// inventoryattributeDescID is the schema descriptor for id field.
	inventoryattributeDescID := inventoryattributeFields[0].Descriptor()
	// inventoryattribute.IDValidator is a validator for the "id" field. It is called by the builders before save.
	inventoryattribute.IDValidator = inventoryattributeDescID.Validators[0].(func(string) error)
	inventorytypeFields := schema.InventoryType{}.Fields()
	_ = inventorytypeFields
	// inventorytypeDescID is the schema descriptor for id field.
	inventorytypeDescID := inventorytypeFields[0].Descriptor()
	// inventorytype.IDValidator is a validator for the "id" field. It is called by the builders before save.
	inventorytype.IDValidator = inventorytypeDescID.Validators[0].(func(string) error)
	likeFields := schema.Like{}.Fields()
	_ = likeFields
	// likeDescCreatedAt is the schema descriptor for createdAt field.
	likeDescCreatedAt := likeFields[1].Descriptor()
	// like.DefaultCreatedAt holds the default value on creation for the createdAt field.
	like.DefaultCreatedAt = likeDescCreatedAt.Default.(func() time.Time)
	// likeDescUpdatedAt is the schema descriptor for updatedAt field.
	likeDescUpdatedAt := likeFields[2].Descriptor()
	// like.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	like.UpdateDefaultUpdatedAt = likeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// likeDescID is the schema descriptor for id field.
	likeDescID := likeFields[0].Descriptor()
	// like.IDValidator is a validator for the "id" field. It is called by the builders before save.
	like.IDValidator = likeDescID.Validators[0].(func(string) error)
	mediaFields := schema.Media{}.Fields()
	_ = mediaFields
	// mediaDescCreatedAt is the schema descriptor for CreatedAt field.
	mediaDescCreatedAt := mediaFields[3].Descriptor()
	// media.DefaultCreatedAt holds the default value on creation for the CreatedAt field.
	media.DefaultCreatedAt = mediaDescCreatedAt.Default.(func() time.Time)
	// mediaDescUpdatedAt is the schema descriptor for UpdatedAt field.
	mediaDescUpdatedAt := mediaFields[4].Descriptor()
	// media.DefaultUpdatedAt holds the default value on creation for the UpdatedAt field.
	media.DefaultUpdatedAt = mediaDescUpdatedAt.Default.(func() time.Time)
	// media.UpdateDefaultUpdatedAt holds the default value on update for the UpdatedAt field.
	media.UpdateDefaultUpdatedAt = mediaDescUpdatedAt.UpdateDefault.(func() time.Time)
	// mediaDescLikeCount is the schema descriptor for likeCount field.
	mediaDescLikeCount := mediaFields[5].Descriptor()
	// media.DefaultLikeCount holds the default value on creation for the likeCount field.
	media.DefaultLikeCount = mediaDescLikeCount.Default.(int)
	// mediaDescDislikeCount is the schema descriptor for dislikeCount field.
	mediaDescDislikeCount := mediaFields[6].Descriptor()
	// media.DefaultDislikeCount holds the default value on creation for the dislikeCount field.
	media.DefaultDislikeCount = mediaDescDislikeCount.Default.(int)
	// mediaDescID is the schema descriptor for id field.
	mediaDescID := mediaFields[0].Descriptor()
	// media.IDValidator is a validator for the "id" field. It is called by the builders before save.
	media.IDValidator = mediaDescID.Validators[0].(func(string) error)
	menuFields := schema.Menu{}.Fields()
	_ = menuFields
	// menuDescID is the schema descriptor for id field.
	menuDescID := menuFields[0].Descriptor()
	// menu.IDValidator is a validator for the "id" field. It is called by the builders before save.
	menu.IDValidator = menuDescID.Validators[0].(func(string) error)
	placeHooks := schema.Place{}.Hooks()
	place.Hooks[0] = placeHooks[0]
	placeFields := schema.Place{}.Fields()
	_ = placeFields
	// placeDescCoverImage is the schema descriptor for cover_image field.
	placeDescCoverImage := placeFields[8].Descriptor()
	// place.DefaultCoverImage holds the default value on creation for the cover_image field.
	place.DefaultCoverImage = placeDescCoverImage.Default.(string)
	// placeDescFollowerCount is the schema descriptor for follower_count field.
	placeDescFollowerCount := placeFields[29].Descriptor()
	// place.DefaultFollowerCount holds the default value on creation for the follower_count field.
	place.DefaultFollowerCount = placeDescFollowerCount.Default.(int)
	// placeDescFollowingCount is the schema descriptor for following_count field.
	placeDescFollowingCount := placeFields[30].Descriptor()
	// place.DefaultFollowingCount holds the default value on creation for the following_count field.
	place.DefaultFollowingCount = placeDescFollowingCount.Default.(int)
	// placeDescIsPremium is the schema descriptor for is_Premium field.
	placeDescIsPremium := placeFields[31].Descriptor()
	// place.DefaultIsPremium holds the default value on creation for the is_Premium field.
	place.DefaultIsPremium = placeDescIsPremium.Default.(bool)
	// placeDescIsPublished is the schema descriptor for is_published field.
	placeDescIsPublished := placeFields[32].Descriptor()
	// place.DefaultIsPublished holds the default value on creation for the is_published field.
	place.DefaultIsPublished = placeDescIsPublished.Default.(bool)
	// placeDescLikedByCurrentUser is the schema descriptor for likedByCurrentUser field.
	placeDescLikedByCurrentUser := placeFields[33].Descriptor()
	// place.DefaultLikedByCurrentUser holds the default value on creation for the likedByCurrentUser field.
	place.DefaultLikedByCurrentUser = placeDescLikedByCurrentUser.Default.(bool)
	// placeDescFollowedByCurrentUser is the schema descriptor for followedByCurrentUser field.
	placeDescFollowedByCurrentUser := placeFields[34].Descriptor()
	// place.DefaultFollowedByCurrentUser holds the default value on creation for the followedByCurrentUser field.
	place.DefaultFollowedByCurrentUser = placeDescFollowedByCurrentUser.Default.(bool)
	// placeDescID is the schema descriptor for id field.
	placeDescID := placeFields[0].Descriptor()
	// place.IDValidator is a validator for the "id" field. It is called by the builders before save.
	place.IDValidator = placeDescID.Validators[0].(func(string) error)
	placeinventoryFields := schema.PlaceInventory{}.Fields()
	_ = placeinventoryFields
	// placeinventoryDescLastUpdated is the schema descriptor for last_updated field.
	placeinventoryDescLastUpdated := placeinventoryFields[11].Descriptor()
	// placeinventory.DefaultLastUpdated holds the default value on creation for the last_updated field.
	placeinventory.DefaultLastUpdated = placeinventoryDescLastUpdated.Default.(func() time.Time)
	// placeinventoryDescID is the schema descriptor for id field.
	placeinventoryDescID := placeinventoryFields[0].Descriptor()
	// placeinventory.IDValidator is a validator for the "id" field. It is called by the builders before save.
	placeinventory.IDValidator = placeinventoryDescID.Validators[0].(func(string) error)
	placeinventoryattributeFields := schema.PlaceInventoryAttribute{}.Fields()
	_ = placeinventoryattributeFields
	// placeinventoryattributeDescID is the schema descriptor for id field.
	placeinventoryattributeDescID := placeinventoryattributeFields[0].Descriptor()
	// placeinventoryattribute.IDValidator is a validator for the "id" field. It is called by the builders before save.
	placeinventoryattribute.IDValidator = placeinventoryattributeDescID.Validators[0].(func(string) error)
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescContent is the schema descriptor for Content field.
	postDescContent := postFields[1].Descriptor()
	// post.ContentValidator is a validator for the "Content" field. It is called by the builders before save.
	post.ContentValidator = postDescContent.Validators[0].(func(string) error)
	// postDescCreatedAt is the schema descriptor for CreatedAt field.
	postDescCreatedAt := postFields[2].Descriptor()
	// post.DefaultCreatedAt holds the default value on creation for the CreatedAt field.
	post.DefaultCreatedAt = postDescCreatedAt.Default.(func() time.Time)
	// postDescUpdatedAt is the schema descriptor for UpdatedAt field.
	postDescUpdatedAt := postFields[3].Descriptor()
	// post.UpdateDefaultUpdatedAt holds the default value on update for the UpdatedAt field.
	post.UpdateDefaultUpdatedAt = postDescUpdatedAt.UpdateDefault.(func() time.Time)
	// postDescLikedByMe is the schema descriptor for LikedByMe field.
	postDescLikedByMe := postFields[5].Descriptor()
	// post.DefaultLikedByMe holds the default value on creation for the LikedByMe field.
	post.DefaultLikedByMe = postDescLikedByMe.Default.(bool)
	// postDescLikeCount is the schema descriptor for LikeCount field.
	postDescLikeCount := postFields[6].Descriptor()
	// post.DefaultLikeCount holds the default value on creation for the LikeCount field.
	post.DefaultLikeCount = postDescLikeCount.Default.(int)
	// postDescCommentCount is the schema descriptor for CommentCount field.
	postDescCommentCount := postFields[7].Descriptor()
	// post.DefaultCommentCount holds the default value on creation for the CommentCount field.
	post.DefaultCommentCount = postDescCommentCount.Default.(int)
	// postDescShareCount is the schema descriptor for ShareCount field.
	postDescShareCount := postFields[8].Descriptor()
	// post.DefaultShareCount holds the default value on creation for the ShareCount field.
	post.DefaultShareCount = postDescShareCount.Default.(int)
	// postDescViewCount is the schema descriptor for ViewCount field.
	postDescViewCount := postFields[9].Descriptor()
	// post.DefaultViewCount holds the default value on creation for the ViewCount field.
	post.DefaultViewCount = postDescViewCount.Default.(int)
	// postDescIsSponsored is the schema descriptor for IsSponsored field.
	postDescIsSponsored := postFields[10].Descriptor()
	// post.DefaultIsSponsored holds the default value on creation for the IsSponsored field.
	post.DefaultIsSponsored = postDescIsSponsored.Default.(bool)
	// postDescIsPromoted is the schema descriptor for IsPromoted field.
	postDescIsPromoted := postFields[11].Descriptor()
	// post.DefaultIsPromoted holds the default value on creation for the IsPromoted field.
	post.DefaultIsPromoted = postDescIsPromoted.Default.(bool)
	// postDescIsBoosted is the schema descriptor for IsBoosted field.
	postDescIsBoosted := postFields[12].Descriptor()
	// post.DefaultIsBoosted holds the default value on creation for the IsBoosted field.
	post.DefaultIsBoosted = postDescIsBoosted.Default.(bool)
	// postDescIsPinned is the schema descriptor for IsPinned field.
	postDescIsPinned := postFields[13].Descriptor()
	// post.DefaultIsPinned holds the default value on creation for the IsPinned field.
	post.DefaultIsPinned = postDescIsPinned.Default.(bool)
	// postDescIsHidden is the schema descriptor for IsHidden field.
	postDescIsHidden := postFields[14].Descriptor()
	// post.DefaultIsHidden holds the default value on creation for the IsHidden field.
	post.DefaultIsHidden = postDescIsHidden.Default.(bool)
	// postDescRelevanceScore is the schema descriptor for RelevanceScore field.
	postDescRelevanceScore := postFields[15].Descriptor()
	// post.DefaultRelevanceScore holds the default value on creation for the RelevanceScore field.
	post.DefaultRelevanceScore = postDescRelevanceScore.Default.(int)
	// postDescID is the schema descriptor for id field.
	postDescID := postFields[0].Descriptor()
	// post.IDValidator is a validator for the "id" field. It is called by the builders before save.
	post.IDValidator = postDescID.Validators[0].(func(string) error)
	ratingHooks := schema.Rating{}.Hooks()
	rating.Hooks[0] = ratingHooks[0]
	ratingFields := schema.Rating{}.Fields()
	_ = ratingFields
	// ratingDescScore is the schema descriptor for score field.
	ratingDescScore := ratingFields[1].Descriptor()
	// rating.ScoreValidator is a validator for the "score" field. It is called by the builders before save.
	rating.ScoreValidator = func() func(int) error {
		validators := ratingDescScore.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
		}
		return func(score int) error {
			for _, fn := range fns {
				if err := fn(score); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// ratingDescRatedAt is the schema descriptor for ratedAt field.
	ratingDescRatedAt := ratingFields[3].Descriptor()
	// rating.DefaultRatedAt holds the default value on creation for the ratedAt field.
	rating.DefaultRatedAt = ratingDescRatedAt.Default.(func() time.Time)
	reservationFields := schema.Reservation{}.Fields()
	_ = reservationFields
	// reservationDescID is the schema descriptor for id field.
	reservationDescID := reservationFields[0].Descriptor()
	// reservation.IDValidator is a validator for the "id" field. It is called by the builders before save.
	reservation.IDValidator = reservationDescID.Validators[0].(func(string) error)
	reviewHooks := schema.Review{}.Hooks()
	review.Hooks[0] = reviewHooks[0]
	reviewFields := schema.Review{}.Fields()
	_ = reviewFields
	// reviewDescScore is the schema descriptor for score field.
	reviewDescScore := reviewFields[1].Descriptor()
	// review.ScoreValidator is a validator for the "score" field. It is called by the builders before save.
	review.ScoreValidator = func() func(float64) error {
		validators := reviewDescScore.Validators
		fns := [...]func(float64) error{
			validators[0].(func(float64) error),
			validators[1].(func(float64) error),
		}
		return func(score float64) error {
			for _, fn := range fns {
				if err := fn(score); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// reviewDescCreatedAt is the schema descriptor for createdAt field.
	reviewDescCreatedAt := reviewFields[3].Descriptor()
	// review.DefaultCreatedAt holds the default value on creation for the createdAt field.
	review.DefaultCreatedAt = reviewDescCreatedAt.Default.(func() time.Time)
	// reviewDescLikeCount is the schema descriptor for likeCount field.
	reviewDescLikeCount := reviewFields[4].Descriptor()
	// review.DefaultLikeCount holds the default value on creation for the likeCount field.
	review.DefaultLikeCount = reviewDescLikeCount.Default.(int)
	// reviewDescDislikeCount is the schema descriptor for dislikeCount field.
	reviewDescDislikeCount := reviewFields[5].Descriptor()
	// review.DefaultDislikeCount holds the default value on creation for the dislikeCount field.
	review.DefaultDislikeCount = reviewDescDislikeCount.Default.(int)
	// reviewDescFlag is the schema descriptor for flag field.
	reviewDescFlag := reviewFields[6].Descriptor()
	// review.DefaultFlag holds the default value on creation for the flag field.
	review.DefaultFlag = reviewDescFlag.Default.(string)
	roomFields := schema.Room{}.Fields()
	_ = roomFields
	// roomDescID is the schema descriptor for id field.
	roomDescID := roomFields[0].Descriptor()
	// room.IDValidator is a validator for the "id" field. It is called by the builders before save.
	room.IDValidator = roomDescID.Validators[0].(func(string) error)
	ticketFields := schema.Ticket{}.Fields()
	_ = ticketFields
	// ticketDescCreatedAt is the schema descriptor for createdAt field.
	ticketDescCreatedAt := ticketFields[1].Descriptor()
	// ticket.DefaultCreatedAt holds the default value on creation for the createdAt field.
	ticket.DefaultCreatedAt = ticketDescCreatedAt.Default.(func() time.Time)
	// ticketDescUpdatedAt is the schema descriptor for updatedAt field.
	ticketDescUpdatedAt := ticketFields[2].Descriptor()
	// ticket.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	ticket.DefaultUpdatedAt = ticketDescUpdatedAt.Default.(func() time.Time)
	// ticket.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	ticket.UpdateDefaultUpdatedAt = ticketDescUpdatedAt.UpdateDefault.(func() time.Time)
	ticketoptionFields := schema.TicketOption{}.Fields()
	_ = ticketoptionFields
	// ticketoptionDescCreatedAt is the schema descriptor for createdAt field.
	ticketoptionDescCreatedAt := ticketoptionFields[1].Descriptor()
	// ticketoption.DefaultCreatedAt holds the default value on creation for the createdAt field.
	ticketoption.DefaultCreatedAt = ticketoptionDescCreatedAt.Default.(func() time.Time)
	// ticketoptionDescUpdatedAt is the schema descriptor for updatedAt field.
	ticketoptionDescUpdatedAt := ticketoptionFields[2].Descriptor()
	// ticketoption.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	ticketoption.DefaultUpdatedAt = ticketoptionDescUpdatedAt.Default.(func() time.Time)
	// ticketoption.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	ticketoption.UpdateDefaultUpdatedAt = ticketoptionDescUpdatedAt.UpdateDefault.(func() time.Time)
	transactionhistoryFields := schema.TransactionHistory{}.Fields()
	_ = transactionhistoryFields
	// transactionhistoryDescDate is the schema descriptor for date field.
	transactionhistoryDescDate := transactionhistoryFields[2].Descriptor()
	// transactionhistory.DefaultDate holds the default value on creation for the date field.
	transactionhistory.DefaultDate = transactionhistoryDescDate.Default.(func() time.Time)
	userHooks := schema.User{}.Hooks()
	user.Hooks[0] = userHooks[0]
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCoverImage is the schema descriptor for cover_image field.
	userDescCoverImage := userFields[4].Descriptor()
	// user.DefaultCoverImage holds the default value on creation for the cover_image field.
	user.DefaultCoverImage = userDescCoverImage.Default.(string)
	// userDescBio is the schema descriptor for bio field.
	userDescBio := userFields[11].Descriptor()
	// user.DefaultBio holds the default value on creation for the bio field.
	user.DefaultBio = userDescBio.Default.(string)
	// userDescFollowerCount is the schema descriptor for follower_count field.
	userDescFollowerCount := userFields[17].Descriptor()
	// user.DefaultFollowerCount holds the default value on creation for the follower_count field.
	user.DefaultFollowerCount = userDescFollowerCount.Default.(int)
	// userDescFollowingCount is the schema descriptor for following_count field.
	userDescFollowingCount := userFields[18].Descriptor()
	// user.DefaultFollowingCount holds the default value on creation for the following_count field.
	user.DefaultFollowingCount = userDescFollowingCount.Default.(int)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.IDValidator is a validator for the "id" field. It is called by the builders before save.
	user.IDValidator = userDescID.Validators[0].(func(string) error)
	userbusinessFields := schema.UserBusiness{}.Fields()
	_ = userbusinessFields
	// userbusinessDescPermissions is the schema descriptor for permissions field.
	userbusinessDescPermissions := userbusinessFields[2].Descriptor()
	// userbusiness.PermissionsValidator is a validator for the "permissions" field. It is called by the builders before save.
	userbusiness.PermissionsValidator = userbusinessDescPermissions.Validators[0].(func(string) error)
	// userbusinessDescID is the schema descriptor for id field.
	userbusinessDescID := userbusinessFields[0].Descriptor()
	// userbusiness.IDValidator is a validator for the "id" field. It is called by the builders before save.
	userbusiness.IDValidator = userbusinessDescID.Validators[0].(func(string) error)
	userfollowbusinessFields := schema.UserFollowBusiness{}.Fields()
	_ = userfollowbusinessFields
	// userfollowbusinessDescCreatedAt is the schema descriptor for CreatedAt field.
	userfollowbusinessDescCreatedAt := userfollowbusinessFields[1].Descriptor()
	// userfollowbusiness.DefaultCreatedAt holds the default value on creation for the CreatedAt field.
	userfollowbusiness.DefaultCreatedAt = userfollowbusinessDescCreatedAt.Default.(func() time.Time)
	// userfollowbusinessDescUpdatedAt is the schema descriptor for UpdatedAt field.
	userfollowbusinessDescUpdatedAt := userfollowbusinessFields[2].Descriptor()
	// userfollowbusiness.UpdateDefaultUpdatedAt holds the default value on update for the UpdatedAt field.
	userfollowbusiness.UpdateDefaultUpdatedAt = userfollowbusinessDescUpdatedAt.UpdateDefault.(func() time.Time)
	userfolloweventFields := schema.UserFollowEvent{}.Fields()
	_ = userfolloweventFields
	// userfolloweventDescCreatedAt is the schema descriptor for createdAt field.
	userfolloweventDescCreatedAt := userfolloweventFields[1].Descriptor()
	// userfollowevent.DefaultCreatedAt holds the default value on creation for the createdAt field.
	userfollowevent.DefaultCreatedAt = userfolloweventDescCreatedAt.Default.(func() time.Time)
	// userfolloweventDescUpdatedAt is the schema descriptor for updatedAt field.
	userfolloweventDescUpdatedAt := userfolloweventFields[2].Descriptor()
	// userfollowevent.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	userfollowevent.DefaultUpdatedAt = userfolloweventDescUpdatedAt.Default.(func() time.Time)
	// userfollowevent.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	userfollowevent.UpdateDefaultUpdatedAt = userfolloweventDescUpdatedAt.UpdateDefault.(func() time.Time)
	userfollowplaceFields := schema.UserFollowPlace{}.Fields()
	_ = userfollowplaceFields
	// userfollowplaceDescCreatedAt is the schema descriptor for CreatedAt field.
	userfollowplaceDescCreatedAt := userfollowplaceFields[1].Descriptor()
	// userfollowplace.DefaultCreatedAt holds the default value on creation for the CreatedAt field.
	userfollowplace.DefaultCreatedAt = userfollowplaceDescCreatedAt.Default.(func() time.Time)
	// userfollowplaceDescUpdatedAt is the schema descriptor for UpdatedAt field.
	userfollowplaceDescUpdatedAt := userfollowplaceFields[2].Descriptor()
	// userfollowplace.UpdateDefaultUpdatedAt holds the default value on update for the UpdatedAt field.
	userfollowplace.UpdateDefaultUpdatedAt = userfollowplaceDescUpdatedAt.UpdateDefault.(func() time.Time)
	userfollowuserFields := schema.UserFollowUser{}.Fields()
	_ = userfollowuserFields
	// userfollowuserDescCreatedAt is the schema descriptor for CreatedAt field.
	userfollowuserDescCreatedAt := userfollowuserFields[1].Descriptor()
	// userfollowuser.DefaultCreatedAt holds the default value on creation for the CreatedAt field.
	userfollowuser.DefaultCreatedAt = userfollowuserDescCreatedAt.Default.(func() time.Time)
	// userfollowuserDescUpdatedAt is the schema descriptor for UpdatedAt field.
	userfollowuserDescUpdatedAt := userfollowuserFields[2].Descriptor()
	// userfollowuser.UpdateDefaultUpdatedAt holds the default value on update for the UpdatedAt field.
	userfollowuser.UpdateDefaultUpdatedAt = userfollowuserDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userfollowuserDescID is the schema descriptor for id field.
	userfollowuserDescID := userfollowuserFields[0].Descriptor()
	// userfollowuser.IDValidator is a validator for the "id" field. It is called by the builders before save.
	userfollowuser.IDValidator = userfollowuserDescID.Validators[0].(func(string) error)
	userlikeplaceFields := schema.UserLikePlace{}.Fields()
	_ = userlikeplaceFields
	// userlikeplaceDescCreatedAt is the schema descriptor for CreatedAt field.
	userlikeplaceDescCreatedAt := userlikeplaceFields[1].Descriptor()
	// userlikeplace.DefaultCreatedAt holds the default value on creation for the CreatedAt field.
	userlikeplace.DefaultCreatedAt = userlikeplaceDescCreatedAt.Default.(func() time.Time)
	// userlikeplaceDescUpdatedAt is the schema descriptor for UpdatedAt field.
	userlikeplaceDescUpdatedAt := userlikeplaceFields[2].Descriptor()
	// userlikeplace.UpdateDefaultUpdatedAt holds the default value on update for the UpdatedAt field.
	userlikeplace.UpdateDefaultUpdatedAt = userlikeplaceDescUpdatedAt.UpdateDefault.(func() time.Time)
}

const (
	Version = "v0.12.3"                                         // Version of ent codegen.
	Sum     = "h1:N5lO2EOrHpCH5HYfiMOCHYbo+oh5M8GjT0/cx5x6xkk=" // Sum of ent codegen.
)
