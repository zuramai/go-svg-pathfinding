package model

import "time"

type LoginToken struct {
	id        int64
	user_id   int64
	token     string
	createdAt time.Time
	updatedAt time.Time
}
