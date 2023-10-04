// Code generated by ent, DO NOT EDIT.

package website

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the website type in the database.
	Label = "website"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDomainName holds the string denoting the domainname field in the database.
	FieldDomainName = "domain_name"
	// FieldHeadingText holds the string denoting the heading_text field in the database.
	FieldHeadingText = "heading_text"
	// FieldBusinessLogo holds the string denoting the business_logo field in the database.
	FieldBusinessLogo = "business_logo"
	// FieldBusinessName holds the string denoting the business_name field in the database.
	FieldBusinessName = "business_name"
	// FieldBannerSectionBackgroundImage holds the string denoting the banner_section_background_image field in the database.
	FieldBannerSectionBackgroundImage = "banner_section_background_image"
	// FieldBannerSectionBackgroundColor holds the string denoting the banner_section_background_color field in the database.
	FieldBannerSectionBackgroundColor = "banner_section_background_color"
	// FieldBannerSectionText holds the string denoting the banner_section_text field in the database.
	FieldBannerSectionText = "banner_section_text"
	// FieldThreeItemsSectionHeadingText holds the string denoting the three_items_section_heading_text field in the database.
	FieldThreeItemsSectionHeadingText = "three_items_section_heading_text"
	// FieldThreeItemsSectionDetailsText holds the string denoting the three_items_section_details_text field in the database.
	FieldThreeItemsSectionDetailsText = "three_items_section_details_text"
	// FieldThreeItemsSectionItemOneText holds the string denoting the three_items_section_item_one_text field in the database.
	FieldThreeItemsSectionItemOneText = "three_items_section_item_one_text"
	// FieldThreeItemsSectionItemTwoText holds the string denoting the three_items_section_item_two_text field in the database.
	FieldThreeItemsSectionItemTwoText = "three_items_section_item_two_text"
	// FieldThreeItemsSectionItemThreeText holds the string denoting the three_items_section_item_three_text field in the database.
	FieldThreeItemsSectionItemThreeText = "three_items_section_item_three_text"
	// FieldBannerTwoSectionBackgroundImage holds the string denoting the banner_two_section_background_image field in the database.
	FieldBannerTwoSectionBackgroundImage = "banner_two_section_background_image"
	// FieldBannerTwoSectionBackgroundColor holds the string denoting the banner_two_section_background_color field in the database.
	FieldBannerTwoSectionBackgroundColor = "banner_two_section_background_color"
	// FieldBannerTwoLeftSectionHeadingText holds the string denoting the banner_two_left_section_heading_text field in the database.
	FieldBannerTwoLeftSectionHeadingText = "banner_two_left_section_heading_text"
	// FieldBannerTwoLeftSectionDetailsText holds the string denoting the banner_two_left_section_details_text field in the database.
	FieldBannerTwoLeftSectionDetailsText = "banner_two_left_section_details_text"
	// FieldBannerTwoLeftSectionButtonText holds the string denoting the banner_two_left_section_button_text field in the database.
	FieldBannerTwoLeftSectionButtonText = "banner_two_left_section_button_text"
	// FieldBannerTwoLeftSectionButtonLink holds the string denoting the banner_two_left_section_button_link field in the database.
	FieldBannerTwoLeftSectionButtonLink = "banner_two_left_section_button_link"
	// FieldBannerTwoRightSideImage holds the string denoting the banner_two_right_side_image field in the database.
	FieldBannerTwoRightSideImage = "banner_two_right_side_image"
	// FieldAchievementsSection holds the string denoting the achievements_section field in the database.
	FieldAchievementsSection = "achievements_section"
	// FieldInventorySectionHeadingText holds the string denoting the inventory_section_heading_text field in the database.
	FieldInventorySectionHeadingText = "inventory_section_heading_text"
	// FieldCreationDate holds the string denoting the creationdate field in the database.
	FieldCreationDate = "creation_date"
	// FieldLastUpdated holds the string denoting the lastupdated field in the database.
	FieldLastUpdated = "last_updated"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldKeywords holds the string denoting the keywords field in the database.
	FieldKeywords = "keywords"
	// FieldLanguage holds the string denoting the language field in the database.
	FieldLanguage = "language"
	// FieldLogo holds the string denoting the logo field in the database.
	FieldLogo = "logo"
	// FieldFavicon holds the string denoting the favicon field in the database.
	FieldFavicon = "favicon"
	// FieldFacebook holds the string denoting the facebook field in the database.
	FieldFacebook = "facebook"
	// FieldTwitter holds the string denoting the twitter field in the database.
	FieldTwitter = "twitter"
	// FieldInstagram holds the string denoting the instagram field in the database.
	FieldInstagram = "instagram"
	// FieldYoutube holds the string denoting the youtube field in the database.
	FieldYoutube = "youtube"
	// FieldLinkedin holds the string denoting the linkedin field in the database.
	FieldLinkedin = "linkedin"
	// FieldPinterest holds the string denoting the pinterest field in the database.
	FieldPinterest = "pinterest"
	// FieldMapCoordinates holds the string denoting the mapcoordinates field in the database.
	FieldMapCoordinates = "map_coordinates"
	// FieldLongitude holds the string denoting the longitude field in the database.
	FieldLongitude = "longitude"
	// FieldLatitude holds the string denoting the latitude field in the database.
	FieldLatitude = "latitude"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldCity holds the string denoting the city field in the database.
	FieldCity = "city"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldCountry holds the string denoting the country field in the database.
	FieldCountry = "country"
	// FieldZipCode holds the string denoting the zipcode field in the database.
	FieldZipCode = "zip_code"
	// FieldPhoneNumber holds the string denoting the phonenumber field in the database.
	FieldPhoneNumber = "phone_number"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldMetaTags holds the string denoting the metatags field in the database.
	FieldMetaTags = "meta_tags"
	// EdgeBusiness holds the string denoting the business edge name in mutations.
	EdgeBusiness = "business"
	// EdgeCustomBlocks holds the string denoting the customblocks edge name in mutations.
	EdgeCustomBlocks = "customBlocks"
	// EdgeAssets holds the string denoting the assets edge name in mutations.
	EdgeAssets = "assets"
	// Table holds the table name of the website in the database.
	Table = "websites"
	// BusinessTable is the table that holds the business relation/edge.
	BusinessTable = "websites"
	// BusinessInverseTable is the table name for the Business entity.
	// It exists in this package in order to avoid circular dependency with the "business" package.
	BusinessInverseTable = "businesses"
	// BusinessColumn is the table column denoting the business relation/edge.
	BusinessColumn = "business_websites"
	// CustomBlocksTable is the table that holds the customBlocks relation/edge.
	CustomBlocksTable = "custom_blocks"
	// CustomBlocksInverseTable is the table name for the CustomBlock entity.
	// It exists in this package in order to avoid circular dependency with the "customblock" package.
	CustomBlocksInverseTable = "custom_blocks"
	// CustomBlocksColumn is the table column denoting the customBlocks relation/edge.
	CustomBlocksColumn = "website_custom_blocks"
	// AssetsTable is the table that holds the assets relation/edge.
	AssetsTable = "media"
	// AssetsInverseTable is the table name for the Media entity.
	// It exists in this package in order to avoid circular dependency with the "media" package.
	AssetsInverseTable = "media"
	// AssetsColumn is the table column denoting the assets relation/edge.
	AssetsColumn = "website_assets"
)

