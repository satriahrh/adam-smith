// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/satriahrh/adam-smith/build/ent/predicate"
	"github.com/satriahrh/adam-smith/build/ent/variation"
)

// VariationDelete is the builder for deleting a Variation entity.
type VariationDelete struct {
	config
	hooks    []Hook
	mutation *VariationMutation
}

// Where adds a new predicate to the delete builder.
func (vd *VariationDelete) Where(ps ...predicate.Variation) *VariationDelete {
	vd.mutation.predicates = append(vd.mutation.predicates, ps...)
	return vd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vd *VariationDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vd.hooks) == 0 {
		affected, err = vd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VariationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vd.mutation = mutation
			affected, err = vd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vd.hooks) - 1; i >= 0; i-- {
			mut = vd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (vd *VariationDelete) ExecX(ctx context.Context) int {
	n, err := vd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vd *VariationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: variation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: variation.FieldID,
			},
		},
	}
	if ps := vd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, vd.driver, _spec)
}

// VariationDeleteOne is the builder for deleting a single Variation entity.
type VariationDeleteOne struct {
	vd *VariationDelete
}

// Exec executes the deletion query.
func (vdo *VariationDeleteOne) Exec(ctx context.Context) error {
	n, err := vdo.vd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{variation.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vdo *VariationDeleteOne) ExecX(ctx context.Context) {
	vdo.vd.ExecX(ctx)
}
