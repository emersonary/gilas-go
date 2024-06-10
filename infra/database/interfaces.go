package database

import (
	modelinternal "gitgub.com/emersonary/gilasw/go/internal/model"
)

type ModelInterface interface {
	FindIDByName(name string) interface{}
	FindAll() interface{}
}

type UserInterface interface {
	ModelInterface
	Create(user *modelinternal.User) error
}

type CategoryInterface interface {
	ModelInterface
	Create(user *modelinternal.Category) error
}

type ChannelInterface interface {
	ModelInterface
	Create(user *modelinternal.Channel) error
}
