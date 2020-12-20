// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/satriahrh/adam-smith/build/ent/outbounddeal"
	"github.com/satriahrh/adam-smith/build/ent/outboundtransaction"
	"github.com/satriahrh/adam-smith/build/ent/predicate"
	"github.com/satriahrh/adam-smith/build/ent/variant"
)

// OutboundDealQuery is the builder for querying OutboundDeal entities.
type OutboundDealQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	predicates []predicate.OutboundDeal
	// eager-loading edges.
	withVariant     *VariantQuery
	withParent      *OutboundDealQuery
	withChildren    *OutboundDealQuery
	withTransaction *OutboundTransactionQuery
	withFKs         bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (odq *OutboundDealQuery) Where(ps ...predicate.OutboundDeal) *OutboundDealQuery {
	odq.predicates = append(odq.predicates, ps...)
	return odq
}

// Limit adds a limit step to the query.
func (odq *OutboundDealQuery) Limit(limit int) *OutboundDealQuery {
	odq.limit = &limit
	return odq
}

// Offset adds an offset step to the query.
func (odq *OutboundDealQuery) Offset(offset int) *OutboundDealQuery {
	odq.offset = &offset
	return odq
}

// Order adds an order step to the query.
func (odq *OutboundDealQuery) Order(o ...OrderFunc) *OutboundDealQuery {
	odq.order = append(odq.order, o...)
	return odq
}

