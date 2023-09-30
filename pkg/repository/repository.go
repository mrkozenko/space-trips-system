package repository

import (
	"github.com/jmoiron/sqlx"
	spaceTripsSystem "github.com/mrkozenko/space-trips-system"
)

type Authorization interface {
	CreateUser(user spaceTripsSystem.User) (int, error)
}

type SpaceShips interface {
}

type SpaceTrips interface {
}

type UserTrips interface {
}

type Repository struct {
	Authorization
	SpaceShips
	SpaceTrips
	UserTrips
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
