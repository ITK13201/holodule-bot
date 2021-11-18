package repository

import "github.com/ITK13201/holodule-bot/domain"

type DistributorRepository interface {
	Store(distributor domain.Distributor) (int, error)
	Update(d domain.Distributor) error
	FindById(id int) (*domain.Distributor, error)
	FindByName(name string) (*domain.Distributor, error)
	FindAll() ([]domain.Distributor, error)
}
