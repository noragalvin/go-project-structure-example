package repository

import (
	"ecommerce-integrations/internal/app/models"
)

type ShopRepository interface {
	BulkInsertShops([]interface{}) ([]models.Shop, error)
}
