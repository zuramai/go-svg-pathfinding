package model

import "time"

type User struct {
	id        int64
	username  string
	password  string
	role      string
	createdAt time.Time
	updatedAt time.Time
}
