package repository

import (
	"github.com/ITK13201/holodule-bot/domain"
)

type VideoRepository interface {
	Store(distributor domain.Video) (int, error)
	FindById(id int) (*domain.Video, error)
	FindAll() ([]domain.Video, error)
	FindComingSoon() ([]domain.Video, error)
}
