package domain

import (
	"database/sql"
	"github.com/ITK13201/holodule-bot/config"
	"time"
)

type Video struct {
	Id          int          `json:"id" db:"id"`
	Distributor Distributor  `json:"distributor" db:"distributor"`
	Url         string       `json:"url" db:"url"`
	Datetime    time.Time    `json:"datetime" db:"datetime"`
	ImageUrl    string       `json:"image_url" db:"image_url"`
	NotifiedAt  sql.NullTime `json:"notified_at" db:"notified_at"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
}

func (v *Video) Utc2jst() *Video {
	v.Datetime = v.Datetime.In(config.JST)
	if v.NotifiedAt.Valid {
		notifiedAt := v.NotifiedAt.Time
		v.NotifiedAt.Time = notifiedAt.In(config.JST)
	}
	return v
}
