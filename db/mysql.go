package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//MySQLConnect function for connect to mysql
func MySQLConnect() (dbs *sql.DB) {
	dbs, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/serviciotest")

	if err != nil {
		panic(err.Error())
	}

	return dbs
}
