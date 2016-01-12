bsql (Build SQL)
===

Build SQL package

快速构建SQL语句, 方便程序自动构建。

### 安装

````
$: go get github.com/liujianping/bsql

````

### 快速开始

- 新增语句(INSERT)

````
import . "github.com/liujianping/bsql"

//! insert sql
t1 := TABLE("accounts")

t1.Column("mailbox").Set("demo@demo.cn")
t1.Column("age").Set(20)
t1.Column("create_at").Set(time.Now())

insert := NewInsertSQL(t1).Statment()
assert.NotNil(t, insert)

log.Println("insert format:", insert.SQLFormat())
log.Println("insert params:", insert.SQLParams())

//! output
//! insert format: INSERT INTO `accounts` ( `mailbox`, `age`, `create_at` ) VALUES ( ?, ?, ? )
//! insert params: [demo@demo.cn 20 2016-01-12 21:59:56.607552428 +0800 CST]

````

- 更新语句(UPDATE)

````
//! update sql
t2 := TABLE("accounts")

t2.Column("mailbox").Set("demo@demo.cn")
t2.Column("age").Set(30)

update := NewUpdateSQL(t2).Where(EQ(t.Column("id").Name(), 10)).Statment()

log.Println("update format:", insert.SQLFormat())
log.Println("update params:", insert.SQLParams())

//! output
//! update format: UPDATE `accounts` SET `mailbox` = ?, `age` = ? WHERE `id` = ?
//! update params: [demo@demo.cn 30 10]

````

- 删除语句(DELETE)

````
//! delete sql
t3 := TABLE("accounts")

update := NewDeleteSQL(t3).Where(EQ(t.Column("id").Name(), 10)).Statment()

log.Println("delete format:", insert.SQLFormat())
log.Println("delete params:", insert.SQLParams())

//! output
//! delete format: DELETE FROM `accounts` WHERE `id` = ?
//! delete params: [10]

````

- 查询语句(SELECT)

````

//! query sql
t4 := TABLE("accounts").AS("a")
t4.Columns("name", "mailbox", "age")

t5 := TABLE("account_comments").AS("c")
t5.Columns("title", "content")

query := bsql.NewQuerySQL(t4).Join(LEFT(t5, EQ(t4.Column("id").Name(), t5.Column("account_id").Name())))
query.Where(EQ(t4.Column("id").Name(), 10))

log.Println("query format:", query.Statment().SQLFormat())
log.Println("query params:", query.Statment().SQLParams())

//! output
//! query format: SELECT a.name, a.mailbox, a.age, a.id , c.title, c.content, c.account_id FROM `accounts` AS a LEFT JOIN `account_comments` AS c ON a.id = c.account_id WHERE a.id = ?
//! delete params: [10]

````

### TODO
