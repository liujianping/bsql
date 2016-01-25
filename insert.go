package bsql

import "strings"

type InsertSQL struct {
	table *Table
}

func NewInsertSQL(t *Table) *InsertSQL {
	return &InsertSQL{
		table: t,
	}
}

func (insert *InsertSQL) Statment() *Statment {
	if insert.table == nil {
		return nil
	}

	if len(insert.table.Values()) == 0 {
		return nil
	}

	fmts := []string{}
	vals := []interface{}{}
	columns := []string{}
	symbols := []string{}
	for _, col := range insert.table.Values() {
		columns = append(columns, col.Name())
		symbols = append(symbols, "?")
		vals = append(vals, col.Get())
	}

	fmts = append(fmts, "INSERT")
	fmts = append(fmts, "INTO")
	fmts = append(fmts, insert.table.SQL())
	fmts = append(fmts, "(")
	fmts = append(fmts, strings.Join(columns, ", "))
	fmts = append(fmts, ")")
	fmts = append(fmts, "VALUES")
	fmts = append(fmts, "(")
	fmts = append(fmts, strings.Join(symbols, ", "))
	fmts = append(fmts, ")")
	return &Statment{
		format: strings.Join(fmts, " "),
		values: vals,
	}
}
