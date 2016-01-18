package bsql

import "strings"

type QuerySQL struct {
	table         *Table
	joins         []*JoinTable
	conditions    []Expression
	order_asc     bool
	order_columns []*Column
	group_columns []*Column
	offset        int64
	limit         int64
}

func NewQuerySQL(table *Table) *QuerySQL {
	return &QuerySQL{
		table:         table,
		joins:         []*JoinTable{},
		conditions:    []Expression{},
		order_asc:     false,
		order_columns: []*Column{},
		group_columns: []*Column{},
	}
}

func (query *QuerySQL) Table(name string) *Table {
	if query.table == nil {
		query.table = TABLE(name)
		return query.table
	}

	if query.table.name == name {
		return query.table
	}

	for _, join := range query.joins {
		if join.table.name == name {
			return join.table
		}
	}
	return nil
}

func (query *QuerySQL) LeftJoin(table *Table) *JoinTable {
	join := LEFT(table)
	query.joins = append(query.joins, join)
	return join
}

func (query *QuerySQL) RightJoin(table *Table) *JoinTable {
	join := RIGHT(table)
	query.joins = append(query.joins, join)
	return join
}

func (query *QuerySQL) InnerJoin(table *Table) *JoinTable {
	join := INNER(table)
	query.joins = append(query.joins, join)
	return join
}

func (query *QuerySQL) Join(join *JoinTable) *QuerySQL {
	query.joins = append(query.joins, join)
	return query
}

type JoinTable struct {
	symbol string
	table  *Table
	on     string
}

func (join *JoinTable) On(exp string) {
	join.on = exp
}

func (join *JoinTable) SQL() string {
	fmts := []string{}
	fmts = append(fmts, join.symbol)
	fmts = append(fmts, join.table.SQL())
	fmts = append(fmts, "ON")
	fmts = append(fmts, join.on)
	return strings.Join(fmts, " ")
}

func LEFT(t *Table) *JoinTable {
	return &JoinTable{
		symbol: "LEFT JOIN",
		table:  t,
		on:     "",
	}
}

func RIGHT(t *Table) *JoinTable {
	return &JoinTable{
		symbol: "RIGHT JOIN",
		table:  t,
		on:     "",
	}
}

func INNER(t *Table) *JoinTable {
	return &JoinTable{
		symbol: "INNER JOIN",
		table:  t,
		on:     "",
	}
}

func (query *QuerySQL) Where(exps ...Expression) *QuerySQL {
	query.conditions = append(query.conditions, exps...)
	return query
}

func (query *QuerySQL) OrderByAsc(columns ...*Column) *QuerySQL {
	if query.order_asc != true {
		query.order_asc = true
		query.order_columns = []*Column{}
	}
	query.order_columns = append(query.order_columns, columns...)
	return query
}

func (query *QuerySQL) OrderByDesc(columns ...*Column) *QuerySQL {
	if query.order_asc != false {
		query.order_asc = false
		query.order_columns = []*Column{}
	}
	query.order_columns = append(query.order_columns, columns...)
	return query
}

func (query *QuerySQL) Limit(page_no, page_size int64) *QuerySQL {
	query.limit = page_size
	query.offset = page_no * page_size
	return query
}

func (query *QuerySQL) GroupBy(columns ...*Column) *QuerySQL {
	query.group_columns = append(query.group_columns, columns...)
	return query
}

func (query *QuerySQL) Statment() *Statment {
	fmts := []string{}
	vals := []interface{}{}

	fmts = append(fmts, "SELECT")
	fmts = append(fmts, query.table.ColumnsSQL())

	for _, join := range query.joins {
		fmts = append(fmts, ",")
		fmts = append(fmts, join.table.ColumnsSQL())
	}

	fmts = append(fmts, "FROM")
	fmts = append(fmts, query.table.SQL())

	for _, join := range query.joins {
		fmts = append(fmts, join.SQL())
	}

	if len(query.conditions) > 0 {
		stmt := AND(query.conditions...).statment()
		if stmt != nil {
			fmts = append(fmts, "WHERE")
			fmts = append(fmts, stmt.SQLFormat())
			vals = append(vals, stmt.SQLParams()...)
		}
	}

	if len(query.group_columns) > 0 {
		fmts = append(fmts, "GROUP BY")

		fields := []string{}
		for _, col := range query.group_columns {
			fields = append(fields, col.Name())
		}
		fmts = append(fmts, strings.Join(fields, ", "))
	}

	if len(query.order_columns) > 0 {
		fmts = append(fmts, "ORDER BY")

		fields := []string{}
		for _, col := range query.order_columns {
			fields = append(fields, col.Name())
		}
		fmts = append(fmts, strings.Join(fields, ", "))

		if query.order_asc {
			fmts = append(fmts, "ASC")
		} else {
			fmts = append(fmts, "DESC")
		}
	}

	if query.limit > 0 {
		fmts = append(fmts, "LIMIT ?,?")
		vals = append(vals, query.offset, query.limit)
	}

	return &Statment{
		format: strings.Join(fmts, " "),
		values: vals,
	}
}

func (query *QuerySQL) CountStatment() *Statment {
	fmts := []string{}
	vals := []interface{}{}

	fmts = append(fmts, "SELECT")
	fmts = append(fmts, "COUNT(*)")
	fmts = append(fmts, "FROM")
	fmts = append(fmts, query.table.SQL())

	for _, join := range query.joins {
		fmts = append(fmts, join.symbol)
		fmts = append(fmts, join.table.SQL())
		fmts = append(fmts, "ON")
		fmts = append(fmts, join.on)
	}

	if len(query.conditions) > 0 {
		stmt := AND(query.conditions...).statment()
		if stmt != nil {
			fmts = append(fmts, "WHERE")
			fmts = append(fmts, stmt.SQLFormat())
			vals = append(vals, stmt.SQLParams()...)
		}
	}

	return &Statment{
		format: strings.Join(fmts, " "),
		values: vals,
	}
}
