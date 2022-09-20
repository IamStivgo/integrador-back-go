package dentist

import (
	"github.com/iamstivgo/integrador-back-go.git/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Dentist, error)
	GetByID(id int) (domain.Dentist, error)
	Create(dentist domain.Dentist) (domain.Dentist, error)
	Update(dentist domain.Dentist) (domain.Dentist, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAll() ([]*domain.Dentist, error) {
}

func (s *service) GetByID(id int) (*domain.Dentist, error) {
	d, err := s.r.GetByID(id)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (s *service) Create(dentist domain.Dentist) (domain.Dentist, error) {
	d, err := s.r.Create(dentist)
	if err != nil {
		return domain.Dentist{}, err
	}
	return d, nil
}
