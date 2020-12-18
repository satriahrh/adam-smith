// Code generated by entc, DO NOT EDIT.

package variation

import (
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/satriahrh/adam-smith/build/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id uint64) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Stock applies equality check predicate on the "stock" field. It's identical to StockEQ.
func Stock(v uint8) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStock), v))
	})
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v uint) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// ImagesIsNil applies the IsNil predicate on the "images" field.
func ImagesIsNil() predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldImages)))
	})
}

// ImagesNotNil applies the NotNil predicate on the "images" field.
func ImagesNotNil() predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldImages)))
	})
}

// StockEQ applies the EQ predicate on the "stock" field.
func StockEQ(v uint8) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStock), v))
	})
}

// StockNEQ applies the NEQ predicate on the "stock" field.
func StockNEQ(v uint8) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStock), v))
	})
}

// StockIn applies the In predicate on the "stock" field.
func StockIn(vs ...uint8) predicate.Variation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Variation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStock), v...))
	})
}

// StockNotIn applies the NotIn predicate on the "stock" field.
func StockNotIn(vs ...uint8) predicate.Variation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Variation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStock), v...))
	})
}

// StockGT applies the GT predicate on the "stock" field.
func StockGT(v uint8) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStock), v))
	})
}

// StockGTE applies the GTE predicate on the "stock" field.
func StockGTE(v uint8) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStock), v))
	})
}

// StockLT applies the LT predicate on the "stock" field.
func StockLT(v uint8) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStock), v))
	})
}

// StockLTE applies the LTE predicate on the "stock" field.
func StockLTE(v uint8) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStock), v))
	})
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v uint) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v uint) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrice), v))
	})
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...uint) predicate.Variation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Variation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPrice), v...))
	})
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...uint) predicate.Variation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Variation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPrice), v...))
	})
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v uint) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrice), v))
	})
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v uint) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrice), v))
	})
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v uint) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrice), v))
	})
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v uint) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrice), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.Variation) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChildrenTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.Variation) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasVariant applies the HasEdge predicate on the "variant" edge.
func HasVariant() predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VariantTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, VariantTable, VariantColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVariantWith applies the HasEdge predicate on the "variant" edge with a given conditions (other predicates).
func HasVariantWith(preds ...predicate.Variant) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VariantInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, VariantTable, VariantColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProduct applies the HasEdge predicate on the "product" edge.
func HasProduct() predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProductTable, ProductColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductWith applies the HasEdge predicate on the "product" edge with a given conditions (other predicates).
func HasProductWith(preds ...predicate.Product) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProductTable, ProductColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOutboundDeals applies the HasEdge predicate on the "outbound_deals" edge.
func HasOutboundDeals() predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OutboundDealsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, OutboundDealsTable, OutboundDealsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOutboundDealsWith applies the HasEdge predicate on the "outbound_deals" edge with a given conditions (other predicates).
func HasOutboundDealsWith(preds ...predicate.OutboundDeal) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OutboundDealsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, OutboundDealsTable, OutboundDealsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Variation) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Variation) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Variation) predicate.Variation {
	return predicate.Variation(func(s *sql.Selector) {
		p(s.Not())
	})
}
