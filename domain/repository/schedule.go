package repository

import "github.com/jackc/pgx/v4"

type ScheduleRepository struct {
	DB *pgx.Conn
}
