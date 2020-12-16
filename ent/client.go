// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/satriahrh/adam-smith/ent/migrate"

	"github.com/satriahrh/adam-smith/ent/brand"
	"github.com/satriahrh/adam-smith/ent/outbounddeal"
	"github.com/satriahrh/adam-smith/ent/outboundshipping"
	"github.com/satriahrh/adam-smith/ent/outboundtransaction"
	"github.com/satriahrh/adam-smith/ent/product"
	"github.com/satriahrh/adam-smith/ent/variant"
	"github.com/satriahrh/adam-smith/ent/variation"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Brand is the client for interacting with the Brand builders.
	Brand *BrandClient
	// OutboundDeal is the client for interacting with the OutboundDeal builders.
	OutboundDeal *OutboundDealClient
	// OutboundShipping is the client for interacting with the OutboundShipping builders.
	OutboundShipping *OutboundShippingClient
	// OutboundTransaction is the client for interacting with the OutboundTransaction builders.
	OutboundTransaction *OutboundTransactionClient
	// Product is the client for interacting with the Product builders.
	Product *ProductClient
	// Variant is the client for interacting with the Variant builders.
	Variant *VariantClient
	// Variation is the client for interacting with the Variation builders.
	Variation *VariationClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Brand = NewBrandClient(c.config)
	c.OutboundDeal = NewOutboundDealClient(c.config)
	c.OutboundShipping = NewOutboundShippingClient(c.config)
	c.OutboundTransaction = NewOutboundTransactionClient(c.config)
	c.Product = NewProductClient(c.config)
	c.Variant = NewVariantClient(c.config)
	c.Variation = NewVariationClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:                 ctx,
		config:              cfg,
		Brand:               NewBrandClient(cfg),
		OutboundDeal:        NewOutboundDealClient(cfg),
		OutboundShipping:    NewOutboundShippingClient(cfg),
		OutboundTransaction: NewOutboundTransactionClient(cfg),
		Product:             NewProductClient(cfg),
		Variant:             NewVariantClient(cfg),
		Variation:           NewVariationClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:              cfg,
		Brand:               NewBrandClient(cfg),
		OutboundDeal:        NewOutboundDealClient(cfg),
		OutboundShipping:    NewOutboundShippingClient(cfg),
		OutboundTransaction: NewOutboundTransactionClient(cfg),
		Product:             NewProductClient(cfg),
		Variant:             NewVariantClient(cfg),
		Variation:           NewVariationClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Brand.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Brand.Use(hooks...)
	c.OutboundDeal.Use(hooks...)
	c.OutboundShipping.Use(hooks...)
	c.OutboundTransaction.Use(hooks...)
	c.Product.Use(hooks...)
	c.Variant.Use(hooks...)
	c.Variation.Use(hooks...)
}

// BrandClient is a client for the Brand schema.
type BrandClient struct {
	config
}

