package store

import "github.com/iamstivgo/integrador-back-go.git/internal/domain"

type DentistInterface interface {
	GetAll() ([]domain.Dentist, error)
	GetByID(id int) (domain.Dentist, error)
	Create(dentist domain.Dentist) error
	Update(dentist domain.Dentist) error
	Delete(id int) error
}
