package seeders

import (
	"gitgub.com/emersonary/gilasw/go/internal/model"
	"gorm.io/gorm"
)

func createviewsifnotexist(db *gorm.DB) {

	db.Exec(`create view if not exists v_message_notifications as
	select
	 message_notifications.id,
	 message_notifications.notification_id,
	 users.name UserName,
	 categories.name CategoryName,
	 channels.name ChannelName,
	 message_notifications.message_text,
	 message_notifications.created_at
	from
	 users,
	 categories,
	 channels,
	 message_notifications
	where message_notifications.user_id = users.id
	 and message_notifications.category_id = categories.id
	 and message_notifications.channel_id = channels.id ;
`)

	db.Exec(`create view if not exists v_messages as
select 
 messages.id,
 categories.name CategoryName,
 messages.message_text,
 messages.created_at
from 
categories,
messages
where messages.category_id = categories.id;`)

}

func migrate(db *gorm.DB) {

	db.AutoMigrate(
		&model.Message{},
		&model.MessageNotification{},
		&model.User{},
		&model.Category{},
		&model.Channel{})

	createviewsifnotexist(db)

}

func LoadOrSeed(db *gorm.DB) {

	migrate(db)
	LoadOrSeedCategories(db)
	LoadOrSeedChannels(db)
	LoadOrSeedUsers(db)

}
