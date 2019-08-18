package model

type Follow struct {
	UserID         string `json:"user_id"`
	FollowedUserID string `json:"followed_user_id"`
}
