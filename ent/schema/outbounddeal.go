package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// OutboundDeal holds the schema definition for the OutboundDeal entity.
type OutboundDeal struct {
	ent.Schema
}

// Fields of the OutboundDeal.
func (OutboundDeal) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("quantity").Comment("Number of purchased product."),
		field.Uint("amount").Comment("Total amount of the deals."),
	}
}

// Edges of the OutboundDeal.
func (OutboundDeal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("variant", Variant.Type).
			Unique(),
		edge.To("children", OutboundDeal.Type).
			From("parent").
			Unique(),
		edge.From("transaction", OutboundTransaction.Type).
			Ref("deals").
			Unique(),
	}
}
