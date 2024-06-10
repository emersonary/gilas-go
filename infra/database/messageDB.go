package database

import (
	"strconv"

	"gitgub.com/emersonary/gilasw/go/global"
	"gitgub.com/emersonary/gilasw/go/internal/model"
	"gorm.io/gorm"
)

type Message struct {
	DB *gorm.DB
}

func NewMessage(db *gorm.DB) *Message {
	return &Message{DB: db}
}

func (c *Message) createNotification(user model.UserComplete, message *model.Message) error {

	for _, channel := range user.Channels {

		messagenotification, err := model.NewMessageNotification(user.ID, channel.ID, message)
		if err != nil {
			return err
		}

		err = c.DB.Create(messagenotification).Error
		if err != nil {
			return err
		}

	}

	return nil

}

func (c *Message) createNotifications(message *model.Message) (*model.Message, error) {

	for _, user := range *global.UsersComplete {

		for _, category := range user.Subscribed {

			if category.ID == message.CategoryID {

				err := c.createNotification(user, message)
				if err != nil {
					return nil, err
				}

			}

		}
	}

	return message, nil
}

func (c *Message) Create(message *model.Message) (*model.Message, error) {
	err := c.DB.Create(message).Error
	if err != nil {
		return nil, err
	}
	return c.createNotifications(message)
}

func (c *Message) GetMessagesLastRows(limit int) (*[]model.MessageView, error) {

	var messages []model.MessageView
	err := c.DB.Raw("select * from v_Messages order by created_at desc limit " + strconv.Itoa(limit)).Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return &messages, nil
}

func (c *Message) GetMessagesNotificationsLastRows(limit int) (*[]model.MessageNotificationView, error) {

	var messages []model.MessageNotificationView
	err := c.DB.Raw("select * from v_Message_Notifications order by created_at desc limit " + strconv.Itoa(limit)).Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return &messages, nil
}

func (c *Message) MessageCount() int64 {
	var count int64
	c.DB.Model(&model.Message{}).Count(&count)
	return count
}

func (c *Message) MessageNotificationCount() int64 {
	var count int64
	c.DB.Model([]model.MessageNotification{}).Count(&count)
	return count
}