// Columns holds all SQL columns for website fields.
var Columns = []string{
	FieldID,
	FieldDomainName,
	FieldHeadingText,
	FieldBusinessLogo,
	FieldBusinessName,
	FieldBannerSectionBackgroundImage,
	FieldBannerSectionBackgroundColor,
	FieldBannerSectionText,
	FieldThreeItemsSectionHeadingText,
	FieldThreeItemsSectionDetailsText,
	FieldThreeItemsSectionItemOneText,
	FieldThreeItemsSectionItemTwoText,
	FieldThreeItemsSectionItemThreeText,
	FieldBannerTwoSectionBackgroundImage,
	FieldBannerTwoSectionBackgroundColor,
	FieldBannerTwoLeftSectionHeadingText,
	FieldBannerTwoLeftSectionDetailsText,
	FieldBannerTwoLeftSectionButtonText,
	FieldBannerTwoLeftSectionButtonLink,
	FieldBannerTwoRightSideImage,
	FieldAchievementsSection,
	FieldInventorySectionHeadingText,
	FieldCreationDate,
	FieldLastUpdated,
	FieldTitle,
	FieldDescription,
	FieldKeywords,
	FieldLanguage,
	FieldLogo,
	FieldFavicon,
	FieldFacebook,
	FieldTwitter,
	FieldInstagram,
	FieldYoutube,
	FieldLinkedin,
	FieldPinterest,
	FieldMapCoordinates,
	FieldLongitude,
	FieldLatitude,
	FieldAddress,
	FieldCity,
	FieldState,
	FieldCountry,
	FieldZipCode,
	FieldPhoneNumber,
	FieldEmail,
	FieldMetaTags,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "websites"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"business_websites",
	"template_websites",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreationDate holds the default value on creation for the "creationDate" field.
	DefaultCreationDate func() time.Time
	// UpdateDefaultLastUpdated holds the default value on update for the "lastUpdated" field.
	UpdateDefaultLastUpdated func() time.Time
)

