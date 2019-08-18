package model

import "time"

type User struct {
	ID          string     `json:"id"`
	DisplayID   string     `json:"display_id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type BatchGetUserResponse struct {
	User   *User    `json:"user"`
	Tweets []*Tweet `json:"tweets"`
}
