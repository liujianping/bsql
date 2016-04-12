package bsql

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTable(t *testing.T) {
	t1 := TABLE("test").As("t1")

	assert.Equal(t, t1.Name(), "`test`")
	assert.Equal(t, t1.SQL(), "`test` AS t1")

	col := t1.Column("name").As("test_name")

	assert.Equal(t, col.Name(), "t1.name")
	assert.Equal(t, col.SQL(), "t1.name AS test_name")
}

func TestInsertSQL(t *testing.T) {
	//! insert sql
	t1 := TABLE("accounts")

	t1.Column("mailbox").Set("demo@demo.cn")
	t1.Column("age").Set(20)
	t1.Column("create_at").Set(time.Now())

	insert := NewInsertSQL(t1).Statment()
	assert.NotNil(t, insert)

	log.Println("insert format:", insert.SQLFormat())
	log.Println("insert params:", insert.SQLParams())
}

func TestUpdateSQL(t *testing.T) {
	//! update sql
	t2 := TABLE("accounts")

	t2.Column("mailbox").Set("demo@demo.cn")
	t2.Column("age").Set(30)

	update := NewUpdateSQL(t2).Where(EQ(t2.Column("id").Name(), 10)).Statment()
	assert.NotNil(t, update)

	log.Println("update format:", update.SQLFormat())
	log.Println("update params:", update.SQLParams())
}

func TestDeleteSQL(t *testing.T) {
	//! delete sql
	t3 := TABLE("accounts")

	del := NewDeleteSQL(t3).Where(EQ(t3.Column("id").Name(), 10)).Statment()

	log.Println("delete format:", del.SQLFormat())
	log.Println("delete params:", del.SQLParams())
}

func TestQuerySQL(t *testing.T) {
	//! query sql
	t4 := TABLE("accounts").As("a")
	t4.Columns("name", "mailbox", "age")

	t5 := TABLE("account_comments").As("c")
	t5.Columns("title", "content")

	query := NewQuerySQL(t4)
	//left join
	query.LeftJoin(t5).On("a.id = c.account_id")

	query.Where(EQ(t4.Column("id").Name(), 10))
	query.OrderByAsc(t4.Column("age"))
	query.OrderByDesc(t5.Column("id"))
	assert.NotNil(t, query)

	log.Println("query format:", query.Statment().SQLFormat())
	log.Println("query params:", query.Statment().SQLParams())
}
