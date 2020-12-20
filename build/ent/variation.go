// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/satriahrh/adam-smith/build/ent/variation"
)

// Variation is the model entity for the Variation schema.
type Variation struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type variation.Type `json:"type,omitempty"`
	// Value holds the value of the "value" field.
	Value string `json:"value,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VariationQuery when eager-loading is set.
	Edges VariationEdges `json:"edges"`
}

// VariationEdges holds the relations/edges for other nodes in the graph.
type VariationEdges struct {
	// Variants holds the value of the variants edge.
	Variants []*Variant
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// VariantsOrErr returns the Variants value or an error if the edge
// was not loaded in eager-loading.
func (e VariationEdges) VariantsOrErr() ([]*Variant, error) {
	if e.loadedTypes[0] {
		return e.Variants, nil
	}
	return nil, &NotLoadedError{edge: "variants"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Variation) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // type
		&sql.NullString{}, // value
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Variation fields.
func (v *Variation) assignValues(values ...interface{}) error {
	if m, n := len(values), len(variation.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	v.ID = uint64(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field type", values[0])
	} else if value.Valid {
		v.Type = variation.Type(value.String)
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field value", values[1])
	} else if value.Valid {
		v.Value = value.String
	}
	return nil
}

// QueryVariants queries the variants edge of the Variation.
func (v *Variation) QueryVariants() *VariantQuery {
	return (&VariationClient{config: v.config}).QueryVariants(v)
}

// Update returns a builder for updating this Variation.
// Note that, you need to call Variation.Unwrap() before calling this method, if this Variation
// was returned from a transaction, and the transaction was committed or rolled back.
func (v *Variation) Update() *VariationUpdateOne {
	return (&VariationClient{config: v.config}).UpdateOne(v)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (v *Variation) Unwrap() *Variation {
	tx, ok := v.config.driver.(*txDriver)
	if !ok {
		panic("ent: Variation is not a transactional entity")
	}
	v.config.driver = tx.drv
	return v
}

// String implements the fmt.Stringer.
func (v *Variation) String() string {
	var builder strings.Builder
	builder.WriteString("Variation(")
	builder.WriteString(fmt.Sprintf("id=%v", v.ID))
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", v.Type))
	builder.WriteString(", value=")
	builder.WriteString(v.Value)
	builder.WriteByte(')')
	return builder.String()
}

// Variations is a parsable slice of Variation.
type Variations []*Variation

func (v Variations) config(cfg config) {
	for _i := range v {
		v[_i].config = cfg
	}
}
