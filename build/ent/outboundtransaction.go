// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/satriahrh/adam-smith/build/ent/outboundshipping"
	"github.com/satriahrh/adam-smith/build/ent/outboundtransaction"
	"github.com/satriahrh/adam-smith/ent/schema"
)

// OutboundTransaction is the model entity for the OutboundTransaction schema.
type OutboundTransaction struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Channel holds the value of the "channel" field.
	Channel outboundtransaction.Channel `json:"channel,omitempty"`
	// Invoice holds the value of the "invoice" field.
	Invoice schema.OutboundTransactionInvoice `json:"invoice,omitempty"`
	// Benefit holds the value of the "benefit" field.
	Benefit schema.OutboundTransactionBenefit `json:"benefit,omitempty"`
	// Cost holds the value of the "cost" field.
	Cost schema.OutboundTransactionCost `json:"cost,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount uint `json:"amount,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OutboundTransactionQuery when eager-loading is set.
	Edges OutboundTransactionEdges `json:"edges"`
}

// OutboundTransactionEdges holds the relations/edges for other nodes in the graph.
type OutboundTransactionEdges struct {
	// Deals holds the value of the deals edge.
	Deals []*OutboundDeal
	// Shipping holds the value of the shipping edge.
	Shipping *OutboundShipping
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// DealsOrErr returns the Deals value or an error if the edge
// was not loaded in eager-loading.
func (e OutboundTransactionEdges) DealsOrErr() ([]*OutboundDeal, error) {
	if e.loadedTypes[0] {
		return e.Deals, nil
	}
	return nil, &NotLoadedError{edge: "deals"}
}

// ShippingOrErr returns the Shipping value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OutboundTransactionEdges) ShippingOrErr() (*OutboundShipping, error) {
	if e.loadedTypes[1] {
		if e.Shipping == nil {
			// The edge shipping was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: outboundshipping.Label}
		}
		return e.Shipping, nil
	}
	return nil, &NotLoadedError{edge: "shipping"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OutboundTransaction) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // channel
		&[]byte{},         // invoice
		&[]byte{},         // benefit
		&[]byte{},         // cost
		&sql.NullInt64{},  // amount
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OutboundTransaction fields.
func (ot *OutboundTransaction) assignValues(values ...interface{}) error {
	if m, n := len(values), len(outboundtransaction.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	ot.ID = uint64(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field channel", values[0])
	} else if value.Valid {
		ot.Channel = outboundtransaction.Channel(value.String)
	}

	if value, ok := values[1].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field invoice", values[1])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &ot.Invoice); err != nil {
			return fmt.Errorf("unmarshal field invoice: %v", err)
		}
	}

	if value, ok := values[2].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field benefit", values[2])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &ot.Benefit); err != nil {
			return fmt.Errorf("unmarshal field benefit: %v", err)
		}
	}

	if value, ok := values[3].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field cost", values[3])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &ot.Cost); err != nil {
			return fmt.Errorf("unmarshal field cost: %v", err)
		}
	}
	if value, ok := values[4].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field amount", values[4])
	} else if value.Valid {
		ot.Amount = uint(value.Int64)
	}
	return nil
}

// QueryDeals queries the deals edge of the OutboundTransaction.
func (ot *OutboundTransaction) QueryDeals() *OutboundDealQuery {
	return (&OutboundTransactionClient{config: ot.config}).QueryDeals(ot)
}

// QueryShipping queries the shipping edge of the OutboundTransaction.
func (ot *OutboundTransaction) QueryShipping() *OutboundShippingQuery {
	return (&OutboundTransactionClient{config: ot.config}).QueryShipping(ot)
}

// Update returns a builder for updating this OutboundTransaction.
// Note that, you need to call OutboundTransaction.Unwrap() before calling this method, if this OutboundTransaction
// was returned from a transaction, and the transaction was committed or rolled back.
func (ot *OutboundTransaction) Update() *OutboundTransactionUpdateOne {
	return (&OutboundTransactionClient{config: ot.config}).UpdateOne(ot)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (ot *OutboundTransaction) Unwrap() *OutboundTransaction {
	tx, ok := ot.config.driver.(*txDriver)
	if !ok {
		panic("ent: OutboundTransaction is not a transactional entity")
	}
	ot.config.driver = tx.drv
	return ot
}

// String implements the fmt.Stringer.
func (ot *OutboundTransaction) String() string {
	var builder strings.Builder
	builder.WriteString("OutboundTransaction(")
	builder.WriteString(fmt.Sprintf("id=%v", ot.ID))
	builder.WriteString(", channel=")
	builder.WriteString(fmt.Sprintf("%v", ot.Channel))
	builder.WriteString(", invoice=")
	builder.WriteString(fmt.Sprintf("%v", ot.Invoice))
	builder.WriteString(", benefit=")
	builder.WriteString(fmt.Sprintf("%v", ot.Benefit))
	builder.WriteString(", cost=")
	builder.WriteString(fmt.Sprintf("%v", ot.Cost))
	builder.WriteString(", amount=")
	builder.WriteString(fmt.Sprintf("%v", ot.Amount))
	builder.WriteByte(')')
	return builder.String()
}

// OutboundTransactions is a parsable slice of OutboundTransaction.
type OutboundTransactions []*OutboundTransaction

func (ot OutboundTransactions) config(cfg config) {
	for _i := range ot {
		ot[_i].config = cfg
	}
}
