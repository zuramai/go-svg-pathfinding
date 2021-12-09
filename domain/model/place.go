package model

import "time"

type Place struct {
	Id          int64
	Code        int64
	Name        string
	Latitude    float64
	Longitude   float64
	X           int
	Y           int
	ImagePath   string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
