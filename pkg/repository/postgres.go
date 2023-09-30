package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable       = "users"
	clientTripsTable = "client_trips"
	spaceTripsTable  = "space_trips"
	spaceshipsTable  = "spaceships"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBNAME   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBNAME, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
