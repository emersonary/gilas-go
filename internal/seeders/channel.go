package seeders

import (
	"gitgub.com/emersonary/gilasw/go/configs"
	"gitgub.com/emersonary/gilasw/go/global"
	"gitgub.com/emersonary/gilasw/go/infra/database"
	"gitgub.com/emersonary/gilasw/go/pkg/model"
	"gorm.io/gorm"
)

func LoadOrSeedChannels(db *gorm.DB) {

	channelDB := database.NewChannel(db)
	var err error
	global.Channels, err = channelDB.FindAll()
	if err != nil {
		panic(err)
	}

	if len(*global.Channels) < 3 {

		channels := configs.DefaultChannels()

		for _, channel := range *channels {

			channel.ID = model.NewID()
			if channelDB.Create(&channel) == nil {

				*global.Channels = append(*global.Channels, channel)

			}

		}

	}

}
