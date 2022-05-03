// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/subscribe-manager/pkg/db/ent/emailsubscriber"
	"github.com/google/uuid"
)

// EmailSubscriber is the model entity for the EmailSubscriber schema.
type EmailSubscriber struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// EmailAddress holds the value of the "email_address" field.
	EmailAddress string `json:"email_address,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EmailSubscriber) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case emailsubscriber.FieldCreatedAt, emailsubscriber.FieldUpdatedAt, emailsubscriber.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case emailsubscriber.FieldEmailAddress:
			values[i] = new(sql.NullString)
		case emailsubscriber.FieldID, emailsubscriber.FieldAppID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type EmailSubscriber", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EmailSubscriber fields.
func (es *EmailSubscriber) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case emailsubscriber.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				es.ID = *value
			}
		case emailsubscriber.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				es.CreatedAt = uint32(value.Int64)
			}
		case emailsubscriber.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				es.UpdatedAt = uint32(value.Int64)
			}
		case emailsubscriber.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				es.DeletedAt = uint32(value.Int64)
			}
		case emailsubscriber.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				es.AppID = *value
			}
		case emailsubscriber.FieldEmailAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email_address", values[i])
			} else if value.Valid {
				es.EmailAddress = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this EmailSubscriber.
// Note that you need to call EmailSubscriber.Unwrap() before calling this method if this EmailSubscriber
// was returned from a transaction, and the transaction was committed or rolled back.
func (es *EmailSubscriber) Update() *EmailSubscriberUpdateOne {
	return (&EmailSubscriberClient{config: es.config}).UpdateOne(es)
}

// Unwrap unwraps the EmailSubscriber entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (es *EmailSubscriber) Unwrap() *EmailSubscriber {
	tx, ok := es.config.driver.(*txDriver)
	if !ok {
		panic("ent: EmailSubscriber is not a transactional entity")
	}
	es.config.driver = tx.drv
	return es
}

// String implements the fmt.Stringer.
func (es *EmailSubscriber) String() string {
	var builder strings.Builder
	builder.WriteString("EmailSubscriber(")
	builder.WriteString(fmt.Sprintf("id=%v", es.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(fmt.Sprintf("%v", es.CreatedAt))
	builder.WriteString(", updated_at=")
	builder.WriteString(fmt.Sprintf("%v", es.UpdatedAt))
	builder.WriteString(", deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", es.DeletedAt))
	builder.WriteString(", app_id=")
	builder.WriteString(fmt.Sprintf("%v", es.AppID))
	builder.WriteString(", email_address=")
	builder.WriteString(es.EmailAddress)
	builder.WriteByte(')')
	return builder.String()
}

// EmailSubscribers is a parsable slice of EmailSubscriber.
type EmailSubscribers []*EmailSubscriber

func (es EmailSubscribers) config(cfg config) {
	for _i := range es {
		es[_i].config = cfg
	}
}
