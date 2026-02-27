package model

import "time"

type UserTag struct {
	ID        int64     `json:"id"`
	UserID    string    `json:"user_id"`
	Name      []string  `json:"name"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}
