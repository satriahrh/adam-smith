package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// OutboundShipping holds the schema definition for the Delivery entity.
type OutboundShipping struct {
	ent.Schema
}

// Fields of the OutboundShipping.
func (OutboundShipping) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("courier").
			Values("self", "jnt", "shopee_express", "jne", "sicepat", "ninja_express", "anteraja", "pos_indonesia"),
		field.String("courier_tracking_code"),
		field.Enum("type").
			Values("pick", "drop"),
		field.Enum("state").
			Values("requested", "sent", "delivered", "finish"),
		field.String("consignee"),
		field.String("consignee_phone"),
		field.String("address"),
		field.Uint("postal_code"),
		field.Uint("cost"),
	}
}

// Edges of the OutboundShipping.
func (OutboundShipping) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("transaction", OutboundTransaction.Type).
			Ref("shipping").
			Unique(),
	}
}
