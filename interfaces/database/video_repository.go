package database

import (
	"github.com/ITK13201/holodule-bot/domain"
)

type VideoRepository struct {
	SqlHandler
}

func (repo *VideoRepository) Store(v domain.Video) (int, error) {
	result, err := repo.Exec(`
		INSERT INTO videos (
			distributor_id,
			url,
			datetime,
			image_url
		) VALUES (
			?,
			?,
			?,
			?
		)
		`,
		v.Distributor.Id,
		v.Url,
		v.Datetime,
		v.ImageUrl,
	)
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

func (repo *VideoRepository) FindById(id int) (*domain.Video, error) {
	var v domain.Video
	err := repo.Get(
		&v,
		`SELECT
			v.id AS id,
			url,
			datetime,
			image_url,
			notified_at,
			v.created_at AS created_at,
			v.updated_at AS updated_at,
			d.id AS 'distributor.id',
			name AS 'distributor.name',
			icon_url AS 'distributor.icon_url',
			d.created_at AS 'distributor.created_at',
			d.updated_at AS 'distributor.updated_at'
		FROM videos AS v
		INNER JOIN distributors AS d ON v.distributor_id = d.id
		WHERE v.id = ?
		`,
		id,
	)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

func (repo *VideoRepository) FindAll() ([]domain.Video, error) {
	var videos []domain.Video
	err := repo.Select(
		&videos,
		`SELECT
			v.id AS id,
			url,
			datetime,
			image_url,
			notified_at,
			v.created_at AS created_at,
			v.updated_at AS updated_at,
			d.id AS 'distributor.id',
			name AS 'distributor.name',
			icon_url AS 'distributor.icon_url',
			d.created_at AS 'distributor.created_at',
			d.updated_at AS 'distributor.updated_at'
		FROM videos AS v
		INNER JOIN distributors AS d ON v.distributor_id = d.id
	`,
	)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (repo *VideoRepository) FindComingSoon() ([]domain.Video, error) {
	var videos []domain.Video
	err := repo.Select(
		&videos,
		`SELECT
			v.id AS id,
			url,
			datetime,
			image_url,
			notified_at,
			v.created_at AS created_at,
			v.updated_at AS updated_at,
			d.id AS 'distributor.id',
			name AS 'distributor.name',
			icon_url AS 'distributor.icon_url',
			d.created_at AS 'distributor.created_at',
			d.updated_at AS 'distributor.updated_at'
		FROM videos AS v
		INNER JOIN distributors AS d ON v.distributor_id = d.id
		WHERE notified_at IS NULL
			AND (NOW() - INTERVAL 2 HOUR) < datetime AND datetime < (NOW() - INTERVAL 1 HOUR)
	`,
	)
	if err != nil {
		return nil, err
	}
	return videos, nil
}
