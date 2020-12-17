// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/satriahrh/adam-smith/build/ent/outboundshipping"
	"github.com/satriahrh/adam-smith/build/ent/outboundtransaction"
)

// OutboundShipping is the model entity for the OutboundShipping schema.
type OutboundShipping struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Courier holds the value of the "courier" field.
	Courier outboundshipping.Courier `json:"courier,omitempty"`
	// CourierTrackingCode holds the value of the "courier_tracking_code" field.
	CourierTrackingCode string `json:"courier_tracking_code,omitempty"`
	// Type holds the value of the "type" field.
	Type outboundshipping.Type `json:"type,omitempty"`
	// State holds the value of the "state" field.
	State outboundshipping.State `json:"state,omitempty"`
	// Consignee holds the value of the "consignee" field.
	Consignee string `json:"consignee,omitempty"`
	// ConsigneePhone holds the value of the "consignee_phone" field.
	ConsigneePhone string `json:"consignee_phone,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// PostalCode holds the value of the "postal_code" field.
	PostalCode uint `json:"postal_code,omitempty"`
	// Cost holds the value of the "cost" field.
	Cost uint `json:"cost,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OutboundShippingQuery when eager-loading is set.
	Edges                         OutboundShippingEdges `json:"edges"`
	outbound_transaction_shipping *int
}

// OutboundShippingEdges holds the relations/edges for other nodes in the graph.
type OutboundShippingEdges struct {
	// Transaction holds the value of the transaction edge.
	Transaction *OutboundTransaction
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TransactionOrErr returns the Transaction value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OutboundShippingEdges) TransactionOrErr() (*OutboundTransaction, error) {
	if e.loadedTypes[0] {
		if e.Transaction == nil {
			// The edge transaction was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: outboundtransaction.Label}
		}
		return e.Transaction, nil
	}
	return nil, &NotLoadedError{edge: "transaction"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OutboundShipping) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // courier
		&sql.NullString{}, // courier_tracking_code
		&sql.NullString{}, // type
		&sql.NullString{}, // state
		&sql.NullString{}, // consignee
		&sql.NullString{}, // consignee_phone
		&sql.NullString{}, // address
		&sql.NullInt64{},  // postal_code
		&sql.NullInt64{},  // cost
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*OutboundShipping) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // outbound_transaction_shipping
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OutboundShipping fields.
func (os *OutboundShipping) assignValues(values ...interface{}) error {
	if m, n := len(values), len(outboundshipping.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	os.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field courier", values[0])
	} else if value.Valid {
		os.Courier = outboundshipping.Courier(value.String)
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field courier_tracking_code", values[1])
	} else if value.Valid {
		os.CourierTrackingCode = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field type", values[2])
	} else if value.Valid {
		os.Type = outboundshipping.Type(value.String)
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field state", values[3])
	} else if value.Valid {
		os.State = outboundshipping.State(value.String)
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field consignee", values[4])
	} else if value.Valid {
		os.Consignee = value.String
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field consignee_phone", values[5])
	} else if value.Valid {
		os.ConsigneePhone = value.String
	}
	if value, ok := values[6].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field address", values[6])
	} else if value.Valid {
		os.Address = value.String
	}
	if value, ok := values[7].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field postal_code", values[7])
	} else if value.Valid {
		os.PostalCode = uint(value.Int64)
	}
	if value, ok := values[8].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field cost", values[8])
	} else if value.Valid {
		os.Cost = uint(value.Int64)
	}
	values = values[9:]
	if len(values) == len(outboundshipping.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field outbound_transaction_shipping", value)
		} else if value.Valid {
			os.outbound_transaction_shipping = new(int)
			*os.outbound_transaction_shipping = int(value.Int64)
		}
	}
	return nil
}

// QueryTransaction queries the transaction edge of the OutboundShipping.
func (os *OutboundShipping) QueryTransaction() *OutboundTransactionQuery {
	return (&OutboundShippingClient{config: os.config}).QueryTransaction(os)
}

// Update returns a builder for updating this OutboundShipping.
// Note that, you need to call OutboundShipping.Unwrap() before calling this method, if this OutboundShipping
// was returned from a transaction, and the transaction was committed or rolled back.
func (os *OutboundShipping) Update() *OutboundShippingUpdateOne {
	return (&OutboundShippingClient{config: os.config}).UpdateOne(os)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (os *OutboundShipping) Unwrap() *OutboundShipping {
	tx, ok := os.config.driver.(*txDriver)
	if !ok {
		panic("ent: OutboundShipping is not a transactional entity")
	}
	os.config.driver = tx.drv
	return os
}

// String implements the fmt.Stringer.
func (os *OutboundShipping) String() string {
	var builder strings.Builder
	builder.WriteString("OutboundShipping(")
	builder.WriteString(fmt.Sprintf("id=%v", os.ID))
	builder.WriteString(", courier=")
	builder.WriteString(fmt.Sprintf("%v", os.Courier))
	builder.WriteString(", courier_tracking_code=")
	builder.WriteString(os.CourierTrackingCode)
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", os.Type))
	builder.WriteString(", state=")
	builder.WriteString(fmt.Sprintf("%v", os.State))
	builder.WriteString(", consignee=")
	builder.WriteString(os.Consignee)
	builder.WriteString(", consignee_phone=")
	builder.WriteString(os.ConsigneePhone)
	builder.WriteString(", address=")
	builder.WriteString(os.Address)
	builder.WriteString(", postal_code=")
	builder.WriteString(fmt.Sprintf("%v", os.PostalCode))
	builder.WriteString(", cost=")
	builder.WriteString(fmt.Sprintf("%v", os.Cost))
	builder.WriteByte(')')
	return builder.String()
}

// OutboundShippings is a parsable slice of OutboundShipping.
type OutboundShippings []*OutboundShipping

func (os OutboundShippings) config(cfg config) {
	for _i := range os {
		os[_i].config = cfg
	}
}