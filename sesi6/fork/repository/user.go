package repository

import (
	"log"

	"fork/model"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func (u *UserRepo) Migrate() {
	err := u.DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("Error migrating user: ", err)
	}
}

func (u *UserRepo) Create(user *model.User) error {
	err := u.DB.Debug().Model(&model.User{}).Create(user).Error
	return err
}

func (u *UserRepo) Get() ([]*model.User, error) {
	users := []*model.User{}
	err := u.DB.Debug().Model(&model.User{}).Find(&users).Error
	return users, err
}
