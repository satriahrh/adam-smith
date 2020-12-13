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

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.JSON("descriptions", []struct {
			Heading string `json:"heading"`
			Body    string `json:"body"`
		}{}),
		field.JSON("images", struct {
			Thumbnail string   `json:"thumbnail"`
			Displays  []string `json:"displays"`
			Variants  []string `json:"variants"`
		}{}),
		field.JSON("marketplaces", struct {
			Tokopedia string `json:"tokopedia"`
			Shopee    string `json:"shopee"`
			Bukalapak string `json:"bukalapak"`
		}{}),
		field.JSON("prices", []uint{}),
		field.JSON("stocks", []uint{}),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("brand", Brand.Type).
			Ref("products"),
	}
}