// NewBrandClient returns a client for the Brand from the given config.
func NewBrandClient(c config) *BrandClient {
	return &BrandClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `brand.Hooks(f(g(h())))`.
func (c *BrandClient) Use(hooks ...Hook) {
	c.hooks.Brand = append(c.hooks.Brand, hooks...)
}

// Create returns a create builder for Brand.
func (c *BrandClient) Create() *BrandCreate {
	mutation := newBrandMutation(c.config, OpCreate)
	return &BrandCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Brand entities.
func (c *BrandClient) CreateBulk(builders ...*BrandCreate) *BrandCreateBulk {
	return &BrandCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Brand.
func (c *BrandClient) Update() *BrandUpdate {
	mutation := newBrandMutation(c.config, OpUpdate)
	return &BrandUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BrandClient) UpdateOne(b *Brand) *BrandUpdateOne {
	mutation := newBrandMutation(c.config, OpUpdateOne, withBrand(b))
	return &BrandUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BrandClient) UpdateOneID(id int) *BrandUpdateOne {
	mutation := newBrandMutation(c.config, OpUpdateOne, withBrandID(id))
	return &BrandUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Brand.
func (c *BrandClient) Delete() *BrandDelete {
	mutation := newBrandMutation(c.config, OpDelete)
	return &BrandDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *BrandClient) DeleteOne(b *Brand) *BrandDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *BrandClient) DeleteOneID(id int) *BrandDeleteOne {
	builder := c.Delete().Where(brand.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BrandDeleteOne{builder}
}

// Query returns a query builder for Brand.
func (c *BrandClient) Query() *BrandQuery {
	return &BrandQuery{config: c.config}
}

// Get returns a Brand entity by its id.
func (c *BrandClient) Get(ctx context.Context, id int) (*Brand, error) {
	return c.Query().Where(brand.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BrandClient) GetX(ctx context.Context, id int) *Brand {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryProducts queries the products edge of a Brand.
func (c *BrandClient) QueryProducts(b *Brand) *ProductQuery {
	query := &ProductQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(brand.Table, brand.FieldID, id),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, brand.ProductsTable, brand.ProductsColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BrandClient) Hooks() []Hook {
	return c.hooks.Brand
}

// OutboundDealClient is a client for the OutboundDeal schema.
type OutboundDealClient struct {
	config
}

// NewOutboundDealClient returns a client for the OutboundDeal from the given config.
func NewOutboundDealClient(c config) *OutboundDealClient {
	return &OutboundDealClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `outbounddeal.Hooks(f(g(h())))`.
func (c *OutboundDealClient) Use(hooks ...Hook) {
	c.hooks.OutboundDeal = append(c.hooks.OutboundDeal, hooks...)
}

// Create returns a create builder for OutboundDeal.
func (c *OutboundDealClient) Create() *OutboundDealCreate {
	mutation := newOutboundDealMutation(c.config, OpCreate)
	return &OutboundDealCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of OutboundDeal entities.
func (c *OutboundDealClient) CreateBulk(builders ...*OutboundDealCreate) *OutboundDealCreateBulk {
	return &OutboundDealCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OutboundDeal.
func (c *OutboundDealClient) Update() *OutboundDealUpdate {
	mutation := newOutboundDealMutation(c.config, OpUpdate)
	return &OutboundDealUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OutboundDealClient) UpdateOne(od *OutboundDeal) *OutboundDealUpdateOne {
	mutation := newOutboundDealMutation(c.config, OpUpdateOne, withOutboundDeal(od))
	return &OutboundDealUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OutboundDealClient) UpdateOneID(id int) *OutboundDealUpdateOne {
	mutation := newOutboundDealMutation(c.config, OpUpdateOne, withOutboundDealID(id))
	return &OutboundDealUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OutboundDeal.
func (c *OutboundDealClient) Delete() *OutboundDealDelete {
	mutation := newOutboundDealMutation(c.config, OpDelete)
	return &OutboundDealDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *OutboundDealClient) DeleteOne(od *OutboundDeal) *OutboundDealDeleteOne {
	return c.DeleteOneID(od.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *OutboundDealClient) DeleteOneID(id int) *OutboundDealDeleteOne {
	builder := c.Delete().Where(outbounddeal.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OutboundDealDeleteOne{builder}
}

// Query returns a query builder for OutboundDeal.
func (c *OutboundDealClient) Query() *OutboundDealQuery {
	return &OutboundDealQuery{config: c.config}
}

// Get returns a OutboundDeal entity by its id.
func (c *OutboundDealClient) Get(ctx context.Context, id int) (*OutboundDeal, error) {
	return c.Query().Where(outbounddeal.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OutboundDealClient) GetX(ctx context.Context, id int) *OutboundDeal {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryVariant queries the variant edge of a OutboundDeal.
func (c *OutboundDealClient) QueryVariant(od *OutboundDeal) *VariationQuery {
	query := &VariationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := od.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(outbounddeal.Table, outbounddeal.FieldID, id),
			sqlgraph.To(variation.Table, variation.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, outbounddeal.VariantTable, outbounddeal.VariantPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(od.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTransaction queries the transaction edge of a OutboundDeal.
func (c *OutboundDealClient) QueryTransaction(od *OutboundDeal) *OutboundTransactionQuery {
	query := &OutboundTransactionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := od.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(outbounddeal.Table, outbounddeal.FieldID, id),
			sqlgraph.To(outboundtransaction.Table, outboundtransaction.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, outbounddeal.TransactionTable, outbounddeal.TransactionPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(od.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OutboundDealClient) Hooks() []Hook {
	return c.hooks.OutboundDeal
}

// OutboundShippingClient is a client for the OutboundShipping schema.
type OutboundShippingClient struct {
	config
}

// NewOutboundShippingClient returns a client for the OutboundShipping from the given config.
func NewOutboundShippingClient(c config) *OutboundShippingClient {
	return &OutboundShippingClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `outboundshipping.Hooks(f(g(h())))`.
func (c *OutboundShippingClient) Use(hooks ...Hook) {
	c.hooks.OutboundShipping = append(c.hooks.OutboundShipping, hooks...)
}

// Create returns a create builder for OutboundShipping.
func (c *OutboundShippingClient) Create() *OutboundShippingCreate {
	mutation := newOutboundShippingMutation(c.config, OpCreate)
	return &OutboundShippingCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of OutboundShipping entities.
func (c *OutboundShippingClient) CreateBulk(builders ...*OutboundShippingCreate) *OutboundShippingCreateBulk {
	return &OutboundShippingCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OutboundShipping.
func (c *OutboundShippingClient) Update() *OutboundShippingUpdate {
	mutation := newOutboundShippingMutation(c.config, OpUpdate)
	return &OutboundShippingUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OutboundShippingClient) UpdateOne(os *OutboundShipping) *OutboundShippingUpdateOne {
	mutation := newOutboundShippingMutation(c.config, OpUpdateOne, withOutboundShipping(os))
	return &OutboundShippingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OutboundShippingClient) UpdateOneID(id int) *OutboundShippingUpdateOne {
	mutation := newOutboundShippingMutation(c.config, OpUpdateOne, withOutboundShippingID(id))
	return &OutboundShippingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OutboundShipping.
func (c *OutboundShippingClient) Delete() *OutboundShippingDelete {
	mutation := newOutboundShippingMutation(c.config, OpDelete)
	return &OutboundShippingDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *OutboundShippingClient) DeleteOne(os *OutboundShipping) *OutboundShippingDeleteOne {
	return c.DeleteOneID(os.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *OutboundShippingClient) DeleteOneID(id int) *OutboundShippingDeleteOne {
	builder := c.Delete().Where(outboundshipping.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OutboundShippingDeleteOne{builder}
}

// Query returns a query builder for OutboundShipping.
func (c *OutboundShippingClient) Query() *OutboundShippingQuery {
	return &OutboundShippingQuery{config: c.config}
}

// Get returns a OutboundShipping entity by its id.
func (c *OutboundShippingClient) Get(ctx context.Context, id int) (*OutboundShipping, error) {
	return c.Query().Where(outboundshipping.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OutboundShippingClient) GetX(ctx context.Context, id int) *OutboundShipping {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTransaction queries the transaction edge of a OutboundShipping.
func (c *OutboundShippingClient) QueryTransaction(os *OutboundShipping) *OutboundTransactionQuery {
	query := &OutboundTransactionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := os.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(outboundshipping.Table, outboundshipping.FieldID, id),
			sqlgraph.To(outboundtransaction.Table, outboundtransaction.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, outboundshipping.TransactionTable, outboundshipping.TransactionPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(os.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OutboundShippingClient) Hooks() []Hook {
	return c.hooks.OutboundShipping
}

// OutboundTransactionClient is a client for the OutboundTransaction schema.
type OutboundTransactionClient struct {
	config
}

// NewOutboundTransactionClient returns a client for the OutboundTransaction from the given config.
func NewOutboundTransactionClient(c config) *OutboundTransactionClient {
	return &OutboundTransactionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `outboundtransaction.Hooks(f(g(h())))`.
func (c *OutboundTransactionClient) Use(hooks ...Hook) {
	c.hooks.OutboundTransaction = append(c.hooks.OutboundTransaction, hooks...)
}

// Create returns a create builder for OutboundTransaction.
func (c *OutboundTransactionClient) Create() *OutboundTransactionCreate {
	mutation := newOutboundTransactionMutation(c.config, OpCreate)
	return &OutboundTransactionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of OutboundTransaction entities.
func (c *OutboundTransactionClient) CreateBulk(builders ...*OutboundTransactionCreate) *OutboundTransactionCreateBulk {
	return &OutboundTransactionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OutboundTransaction.
func (c *OutboundTransactionClient) Update() *OutboundTransactionUpdate {
	mutation := newOutboundTransactionMutation(c.config, OpUpdate)
	return &OutboundTransactionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OutboundTransactionClient) UpdateOne(ot *OutboundTransaction) *OutboundTransactionUpdateOne {
	mutation := newOutboundTransactionMutation(c.config, OpUpdateOne, withOutboundTransaction(ot))
	return &OutboundTransactionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OutboundTransactionClient) UpdateOneID(id int) *OutboundTransactionUpdateOne {
	mutation := newOutboundTransactionMutation(c.config, OpUpdateOne, withOutboundTransactionID(id))
	return &OutboundTransactionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OutboundTransaction.
func (c *OutboundTransactionClient) Delete() *OutboundTransactionDelete {
	mutation := newOutboundTransactionMutation(c.config, OpDelete)
	return &OutboundTransactionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *OutboundTransactionClient) DeleteOne(ot *OutboundTransaction) *OutboundTransactionDeleteOne {
	return c.DeleteOneID(ot.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *OutboundTransactionClient) DeleteOneID(id int) *OutboundTransactionDeleteOne {
	builder := c.Delete().Where(outboundtransaction.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OutboundTransactionDeleteOne{builder}
}

// Query returns a query builder for OutboundTransaction.
func (c *OutboundTransactionClient) Query() *OutboundTransactionQuery {
	return &OutboundTransactionQuery{config: c.config}
}

// Get returns a OutboundTransaction entity by its id.
func (c *OutboundTransactionClient) Get(ctx context.Context, id int) (*OutboundTransaction, error) {
	return c.Query().Where(outboundtransaction.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OutboundTransactionClient) GetX(ctx context.Context, id int) *OutboundTransaction {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryShipping queries the shipping edge of a OutboundTransaction.
func (c *OutboundTransactionClient) QueryShipping(ot *OutboundTransaction) *OutboundShippingQuery {
	query := &OutboundShippingQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ot.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(outboundtransaction.Table, outboundtransaction.FieldID, id),
			sqlgraph.To(outboundshipping.Table, outboundshipping.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, outboundtransaction.ShippingTable, outboundtransaction.ShippingPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(ot.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDeals queries the deals edge of a OutboundTransaction.
func (c *OutboundTransactionClient) QueryDeals(ot *OutboundTransaction) *OutboundDealQuery {
	query := &OutboundDealQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ot.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(outboundtransaction.Table, outboundtransaction.FieldID, id),
			sqlgraph.To(outbounddeal.Table, outbounddeal.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, outboundtransaction.DealsTable, outboundtransaction.DealsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(ot.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OutboundTransactionClient) Hooks() []Hook {
	return c.hooks.OutboundTransaction
}

// ProductClient is a client for the Product schema.
type ProductClient struct {
	config
}

// NewProductClient returns a client for the Product from the given config.
func NewProductClient(c config) *ProductClient {
	return &ProductClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `product.Hooks(f(g(h())))`.
func (c *ProductClient) Use(hooks ...Hook) {
	c.hooks.Product = append(c.hooks.Product, hooks...)
}

// Create returns a create builder for Product.
func (c *ProductClient) Create() *ProductCreate {
	mutation := newProductMutation(c.config, OpCreate)
	return &ProductCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Product entities.
func (c *ProductClient) CreateBulk(builders ...*ProductCreate) *ProductCreateBulk {
	return &ProductCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Product.
func (c *ProductClient) Update() *ProductUpdate {
	mutation := newProductMutation(c.config, OpUpdate)
	return &ProductUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProductClient) UpdateOne(pr *Product) *ProductUpdateOne {
	mutation := newProductMutation(c.config, OpUpdateOne, withProduct(pr))
	return &ProductUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProductClient) UpdateOneID(id int) *ProductUpdateOne {
	mutation := newProductMutation(c.config, OpUpdateOne, withProductID(id))
	return &ProductUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Product.
func (c *ProductClient) Delete() *ProductDelete {
	mutation := newProductMutation(c.config, OpDelete)
	return &ProductDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ProductClient) DeleteOne(pr *Product) *ProductDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ProductClient) DeleteOneID(id int) *ProductDeleteOne {
	builder := c.Delete().Where(product.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProductDeleteOne{builder}
}

// Query returns a query builder for Product.
func (c *ProductClient) Query() *ProductQuery {
	return &ProductQuery{config: c.config}
}

// Get returns a Product entity by its id.
func (c *ProductClient) Get(ctx context.Context, id int) (*Product, error) {
	return c.Query().Where(product.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProductClient) GetX(ctx context.Context, id int) *Product {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryVariations queries the variations edge of a Product.
func (c *ProductClient) QueryVariations(pr *Product) *VariationQuery {
	query := &VariationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(product.Table, product.FieldID, id),
			sqlgraph.To(variation.Table, variation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, product.VariationsTable, product.VariationsColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBrand queries the brand edge of a Product.
func (c *ProductClient) QueryBrand(pr *Product) *BrandQuery {
	query := &BrandQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(product.Table, product.FieldID, id),
			sqlgraph.To(brand.Table, brand.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, product.BrandTable, product.BrandColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ProductClient) Hooks() []Hook {
	return c.hooks.Product
}

// VariantClient is a client for the Variant schema.
type VariantClient struct {
	config
}

// NewVariantClient returns a client for the Variant from the given config.
func NewVariantClient(c config) *VariantClient {
	return &VariantClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `variant.Hooks(f(g(h())))`.
func (c *VariantClient) Use(hooks ...Hook) {
	c.hooks.Variant = append(c.hooks.Variant, hooks...)
}

// Create returns a create builder for Variant.
func (c *VariantClient) Create() *VariantCreate {
	mutation := newVariantMutation(c.config, OpCreate)
	return &VariantCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Variant entities.
func (c *VariantClient) CreateBulk(builders ...*VariantCreate) *VariantCreateBulk {
	return &VariantCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Variant.
func (c *VariantClient) Update() *VariantUpdate {
	mutation := newVariantMutation(c.config, OpUpdate)
	return &VariantUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VariantClient) UpdateOne(v *Variant) *VariantUpdateOne {
	mutation := newVariantMutation(c.config, OpUpdateOne, withVariant(v))
	return &VariantUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VariantClient) UpdateOneID(id int) *VariantUpdateOne {
	mutation := newVariantMutation(c.config, OpUpdateOne, withVariantID(id))
	return &VariantUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Variant.
func (c *VariantClient) Delete() *VariantDelete {
	mutation := newVariantMutation(c.config, OpDelete)
	return &VariantDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *VariantClient) DeleteOne(v *Variant) *VariantDeleteOne {
	return c.DeleteOneID(v.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *VariantClient) DeleteOneID(id int) *VariantDeleteOne {
	builder := c.Delete().Where(variant.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VariantDeleteOne{builder}
}

// Query returns a query builder for Variant.
func (c *VariantClient) Query() *VariantQuery {
	return &VariantQuery{config: c.config}
}

// Get returns a Variant entity by its id.
func (c *VariantClient) Get(ctx context.Context, id int) (*Variant, error) {
	return c.Query().Where(variant.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VariantClient) GetX(ctx context.Context, id int) *Variant {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryVariantUses queries the variant_uses edge of a Variant.
func (c *VariantClient) QueryVariantUses(v *Variant) *VariationQuery {
	query := &VariationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(variant.Table, variant.FieldID, id),
			sqlgraph.To(variation.Table, variation.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, variant.VariantUsesTable, variant.VariantUsesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *VariantClient) Hooks() []Hook {
	return c.hooks.Variant
}

// VariationClient is a client for the Variation schema.
type VariationClient struct {
	config
}

// NewVariationClient returns a client for the Variation from the given config.
func NewVariationClient(c config) *VariationClient {
	return &VariationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `variation.Hooks(f(g(h())))`.
func (c *VariationClient) Use(hooks ...Hook) {
	c.hooks.Variation = append(c.hooks.Variation, hooks...)
}

// Create returns a create builder for Variation.
func (c *VariationClient) Create() *VariationCreate {
	mutation := newVariationMutation(c.config, OpCreate)
	return &VariationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Variation entities.
func (c *VariationClient) CreateBulk(builders ...*VariationCreate) *VariationCreateBulk {
	return &VariationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Variation.
func (c *VariationClient) Update() *VariationUpdate {
	mutation := newVariationMutation(c.config, OpUpdate)
	return &VariationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VariationClient) UpdateOne(v *Variation) *VariationUpdateOne {
	mutation := newVariationMutation(c.config, OpUpdateOne, withVariation(v))
	return &VariationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VariationClient) UpdateOneID(id int) *VariationUpdateOne {
	mutation := newVariationMutation(c.config, OpUpdateOne, withVariationID(id))
	return &VariationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Variation.
func (c *VariationClient) Delete() *VariationDelete {
	mutation := newVariationMutation(c.config, OpDelete)
	return &VariationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *VariationClient) DeleteOne(v *Variation) *VariationDeleteOne {
	return c.DeleteOneID(v.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *VariationClient) DeleteOneID(id int) *VariationDeleteOne {
	builder := c.Delete().Where(variation.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VariationDeleteOne{builder}
}

// Query returns a query builder for Variation.
func (c *VariationClient) Query() *VariationQuery {
	return &VariationQuery{config: c.config}
}

// Get returns a Variation entity by its id.
func (c *VariationClient) Get(ctx context.Context, id int) (*Variation, error) {
	return c.Query().Where(variation.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VariationClient) GetX(ctx context.Context, id int) *Variation {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a Variation.
func (c *VariationClient) QueryParent(v *Variation) *VariationQuery {
	query := &VariationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(variation.Table, variation.FieldID, id),
			sqlgraph.To(variation.Table, variation.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, variation.ParentTable, variation.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryChildren queries the children edge of a Variation.
func (c *VariationClient) QueryChildren(v *Variation) *VariationQuery {
	query := &VariationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(variation.Table, variation.FieldID, id),
			sqlgraph.To(variation.Table, variation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, variation.ChildrenTable, variation.ChildrenColumn),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryProduct queries the product edge of a Variation.
func (c *VariationClient) QueryProduct(v *Variation) *ProductQuery {
	query := &ProductQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(variation.Table, variation.FieldID, id),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, variation.ProductTable, variation.ProductColumn),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryVariant queries the variant edge of a Variation.
func (c *VariationClient) QueryVariant(v *Variation) *VariantQuery {
	query := &VariantQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(variation.Table, variation.FieldID, id),
			sqlgraph.To(variant.Table, variant.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, variation.VariantTable, variation.VariantPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOutboundDeals queries the outbound_deals edge of a Variation.
func (c *VariationClient) QueryOutboundDeals(v *Variation) *OutboundDealQuery {
	query := &OutboundDealQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(variation.Table, variation.FieldID, id),
			sqlgraph.To(outbounddeal.Table, outbounddeal.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, variation.OutboundDealsTable, variation.OutboundDealsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *VariationClient) Hooks() []Hook {
	return c.hooks.Variation
}
