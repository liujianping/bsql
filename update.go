package bsql

import (
	"fmt"
	"strings"
)

type UpdateSQL struct {
	table      *Table
	conditions []Expression
}

func NewUpdateSQL(t *Table) *UpdateSQL {
	return &UpdateSQL{
		table:      t,
		conditions: []Expression{},
	}
}

func (update *UpdateSQL) Where(exps ...Expression) *UpdateSQL {
	update.conditions = append(update.conditions, exps...)
	return update
}

func (update *UpdateSQL) Statment() *Statment {
	if update.table == nil {
		return nil
	}

	if len(update.table.Values()) == 0 {
		return nil
	}

	fmts := []string{}
	vals := []interface{}{}

	fmts = append(fmts, "UPDATE")
	fmts = append(fmts, update.table.SQL())
	fmts = append(fmts, "SET")

	symbols := []string{}
	for _, col := range update.table.Values() {
		symbols = append(symbols, fmt.Sprintf("%s = ?", col.Name()))
		vals = append(vals, col.Get())
	}
	fmts = append(fmts, strings.Join(symbols, ", "))

	if len(update.conditions) > 0 {
		condition := AND(update.conditions...).statment()
		if condition != nil {
			fmts = append(fmts, "WHERE")
			fmts = append(fmts, condition.SQLFormat())
			vals = append(vals, condition.SQLParams()...)
		}
	}

	return &Statment{
		format: strings.Join(fmts, " "),
		values: vals,
	}
}
