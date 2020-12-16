// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/satriahrh/adam-smith/ent/outbounddeal"
	"github.com/satriahrh/adam-smith/ent/outboundtransaction"
	"github.com/satriahrh/adam-smith/ent/predicate"
	"github.com/satriahrh/adam-smith/ent/variation"
)

// OutboundDealUpdate is the builder for updating OutboundDeal entities.
type OutboundDealUpdate struct {
	config
	hooks    []Hook
	mutation *OutboundDealMutation
}

// Where adds a new predicate for the builder.
func (odu *OutboundDealUpdate) Where(ps ...predicate.OutboundDeal) *OutboundDealUpdate {
	odu.mutation.predicates = append(odu.mutation.predicates, ps...)
	return odu
}

// SetQuantity sets the quantity field.
func (odu *OutboundDealUpdate) SetQuantity(u uint) *OutboundDealUpdate {
	odu.mutation.ResetQuantity()
	odu.mutation.SetQuantity(u)
	return odu
}

// AddQuantity adds u to quantity.
func (odu *OutboundDealUpdate) AddQuantity(u uint) *OutboundDealUpdate {
	odu.mutation.AddQuantity(u)
	return odu
}

// SetAmount sets the amount field.
func (odu *OutboundDealUpdate) SetAmount(u uint) *OutboundDealUpdate {
	odu.mutation.ResetAmount()
	odu.mutation.SetAmount(u)
	return odu
}

// AddAmount adds u to amount.
func (odu *OutboundDealUpdate) AddAmount(u uint) *OutboundDealUpdate {
	odu.mutation.AddAmount(u)
	return odu
}

// AddVariantIDs adds the variant edge to Variation by ids.
func (odu *OutboundDealUpdate) AddVariantIDs(ids ...int) *OutboundDealUpdate {
	odu.mutation.AddVariantIDs(ids...)
	return odu
}

// AddVariant adds the variant edges to Variation.
func (odu *OutboundDealUpdate) AddVariant(v ...*Variation) *OutboundDealUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return odu.AddVariantIDs(ids...)
}

// AddTransactionIDs adds the transaction edge to OutboundTransaction by ids.
func (odu *OutboundDealUpdate) AddTransactionIDs(ids ...int) *OutboundDealUpdate {
	odu.mutation.AddTransactionIDs(ids...)
	return odu
}

// AddTransaction adds the transaction edges to OutboundTransaction.
func (odu *OutboundDealUpdate) AddTransaction(o ...*OutboundTransaction) *OutboundDealUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return odu.AddTransactionIDs(ids...)
}

// Mutation returns the OutboundDealMutation object of the builder.
func (odu *OutboundDealUpdate) Mutation() *OutboundDealMutation {
	return odu.mutation
}

// ClearVariant clears all "variant" edges to type Variation.
func (odu *OutboundDealUpdate) ClearVariant() *OutboundDealUpdate {
	odu.mutation.ClearVariant()
	return odu
}

// RemoveVariantIDs removes the variant edge to Variation by ids.
func (odu *OutboundDealUpdate) RemoveVariantIDs(ids ...int) *OutboundDealUpdate {
	odu.mutation.RemoveVariantIDs(ids...)
	return odu
}

// RemoveVariant removes variant edges to Variation.
func (odu *OutboundDealUpdate) RemoveVariant(v ...*Variation) *OutboundDealUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return odu.RemoveVariantIDs(ids...)
}

// ClearTransaction clears all "transaction" edges to type OutboundTransaction.
func (odu *OutboundDealUpdate) ClearTransaction() *OutboundDealUpdate {
	odu.mutation.ClearTransaction()
	return odu
}

// RemoveTransactionIDs removes the transaction edge to OutboundTransaction by ids.
func (odu *OutboundDealUpdate) RemoveTransactionIDs(ids ...int) *OutboundDealUpdate {
	odu.mutation.RemoveTransactionIDs(ids...)
	return odu
}

