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
)

// BrandCreate is the builder for creating a Brand entity.
type BrandCreate struct {
	config
	mutation *BrandMutation
	hooks    []Hook
}

// SetCode sets the code field.
func (bc *BrandCreate) SetCode(s string) *BrandCreate {
	bc.mutation.SetCode(s)
	return bc
}

// SetName sets the name field.
func (bc *BrandCreate) SetName(s string) *BrandCreate {
	bc.mutation.SetName(s)
	return bc
}

// AddProductIDs adds the products edge to Product by ids.
func (bc *BrandCreate) AddProductIDs(ids ...int) *BrandCreate {
	bc.mutation.AddProductIDs(ids...)
	return bc
}

// AddProducts adds the products edges to Product.
func (bc *BrandCreate) AddProducts(p ...*Product) *BrandCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bc.AddProductIDs(ids...)
}

// Mutation returns the BrandMutation object of the builder.
func (bc *BrandCreate) Mutation() *BrandMutation {
	return bc.mutation
}

// Save creates the Brand in the database.
func (bc *BrandCreate) Save(ctx context.Context) (*Brand, error) {
	var (
		err  error
		node *Brand
	)
	if len(bc.hooks) == 0 {
		if err = bc.check(); err != nil {
			return nil, err
		}
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BrandMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bc.check(); err != nil {
				return nil, err
			}
			bc.mutation = mutation
			node, err = bc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			mut = bc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BrandCreate) SaveX(ctx context.Context) *Brand {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (bc *BrandCreate) check() error {
	if _, ok := bc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New("ent: missing required field \"code\"")}
	}
	if v, ok := bc.mutation.Code(); ok {
		if err := brand.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf("ent: validator failed for field \"code\": %w", err)}
		}
	}
	if _, ok := bc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := bc.mutation.Name(); ok {
		if err := brand.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (bc *BrandCreate) sqlSave(ctx context.Context) (*Brand, error) {
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bc *BrandCreate) createSpec() (*Brand, *sqlgraph.CreateSpec) {
	var (
		_node = &Brand{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: brand.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: brand.FieldID,
			},
		}
	)
	if value, ok := bc.mutation.Code(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: brand.FieldCode,
		})
		_node.Code = value
	}
	if value, ok := bc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: brand.FieldName,
		})
		_node.Name = value
	}
	if nodes := bc.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   brand.ProductsTable,
			Columns: []string{brand.ProductsColumn},
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
	return _node, _spec
}

// BrandCreateBulk is the builder for creating a bulk of Brand entities.
type BrandCreateBulk struct {
	config
	builders []*BrandCreate
}

// Save creates the Brand entities in the database.
func (bcb *BrandCreateBulk) Save(ctx context.Context) ([]*Brand, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Brand, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BrandMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (bcb *BrandCreateBulk) SaveX(ctx context.Context) []*Brand {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
