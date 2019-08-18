package repository

import (
	"github.com/0g3/treasure-app/app/api/v1/model"
	"github.com/0g3/treasure-app/app/db"
	"github.com/jinzhu/gorm"
)

var User = newUserRepository(db.Conn)

type UserRepository interface {
	Get(id string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) Get(id string) (*model.User, error) {
	u := new(model.User)
	err := r.db.Where("id = ?", id).First(u).Error
	return u, err
}

func newUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
