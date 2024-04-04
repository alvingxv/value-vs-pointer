package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	user     = "root"
	pass     = "root"
	port     = "3306"
	host     = "localhost"
	database = "playground"
)

var (
	db  *sql.DB
	err error
)

func handleDatabaseConnection() {
	dbinfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, database)

	db, err := sql.Open("mysql", dbinfo)
	if err != nil {
		log.Panic(err)
	}

	if err != nil {
		log.Panic("error occured while trying to validate database arguments:", err)
	}

	err = db.Ping()

	if err != nil {
		log.Panic("error occured while trying to connect to database:", err)
	}

}

func InitiliazeDatabase() {
	handleDatabaseConnection()
}

func GetDatabaseInstance() *sql.DB {
	return db
}
