package repository

import (
	"ecommerce-integrations/internal/app/models"
)

type WebhookRepository interface {
	BulkInsertWebhooks([]interface{}) ([]models.Webhook, error)
}
