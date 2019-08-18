package repository

import (
	"github.com/0g3/treasure-app/app/api/v1/model"
	"github.com/0g3/treasure-app/app/db"
	"github.com/jinzhu/gorm"
)

var Follow = newFollowRepository(db.Conn)

type FollowRepository interface {
	Create(*model.Follow) error
	ListByUserID(userID string) ([]*model.Follow, error)
	ListByFollowedUserID(followedUserID string) ([]*model.Follow, error)
}

type followRepository struct {
	db *gorm.DB
}

func (r *followRepository) Create(f *model.Follow) error {
	return r.db.Create(f).Error
}

func (r *followRepository) ListByUserID(userID string) ([]*model.Follow, error) {
	followers := make([]*model.Follow, 0)
	err := r.db.Where("user_id = ?", userID).Find(&followers).Error
	return followers, err
}

func (r *followRepository) ListByFollowedUserID(followedUserID string) ([]*model.Follow, error) {
	followers := make([]*model.Follow, 0)
	err := r.db.Where("followed_user_id = ?", followedUserID).Find(&followers).Error
	return followers, err
}

func newFollowRepository(db *gorm.DB) FollowRepository {
	return &followRepository{
		db: db,
	}
}
