package repository

import (
	"github.com/0g3/treasure-app/app/api/v1/model"
	"github.com/0g3/treasure-app/app/db"
	"github.com/jinzhu/gorm"
)

var Tweet = newTweetRepository(db.Conn)

type TweetRepository interface {
	Create(*model.Tweet) error
	ListByUserID(userID string) ([]*model.Tweet, error)
}

type tweetRepository struct {
	db *gorm.DB
}

func (r *tweetRepository) Create(t *model.Tweet) error {
	return r.db.Create(t).Error
}

func (r *tweetRepository) ListByUserID(userID string) ([]*model.Tweet, error) {
	users := make([]*model.Tweet, 0)
	err := r.db.Where("user_id = ?", userID).Find(&users).Error
	return users, err
}

func newTweetRepository(db *gorm.DB) TweetRepository {
	return &tweetRepository{
		db: db,
	}
}
