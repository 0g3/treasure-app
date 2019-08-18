package repository

import (
	"github.com/0g3/treasure-app/app/api/v1/model"
	"github.com/0g3/treasure-app/app/db"
	"github.com/jinzhu/gorm"
)

var Follow = newFollowRepository(db.Conn)

type FollowRepository interface {
	Create(*model.Follow) error
}

type followRepository struct {
	db *gorm.DB
}

func (r *followRepository) Create(f *model.Follow) error {
	return r.db.Create(f).Error
}

func newFollowRepository(db *gorm.DB) FollowRepository {
	return &followRepository{
		db: db,
	}
}
