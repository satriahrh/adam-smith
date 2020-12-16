// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/satriahrh/adam-smith/ent/outbounddeal"
	"github.com/satriahrh/adam-smith/ent/predicate"
	"github.com/satriahrh/adam-smith/ent/product"
	"github.com/satriahrh/adam-smith/ent/variant"
	"github.com/satriahrh/adam-smith/ent/variation"
)

// VariationUpdate is the builder for updating Variation entities.
type VariationUpdate struct {
	config
	hooks    []Hook
	mutation *VariationMutation
}

// Where adds a new predicate for the builder.
func (vu *VariationUpdate) Where(ps ...predicate.Variation) *VariationUpdate {
	vu.mutation.predicates = append(vu.mutation.predicates, ps...)
	return vu
}

// SetImages sets the images field.
func (vu *VariationUpdate) SetImages(s []string) *VariationUpdate {
	vu.mutation.SetImages(s)
	return vu
}

// ClearImages clears the value of images.
func (vu *VariationUpdate) ClearImages() *VariationUpdate {
	vu.mutation.ClearImages()
	return vu
}

// SetStock sets the stock field.
func (vu *VariationUpdate) SetStock(u uint8) *VariationUpdate {
	vu.mutation.ResetStock()
	vu.mutation.SetStock(u)
	return vu
}

// SetNillableStock sets the stock field if the given value is not nil.
func (vu *VariationUpdate) SetNillableStock(u *uint8) *VariationUpdate {
	if u != nil {
		vu.SetStock(*u)
	}
	return vu
}

// AddStock adds u to stock.
func (vu *VariationUpdate) AddStock(u uint8) *VariationUpdate {
	vu.mutation.AddStock(u)
	return vu
}

// SetPrice sets the price field.
func (vu *VariationUpdate) SetPrice(u uint) *VariationUpdate {
	vu.mutation.ResetPrice()
	vu.mutation.SetPrice(u)
	return vu
}

// SetNillablePrice sets the price field if the given value is not nil.
func (vu *VariationUpdate) SetNillablePrice(u *uint) *VariationUpdate {
	if u != nil {
		vu.SetPrice(*u)
	}
	return vu
}

// AddPrice adds u to price.
func (vu *VariationUpdate) AddPrice(u uint) *VariationUpdate {
	vu.mutation.AddPrice(u)
	return vu
}

// SetParentID sets the parent edge to Variation by id.
func (vu *VariationUpdate) SetParentID(id int) *VariationUpdate {
	vu.mutation.SetParentID(id)
	return vu
}

// SetNillableParentID sets the parent edge to Variation by id if the given value is not nil.
func (vu *VariationUpdate) SetNillableParentID(id *int) *VariationUpdate {
	if id != nil {
		vu = vu.SetParentID(*id)
	}
	return vu
}

// SetParent sets the parent edge to Variation.
func (vu *VariationUpdate) SetParent(v *Variation) *VariationUpdate {
	return vu.SetParentID(v.ID)
}

// AddChildIDs adds the children edge to Variation by ids.
func (vu *VariationUpdate) AddChildIDs(ids ...int) *VariationUpdate {
	vu.mutation.AddChildIDs(ids...)
	return vu
}

// AddChildren adds the children edges to Variation.
func (vu *VariationUpdate) AddChildren(v ...*Variation) *VariationUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vu.AddChildIDs(ids...)
}

// SetProductID sets the product edge to Product by id.
func (vu *VariationUpdate) SetProductID(id int) *VariationUpdate {
	vu.mutation.SetProductID(id)
	return vu
}

// SetNillableProductID sets the product edge to Product by id if the given value is not nil.
func (vu *VariationUpdate) SetNillableProductID(id *int) *VariationUpdate {
	if id != nil {
		vu = vu.SetProductID(*id)
	}
	return vu
}

// SetProduct sets the product edge to Product.
func (vu *VariationUpdate) SetProduct(p *Product) *VariationUpdate {
	return vu.SetProductID(p.ID)
}

