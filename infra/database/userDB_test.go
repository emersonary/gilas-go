package database

import (
	"testing"

	"gitgub.com/emersonary/gilasw/go/internal/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&model.User{})

	user, err := model.NewUser("Emerson", "senior@emersonary.dev", "+5511943249844")
	assert.NotNil(t, user)
	assert.Nil(t, err)

	userDB := NewUser(db)
	err = userDB.Create(user)
	assert.Nil(t, err)

	seluser, err := userDB.FindByName("Emerson")

	assert.NotNil(t, seluser)
	assert.Nil(t, err)

}