// OrderOption defines the ordering options for the Website queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDomainName orders the results by the domainName field.
func ByDomainName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDomainName, opts...).ToFunc()
}

// ByHeadingText orders the results by the heading_text field.
func ByHeadingText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHeadingText, opts...).ToFunc()
}

// ByBusinessLogo orders the results by the business_logo field.
func ByBusinessLogo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBusinessLogo, opts...).ToFunc()
}

// ByBusinessName orders the results by the business_name field.
func ByBusinessName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBusinessName, opts...).ToFunc()
}

// ByBannerSectionBackgroundImage orders the results by the banner_section_background_image field.
func ByBannerSectionBackgroundImage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBannerSectionBackgroundImage, opts...).ToFunc()
}

// ByBannerSectionBackgroundColor orders the results by the banner_section_background_color field.
func ByBannerSectionBackgroundColor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBannerSectionBackgroundColor, opts...).ToFunc()
}

// ByBannerSectionText orders the results by the banner_section_text field.
func ByBannerSectionText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBannerSectionText, opts...).ToFunc()
}

// ByThreeItemsSectionHeadingText orders the results by the three_items_section_heading_text field.
func ByThreeItemsSectionHeadingText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldThreeItemsSectionHeadingText, opts...).ToFunc()
}

// ByThreeItemsSectionDetailsText orders the results by the three_items_section_details_text field.
func ByThreeItemsSectionDetailsText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldThreeItemsSectionDetailsText, opts...).ToFunc()
}

// ByThreeItemsSectionItemOneText orders the results by the three_items_section_item_one_text field.
func ByThreeItemsSectionItemOneText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldThreeItemsSectionItemOneText, opts...).ToFunc()
}

// ByThreeItemsSectionItemTwoText orders the results by the three_items_section_item_two_text field.
func ByThreeItemsSectionItemTwoText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldThreeItemsSectionItemTwoText, opts...).ToFunc()
}

// ByThreeItemsSectionItemThreeText orders the results by the three_items_section_item_three_text field.
func ByThreeItemsSectionItemThreeText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldThreeItemsSectionItemThreeText, opts...).ToFunc()
}

// ByBannerTwoSectionBackgroundImage orders the results by the banner_two_section_background_image field.
func ByBannerTwoSectionBackgroundImage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBannerTwoSectionBackgroundImage, opts...).ToFunc()
}

// ByBannerTwoSectionBackgroundColor orders the results by the banner_two_section_background_color field.
func ByBannerTwoSectionBackgroundColor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBannerTwoSectionBackgroundColor, opts...).ToFunc()
}

// ByBannerTwoLeftSectionHeadingText orders the results by the banner_two_left_section_heading_text field.
func ByBannerTwoLeftSectionHeadingText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBannerTwoLeftSectionHeadingText, opts...).ToFunc()
}

