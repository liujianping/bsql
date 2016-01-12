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

log.Println("insert format:", insert.SQLFormat())
log.Println("insert params:", insert.SQLParams())

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

````

- 删除语句(DELETE)

````
//! delete sql
t3 := TABLE("accounts")

update := NewDeleteSQL(t3).Where(EQ(t.Column("id").Name(), 10)).Statment()

log.Println("delete format:", insert.SQLFormat())
log.Println("delete params:", insert.SQLParams())

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

````