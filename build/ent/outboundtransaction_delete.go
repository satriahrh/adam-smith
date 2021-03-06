// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/satriahrh/adam-smith/build/ent/outboundtransaction"
	"github.com/satriahrh/adam-smith/build/ent/predicate"
)

// OutboundTransactionDelete is the builder for deleting a OutboundTransaction entity.
type OutboundTransactionDelete struct {
	config
	hooks    []Hook
	mutation *OutboundTransactionMutation
}

// Where adds a new predicate to the delete builder.
func (otd *OutboundTransactionDelete) Where(ps ...predicate.OutboundTransaction) *OutboundTransactionDelete {
	otd.mutation.predicates = append(otd.mutation.predicates, ps...)
	return otd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (otd *OutboundTransactionDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(otd.hooks) == 0 {
		affected, err = otd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutboundTransactionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			otd.mutation = mutation
			affected, err = otd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(otd.hooks) - 1; i >= 0; i-- {
			mut = otd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, otd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (otd *OutboundTransactionDelete) ExecX(ctx context.Context) int {
	n, err := otd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (otd *OutboundTransactionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: outboundtransaction.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outboundtransaction.FieldID,
			},
		},
	}
	if ps := otd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, otd.driver, _spec)
}

// OutboundTransactionDeleteOne is the builder for deleting a single OutboundTransaction entity.
type OutboundTransactionDeleteOne struct {
	otd *OutboundTransactionDelete
}

// Exec executes the deletion query.
func (otdo *OutboundTransactionDeleteOne) Exec(ctx context.Context) error {
	n, err := otdo.otd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{outboundtransaction.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (otdo *OutboundTransactionDeleteOne) ExecX(ctx context.Context) {
	otdo.otd.ExecX(ctx)
}
