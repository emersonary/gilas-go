package database

import (
	"testing"

	"gitgub.com/emersonary/gilasw/go/internal/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateCategory(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&model.Category{})

	category, err := model.NewCategory("Category A")
	assert.NotNil(t, category)
	assert.Nil(t, err)

	categoryDB := NewCategory(db)
	err = categoryDB.Create(category)
	assert.Nil(t, err)

	seluser, err := categoryDB.FindByName("Category A")

	assert.NotNil(t, seluser)
	assert.Nil(t, err)

}
