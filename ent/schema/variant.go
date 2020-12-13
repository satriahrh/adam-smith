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
func (v Variant) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").
			Values(v.EnumType()...),
		field.String("value"),
	}
}

// Edges of the Variant.
func (Variant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("variant_uses", ProductVariant.Type),
	}
}

// EnumType enum of the field type
func (Variant) EnumType() []string {
	return []string{"color", "size"}
}
