package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"people-api/internal/config"
	"people-api/internal/models/nationality"
	"people-api/internal/models/person"
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

func (s *Storage) GetNationalityIdByName(name string) (int64, error) {
	const op = "storage.postgres.GetNationalityIdByName"

	var id int64
	err := s.db.QueryRow("SELECT id FROM nationality WHERE name=$1", name).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *Storage) AddNationality(nat nationality.Nationality) error {
	const op = "storage.postgres.AddNationality"

	err := s.db.QueryRow("INSERT INTO nationalities VALUES ($1)", nat.Name)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) SavePerson(person person.Person) error {
	const op = "storage.postgres.SavePerson"

	stmt, err := s.db.Prepare("INSERT INTO persons(name, surname, patronymic, age, sex, nationality_id) VALUES (?,?,?,?,?,?)")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(
		person.Name,
		person.Surname,
		person.Patronymic,
		person.Age,
		person.Sex,
		person.Nationality.Id,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
func (s *Storage) SavePersons(persons []person.Person) error {
	const op = "storage.postgres.SavePersons"

	for _, p := range persons {
		stmt, err := s.db.Prepare("INSERT INTO persons(name, surname, patronymic, age, sex, nationality_id) VALUES (?,?,?,?,?,?)")
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		_, err = stmt.Exec(
			p.Name,
			p.Surname,
			p.Patronymic,
			p.Age,
			p.Sex,
			p.Nationality.Id,
		)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	return nil
}

func (s *Storage) DeletePersonById(id int64) error {
	const op = "storage.postgres.DeletePersonById"

	stmt, err := s.db.Prepare("DELETE FROM persons WHERE id = ?")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Storage) UpdatePerson(new person.Person, old person.Person) error {
	const op = "storage.postgres.UpdatePerson"

	stmt, err := s.db.Prepare("UPDATE persons SET name=?, surname=?, patronymic=?, sex=?, nationality_id=?, age=? WHERE id=?")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(
		new.Name,
		new.Surname,
		new.Patronymic,
		new.Sex,
		new.Nationality.Id,
		new.Age,
		old.Id,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
