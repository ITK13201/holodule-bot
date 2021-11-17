package repository

import (
	"database/sql"
	"github.com/ITK13201/holodule-bot/domain"
)

type VideoRepository interface {
	Store(distributor domain.Video) (int, error)
	UpdateNotifiedAt(id int) error
	FindById(id int) (*domain.Video, error)
	FindBy3(distributorId int, url string, datetime sql.NullTime) (*domain.Video, error)
	FindAll() ([]domain.Video, error)
	FindComingSoon() ([]domain.Video, error)
}
