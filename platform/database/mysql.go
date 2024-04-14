package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	// _ "github.com/go-sql-driver/mysql" // load driver for Mysql
	"github.com/golden-ocean/ocean-admin/pkg/utils"
	"github.com/jmoiron/sqlx"
	"github.com/qustavo/sqlhooks/v2"
)

type Hooks struct{}
type contextKey string

const beginKey contextKey = "begin"

// Before hook will print the query with it's args and return the context with the timestamp
func (h *Hooks) Before(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	// fmt.Printf("> %s %q", query, args)
	fmt.Printf("before> q=%s, args = %+v\n", query, args)
	return context.WithValue(ctx, beginKey, time.Now()), nil
}

// After hook will get the timestamp registered on the Before hook and print the elapsed time
func (h *Hooks) After(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	begin := ctx.Value(beginKey).(time.Time)
	fmt.Printf(". took: %s\n", time.Since(begin))
	return ctx, nil
}

func MySQLConnection() (*sqlx.DB, error) {
	dsn, err := utils.ConnectionURLBuilder("mysql")
	if err != nil {
		return nil, err
	}
	sql.Register("mysqlWithHooks", sqlhooks.Wrap(&mysql.MySQLDriver{}, &Hooks{}))
	db, err := sql.Open("mysqlWithHooks", dsn) // db is *sql.DB
	sqlx_db := sqlx.NewDb(db, "mysql")
	// db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("错误,无法连接数据库, %w", err)
	}
	// db.SetMaxOpenConns(25)   // 最大打开的连接数 不超过数据库服务自身支持的并发连接数
	// db.SetMaxIdleConns(5)    // 最大闲置的连接数 一般建议maxIdleConns的值为MaxOpenConns的1/2
	// db.SetConnMaxLifetime(2) // 连接的最大可复用时间 不超过数据库的超时参数值。
	// 判断连接的数据库
	err = db.Ping()
	if err != nil {
		defer db.Close()
		fmt.Printf("连接 %s 失败, Error: %v \n", dsn, err)
	}

	return sqlx_db, nil
}
