package model

import "time"

type Place struct {
	Id          int64     `json:"id"`
	Code        string    `json:"-"`
	Name        string    `json:"name"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	X           int       `json:"x"`
	Y           int       `json:"y"`
	ImagePath   string    `json:"image_path"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
