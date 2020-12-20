package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
)

// Variation holds the schema definition for the Variation entity.
type Variation struct {
	ent.Schema
}

// Fields of the Variant.
func (v Variation) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").
			NotEmpty(),
		field.String("value").
			NotEmpty(),
	}
}

// Edges of the Variation.
func (Variation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("variants", Variant.Type).
			Ref("variation"),
	}
}

// Indexes of the Variation.
func (Variation) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("value", "type").Unique(),
	}
}

// EnumType enum of the field type
func (Variation) EnumType() []string {
	return []string{"color", "size"}
}
