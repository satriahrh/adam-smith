// Code generated by entc, DO NOT EDIT.

package variant

import (
	"fmt"
)

const (
	// Label holds the string label denoting the variant type in the database.
	Label = "variant"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"

	// EdgeVariantUses holds the string denoting the variant_uses edge name in mutations.
	EdgeVariantUses = "variant_uses"

	// Table holds the table name of the variant in the database.
	Table = "variants"
	// VariantUsesTable is the table the holds the variant_uses relation/edge. The primary key declared below.
	VariantUsesTable = "variant_variant_uses"
	// VariantUsesInverseTable is the table name for the Variation entity.
	// It exists in this package in order to avoid circular dependency with the "variation" package.
	VariantUsesInverseTable = "variations"
)

// Columns holds all SQL columns for variant fields.
var Columns = []string{
	FieldID,
	FieldType,
	FieldValue,
}

var (
	// VariantUsesPrimaryKey and VariantUsesColumn2 are the table columns denoting the
	// primary key for the variant_uses relation (M2M).
	VariantUsesPrimaryKey = []string{"variant_id", "variation_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Type defines the type for the type enum field.
type Type string

// Type values.
const (
	TypeColor Type = "color"
	TypeSize  Type = "size"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeColor, TypeSize:
		return nil
	default:
		return fmt.Errorf("variant: invalid enum value for type field: %q", _type)
	}
}