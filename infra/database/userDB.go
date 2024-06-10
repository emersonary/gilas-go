package database

import (
	"gitgub.com/emersonary/gilasw/go/internal/model"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}

func (u *User) Create(user *model.User) error {
	return u.DB.Create(user).Error
}

func (u *User) FindByName(name string) (*model.User, error) {
	var user model.User

	if err := u.DB.Where("name = $1", name).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil

}

func (u *User) FindAll() (*[]model.User, error) {

	var users *[]model.User

	u.DB.Find(&users)

	return users, nil

}