// RemoveTransaction removes transaction edges to OutboundTransaction.
func (odu *OutboundDealUpdate) RemoveTransaction(o ...*OutboundTransaction) *OutboundDealUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return odu.RemoveTransactionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (odu *OutboundDealUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(odu.hooks) == 0 {
		affected, err = odu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutboundDealMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			odu.mutation = mutation
			affected, err = odu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(odu.hooks) - 1; i >= 0; i-- {
			mut = odu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, odu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (odu *OutboundDealUpdate) SaveX(ctx context.Context) int {
	affected, err := odu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (odu *OutboundDealUpdate) Exec(ctx context.Context) error {
	_, err := odu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (odu *OutboundDealUpdate) ExecX(ctx context.Context) {
	if err := odu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (odu *OutboundDealUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outbounddeal.Table,
			Columns: outbounddeal.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outbounddeal.FieldID,
			},
		},
	}
	if ps := odu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := odu.mutation.Quantity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: outbounddeal.FieldQuantity,
		})
	}
	if value, ok := odu.mutation.AddedQuantity(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: outbounddeal.FieldQuantity,
		})
	}
	if value, ok := odu.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: outbounddeal.FieldAmount,
		})
	}
	if value, ok := odu.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: outbounddeal.FieldAmount,
		})
	}
	if odu.mutation.VariantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   outbounddeal.VariantTable,
			Columns: outbounddeal.VariantPrimaryKey,
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
	if nodes := odu.mutation.RemovedVariantIDs(); len(nodes) > 0 && !odu.mutation.VariantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   outbounddeal.VariantTable,
			Columns: outbounddeal.VariantPrimaryKey,
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
	if nodes := odu.mutation.VariantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   outbounddeal.VariantTable,
			Columns: outbounddeal.VariantPrimaryKey,
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
	if odu.mutation.TransactionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   outbounddeal.TransactionTable,
			Columns: outbounddeal.TransactionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outboundtransaction.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := odu.mutation.RemovedTransactionIDs(); len(nodes) > 0 && !odu.mutation.TransactionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   outbounddeal.TransactionTable,
			Columns: outbounddeal.TransactionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outboundtransaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := odu.mutation.TransactionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   outbounddeal.TransactionTable,
			Columns: outbounddeal.TransactionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outboundtransaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, odu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outbounddeal.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// OutboundDealUpdateOne is the builder for updating a single OutboundDeal entity.
type OutboundDealUpdateOne struct {
	config
	hooks    []Hook
	mutation *OutboundDealMutation
}

// SetQuantity sets the quantity field.
func (oduo *OutboundDealUpdateOne) SetQuantity(u uint) *OutboundDealUpdateOne {
	oduo.mutation.ResetQuantity()
	oduo.mutation.SetQuantity(u)
	return oduo
}

// AddQuantity adds u to quantity.
func (oduo *OutboundDealUpdateOne) AddQuantity(u uint) *OutboundDealUpdateOne {
	oduo.mutation.AddQuantity(u)
	return oduo
}

// SetAmount sets the amount field.
func (oduo *OutboundDealUpdateOne) SetAmount(u uint) *OutboundDealUpdateOne {
	oduo.mutation.ResetAmount()
	oduo.mutation.SetAmount(u)
	return oduo
}

// AddAmount adds u to amount.
func (oduo *OutboundDealUpdateOne) AddAmount(u uint) *OutboundDealUpdateOne {
	oduo.mutation.AddAmount(u)
	return oduo
}

// AddVariantIDs adds the variant edge to Variation by ids.
func (oduo *OutboundDealUpdateOne) AddVariantIDs(ids ...int) *OutboundDealUpdateOne {
	oduo.mutation.AddVariantIDs(ids...)
	return oduo
}

// AddVariant adds the variant edges to Variation.
func (oduo *OutboundDealUpdateOne) AddVariant(v ...*Variation) *OutboundDealUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return oduo.AddVariantIDs(ids...)
}

// AddTransactionIDs adds the transaction edge to OutboundTransaction by ids.
func (oduo *OutboundDealUpdateOne) AddTransactionIDs(ids ...int) *OutboundDealUpdateOne {
	oduo.mutation.AddTransactionIDs(ids...)
	return oduo
}

// AddTransaction adds the transaction edges to OutboundTransaction.
func (oduo *OutboundDealUpdateOne) AddTransaction(o ...*OutboundTransaction) *OutboundDealUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oduo.AddTransactionIDs(ids...)
}

