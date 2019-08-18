package repository

import (
	"github.com/0g3/treasure-app/app/api/v1/model"
	"github.com/0g3/treasure-app/app/db"
	"github.com/jinzhu/gorm"
)

var Tweet = newTweetRepository(db.Conn)

type TweetRepository interface {
	Create(*model.Tweet) error
}

type tweetRepository struct {
	db *gorm.DB
}

func (r *tweetRepository) Create(t *model.Tweet) error {
	return r.db.Create(t).Error
}

func newTweetRepository(db *gorm.DB) TweetRepository {
	return &tweetRepository{
		db: db,
	}
}