// QueryVariant chains the current query on the variant edge.
func (odq *OutboundDealQuery) QueryVariant() *VariantQuery {
	query := &VariantQuery{config: odq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := odq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := odq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outbounddeal.Table, outbounddeal.FieldID, selector),
			sqlgraph.To(variant.Table, variant.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, outbounddeal.VariantTable, outbounddeal.VariantColumn),
		)
		fromU = sqlgraph.SetNeighbors(odq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryParent chains the current query on the parent edge.
func (odq *OutboundDealQuery) QueryParent() *OutboundDealQuery {
	query := &OutboundDealQuery{config: odq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := odq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := odq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outbounddeal.Table, outbounddeal.FieldID, selector),
			sqlgraph.To(outbounddeal.Table, outbounddeal.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, outbounddeal.ParentTable, outbounddeal.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(odq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the children edge.
func (odq *OutboundDealQuery) QueryChildren() *OutboundDealQuery {
	query := &OutboundDealQuery{config: odq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := odq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := odq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outbounddeal.Table, outbounddeal.FieldID, selector),
			sqlgraph.To(outbounddeal.Table, outbounddeal.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, outbounddeal.ChildrenTable, outbounddeal.ChildrenColumn),
		)
		fromU = sqlgraph.SetNeighbors(odq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTransaction chains the current query on the transaction edge.
func (odq *OutboundDealQuery) QueryTransaction() *OutboundTransactionQuery {
	query := &OutboundTransactionQuery{config: odq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := odq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := odq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(outbounddeal.Table, outbounddeal.FieldID, selector),
			sqlgraph.To(outboundtransaction.Table, outboundtransaction.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, outbounddeal.TransactionTable, outbounddeal.TransactionColumn),
		)
		fromU = sqlgraph.SetNeighbors(odq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OutboundDeal entity in the query. Returns *NotFoundError when no outbounddeal was found.
func (odq *OutboundDealQuery) First(ctx context.Context) (*OutboundDeal, error) {
	nodes, err := odq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{outbounddeal.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (odq *OutboundDealQuery) FirstX(ctx context.Context) *OutboundDeal {
	node, err := odq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OutboundDeal id in the query. Returns *NotFoundError when no id was found.
func (odq *OutboundDealQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = odq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{outbounddeal.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (odq *OutboundDealQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := odq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only OutboundDeal entity in the query, returns an error if not exactly one entity was returned.
func (odq *OutboundDealQuery) Only(ctx context.Context) (*OutboundDeal, error) {
	nodes, err := odq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{outbounddeal.Label}
	default:
		return nil, &NotSingularError{outbounddeal.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (odq *OutboundDealQuery) OnlyX(ctx context.Context) *OutboundDeal {
	node, err := odq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only OutboundDeal id in the query, returns an error if not exactly one id was returned.
func (odq *OutboundDealQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = odq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{outbounddeal.Label}
	default:
		err = &NotSingularError{outbounddeal.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (odq *OutboundDealQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := odq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OutboundDeals.
func (odq *OutboundDealQuery) All(ctx context.Context) ([]*OutboundDeal, error) {
	if err := odq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return odq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (odq *OutboundDealQuery) AllX(ctx context.Context) []*OutboundDeal {
	nodes, err := odq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OutboundDeal ids.
func (odq *OutboundDealQuery) IDs(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	if err := odq.Select(outbounddeal.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (odq *OutboundDealQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := odq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (odq *OutboundDealQuery) Count(ctx context.Context) (int, error) {
	if err := odq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return odq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (odq *OutboundDealQuery) CountX(ctx context.Context) int {
	count, err := odq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (odq *OutboundDealQuery) Exist(ctx context.Context) (bool, error) {
	if err := odq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return odq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (odq *OutboundDealQuery) ExistX(ctx context.Context) bool {
	exist, err := odq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (odq *OutboundDealQuery) Clone() *OutboundDealQuery {
	if odq == nil {
		return nil
	}
	return &OutboundDealQuery{
		config:          odq.config,
		limit:           odq.limit,
		offset:          odq.offset,
		order:           append([]OrderFunc{}, odq.order...),
		predicates:      append([]predicate.OutboundDeal{}, odq.predicates...),
		withVariant:     odq.withVariant.Clone(),
		withParent:      odq.withParent.Clone(),
		withChildren:    odq.withChildren.Clone(),
		withTransaction: odq.withTransaction.Clone(),
		// clone intermediate query.
		sql:  odq.sql.Clone(),
		path: odq.path,
	}
}

//  WithVariant tells the query-builder to eager-loads the nodes that are connected to
// the "variant" edge. The optional arguments used to configure the query builder of the edge.
func (odq *OutboundDealQuery) WithVariant(opts ...func(*VariantQuery)) *OutboundDealQuery {
	query := &VariantQuery{config: odq.config}
	for _, opt := range opts {
		opt(query)
	}
	odq.withVariant = query
	return odq
}

//  WithParent tells the query-builder to eager-loads the nodes that are connected to
// the "parent" edge. The optional arguments used to configure the query builder of the edge.
func (odq *OutboundDealQuery) WithParent(opts ...func(*OutboundDealQuery)) *OutboundDealQuery {
	query := &OutboundDealQuery{config: odq.config}
	for _, opt := range opts {
		opt(query)
	}
	odq.withParent = query
	return odq
}

//  WithChildren tells the query-builder to eager-loads the nodes that are connected to
// the "children" edge. The optional arguments used to configure the query builder of the edge.
func (odq *OutboundDealQuery) WithChildren(opts ...func(*OutboundDealQuery)) *OutboundDealQuery {
	query := &OutboundDealQuery{config: odq.config}
	for _, opt := range opts {
		opt(query)
	}
	odq.withChildren = query
	return odq
}

//  WithTransaction tells the query-builder to eager-loads the nodes that are connected to
// the "transaction" edge. The optional arguments used to configure the query builder of the edge.
func (odq *OutboundDealQuery) WithTransaction(opts ...func(*OutboundTransactionQuery)) *OutboundDealQuery {
	query := &OutboundTransactionQuery{config: odq.config}
	for _, opt := range opts {
		opt(query)
	}
	odq.withTransaction = query
	return odq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Quantity uint `json:"quantity,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OutboundDeal.Query().
//		GroupBy(outbounddeal.FieldQuantity).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (odq *OutboundDealQuery) GroupBy(field string, fields ...string) *OutboundDealGroupBy {
	group := &OutboundDealGroupBy{config: odq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := odq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return odq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Quantity uint `json:"quantity,omitempty"`
//	}
//
//	client.OutboundDeal.Query().
//		Select(outbounddeal.FieldQuantity).
//		Scan(ctx, &v)
//
func (odq *OutboundDealQuery) Select(field string, fields ...string) *OutboundDealSelect {
	selector := &OutboundDealSelect{config: odq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := odq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return odq.sqlQuery(), nil
	}
	return selector
}

func (odq *OutboundDealQuery) prepareQuery(ctx context.Context) error {
	if odq.path != nil {
		prev, err := odq.path(ctx)
		if err != nil {
			return err
		}
		odq.sql = prev
	}
	return nil
}

func (odq *OutboundDealQuery) sqlAll(ctx context.Context) ([]*OutboundDeal, error) {
	var (
		nodes       = []*OutboundDeal{}
		withFKs     = odq.withFKs
		_spec       = odq.querySpec()
		loadedTypes = [4]bool{
			odq.withVariant != nil,
			odq.withParent != nil,
			odq.withChildren != nil,
			odq.withTransaction != nil,
		}
	)
	if odq.withVariant != nil || odq.withParent != nil || odq.withTransaction != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, outbounddeal.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &OutboundDeal{config: odq.config}
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
	if err := sqlgraph.QueryNodes(ctx, odq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := odq.withVariant; query != nil {
		ids := make([]uint64, 0, len(nodes))
		nodeids := make(map[uint64][]*OutboundDeal)
		for i := range nodes {
			if fk := nodes[i].outbound_deal_variant; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(variant.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "outbound_deal_variant" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Variant = n
			}
		}
	}

	if query := odq.withParent; query != nil {
		ids := make([]uint64, 0, len(nodes))
		nodeids := make(map[uint64][]*OutboundDeal)
		for i := range nodes {
			if fk := nodes[i].outbound_deal_children; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(outbounddeal.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "outbound_deal_children" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := odq.withChildren; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uint64]*OutboundDeal)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Children = []*OutboundDeal{}
		}
		query.withFKs = true
		query.Where(predicate.OutboundDeal(func(s *sql.Selector) {
			s.Where(sql.InValues(outbounddeal.ChildrenColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.outbound_deal_children
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "outbound_deal_children" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "outbound_deal_children" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Children = append(node.Edges.Children, n)
		}
	}

	if query := odq.withTransaction; query != nil {
		ids := make([]uint64, 0, len(nodes))
		nodeids := make(map[uint64][]*OutboundDeal)
		for i := range nodes {
			if fk := nodes[i].outbound_transaction_deals; fk != nil {
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
				return nil, fmt.Errorf(`unexpected foreign-key "outbound_transaction_deals" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Transaction = n
			}
		}
	}

	return nodes, nil
}

func (odq *OutboundDealQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := odq.querySpec()
	return sqlgraph.CountNodes(ctx, odq.driver, _spec)
}

func (odq *OutboundDealQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := odq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (odq *OutboundDealQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outbounddeal.Table,
			Columns: outbounddeal.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: outbounddeal.FieldID,
			},
		},
		From:   odq.sql,
		Unique: true,
	}
	if ps := odq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := odq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := odq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := odq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, outbounddeal.ValidColumn)
			}
		}
	}
	return _spec
}

func (odq *OutboundDealQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(odq.driver.Dialect())
	t1 := builder.Table(outbounddeal.Table)
	selector := builder.Select(t1.Columns(outbounddeal.Columns...)...).From(t1)
	if odq.sql != nil {
		selector = odq.sql
		selector.Select(selector.Columns(outbounddeal.Columns...)...)
	}
	for _, p := range odq.predicates {
		p(selector)
	}
	for _, p := range odq.order {
		p(selector, outbounddeal.ValidColumn)
	}
	if offset := odq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := odq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OutboundDealGroupBy is the builder for group-by OutboundDeal entities.
type OutboundDealGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (odgb *OutboundDealGroupBy) Aggregate(fns ...AggregateFunc) *OutboundDealGroupBy {
	odgb.fns = append(odgb.fns, fns...)
	return odgb
}

// Scan applies the group-by query and scan the result into the given value.
func (odgb *OutboundDealGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := odgb.path(ctx)
	if err != nil {
		return err
	}
	odgb.sql = query
	return odgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (odgb *OutboundDealGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := odgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (odgb *OutboundDealGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(odgb.fields) > 1 {
		return nil, errors.New("ent: OutboundDealGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := odgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (odgb *OutboundDealGroupBy) StringsX(ctx context.Context) []string {
	v, err := odgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (odgb *OutboundDealGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = odgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{outbounddeal.Label}
	default:
		err = fmt.Errorf("ent: OutboundDealGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (odgb *OutboundDealGroupBy) StringX(ctx context.Context) string {
	v, err := odgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (odgb *OutboundDealGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(odgb.fields) > 1 {
		return nil, errors.New("ent: OutboundDealGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := odgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (odgb *OutboundDealGroupBy) IntsX(ctx context.Context) []int {
	v, err := odgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (odgb *OutboundDealGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = odgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{outbounddeal.Label}
	default:
		err = fmt.Errorf("ent: OutboundDealGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (odgb *OutboundDealGroupBy) IntX(ctx context.Context) int {
	v, err := odgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (odgb *OutboundDealGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(odgb.fields) > 1 {
		return nil, errors.New("ent: OutboundDealGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := odgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (odgb *OutboundDealGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := odgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (odgb *OutboundDealGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = odgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{outbounddeal.Label}
	default:
		err = fmt.Errorf("ent: OutboundDealGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (odgb *OutboundDealGroupBy) Float64X(ctx context.Context) float64 {
	v, err := odgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (odgb *OutboundDealGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(odgb.fields) > 1 {
		return nil, errors.New("ent: OutboundDealGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := odgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (odgb *OutboundDealGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := odgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (odgb *OutboundDealGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = odgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{outbounddeal.Label}
	default:
		err = fmt.Errorf("ent: OutboundDealGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (odgb *OutboundDealGroupBy) BoolX(ctx context.Context) bool {
	v, err := odgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (odgb *OutboundDealGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range odgb.fields {
		if !outbounddeal.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := odgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := odgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (odgb *OutboundDealGroupBy) sqlQuery() *sql.Selector {
	selector := odgb.sql
	columns := make([]string, 0, len(odgb.fields)+len(odgb.fns))
	columns = append(columns, odgb.fields...)
	for _, fn := range odgb.fns {
		columns = append(columns, fn(selector, outbounddeal.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(odgb.fields...)
}

// OutboundDealSelect is the builder for select fields of OutboundDeal entities.
type OutboundDealSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ods *OutboundDealSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ods.path(ctx)
	if err != nil {
		return err
	}
	ods.sql = query
	return ods.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ods *OutboundDealSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ods.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ods *OutboundDealSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ods.fields) > 1 {
		return nil, errors.New("ent: OutboundDealSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ods.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ods *OutboundDealSelect) StringsX(ctx context.Context) []string {
	v, err := ods.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ods *OutboundDealSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ods.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{outbounddeal.Label}
	default:
		err = fmt.Errorf("ent: OutboundDealSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ods *OutboundDealSelect) StringX(ctx context.Context) string {
	v, err := ods.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ods *OutboundDealSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ods.fields) > 1 {
		return nil, errors.New("ent: OutboundDealSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ods.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ods *OutboundDealSelect) IntsX(ctx context.Context) []int {
	v, err := ods.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ods *OutboundDealSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ods.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{outbounddeal.Label}
	default:
		err = fmt.Errorf("ent: OutboundDealSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ods *OutboundDealSelect) IntX(ctx context.Context) int {
	v, err := ods.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ods *OutboundDealSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ods.fields) > 1 {
		return nil, errors.New("ent: OutboundDealSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ods.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ods *OutboundDealSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ods.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ods *OutboundDealSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ods.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{outbounddeal.Label}
	default:
		err = fmt.Errorf("ent: OutboundDealSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ods *OutboundDealSelect) Float64X(ctx context.Context) float64 {
	v, err := ods.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ods *OutboundDealSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ods.fields) > 1 {
		return nil, errors.New("ent: OutboundDealSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ods.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ods *OutboundDealSelect) BoolsX(ctx context.Context) []bool {
	v, err := ods.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ods *OutboundDealSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ods.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{outbounddeal.Label}
	default:
		err = fmt.Errorf("ent: OutboundDealSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ods *OutboundDealSelect) BoolX(ctx context.Context) bool {
	v, err := ods.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ods *OutboundDealSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ods.fields {
		if !outbounddeal.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := ods.sqlQuery().Query()
	if err := ods.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ods *OutboundDealSelect) sqlQuery() sql.Querier {
	selector := ods.sql
	selector.Select(selector.Columns(ods.fields...)...)
	return selector
}
