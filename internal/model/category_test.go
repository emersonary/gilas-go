package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCategory(t *testing.T) {
	user, err := NewCategory("Category A")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Name)
	assert.Equal(t, "Category A", user.Name)

}

func TestNewCategoryEmptyName(t *testing.T) {
	user, err := NewCategory("")
	assert.Equal(t, err, ErrEmptyName)
	assert.Nil(t, user)

}

func TestNewCategoryNameTooLong(t *testing.T) {
	user, err := NewCategory("012345678901234567890123456789012345678901234567891")
	assert.Equal(t, err, ErrNameTooLong)
	assert.Nil(t, user)

}
