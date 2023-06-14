// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"
	"placio-app/ent"
)

// The AccountSettingsFunc type is an adapter to allow the use of ordinary
// function as AccountSettings mutator.
type AccountSettingsFunc func(context.Context, *ent.AccountSettingsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AccountSettingsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AccountSettingsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AccountSettingsMutation", m)
}

// The BookingFunc type is an adapter to allow the use of ordinary
// function as Booking mutator.
type BookingFunc func(context.Context, *ent.BookingMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BookingFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.BookingMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BookingMutation", m)
}

// The BusinessFunc type is an adapter to allow the use of ordinary
// function as Business mutator.
type BusinessFunc func(context.Context, *ent.BusinessMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BusinessFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.BusinessMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BusinessMutation", m)
}

// The BusinessFollowBusinessFunc type is an adapter to allow the use of ordinary
// function as BusinessFollowBusiness mutator.
type BusinessFollowBusinessFunc func(context.Context, *ent.BusinessFollowBusinessMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BusinessFollowBusinessFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.BusinessFollowBusinessMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BusinessFollowBusinessMutation", m)
}

// The BusinessFollowUserFunc type is an adapter to allow the use of ordinary
// function as BusinessFollowUser mutator.
type BusinessFollowUserFunc func(context.Context, *ent.BusinessFollowUserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BusinessFollowUserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.BusinessFollowUserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BusinessFollowUserMutation", m)
}

// The ChatFunc type is an adapter to allow the use of ordinary
// function as Chat mutator.
type ChatFunc func(context.Context, *ent.ChatMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ChatFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ChatMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ChatMutation", m)
}

// The CommentFunc type is an adapter to allow the use of ordinary
// function as Comment mutator.
type CommentFunc func(context.Context, *ent.CommentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CommentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CommentMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CommentMutation", m)
}

// The LikeFunc type is an adapter to allow the use of ordinary
// function as Like mutator.
type LikeFunc func(context.Context, *ent.LikeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f LikeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.LikeMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.LikeMutation", m)
}

// The MediaFunc type is an adapter to allow the use of ordinary
// function as Media mutator.
type MediaFunc func(context.Context, *ent.MediaMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MediaFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.MediaMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MediaMutation", m)
}

// The OrderFunc type is an adapter to allow the use of ordinary
// function as Order mutator.
type OrderFunc func(context.Context, *ent.OrderMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrderFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.OrderMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrderMutation", m)
}

// The PaymentFunc type is an adapter to allow the use of ordinary
// function as Payment mutator.
type PaymentFunc func(context.Context, *ent.PaymentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PaymentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PaymentMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PaymentMutation", m)
}

// The PostFunc type is an adapter to allow the use of ordinary
// function as Post mutator.
type PostFunc func(context.Context, *ent.PostMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PostFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PostMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PostMutation", m)
}

// The RatingFunc type is an adapter to allow the use of ordinary
// function as Rating mutator.
type RatingFunc func(context.Context, *ent.RatingMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RatingFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.RatingMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RatingMutation", m)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
}

// The UserBusinessFunc type is an adapter to allow the use of ordinary
// function as UserBusiness mutator.
type UserBusinessFunc func(context.Context, *ent.UserBusinessMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserBusinessFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserBusinessMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserBusinessMutation", m)
}

// The UserFollowBusinessFunc type is an adapter to allow the use of ordinary
// function as UserFollowBusiness mutator.
type UserFollowBusinessFunc func(context.Context, *ent.UserFollowBusinessMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFollowBusinessFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserFollowBusinessMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserFollowBusinessMutation", m)
}

// The UserFollowUserFunc type is an adapter to allow the use of ordinary
// function as UserFollowUser mutator.
type UserFollowUserFunc func(context.Context, *ent.UserFollowUserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFollowUserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserFollowUserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserFollowUserMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
