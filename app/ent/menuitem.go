// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"placio-app/ent/menuitem"
	"placio-app/ent/placeinventory"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// MenuItem is the model entity for the MenuItem schema.
type MenuItem struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Price holds the value of the "price" field.
	Price float64 `json:"price,omitempty"`
	// PreparationTime holds the value of the "preparation_time" field.
	PreparationTime int `json:"preparation_time,omitempty"`
	// Options holds the value of the "options" field.
	Options []string `json:"options,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt string `json:"deleted_at,omitempty"`
	// IsDeleted holds the value of the "is_deleted" field.
	IsDeleted bool `json:"is_deleted,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MenuItemQuery when eager-loading is set.
	Edges        MenuItemEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MenuItemEdges holds the relations/edges for other nodes in the graph.
type MenuItemEdges struct {
	// Menu holds the value of the menu edge.
	Menu []*Menu `json:"menu,omitempty"`
	// Inventory holds the value of the inventory edge.
	Inventory *PlaceInventory `json:"inventory,omitempty"`
	// Media holds the value of the media edge.
	Media []*Media `json:"media,omitempty"`
	// OrderItems holds the value of the order_items edge.
	OrderItems []*OrderItem `json:"order_items,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// MenuOrErr returns the Menu value or an error if the edge
// was not loaded in eager-loading.
func (e MenuItemEdges) MenuOrErr() ([]*Menu, error) {
	if e.loadedTypes[0] {
		return e.Menu, nil
	}
	return nil, &NotLoadedError{edge: "menu"}
}

// InventoryOrErr returns the Inventory value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MenuItemEdges) InventoryOrErr() (*PlaceInventory, error) {
	if e.loadedTypes[1] {
		if e.Inventory == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: placeinventory.Label}
		}
		return e.Inventory, nil
	}
	return nil, &NotLoadedError{edge: "inventory"}
}

// MediaOrErr returns the Media value or an error if the edge
// was not loaded in eager-loading.
func (e MenuItemEdges) MediaOrErr() ([]*Media, error) {
	if e.loadedTypes[2] {
		return e.Media, nil
	}
	return nil, &NotLoadedError{edge: "media"}
}

// OrderItemsOrErr returns the OrderItems value or an error if the edge
// was not loaded in eager-loading.
func (e MenuItemEdges) OrderItemsOrErr() ([]*OrderItem, error) {
	if e.loadedTypes[3] {
		return e.OrderItems, nil
	}
	return nil, &NotLoadedError{edge: "order_items"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MenuItem) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case menuitem.FieldOptions:
			values[i] = new([]byte)
		case menuitem.FieldIsDeleted:
			values[i] = new(sql.NullBool)
		case menuitem.FieldPrice:
			values[i] = new(sql.NullFloat64)
		case menuitem.FieldPreparationTime:
			values[i] = new(sql.NullInt64)
		case menuitem.FieldID, menuitem.FieldName, menuitem.FieldDescription, menuitem.FieldDeletedAt:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MenuItem fields.
func (mi *MenuItem) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case menuitem.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				mi.ID = value.String
			}
		case menuitem.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				mi.Name = value.String
			}
		case menuitem.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				mi.Description = value.String
			}
		case menuitem.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				mi.Price = value.Float64
			}
		case menuitem.FieldPreparationTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field preparation_time", values[i])
			} else if value.Valid {
				mi.PreparationTime = int(value.Int64)
			}
		case menuitem.FieldOptions:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field options", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &mi.Options); err != nil {
					return fmt.Errorf("unmarshal field options: %w", err)
				}
			}
		case menuitem.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				mi.DeletedAt = value.String
			}
		case menuitem.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				mi.IsDeleted = value.Bool
			}
		default:
			mi.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MenuItem.
// This includes values selected through modifiers, order, etc.
func (mi *MenuItem) Value(name string) (ent.Value, error) {
	return mi.selectValues.Get(name)
}

// QueryMenu queries the "menu" edge of the MenuItem entity.
func (mi *MenuItem) QueryMenu() *MenuQuery {
	return NewMenuItemClient(mi.config).QueryMenu(mi)
}

// QueryInventory queries the "inventory" edge of the MenuItem entity.
func (mi *MenuItem) QueryInventory() *PlaceInventoryQuery {
	return NewMenuItemClient(mi.config).QueryInventory(mi)
}

// QueryMedia queries the "media" edge of the MenuItem entity.
func (mi *MenuItem) QueryMedia() *MediaQuery {
	return NewMenuItemClient(mi.config).QueryMedia(mi)
}

// QueryOrderItems queries the "order_items" edge of the MenuItem entity.
func (mi *MenuItem) QueryOrderItems() *OrderItemQuery {
	return NewMenuItemClient(mi.config).QueryOrderItems(mi)
}

// Update returns a builder for updating this MenuItem.
// Note that you need to call MenuItem.Unwrap() before calling this method if this MenuItem
// was returned from a transaction, and the transaction was committed or rolled back.
func (mi *MenuItem) Update() *MenuItemUpdateOne {
	return NewMenuItemClient(mi.config).UpdateOne(mi)
}

// Unwrap unwraps the MenuItem entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mi *MenuItem) Unwrap() *MenuItem {
	_tx, ok := mi.config.driver.(*txDriver)
	if !ok {
		panic("ent: MenuItem is not a transactional entity")
	}
	mi.config.driver = _tx.drv
	return mi
}

// String implements the fmt.Stringer.
func (mi *MenuItem) String() string {
	var builder strings.Builder
	builder.WriteString("MenuItem(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mi.ID))
	builder.WriteString("name=")
	builder.WriteString(mi.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(mi.Description)
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", mi.Price))
	builder.WriteString(", ")
	builder.WriteString("preparation_time=")
	builder.WriteString(fmt.Sprintf("%v", mi.PreparationTime))
	builder.WriteString(", ")
	builder.WriteString("options=")
	builder.WriteString(fmt.Sprintf("%v", mi.Options))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(mi.DeletedAt)
	builder.WriteString(", ")
	builder.WriteString("is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", mi.IsDeleted))
	builder.WriteByte(')')
	return builder.String()
}

// MenuItems is a parsable slice of MenuItem.
type MenuItems []*MenuItem
