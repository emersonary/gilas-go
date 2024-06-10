package seeders

import (
	"gitgub.com/emersonary/gilasw/go/configs"
	"gitgub.com/emersonary/gilasw/go/global"
	"gitgub.com/emersonary/gilasw/go/infra/database"
	modelinternal "gitgub.com/emersonary/gilasw/go/internal/model"
	"gitgub.com/emersonary/gilasw/go/pkg/model"
	"gorm.io/gorm"
)

func SubscribedFromUserName(userName string) []modelinternal.Category {

	switch userName {
	case "User 1":
		return []modelinternal.Category{(*global.Categories)[0]}
	case "User 2":
		return []modelinternal.Category{(*global.Categories)[0], (*global.Categories)[1]}
	case "User 3":
		return []modelinternal.Category{(*global.Categories)[0], (*global.Categories)[1], (*global.Categories)[2]}
	default:
		return []modelinternal.Category{}
	}
}

func ChannelsFromUserName(userName string) []modelinternal.Channel {

	switch userName {
	case "User 3":
		return []modelinternal.Channel{(*global.Channels)[0]}
	case "User 2":
		return []modelinternal.Channel{(*global.Channels)[0], (*global.Channels)[1]}
	case "User 1":
		return []modelinternal.Channel{(*global.Channels)[0], (*global.Channels)[1], (*global.Channels)[2]}
	default:
		return []modelinternal.Channel{}
	}
}

func UserCompleteFromUser(user modelinternal.User) modelinternal.UserComplete {

	return modelinternal.UserComplete{
		User:       user,
		Subscribed: SubscribedFromUserName(user.Name),
		Channels:   ChannelsFromUserName(user.Name),
	}

}

func loadUsersComplete() {

	global.UsersComplete = &[]modelinternal.UserComplete{}

	for _, user := range *global.Users {

		*global.UsersComplete = append(*global.UsersComplete, UserCompleteFromUser(user))

	}

}

func LoadOrSeedUsers(db *gorm.DB) {

	userDB := database.NewUser(db)
	var err error
	global.Users, err = userDB.FindAll()
	if err != nil {
		panic(err)
	}

	if len(*global.Users) < 3 {

		users := configs.DefaultUsers()

		for _, user := range *users {

			user.ID = model.NewID()
			if userDB.Create(&user) == nil {

				*global.Users = append(*global.Users, user)

			}

		}

	}

	loadUsersComplete()

}
