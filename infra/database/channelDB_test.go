package database

import (
	"testing"

	"gitgub.com/emersonary/gilasw/go/internal/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateChannel(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&model.Channel{})

	channel, err := model.NewChannel("Channel A")
	assert.NotNil(t, channel)
	assert.Nil(t, err)

	channelDB := NewChannel(db)
	err = channelDB.Create(channel)
	assert.Nil(t, err)

	seluser, err := channelDB.FindByName("Channel A")

	assert.NotNil(t, seluser)
	assert.Nil(t, err)

}
