package database

import (
	"log"

	"github.com/jmoiron/sqlx"
)

// OpenDBConnection func for opening database connection.
func OpenDBConnection() *sqlx.DB {
	db, err := MySQLConnection()
	if err != nil {
		log.Fatal("Mysql数据库连接出错!", err)
	}
	return db
}
