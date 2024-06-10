package seeders

import (
	"gitgub.com/emersonary/gilasw/go/configs"
	"gitgub.com/emersonary/gilasw/go/global"
	"gitgub.com/emersonary/gilasw/go/infra/database"
	"gitgub.com/emersonary/gilasw/go/pkg/model"
	"gorm.io/gorm"
)

func LoadOrSeedCategories(db *gorm.DB) {

	categoryDB := database.NewCategory(db)
	var err error
	global.Categories, err = categoryDB.FindAll()
	if err != nil {
		panic(err)
	}

	if len(*global.Categories) < 3 {

		categories := configs.DefaultCategories()

		for _, category := range *categories {

			category.ID = model.NewID()
			if categoryDB.Create(&category) == nil {

				*global.Categories = append(*global.Categories, category)

			}

		}

	}

}
