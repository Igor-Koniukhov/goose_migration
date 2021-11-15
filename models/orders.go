package models

import "time"

type Orders struct {
	ID          int       `json:"id"`
	OrderedTime time.Time `json:"ordered_time"`
	Status      string    `json:"status"`
	UserID      int       `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
