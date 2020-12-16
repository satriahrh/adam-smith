// Code generated by entc, DO NOT EDIT.

package outboundtransaction

import (
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/satriahrh/adam-smith/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.OutboundTransaction {
	return predicate.OutboundTransaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OutboundTransaction {
	return predicate.OutboundTransaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OutboundTransaction {
	return predicate.OutboundTransaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OutboundTransaction {
	return predicate.OutboundTransaction(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.OutboundTransaction {
	return predicate.OutboundTransaction(func(s *sql.Selector) {
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
func IDGT(id int) predicate.OutboundTransaction {
	return predicate.OutboundTransaction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OutboundTransaction {
	return predicate.OutboundTransaction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OutboundTransaction {
	return predicate.OutboundTransaction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OutboundTransaction {
	return predicate.OutboundTransaction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v uint) predicate.OutboundTransaction {
	return predicate.OutboundTransaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// ChannelEQ applies the EQ predicate on the "channel" field.
func ChannelEQ(v Channel) predicate.OutboundTransaction {
	return predicate.OutboundTransaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChannel), v))
	})
}

// ChannelNEQ applies the NEQ predicate on the "channel" field.
func ChannelNEQ(v Channel) predicate.OutboundTransaction {
	return predicate.OutboundTransaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldChannel), v))
	})
}

// ChannelIn applies the In predicate on the "channel" field.
func ChannelIn(vs ...Channel) predicate.OutboundTransaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutboundTransaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldChannel), v...))
	})
}

// ChannelNotIn applies the NotIn predicate on the "channel" field.
func ChannelNotIn(vs ...Channel) predicate.OutboundTransaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutboundTransaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldChannel), v...))
	})
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v uint) predicate.OutboundTransaction {
	return predicate.OutboundTransaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v uint) predicate.OutboundTransaction {
	return predicate.OutboundTransaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount), v))
	})
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...uint) predicate.OutboundTransaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutboundTransaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAmount), v...))
	})
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...uint) predicate.OutboundTransaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OutboundTransaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAmount), v...))
	})
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v uint) predicate.OutboundTransaction {
	return predicate.OutboundTransaction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount), v))
	})
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v uint) predicate.OutboundTransaction {
	return predicate.OutboundTransaction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount), v))
	})
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v uint) predicate.OutboundTransaction {
	return predicate.OutboundTransaction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount), v))
	})
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v uint) predicate.OutboundTransaction {
	return predicate.OutboundTransaction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount), v))
	})
}

// HasDeals applies the HasEdge predicate on the "deals" edge.
func HasDeals() predicate.OutboundTransaction {
	return predicate.OutboundTransaction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DealsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DealsTable, DealsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDealsWith applies the HasEdge predicate on the "deals" edge with a given conditions (other predicates).
func HasDealsWith(preds ...predicate.OutboundDeal) predicate.OutboundTransaction {
	return predicate.OutboundTransaction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DealsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DealsTable, DealsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasShipping applies the HasEdge predicate on the "shipping" edge.
func HasShipping() predicate.OutboundTransaction {
	return predicate.OutboundTransaction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ShippingTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ShippingTable, ShippingColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasShippingWith applies the HasEdge predicate on the "shipping" edge with a given conditions (other predicates).
func HasShippingWith(preds ...predicate.OutboundShipping) predicate.OutboundTransaction {
	return predicate.OutboundTransaction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ShippingInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ShippingTable, ShippingColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.OutboundTransaction) predicate.OutboundTransaction {
	return predicate.OutboundTransaction(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.OutboundTransaction) predicate.OutboundTransaction {
	return predicate.OutboundTransaction(func(s *sql.Selector) {
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
func Not(p predicate.OutboundTransaction) predicate.OutboundTransaction {
	return predicate.OutboundTransaction(func(s *sql.Selector) {
		p(s.Not())
	})
}