// ByBannerTwoLeftSectionDetailsText orders the results by the banner_two_left_section_details_text field.
func ByBannerTwoLeftSectionDetailsText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBannerTwoLeftSectionDetailsText, opts...).ToFunc()
}

// ByBannerTwoLeftSectionButtonText orders the results by the banner_two_left_section_button_text field.
func ByBannerTwoLeftSectionButtonText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBannerTwoLeftSectionButtonText, opts...).ToFunc()
}

// ByBannerTwoLeftSectionButtonLink orders the results by the banner_two_left_section_button_link field.
func ByBannerTwoLeftSectionButtonLink(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBannerTwoLeftSectionButtonLink, opts...).ToFunc()
}

// ByBannerTwoRightSideImage orders the results by the banner_two_right_side_image field.
func ByBannerTwoRightSideImage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBannerTwoRightSideImage, opts...).ToFunc()
}

// ByInventorySectionHeadingText orders the results by the Inventory_section_heading_text field.
func ByInventorySectionHeadingText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInventorySectionHeadingText, opts...).ToFunc()
}

// ByCreationDate orders the results by the creationDate field.
func ByCreationDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreationDate, opts...).ToFunc()
}

// ByLastUpdated orders the results by the lastUpdated field.
func ByLastUpdated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastUpdated, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByKeywords orders the results by the keywords field.
func ByKeywords(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKeywords, opts...).ToFunc()
}

// ByLanguage orders the results by the language field.
func ByLanguage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLanguage, opts...).ToFunc()
}

// ByLogo orders the results by the logo field.
func ByLogo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLogo, opts...).ToFunc()
}

// ByFavicon orders the results by the favicon field.
func ByFavicon(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFavicon, opts...).ToFunc()
}

// ByFacebook orders the results by the facebook field.
func ByFacebook(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFacebook, opts...).ToFunc()
}

// ByTwitter orders the results by the twitter field.
func ByTwitter(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTwitter, opts...).ToFunc()
}

// ByInstagram orders the results by the instagram field.
func ByInstagram(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInstagram, opts...).ToFunc()
}

// ByYoutube orders the results by the youtube field.
func ByYoutube(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldYoutube, opts...).ToFunc()
}

// ByLinkedin orders the results by the linkedin field.
func ByLinkedin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLinkedin, opts...).ToFunc()
}

// ByPinterest orders the results by the pinterest field.
func ByPinterest(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPinterest, opts...).ToFunc()
}

// ByLongitude orders the results by the longitude field.
func ByLongitude(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLongitude, opts...).ToFunc()
}

// ByLatitude orders the results by the latitude field.
func ByLatitude(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLatitude, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// ByCity orders the results by the city field.
func ByCity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCity, opts...).ToFunc()
}

// ByState orders the results by the state field.
func ByState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldState, opts...).ToFunc()
}

// ByCountry orders the results by the country field.
func ByCountry(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCountry, opts...).ToFunc()
}

// ByZipCode orders the results by the zipCode field.
func ByZipCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldZipCode, opts...).ToFunc()
}

// ByPhoneNumber orders the results by the phoneNumber field.
func ByPhoneNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhoneNumber, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByBusinessField orders the results by business field.
func ByBusinessField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBusinessStep(), sql.OrderByField(field, opts...))
	}
}

// ByCustomBlocksCount orders the results by customBlocks count.
func ByCustomBlocksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCustomBlocksStep(), opts...)
	}
}

// ByCustomBlocks orders the results by customBlocks terms.
func ByCustomBlocks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCustomBlocksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAssetsCount orders the results by assets count.
func ByAssetsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAssetsStep(), opts...)
	}
}

// ByAssets orders the results by assets terms.
func ByAssets(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAssetsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newBusinessStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BusinessInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, BusinessTable, BusinessColumn),
	)
}
func newCustomBlocksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CustomBlocksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CustomBlocksTable, CustomBlocksColumn),
	)
}
func newAssetsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AssetsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AssetsTable, AssetsColumn),
	)
}
