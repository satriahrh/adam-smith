// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/satriahrh/adam-smith/build/ent/outbounddeal"
	"github.com/satriahrh/adam-smith/build/ent/product"
	"github.com/satriahrh/adam-smith/build/ent/variant"
	"github.com/satriahrh/adam-smith/build/ent/variation"
)

// VariationCreate is the builder for creating a Variation entity.
type VariationCreate struct {
	config
	mutation *VariationMutation
	hooks    []Hook
}

// SetImages sets the images field.
func (vc *VariationCreate) SetImages(s []string) *VariationCreate {
	vc.mutation.SetImages(s)
	return vc
}

// SetStock sets the stock field.
func (vc *VariationCreate) SetStock(u uint8) *VariationCreate {
	vc.mutation.SetStock(u)
	return vc
}

// SetNillableStock sets the stock field if the given value is not nil.
func (vc *VariationCreate) SetNillableStock(u *uint8) *VariationCreate {
	if u != nil {
		vc.SetStock(*u)
	}
	return vc
}

// SetPrice sets the price field.
func (vc *VariationCreate) SetPrice(u uint) *VariationCreate {
	vc.mutation.SetPrice(u)
	return vc
}

// SetNillablePrice sets the price field if the given value is not nil.
func (vc *VariationCreate) SetNillablePrice(u *uint) *VariationCreate {
	if u != nil {
		vc.SetPrice(*u)
	}
	return vc
}

// SetParentID sets the parent edge to Variation by id.
func (vc *VariationCreate) SetParentID(id int) *VariationCreate {
	vc.mutation.SetParentID(id)
	return vc
}

// SetNillableParentID sets the parent edge to Variation by id if the given value is not nil.
func (vc *VariationCreate) SetNillableParentID(id *int) *VariationCreate {
	if id != nil {
		vc = vc.SetParentID(*id)
	}
	return vc
}

// SetParent sets the parent edge to Variation.
func (vc *VariationCreate) SetParent(v *Variation) *VariationCreate {
	return vc.SetParentID(v.ID)
}

// AddChildIDs adds the children edge to Variation by ids.
func (vc *VariationCreate) AddChildIDs(ids ...int) *VariationCreate {
	vc.mutation.AddChildIDs(ids...)
	return vc
}

// AddChildren adds the children edges to Variation.
func (vc *VariationCreate) AddChildren(v ...*Variation) *VariationCreate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vc.AddChildIDs(ids...)
}

// SetProductID sets the product edge to Product by id.
func (vc *VariationCreate) SetProductID(id int) *VariationCreate {
	vc.mutation.SetProductID(id)
	return vc
}

// SetNillableProductID sets the product edge to Product by id if the given value is not nil.
func (vc *VariationCreate) SetNillableProductID(id *int) *VariationCreate {
	if id != nil {
		vc = vc.SetProductID(*id)
	}
	return vc
}

// SetProduct sets the product edge to Product.
func (vc *VariationCreate) SetProduct(p *Product) *VariationCreate {
	return vc.SetProductID(p.ID)
}

// AddVariantIDs adds the variant edge to Variant by ids.
func (vc *VariationCreate) AddVariantIDs(ids ...int) *VariationCreate {
	vc.mutation.AddVariantIDs(ids...)
	return vc
}

// AddVariant adds the variant edges to Variant.
func (vc *VariationCreate) AddVariant(v ...*Variant) *VariationCreate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vc.AddVariantIDs(ids...)
}

// AddOutboundDealIDs adds the outbound_deals edge to OutboundDeal by ids.
func (vc *VariationCreate) AddOutboundDealIDs(ids ...int) *VariationCreate {
	vc.mutation.AddOutboundDealIDs(ids...)
	return vc
}

// AddOutboundDeals adds the outbound_deals edges to OutboundDeal.
func (vc *VariationCreate) AddOutboundDeals(o ...*OutboundDeal) *VariationCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return vc.AddOutboundDealIDs(ids...)
}

// Mutation returns the VariationMutation object of the builder.
func (vc *VariationCreate) Mutation() *VariationMutation {
	return vc.mutation
}

// Save creates the Variation in the database.
func (vc *VariationCreate) Save(ctx context.Context) (*Variation, error) {
	var (
		err  error
		node *Variation
	)
	vc.defaults()
	if len(vc.hooks) == 0 {
		if err = vc.check(); err != nil {
			return nil, err
		}
		node, err = vc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VariationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vc.check(); err != nil {
				return nil, err
			}
			vc.mutation = mutation
			node, err = vc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(vc.hooks) - 1; i >= 0; i-- {
			mut = vc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VariationCreate) SaveX(ctx context.Context) *Variation {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (vc *VariationCreate) defaults() {
	if _, ok := vc.mutation.Stock(); !ok {
		v := variation.DefaultStock
		vc.mutation.SetStock(v)
	}
	if _, ok := vc.mutation.Price(); !ok {
		v := variation.DefaultPrice
		vc.mutation.SetPrice(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VariationCreate) check() error {
	if _, ok := vc.mutation.Stock(); !ok {
		return &ValidationError{Name: "stock", err: errors.New("ent: missing required field \"stock\"")}
	}
	if _, ok := vc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New("ent: missing required field \"price\"")}
	}
	return nil
}

func (vc *VariationCreate) sqlSave(ctx context.Context) (*Variation, error) {
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (vc *VariationCreate) createSpec() (*Variation, *sqlgraph.CreateSpec) {
	var (
		_node = &Variation{config: vc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: variation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: variation.FieldID,
			},
		}
	)
	if value, ok := vc.mutation.Images(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: variation.FieldImages,
		})
		_node.Images = value
	}
	if value, ok := vc.mutation.Stock(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: variation.FieldStock,
		})
		_node.Stock = value
	}
	if value, ok := vc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: variation.FieldPrice,
		})
		_node.Price = value
	}
	if nodes := vc.mutation.ParentIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.ChildrenIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.ProductIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.VariantIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.OutboundDealsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   variation.OutboundDealsTable,
			Columns: []string{variation.OutboundDealsColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VariationCreateBulk is the builder for creating a bulk of Variation entities.
type VariationCreateBulk struct {
	config
	builders []*VariationCreate
}

// Save creates the Variation entities in the database.
func (vcb *VariationCreateBulk) Save(ctx context.Context) ([]*Variation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Variation, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VariationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (vcb *VariationCreateBulk) SaveX(ctx context.Context) []*Variation {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
