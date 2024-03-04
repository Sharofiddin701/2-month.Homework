package storage

import (
	"database/sql"
	"fmt"
	"rent-car/config"
)

type Store struct {
	DB       *sql.DB
	Customer customerRepo
}

func New(cfg config.Config) (Store, error) {
	url := fmt.Sprintf(`host=%s port=%v user=%s password=%s database=%s sslmode=disable`,
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return Store{}, err
	}

	NewCustomer := NewCustomer(db)
	return Store{
		DB:       db,
		Customer: NewCustomer,
	}, nil

}
