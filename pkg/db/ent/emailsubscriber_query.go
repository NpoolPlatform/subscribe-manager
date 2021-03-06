// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/subscribe-manager/pkg/db/ent/emailsubscriber"
	"github.com/NpoolPlatform/subscribe-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// EmailSubscriberQuery is the builder for querying EmailSubscriber entities.
type EmailSubscriberQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.EmailSubscriber
	modifiers  []func(s *sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EmailSubscriberQuery builder.
func (esq *EmailSubscriberQuery) Where(ps ...predicate.EmailSubscriber) *EmailSubscriberQuery {
	esq.predicates = append(esq.predicates, ps...)
	return esq
}

// Limit adds a limit step to the query.
func (esq *EmailSubscriberQuery) Limit(limit int) *EmailSubscriberQuery {
	esq.limit = &limit
	return esq
}

// Offset adds an offset step to the query.
func (esq *EmailSubscriberQuery) Offset(offset int) *EmailSubscriberQuery {
	esq.offset = &offset
	return esq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (esq *EmailSubscriberQuery) Unique(unique bool) *EmailSubscriberQuery {
	esq.unique = &unique
	return esq
}

// Order adds an order step to the query.
func (esq *EmailSubscriberQuery) Order(o ...OrderFunc) *EmailSubscriberQuery {
	esq.order = append(esq.order, o...)
	return esq
}

// First returns the first EmailSubscriber entity from the query.
// Returns a *NotFoundError when no EmailSubscriber was found.
func (esq *EmailSubscriberQuery) First(ctx context.Context) (*EmailSubscriber, error) {
	nodes, err := esq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{emailsubscriber.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (esq *EmailSubscriberQuery) FirstX(ctx context.Context) *EmailSubscriber {
	node, err := esq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EmailSubscriber ID from the query.
// Returns a *NotFoundError when no EmailSubscriber ID was found.
func (esq *EmailSubscriberQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = esq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{emailsubscriber.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (esq *EmailSubscriberQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := esq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EmailSubscriber entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EmailSubscriber entity is found.
// Returns a *NotFoundError when no EmailSubscriber entities are found.
func (esq *EmailSubscriberQuery) Only(ctx context.Context) (*EmailSubscriber, error) {
	nodes, err := esq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{emailsubscriber.Label}
	default:
		return nil, &NotSingularError{emailsubscriber.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (esq *EmailSubscriberQuery) OnlyX(ctx context.Context) *EmailSubscriber {
	node, err := esq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EmailSubscriber ID in the query.
// Returns a *NotSingularError when more than one EmailSubscriber ID is found.
// Returns a *NotFoundError when no entities are found.
func (esq *EmailSubscriberQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = esq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{emailsubscriber.Label}
	default:
		err = &NotSingularError{emailsubscriber.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (esq *EmailSubscriberQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := esq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EmailSubscribers.
func (esq *EmailSubscriberQuery) All(ctx context.Context) ([]*EmailSubscriber, error) {
	if err := esq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return esq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (esq *EmailSubscriberQuery) AllX(ctx context.Context) []*EmailSubscriber {
	nodes, err := esq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EmailSubscriber IDs.
func (esq *EmailSubscriberQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := esq.Select(emailsubscriber.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (esq *EmailSubscriberQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := esq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (esq *EmailSubscriberQuery) Count(ctx context.Context) (int, error) {
	if err := esq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return esq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (esq *EmailSubscriberQuery) CountX(ctx context.Context) int {
	count, err := esq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (esq *EmailSubscriberQuery) Exist(ctx context.Context) (bool, error) {
	if err := esq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return esq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (esq *EmailSubscriberQuery) ExistX(ctx context.Context) bool {
	exist, err := esq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EmailSubscriberQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (esq *EmailSubscriberQuery) Clone() *EmailSubscriberQuery {
	if esq == nil {
		return nil
	}
	return &EmailSubscriberQuery{
		config:     esq.config,
		limit:      esq.limit,
		offset:     esq.offset,
		order:      append([]OrderFunc{}, esq.order...),
		predicates: append([]predicate.EmailSubscriber{}, esq.predicates...),
		// clone intermediate query.
		sql:    esq.sql.Clone(),
		path:   esq.path,
		unique: esq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EmailSubscriber.Query().
//		GroupBy(emailsubscriber.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (esq *EmailSubscriberQuery) GroupBy(field string, fields ...string) *EmailSubscriberGroupBy {
	group := &EmailSubscriberGroupBy{config: esq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := esq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return esq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//	}
//
//	client.EmailSubscriber.Query().
//		Select(emailsubscriber.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (esq *EmailSubscriberQuery) Select(fields ...string) *EmailSubscriberSelect {
	esq.fields = append(esq.fields, fields...)
	return &EmailSubscriberSelect{EmailSubscriberQuery: esq}
}

func (esq *EmailSubscriberQuery) prepareQuery(ctx context.Context) error {
	for _, f := range esq.fields {
		if !emailsubscriber.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if esq.path != nil {
		prev, err := esq.path(ctx)
		if err != nil {
			return err
		}
		esq.sql = prev
	}
	if emailsubscriber.Policy == nil {
		return errors.New("ent: uninitialized emailsubscriber.Policy (forgotten import ent/runtime?)")
	}
	if err := emailsubscriber.Policy.EvalQuery(ctx, esq); err != nil {
		return err
	}
	return nil
}

func (esq *EmailSubscriberQuery) sqlAll(ctx context.Context) ([]*EmailSubscriber, error) {
	var (
		nodes = []*EmailSubscriber{}
		_spec = esq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &EmailSubscriber{config: esq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if len(esq.modifiers) > 0 {
		_spec.Modifiers = esq.modifiers
	}
	if err := sqlgraph.QueryNodes(ctx, esq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (esq *EmailSubscriberQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := esq.querySpec()
	if len(esq.modifiers) > 0 {
		_spec.Modifiers = esq.modifiers
	}
	_spec.Node.Columns = esq.fields
	if len(esq.fields) > 0 {
		_spec.Unique = esq.unique != nil && *esq.unique
	}
	return sqlgraph.CountNodes(ctx, esq.driver, _spec)
}

func (esq *EmailSubscriberQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := esq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (esq *EmailSubscriberQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   emailsubscriber.Table,
			Columns: emailsubscriber.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: emailsubscriber.FieldID,
			},
		},
		From:   esq.sql,
		Unique: true,
	}
	if unique := esq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := esq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, emailsubscriber.FieldID)
		for i := range fields {
			if fields[i] != emailsubscriber.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := esq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := esq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := esq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := esq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (esq *EmailSubscriberQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(esq.driver.Dialect())
	t1 := builder.Table(emailsubscriber.Table)
	columns := esq.fields
	if len(columns) == 0 {
		columns = emailsubscriber.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if esq.sql != nil {
		selector = esq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if esq.unique != nil && *esq.unique {
		selector.Distinct()
	}
	for _, m := range esq.modifiers {
		m(selector)
	}
	for _, p := range esq.predicates {
		p(selector)
	}
	for _, p := range esq.order {
		p(selector)
	}
	if offset := esq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := esq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (esq *EmailSubscriberQuery) ForUpdate(opts ...sql.LockOption) *EmailSubscriberQuery {
	if esq.driver.Dialect() == dialect.Postgres {
		esq.Unique(false)
	}
	esq.modifiers = append(esq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return esq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (esq *EmailSubscriberQuery) ForShare(opts ...sql.LockOption) *EmailSubscriberQuery {
	if esq.driver.Dialect() == dialect.Postgres {
		esq.Unique(false)
	}
	esq.modifiers = append(esq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return esq
}

// EmailSubscriberGroupBy is the group-by builder for EmailSubscriber entities.
type EmailSubscriberGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (esgb *EmailSubscriberGroupBy) Aggregate(fns ...AggregateFunc) *EmailSubscriberGroupBy {
	esgb.fns = append(esgb.fns, fns...)
	return esgb
}

// Scan applies the group-by query and scans the result into the given value.
func (esgb *EmailSubscriberGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := esgb.path(ctx)
	if err != nil {
		return err
	}
	esgb.sql = query
	return esgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (esgb *EmailSubscriberGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := esgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (esgb *EmailSubscriberGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(esgb.fields) > 1 {
		return nil, errors.New("ent: EmailSubscriberGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := esgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (esgb *EmailSubscriberGroupBy) StringsX(ctx context.Context) []string {
	v, err := esgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (esgb *EmailSubscriberGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = esgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{emailsubscriber.Label}
	default:
		err = fmt.Errorf("ent: EmailSubscriberGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (esgb *EmailSubscriberGroupBy) StringX(ctx context.Context) string {
	v, err := esgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (esgb *EmailSubscriberGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(esgb.fields) > 1 {
		return nil, errors.New("ent: EmailSubscriberGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := esgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (esgb *EmailSubscriberGroupBy) IntsX(ctx context.Context) []int {
	v, err := esgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (esgb *EmailSubscriberGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = esgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{emailsubscriber.Label}
	default:
		err = fmt.Errorf("ent: EmailSubscriberGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (esgb *EmailSubscriberGroupBy) IntX(ctx context.Context) int {
	v, err := esgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (esgb *EmailSubscriberGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(esgb.fields) > 1 {
		return nil, errors.New("ent: EmailSubscriberGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := esgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (esgb *EmailSubscriberGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := esgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (esgb *EmailSubscriberGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = esgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{emailsubscriber.Label}
	default:
		err = fmt.Errorf("ent: EmailSubscriberGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (esgb *EmailSubscriberGroupBy) Float64X(ctx context.Context) float64 {
	v, err := esgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (esgb *EmailSubscriberGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(esgb.fields) > 1 {
		return nil, errors.New("ent: EmailSubscriberGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := esgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (esgb *EmailSubscriberGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := esgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (esgb *EmailSubscriberGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = esgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{emailsubscriber.Label}
	default:
		err = fmt.Errorf("ent: EmailSubscriberGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (esgb *EmailSubscriberGroupBy) BoolX(ctx context.Context) bool {
	v, err := esgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (esgb *EmailSubscriberGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range esgb.fields {
		if !emailsubscriber.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := esgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := esgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (esgb *EmailSubscriberGroupBy) sqlQuery() *sql.Selector {
	selector := esgb.sql.Select()
	aggregation := make([]string, 0, len(esgb.fns))
	for _, fn := range esgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(esgb.fields)+len(esgb.fns))
		for _, f := range esgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(esgb.fields...)...)
}

// EmailSubscriberSelect is the builder for selecting fields of EmailSubscriber entities.
type EmailSubscriberSelect struct {
	*EmailSubscriberQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ess *EmailSubscriberSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ess.prepareQuery(ctx); err != nil {
		return err
	}
	ess.sql = ess.EmailSubscriberQuery.sqlQuery(ctx)
	return ess.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ess *EmailSubscriberSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ess.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ess *EmailSubscriberSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ess.fields) > 1 {
		return nil, errors.New("ent: EmailSubscriberSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ess.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ess *EmailSubscriberSelect) StringsX(ctx context.Context) []string {
	v, err := ess.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ess *EmailSubscriberSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ess.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{emailsubscriber.Label}
	default:
		err = fmt.Errorf("ent: EmailSubscriberSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ess *EmailSubscriberSelect) StringX(ctx context.Context) string {
	v, err := ess.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ess *EmailSubscriberSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ess.fields) > 1 {
		return nil, errors.New("ent: EmailSubscriberSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ess.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ess *EmailSubscriberSelect) IntsX(ctx context.Context) []int {
	v, err := ess.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ess *EmailSubscriberSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ess.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{emailsubscriber.Label}
	default:
		err = fmt.Errorf("ent: EmailSubscriberSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ess *EmailSubscriberSelect) IntX(ctx context.Context) int {
	v, err := ess.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ess *EmailSubscriberSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ess.fields) > 1 {
		return nil, errors.New("ent: EmailSubscriberSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ess.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ess *EmailSubscriberSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ess.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ess *EmailSubscriberSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ess.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{emailsubscriber.Label}
	default:
		err = fmt.Errorf("ent: EmailSubscriberSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ess *EmailSubscriberSelect) Float64X(ctx context.Context) float64 {
	v, err := ess.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ess *EmailSubscriberSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ess.fields) > 1 {
		return nil, errors.New("ent: EmailSubscriberSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ess.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ess *EmailSubscriberSelect) BoolsX(ctx context.Context) []bool {
	v, err := ess.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ess *EmailSubscriberSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ess.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{emailsubscriber.Label}
	default:
		err = fmt.Errorf("ent: EmailSubscriberSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ess *EmailSubscriberSelect) BoolX(ctx context.Context) bool {
	v, err := ess.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ess *EmailSubscriberSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ess.sql.Query()
	if err := ess.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
