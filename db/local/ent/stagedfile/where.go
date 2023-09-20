// Code generated by ent, DO NOT EDIT.

package stagedfile

import (
	"entgo.io/ent/dialect/sql"
	"github.com/Creaft-JP/tit/db/local/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldLTE(FieldID, id))
}

// Path applies equality check predicate on the "path" field. It's identical to PathEQ.
func Path(v string) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldEQ(FieldPath, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldEQ(FieldContent, v))
}

// PathEQ applies the EQ predicate on the "path" field.
func PathEQ(v string) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldEQ(FieldPath, v))
}

// PathNEQ applies the NEQ predicate on the "path" field.
func PathNEQ(v string) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldNEQ(FieldPath, v))
}

// PathIn applies the In predicate on the "path" field.
func PathIn(vs ...string) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldIn(FieldPath, vs...))
}

// PathNotIn applies the NotIn predicate on the "path" field.
func PathNotIn(vs ...string) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldNotIn(FieldPath, vs...))
}

// PathGT applies the GT predicate on the "path" field.
func PathGT(v string) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldGT(FieldPath, v))
}

// PathGTE applies the GTE predicate on the "path" field.
func PathGTE(v string) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldGTE(FieldPath, v))
}

// PathLT applies the LT predicate on the "path" field.
func PathLT(v string) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldLT(FieldPath, v))
}

// PathLTE applies the LTE predicate on the "path" field.
func PathLTE(v string) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldLTE(FieldPath, v))
}

// PathContains applies the Contains predicate on the "path" field.
func PathContains(v string) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldContains(FieldPath, v))
}

// PathHasPrefix applies the HasPrefix predicate on the "path" field.
func PathHasPrefix(v string) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldHasPrefix(FieldPath, v))
}

// PathHasSuffix applies the HasSuffix predicate on the "path" field.
func PathHasSuffix(v string) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldHasSuffix(FieldPath, v))
}

// PathEqualFold applies the EqualFold predicate on the "path" field.
func PathEqualFold(v string) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldEqualFold(FieldPath, v))
}

// PathContainsFold applies the ContainsFold predicate on the "path" field.
func PathContainsFold(v string) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldContainsFold(FieldPath, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.StagedFile {
	return predicate.StagedFile(sql.FieldContainsFold(FieldContent, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.StagedFile) predicate.StagedFile {
	return predicate.StagedFile(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.StagedFile) predicate.StagedFile {
	return predicate.StagedFile(func(s *sql.Selector) {
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
func Not(p predicate.StagedFile) predicate.StagedFile {
	return predicate.StagedFile(func(s *sql.Selector) {
		p(s.Not())
	})
}
