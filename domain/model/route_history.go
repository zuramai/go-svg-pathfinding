package model

import "time"

type RouteHistory struct {
	id          int64
	fromPlaceID int64
	toPlaceID   int64
	userID      int64
	routes      string
	createdAt   time.Time
	updatedAt   time.Time
}
