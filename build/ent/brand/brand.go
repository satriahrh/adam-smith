// Code generated by entc, DO NOT EDIT.

package brand

const (
	// Label holds the string label denoting the brand type in the database.
	Label = "brand"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"

	// EdgeProducts holds the string denoting the products edge name in mutations.
	EdgeProducts = "products"

	// Table holds the table name of the brand in the database.
	Table = "brands"
	// ProductsTable is the table the holds the products relation/edge.
	ProductsTable = "products"
	// ProductsInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductsInverseTable = "products"
	// ProductsColumn is the table column denoting the products relation/edge.
	ProductsColumn = "brand_products"
)

// Columns holds all SQL columns for brand fields.
var Columns = []string{
	FieldID,
	FieldCode,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)
