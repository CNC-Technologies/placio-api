// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"placio-app/ent/featurerelease"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// FeatureRelease is the model entity for the FeatureRelease schema.
type FeatureRelease struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// FeatureName holds the value of the "feature_name" field.
	FeatureName string `json:"feature_name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// State holds the value of the "state" field.
	State featurerelease.State `json:"state,omitempty"`
	// ReleaseDate holds the value of the "release_date" field.
	ReleaseDate time.Time `json:"release_date,omitempty"`
	// EligibilityRules holds the value of the "eligibility_rules" field.
	EligibilityRules map[string]interface{} `json:"eligibility_rules,omitempty"`
	// DocumentationLink holds the value of the "documentation_link" field.
	DocumentationLink string `json:"documentation_link,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FeatureRelease) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case featurerelease.FieldEligibilityRules, featurerelease.FieldMetadata:
			values[i] = new([]byte)
		case featurerelease.FieldID, featurerelease.FieldFeatureName, featurerelease.FieldDescription, featurerelease.FieldState, featurerelease.FieldDocumentationLink:
			values[i] = new(sql.NullString)
		case featurerelease.FieldReleaseDate:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FeatureRelease fields.
func (fr *FeatureRelease) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case featurerelease.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				fr.ID = value.String
			}
		case featurerelease.FieldFeatureName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field feature_name", values[i])
			} else if value.Valid {
				fr.FeatureName = value.String
			}
		case featurerelease.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				fr.Description = value.String
			}
		case featurerelease.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				fr.State = featurerelease.State(value.String)
			}
		case featurerelease.FieldReleaseDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field release_date", values[i])
			} else if value.Valid {
				fr.ReleaseDate = value.Time
			}
		case featurerelease.FieldEligibilityRules:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field eligibility_rules", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &fr.EligibilityRules); err != nil {
					return fmt.Errorf("unmarshal field eligibility_rules: %w", err)
				}
			}
		case featurerelease.FieldDocumentationLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field documentation_link", values[i])
			} else if value.Valid {
				fr.DocumentationLink = value.String
			}
		case featurerelease.FieldMetadata:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field metadata", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &fr.Metadata); err != nil {
					return fmt.Errorf("unmarshal field metadata: %w", err)
				}
			}
		default:
			fr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the FeatureRelease.
// This includes values selected through modifiers, order, etc.
func (fr *FeatureRelease) Value(name string) (ent.Value, error) {
	return fr.selectValues.Get(name)
}

// Update returns a builder for updating this FeatureRelease.
// Note that you need to call FeatureRelease.Unwrap() before calling this method if this FeatureRelease
// was returned from a transaction, and the transaction was committed or rolled back.
func (fr *FeatureRelease) Update() *FeatureReleaseUpdateOne {
	return NewFeatureReleaseClient(fr.config).UpdateOne(fr)
}

// Unwrap unwraps the FeatureRelease entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fr *FeatureRelease) Unwrap() *FeatureRelease {
	_tx, ok := fr.config.driver.(*txDriver)
	if !ok {
		panic("ent: FeatureRelease is not a transactional entity")
	}
	fr.config.driver = _tx.drv
	return fr
}

// String implements the fmt.Stringer.
func (fr *FeatureRelease) String() string {
	var builder strings.Builder
	builder.WriteString("FeatureRelease(")
	builder.WriteString(fmt.Sprintf("id=%v, ", fr.ID))
	builder.WriteString("feature_name=")
	builder.WriteString(fr.FeatureName)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(fr.Description)
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(fmt.Sprintf("%v", fr.State))
	builder.WriteString(", ")
	builder.WriteString("release_date=")
	builder.WriteString(fr.ReleaseDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("eligibility_rules=")
	builder.WriteString(fmt.Sprintf("%v", fr.EligibilityRules))
	builder.WriteString(", ")
	builder.WriteString("documentation_link=")
	builder.WriteString(fr.DocumentationLink)
	builder.WriteString(", ")
	builder.WriteString("metadata=")
	builder.WriteString(fmt.Sprintf("%v", fr.Metadata))
	builder.WriteByte(')')
	return builder.String()
}

// FeatureReleases is a parsable slice of FeatureRelease.
type FeatureReleases []*FeatureRelease