package repository

import (
	"ecommerce-integrations/driver"
	"ecommerce-integrations/internal/app/models"

	"github.com/mitchellh/mapstructure"
)

type shopImpl struct {
	database *driver.Database
}

func NewShopRepository(db *driver.Database) ShopRepository {
	return &shopImpl{
		database: db,
	}
}
func (instance *shopImpl) BulkInsertShops(someArr []interface{}) ([]models.Shop, error) {
	shops := []models.Shop{}
	mapstructure.Decode(someArr, &shops)
	// err := instance.database.Conn.Create(&shops).Error
	// if err != nil {
	// 	return nil, err
	// }
	return shops, nil
}
