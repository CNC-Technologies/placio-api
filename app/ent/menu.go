// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"placio-app/ent/menu"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Menu is the model entity for the Menu schema.
type Menu struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt string `json:"deleted_at,omitempty"`
	// IsDeleted holds the value of the "is_deleted" field.
	IsDeleted bool `json:"is_deleted,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Options holds the value of the "options" field.
	Options string `json:"options,omitempty"`
	// FoodType holds the value of the "foodType" field.
	FoodType menu.FoodType `json:"foodType,omitempty"`
	// MenuItemType holds the value of the "menuItemType" field.
	MenuItemType menu.MenuItemType `json:"menuItemType,omitempty"`
	// DrinkType holds the value of the "drinkType" field.
	DrinkType menu.DrinkType `json:"drinkType,omitempty"`
	// DietaryType holds the value of the "dietaryType" field.
	DietaryType menu.DietaryType `json:"dietaryType,omitempty"`
	// IsAvailable holds the value of the "is_available" field.
	IsAvailable bool `json:"is_available,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MenuQuery when eager-loading is set.
	Edges        MenuEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MenuEdges holds the relations/edges for other nodes in the graph.
type MenuEdges struct {
	// Place holds the value of the place edge.
	Place []*Place `json:"place,omitempty"`
	// Categories holds the value of the categories edge.
	Categories []*Category `json:"categories,omitempty"`
	// MenuItems holds the value of the menu_items edge.
	MenuItems []*MenuItem `json:"menu_items,omitempty"`
	// Media holds the value of the media edge.
	Media []*Media `json:"media,omitempty"`
	// CreatedBy holds the value of the created_by edge.
	CreatedBy []*User `json:"created_by,omitempty"`
	// UpdatedBy holds the value of the updated_by edge.
	UpdatedBy []*User `json:"updated_by,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// PlaceOrErr returns the Place value or an error if the edge
// was not loaded in eager-loading.
func (e MenuEdges) PlaceOrErr() ([]*Place, error) {
	if e.loadedTypes[0] {
		return e.Place, nil
	}
	return nil, &NotLoadedError{edge: "place"}
}

// CategoriesOrErr returns the Categories value or an error if the edge
// was not loaded in eager-loading.
func (e MenuEdges) CategoriesOrErr() ([]*Category, error) {
	if e.loadedTypes[1] {
		return e.Categories, nil
	}
	return nil, &NotLoadedError{edge: "categories"}
}

// MenuItemsOrErr returns the MenuItems value or an error if the edge
// was not loaded in eager-loading.
func (e MenuEdges) MenuItemsOrErr() ([]*MenuItem, error) {
	if e.loadedTypes[2] {
		return e.MenuItems, nil
	}
	return nil, &NotLoadedError{edge: "menu_items"}
}

// MediaOrErr returns the Media value or an error if the edge
// was not loaded in eager-loading.
func (e MenuEdges) MediaOrErr() ([]*Media, error) {
	if e.loadedTypes[3] {
		return e.Media, nil
	}
	return nil, &NotLoadedError{edge: "media"}
}

// CreatedByOrErr returns the CreatedBy value or an error if the edge
// was not loaded in eager-loading.
func (e MenuEdges) CreatedByOrErr() ([]*User, error) {
	if e.loadedTypes[4] {
		return e.CreatedBy, nil
	}
	return nil, &NotLoadedError{edge: "created_by"}
}

// UpdatedByOrErr returns the UpdatedBy value or an error if the edge
// was not loaded in eager-loading.
func (e MenuEdges) UpdatedByOrErr() ([]*User, error) {
	if e.loadedTypes[5] {
		return e.UpdatedBy, nil
	}
	return nil, &NotLoadedError{edge: "updated_by"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Menu) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case menu.FieldIsDeleted, menu.FieldIsAvailable:
			values[i] = new(sql.NullBool)
		case menu.FieldID, menu.FieldName, menu.FieldDeletedAt, menu.FieldDescription, menu.FieldOptions, menu.FieldFoodType, menu.FieldMenuItemType, menu.FieldDrinkType, menu.FieldDietaryType:
			values[i] = new(sql.NullString)
		case menu.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Menu fields.
func (m *Menu) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case menu.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				m.ID = value.String
			}
		case menu.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		case menu.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				m.DeletedAt = value.String
			}
		case menu.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				m.IsDeleted = value.Bool
			}
		case menu.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				m.Description = value.String
			}
		case menu.FieldOptions:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field options", values[i])
			} else if value.Valid {
				m.Options = value.String
			}
		case menu.FieldFoodType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field foodType", values[i])
			} else if value.Valid {
				m.FoodType = menu.FoodType(value.String)
			}
		case menu.FieldMenuItemType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field menuItemType", values[i])
			} else if value.Valid {
				m.MenuItemType = menu.MenuItemType(value.String)
			}
		case menu.FieldDrinkType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field drinkType", values[i])
			} else if value.Valid {
				m.DrinkType = menu.DrinkType(value.String)
			}
		case menu.FieldDietaryType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dietaryType", values[i])
			} else if value.Valid {
				m.DietaryType = menu.DietaryType(value.String)
			}
		case menu.FieldIsAvailable:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_available", values[i])
			} else if value.Valid {
				m.IsAvailable = value.Bool
			}
		case menu.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				m.UpdatedAt = value.Time
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Menu.
// This includes values selected through modifiers, order, etc.
func (m *Menu) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryPlace queries the "place" edge of the Menu entity.
func (m *Menu) QueryPlace() *PlaceQuery {
	return NewMenuClient(m.config).QueryPlace(m)
}

// QueryCategories queries the "categories" edge of the Menu entity.
func (m *Menu) QueryCategories() *CategoryQuery {
	return NewMenuClient(m.config).QueryCategories(m)
}

// QueryMenuItems queries the "menu_items" edge of the Menu entity.
func (m *Menu) QueryMenuItems() *MenuItemQuery {
	return NewMenuClient(m.config).QueryMenuItems(m)
}

// QueryMedia queries the "media" edge of the Menu entity.
func (m *Menu) QueryMedia() *MediaQuery {
	return NewMenuClient(m.config).QueryMedia(m)
}

// QueryCreatedBy queries the "created_by" edge of the Menu entity.
func (m *Menu) QueryCreatedBy() *UserQuery {
	return NewMenuClient(m.config).QueryCreatedBy(m)
}

// QueryUpdatedBy queries the "updated_by" edge of the Menu entity.
func (m *Menu) QueryUpdatedBy() *UserQuery {
	return NewMenuClient(m.config).QueryUpdatedBy(m)
}

// Update returns a builder for updating this Menu.
// Note that you need to call Menu.Unwrap() before calling this method if this Menu
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Menu) Update() *MenuUpdateOne {
	return NewMenuClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Menu entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Menu) Unwrap() *Menu {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Menu is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Menu) String() string {
	var builder strings.Builder
	builder.WriteString("Menu(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("name=")
	builder.WriteString(m.Name)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(m.DeletedAt)
	builder.WriteString(", ")
	builder.WriteString("is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", m.IsDeleted))
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(m.Description)
	builder.WriteString(", ")
	builder.WriteString("options=")
	builder.WriteString(m.Options)
	builder.WriteString(", ")
	builder.WriteString("foodType=")
	builder.WriteString(fmt.Sprintf("%v", m.FoodType))
	builder.WriteString(", ")
	builder.WriteString("menuItemType=")
	builder.WriteString(fmt.Sprintf("%v", m.MenuItemType))
	builder.WriteString(", ")
	builder.WriteString("drinkType=")
	builder.WriteString(fmt.Sprintf("%v", m.DrinkType))
	builder.WriteString(", ")
	builder.WriteString("dietaryType=")
	builder.WriteString(fmt.Sprintf("%v", m.DietaryType))
	builder.WriteString(", ")
	builder.WriteString("is_available=")
	builder.WriteString(fmt.Sprintf("%v", m.IsAvailable))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(m.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Menus is a parsable slice of Menu.
type Menus []*Menu
