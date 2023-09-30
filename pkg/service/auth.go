package service

import (
	"crypto/sha1"
	"fmt"

	spaceTripsSystem "github.com/mrkozenko/space-trips-system"
	"github.com/mrkozenko/space-trips-system/pkg/repository"
)

const salt = "fng43j7tgq4hewfidc"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}
func (s *AuthService) CreateUser(user spaceTripsSystem.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
