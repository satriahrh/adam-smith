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

// ProductMarketplaces marketplace details of the product being sold
type ProductMarketplaces struct {
	Tokopedia string `json:"tokopedia,omitempty"`
	Shopee    string `json:"shopee,omitempty"`
	Bukalapak string `json:"bukalapak,omitempty"`
}

// ProductImages detail images of the product
type ProductImages struct {
	Thumbnail string   `json:"thumbnail,omitempty"`
	Displays  []string `json:"displays,omitempty"`
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("sku").
			Unique().Immutable(),
		field.String("name"),
		field.JSON("descriptions", []ProductDescription{}),
		field.JSON("images", ProductImages{}),
		field.JSON("marketplaces", ProductMarketplaces{}),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("variations", Variation.Type),
		edge.From("brand", Brand.Type).
			Ref("products").
			Unique(),
	}
}
