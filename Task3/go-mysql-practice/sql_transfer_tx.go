package main

import (
	"context"
	"database/sql"
	"fmt"

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

	// Setup tables
	setupAcc(ctx, db)
	// Perform transactions
	transfer(ctx, db, 1, 2, 100)    // Success
	transfer(ctx, db, 1, 2, 999999) // Failure (balance check)
}

func setupAcc(ctx context.Context, db *sql.DB) {
	must(exec(ctx, db, `CREATE TABLE IF NOT EXISTS accounts (
		id INTEGER PRIMARY KEY, balance INTEGER NOT NULL
	)`))

	must(exec(ctx, db, `CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		from_account_id INTEGER, to_account_id INTEGER, amount INTEGER
	)`))

	db.ExecContext(ctx, `DELETE FROM accounts`)
	db.ExecContext(ctx, `DELETE FROM transactions`)
	must(exec(ctx, db, `INSERT INTO accounts(id, balance) VALUES(1, 500), (2, 200)`))
}

func transfer(ctx context.Context, db *sql.DB, fromID, toID, amount int64) {
	tx, err := db.BeginTx(ctx, nil)
	must(err)
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var fromBal int64
	err = tx.QueryRowContext(ctx, `SELECT balance FROM accounts WHERE id=?`, fromID).Scan(&fromBal)
	if err != nil {
		p("查询来源失败: %v", err)
		return
	}
	if fromBal < amount {
		p("余额不足，回滚")
		tx.Rollback()
		return
	}

	// Update balances
	_, err = tx.ExecContext(ctx, `UPDATE accounts SET balance=balance-? WHERE id=?`, amount, fromID)
	must(err)
	_, err = tx.ExecContext(ctx, `UPDATE accounts SET balance=balance+? WHERE id=?`, amount, toID)
	must(err)
	_, err = tx.ExecContext(ctx, `INSERT INTO transactions(from_account_id, to_account_id, amount) VALUES(?, ?, ?)`, fromID, toID, amount)
	must(err)

	err = tx.Commit()
	if err != nil {
		p("提交失败: %v", err)
		return
	}

	// Print balance after transaction
	var fromFinal, toFinal int64
	db.QueryRowContext(ctx, `SELECT balance FROM accounts WHERE id=?`, fromID).Scan(&fromFinal)
	db.QueryRowContext(ctx, `SELECT balance FROM accounts WHERE id=?`, toID).Scan(&toFinal)
	p("转账成功: A=%d, B=%d", fromFinal, toFinal)
}

func exec(ctx context.Context, db *sql.DB, query string) error {
	_, err := db.ExecContext(ctx, query)
	return err
}
