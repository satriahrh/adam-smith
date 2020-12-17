// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/satriahrh/adam-smith/ent/brand"
	"github.com/satriahrh/adam-smith/ent/schema"
	"github.com/satriahrh/adam-smith/ent/variation"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	brandFields := schema.Brand{}.Fields()
	_ = brandFields
	// brandDescCode is the schema descriptor for code field.
	brandDescCode := brandFields[0].Descriptor()
	// brand.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	brand.CodeValidator = brandDescCode.Validators[0].(func(string) error)
	// brandDescName is the schema descriptor for name field.
	brandDescName := brandFields[1].Descriptor()
	// brand.NameValidator is a validator for the "name" field. It is called by the builders before save.
	brand.NameValidator = brandDescName.Validators[0].(func(string) error)
	variationFields := schema.Variation{}.Fields()
	_ = variationFields
	// variationDescStock is the schema descriptor for stock field.
	variationDescStock := variationFields[1].Descriptor()
	// variation.DefaultStock holds the default value on creation for the stock field.
	variation.DefaultStock = variationDescStock.Default.(uint8)
	// variationDescPrice is the schema descriptor for price field.
	variationDescPrice := variationFields[2].Descriptor()
	// variation.DefaultPrice holds the default value on creation for the price field.
	variation.DefaultPrice = variationDescPrice.Default.(uint)
}
