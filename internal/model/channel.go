package model

import "gitgub.com/emersonary/gilasw/go/pkg/model"

type Channel struct {
	ID   model.ID `json:"id" gorm:"primaryKey"`
	Name string   `json:"name" gorm:"not null;uniqueIndex"`
}

func NewChannel(name string) (*Channel, error) {

	if name == "" {
		return nil, ErrEmptyName
	}

	if len(name) > 50 {
		return nil, ErrNameTooLong
	}

	return &Channel{
		ID:   model.NewID(),
		Name: name,
	}, nil
}
