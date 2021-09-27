package repository

import (
	"go-project-structure-example/driver"
	"go-project-structure-example/internal/app/custom_types"
	"go-project-structure-example/internal/app/models"

	"github.com/mitchellh/mapstructure"
)

type ProductRepository interface {
	FindProductByID(uint, []string) (models.Product, error)
	GetProductsWithPaginate(custom_types.Paginate) ([]models.Product, error)
	BulkInsertProducts([]interface{}) ([]models.Product, error)
}

type productImpl struct {
	database *driver.Database
}

func NewProductRepository(db *driver.Database) ProductRepository {
	return &productImpl{
		database: db,
	}
}

func (this *productImpl) FindProductByID(productID uint, fields []string) (models.Product, error) {
	product := models.Product{}
	err := this.database.Conn.First(&product).Error
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (this *productImpl) GetProductsWithPaginate(paginate custom_types.Paginate) ([]models.Product, error) {
	products := []models.Product{}
	err := this.database.Conn.Offset(paginate.Offset).Limit(paginate.Limit).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, err
}

func (this *productImpl) BulkInsertProducts(someArr []interface{}) ([]models.Product, error) {
	products := []models.Product{}
	mapstructure.Decode(someArr, &products)
	err := this.database.Conn.Create(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
