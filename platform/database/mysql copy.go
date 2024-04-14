package database

// import (
// 	"database/sql"
// 	"fmt"
// 	"os"

// 	_ "github.com/go-sql-driver/mysql" // load driver for Mysql
// 	"github.com/golden-ocean/ocean-admin/pkg/utils"
// 	"github.com/jmoiron/sqlx"
// 	"github.com/rs/zerolog"
// 	sqldblogger "github.com/simukti/sqldb-logger"
// 	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
// )

// func MySQLConnection() (*sqlx.DB, error) {
// 	mysqlConnURL, err := utils.ConnectionURLBuilder("mysql")
// 	if err != nil {
// 		return nil, err
// 	}
// 	db, err := sql.Open("mysql", mysqlConnURL) // db is *sql.DB
// 	loggerAdapter := zerologadapter.New(zerolog.New(os.Stdout))
// 	loggerOptions := []sqldblogger.Option{
// 		sqldblogger.WithSQLQueryFieldname("sql"),
// 		sqldblogger.WithWrapResult(false),
// 		sqldblogger.WithExecerLevel(sqldblogger.LevelDebug),
// 		sqldblogger.WithQueryerLevel(sqldblogger.LevelDebug),
// 		sqldblogger.WithPreparerLevel(sqldblogger.LevelDebug),
// 	}
// 	db = sqldblogger.OpenDriver(mysqlConnURL, db.Driver(), loggerAdapter, loggerOptions... /*, using_default_options*/) // db is STILL *sql.DB
// 	sqlx_db := sqlx.NewDb(db, "mysql")
// 	// db, err := sqlx.Connect("mysql", mysqlConnURL)
// 	if err != nil {
// 		return nil, fmt.Errorf("错误,无法连接数据库, %w", err)
// 	}

// 	// db.SetMaxOpenConns(25)           // 最大打开的连接数 不超过数据库服务自身支持的并发连接数
// 	// db.SetMaxIdleConns(5)            // 最大闲置的连接数 一般建议maxIdleConns的值为MaxOpenConns的1/2
// 	// db.SetConnMaxLifetime(time.Hour) // 连接的最大可复用时间 不超过数据库的超时参数值。

// 	return sqlx_db, nil
// }
