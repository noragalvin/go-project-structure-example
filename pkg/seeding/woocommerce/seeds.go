package woocommerce

import (
	"gorm.io/gorm"
)

type Seed struct {
	Name string
	Run  func(*gorm.DB) error
}

func All() []Seed {
	return []Seed{
		{
			Name: "CreateShop",
			Run: func(db *gorm.DB) error {
				return createShop(db)
			},
		},
	}
}

// func bulkInsert(db *gorm.DB, fakeData []interface{}, T any) error {
// 	err := db.Model(&T).Create(&fakeData).Error
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
