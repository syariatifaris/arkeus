package entity

import (
	"database/sql"

	"github.com/lib/pq"
)

const (
	ColCreateTime = "create_time"
	ColUpdateTime = "update_time"
	ColCreateBy   = "create_by"
	ColUpdateBy   = "update_by"
)

//Base Entity
type BaseEntity struct {
	CreateTime pq.NullTime   `json:"-" db:"create_time"`
	UpdateTime pq.NullTime   `json:"-" db:"update_time"`
	CreateBy   sql.NullInt64 `json:"-" db:"create_by"`
	UpdateBy   sql.NullInt64 `json:"-" db:"update_by"`
}

//Base Authentication Request
type BaseAuthRequest struct {
}

//Base Response
type BaseResponse struct {
	ErrMessage string `json:"err_message"`
	Status     string `json:"status"`
}
