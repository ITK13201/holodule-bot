package interactor

import (
	"github.com/ITK13201/holodule-bot/domain"
	"github.com/ITK13201/holodule-bot/interfaces/database"
)

type DistributorInteractor struct {
	repository database.DistributorRepository
}

func NewDistributorInteractor(sqlHandler database.SqlHandler) *DistributorInteractor {
	return &DistributorInteractor{
		repository: database.DistributorRepository{
			SqlHandler: sqlHandler,
		},
	}
}

func (interactor *DistributorInteractor) Add(d domain.Distributor) (*domain.Distributor, error) {
	id, err := interactor.repository.Store(d)
	if err != nil {
		return nil, err
	}
	distributor, err := interactor.repository.FindById(id)
	return distributor, err
}

func (interactor *DistributorInteractor) GetById(id int) (*domain.Distributor, error) {
	d, err := interactor.repository.FindById(id)
	return d, err
}

func (interactor *DistributorInteractor) GetByName(name string) (*domain.Distributor, error) {
	d, err := interactor.repository.FindByName(name)
	return d, err
}

func (interactor *DistributorInteractor) GetAll() ([]domain.Distributor, error) {
	distributors, err := interactor.repository.FindAll()
	return distributors, err
}
