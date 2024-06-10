package model

import "gitgub.com/emersonary/gilasw/go/pkg/model"

type Category struct {
	ID   model.ID `json:"id" gorm:"primaryKey"`
	Name string   `json:"name" gorm:"not null;uniqueIndex"`
}

func NewCategory(name string) (*Category, error) {

	if name == "" {
		return nil, ErrEmptyName
	}

	if len(name) > 50 {
		return nil, ErrNameTooLong
	}

	return &Category{
		ID:   model.NewID(),
		Name: name,
	}, nil
}
