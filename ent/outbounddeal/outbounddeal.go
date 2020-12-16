// Code generated by entc, DO NOT EDIT.

package outbounddeal

const (
	// Label holds the string label denoting the outbounddeal type in the database.
	Label = "outbound_deal"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"

	// EdgeVariation holds the string denoting the variation edge name in mutations.
	EdgeVariation = "variation"
	// EdgeTransaction holds the string denoting the transaction edge name in mutations.
	EdgeTransaction = "transaction"

	// Table holds the table name of the outbounddeal in the database.
	Table = "outbound_deals"
	// VariationTable is the table the holds the variation relation/edge.
	VariationTable = "outbound_deals"
	// VariationInverseTable is the table name for the Variation entity.
	// It exists in this package in order to avoid circular dependency with the "variation" package.
	VariationInverseTable = "variations"
	// VariationColumn is the table column denoting the variation relation/edge.
	VariationColumn = "outbound_deal_variation"
	// TransactionTable is the table the holds the transaction relation/edge. The primary key declared below.
	TransactionTable = "outbound_transaction_deals"
	// TransactionInverseTable is the table name for the OutboundTransaction entity.
	// It exists in this package in order to avoid circular dependency with the "outboundtransaction" package.
	TransactionInverseTable = "outbound_transactions"
)

// Columns holds all SQL columns for outbounddeal fields.
var Columns = []string{
	FieldID,
	FieldQuantity,
	FieldAmount,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the OutboundDeal type.
var ForeignKeys = []string{
	"outbound_deal_variation",
}

var (
	// TransactionPrimaryKey and TransactionColumn2 are the table columns denoting the
	// primary key for the transaction relation (M2M).
	TransactionPrimaryKey = []string{"outbound_transaction_id", "outbound_deal_id"}
)

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
