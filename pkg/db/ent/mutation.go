// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/NpoolPlatform/subscribe-manager/pkg/db/ent/emailsubscriber"
	"github.com/NpoolPlatform/subscribe-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeEmailSubscriber = "EmailSubscriber"
)

// EmailSubscriberMutation represents an operation that mutates the EmailSubscriber nodes in the graph.
type EmailSubscriberMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	created_at    *uint32
	addcreated_at *int32
	updated_at    *uint32
	addupdated_at *int32
	deleted_at    *uint32
	adddeleted_at *int32
	app_id        *uuid.UUID
	email_address *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*EmailSubscriber, error)
	predicates    []predicate.EmailSubscriber
}

var _ ent.Mutation = (*EmailSubscriberMutation)(nil)

// emailsubscriberOption allows management of the mutation configuration using functional options.
type emailsubscriberOption func(*EmailSubscriberMutation)

// newEmailSubscriberMutation creates new mutation for the EmailSubscriber entity.
func newEmailSubscriberMutation(c config, op Op, opts ...emailsubscriberOption) *EmailSubscriberMutation {
	m := &EmailSubscriberMutation{
		config:        c,
		op:            op,
		typ:           TypeEmailSubscriber,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withEmailSubscriberID sets the ID field of the mutation.
func withEmailSubscriberID(id uuid.UUID) emailsubscriberOption {
	return func(m *EmailSubscriberMutation) {
		var (
			err   error
			once  sync.Once
			value *EmailSubscriber
		)
		m.oldValue = func(ctx context.Context) (*EmailSubscriber, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().EmailSubscriber.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withEmailSubscriber sets the old EmailSubscriber of the mutation.
func withEmailSubscriber(node *EmailSubscriber) emailsubscriberOption {
	return func(m *EmailSubscriberMutation) {
		m.oldValue = func(context.Context) (*EmailSubscriber, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m EmailSubscriberMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m EmailSubscriberMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of EmailSubscriber entities.
func (m *EmailSubscriberMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *EmailSubscriberMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *EmailSubscriberMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().EmailSubscriber.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *EmailSubscriberMutation) SetCreatedAt(u uint32) {
	m.created_at = &u
	m.addcreated_at = nil
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *EmailSubscriberMutation) CreatedAt() (r uint32, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the EmailSubscriber entity.
// If the EmailSubscriber object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmailSubscriberMutation) OldCreatedAt(ctx context.Context) (v uint32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// AddCreatedAt adds u to the "created_at" field.
func (m *EmailSubscriberMutation) AddCreatedAt(u int32) {
	if m.addcreated_at != nil {
		*m.addcreated_at += u
	} else {
		m.addcreated_at = &u
	}
}

// AddedCreatedAt returns the value that was added to the "created_at" field in this mutation.
func (m *EmailSubscriberMutation) AddedCreatedAt() (r int32, exists bool) {
	v := m.addcreated_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *EmailSubscriberMutation) ResetCreatedAt() {
	m.created_at = nil
	m.addcreated_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *EmailSubscriberMutation) SetUpdatedAt(u uint32) {
	m.updated_at = &u
	m.addupdated_at = nil
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *EmailSubscriberMutation) UpdatedAt() (r uint32, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the EmailSubscriber entity.
// If the EmailSubscriber object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmailSubscriberMutation) OldUpdatedAt(ctx context.Context) (v uint32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// AddUpdatedAt adds u to the "updated_at" field.
func (m *EmailSubscriberMutation) AddUpdatedAt(u int32) {
	if m.addupdated_at != nil {
		*m.addupdated_at += u
	} else {
		m.addupdated_at = &u
	}
}

// AddedUpdatedAt returns the value that was added to the "updated_at" field in this mutation.
func (m *EmailSubscriberMutation) AddedUpdatedAt() (r int32, exists bool) {
	v := m.addupdated_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *EmailSubscriberMutation) ResetUpdatedAt() {
	m.updated_at = nil
	m.addupdated_at = nil
}

// SetDeletedAt sets the "deleted_at" field.
func (m *EmailSubscriberMutation) SetDeletedAt(u uint32) {
	m.deleted_at = &u
	m.adddeleted_at = nil
}

// DeletedAt returns the value of the "deleted_at" field in the mutation.
func (m *EmailSubscriberMutation) DeletedAt() (r uint32, exists bool) {
	v := m.deleted_at
	if v == nil {
		return
	}
	return *v, true
}

// OldDeletedAt returns the old "deleted_at" field's value of the EmailSubscriber entity.
// If the EmailSubscriber object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmailSubscriberMutation) OldDeletedAt(ctx context.Context) (v uint32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDeletedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDeletedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDeletedAt: %w", err)
	}
	return oldValue.DeletedAt, nil
}

// AddDeletedAt adds u to the "deleted_at" field.
func (m *EmailSubscriberMutation) AddDeletedAt(u int32) {
	if m.adddeleted_at != nil {
		*m.adddeleted_at += u
	} else {
		m.adddeleted_at = &u
	}
}

// AddedDeletedAt returns the value that was added to the "deleted_at" field in this mutation.
func (m *EmailSubscriberMutation) AddedDeletedAt() (r int32, exists bool) {
	v := m.adddeleted_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetDeletedAt resets all changes to the "deleted_at" field.
func (m *EmailSubscriberMutation) ResetDeletedAt() {
	m.deleted_at = nil
	m.adddeleted_at = nil
}

// SetAppID sets the "app_id" field.
func (m *EmailSubscriberMutation) SetAppID(u uuid.UUID) {
	m.app_id = &u
}

// AppID returns the value of the "app_id" field in the mutation.
func (m *EmailSubscriberMutation) AppID() (r uuid.UUID, exists bool) {
	v := m.app_id
	if v == nil {
		return
	}
	return *v, true
}

// OldAppID returns the old "app_id" field's value of the EmailSubscriber entity.
// If the EmailSubscriber object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmailSubscriberMutation) OldAppID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAppID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAppID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAppID: %w", err)
	}
	return oldValue.AppID, nil
}

// ResetAppID resets all changes to the "app_id" field.
func (m *EmailSubscriberMutation) ResetAppID() {
	m.app_id = nil
}

// SetEmailAddress sets the "email_address" field.
func (m *EmailSubscriberMutation) SetEmailAddress(s string) {
	m.email_address = &s
}

// EmailAddress returns the value of the "email_address" field in the mutation.
func (m *EmailSubscriberMutation) EmailAddress() (r string, exists bool) {
	v := m.email_address
	if v == nil {
		return
	}
	return *v, true
}

// OldEmailAddress returns the old "email_address" field's value of the EmailSubscriber entity.
// If the EmailSubscriber object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmailSubscriberMutation) OldEmailAddress(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmailAddress is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmailAddress requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmailAddress: %w", err)
	}
	return oldValue.EmailAddress, nil
}

// ResetEmailAddress resets all changes to the "email_address" field.
func (m *EmailSubscriberMutation) ResetEmailAddress() {
	m.email_address = nil
}

// Where appends a list predicates to the EmailSubscriberMutation builder.
func (m *EmailSubscriberMutation) Where(ps ...predicate.EmailSubscriber) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *EmailSubscriberMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (EmailSubscriber).
func (m *EmailSubscriberMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *EmailSubscriberMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.created_at != nil {
		fields = append(fields, emailsubscriber.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, emailsubscriber.FieldUpdatedAt)
	}
	if m.deleted_at != nil {
		fields = append(fields, emailsubscriber.FieldDeletedAt)
	}
	if m.app_id != nil {
		fields = append(fields, emailsubscriber.FieldAppID)
	}
	if m.email_address != nil {
		fields = append(fields, emailsubscriber.FieldEmailAddress)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *EmailSubscriberMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case emailsubscriber.FieldCreatedAt:
		return m.CreatedAt()
	case emailsubscriber.FieldUpdatedAt:
		return m.UpdatedAt()
	case emailsubscriber.FieldDeletedAt:
		return m.DeletedAt()
	case emailsubscriber.FieldAppID:
		return m.AppID()
	case emailsubscriber.FieldEmailAddress:
		return m.EmailAddress()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *EmailSubscriberMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case emailsubscriber.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case emailsubscriber.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case emailsubscriber.FieldDeletedAt:
		return m.OldDeletedAt(ctx)
	case emailsubscriber.FieldAppID:
		return m.OldAppID(ctx)
	case emailsubscriber.FieldEmailAddress:
		return m.OldEmailAddress(ctx)
	}
	return nil, fmt.Errorf("unknown EmailSubscriber field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *EmailSubscriberMutation) SetField(name string, value ent.Value) error {
	switch name {
	case emailsubscriber.FieldCreatedAt:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case emailsubscriber.FieldUpdatedAt:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case emailsubscriber.FieldDeletedAt:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDeletedAt(v)
		return nil
	case emailsubscriber.FieldAppID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAppID(v)
		return nil
	case emailsubscriber.FieldEmailAddress:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmailAddress(v)
		return nil
	}
	return fmt.Errorf("unknown EmailSubscriber field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *EmailSubscriberMutation) AddedFields() []string {
	var fields []string
	if m.addcreated_at != nil {
		fields = append(fields, emailsubscriber.FieldCreatedAt)
	}
	if m.addupdated_at != nil {
		fields = append(fields, emailsubscriber.FieldUpdatedAt)
	}
	if m.adddeleted_at != nil {
		fields = append(fields, emailsubscriber.FieldDeletedAt)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *EmailSubscriberMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case emailsubscriber.FieldCreatedAt:
		return m.AddedCreatedAt()
	case emailsubscriber.FieldUpdatedAt:
		return m.AddedUpdatedAt()
	case emailsubscriber.FieldDeletedAt:
		return m.AddedDeletedAt()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *EmailSubscriberMutation) AddField(name string, value ent.Value) error {
	switch name {
	case emailsubscriber.FieldCreatedAt:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddCreatedAt(v)
		return nil
	case emailsubscriber.FieldUpdatedAt:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddUpdatedAt(v)
		return nil
	case emailsubscriber.FieldDeletedAt:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddDeletedAt(v)
		return nil
	}
	return fmt.Errorf("unknown EmailSubscriber numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *EmailSubscriberMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *EmailSubscriberMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *EmailSubscriberMutation) ClearField(name string) error {
	return fmt.Errorf("unknown EmailSubscriber nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *EmailSubscriberMutation) ResetField(name string) error {
	switch name {
	case emailsubscriber.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case emailsubscriber.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case emailsubscriber.FieldDeletedAt:
		m.ResetDeletedAt()
		return nil
	case emailsubscriber.FieldAppID:
		m.ResetAppID()
		return nil
	case emailsubscriber.FieldEmailAddress:
		m.ResetEmailAddress()
		return nil
	}
	return fmt.Errorf("unknown EmailSubscriber field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *EmailSubscriberMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *EmailSubscriberMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *EmailSubscriberMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *EmailSubscriberMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *EmailSubscriberMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *EmailSubscriberMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *EmailSubscriberMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown EmailSubscriber unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *EmailSubscriberMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown EmailSubscriber edge %s", name)
}
