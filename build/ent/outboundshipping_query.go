// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/satriahrh/adam-smith/build/ent/outboundshipping"
	"github.com/satriahrh/adam-smith/build/ent/outboundtransaction"
	"github.com/satriahrh/adam-smith/build/ent/predicate"
)

// OutboundShippingQuery is the builder for querying OutboundShipping entities.
type OutboundShippingQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	predicates []predicate.OutboundShipping
	// eager-loading edges.
	withTransaction *OutboundTransactionQuery
	withFKs         bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (osq *OutboundShippingQuery) Where(ps ...predicate.OutboundShipping) *OutboundShippingQuery {
	osq.predicates = append(osq.predicates, ps...)
	return osq
}

// Limit adds a limit step to the query.
func (osq *OutboundShippingQuery) Limit(limit int) *OutboundShippingQuery {
	osq.limit = &limit
	return osq
}

// Offset adds an offset step to the query.
func (osq *OutboundShippingQuery) Offset(offset int) *OutboundShippingQuery {
	osq.offset = &offset
	return osq
}

// Order adds an order step to the query.
func (osq *OutboundShippingQuery) Order(o ...OrderFunc) *OutboundShippingQuery {
	osq.order = append(osq.order, o...)
	return osq
}

// QueryTransaction chains the current query on the transaction edge.
func (osq *OutboundShippingQuery) QueryTransaction() *OutboundTransactionQuery {
	query := &OutboundTransactionQuery{config: osq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := osq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := osq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outboundshipping.Table, outboundshipping.FieldID, selector),
			sqlgraph.To(outboundtransaction.Table, outboundtransaction.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, outboundshipping.TransactionTable, outboundshipping.TransactionColumn),
		)
		fromU = sqlgraph.SetNeighbors(osq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OutboundShipping entity in the query. Returns *NotFoundError when no outboundshipping was found.
func (osq *OutboundShippingQuery) First(ctx context.Context) (*OutboundShipping, error) {
	nodes, err := osq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{outboundshipping.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (osq *OutboundShippingQuery) FirstX(ctx context.Context) *OutboundShipping {
	node, err := osq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OutboundShipping id in the query. Returns *NotFoundError when no id was found.
func (osq *OutboundShippingQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = osq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{outboundshipping.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (osq *OutboundShippingQuery) FirstIDX(ctx context.Context) int {
	id, err := osq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only OutboundShipping entity in the query, returns an error if not exactly one entity was returned.
func (osq *OutboundShippingQuery) Only(ctx context.Context) (*OutboundShipping, error) {
	nodes, err := osq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{outboundshipping.Label}
	default:
		return nil, &NotSingularError{outboundshipping.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (osq *OutboundShippingQuery) OnlyX(ctx context.Context) *OutboundShipping {
	node, err := osq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only OutboundShipping id in the query, returns an error if not exactly one id was returned.
func (osq *OutboundShippingQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = osq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{outboundshipping.Label}
	default:
		err = &NotSingularError{outboundshipping.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (osq *OutboundShippingQuery) OnlyIDX(ctx context.Context) int {
	id, err := osq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OutboundShippings.
func (osq *OutboundShippingQuery) All(ctx context.Context) ([]*OutboundShipping, error) {
	if err := osq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return osq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (osq *OutboundShippingQuery) AllX(ctx context.Context) []*OutboundShipping {
	nodes, err := osq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OutboundShipping ids.
func (osq *OutboundShippingQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := osq.Select(outboundshipping.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (osq *OutboundShippingQuery) IDsX(ctx context.Context) []int {
	ids, err := osq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (osq *OutboundShippingQuery) Count(ctx context.Context) (int, error) {
	if err := osq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return osq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (osq *OutboundShippingQuery) CountX(ctx context.Context) int {
	count, err := osq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (osq *OutboundShippingQuery) Exist(ctx context.Context) (bool, error) {
	if err := osq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return osq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (osq *OutboundShippingQuery) ExistX(ctx context.Context) bool {
	exist, err := osq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (osq *OutboundShippingQuery) Clone() *OutboundShippingQuery {
	if osq == nil {
		return nil
	}
	return &OutboundShippingQuery{
		config:          osq.config,
		limit:           osq.limit,
		offset:          osq.offset,
		order:           append([]OrderFunc{}, osq.order...),
		predicates:      append([]predicate.OutboundShipping{}, osq.predicates...),
		withTransaction: osq.withTransaction.Clone(),
		// clone intermediate query.
		sql:  osq.sql.Clone(),
		path: osq.path,
	}
}

//  WithTransaction tells the query-builder to eager-loads the nodes that are connected to
// the "transaction" edge. The optional arguments used to configure the query builder of the edge.
func (osq *OutboundShippingQuery) WithTransaction(opts ...func(*OutboundTransactionQuery)) *OutboundShippingQuery {
	query := &OutboundTransactionQuery{config: osq.config}
	for _, opt := range opts {
		opt(query)
	}
	osq.withTransaction = query
	return osq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Courier outboundshipping.Courier `json:"courier,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OutboundShipping.Query().
//		GroupBy(outboundshipping.FieldCourier).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (osq *OutboundShippingQuery) GroupBy(field string, fields ...string) *OutboundShippingGroupBy {
	group := &OutboundShippingGroupBy{config: osq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := osq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return osq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Courier outboundshipping.Courier `json:"courier,omitempty"`
//	}
//
//	client.OutboundShipping.Query().
//		Select(outboundshipping.FieldCourier).
//		Scan(ctx, &v)
//
func (osq *OutboundShippingQuery) Select(field string, fields ...string) *OutboundShippingSelect {
	selector := &OutboundShippingSelect{config: osq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := osq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return osq.sqlQuery(), nil
	}
	return selector
}

func (osq *OutboundShippingQuery) prepareQuery(ctx context.Context) error {
	if osq.path != nil {
		prev, err := osq.path(ctx)
		if err != nil {
			return err
		}
		osq.sql = prev
	}
	return nil
}

func (osq *OutboundShippingQuery) sqlAll(ctx context.Context) ([]*OutboundShipping, error) {
	var (
		nodes       = []*OutboundShipping{}
		withFKs     = osq.withFKs
		_spec       = osq.querySpec()
		loadedTypes = [1]bool{
			osq.withTransaction != nil,
		}
	)
	if osq.withTransaction != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, outboundshipping.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &OutboundShipping{config: osq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, osq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := osq.withTransaction; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OutboundShipping)
		for i := range nodes {
			if fk := nodes[i].outbound_transaction_shipping; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(outboundtransaction.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "outbound_transaction_shipping" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Transaction = n
			}
		}
	}

	return nodes, nil
}

func (osq *OutboundShippingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := osq.querySpec()
	return sqlgraph.CountNodes(ctx, osq.driver, _spec)
}

func (osq *OutboundShippingQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := osq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (osq *OutboundShippingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outboundshipping.Table,
			Columns: outboundshipping.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outboundshipping.FieldID,
			},
		},
		From:   osq.sql,
		Unique: true,
	}
	if ps := osq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := osq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := osq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := osq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, outboundshipping.ValidColumn)
			}
		}
	}
	return _spec
}

func (osq *OutboundShippingQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(osq.driver.Dialect())
	t1 := builder.Table(outboundshipping.Table)
	selector := builder.Select(t1.Columns(outboundshipping.Columns...)...).From(t1)
	if osq.sql != nil {
		selector = osq.sql
		selector.Select(selector.Columns(outboundshipping.Columns...)...)
	}
	for _, p := range osq.predicates {
		p(selector)
	}
	for _, p := range osq.order {
		p(selector, outboundshipping.ValidColumn)
	}
	if offset := osq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := osq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OutboundShippingGroupBy is the builder for group-by OutboundShipping entities.
type OutboundShippingGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (osgb *OutboundShippingGroupBy) Aggregate(fns ...AggregateFunc) *OutboundShippingGroupBy {
	osgb.fns = append(osgb.fns, fns...)
	return osgb
}

// Scan applies the group-by query and scan the result into the given value.
func (osgb *OutboundShippingGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := osgb.path(ctx)
	if err != nil {
		return err
	}
	osgb.sql = query
	return osgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (osgb *OutboundShippingGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := osgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (osgb *OutboundShippingGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(osgb.fields) > 1 {
		return nil, errors.New("ent: OutboundShippingGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := osgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (osgb *OutboundShippingGroupBy) StringsX(ctx context.Context) []string {
	v, err := osgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (osgb *OutboundShippingGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = osgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{outboundshipping.Label}
	default:
		err = fmt.Errorf("ent: OutboundShippingGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (osgb *OutboundShippingGroupBy) StringX(ctx context.Context) string {
	v, err := osgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (osgb *OutboundShippingGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(osgb.fields) > 1 {
		return nil, errors.New("ent: OutboundShippingGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := osgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (osgb *OutboundShippingGroupBy) IntsX(ctx context.Context) []int {
	v, err := osgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (osgb *OutboundShippingGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = osgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{outboundshipping.Label}
	default:
		err = fmt.Errorf("ent: OutboundShippingGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (osgb *OutboundShippingGroupBy) IntX(ctx context.Context) int {
	v, err := osgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (osgb *OutboundShippingGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(osgb.fields) > 1 {
		return nil, errors.New("ent: OutboundShippingGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := osgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (osgb *OutboundShippingGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := osgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (osgb *OutboundShippingGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = osgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{outboundshipping.Label}
	default:
		err = fmt.Errorf("ent: OutboundShippingGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (osgb *OutboundShippingGroupBy) Float64X(ctx context.Context) float64 {
	v, err := osgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (osgb *OutboundShippingGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(osgb.fields) > 1 {
		return nil, errors.New("ent: OutboundShippingGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := osgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (osgb *OutboundShippingGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := osgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (osgb *OutboundShippingGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = osgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{outboundshipping.Label}
	default:
		err = fmt.Errorf("ent: OutboundShippingGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (osgb *OutboundShippingGroupBy) BoolX(ctx context.Context) bool {
	v, err := osgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (osgb *OutboundShippingGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range osgb.fields {
		if !outboundshipping.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := osgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := osgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (osgb *OutboundShippingGroupBy) sqlQuery() *sql.Selector {
	selector := osgb.sql
	columns := make([]string, 0, len(osgb.fields)+len(osgb.fns))
	columns = append(columns, osgb.fields...)
	for _, fn := range osgb.fns {
		columns = append(columns, fn(selector, outboundshipping.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(osgb.fields...)
}

// OutboundShippingSelect is the builder for select fields of OutboundShipping entities.
type OutboundShippingSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (oss *OutboundShippingSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := oss.path(ctx)
	if err != nil {
		return err
	}
	oss.sql = query
	return oss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (oss *OutboundShippingSelect) ScanX(ctx context.Context, v interface{}) {
	if err := oss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (oss *OutboundShippingSelect) Strings(ctx context.Context) ([]string, error) {
	if len(oss.fields) > 1 {
		return nil, errors.New("ent: OutboundShippingSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := oss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (oss *OutboundShippingSelect) StringsX(ctx context.Context) []string {
	v, err := oss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (oss *OutboundShippingSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = oss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{outboundshipping.Label}
	default:
		err = fmt.Errorf("ent: OutboundShippingSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (oss *OutboundShippingSelect) StringX(ctx context.Context) string {
	v, err := oss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (oss *OutboundShippingSelect) Ints(ctx context.Context) ([]int, error) {
	if len(oss.fields) > 1 {
		return nil, errors.New("ent: OutboundShippingSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := oss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (oss *OutboundShippingSelect) IntsX(ctx context.Context) []int {
	v, err := oss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (oss *OutboundShippingSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = oss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{outboundshipping.Label}
	default:
		err = fmt.Errorf("ent: OutboundShippingSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (oss *OutboundShippingSelect) IntX(ctx context.Context) int {
	v, err := oss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (oss *OutboundShippingSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(oss.fields) > 1 {
		return nil, errors.New("ent: OutboundShippingSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := oss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (oss *OutboundShippingSelect) Float64sX(ctx context.Context) []float64 {
	v, err := oss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (oss *OutboundShippingSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = oss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{outboundshipping.Label}
	default:
		err = fmt.Errorf("ent: OutboundShippingSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (oss *OutboundShippingSelect) Float64X(ctx context.Context) float64 {
	v, err := oss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (oss *OutboundShippingSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(oss.fields) > 1 {
		return nil, errors.New("ent: OutboundShippingSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := oss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (oss *OutboundShippingSelect) BoolsX(ctx context.Context) []bool {
	v, err := oss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (oss *OutboundShippingSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = oss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{outboundshipping.Label}
	default:
		err = fmt.Errorf("ent: OutboundShippingSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (oss *OutboundShippingSelect) BoolX(ctx context.Context) bool {
	v, err := oss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (oss *OutboundShippingSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range oss.fields {
		if !outboundshipping.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := oss.sqlQuery().Query()
	if err := oss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (oss *OutboundShippingSelect) sqlQuery() sql.Querier {
	selector := oss.sql
	selector.Select(selector.Columns(oss.fields...)...)
	return selector
}
