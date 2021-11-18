package interactor

import (
	"database/sql"
	"github.com/ITK13201/holodule-bot/domain"
	"github.com/ITK13201/holodule-bot/interfaces/database"
)

type VideoInteractor struct {
	repository database.VideoRepository
}

func NewVideoInteractor(sqlHandler database.SqlHandler) *VideoInteractor {
	return &VideoInteractor{
		repository: database.VideoRepository{
			SqlHandler: sqlHandler,
		},
	}
}

func (interactor *VideoInteractor) Add(v domain.Video) (*domain.Video, error) {
	id, err := interactor.repository.Store(v)
	if err != nil {
		return nil, err
	}
	video, err := interactor.repository.FindById(id)
	return video, err
}

func (interactor *VideoInteractor) UpdateNotifiedAt(id int) error {
	err := interactor.repository.UpdateNotifiedAt(id)
	return err
}

func (interactor *VideoInteractor) GetById(id int) (*domain.Video, error) {
	video, err := interactor.repository.FindById(id)
	return video, err
}

func (interactor *VideoInteractor) GetBy3(distributorId int, url string, datetime sql.NullTime) (*domain.Video, error) {
	video, err := interactor.repository.FindBy3(distributorId, url, datetime)
	return video, err
}

func (interactor *VideoInteractor) GetAll() ([]domain.Video, error) {
	videos, err := interactor.repository.FindAll()
	return videos, err
}

func (interactor *VideoInteractor) GetComingSoon() ([]domain.Video, error) {
	videos, err := interactor.repository.FindComingSoon()
	return videos, err
}
