package repo

import (
	"github.com/tokopedia/sqlt"
)

//Base Database Repository
type BaseDatabaseRepository interface {
	GetDB() *sqlt.DB
}