// AddVariantIDs adds the variant edge to Variant by ids.
func (vu *VariationUpdate) AddVariantIDs(ids ...int) *VariationUpdate {
	vu.mutation.AddVariantIDs(ids...)
	return vu
}

// AddVariant adds the variant edges to Variant.
func (vu *VariationUpdate) AddVariant(v ...*Variant) *VariationUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vu.AddVariantIDs(ids...)
}

// AddOutboundDealIDs adds the outbound_deals edge to OutboundDeal by ids.
func (vu *VariationUpdate) AddOutboundDealIDs(ids ...int) *VariationUpdate {
	vu.mutation.AddOutboundDealIDs(ids...)
	return vu
}

// AddOutboundDeals adds the outbound_deals edges to OutboundDeal.
func (vu *VariationUpdate) AddOutboundDeals(o ...*OutboundDeal) *VariationUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return vu.AddOutboundDealIDs(ids...)
}

// Mutation returns the VariationMutation object of the builder.
func (vu *VariationUpdate) Mutation() *VariationMutation {
	return vu.mutation
}

// ClearParent clears the "parent" edge to type Variation.
func (vu *VariationUpdate) ClearParent() *VariationUpdate {
	vu.mutation.ClearParent()
	return vu
}

// ClearChildren clears all "children" edges to type Variation.
func (vu *VariationUpdate) ClearChildren() *VariationUpdate {
	vu.mutation.ClearChildren()
	return vu
}

// RemoveChildIDs removes the children edge to Variation by ids.
func (vu *VariationUpdate) RemoveChildIDs(ids ...int) *VariationUpdate {
	vu.mutation.RemoveChildIDs(ids...)
	return vu
}

// RemoveChildren removes children edges to Variation.
func (vu *VariationUpdate) RemoveChildren(v ...*Variation) *VariationUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vu.RemoveChildIDs(ids...)
}

// ClearProduct clears the "product" edge to type Product.
func (vu *VariationUpdate) ClearProduct() *VariationUpdate {
	vu.mutation.ClearProduct()
	return vu
}

// ClearVariant clears all "variant" edges to type Variant.
func (vu *VariationUpdate) ClearVariant() *VariationUpdate {
	vu.mutation.ClearVariant()
	return vu
}

// RemoveVariantIDs removes the variant edge to Variant by ids.
func (vu *VariationUpdate) RemoveVariantIDs(ids ...int) *VariationUpdate {
	vu.mutation.RemoveVariantIDs(ids...)
	return vu
}

// RemoveVariant removes variant edges to Variant.
func (vu *VariationUpdate) RemoveVariant(v ...*Variant) *VariationUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vu.RemoveVariantIDs(ids...)
}

// ClearOutboundDeals clears all "outbound_deals" edges to type OutboundDeal.
func (vu *VariationUpdate) ClearOutboundDeals() *VariationUpdate {
	vu.mutation.ClearOutboundDeals()
	return vu
}

// RemoveOutboundDealIDs removes the outbound_deals edge to OutboundDeal by ids.
func (vu *VariationUpdate) RemoveOutboundDealIDs(ids ...int) *VariationUpdate {
	vu.mutation.RemoveOutboundDealIDs(ids...)
	return vu
}

