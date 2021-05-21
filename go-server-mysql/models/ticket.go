package models

import "time"

type Ticket struct {
	ID        int64 `json:"id" gorm:"primary_key;auto_increment"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Usuario   string `json:"usuario"`
	Status    bool   `json:"status"`
}
