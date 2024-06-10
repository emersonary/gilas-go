package database

import (
	"gitgub.com/emersonary/gilasw/go/internal/model"
	"gorm.io/gorm"
)

type Channel struct {
	DB *gorm.DB
}

func NewChannel(db *gorm.DB) *Channel {
	return &Channel{DB: db}
}

func (c *Channel) Create(channel *model.Channel) error {
	return c.DB.Create(channel).Error
}

func (c *Channel) FindByName(name string) (*model.Channel, error) {
	var channel model.Channel

	if err := c.DB.Where("name = $1", name).First(&channel).Error; err != nil {
		return nil, err
	}

	return &channel, nil

}

func (c *Channel) FindAll() (*[]model.Channel, error) {

	var channels *[]model.Channel

	c.DB.Find(&channels)

	return channels, nil

}
