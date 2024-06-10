package model

import (
	"time"

	"gitgub.com/emersonary/gilasw/go/pkg/model"
)

type Message struct {
	ID          model.ID  `json:"id" gorm:"primaryKey"`
	Category    Category  `json:"category" gorm:"primaryKey"`
	CategoryID  model.ID  `json:"categoryid" gorm:"foreignKey"`
	MessageText string    `json:"messagetext" gorm:"not null"`
	CreatedAt   time.Time `json:"createdat"`
}

type MessageView struct {
	ID           model.ID  `json:"id"`
	CategoryName string    `json:"category"`
	MessageText  string    `json:"messagetext"`
	CreatedAt    time.Time `json:"createdat"`
}

func NewMessage(categoryID model.ID, messageText string) (*Message, error) {

	if messageText == "" {
		return nil, ErrEmptyName
	}

	if len(messageText) > 300 {
		return nil, ErrNameTooLong
	}

	return &Message{
		ID:          model.NewID(),
		CategoryID:  categoryID,
		MessageText: messageText,
		CreatedAt:   time.Now(),
	}, nil
}

type MessageNotification struct {
	Message
	NotificationID model.ID `json:"id" gorm:"primaryKey"`
	Channel        Channel
	ChannelID      model.ID `json:"channelid" gorm:"foreignKey"`
	User           User
	UserID         model.ID `json:"categoryid" gorm:"foreignKey"`
}

type MessageNotificationView struct {
	ID             model.ID  `json:"id"`
	NotificationID model.ID  `json:"notificationid"`
	UserName       string    `json:"user"`
	CategoryName   string    `json:"category"`
	ChannelName    string    `json:"channel"`
	MessageText    string    `json:"messagetext"`
	CreatedAt      time.Time `json:"createdat"`
}

func NewMessageNotification(userID, channelID model.ID, message *Message) (*MessageNotification, error) {
	return &MessageNotification{
		Message:        *message,
		NotificationID: model.NewID(),
		UserID:         userID,
		ChannelID:      channelID,
	}, nil
}
