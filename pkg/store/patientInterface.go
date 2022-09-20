package store

import "github.com/iamstivgo/integrador-back-go.git/internal/domain"

type PatientInterface interface {
	GetAll() ([]domain.Patient, error)
	GetByID(id int) (domain.Patient, error)
	Create(patient domain.Patient) error
	Update(patient domain.Patient) error
	Delete(id int) error
}