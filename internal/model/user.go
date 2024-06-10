package model

import "gitgub.com/emersonary/gilasw/go/pkg/model"

type User struct {
	ID    model.ID `json:"id" gorm:"primaryKey"`
	Name  string   `json:"name" gorm:"not null;uniqueIndex"`
	Email string   `json:"email" gorm:"not null"`
	Phone string   `json:"phone" gorm:"not null"`
}

type UserComplete = struct {
	User
	Subscribed []Category `json:"subscribed"`
	Channels   []Channel  `json:"channels"`
}

func NewUser(name, email, phone string) (*User, error) {

	if name == "" {
		return nil, ErrEmptyName
	}

	if len(name) > 50 {
		return nil, ErrNameTooLong
	}

	return &User{
		ID:    model.NewID(),
		Name:  name,
		Email: email,
		Phone: phone,
	}, nil
}
