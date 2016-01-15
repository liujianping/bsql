package bsql

import (
	"fmt"
	"strings"
)

type Column struct {
	table *Table
	name  string
	alias string
	value interface{}
}

func COLUMN(tb *Table, name string) *Column {
	return &Column{
		table: tb,
		name:  strings.ToLower(name),
	}
}

func (col *Column) Name() string {
	if col.table.alias != "" {
		return fmt.Sprintf("%s.%s", col.table.alias, col.name)
	}
	return SQLQuote(col.name)
}

func (col *Column) As(as string) *Column {
	col.alias = as
	return col
}

func (col *Column) Set(val interface{}) {
	col.value = val
}

func (col *Column) Get() interface{} {
	return col.value
}

func (col *Column) SQL() string {
	if col.alias != "" {
		return fmt.Sprintf("%s AS %s", col.Name(), col.alias)
	}
	return col.Name()
}

type Table struct {
	name    string
	alias   string
	primary string
	columns map[string]*Column
}

func TABLE(name string) *Table {
	return &Table{
		name:    strings.ToLower(name),
		primary: "",
		columns: make(map[string]*Column),
	}
}

func (t *Table) Name() string {
	return SQLQuote(t.name)
}

func (t *Table) Alias() string {
	return t.alias
}

func (t *Table) PrimaryKey(pk string) *Table {
	t.primary = pk
	return t
}

func (t *Table) As(as string) *Table {
	t.alias = as
	return t
}

func (t *Table) SQL() string {
	if t.alias != "" {
		return fmt.Sprintf("%s AS %s", t.Name(), t.alias)
	}
	return t.Name()
}

func (t *Table) ColumnsSQL() string {
	if len(t.columns) > 0 {
		columns := []string{}
		for _, col := range t.columns {
			columns = append(columns, col.SQL())
		}
		return strings.Join(columns, ", ")
	}
	return "*"
}

func (t *Table) Columns(columns ...string) *Table {
	for _, col := range columns {
		t.columns[col] = COLUMN(t, col)
	}
	return t
}

func (t *Table) Values() []*Column {
	values := []*Column{}
	for _, col := range t.columns {
		if col.Get() != nil {
			values = append(values, col)
		}
	}
	return values
}

func (t *Table) Column(name string) *Column {
	if col, ok := t.columns[name]; ok {
		return col
	}
	col := COLUMN(t, name)
	t.columns[name] = col
	return col
}
