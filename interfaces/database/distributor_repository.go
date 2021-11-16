package database

import (
	"github.com/ITK13201/holodule-bot/domain"
)

type DistributorRepository struct {
	SqlHandler
}

func (repo *DistributorRepository) Store(d domain.Distributor) (int, error) {
	result, err := repo.NamedExec(`
		INSERT INTO distributors (
			name,
			icon_url
		) VALUES (
			:name,
			:icon_url
		)
	`, d)
	if err != nil {
		return 0, err
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	id := int(id64)
	return id, err
}

func (repo *DistributorRepository) FindById(id int) (*domain.Distributor, error) {
	var d domain.Distributor
	err := repo.Get(
		&d,
		"SELECT * FROM distributors WHERE id = ?",
		id,
	)
	if err != nil {
		return nil, err
	}

	return &d, nil
}

func (repo *DistributorRepository) FindByName(name string) (*domain.Distributor, error) {
	var d domain.Distributor
	err := repo.Get(
		&d,
		"SELECT * FROM distributors WHERE name = ?",
		name,
	)
	if err != nil {
		return nil, err
	}

	return &d, nil
}

func (repo *DistributorRepository) FindAll() ([]domain.Distributor, error) {
	var distributors []domain.Distributor
	err := repo.Select(
		&distributors,
		"SELECT * FROM users",
	)
	if err != nil {
		return nil, err
	}
	return distributors, nil
}
