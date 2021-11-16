package domain

import "time"

type Video struct {
	Id          int         `json:"id" db:"id"`
	Distributor Distributor `json:"distributor" db:"distributor"`
	Url         string      `json:"url" db:"url"`
	Datetime    time.Time   `json:"datetime" db:"datetime"`
	ImageUrl    string      `json:"image_url" db:"image_url"`
	CreatedAt   time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at" db:"updated_at"`
}
