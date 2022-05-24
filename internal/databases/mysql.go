package internal

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLClient struct {
	*sql.DB
}

func NewMySQLClient() *MySQLClient {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("HOSTNAME"), os.Getenv("PORT"), os.Getenv("DATABASE"))
	db, err := sql.Open("mysql", dns)
	if err != nil {
		panic(err)
	}

	return &MySQLClient{db}
}
