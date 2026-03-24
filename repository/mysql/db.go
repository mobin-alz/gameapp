package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLDB struct {
	db *sql.DB
}

func New() *MySQLDB {
	db, err := sql.Open("mysql", "gameapp:gameappPassword@(localhost:3306)/gameapp_db")
	if err != nil {
		panic(fmt.Errorf("can't open mysql db: %v\n", err))
	}
	db.SetConnMaxLifetime(time.Minute * 60)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	return &MySQLDB{db}
}
