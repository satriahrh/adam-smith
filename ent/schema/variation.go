package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Variation holds the schema definition for the Variation entity.
type Variation struct {
	ent.Schema
}

// Fields of the Variation.
func (Variation) Fields() []ent.Field {
	return []ent.Field{
		field.Strings("images").
			Optional(),
		field.Uint8("stock").
			Default(0),
		field.Uint("price").
			Default(0),
	}
}

// Edges of the Variation.
func (Variation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Variation.Type).
			From("parent").
			Unique(),
		edge.From("product", Product.Type).
			Ref("variations").
			Unique(),
		edge.From("variant", Variant.Type).
			Ref("variations"),
		edge.From("outbound_deals", OutboundDeal.Type).
			Ref("variation"),
	}
}
