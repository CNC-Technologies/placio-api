package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type UserFollowUser struct {
	ent.Schema
}

// Fields of the UserFollowUser.
func (UserFollowUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").UpdateDefault(time.Now),
	}
}

// Edges of the UserFollowUser.
func (UserFollowUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("follower", User.Type).
			Ref("followedUsers").
			Unique(),
		edge.From("followed", User.Type).
			Ref("followerUsers").
			Unique(),
	}
}

type UserFollowBusiness struct {
	ent.Schema
}

// Fields of the UserFollowBusiness.
func (UserFollowBusiness) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable(),
		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").UpdateDefault(time.Now),
	}
}

// Edges of the UserFollowBusiness.
func (UserFollowBusiness) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("followedBusinesses").
			Unique(),
		edge.From("business", Business.Type).
			Ref("followerUsers").
			Unique(),
	}
}

type BusinessFollowBusiness struct {
	ent.Schema
}

// Fields of the BusinessFollowBusiness.
func (BusinessFollowBusiness) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable(),
		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").UpdateDefault(time.Now),
	}
}

// Edges of the BusinessFollowBusiness.
func (BusinessFollowBusiness) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("follower", Business.Type).
			Ref("followedBusinesses").
			Unique(),
		edge.From("followed", Business.Type).
			Ref("followerBusinesses").
			Unique(),
	}
}

type BusinessFollowUser struct {
	ent.Schema
}

// Fields of the BusinessFollowUser.
func (BusinessFollowUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable(),
		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").UpdateDefault(time.Now),
	}
}

// Edges of the BusinessFollowUser.
func (BusinessFollowUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("business", Business.Type).
			Ref("followedUsers").
			Unique(),
		edge.From("user", User.Type).
			Ref("followerBusinesses").
			Unique(),
	}
}

// UserFollowPlace holds the schema definition for the UserFollowPlace entity.
type UserFollowPlace struct {
	ent.Schema
}

func (UserFollowPlace) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable(),
		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").UpdateDefault(time.Now),
	}
}

func (UserFollowPlace) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("followedPlaces").Unique(),
		edge.To("place", Place.Type).Unique(),
	}
}

// UserFollowEvent holds the schema definition for the UserFollowEvent entity.
type UserFollowEvent struct {
	ent.Schema
}

// Fields of the UserFollowEvent.
func (UserFollowEvent) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable(),
		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the UserFollowEvent.
func (UserFollowEvent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("userFollowEvents").
			Unique().
			Required(),
		edge.To("event", Event.Type).
			Unique().
			Required(),
	}
}

// BusinessFollowEvent holds the schema definition for the BusinessFollowEvent entity.
type BusinessFollowEvent struct {
	ent.Schema
}

// Fields of the BusinessFollowEvent.
func (BusinessFollowEvent) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable(),
		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the BusinessFollowEvent.
func (BusinessFollowEvent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("business", Business.Type).
			Ref("businessFollowEvents").
			Unique().
			Required(),
		edge.To("event", Event.Type).
			Unique().
			Required(),
	}
}
