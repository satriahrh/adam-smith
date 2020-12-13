package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// ProductDescription description of a product
type ProductDescription struct {
	Heading string `json:"heading"`
	Body    string `json:"body"`
}

// ProductMarketplace marketplace details of the product being sold
type ProductMarketplace struct {
	Tokopedia string `json:"tokopedia"`
	Shopee    string `json:"shopee"`
	Bukalapak string `json:"bukalapak"`
}

// ProductImages detail images of the product
type ProductImages struct {
	Thumbnail string   `json:"thumbnail"`
	Displays  []string `json:"displays"`
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.JSON("descriptions", []ProductDescription{}),
		field.JSON("images", ProductImages{}),
		field.JSON("marketplaces", ProductMarketplace{}),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("brand", Brand.Type).
			Ref("products"),
		edge.To("variations", ProductVariant.Type),
	}
}
