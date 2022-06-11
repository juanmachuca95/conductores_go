package datastore

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLClient() *sql.DB {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("HOSTNAME"), os.Getenv("PORT"), os.Getenv("DATABASE"))
	db, err := sql.Open("mysql", dns)
	if err != nil {
		panic(err)
	}

	return db
}
