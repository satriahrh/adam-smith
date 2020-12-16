// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/satriahrh/adam-smith/ent/outbounddeal"
	"github.com/satriahrh/adam-smith/ent/variation"
)

// OutboundDeal is the model entity for the OutboundDeal schema.
type OutboundDeal struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Quantity holds the value of the "quantity" field.
	Quantity uint `json:"quantity,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount uint `json:"amount,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OutboundDealQuery when eager-loading is set.
	Edges                   OutboundDealEdges `json:"edges"`
	outbound_deal_variation *int
}

// OutboundDealEdges holds the relations/edges for other nodes in the graph.
type OutboundDealEdges struct {
	// Variation holds the value of the variation edge.
	Variation *Variation
	// Transaction holds the value of the transaction edge.
	Transaction []*OutboundTransaction
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// VariationOrErr returns the Variation value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OutboundDealEdges) VariationOrErr() (*Variation, error) {
	if e.loadedTypes[0] {
		if e.Variation == nil {
			// The edge variation was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: variation.Label}
		}
		return e.Variation, nil
	}
	return nil, &NotLoadedError{edge: "variation"}
}

// TransactionOrErr returns the Transaction value or an error if the edge
// was not loaded in eager-loading.
func (e OutboundDealEdges) TransactionOrErr() ([]*OutboundTransaction, error) {
	if e.loadedTypes[1] {
		return e.Transaction, nil
	}
	return nil, &NotLoadedError{edge: "transaction"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OutboundDeal) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullInt64{}, // quantity
		&sql.NullInt64{}, // amount
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*OutboundDeal) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // outbound_deal_variation
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OutboundDeal fields.
func (od *OutboundDeal) assignValues(values ...interface{}) error {
	if m, n := len(values), len(outbounddeal.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	od.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field quantity", values[0])
	} else if value.Valid {
		od.Quantity = uint(value.Int64)
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field amount", values[1])
	} else if value.Valid {
		od.Amount = uint(value.Int64)
	}
	values = values[2:]
	if len(values) == len(outbounddeal.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field outbound_deal_variation", value)
		} else if value.Valid {
			od.outbound_deal_variation = new(int)
			*od.outbound_deal_variation = int(value.Int64)
		}
	}
	return nil
}

// QueryVariation queries the variation edge of the OutboundDeal.
func (od *OutboundDeal) QueryVariation() *VariationQuery {
	return (&OutboundDealClient{config: od.config}).QueryVariation(od)
}

// QueryTransaction queries the transaction edge of the OutboundDeal.
func (od *OutboundDeal) QueryTransaction() *OutboundTransactionQuery {
	return (&OutboundDealClient{config: od.config}).QueryTransaction(od)
}

// Update returns a builder for updating this OutboundDeal.
// Note that, you need to call OutboundDeal.Unwrap() before calling this method, if this OutboundDeal
// was returned from a transaction, and the transaction was committed or rolled back.
func (od *OutboundDeal) Update() *OutboundDealUpdateOne {
	return (&OutboundDealClient{config: od.config}).UpdateOne(od)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (od *OutboundDeal) Unwrap() *OutboundDeal {
	tx, ok := od.config.driver.(*txDriver)
	if !ok {
		panic("ent: OutboundDeal is not a transactional entity")
	}
	od.config.driver = tx.drv
	return od
}

// String implements the fmt.Stringer.
func (od *OutboundDeal) String() string {
	var builder strings.Builder
	builder.WriteString("OutboundDeal(")
	builder.WriteString(fmt.Sprintf("id=%v", od.ID))
	builder.WriteString(", quantity=")
	builder.WriteString(fmt.Sprintf("%v", od.Quantity))
	builder.WriteString(", amount=")
	builder.WriteString(fmt.Sprintf("%v", od.Amount))
	builder.WriteByte(')')
	return builder.String()
}

// OutboundDeals is a parsable slice of OutboundDeal.
type OutboundDeals []*OutboundDeal

func (od OutboundDeals) config(cfg config) {
	for _i := range od {
		od[_i].config = cfg
	}
}
