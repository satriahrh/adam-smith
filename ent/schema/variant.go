package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Variant holds the schema definition for the Variant entity.
type Variant struct {
	ent.Schema
}

// Fields of the Variant.
func (Variant) Fields() []ent.Field {
	return []ent.Field{
		field.Strings("images").
			Optional(),
		field.Uint8("stock").
			Default(0),
		field.Uint("price").
			Default(0),
	}
}

// Edges of the Variant.
func (Variant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Variant.Type).
			From("parent").
			Unique(),
		edge.To("variation", Variation.Type).
			Unique(),
		edge.From("product", Product.Type).
			Ref("variants").
			Unique(),
		edge.From("outbound_deals", OutboundDeal.Type).
			Ref("variant"),
	}
}
