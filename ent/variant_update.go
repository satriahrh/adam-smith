// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/satriahrh/adam-smith/ent/predicate"
	"github.com/satriahrh/adam-smith/ent/variant"
	"github.com/satriahrh/adam-smith/ent/variation"
)

// VariantUpdate is the builder for updating Variant entities.
type VariantUpdate struct {
	config
	hooks    []Hook
	mutation *VariantMutation
}

// Where adds a new predicate for the builder.
func (vu *VariantUpdate) Where(ps ...predicate.Variant) *VariantUpdate {
	vu.mutation.predicates = append(vu.mutation.predicates, ps...)
	return vu
}

// SetType sets the type field.
func (vu *VariantUpdate) SetType(v variant.Type) *VariantUpdate {
	vu.mutation.SetType(v)
	return vu
}

// SetValue sets the value field.
func (vu *VariantUpdate) SetValue(s string) *VariantUpdate {
	vu.mutation.SetValue(s)
	return vu
}

// AddVariantUseIDs adds the variant_uses edge to Variation by ids.
func (vu *VariantUpdate) AddVariantUseIDs(ids ...int) *VariantUpdate {
	vu.mutation.AddVariantUseIDs(ids...)
	return vu
}

// AddVariantUses adds the variant_uses edges to Variation.
func (vu *VariantUpdate) AddVariantUses(v ...*Variation) *VariantUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vu.AddVariantUseIDs(ids...)
}

// Mutation returns the VariantMutation object of the builder.
func (vu *VariantUpdate) Mutation() *VariantMutation {
	return vu.mutation
}

// ClearVariantUses clears all "variant_uses" edges to type Variation.
func (vu *VariantUpdate) ClearVariantUses() *VariantUpdate {
	vu.mutation.ClearVariantUses()
	return vu
}

// RemoveVariantUseIDs removes the variant_uses edge to Variation by ids.
func (vu *VariantUpdate) RemoveVariantUseIDs(ids ...int) *VariantUpdate {
	vu.mutation.RemoveVariantUseIDs(ids...)
	return vu
}

// RemoveVariantUses removes variant_uses edges to Variation.
func (vu *VariantUpdate) RemoveVariantUses(v ...*Variation) *VariantUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vu.RemoveVariantUseIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vu *VariantUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vu.hooks) == 0 {
		if err = vu.check(); err != nil {
			return 0, err
		}
		affected, err = vu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VariantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vu.check(); err != nil {
				return 0, err
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
func (vu *VariantUpdate) SaveX(ctx context.Context) int {
	affected, err := vu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vu *VariantUpdate) Exec(ctx context.Context) error {
	_, err := vu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vu *VariantUpdate) ExecX(ctx context.Context) {
	if err := vu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vu *VariantUpdate) check() error {
	if v, ok := vu.mutation.GetType(); ok {
		if err := variant.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	return nil
}

func (vu *VariantUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   variant.Table,
			Columns: variant.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: variant.FieldID,
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
	if value, ok := vu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: variant.FieldType,
		})
	}
	if value, ok := vu.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: variant.FieldValue,
		})
	}
	if vu.mutation.VariantUsesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   variant.VariantUsesTable,
			Columns: variant.VariantUsesPrimaryKey,
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
	if nodes := vu.mutation.RemovedVariantUsesIDs(); len(nodes) > 0 && !vu.mutation.VariantUsesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   variant.VariantUsesTable,
			Columns: variant.VariantUsesPrimaryKey,
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
	if nodes := vu.mutation.VariantUsesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   variant.VariantUsesTable,
			Columns: variant.VariantUsesPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, vu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{variant.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// VariantUpdateOne is the builder for updating a single Variant entity.
type VariantUpdateOne struct {
	config
	hooks    []Hook
	mutation *VariantMutation
}

// SetType sets the type field.
func (vuo *VariantUpdateOne) SetType(v variant.Type) *VariantUpdateOne {
	vuo.mutation.SetType(v)
	return vuo
}

// SetValue sets the value field.
func (vuo *VariantUpdateOne) SetValue(s string) *VariantUpdateOne {
	vuo.mutation.SetValue(s)
	return vuo
}

// AddVariantUseIDs adds the variant_uses edge to Variation by ids.
func (vuo *VariantUpdateOne) AddVariantUseIDs(ids ...int) *VariantUpdateOne {
	vuo.mutation.AddVariantUseIDs(ids...)
	return vuo
}

// AddVariantUses adds the variant_uses edges to Variation.
func (vuo *VariantUpdateOne) AddVariantUses(v ...*Variation) *VariantUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vuo.AddVariantUseIDs(ids...)
}

// Mutation returns the VariantMutation object of the builder.
func (vuo *VariantUpdateOne) Mutation() *VariantMutation {
	return vuo.mutation
}

// ClearVariantUses clears all "variant_uses" edges to type Variation.
func (vuo *VariantUpdateOne) ClearVariantUses() *VariantUpdateOne {
	vuo.mutation.ClearVariantUses()
	return vuo
}

// RemoveVariantUseIDs removes the variant_uses edge to Variation by ids.
func (vuo *VariantUpdateOne) RemoveVariantUseIDs(ids ...int) *VariantUpdateOne {
	vuo.mutation.RemoveVariantUseIDs(ids...)
	return vuo
}

// RemoveVariantUses removes variant_uses edges to Variation.
func (vuo *VariantUpdateOne) RemoveVariantUses(v ...*Variation) *VariantUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vuo.RemoveVariantUseIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (vuo *VariantUpdateOne) Save(ctx context.Context) (*Variant, error) {
	var (
		err  error
		node *Variant
	)
	if len(vuo.hooks) == 0 {
		if err = vuo.check(); err != nil {
			return nil, err
		}
		node, err = vuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VariantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vuo.check(); err != nil {
				return nil, err
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
func (vuo *VariantUpdateOne) SaveX(ctx context.Context) *Variant {
	node, err := vuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vuo *VariantUpdateOne) Exec(ctx context.Context) error {
	_, err := vuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vuo *VariantUpdateOne) ExecX(ctx context.Context) {
	if err := vuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vuo *VariantUpdateOne) check() error {
	if v, ok := vuo.mutation.GetType(); ok {
		if err := variant.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	return nil
}

func (vuo *VariantUpdateOne) sqlSave(ctx context.Context) (_node *Variant, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   variant.Table,
			Columns: variant.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: variant.FieldID,
			},
		},
	}
	id, ok := vuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Variant.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := vuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: variant.FieldType,
		})
	}
	if value, ok := vuo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: variant.FieldValue,
		})
	}
	if vuo.mutation.VariantUsesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   variant.VariantUsesTable,
			Columns: variant.VariantUsesPrimaryKey,
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
	if nodes := vuo.mutation.RemovedVariantUsesIDs(); len(nodes) > 0 && !vuo.mutation.VariantUsesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   variant.VariantUsesTable,
			Columns: variant.VariantUsesPrimaryKey,
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
	if nodes := vuo.mutation.VariantUsesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   variant.VariantUsesTable,
			Columns: variant.VariantUsesPrimaryKey,
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
	_node = &Variant{config: vuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, vuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{variant.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}