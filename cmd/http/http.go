package http

import (
	"context"
	"fmt"
	"os"
	"wsc2017/app"

	"github.com/jackc/pgx/v4"
)

type HttpServer struct {
	app app.Application
	db  *pgx.Conn
}

func (http *HttpServer) Connect() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
}
