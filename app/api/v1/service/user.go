package service

import (
	"github.com/0g3/treasure-app/app/api/v1/model"
	"github.com/0g3/treasure-app/app/api/v1/repository"
	"github.com/jinzhu/gorm"
)

func BatchGetUser(userID string) (*model.BatchGetUserResponse, error) {
	u, err := repository.User.Get(userID)
	if err != nil {
		return nil, err
	}

	tweets, err := repository.Tweet.ListByUserID(userID)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return nil, err
	}

	return &model.BatchGetUserResponse{
		User:   u,
		Tweets: tweets,
	}, err
}
