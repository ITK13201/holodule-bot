package domain

import "time"

type Distributor struct {
	Id        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	IconUrl   string    `json:"icon_url" db:"icon_url"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
