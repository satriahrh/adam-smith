package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// OutboundTransaction holds the schema definition for the OutboundTransaction entity.
type OutboundTransaction struct {
	ent.Schema
}

type OutboundTransactionInvoice struct {
	SubAmount int `json:"sub_amount"`
	Shipping  int `json:"shipping"`
}

type OutboundTransactionBenefit struct {
	DiscountAmount uint `json:"discount_amount"`
	GiftAmount     uint `json:"gift_amount"`
	Rounding       int  `json:"rounding"`
}

type OutboundTransactionCost struct {
	MarketplaceAdmin     uint `json:"marketplace_admin"`
	MarketplacePromotion uint `json:"marketplace_promotion"`
	DiscountedAmount     uint `json:"discounted_amount"`
	Shipping             uint `json:"shipping"`
}

// Fields of the OutboundTransaction.
func (OutboundTransaction) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("channel").
			Values("shopee", "direct"),
		field.JSON("invoice", OutboundTransactionInvoice{}),
		field.JSON("benefit", OutboundTransactionBenefit{}),
		field.JSON("cost", OutboundTransactionCost{}),
		field.Uint("amount"),
	}
}

// Edges of the OutboundTransaction.
func (OutboundTransaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("deals", OutboundDeal.Type),
		edge.To("shipping", OutboundShipping.Type).
			Unique(),
	}
}
