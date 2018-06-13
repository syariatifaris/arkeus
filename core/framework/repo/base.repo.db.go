package repo

import (
	"github.com/jmoiron/sqlx"
)

//Base Database Repository
type BaseDatabaseRepository interface {
	GetDB() *sqlx.DB
}
