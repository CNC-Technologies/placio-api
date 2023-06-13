// Code generated by ent, DO NOT EDIT.

package user

import (
	"placio-app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldID, id))
}

// Auth0ID applies equality check predicate on the "auth0_id" field. It's identical to Auth0IDEQ.
func Auth0ID(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAuth0ID, v))
}

// Auth0IDEQ applies the EQ predicate on the "auth0_id" field.
func Auth0IDEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAuth0ID, v))
}

// Auth0IDNEQ applies the NEQ predicate on the "auth0_id" field.
func Auth0IDNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldAuth0ID, v))
}

// Auth0IDIn applies the In predicate on the "auth0_id" field.
func Auth0IDIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldAuth0ID, vs...))
}

// Auth0IDNotIn applies the NotIn predicate on the "auth0_id" field.
func Auth0IDNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldAuth0ID, vs...))
}

// Auth0IDGT applies the GT predicate on the "auth0_id" field.
func Auth0IDGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldAuth0ID, v))
}

// Auth0IDGTE applies the GTE predicate on the "auth0_id" field.
func Auth0IDGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldAuth0ID, v))
}

// Auth0IDLT applies the LT predicate on the "auth0_id" field.
func Auth0IDLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldAuth0ID, v))
}

// Auth0IDLTE applies the LTE predicate on the "auth0_id" field.
func Auth0IDLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldAuth0ID, v))
}

// Auth0IDContains applies the Contains predicate on the "auth0_id" field.
func Auth0IDContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldAuth0ID, v))
}

// Auth0IDHasPrefix applies the HasPrefix predicate on the "auth0_id" field.
func Auth0IDHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldAuth0ID, v))
}

// Auth0IDHasSuffix applies the HasSuffix predicate on the "auth0_id" field.
func Auth0IDHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldAuth0ID, v))
}

// Auth0IDEqualFold applies the EqualFold predicate on the "auth0_id" field.
func Auth0IDEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldAuth0ID, v))
}

// Auth0IDContainsFold applies the ContainsFold predicate on the "auth0_id" field.
func Auth0IDContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldAuth0ID, v))
}

// Auth0DataIsNil applies the IsNil predicate on the "auth0_data" field.
func Auth0DataIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldAuth0Data))
}

// Auth0DataNotNil applies the NotNil predicate on the "auth0_data" field.
func Auth0DataNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldAuth0Data))
}

// HasUserBusinesses applies the HasEdge predicate on the "userBusinesses" edge.
func HasUserBusinesses() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserBusinessesTable, UserBusinessesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserBusinessesWith applies the HasEdge predicate on the "userBusinesses" edge with a given conditions (other predicates).
func HasUserBusinessesWith(preds ...predicate.UserBusiness) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newUserBusinessesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasComments applies the HasEdge predicate on the "comments" edge.
func HasComments() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CommentsTable, CommentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCommentsWith applies the HasEdge predicate on the "comments" edge with a given conditions (other predicates).
func HasCommentsWith(preds ...predicate.Comment) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newCommentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLikes applies the HasEdge predicate on the "likes" edge.
func HasLikes() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LikesTable, LikesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLikesWith applies the HasEdge predicate on the "likes" edge with a given conditions (other predicates).
func HasLikesWith(preds ...predicate.Like) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newLikesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPosts applies the HasEdge predicate on the "posts" edge.
func HasPosts() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PostsTable, PostsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPostsWith applies the HasEdge predicate on the "posts" edge with a given conditions (other predicates).
func HasPostsWith(preds ...predicate.Post) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newPostsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
