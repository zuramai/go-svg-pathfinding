package model

import "time"

type Place struct {
	id          int64
	name        string
	latitude    float64
	longitude   float64
	x           int
	y           int
	image_path  string
	description string
	createdAt   time.Time
	updatedAt   time.Time
}
