package bsql

import "strings"

type SQL interface {
	Statment() *Statment
}

type Statment struct {
	format string
	values []interface{}
}

func (statment *Statment) SQLFormat() string {
	return statment.format
}

func (statment *Statment) SQLParams() []interface{} {
	return statment.values
}

func (statment *Statment) Join(join *Statment, sep string) *Statment {
	fmts := []string{statment.format, join.format}
	return &Statment{
		format: strings.Join(fmts, sep),
		values: append(statment.values, join),
	}
}

func SQLQuote(name string) string {
	return "`" + name + "`"
}

func Join(statments []*Statment, sep string) *Statment {
	if len(statments) == 0 {
		return nil
	}

	if len(statments) == 1 {
		return statments[0]
	}

	var fmts []string
	var vals []interface{}
	for _, statment := range statments {
		fmts = append(fmts, statment.format)
		vals = append(vals, statment.values...)
	}

	return &Statment{
		format: strings.Join(fmts, sep),
		values: vals,
	}
}
