package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewChannel(t *testing.T) {

	user, err := NewCategory("Channel A")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Name)
	assert.Equal(t, "Channel A", user.Name)

}

func TestNewChannelEmptyName(t *testing.T) {
	user, err := NewChannel("")
	assert.Equal(t, err, ErrEmptyName)
	assert.Nil(t, user)

}

func TestNewChannelNameTooLong(t *testing.T) {
	user, err := NewChannel("012345678901234567890123456789012345678901234567891")
	assert.Equal(t, err, ErrNameTooLong)
	assert.Nil(t, user)

}
