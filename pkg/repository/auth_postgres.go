package repository

import (
	"fmt"

	spaceTripsSystem "github.com/mrkozenko/space-trips-system"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user spaceTripsSystem.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (full_name, email, phone, password_hash, balance) values ($1,$2,$3,$4,0) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Email, user.Phone, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
