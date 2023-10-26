package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"people-api/internal/config"
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

func (s *Storage) SavePerson(p person.Person) error {
	const op = "storage.postgres.SavePerson"

	stmt, err := s.db.Prepare("INSERT INTO persons(name, surname, patronymic, age, sex, nationality) VALUES ($1,$2,$3,$4,$5,$6)")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(
		p.Name,
		p.Surname,
		p.Patronymic,
		p.Age,
		p.Sex,
		p.Nationality,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
func (s *Storage) SavePersons(persons []person.Person) error {
	const op = "storage.postgres.SavePersons"

	for _, p := range persons {
		stmt, err := s.db.Prepare("INSERT INTO persons(name, surname, patronymic, age, sex, nationality) VALUES ($1,$2,$3,$4,$5,$6)")
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		_, err = stmt.Exec(
			p.Name,
			p.Surname,
			p.Patronymic,
			p.Age,
			p.Sex,
			p.Nationality,
		)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	return nil
}

func (s *Storage) DeletePerson(id int64) error {
	const op = "storage.postgres.DeletePersonById"

	stmt, err := s.db.Prepare("DELETE FROM persons WHERE id = $1")
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

	stmt, err := s.db.Prepare("UPDATE persons SET name=$1, surname=$2, patronymic=$3, sex=$4, nationality=$5, age=$6 WHERE id=$7")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(
		new.Name,
		new.Surname,
		new.Patronymic,
		new.Sex,
		new.Nationality,
		new.Age,
		old.Id,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Storage) GetPerson(id int) (*person.Person, error) {
	const op = "storage.postgres.GetPerson"

	row := s.db.QueryRow("SELECT name, surname, patronymic, age, sex, nationality FROM persons WHERE id = $1", id)
	var p person.Person
	if err := row.Scan(&p.Name, &p.Surname, &p.Patronymic, &p.Age, &p.Sex, &p.Nationality); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &p, nil
}

func (s *Storage) GetCountOfPersons() (int, error) {
	const op = "storage.postgres.GetCountOfPersons"
	var count int
	if err := s.db.QueryRow("SELECT count(*) FROM persons").Scan(&count); err != nil {
		return count, fmt.Errorf("%s: %w", op, err)
	}
	return count, nil
}

func (s *Storage) GetPersons(limit int, offset int) ([]person.Person, error) {
	const op = "storage.postgres.GetPersons"

	rows, err := s.db.Query("SELECT * FROM persons LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return []person.Person{}, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var persons []person.Person
	for rows.Next() {
		var p person.Person
		if err := rows.Scan(&p.Id, &p.Name, &p.Surname, &p.Patronymic, &p.Age, &p.Sex, &p.Nationality); err != nil {
			return []person.Person{}, fmt.Errorf("%s: %w", op, err)
		}
		persons = append(persons, p)
	}

	return persons, nil
}
