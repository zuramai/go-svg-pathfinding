package model

import (
	"github.com/jackc/pgtype"
)

type User struct {
	Id        int64            `json:"id"`
	Username  string           `json:"username"`
	Password  string           `json:"password"`
	Role      string           `json:"role"`
	CreatedAt pgtype.Timestamp `json:"-"`
	UpdatedAt pgtype.Timestamp `json:"-"`
}
