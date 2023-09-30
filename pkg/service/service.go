package service

import (
	spaceTripsSystem "github.com/mrkozenko/space-trips-system"
	"github.com/mrkozenko/space-trips-system/pkg/repository"
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

type Service struct {
	Authorization
	SpaceShips
	SpaceTrips
	UserTrips
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
