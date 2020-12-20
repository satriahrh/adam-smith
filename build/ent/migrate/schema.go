// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// BrandsColumns holds the columns for the "brands" table.
	BrandsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "code", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
	}
	// BrandsTable holds the schema information for the "brands" table.
	BrandsTable = &schema.Table{
		Name:        "brands",
		Columns:     BrandsColumns,
		PrimaryKey:  []*schema.Column{BrandsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// OutboundDealsColumns holds the columns for the "outbound_deals" table.
	OutboundDealsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "quantity", Type: field.TypeUint},
		{Name: "amount", Type: field.TypeUint},
		{Name: "outbound_deal_variant", Type: field.TypeUint64, Nullable: true},
		{Name: "outbound_deal_children", Type: field.TypeUint64, Nullable: true},
		{Name: "outbound_transaction_deals", Type: field.TypeUint64, Nullable: true},
	}
	// OutboundDealsTable holds the schema information for the "outbound_deals" table.
	OutboundDealsTable = &schema.Table{
		Name:       "outbound_deals",
		Columns:    OutboundDealsColumns,
		PrimaryKey: []*schema.Column{OutboundDealsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "outbound_deals_variants_variant",
				Columns: []*schema.Column{OutboundDealsColumns[3]},

				RefColumns: []*schema.Column{VariantsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "outbound_deals_outbound_deals_children",
				Columns: []*schema.Column{OutboundDealsColumns[4]},

				RefColumns: []*schema.Column{OutboundDealsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "outbound_deals_outbound_transactions_deals",
				Columns: []*schema.Column{OutboundDealsColumns[5]},

				RefColumns: []*schema.Column{OutboundTransactionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OutboundShippingsColumns holds the columns for the "outbound_shippings" table.
	OutboundShippingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "courier", Type: field.TypeEnum, Enums: []string{"self", "jnt", "shopee_express", "jne", "sicepat", "ninja_express", "anteraja", "pos_indonesia"}},
		{Name: "courier_tracking_code", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"pick", "drop"}},
		{Name: "state", Type: field.TypeEnum, Enums: []string{"requested", "sent", "delivered", "finish"}},
		{Name: "consignee", Type: field.TypeString},
		{Name: "consignee_phone", Type: field.TypeString},
		{Name: "address", Type: field.TypeString},
		{Name: "postal_code", Type: field.TypeUint},
		{Name: "cost", Type: field.TypeUint},
		{Name: "outbound_transaction_shipping", Type: field.TypeUint64, Unique: true, Nullable: true},
	}
	// OutboundShippingsTable holds the schema information for the "outbound_shippings" table.
	OutboundShippingsTable = &schema.Table{
		Name:       "outbound_shippings",
		Columns:    OutboundShippingsColumns,
		PrimaryKey: []*schema.Column{OutboundShippingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "outbound_shippings_outbound_transactions_shipping",
				Columns: []*schema.Column{OutboundShippingsColumns[10]},

				RefColumns: []*schema.Column{OutboundTransactionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OutboundTransactionsColumns holds the columns for the "outbound_transactions" table.
	OutboundTransactionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "channel", Type: field.TypeEnum, Enums: []string{"shopee", "direct"}},
		{Name: "invoice", Type: field.TypeJSON},
		{Name: "benefit", Type: field.TypeJSON},
		{Name: "cost", Type: field.TypeJSON},
		{Name: "amount", Type: field.TypeUint},
	}
	// OutboundTransactionsTable holds the schema information for the "outbound_transactions" table.
	OutboundTransactionsTable = &schema.Table{
		Name:        "outbound_transactions",
		Columns:     OutboundTransactionsColumns,
		PrimaryKey:  []*schema.Column{OutboundTransactionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ProductsColumns holds the columns for the "products" table.
	ProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "sku", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "descriptions", Type: field.TypeJSON},
		{Name: "images", Type: field.TypeJSON},
		{Name: "marketplaces", Type: field.TypeJSON},
		{Name: "brand_products", Type: field.TypeUint64, Nullable: true},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:       "products",
		Columns:    ProductsColumns,
		PrimaryKey: []*schema.Column{ProductsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "products_brands_products",
				Columns: []*schema.Column{ProductsColumns[6]},

				RefColumns: []*schema.Column{BrandsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// VariantsColumns holds the columns for the "variants" table.
	VariantsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "images", Type: field.TypeJSON, Nullable: true},
		{Name: "stock", Type: field.TypeUint8},
		{Name: "price", Type: field.TypeUint},
		{Name: "product_variants", Type: field.TypeUint64, Nullable: true},
		{Name: "variant_children", Type: field.TypeUint64, Nullable: true},
		{Name: "variant_variation", Type: field.TypeUint64, Nullable: true},
	}
	// VariantsTable holds the schema information for the "variants" table.
	VariantsTable = &schema.Table{
		Name:       "variants",
		Columns:    VariantsColumns,
		PrimaryKey: []*schema.Column{VariantsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "variants_products_variants",
				Columns: []*schema.Column{VariantsColumns[4]},

				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "variants_variants_children",
				Columns: []*schema.Column{VariantsColumns[5]},

				RefColumns: []*schema.Column{VariantsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "variants_variations_variation",
				Columns: []*schema.Column{VariantsColumns[6]},

				RefColumns: []*schema.Column{VariationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// VariationsColumns holds the columns for the "variations" table.
	VariationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"color", "size"}},
		{Name: "value", Type: field.TypeString},
	}
	// VariationsTable holds the schema information for the "variations" table.
	VariationsTable = &schema.Table{
		Name:        "variations",
		Columns:     VariationsColumns,
		PrimaryKey:  []*schema.Column{VariationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "variation_value",
				Unique:  false,
				Columns: []*schema.Column{VariationsColumns[2]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BrandsTable,
		OutboundDealsTable,
		OutboundShippingsTable,
		OutboundTransactionsTable,
		ProductsTable,
		VariantsTable,
		VariationsTable,
	}
)

func init() {
	OutboundDealsTable.ForeignKeys[0].RefTable = VariantsTable
	OutboundDealsTable.ForeignKeys[1].RefTable = OutboundDealsTable
	OutboundDealsTable.ForeignKeys[2].RefTable = OutboundTransactionsTable
	OutboundShippingsTable.ForeignKeys[0].RefTable = OutboundTransactionsTable
	ProductsTable.ForeignKeys[0].RefTable = BrandsTable
	VariantsTable.ForeignKeys[0].RefTable = ProductsTable
	VariantsTable.ForeignKeys[1].RefTable = VariantsTable
	VariantsTable.ForeignKeys[2].RefTable = VariationsTable
}
