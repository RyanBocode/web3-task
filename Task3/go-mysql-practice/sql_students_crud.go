package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func p(format string, args ...any) {
	fmt.Printf(format+"\n", args...)
}

func main() {
	ctx := context.Background()

	// 使用SQLite内存数据库，不需要外部服务
	db, err := sql.Open("sqlite3", ":memory:")
	must(err)
	defer db.Close()
	must(db.Ping())

	// Step 1: Setup Students Table
	setupStudents(ctx, db)
	// Step 2: Perform CRUD
	doCRUD(ctx, db)
}

func setupStudents(ctx context.Context, db *sql.DB) {
	ddl := `
		CREATE TABLE IF NOT EXISTS students (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			age INTEGER,
			grade TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`
	must(execMulti(ctx, db, ddl))
}

func doCRUD(ctx context.Context, db *sql.DB) {
	// Insert record
	res, err := db.ExecContext(ctx, `INSERT INTO students(name, age, grade) VALUES(?, ?, ?)`, "张三", 20, "三年级")
	must(err)
	id, _ := res.LastInsertId()
	p("插入成功: id=%d", id)

	// Query students older than 18
	rows, err := db.QueryContext(ctx, `SELECT id, name, age, grade FROM students WHERE age > ?`, 18)
	must(err)
	defer rows.Close()
	p("查询(年龄>18)：")
	for rows.Next() {
		var id, age int
		var name, grade string
		must(rows.Scan(&id, &name, &age, &grade))
		p("  id=%d name=%s age=%d grade=%s", id, name, age, grade)
	}

	// Update record (张三's grade)
	_, err = db.ExecContext(ctx, `UPDATE students SET grade=? WHERE name=?`, "四年级", "张三")
	must(err)
	p("更新成功: 张三 -> 四年级")

	// Delete record where age < 15
	res, err = db.ExecContext(ctx, `DELETE FROM students WHERE age < ?`, 15)
	must(err)
	affected, _ := res.RowsAffected()
	p("删除成功: 受影响行=%d", affected)
}

func execMulti(ctx context.Context, db *sql.DB, sqls string) error {
	for _, s := range strings.Split(sqls, ";") {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		if _, err := db.ExecContext(ctx, s); err != nil {
			return err
		}
	}
	return nil
}
