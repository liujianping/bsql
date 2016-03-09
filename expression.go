package bsql

import (
	"fmt"
	"strings"
)

type Expression struct {
	insides []Expression
	field   string
	symbol  string
	value   interface{}
}

func (exp Expression) SQL() (string, error) {
	val, ok := exp.value.(string)
	if !ok {
		return "", fmt.Errorf("expression value (%v) not string type", exp.value)
	}

	switch strings.ToUpper(exp.symbol) {
	case "AND", "OR", "LIKE", "IN":
		return "", fmt.Errorf("expression (%s) unsupport SQL out", exp.symbol)
	}
	return fmt.Sprintf("%s %s %s", SQLQuote(exp.field), exp.symbol, val), nil
}

func BinaryExpression(field, symbol string, value interface{}) Expression {
	switch strings.ToUpper(symbol) {
	case "LIKE":
		return LIKE(field, value)
	case "LLIKE":
		return LLIKE(field, value)
	case "RLIKE":
		return RLIKE(field, value)
	case "EQ":
		return EQ(field, value)
	case "NE":
		return NE(field, value)
	case "IN":
		return IN(field, value)
	case "GT":
		return GT(field, value)
	case "GE":
		return GE(field, value)
	case "LT":
		return LT(field, value)
	case "LE":
		return LE(field, value)
	}
	return NONE(field, value)
}

func (exp Expression) statment() *Statment {
	switch strings.ToUpper(exp.symbol) {
	case "AND":
		var statments []*Statment
		for _, inside := range exp.insides {
			if inside.statment() != nil {
				statments = append(statments, inside.statment())
			}
		}
		return Brace(Join(statments, ") AND ("))
	case "OR":
		var statments []*Statment
		for _, inside := range exp.insides {
			if inside.statment() != nil {
				statments = append(statments, inside.statment())
			}
		}
		return Brace(Join(statments, ") OR ("))
	case "LIKE":
		return &Statment{
			format: fmt.Sprintf("%s LIKE ?", SQLQuote(exp.field)),
			values: []interface{}{exp.value},
		}
	case "IN":
		return nil
	default:
		return &Statment{
			format: fmt.Sprintf("%s %s ?", SQLQuote(exp.field), exp.symbol),
			values: []interface{}{exp.value},
		}
	}
	return nil
}

func AND(exps ...Expression) Expression {
	return Expression{
		insides: exps,
		symbol:  "AND",
	}
}

func OR(exps ...Expression) Expression {
	return Expression{
		insides: exps,
		symbol:  "OR",
	}
}

func NONE(field string, value interface{}) Expression {
	return Expression{
		field:  "",
		symbol: "",
		value:  "",
	}
}

func RANGE(field string, start, end interface{}) Expression {
	return AND(GT(field, start), LT(field, end))
}

func LIKE(field string, value interface{}) Expression {
	return Expression{
		field:  field,
		symbol: "LIKE",
		value:  fmt.Sprintf("%%%s%%", value),
	}
}

func LLIKE(field string, value interface{}) Expression {
	return Expression{
		field:  field,
		symbol: "LIKE",
		value:  fmt.Sprintf("%%%s", value),
	}
}

func RLIKE(field string, value interface{}) Expression {
	return Expression{
		field:  field,
		symbol: "LIKE",
		value:  fmt.Sprintf("%s%%", value),
	}
}

func EQ(field string, value interface{}) Expression {
	return Expression{
		field:  field,
		symbol: "=",
		value:  value,
	}
}

func NE(field string, value interface{}) Expression {
	return Expression{
		field:  field,
		symbol: "!=",
		value:  value,
	}
}

func IN(field string, value interface{}) Expression {
	return Expression{
		field:  field,
		symbol: "IN",
		value:  value,
	}
}

func GT(field string, value interface{}) Expression {
	return Expression{
		field:  field,
		symbol: ">",
		value:  value,
	}
}

func GE(field string, value interface{}) Expression {
	return Expression{
		field:  field,
		symbol: ">=",
		value:  value,
	}
}

func LT(field string, value interface{}) Expression {
	return Expression{
		field:  field,
		symbol: "<",
		value:  value,
	}
}

func LE(field string, value interface{}) Expression {
	return Expression{
		field:  field,
		symbol: "<=",
		value:  value,
	}
}
