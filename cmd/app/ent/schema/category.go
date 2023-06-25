package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			Unique().
			Immutable(),
		field.String("name").Unique(),
		field.String("image").Optional(),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("categoryAssignments", CategoryAssignment.Type),
	}
}

type CategoryAssignment struct {
	ent.Schema
}

// Fields of the CategoryAssignment.
func (CategoryAssignment) Fields() []ent.Field {
	return []ent.Field{
		field.String("entity_id").
			Comment("This represents the ID of User, Business or Place entity").Optional(),
		field.String("entity_type").
			Comment("This represents the type of entity: User, Business, or Place").Optional(),
		field.String("category_id").
			Comment("This represents the ID of the category").Optional(),
	}
}

// Edges of the CategoryAssignment.
func (CategoryAssignment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("categoryAssignments").
			Unique().
			Field("entity_id"),
		edge.From("business", Business.Type).
			Ref("categoryAssignments").
			Unique().
			Field("entity_id"),
		edge.From("place", Place.Type).
			Ref("categoryAssignments").
			Unique().
			Field("entity_id"),
		edge.From("category", Category.Type).
			Ref("categoryAssignments").
			Unique().
			Field("category_id"),
	}
}
