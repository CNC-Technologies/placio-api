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

// The AccountWalletFunc type is an adapter to allow the use of ordinary
// function as AccountWallet mutator.
type AccountWalletFunc func(context.Context, *ent.AccountWalletMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AccountWalletFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AccountWalletMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AccountWalletMutation", m)
}

// The AmenityFunc type is an adapter to allow the use of ordinary
// function as Amenity mutator.
type AmenityFunc func(context.Context, *ent.AmenityMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AmenityFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AmenityMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AmenityMutation", m)
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

// The BusinessFollowEventFunc type is an adapter to allow the use of ordinary
// function as BusinessFollowEvent mutator.
type BusinessFollowEventFunc func(context.Context, *ent.BusinessFollowEventMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BusinessFollowEventFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.BusinessFollowEventMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BusinessFollowEventMutation", m)
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

// The CategoryFunc type is an adapter to allow the use of ordinary
// function as Category mutator.
type CategoryFunc func(context.Context, *ent.CategoryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CategoryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CategoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CategoryMutation", m)
}

// The CategoryAssignmentFunc type is an adapter to allow the use of ordinary
// function as CategoryAssignment mutator.
type CategoryAssignmentFunc func(context.Context, *ent.CategoryAssignmentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CategoryAssignmentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CategoryAssignmentMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CategoryAssignmentMutation", m)
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

// The CustomBlockFunc type is an adapter to allow the use of ordinary
// function as CustomBlock mutator.
type CustomBlockFunc func(context.Context, *ent.CustomBlockMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CustomBlockFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CustomBlockMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CustomBlockMutation", m)
}

// The EventFunc type is an adapter to allow the use of ordinary
// function as Event mutator.
type EventFunc func(context.Context, *ent.EventMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EventFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.EventMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EventMutation", m)
}

// The EventOrganizerFunc type is an adapter to allow the use of ordinary
// function as EventOrganizer mutator.
type EventOrganizerFunc func(context.Context, *ent.EventOrganizerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EventOrganizerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.EventOrganizerMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EventOrganizerMutation", m)
}

// The FAQFunc type is an adapter to allow the use of ordinary
// function as FAQ mutator.
type FAQFunc func(context.Context, *ent.FAQMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FAQFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.FAQMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FAQMutation", m)
}

// The FeatureReleaseFunc type is an adapter to allow the use of ordinary
// function as FeatureRelease mutator.
type FeatureReleaseFunc func(context.Context, *ent.FeatureReleaseMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FeatureReleaseFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.FeatureReleaseMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FeatureReleaseMutation", m)
}

// The FitnessFunc type is an adapter to allow the use of ordinary
// function as Fitness mutator.
type FitnessFunc func(context.Context, *ent.FitnessMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FitnessFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.FitnessMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FitnessMutation", m)
}

// The HelpFunc type is an adapter to allow the use of ordinary
// function as Help mutator.
type HelpFunc func(context.Context, *ent.HelpMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f HelpFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.HelpMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.HelpMutation", m)
}

// The InventoryAttributeFunc type is an adapter to allow the use of ordinary
// function as InventoryAttribute mutator.
type InventoryAttributeFunc func(context.Context, *ent.InventoryAttributeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f InventoryAttributeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.InventoryAttributeMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.InventoryAttributeMutation", m)
}

// The InventoryTypeFunc type is an adapter to allow the use of ordinary
// function as InventoryType mutator.
type InventoryTypeFunc func(context.Context, *ent.InventoryTypeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f InventoryTypeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.InventoryTypeMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.InventoryTypeMutation", m)
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

// The MenuFunc type is an adapter to allow the use of ordinary
// function as Menu mutator.
type MenuFunc func(context.Context, *ent.MenuMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MenuFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.MenuMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MenuMutation", m)
}

// The MenuItemFunc type is an adapter to allow the use of ordinary
// function as MenuItem mutator.
type MenuItemFunc func(context.Context, *ent.MenuItemMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MenuItemFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.MenuItemMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MenuItemMutation", m)
}

// The NotificationFunc type is an adapter to allow the use of ordinary
// function as Notification mutator.
type NotificationFunc func(context.Context, *ent.NotificationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f NotificationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.NotificationMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.NotificationMutation", m)
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

// The OrderItemFunc type is an adapter to allow the use of ordinary
// function as OrderItem mutator.
type OrderItemFunc func(context.Context, *ent.OrderItemMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrderItemFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.OrderItemMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrderItemMutation", m)
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

// The PermissionFunc type is an adapter to allow the use of ordinary
// function as Permission mutator.
type PermissionFunc func(context.Context, *ent.PermissionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PermissionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PermissionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PermissionMutation", m)
}

// The PlaceFunc type is an adapter to allow the use of ordinary
// function as Place mutator.
type PlaceFunc func(context.Context, *ent.PlaceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PlaceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PlaceMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PlaceMutation", m)
}

// The PlaceInventoryFunc type is an adapter to allow the use of ordinary
// function as PlaceInventory mutator.
type PlaceInventoryFunc func(context.Context, *ent.PlaceInventoryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PlaceInventoryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PlaceInventoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PlaceInventoryMutation", m)
}

// The PlaceInventoryAttributeFunc type is an adapter to allow the use of ordinary
// function as PlaceInventoryAttribute mutator.
type PlaceInventoryAttributeFunc func(context.Context, *ent.PlaceInventoryAttributeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PlaceInventoryAttributeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PlaceInventoryAttributeMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PlaceInventoryAttributeMutation", m)
}

// The PlaceTableFunc type is an adapter to allow the use of ordinary
// function as PlaceTable mutator.
type PlaceTableFunc func(context.Context, *ent.PlaceTableMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PlaceTableFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PlaceTableMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PlaceTableMutation", m)
}

// The PlanFunc type is an adapter to allow the use of ordinary
// function as Plan mutator.
type PlanFunc func(context.Context, *ent.PlanMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PlanFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PlanMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PlanMutation", m)
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

// The PriceFunc type is an adapter to allow the use of ordinary
// function as Price mutator.
type PriceFunc func(context.Context, *ent.PriceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PriceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PriceMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PriceMutation", m)
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

// The ReactionFunc type is an adapter to allow the use of ordinary
// function as Reaction mutator.
type ReactionFunc func(context.Context, *ent.ReactionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ReactionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ReactionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ReactionMutation", m)
}

// The ReservationFunc type is an adapter to allow the use of ordinary
// function as Reservation mutator.
type ReservationFunc func(context.Context, *ent.ReservationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ReservationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ReservationMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ReservationMutation", m)
}

// The ReservationBlockFunc type is an adapter to allow the use of ordinary
// function as ReservationBlock mutator.
type ReservationBlockFunc func(context.Context, *ent.ReservationBlockMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ReservationBlockFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ReservationBlockMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ReservationBlockMutation", m)
}

// The ResourseFunc type is an adapter to allow the use of ordinary
// function as Resourse mutator.
type ResourseFunc func(context.Context, *ent.ResourseMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ResourseFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ResourseMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ResourseMutation", m)
}

// The ReviewFunc type is an adapter to allow the use of ordinary
// function as Review mutator.
type ReviewFunc func(context.Context, *ent.ReviewMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ReviewFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ReviewMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ReviewMutation", m)
}

// The RoomFunc type is an adapter to allow the use of ordinary
// function as Room mutator.
type RoomFunc func(context.Context, *ent.RoomMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RoomFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.RoomMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RoomMutation", m)
}

// The RoomCategoryFunc type is an adapter to allow the use of ordinary
// function as RoomCategory mutator.
type RoomCategoryFunc func(context.Context, *ent.RoomCategoryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RoomCategoryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.RoomCategoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RoomCategoryMutation", m)
}

// The StaffFunc type is an adapter to allow the use of ordinary
// function as Staff mutator.
type StaffFunc func(context.Context, *ent.StaffMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f StaffFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.StaffMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.StaffMutation", m)
}

// The SubscriptionFunc type is an adapter to allow the use of ordinary
// function as Subscription mutator.
type SubscriptionFunc func(context.Context, *ent.SubscriptionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SubscriptionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.SubscriptionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SubscriptionMutation", m)
}

// The TemplateFunc type is an adapter to allow the use of ordinary
// function as Template mutator.
type TemplateFunc func(context.Context, *ent.TemplateMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TemplateFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.TemplateMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TemplateMutation", m)
}

// The TicketFunc type is an adapter to allow the use of ordinary
// function as Ticket mutator.
type TicketFunc func(context.Context, *ent.TicketMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TicketFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.TicketMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TicketMutation", m)
}

// The TicketOptionFunc type is an adapter to allow the use of ordinary
// function as TicketOption mutator.
type TicketOptionFunc func(context.Context, *ent.TicketOptionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TicketOptionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.TicketOptionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TicketOptionMutation", m)
}

// The TrainerFunc type is an adapter to allow the use of ordinary
// function as Trainer mutator.
type TrainerFunc func(context.Context, *ent.TrainerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TrainerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.TrainerMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TrainerMutation", m)
}

// The TransactionHistoryFunc type is an adapter to allow the use of ordinary
// function as TransactionHistory mutator.
type TransactionHistoryFunc func(context.Context, *ent.TransactionHistoryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TransactionHistoryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.TransactionHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TransactionHistoryMutation", m)
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

// The UserFollowEventFunc type is an adapter to allow the use of ordinary
// function as UserFollowEvent mutator.
type UserFollowEventFunc func(context.Context, *ent.UserFollowEventMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFollowEventFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserFollowEventMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserFollowEventMutation", m)
}

// The UserFollowPlaceFunc type is an adapter to allow the use of ordinary
// function as UserFollowPlace mutator.
type UserFollowPlaceFunc func(context.Context, *ent.UserFollowPlaceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFollowPlaceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserFollowPlaceMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserFollowPlaceMutation", m)
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

// The UserLikePlaceFunc type is an adapter to allow the use of ordinary
// function as UserLikePlace mutator.
type UserLikePlaceFunc func(context.Context, *ent.UserLikePlaceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserLikePlaceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserLikePlaceMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserLikePlaceMutation", m)
}

// The WebsiteFunc type is an adapter to allow the use of ordinary
// function as Website mutator.
type WebsiteFunc func(context.Context, *ent.WebsiteMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f WebsiteFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.WebsiteMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.WebsiteMutation", m)
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
