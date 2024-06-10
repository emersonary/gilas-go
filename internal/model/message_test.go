package model

import (
	"testing"
	"time"

	"gitgub.com/emersonary/gilasw/go/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestNewMessage(t *testing.T) {
	categoryID := model.NewID()
	message, err := NewMessage(categoryID, "TESTE DE MSG")
	assert.Nil(t, err)
	assert.NotNil(t, message)
	assert.NotEmpty(t, message.ID)
	assert.NotEmpty(t, message.CategoryID)
	assert.Equal(t, categoryID, message.CategoryID)
	assert.Equal(t, "TESTE DE MSG", message.MessageText)
	assert.Less(t, time.Now().UnixNano()-message.CreatedAt.UnixNano(), time.Nanosecond*10)
}

func TestNewMessageEmptyName(t *testing.T) {
	user, err := NewMessage(model.NewID(), "")
	assert.Equal(t, err, ErrEmptyName)
	assert.Nil(t, user)

}

func TestNewMessageNameTooLong(t *testing.T) {
	user, err := NewMessage(model.NewID(), `012345678901234567890123456789012345678901234567891
	012345678901234567890123456789012345678901234567891
	012345678901234567890123456789012345678901234567891
	012345678901234567890123456789012345678901234567891
	012345678901234567890123456789012345678901234567891
	012345678901234567890123456789012345678901234567891
	012345678901234567890123456789012345678901234567891
	012345678901234567890123456789012345678901234567891
	012345678901234567890123456789012345678901234567891`)
	assert.Equal(t, err, ErrNameTooLong)
	assert.Nil(t, user)

}
