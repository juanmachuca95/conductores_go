package repository

import (
	"database/sql"

	"github.com/juanmachuca95/conductores_go/usescases/repository"
)

type dbRepository struct {
	db *sql.DB
}

func NewDBRepository(db *sql.DB) repository.DBRepository {
	return dbRepository{db: db}
}
