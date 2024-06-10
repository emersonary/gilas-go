package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "j@j.com", "5511943249844")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Name)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "j@j.com", user.Email)
}

func TestNewUserEmptyName(t *testing.T) {
	user, err := NewUser("", "1", "1")
	assert.Equal(t, err, ErrEmptyName)
	assert.Nil(t, user)

}

func TestNewUserNameTooLong(t *testing.T) {
	user, err := NewUser("012345678901234567890123456789012345678901234567891", "1", "1")
	assert.Equal(t, err, ErrNameTooLong)
	assert.Nil(t, user)

}
