package mysql

import (
    "database/sql"
    "log"
	"net/url"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
    var err error

    dbHost := "localhost"
	dbPort := "3306"
	dbUser := "root"
	dbPass := "biteme10"
	dbName := "ibadf"
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	db, err = sql.Open(`mysql`, dsn)
    if err != nil {
        log.Panic(err)
    }

    if err = db.Ping(); err != nil {
        log.Panic(err)
    }
}