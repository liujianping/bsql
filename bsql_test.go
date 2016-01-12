package bsql

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTable(t *testing.T) {
	t1 := TABLE("test").AS("t1")

	assert.Equal(t, t1.Name(), "`test`")
	assert.Equal(t, t1.SQL(), "`test` AS t1")

	col := t1.Column("name").AS("test_name")

	assert.Equal(t, col.Name(), "t1.name")
	assert.Equal(t, col.SQL(), "t1.name AS test_name")
}

func TestInsertSQL(t *testing.T) {
	t1 := TABLE("account")

	t1.Column("mailbox").Set("ljp@sina.com")
	t1.Column("age").Set(30)
	t1.Column("create_at").Set(time.Now())

	insert := NewInsertSQL(t1)
	log.Println(insert.Statment())
	assert.Equal(t, "INSERT INTO `account` ( `mailbox`, `age`, `create_at` ) VALUES ( ?, ?, ? )", insert.Statment().SQLFormat())
}

func TestUpdateSQL(t *testing.T) {
	t1 := TABLE("account")

	t1.Column("mailbox").Set("ljp@sina.com")
	t1.Column("age").Set(30)
	t1.Column("create_at").Set(time.Now())

	update := NewUpdateSQL(t1)
	log.Println(update.Statment())
	assert.Equal(t, "UPDATE `account` SET `mailbox` = ?, `age` = ?, `create_at` = ?", update.Statment().SQLFormat())
}

func TestDeleteSQL(t *testing.T) {
	t1 := TABLE("account")
	del := NewDeleteSQL(t1)
	del.Where(EQ(t1.Column("id").Name(), 10))
	log.Println(del.Statment())
	assert.Equal(t, "DELETE FROM `account` WHERE `id` = ?", del.Statment().SQLFormat())
}

func TestQuerySQL(t *testing.T) {
	t1 := TABLE("accounts").AS("a")
	t1.Columns("name", "mailbox", "age")
	t1.Column("name").AS("account_name")

	query := NewQuerySQL(t1)

	t2 := TABLE("account_posts").AS("p")
	t2.Columns("title", "create_at")

	query.Join(LEFT(t2, EQ(t1.Column("id").Name(), t2.Column("account_id").Name())))

	log.Println(query.Statment())

	assert.NotNil(t, query)

	t3 := query.Table("account_posts")
	assert.NotNil(t, t3)

	log.Println(t3.SQL())
	log.Println(t3.ColumnsSQL())
}
