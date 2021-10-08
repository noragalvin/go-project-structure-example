package repository

import (
	"ecommerce-integrations/internal/app/custom_types"
	"ecommerce-integrations/internal/app/models"
)

type ProductRepository interface {
	FindProductByID(uint, []string) (models.Product, error)
	GetProductsWithPaginate(custom_types.Paginate) ([]models.Product, error)
	BulkInsertProducts([]interface{}) ([]models.Product, error)
}
