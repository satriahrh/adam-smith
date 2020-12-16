// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/satriahrh/adam-smith/ent/brand"
	"github.com/satriahrh/adam-smith/ent/product"
	"github.com/satriahrh/adam-smith/ent/schema"
	"github.com/satriahrh/adam-smith/ent/variation"
)

// ProductCreate is the builder for creating a Product entity.
type ProductCreate struct {
	config
	mutation *ProductMutation
	hooks    []Hook
}

// SetSku sets the sku field.
func (pc *ProductCreate) SetSku(s string) *ProductCreate {
	pc.mutation.SetSku(s)
	return pc
}

// SetName sets the name field.
func (pc *ProductCreate) SetName(s string) *ProductCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetDescriptions sets the descriptions field.
func (pc *ProductCreate) SetDescriptions(sd []schema.ProductDescription) *ProductCreate {
	pc.mutation.SetDescriptions(sd)
	return pc
}

// SetImages sets the images field.
func (pc *ProductCreate) SetImages(si schema.ProductImages) *ProductCreate {
	pc.mutation.SetImages(si)
	return pc
}

// SetMarketplaces sets the marketplaces field.
func (pc *ProductCreate) SetMarketplaces(sm schema.ProductMarketplace) *ProductCreate {
	pc.mutation.SetMarketplaces(sm)
	return pc
}

// AddVariationIDs adds the variations edge to Variation by ids.
func (pc *ProductCreate) AddVariationIDs(ids ...int) *ProductCreate {
	pc.mutation.AddVariationIDs(ids...)
	return pc
}

// AddVariations adds the variations edges to Variation.
func (pc *ProductCreate) AddVariations(v ...*Variation) *ProductCreate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return pc.AddVariationIDs(ids...)
}

// SetBrandID sets the brand edge to Brand by id.
func (pc *ProductCreate) SetBrandID(id int) *ProductCreate {
	pc.mutation.SetBrandID(id)
	return pc
}

// SetNillableBrandID sets the brand edge to Brand by id if the given value is not nil.
func (pc *ProductCreate) SetNillableBrandID(id *int) *ProductCreate {
	if id != nil {
		pc = pc.SetBrandID(*id)
	}
	return pc
}

// SetBrand sets the brand edge to Brand.
func (pc *ProductCreate) SetBrand(b *Brand) *ProductCreate {
	return pc.SetBrandID(b.ID)
}

// Mutation returns the ProductMutation object of the builder.
func (pc *ProductCreate) Mutation() *ProductMutation {
	return pc.mutation
}

// Save creates the Product in the database.
func (pc *ProductCreate) Save(ctx context.Context) (*Product, error) {
	var (
		err  error
		node *Product
	)
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProductCreate) SaveX(ctx context.Context) *Product {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProductCreate) check() error {
	if _, ok := pc.mutation.Sku(); !ok {
		return &ValidationError{Name: "sku", err: errors.New("ent: missing required field \"sku\"")}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := pc.mutation.Descriptions(); !ok {
		return &ValidationError{Name: "descriptions", err: errors.New("ent: missing required field \"descriptions\"")}
	}
	if _, ok := pc.mutation.Images(); !ok {
		return &ValidationError{Name: "images", err: errors.New("ent: missing required field \"images\"")}
	}
	if _, ok := pc.mutation.Marketplaces(); !ok {
		return &ValidationError{Name: "marketplaces", err: errors.New("ent: missing required field \"marketplaces\"")}
	}
	return nil
}

func (pc *ProductCreate) sqlSave(ctx context.Context) (*Product, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pc *ProductCreate) createSpec() (*Product, *sqlgraph.CreateSpec) {
	var (
		_node = &Product{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: product.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: product.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Sku(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldSku,
		})
		_node.Sku = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldName,
		})
		_node.Name = value
	}
	if value, ok := pc.mutation.Descriptions(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: product.FieldDescriptions,
		})
		_node.Descriptions = value
	}
	if value, ok := pc.mutation.Images(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: product.FieldImages,
		})
		_node.Images = value
	}
	if value, ok := pc.mutation.Marketplaces(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: product.FieldMarketplaces,
		})
		_node.Marketplaces = value
	}
	if nodes := pc.mutation.VariationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.VariationsTable,
			Columns: []string{product.VariationsColumn},
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
	if nodes := pc.mutation.BrandIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.BrandTable,
			Columns: []string{product.BrandColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: brand.FieldID,
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

// ProductCreateBulk is the builder for creating a bulk of Product entities.
type ProductCreateBulk struct {
	config
	builders []*ProductCreate
}

// Save creates the Product entities in the database.
func (pcb *ProductCreateBulk) Save(ctx context.Context) ([]*Product, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Product, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (pcb *ProductCreateBulk) SaveX(ctx context.Context) []*Product {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
