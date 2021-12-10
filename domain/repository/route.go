package repository

import "go.mongodb.org/mongo-driver/mongo"

type RouteRepository struct {
	DB *mongo.Database
}
