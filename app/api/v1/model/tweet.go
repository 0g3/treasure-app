package model

import "time"

type CreateTweetRequest struct {
	Body string `json:"body"`
}

type Tweet struct {
	ID        uint64    `json:"id"`
	UserID    string    `json:"user_id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}
