package model

import "time"

type LoginToken struct {
	Id        int64
	UserId    int64
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
