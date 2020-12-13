package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// ProductVariant holds the schema definition for the ProductVariant entity.
type ProductVariant struct {
	ent.Schema
}

// Fields of the ProductVariant.
func (ProductVariant) Fields() []ent.Field {
	return []ent.Field{
		field.Uint8("level"),
		field.Strings("images"),
		field.Uint8("stock"),
		field.Uint("price"),
	}
}

// Edges of the ProductVariant.
func (ProductVariant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", ProductVariant.Type).
			From("parent").
			Unique(),
		edge.From("product", Product.Type).
			Ref("variants"),
		edge.From("variant", Variant.Type).
			Ref("variant_uses"),
	}
}
