package repository

import (
	"ecommerce-integrations/driver"
	"ecommerce-integrations/internal/app/custom_types"
	"ecommerce-integrations/internal/app/models"

	"github.com/mitchellh/mapstructure"
)

type productImpl struct {
	database *driver.Database
}

func NewProductRepository(db *driver.Database) ProductRepository {
	return &productImpl{
		database: db,
	}
}

func (instance *productImpl) FindProductByID(productID uint, fields []string) (models.Product, error) {
	product := models.Product{}
	err := instance.database.Conn.First(&product).Error
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (instance *productImpl) GetProductsWithPaginate(paginate custom_types.Paginate) ([]models.Product, error) {
	products := []models.Product{}
	err := instance.database.Conn.Offset(paginate.Offset).Limit(paginate.Limit).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, err
}

func (instance *productImpl) BulkInsertProducts(someArr []interface{}) ([]models.Product, error) {
	products := []models.Product{}
	mapstructure.Decode(someArr, &products)
	err := instance.database.Conn.Create(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
