// Code generated by entc, DO NOT EDIT.

package outboundshipping

import (
	"fmt"
)

const (
	// Label holds the string label denoting the outboundshipping type in the database.
	Label = "outbound_shipping"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCourier holds the string denoting the courier field in the database.
	FieldCourier = "courier"
	// FieldCourierTrackingCode holds the string denoting the courier_tracking_code field in the database.
	FieldCourierTrackingCode = "courier_tracking_code"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldConsignee holds the string denoting the consignee field in the database.
	FieldConsignee = "consignee"
	// FieldConsigneePhone holds the string denoting the consignee_phone field in the database.
	FieldConsigneePhone = "consignee_phone"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldPostalCode holds the string denoting the postal_code field in the database.
	FieldPostalCode = "postal_code"
	// FieldCost holds the string denoting the cost field in the database.
	FieldCost = "cost"

	// EdgeTransaction holds the string denoting the transaction edge name in mutations.
	EdgeTransaction = "transaction"

	// Table holds the table name of the outboundshipping in the database.
	Table = "outbound_shippings"
	// TransactionTable is the table the holds the transaction relation/edge.
	TransactionTable = "outbound_shippings"
	// TransactionInverseTable is the table name for the OutboundTransaction entity.
	// It exists in this package in order to avoid circular dependency with the "outboundtransaction" package.
	TransactionInverseTable = "outbound_transactions"
	// TransactionColumn is the table column denoting the transaction relation/edge.
	TransactionColumn = "outbound_transaction_shipping"
)

// Columns holds all SQL columns for outboundshipping fields.
var Columns = []string{
	FieldID,
	FieldCourier,
	FieldCourierTrackingCode,
	FieldType,
	FieldState,
	FieldConsignee,
	FieldConsigneePhone,
	FieldAddress,
	FieldPostalCode,
	FieldCost,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the OutboundShipping type.
var ForeignKeys = []string{
	"outbound_transaction_shipping",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Courier defines the type for the courier enum field.
type Courier string

// Courier values.
const (
	CourierSelf          Courier = "self"
	CourierJnt           Courier = "jnt"
	CourierShopeeExpress Courier = "shopee_express"
	CourierJne           Courier = "jne"
	CourierSicepat       Courier = "sicepat"
	CourierNinjaExpress  Courier = "ninja_express"
	CourierAnteraja      Courier = "anteraja"
	CourierPosIndonesia  Courier = "pos_indonesia"
)

func (c Courier) String() string {
	return string(c)
}

// CourierValidator is a validator for the "courier" field enum values. It is called by the builders before save.
func CourierValidator(c Courier) error {
	switch c {
	case CourierSelf, CourierJnt, CourierShopeeExpress, CourierJne, CourierSicepat, CourierNinjaExpress, CourierAnteraja, CourierPosIndonesia:
		return nil
	default:
		return fmt.Errorf("outboundshipping: invalid enum value for courier field: %q", c)
	}
}

// Type defines the type for the type enum field.
type Type string

// Type values.
const (
	TypePick Type = "pick"
	TypeDrop Type = "drop"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypePick, TypeDrop:
		return nil
	default:
		return fmt.Errorf("outboundshipping: invalid enum value for type field: %q", _type)
	}
}

// State defines the type for the state enum field.
type State string

// State values.
const (
	StateRequested State = "requested"
	StateSent      State = "sent"
	StateDelivered State = "delivered"
	StateFinish    State = "finish"
)

func (s State) String() string {
	return string(s)
}

// StateValidator is a validator for the "state" field enum values. It is called by the builders before save.
func StateValidator(s State) error {
	switch s {
	case StateRequested, StateSent, StateDelivered, StateFinish:
		return nil
	default:
		return fmt.Errorf("outboundshipping: invalid enum value for state field: %q", s)
	}
}
