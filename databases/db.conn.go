package databases

import (
	"database/sql"

	_"github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func DBInit()  {
	conf := GetConfig()

	// data source name
	connectionSTR := conf.DB_USER + ":" + conf.DB_PASS + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME

	db, err = sql.Open("mysql", connectionSTR)
	if err != nil {
		panic("connectionSTR error")
	}

	err = db.Ping()
	if err != nil {
		panic("Invalid DSN")
	}
}

func CreateConn() *sql.DB {
	return db
}