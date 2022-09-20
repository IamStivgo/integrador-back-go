package store

import (
	"database/sql"

	"github.com/iamstivgo/integrador-back-go.git/internal/domain"
)

type DentistSqlStore struct {
	db *sql.DB
}

func NewDentistSqlStore(db *sql.DB) DentistSqlStore {
	return &DentistSqlStore{db: db}
}

func (s *DentistSqlStore) GetAll() ([]domain.Dentist, error) {
}

func (s *DentistSqlStore) GetByID(id int) (domain.Dentist, error) {
	var dentist domain.Dentist
	query := `SELECT id, first_name, last_name, enrollment FROM dentist WHERE id = ?`
	row := s.db.QueryRow(query, id)
	err := row.Scan(&dentist.ID, &dentist.FirstName, &dentist.LastName, &dentist.Enrollment)
	if err != nil {
		return domain.Dentist{} , err
	}
	return dentist, nil
}

func (s *DentistSqlStore) Create(dentist domain.Dentist) error {
	query := `INSERT INTO dentist (first_name, last_name, enrollment) VALUES (?, ?, ?)`
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(dentist.FirstName, dentist.LastName, dentist.Enrollment)
	if err != nil {
		return err
	}
	_,err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *DentistSqlStore) Update(dentist domain.Dentist) error {
	query := `UPDATE dentist SET first_name = ?, last_name = ?, enrollment = ? WHERE id = ?`
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(dentist.FirstName, dentist.LastName, dentist.Enrollment, dentist.ID)
	if err != nil {
		return err
	}
	_,err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *DentistSqlStore) Delete(id int) error {
	query := `DELETE FROM dentist WHERE id = ?`
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	_,err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
