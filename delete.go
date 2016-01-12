package bsql

import "strings"

type DeleteSQL struct {
	table      *Table
	conditions []Expression
}

func NewDeleteSQL(t *Table) *DeleteSQL {
	return &DeleteSQL{
		table:      t,
		conditions: []Expression{},
	}
}

func (del *DeleteSQL) Where(exps ...Expression) *DeleteSQL {
	del.conditions = append(del.conditions, exps...)
	return del
}

func (del *DeleteSQL) Statment() *Statment {
	if del.table == nil {
		return nil
	}

	fmts := []string{}
	vals := []interface{}{}

	fmts = append(fmts, "DELETE")
	fmts = append(fmts, "FROM")
	fmts = append(fmts, del.table.SQL())

	if len(del.conditions) > 0 {
		condition := AND(del.conditions...).statment()
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