// RemoveOutboundDeals removes outbound_deals edges to OutboundDeal.
func (vu *VariationUpdate) RemoveOutboundDeals(o ...*OutboundDeal) *VariationUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return vu.RemoveOutboundDealIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vu *VariationUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vu.hooks) == 0 {
		affected, err = vu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VariationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vu.mutation = mutation
			affected, err = vu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vu.hooks) - 1; i >= 0; i-- {
			mut = vu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (vu *VariationUpdate) SaveX(ctx context.Context) int {
	affected, err := vu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vu *VariationUpdate) Exec(ctx context.Context) error {
	_, err := vu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vu *VariationUpdate) ExecX(ctx context.Context) {
	if err := vu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (vu *VariationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   variation.Table,
			Columns: variation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: variation.FieldID,
			},
		},
	}
	if ps := vu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vu.mutation.Images(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: variation.FieldImages,
		})
	}
	if vu.mutation.ImagesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: variation.FieldImages,
		})
	}
	if value, ok := vu.mutation.Stock(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: variation.FieldStock,
		})
	}
	if value, ok := vu.mutation.AddedStock(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: variation.FieldStock,
		})
	}
	if value, ok := vu.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: variation.FieldPrice,
		})
	}
	if value, ok := vu.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: variation.FieldPrice,
		})
	}
	if vu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   variation.ParentTable,
			Columns: []string{variation.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: variation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   variation.ParentTable,
			Columns: []string{variation.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: variation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   variation.ChildrenTable,
			Columns: []string{variation.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: variation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !vu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   variation.ChildrenTable,
			Columns: []string{variation.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: variation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   variation.ChildrenTable,
			Columns: []string{variation.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: variation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vu.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   variation.ProductTable,
			Columns: []string{variation.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   variation.ProductTable,
			Columns: []string{variation.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vu.mutation.VariantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   variation.VariantTable,
			Columns: variation.VariantPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: variant.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.RemovedVariantIDs(); len(nodes) > 0 && !vu.mutation.VariantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   variation.VariantTable,
			Columns: variation.VariantPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: variant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.VariantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   variation.VariantTable,
			Columns: variation.VariantPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: variant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vu.mutation.OutboundDealsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   variation.OutboundDealsTable,
			Columns: variation.OutboundDealsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outbounddeal.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.RemovedOutboundDealsIDs(); len(nodes) > 0 && !vu.mutation.OutboundDealsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   variation.OutboundDealsTable,
			Columns: variation.OutboundDealsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outbounddeal.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.OutboundDealsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   variation.OutboundDealsTable,
			Columns: variation.OutboundDealsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outbounddeal.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{variation.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// VariationUpdateOne is the builder for updating a single Variation entity.
type VariationUpdateOne struct {
	config
	hooks    []Hook
	mutation *VariationMutation
}

// SetImages sets the images field.
func (vuo *VariationUpdateOne) SetImages(s []string) *VariationUpdateOne {
	vuo.mutation.SetImages(s)
	return vuo
}

// ClearImages clears the value of images.
func (vuo *VariationUpdateOne) ClearImages() *VariationUpdateOne {
	vuo.mutation.ClearImages()
	return vuo
}

// SetStock sets the stock field.
func (vuo *VariationUpdateOne) SetStock(u uint8) *VariationUpdateOne {
	vuo.mutation.ResetStock()
	vuo.mutation.SetStock(u)
	return vuo
}

// SetNillableStock sets the stock field if the given value is not nil.
func (vuo *VariationUpdateOne) SetNillableStock(u *uint8) *VariationUpdateOne {
	if u != nil {
		vuo.SetStock(*u)
	}
	return vuo
}

// AddStock adds u to stock.
func (vuo *VariationUpdateOne) AddStock(u uint8) *VariationUpdateOne {
	vuo.mutation.AddStock(u)
	return vuo
}

// SetPrice sets the price field.
func (vuo *VariationUpdateOne) SetPrice(u uint) *VariationUpdateOne {
	vuo.mutation.ResetPrice()
	vuo.mutation.SetPrice(u)
	return vuo
}

// SetNillablePrice sets the price field if the given value is not nil.
func (vuo *VariationUpdateOne) SetNillablePrice(u *uint) *VariationUpdateOne {
	if u != nil {
		vuo.SetPrice(*u)
	}
	return vuo
}

// AddPrice adds u to price.
func (vuo *VariationUpdateOne) AddPrice(u uint) *VariationUpdateOne {
	vuo.mutation.AddPrice(u)
	return vuo
}

// SetParentID sets the parent edge to Variation by id.
func (vuo *VariationUpdateOne) SetParentID(id int) *VariationUpdateOne {
	vuo.mutation.SetParentID(id)
	return vuo
}

// SetNillableParentID sets the parent edge to Variation by id if the given value is not nil.
func (vuo *VariationUpdateOne) SetNillableParentID(id *int) *VariationUpdateOne {
	if id != nil {
		vuo = vuo.SetParentID(*id)
	}
	return vuo
}

// SetParent sets the parent edge to Variation.
func (vuo *VariationUpdateOne) SetParent(v *Variation) *VariationUpdateOne {
	return vuo.SetParentID(v.ID)
}

// AddChildIDs adds the children edge to Variation by ids.
func (vuo *VariationUpdateOne) AddChildIDs(ids ...int) *VariationUpdateOne {
	vuo.mutation.AddChildIDs(ids...)
	return vuo
}

// AddChildren adds the children edges to Variation.
func (vuo *VariationUpdateOne) AddChildren(v ...*Variation) *VariationUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vuo.AddChildIDs(ids...)
}

// SetProductID sets the product edge to Product by id.
func (vuo *VariationUpdateOne) SetProductID(id int) *VariationUpdateOne {
	vuo.mutation.SetProductID(id)
	return vuo
}

// SetNillableProductID sets the product edge to Product by id if the given value is not nil.
func (vuo *VariationUpdateOne) SetNillableProductID(id *int) *VariationUpdateOne {
	if id != nil {
		vuo = vuo.SetProductID(*id)
	}
	return vuo
}

// SetProduct sets the product edge to Product.
func (vuo *VariationUpdateOne) SetProduct(p *Product) *VariationUpdateOne {
	return vuo.SetProductID(p.ID)
}

// AddVariantIDs adds the variant edge to Variant by ids.
func (vuo *VariationUpdateOne) AddVariantIDs(ids ...int) *VariationUpdateOne {
	vuo.mutation.AddVariantIDs(ids...)
	return vuo
}

// AddVariant adds the variant edges to Variant.
func (vuo *VariationUpdateOne) AddVariant(v ...*Variant) *VariationUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vuo.AddVariantIDs(ids...)
}

// AddOutboundDealIDs adds the outbound_deals edge to OutboundDeal by ids.
func (vuo *VariationUpdateOne) AddOutboundDealIDs(ids ...int) *VariationUpdateOne {
	vuo.mutation.AddOutboundDealIDs(ids...)
	return vuo
}

// AddOutboundDeals adds the outbound_deals edges to OutboundDeal.
func (vuo *VariationUpdateOne) AddOutboundDeals(o ...*OutboundDeal) *VariationUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return vuo.AddOutboundDealIDs(ids...)
}

// Mutation returns the VariationMutation object of the builder.
func (vuo *VariationUpdateOne) Mutation() *VariationMutation {
	return vuo.mutation
}

// ClearParent clears the "parent" edge to type Variation.
func (vuo *VariationUpdateOne) ClearParent() *VariationUpdateOne {
	vuo.mutation.ClearParent()
	return vuo
}

// ClearChildren clears all "children" edges to type Variation.
func (vuo *VariationUpdateOne) ClearChildren() *VariationUpdateOne {
	vuo.mutation.ClearChildren()
	return vuo
}

// RemoveChildIDs removes the children edge to Variation by ids.
func (vuo *VariationUpdateOne) RemoveChildIDs(ids ...int) *VariationUpdateOne {
	vuo.mutation.RemoveChildIDs(ids...)
	return vuo
}

// RemoveChildren removes children edges to Variation.
func (vuo *VariationUpdateOne) RemoveChildren(v ...*Variation) *VariationUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vuo.RemoveChildIDs(ids...)
}

// ClearProduct clears the "product" edge to type Product.
func (vuo *VariationUpdateOne) ClearProduct() *VariationUpdateOne {
	vuo.mutation.ClearProduct()
	return vuo
}

// ClearVariant clears all "variant" edges to type Variant.
func (vuo *VariationUpdateOne) ClearVariant() *VariationUpdateOne {
	vuo.mutation.ClearVariant()
	return vuo
}

// RemoveVariantIDs removes the variant edge to Variant by ids.
func (vuo *VariationUpdateOne) RemoveVariantIDs(ids ...int) *VariationUpdateOne {
	vuo.mutation.RemoveVariantIDs(ids...)
	return vuo
}

// RemoveVariant removes variant edges to Variant.
func (vuo *VariationUpdateOne) RemoveVariant(v ...*Variant) *VariationUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vuo.RemoveVariantIDs(ids...)
}

// ClearOutboundDeals clears all "outbound_deals" edges to type OutboundDeal.
func (vuo *VariationUpdateOne) ClearOutboundDeals() *VariationUpdateOne {
	vuo.mutation.ClearOutboundDeals()
	return vuo
}

// RemoveOutboundDealIDs removes the outbound_deals edge to OutboundDeal by ids.
func (vuo *VariationUpdateOne) RemoveOutboundDealIDs(ids ...int) *VariationUpdateOne {
	vuo.mutation.RemoveOutboundDealIDs(ids...)
	return vuo
}

// RemoveOutboundDeals removes outbound_deals edges to OutboundDeal.
func (vuo *VariationUpdateOne) RemoveOutboundDeals(o ...*OutboundDeal) *VariationUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return vuo.RemoveOutboundDealIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (vuo *VariationUpdateOne) Save(ctx context.Context) (*Variation, error) {
	var (
		err  error
		node *Variation
	)
	if len(vuo.hooks) == 0 {
		node, err = vuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VariationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vuo.mutation = mutation
			node, err = vuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(vuo.hooks) - 1; i >= 0; i-- {
			mut = vuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (vuo *VariationUpdateOne) SaveX(ctx context.Context) *Variation {
	node, err := vuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vuo *VariationUpdateOne) Exec(ctx context.Context) error {
	_, err := vuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vuo *VariationUpdateOne) ExecX(ctx context.Context) {
	if err := vuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (vuo *VariationUpdateOne) sqlSave(ctx context.Context) (_node *Variation, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   variation.Table,
			Columns: variation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: variation.FieldID,
			},
		},
	}
	id, ok := vuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Variation.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := vuo.mutation.Images(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: variation.FieldImages,
		})
	}
	if vuo.mutation.ImagesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: variation.FieldImages,
		})
	}
	if value, ok := vuo.mutation.Stock(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: variation.FieldStock,
		})
	}
	if value, ok := vuo.mutation.AddedStock(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: variation.FieldStock,
		})
	}
	if value, ok := vuo.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: variation.FieldPrice,
		})
	}
	if value, ok := vuo.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: variation.FieldPrice,
		})
	}
	if vuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   variation.ParentTable,
			Columns: []string{variation.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: variation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   variation.ParentTable,
			Columns: []string{variation.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: variation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   variation.ChildrenTable,
			Columns: []string{variation.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: variation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !vuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   variation.ChildrenTable,
			Columns: []string{variation.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: variation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   variation.ChildrenTable,
			Columns: []string{variation.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: variation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vuo.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   variation.ProductTable,
			Columns: []string{variation.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   variation.ProductTable,
			Columns: []string{variation.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vuo.mutation.VariantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   variation.VariantTable,
			Columns: variation.VariantPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: variant.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.RemovedVariantIDs(); len(nodes) > 0 && !vuo.mutation.VariantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   variation.VariantTable,
			Columns: variation.VariantPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: variant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.VariantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   variation.VariantTable,
			Columns: variation.VariantPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: variant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vuo.mutation.OutboundDealsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   variation.OutboundDealsTable,
			Columns: variation.OutboundDealsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outbounddeal.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.RemovedOutboundDealsIDs(); len(nodes) > 0 && !vuo.mutation.OutboundDealsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   variation.OutboundDealsTable,
			Columns: variation.OutboundDealsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outbounddeal.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.OutboundDealsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   variation.OutboundDealsTable,
			Columns: variation.OutboundDealsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outbounddeal.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Variation{config: vuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, vuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{variation.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
