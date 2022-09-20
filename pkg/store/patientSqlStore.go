package store

import (
	"database/sql"

	"github.com/iamstivgo/integrador-back-go.git/internal/domain"
)

type PatientSqlStore struct {
	db *sql.DB
}

func NewPatientSqlStore(db *sql.DB) PatientSqlStore {
	return &PatientSqlStore{db: db}
}

func (s *PatientSqlStore) GetAll() ([]*domain.Patient, error) {
}

func (s *PatientSqlStore) GetByID(id int) (domain.Patient, error) {
	var patient domain.Patient
	query := `SELECT id, first_name, last_name, dni, address, discharge_date FROM patient WHERE id = ?`
	row := s.db.QueryRow(query, id)
	err := row.Scan(&patient.ID, &patient.FirstName, &patient.LastName, &patient.DNI, &patient.Address, &patient.Discharge_date)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (s *PatientSqlStore) Create(patient domain.Patient) error {
	query := `INSERT INTO patient (first_name, last_name, dni, address, discharge_date) VALUES (?, ?, ?, ?, ?)`
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(patient.FirstName, patient.LastName, patient.DNI, patient.Address, patient.Discharge_date)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *PatientSqlStore) Update(patient domain.Patient) error {
	query := `UPDATE patient SET first_name = ?, last_name = ?, dni = ?, address = ?, discharge_date = ? WHERE id = ?`
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(patient.FirstName, patient.LastName, patient.DNI, patient.Address, patient.Discharge_date, patient.ID)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *PatientSqlStore) Delete(id int) error {
	query := `DELETE FROM patient WHERE id = ?`
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