// Mutation returns the OutboundDealMutation object of the builder.
func (oduo *OutboundDealUpdateOne) Mutation() *OutboundDealMutation {
	return oduo.mutation
}

// ClearVariant clears all "variant" edges to type Variation.
func (oduo *OutboundDealUpdateOne) ClearVariant() *OutboundDealUpdateOne {
	oduo.mutation.ClearVariant()
	return oduo
}

// RemoveVariantIDs removes the variant edge to Variation by ids.
func (oduo *OutboundDealUpdateOne) RemoveVariantIDs(ids ...int) *OutboundDealUpdateOne {
	oduo.mutation.RemoveVariantIDs(ids...)
	return oduo
}

// RemoveVariant removes variant edges to Variation.
func (oduo *OutboundDealUpdateOne) RemoveVariant(v ...*Variation) *OutboundDealUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return oduo.RemoveVariantIDs(ids...)
}

// ClearTransaction clears all "transaction" edges to type OutboundTransaction.
func (oduo *OutboundDealUpdateOne) ClearTransaction() *OutboundDealUpdateOne {
	oduo.mutation.ClearTransaction()
	return oduo
}

// RemoveTransactionIDs removes the transaction edge to OutboundTransaction by ids.
func (oduo *OutboundDealUpdateOne) RemoveTransactionIDs(ids ...int) *OutboundDealUpdateOne {
	oduo.mutation.RemoveTransactionIDs(ids...)
	return oduo
}

// RemoveTransaction removes transaction edges to OutboundTransaction.
func (oduo *OutboundDealUpdateOne) RemoveTransaction(o ...*OutboundTransaction) *OutboundDealUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oduo.RemoveTransactionIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (oduo *OutboundDealUpdateOne) Save(ctx context.Context) (*OutboundDeal, error) {
	var (
		err  error
		node *OutboundDeal
	)
	if len(oduo.hooks) == 0 {
		node, err = oduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutboundDealMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oduo.mutation = mutation
			node, err = oduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(oduo.hooks) - 1; i >= 0; i-- {
			mut = oduo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oduo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (oduo *OutboundDealUpdateOne) SaveX(ctx context.Context) *OutboundDeal {
	node, err := oduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oduo *OutboundDealUpdateOne) Exec(ctx context.Context) error {
	_, err := oduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oduo *OutboundDealUpdateOne) ExecX(ctx context.Context) {
	if err := oduo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (oduo *OutboundDealUpdateOne) sqlSave(ctx context.Context) (_node *OutboundDeal, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outbounddeal.Table,
			Columns: outbounddeal.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outbounddeal.FieldID,
			},
		},
	}
	id, ok := oduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing OutboundDeal.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := oduo.mutation.Quantity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: outbounddeal.FieldQuantity,
		})
	}
	if value, ok := oduo.mutation.AddedQuantity(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: outbounddeal.FieldQuantity,
		})
	}
	if value, ok := oduo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: outbounddeal.FieldAmount,
		})
	}
	if value, ok := oduo.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: outbounddeal.FieldAmount,
		})
	}
	if oduo.mutation.VariantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   outbounddeal.VariantTable,
			Columns: outbounddeal.VariantPrimaryKey,
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
	if nodes := oduo.mutation.RemovedVariantIDs(); len(nodes) > 0 && !oduo.mutation.VariantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   outbounddeal.VariantTable,
			Columns: outbounddeal.VariantPrimaryKey,
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
	if nodes := oduo.mutation.VariantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   outbounddeal.VariantTable,
			Columns: outbounddeal.VariantPrimaryKey,
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
	if oduo.mutation.TransactionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   outbounddeal.TransactionTable,
			Columns: outbounddeal.TransactionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outboundtransaction.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oduo.mutation.RemovedTransactionIDs(); len(nodes) > 0 && !oduo.mutation.TransactionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   outbounddeal.TransactionTable,
			Columns: outbounddeal.TransactionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outboundtransaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oduo.mutation.TransactionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   outbounddeal.TransactionTable,
			Columns: outbounddeal.TransactionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: outboundtransaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OutboundDeal{config: oduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, oduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outbounddeal.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
