package configs

import (
	"gitgub.com/emersonary/gilasw/go/internal/model"
)

func DefaultCategories() *[]model.Category {

	return &[]model.Category{{Name: "Sports"},
		{Name: "Finance"},
		{Name: "Movies"}}
}

func DefaultChannels() *[]model.Channel {

	return &[]model.Channel{{Name: "SMS"},
		{Name: "E-mail"},
		{Name: "Push Notification"}}
}

func DefaultUsers() *[]model.User {

	return &[]model.User{{
		Name:  "User 1",
		Email: "user1@fakemail.com",
		Phone: "551199999991",
	},
		{
			Name:  "User 2",
			Email: "user2@fakemail.com",
			Phone: "551199999992",
		},
		{
			Name:  "User 3",
			Email: "user3@fakemail.com",
			Phone: "551199999993",
		}}
}
