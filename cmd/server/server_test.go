package main

import (
	"testing"

	"gitgub.com/emersonary/gilasw/go/infra/database"
	"gitgub.com/emersonary/gilasw/go/internal/model"
	"gitgub.com/emersonary/gilasw/go/internal/seeders"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateMessage(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	seeders.LoadOrSeed(db)

	categoryDB := database.NewCategory(db)
	category, err := categoryDB.FindByName("Sports")
	assert.NotNil(t, categoryDB)
	assert.Nil(t, err)

	message, err := model.NewMessage(category.ID, "TESTE DE MSG")
	assert.NotNil(t, message)
	assert.Nil(t, err)

	messageDB := database.NewMessage(db)
	messagereturn, err := messageDB.Create(message)
	assert.NotNil(t, messagereturn)
	assert.Nil(t, err)

	assert.EqualValues(t, 1, messageDB.MessageCount())
	assert.EqualValues(t, 6, messageDB.MessageNotificationCount())

}
