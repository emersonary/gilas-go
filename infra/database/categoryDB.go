package database

import (
	"gitgub.com/emersonary/gilasw/go/internal/model"
	pkgmodel "gitgub.com/emersonary/gilasw/go/pkg/model"
	"gorm.io/gorm"
)

type Category struct {
	DB *gorm.DB
}

func NewCategory(db *gorm.DB) *Category {
	return &Category{DB: db}
}

func (c *Category) Create(category *model.Category) error {
	return c.DB.Create(category).Error
}

func (c *Category) FindByID(id pkgmodel.ID) (*model.Category, error) {
	var category model.Category

	if err := c.DB.Where("id = $1", id).First(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil

}

func (c *Category) FindByName(name string) (*model.Category, error) {
	var category model.Category

	if err := c.DB.Where("name = $1", name).First(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil

}

func (c *Category) Store(name string) (*model.Category, error) {
	var category model.Category

	if err := c.DB.Where("name = $1", name).First(&category).Error; err != nil {
		panic(err)
	}

	return &category, nil

}

func (c *Category) FindAll() (*[]model.Category, error) {

	var categories *[]model.Category

	c.DB.Find(&categories)

	return categories, nil

}
