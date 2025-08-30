package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()

	// 使用SQLite内存数据库，不需要外部服务
	sqldb, err := sql.Open("sqlite3", ":memory:")
	must(err)
	defer sqldb.Close()

	must(sqldb.Ping())

	db := sqlx.NewDb(sqldb, "sqlite3")

	// Step 1: Setup employees table
	setupEmployees(ctx, db)

	// Step 2: Run queries
	runEmployeeQueries(ctx, db)
}

func setupEmployees(ctx context.Context, db *sqlx.DB) {
	createSQL := "CREATE TABLE IF NOT EXISTS employees (id INTEGER PRIMARY KEY, name TEXT, department TEXT, salary REAL)"
	must(exec(ctx, db.DB, createSQL))

	db.ExecContext(ctx, "DELETE FROM employees")
	insertSQL := "INSERT INTO employees(id, name, department, salary) VALUES (1, 'Alice', '技术部', 25000), (2, 'Bob', '市场部', 18000), (3, 'Carol', '技术部', 32000)"
	db.ExecContext(ctx, insertSQL)
}

func runEmployeeQueries(ctx context.Context, db *sqlx.DB) {
	// 查询技术部员工
	var techs []Employee
	must(db.Select(&techs, `SELECT id, name, department, salary FROM employees WHERE department=?`, "技术部"))

	fmt.Println("=== 技术部员工 ===")
	for _, emp := range techs {
		fmt.Printf("ID: %d, 姓名: %s, 部门: %s, 工资: %.2f\n", emp.ID, emp.Name, emp.Department, emp.Salary)
	}

	// 查询工资最高的员工
	var top Employee
	must(db.Get(&top, `SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1`))

	fmt.Println("\n=== 工资最高的员工 ===")
	fmt.Printf("ID: %d, 姓名: %s, 部门: %s, 工资: %.2f\n", top.ID, top.Name, top.Department, top.Salary)

	// 查询所有员工
	var allEmployees []Employee
	must(db.Select(&allEmployees, `SELECT id, name, department, salary FROM employees ORDER BY id`))

	fmt.Println("\n=== 所有员工信息 ===")
	for _, emp := range allEmployees {
		fmt.Printf("ID: %d, 姓名: %s, 部门: %s, 工资: %.2f\n", emp.ID, emp.Name, emp.Department, emp.Salary)
	}
}

func exec(ctx context.Context, db *sql.DB, query string) error {
	_, err := db.ExecContext(ctx, query)
	return err
}
