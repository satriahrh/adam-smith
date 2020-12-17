// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/satriahrh/adam-smith/build/ent/variant"
)

// Variant is the model entity for the Variant schema.
type Variant struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type variant.Type `json:"type,omitempty"`
	// Value holds the value of the "value" field.
	Value string `json:"value,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VariantQuery when eager-loading is set.
	Edges VariantEdges `json:"edges"`
}

// VariantEdges holds the relations/edges for other nodes in the graph.
type VariantEdges struct {
	// Variations holds the value of the variations edge.
	Variations []*Variation
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// VariationsOrErr returns the Variations value or an error if the edge
// was not loaded in eager-loading.
func (e VariantEdges) VariationsOrErr() ([]*Variation, error) {
	if e.loadedTypes[0] {
		return e.Variations, nil
	}
	return nil, &NotLoadedError{edge: "variations"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Variant) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // type
		&sql.NullString{}, // value
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Variant fields.
func (v *Variant) assignValues(values ...interface{}) error {
	if m, n := len(values), len(variant.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	v.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field type", values[0])
	} else if value.Valid {
		v.Type = variant.Type(value.String)
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field value", values[1])
	} else if value.Valid {
		v.Value = value.String
	}
	return nil
}

// QueryVariations queries the variations edge of the Variant.
func (v *Variant) QueryVariations() *VariationQuery {
	return (&VariantClient{config: v.config}).QueryVariations(v)
}

// Update returns a builder for updating this Variant.
// Note that, you need to call Variant.Unwrap() before calling this method, if this Variant
// was returned from a transaction, and the transaction was committed or rolled back.
func (v *Variant) Update() *VariantUpdateOne {
	return (&VariantClient{config: v.config}).UpdateOne(v)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (v *Variant) Unwrap() *Variant {
	tx, ok := v.config.driver.(*txDriver)
	if !ok {
		panic("ent: Variant is not a transactional entity")
	}
	v.config.driver = tx.drv
	return v
}

// String implements the fmt.Stringer.
func (v *Variant) String() string {
	var builder strings.Builder
	builder.WriteString("Variant(")
	builder.WriteString(fmt.Sprintf("id=%v", v.ID))
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", v.Type))
	builder.WriteString(", value=")
	builder.WriteString(v.Value)
	builder.WriteByte(')')
	return builder.String()
}

// Variants is a parsable slice of Variant.
type Variants []*Variant

func (v Variants) config(cfg config) {
	for _i := range v {
		v[_i].config = cfg
	}
}