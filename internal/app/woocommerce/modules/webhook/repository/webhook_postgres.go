package repository

import (
	"ecommerce-integrations/driver"
	"ecommerce-integrations/internal/app/models"

	"github.com/mitchellh/mapstructure"
)

type webhookImpl struct {
	database *driver.Database
}

func NewWebhookRepository(db *driver.Database) WebhookRepository {
	return &webhookImpl{
		database: db,
	}
}
func (instance *webhookImpl) BulkInsertWebhooks(someArr []interface{}) ([]models.Webhook, error) {
	webhooks := []models.Webhook{}
	mapstructure.Decode(someArr, &webhooks)
	// err := instance.database.Conn.Create(&webhooks).Error
	// if err != nil {
	// 	return nil, err
	// }
	return webhooks, nil
}
