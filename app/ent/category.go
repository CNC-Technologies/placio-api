// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"placio-app/ent/category"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Category is the model entity for the Category schema.
type Category struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Image holds the value of the "image" field.
	Image string `json:"image,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Icon holds the value of the "icon" field.
	Icon string `json:"icon,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// ParentID holds the value of the "parent_id" field.
	ParentID string `json:"parent_id,omitempty"`
	// ParentName holds the value of the "parent_name" field.
	ParentName string `json:"parent_name,omitempty"`
	// ParentImage holds the value of the "parent_image" field.
	ParentImage string `json:"parent_image,omitempty"`
	// ParentDescription holds the value of the "parent_description" field.
	ParentDescription string `json:"parent_description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CategoryQuery when eager-loading is set.
	Edges                  CategoryEdges `json:"edges"`
	business_categories    *string
	event_event_categories *string
	place_categories       *string
	post_categories        *string
	user_categories        *string
	selectValues           sql.SelectValues
}

// CategoryEdges holds the relations/edges for other nodes in the graph.
type CategoryEdges struct {
	// CategoryAssignments holds the value of the categoryAssignments edge.
	CategoryAssignments []*CategoryAssignment `json:"categoryAssignments,omitempty"`
	// PlaceInventories holds the value of the place_inventories edge.
	PlaceInventories []*PlaceInventory `json:"place_inventories,omitempty"`
	// Media holds the value of the media edge.
	Media []*Media `json:"media,omitempty"`
	// Menus holds the value of the menus edge.
	Menus []*Menu `json:"menus,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// CategoryAssignmentsOrErr returns the CategoryAssignments value or an error if the edge
// was not loaded in eager-loading.
func (e CategoryEdges) CategoryAssignmentsOrErr() ([]*CategoryAssignment, error) {
	if e.loadedTypes[0] {
		return e.CategoryAssignments, nil
	}
	return nil, &NotLoadedError{edge: "categoryAssignments"}
}

// PlaceInventoriesOrErr returns the PlaceInventories value or an error if the edge
// was not loaded in eager-loading.
func (e CategoryEdges) PlaceInventoriesOrErr() ([]*PlaceInventory, error) {
	if e.loadedTypes[1] {
		return e.PlaceInventories, nil
	}
	return nil, &NotLoadedError{edge: "place_inventories"}
}

// MediaOrErr returns the Media value or an error if the edge
// was not loaded in eager-loading.
func (e CategoryEdges) MediaOrErr() ([]*Media, error) {
	if e.loadedTypes[2] {
		return e.Media, nil
	}
	return nil, &NotLoadedError{edge: "media"}
}

// MenusOrErr returns the Menus value or an error if the edge
// was not loaded in eager-loading.
func (e CategoryEdges) MenusOrErr() ([]*Menu, error) {
	if e.loadedTypes[3] {
		return e.Menus, nil
	}
	return nil, &NotLoadedError{edge: "menus"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Category) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case category.FieldID, category.FieldName, category.FieldImage, category.FieldDescription, category.FieldIcon, category.FieldType, category.FieldParentID, category.FieldParentName, category.FieldParentImage, category.FieldParentDescription:
			values[i] = new(sql.NullString)
		case category.ForeignKeys[0]: // business_categories
			values[i] = new(sql.NullString)
		case category.ForeignKeys[1]: // event_event_categories
			values[i] = new(sql.NullString)
		case category.ForeignKeys[2]: // place_categories
			values[i] = new(sql.NullString)
		case category.ForeignKeys[3]: // post_categories
			values[i] = new(sql.NullString)
		case category.ForeignKeys[4]: // user_categories
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Category fields.
func (c *Category) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case category.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				c.ID = value.String
			}
		case category.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case category.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[i])
			} else if value.Valid {
				c.Image = value.String
			}
		case category.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				c.Description = value.String
			}
		case category.FieldIcon:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field icon", values[i])
			} else if value.Valid {
				c.Icon = value.String
			}
		case category.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				c.Type = value.String
			}
		case category.FieldParentID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				c.ParentID = value.String
			}
		case category.FieldParentName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field parent_name", values[i])
			} else if value.Valid {
				c.ParentName = value.String
			}
		case category.FieldParentImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field parent_image", values[i])
			} else if value.Valid {
				c.ParentImage = value.String
			}
		case category.FieldParentDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field parent_description", values[i])
			} else if value.Valid {
				c.ParentDescription = value.String
			}
		case category.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field business_categories", values[i])
			} else if value.Valid {
				c.business_categories = new(string)
				*c.business_categories = value.String
			}
		case category.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field event_event_categories", values[i])
			} else if value.Valid {
				c.event_event_categories = new(string)
				*c.event_event_categories = value.String
			}
		case category.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field place_categories", values[i])
			} else if value.Valid {
				c.place_categories = new(string)
				*c.place_categories = value.String
			}
		case category.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field post_categories", values[i])
			} else if value.Valid {
				c.post_categories = new(string)
				*c.post_categories = value.String
			}
		case category.ForeignKeys[4]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_categories", values[i])
			} else if value.Valid {
				c.user_categories = new(string)
				*c.user_categories = value.String
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Category.
// This includes values selected through modifiers, order, etc.
func (c *Category) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryCategoryAssignments queries the "categoryAssignments" edge of the Category entity.
func (c *Category) QueryCategoryAssignments() *CategoryAssignmentQuery {
	return NewCategoryClient(c.config).QueryCategoryAssignments(c)
}

// QueryPlaceInventories queries the "place_inventories" edge of the Category entity.
func (c *Category) QueryPlaceInventories() *PlaceInventoryQuery {
	return NewCategoryClient(c.config).QueryPlaceInventories(c)
}

// QueryMedia queries the "media" edge of the Category entity.
func (c *Category) QueryMedia() *MediaQuery {
	return NewCategoryClient(c.config).QueryMedia(c)
}

// QueryMenus queries the "menus" edge of the Category entity.
func (c *Category) QueryMenus() *MenuQuery {
	return NewCategoryClient(c.config).QueryMenus(c)
}

// Update returns a builder for updating this Category.
// Note that you need to call Category.Unwrap() before calling this method if this Category
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Category) Update() *CategoryUpdateOne {
	return NewCategoryClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Category entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Category) Unwrap() *Category {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Category is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Category) String() string {
	var builder strings.Builder
	builder.WriteString("Category(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("image=")
	builder.WriteString(c.Image)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(c.Description)
	builder.WriteString(", ")
	builder.WriteString("icon=")
	builder.WriteString(c.Icon)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(c.Type)
	builder.WriteString(", ")
	builder.WriteString("parent_id=")
	builder.WriteString(c.ParentID)
	builder.WriteString(", ")
	builder.WriteString("parent_name=")
	builder.WriteString(c.ParentName)
	builder.WriteString(", ")
	builder.WriteString("parent_image=")
	builder.WriteString(c.ParentImage)
	builder.WriteString(", ")
	builder.WriteString("parent_description=")
	builder.WriteString(c.ParentDescription)
	builder.WriteByte(')')
	return builder.String()
}

// Categories is a parsable slice of Category.
type Categories []*Category
