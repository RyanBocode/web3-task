package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
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

	// Step 1: Setup books table
	setupBooks(ctx, db)

	// Step 2: Run query to get books with price > 50
	runBookQuery(ctx, db)
}

func setupBooks(ctx context.Context, db *sqlx.DB) {
	createSQL := "CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, title TEXT, author TEXT, price REAL)"
	must(exec(ctx, db.DB, createSQL))

	db.ExecContext(ctx, "DELETE FROM books")
	insertSQL := "INSERT INTO books(id, title, author, price) VALUES (1, 'Go 语言实战', '张作者', 68.0), (2, '数据库系统概念', '韩作者', 88.0), (3, '算法图解', '李作者', 45.0)"
	db.ExecContext(ctx, insertSQL)
}

func runBookQuery(ctx context.Context, db *sqlx.DB) {
	// 查询价格大于50的书籍
	var books []Book
	must(db.Select(&books, `SELECT id, title, author, price FROM books WHERE price > ? ORDER BY price DESC`, 50))

	fmt.Println("=== 价格 > 50 的书籍 ===")
	for _, book := range books {
		fmt.Printf("ID: %d, 书名: %s, 作者: %s, 价格: %.2f\n", book.ID, book.Title, book.Author, book.Price)
	}

	// 查询所有书籍
	var allBooks []Book
	must(db.Select(&allBooks, `SELECT id, title, author, price FROM books ORDER BY id`))

	fmt.Println("\n=== 所有书籍信息 ===")
	for _, book := range allBooks {
		fmt.Printf("ID: %d, 书名: %s, 作者: %s, 价格: %.2f\n", book.ID, book.Title, book.Author, book.Price)
	}
}

func exec(ctx context.Context, db *sql.DB, query string) error {
	_, err := db.ExecContext(ctx, query)
	return err
}
