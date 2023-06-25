package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Business holds the schema definition for the Business entity.
type Business struct {
	ent.Schema
}

// Fields of the Business.
func (Business) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.String("name"),
		field.String("search_text").Optional(),
		field.Float("relevance_score").Optional(),
	}
}

// Edges of the Business.
func (Business) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("userBusinesses", UserBusiness.Type),
		edge.To("business_account_settings", AccountSettings.Type).
			Unique(),
		edge.To("posts", Post.Type),
		edge.To("followedUsers", BusinessFollowUser.Type),
		edge.To("followerUsers", UserFollowBusiness.Type),
		edge.To("followedBusinesses", BusinessFollowBusiness.Type),
		edge.To("followerBusinesses", BusinessFollowBusiness.Type),
		edge.To("places", Place.Type),
		edge.To("categories", Category.Type),
		edge.To("categoryAssignments", CategoryAssignment.Type),
	}
}
