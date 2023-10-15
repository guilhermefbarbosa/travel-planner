package database

import (
	"fmt"
	"gorm.io/gorm"
)

type Operator string

const (
	Equal              Operator = "="
	NotEqual           Operator = "!="
	GreaterThan        Operator = ">"
	GreaterThanOrEqual Operator = ">="
	LessThan           Operator = "<"
	LessThanOrEqual    Operator = "<="
	Like               Operator = "LIKE"
	ILike              Operator = "ILIKE"
	In                 Operator = "IN"
	NotIn              Operator = "NOT IN"
	Wildcard           Operator = "%"
)

type Clause interface {
	Apply(db *gorm.DB) *gorm.DB
	toString() string
}

type FieldClause struct {
	Field    string
	Operator Operator
	Value    any
}

func (f FieldClause) toString() string {
	return fmt.Sprintf("(%s %s ?)", f.Field, f.Operator)
}

func (f FieldClause) Apply(db *gorm.DB) *gorm.DB {
	return db.Where(f.toString(), f.Value)
}

type CombinedClause struct {
	LeftClause  Clause
	Operator    Operator
	RightClause Clause
}

func (c CombinedClause) toString() string {
	return fmt.Sprintf("(%s %s %s)", c.LeftClause.toString(), c.Operator, c.RightClause.toString())
}

func (c CombinedClause) Apply(db *gorm.DB) *gorm.DB {
	return db.Where(c.toString())
}
