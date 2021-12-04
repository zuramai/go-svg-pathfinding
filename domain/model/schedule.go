package model

import "time"

type Schedule struct {
	id            int64
	Type          string
	line          int
	fromPlaceId   int64
	toPlaceId     int64
	departureTime time.Time
}
