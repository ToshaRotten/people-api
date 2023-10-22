package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"people-api/internal/config"
	"people-api/internal/models/nationality"
)

type Storage struct {
	db *sql.DB
}

func New(config config.Storage) (*Storage, error) {
	// op is used for mark errors
	const op = "storage.postgres.New"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.Username, config.Password, config.Database)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) CheckNationality(nationality nationality.Nationality) (bool, error) {
	const op = "storage.postgres.CheckNationality"

	var hasNat bool
	err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM nationality WHERE name=$1 AND nationality IS NOT NULL)", nationality.Name).Scan(&hasNat)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return hasNat, nil
}

func (s *Storage) GetNationalityIdByName(name string) (int, error) {
	const op = "storage.postgres.GetNationalityIdByName"

	var id int
	err := s.db.QueryRow("SELECT id FROM nationality WHERE name=$1", name).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

//func (s *Storage) AddNationality(name string) error {
//
//}

//func (s *Storage) SaveUser(user user.User) error {
//	const op = "storage.postgres.SaveUser"
//
//	stmt, err := s.db.Prepare("INSERT INTO users() VALUES (&)")
//
//	return nil
//}
