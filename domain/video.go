package domain

import (
	"database/sql"
	"time"
)

type Video struct {
	Id          int          `json:"id" db:"id"`
	Distributor Distributor  `json:"distributor" db:"distributor"`
	Url         string       `json:"url" db:"url"`
	Datetime    sql.NullTime `json:"datetime" db:"datetime"`
	ImageUrl    string       `json:"image_url" db:"image_url"`
	NotifiedAt  sql.NullTime `json:"notified_at" db:"notified_at"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
}
