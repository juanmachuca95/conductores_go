package internal

import (
	"database/sql"
	"fmt"
	"os"
)

type MySQLClient struct {
	*sql.DB
}

func NewMySQLClient() *MySQLClient {
	dns := fmt.Sprintf("%s:%stcp(%s:%s)/%s", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("HOSTNAME"), os.Getenv("PORT"), os.Getenv("DATABASE"))
	db, err := sql.Open("mysql", dns)
	if err != nil {
		panic("No fue posible conectar con la base de datos")
	}

	return &MySQLClient{db}
}
