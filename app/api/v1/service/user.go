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

	// TODO: refactoring
	followList, err := repository.Follow.ListByUserID(userID)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return nil, err
	}
	follows := make([]*model.User, 0, len(followList))
	for _, f := range followList {
		follow, err := repository.User.Get(f.FollowedUserID)
		if err != nil {
			return nil, err
		}
		follows = append(follows, follow)
	}

	// TODO: refactoring
	followedList, err := repository.Follow.ListByFollowedUserID(userID)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return nil, err
	}
	followers := make([]*model.User, 0, len(followedList))
	for _, f := range followedList {
		follower, err := repository.User.Get(f.UserID)
		if err != nil {
			return nil, err
		}
		followers = append(followers, follower)
	}

	return &model.BatchGetUserResponse{
		User:      u,
		Tweets:    tweets,
		Follows:   follows,
		Followers: followers,
	}, err
}
